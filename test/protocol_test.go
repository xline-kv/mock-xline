package test

import (
	"context"
	"fmt"
	"net"
	"testing"
	// "time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	"github.com/xline-kv/mock-xline/server"
	"github.com/xline-kv/mock-xline/client"
)

type ProtocolServerTestSuite struct {
	suite.Suite
}

func (suite *ProtocolServerTestSuite) SetupTest() {
	go func() {
		fmt.Println("set up")
		lis, err := net.Listen("tcp", "127.0.0.1:5555")
		if err != nil {
			fmt.Printf("failed to listen: %v\n", err)
		}
		s := grpc.NewServer()
		curppb.RegisterProtocolServer(s, &server.ProtocolServer{})
		//...
		fmt.Printf("server listening at %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v\n", err)
		}
	}()
}

func (suite *ProtocolServerTestSuite) TestTemp() {
	//  time.Sleep(1 * time.Second)
	cli, err := client.NewProtocolClient()
	if err != nil {
		panic(err)
	}

	// request
	token := ""
	cmd := &xlinepb.Command{
		Keys: []*xlinepb.KeyRange{
			{
				Key:      []byte("key"),
				RangeEnd: []byte("key"),
			},
		},
		Request: &xlinepb.RequestWithToken{
			Token: &token,
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("key"),
					Value: []byte("value"),
				},
			},
		},
		ProposeId: &xlinepb.ProposeId{
			ClientId: 1,
			SeqNum:   1,
		},
	}
	bcmd, err := proto.Marshal(cmd)
	if err != nil {
		panic(err)
	}
	req := &curppb.ProposeRequest{
		Command:        bcmd,
		ClusterVersion: 1,
	}


	res, err := cli.Propose(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("final", res)
}

func TestProtocolServer(t *testing.T) {
	suite.Run(t, new(ProtocolServerTestSuite))
}
