package data

import (
	"encoding/json"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"

	// "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func NewFetchClusterDataMap() map[Request]*Response[curppb.FetchClusterResponse] {
	// fetch cluster request
	fetchClusterRequest := &curppb.FetchClusterRequest{}
	byteFetchClusterRequest, err := json.Marshal(fetchClusterRequest)
	if err != nil {
		panic(err)
	}
	stringFetchClusterRequest := string(byteFetchClusterRequest)

	// fetch cluster response 1
	var leaderId uint64 = 1
	res1 := &curppb.FetchClusterResponse{
		LeaderId:  &leaderId,
		Term:      1,
		ClusterId: 1,
		Members: []*curppb.Member{
			{
				Id:    1,
				Name:  "node1",
				Addrs: []string{"127.0.0.1:5555"},
			},
		},
		ClusterVersion: 0,
	}
	res2 := &curppb.FetchClusterResponse{
		LeaderId:  &leaderId,
		Term:      1,
		ClusterId: 1,
		Members: []*curppb.Member{
			{
				Id:    1,
				Name:  "node1",
				Addrs: []string{"127.0.0.1:5555"},
			},
		},
		ClusterVersion: 1,
	}
	response1 := Response[curppb.FetchClusterResponse]{
		Index: 0,
		Res: []ResponseWrapper[curppb.FetchClusterResponse]{
			{
				Res: *res1,
			},
			{
				Res: *res2,
			},
		},
	}

	return map[Request]*Response[curppb.FetchClusterResponse]{
		stringFetchClusterRequest: &response1,
	}
}

func NewProposeDataMap() map[Request]*Response[curppb.ProposeResponse] {
	var leaser_id uint64 = 1
	er := &xlinepb.CommandResponse{
		ResponseWrapper: &xlinepb.CommandResponse_PutResponse{
			PutResponse: &xlinepb.PutResponse{
				Header: &xlinepb.ResponseHeader{
					ClusterId: 1,
					MemberId: 1,
				},
			},
		},
	}
	ber, err := proto.Marshal(er)
	if err != nil {
		panic(ber)
	}
	res := &curppb.ProposeResponse{
		LeaderId: &leaser_id,
		Term:     1,
		ExeResult: &curppb.ProposeResponse_Result{
			Result: &curppb.CmdResult{
				Result: &curppb.CmdResult_Ok{
					Ok: ber,
				},
			},
		},
	}

	return map[Request]*Response[curppb.ProposeResponse]{
		"default": {
			Index: 0,
			Res: []ResponseWrapper[curppb.ProposeResponse]{
				{
					Err: status.Errorf(codes.FailedPrecondition, "wrong cluster version"),
				},
				{
					Res: *res,
				},
			},
		},
	}
}

func NewWaitSyncedDataMap() map[Request]*Response[curppb.WaitSyncedResponse] {
	return map[Request]*Response[curppb.WaitSyncedResponse]{
		"default": {
			Index: 0,
			Res: []ResponseWrapper[curppb.WaitSyncedResponse]{
				{
					Err: status.Errorf(codes.FailedPrecondition, "wrong cluster version"),
				},
			},
		},
	}
}
