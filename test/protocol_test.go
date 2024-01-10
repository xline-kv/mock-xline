package test

import (
	"context"
	"testing"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type ProtocolServerTestSuite struct {
	suite.Suite

	cli curppb.ProtocolClient
}

func (suite *ProtocolServerTestSuite) SetupSuite() {
	addr := "127.0.0.1:5555"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	suite.NoError(err)

	suite.cli = curppb.NewProtocolClient(conn)
}

func (suite *ProtocolServerTestSuite) TestFetchCluster() {
	res, err := suite.cli.FetchCluster(context.TODO(), &curppb.FetchClusterRequest{})
	suite.NoError(err)
	suite.IsType(&curppb.FetchClusterResponse{}, res)
}

func (suite *ProtocolServerTestSuite) TestPut() {
	cmd := &xlinepb.Command{
		Request: &xlinepb.RequestWithToken{
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("key"),
					Value: []byte("val"),
				},
			},
		},
	}
	pid := &curppb.ProposeId{
		ClientId: 1357,
		SeqNum:   2468,
	}

	bcmd, err := proto.Marshal(cmd)
	suite.NoError(err)

	req := &curppb.ProposeRequest{
		ProposeId: pid,
		Command:   bcmd,
	}

	ppsRes, err := suite.cli.Propose(context.TODO(), req)
	suite.NoError(err)
	suite.IsType(&curppb.ProposeResponse{}, ppsRes)

	wsRes, err := suite.cli.WaitSynced(context.TODO(), &curppb.WaitSyncedRequest{ProposeId: pid})
	suite.NoError(err)
	suite.IsType(&curppb.WaitSyncedResponse{}, wsRes)
}

func TestProtocolServerTestSuite(t *testing.T) {
	suite.Run(t, new(ProtocolServerTestSuite))
}
