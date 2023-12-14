package data

import (
	// "encoding/json"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	// xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	// "google.golang.org/protobuf/proto"
)

type Request = string

type ResponseWrapper[T curppb.FetchClusterResponse | curppb.ProposeResponse | curppb.WaitSyncedResponse] struct {
	Res T
	Err error
}

type Response[T curppb.FetchClusterResponse | curppb.ProposeResponse | curppb.WaitSyncedResponse] struct {
	Index uint
	Res   []ResponseWrapper[T]
}

// func NewResponse(res interface{}, err error) Response {
// 	return Response{Res: res, Err: err}
// }

// func NewProposeRequest(cmd *xlinepb.Command, clusterVer uint64) (*Request, error) {
// 	bcmd, err := proto.Marshal(cmd)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req := &curppb.ProposeRequest{
// 		Command:        bcmd,
// 		ClusterVersion: clusterVer,
// 	}
// 	breq, err := json.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sreq := string(breq)
// 	return &sreq, nil
// }
