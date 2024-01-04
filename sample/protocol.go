package sample

import (
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	"google.golang.org/protobuf/proto"
)

type testName = string
type serverId = uint64

type fetchClusterRequest = string
type command = string
type proposeId = string

type FetchClusterMap = map[fetchClusterRequest]*curppb.FetchClusterResponse
type ProposeMap = map[command]*curppb.ProposeResponse
type CommandMap = map[proposeId]*xlinepb.Command
type WaitSyncedMap = map[command]*curppb.WaitSyncedResponse

func NewFetchClusterMap(name testName) FetchClusterMap {
	req := &curppb.FetchClusterRequest{}

	var leaderId serverId = 1
	res := &curppb.FetchClusterResponse{
		LeaderId:  &leaderId,
		Term:      1,
		ClusterId: 1,
		Members: []*curppb.Member{
			{
				Id:    1,
				Name:  "node1",
				Addrs: []string{"127.0.0.1:2379"},
			},
		},
		ClusterVersion: 1,
	}

	switch name {
	case "protocol_client":
		return FetchClusterMap{
			req.String(): res,
		}
	default:
		panic("unknown test case")
	}
}

func NewProposeMap(name testName) ProposeMap {
	cmd := &xlinepb.Command{
		Request: &xlinepb.RequestWithToken{
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("test"),
					Value: []byte("protocol_client"),
				},
			},
		},
	}

	er := &xlinepb.PutResponse{}
	ber, err := proto.Marshal(er)
	if err != nil {
		panic(err)
	}
	res := &curppb.ProposeResponse{
		Result: &curppb.CmdResult{
			Result: &curppb.CmdResult_Ok{
				Ok: ber,
			},
		},
	}

	switch name {
	case "protocol_client":
		return ProposeMap{
			cmd.String(): res,
		}
	default:
		panic("unknown test case")
	}
}

func NewWaitSyncedMap(name testName) WaitSyncedMap {
	cmd := &xlinepb.Command{
		Request: &xlinepb.RequestWithToken{
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("test"),
					Value: []byte("protocol_client"),
				},
			},
		},
	}

	er := &xlinepb.PutResponse{}
	ber, err := proto.Marshal(er)
	if err != nil {
		panic(ber)
	}

	res := &curppb.WaitSyncedResponse{
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

	switch name {
	case "protocol_client":
		return WaitSyncedMap{
			cmd.String(): res,
		}
	default:
		panic("unknown test case")
	}
}
