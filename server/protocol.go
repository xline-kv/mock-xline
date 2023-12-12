package server

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/xline-kv/mock-xline/data"
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
)

type ProtocolServer struct {
	curppb.ProtocolServer

	dataMap  map[*curppb.ProposeRequest]*curppb.ProposeResponse
}

func (s *ProtocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	m := data.NewProtocolDataMap()

	breq, err := sonic.Marshal(req)
	if err != nil {
		panic(err)
	}

	res, ok := m[string(breq)]
	if !ok {
		panic("map error")
	}

	if res.Err != nil {
		return nil, res.Err
	}
	return res.Res.(*curppb.ProposeResponse), nil
}

// func (s *ProtocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
// 	var leaderId uint64 = 1
// 	res := &curppb.FetchClusterResponse{
// 		LeaderId: &leaderId,
// 		Term: 1,
// 		ClusterId: 1,
// 		Members: []*curppb.Member{
// 			{
// 				Id: 1,
// 				Name: "node1",
// 				Addrs: []string{"127.0.0.1:5555"},
// 			},
// 		},
// 	}
// 	return res, nil
// }
