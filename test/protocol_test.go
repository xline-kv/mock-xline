package test_e2e

import (
	"context"
	"testing"

	curppb "github.com/xline-kv/mock-xline/api/gen/curppb"
	"github.com/xline-kv/mock-xline/api/gen/mockpb"
	xlinepb "github.com/xline-kv/mock-xline/api/gen/xlinepb"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type ProtocolServerTestSuite struct {
	suite.Suite

	// mock client
	mockCli mockpb.MockClient
	// protocol client
	protocolCli curppb.ProtocolClient

	// fetch cluster request
	fcreq *curppb.FetchClusterRequest
	// fetch cluster response
	fcres *curppb.FetchClusterResponse
	// serialized command
	bcmd []byte
	// propose response
	pres *curppb.ProposeResponse
	// wait synced response
	wsres *curppb.WaitSyncedResponse
}

func (suite *ProtocolServerTestSuite) SetupTest() {
	addr := "127.0.0.1:5555"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	suite.NoError(err)

	// build clients
	suite.mockCli = mockpb.NewMockClient(conn)
	suite.protocolCli = curppb.NewProtocolClient(conn)

	// build fetch cluster request
	suite.fcreq = &curppb.FetchClusterRequest{}

	// build fetch cluster response
	var leaderId uint64 = 123456
	suite.fcres = &curppb.FetchClusterResponse{
		LeaderId:  &leaderId,
		Term:      1,
		ClusterId: 123456,
		Members: []*curppb.Member{
			{
				Id:    123456,
				Name:  "node1",
				Addrs: []string{"127.0.0.1:5555"},
			},
		},
		ClusterVersion: 1,
	}

	// build command
	cmd := xlinepb.Command{
		Keys: []*xlinepb.KeyRange{
			{
				Key:      []byte("put"),
				RangeEnd: []byte("put"),
			},
		},
		Request: &xlinepb.RequestWithToken{
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("put"),
					Value: []byte("val"),
				},
			},
		},
	}
	bcmd, err := proto.Marshal(&cmd)
	suite.NoError(err)
	suite.bcmd = bcmd

	// build propose response
	er := xlinepb.CommandResponse{
		ResponseWrapper: &xlinepb.CommandResponse_PutResponse{
			PutResponse: &xlinepb.PutResponse{
				Header: &xlinepb.ResponseHeader{
					ClusterId: 1,
					MemberId:  123456,
					Revision:  1,
				},
			},
		},
	}
	ber, err := proto.Marshal(&er)
	suite.NoError(err)

	suite.pres = &curppb.ProposeResponse{
		Result: &curppb.CmdResult{
			Result: &curppb.CmdResult_Ok{
				Ok: ber,
			},
		},
	}

	// build wait synced response
	suite.wsres = &curppb.WaitSyncedResponse{
		AfterSyncResult: &curppb.CmdResult{
			Result: &curppb.CmdResult_Ok{
				Ok: ber,
			},
		},
		ExeResult: &curppb.CmdResult{
			Result: &curppb.CmdResult_Ok{
				Ok: ber,
			},
		},
	}
}

func (suite *ProtocolServerTestSuite) TestFetchCluster() {
	bfcreq, err := proto.Marshal(suite.fcreq)
	suite.NoError(err)
	bfcres, err := proto.Marshal(suite.fcres)
	suite.NoError(err)

	_, err = suite.mockCli.MockFetchCluster(context.TODO(), &mockpb.MockRequest{
		Request: bfcreq,
		Response: []*mockpb.Response{
			{
				Response: &mockpb.Response_Success{
					Success: bfcres,
				},
			},
		},
	})
	suite.NoError(err)

	res, err := suite.protocolCli.FetchCluster(context.TODO(), suite.fcreq)
	suite.NoError(err)
	suite.Equal(suite.fcres.String(), res.String())
}

func (suite *ProtocolServerTestSuite) TestPropose() {
	bpres, err := proto.Marshal(suite.pres)
	suite.NoError(err)

	_, err = suite.mockCli.MockPropose(context.TODO(), &mockpb.MockRequest{
		Request: suite.bcmd,
		Response: []*mockpb.Response{
			{
				Response: &mockpb.Response_Success{
					Success: bpres,
				},
			},
		},
	})
	suite.NoError(err)

	preq := &curppb.ProposeRequest{ProposeId: &curppb.ProposeId{ClientId: 1, SeqNum: 123456}, Command: suite.bcmd, ClusterVersion: 1}
	res, err := suite.protocolCli.Propose(context.TODO(), preq)
	suite.NoError(err)
	suite.Equal(suite.pres.String(), res.String())
}

func (suite *ProtocolServerTestSuite) TestWaitSynced() {
	bwsres, err := proto.Marshal(suite.wsres)
	suite.NoError(err)

	_, err = suite.mockCli.MockWaitSynced(context.TODO(), &mockpb.MockRequest{
		Request: suite.bcmd,
		Response: []*mockpb.Response{
			{
				Response: &mockpb.Response_Success{
					Success: bwsres,
				},
			},
		},
	})
	suite.NoError(err)

	wsreq := &curppb.WaitSyncedRequest{ProposeId: &curppb.ProposeId{ClientId: 1, SeqNum: 123456}, ClusterVersion: 1}
	res, err := suite.protocolCli.WaitSynced(context.TODO(), wsreq)
	suite.NoError(err)
	suite.Equal(suite.wsres.String(), res.String())
}

func TestProtocolServerTestSuite(t *testing.T) {
	suite.Run(t, new(ProtocolServerTestSuite))
}
