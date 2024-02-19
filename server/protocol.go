package server

import (
	"context"
	"time"
	"fmt"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type protocolServer struct {
	curppb.ProtocolServer

	fetchClusterOpCnt uint
	fetchClusterOpMap map[string][]FetchClusterResponse
	proposeOpCnt      uint
	proposeOpMap      map[string]ProposeResponse
	waitSyncedOpCnt   uint
	waitSyncedOpMap   map[string]WaitSyncedResponse
}

func NewProtocolServer(
	fetchClusterOpMap map[string][]FetchClusterResponse,
	proposeOpMap map[string]ProposeResponse,
	waitSyncedOpMap map[string]WaitSyncedResponse,
) *protocolServer {
	return &protocolServer{
		fetchClusterOpCnt: 0,
		fetchClusterOpMap: fetchClusterOpMap,
		proposeOpMap:      proposeOpMap,
		waitSyncedOpMap:   waitSyncedOpMap,
	}
}

func (s *protocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
	fmt.Println("fetch cluster request:", req)

	if resps, ok := s.fetchClusterOpMap[req.String()]; ok {
		res := resps[s.fetchClusterOpCnt]
		s.fetchClusterOpCnt++
		if res.Data != nil {
			return res.Data, nil
		} else if res.Error != nil {
			return nil, res.Error.Err()
		} else {
			panic("both data and error are nil")
		}
	}

	return nil, status.New(codes.NotFound, "not found").Err()
}

func (s *protocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	fmt.Println("propose request:", req.String())

	if res, ok := s.proposeOpMap[req.String()]; ok {
		if res.Error == nil {
			if res.Data != nil {
				switch res.Data.Type {
				case "range":
					er := &xlinepb.CommandResponse{
						ResponseWrapper: &xlinepb.CommandResponse_RangeResponse{
							RangeResponse: res.Data.Range,
						},
					}
					ber, err := proto.Marshal(er)
					if err != nil {
						panic(err)
					}
					return &curppb.ProposeResponse{
						Result: &curppb.CmdResult{
							Result: &curppb.CmdResult_Ok{
								Ok: ber,
							},
						},
					}, nil
				}
			}
			return &curppb.ProposeResponse{}, nil
		}
		return nil, res.Error.Err()
	}

	s.proposeOpCnt++
	switch s.proposeOpCnt {
	case 1:
		err := RpcError{Code: 13, Message: "shutting down"}
		return nil, err.Err()
	case 2, 3, 4:
		err := RpcError{Code: 13, Message: "key conflict"}
		return nil, err.Err()
	}

	return nil, status.New(codes.NotFound, "not found").Err()
}

func (s *protocolServer) WaitSynced(ctx context.Context, req *curppb.WaitSyncedRequest) (*curppb.WaitSyncedResponse, error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("wait synced:", req.String())

	if res, ok := s.waitSyncedOpMap[req.String()]; ok {
		if res.Error == nil {
			switch res.Data.ExecuteResult.Type {
			case "range":
				er := &xlinepb.CommandResponse{
					ResponseWrapper: &xlinepb.CommandResponse_RangeResponse{
						RangeResponse: res.Data.ExecuteResult.Range,
					},
				}
				ber, err := proto.Marshal(er)
				if err != nil {
					panic(err)
				}
				basr, err := proto.Marshal(res.Data.AfterSyncResult)
				if err != nil {
					panic(err)
				}
				return &curppb.WaitSyncedResponse{
					AfterSyncResult: &curppb.CmdResult{
						Result: &curppb.CmdResult_Ok{
							Ok: basr,
						},
					},
					ExeResult: &curppb.CmdResult{
						Result: &curppb.CmdResult_Ok{
							Ok: ber,
						},
					},
				}, nil
			}
		}
		return nil, res.Error.Err()
	}

	s.waitSyncedOpCnt++
	switch s.waitSyncedOpCnt {
	case 1:
		err := RpcError{Code: 13, Message: "shutting down"}
		return nil, err.Err()
	case 2, 3, 4:
		err := RpcError{Code: 13, Message: "key conflict"}
		return nil, err.Err()
	}

	return nil, status.New(codes.NotFound, "not found").Err()
}
