package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	// "fmt"
	// "text/template/parse"

	// "github.com/bytedance/sonic"
	"github.com/xline-kv/mock-xline/data"
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
)

type ProtocolServer struct {
	curppb.ProtocolServer

	FetchClusterMap map[data.Request]*data.Response[curppb.FetchClusterResponse]
	ProposeMap map[data.Request]*data.Response[curppb.ProposeResponse]
	WaitSyncedMap map[data.Request]*data.Response[curppb.WaitSyncedResponse]
}

func (s *ProtocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
	fmt.Println("fetch cluster")
	breq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	sreq := string(breq)
	if res, ok := s.FetchClusterMap[sreq]; ok {
		var resp *curppb.FetchClusterResponse
		if len(res.Res)-1 < int(res.Index) {
			panic("index error")
		} else if len(res.Res)-1 == int(res.Index) {
			resp = &res.Res[res.Index].Res
		} else {
			resp = &res.Res[res.Index].Res
			res.Index++
		}
		return resp, nil
	}
	panic("data error")
}

func (s *ProtocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	fmt.Println("propose")
	breq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	sreq := string(breq)
	if _, ok := s.ProposeMap[sreq]; ok {

	}
	if res, ok := s.ProposeMap["default"]; ok {
		var resp *curppb.ProposeResponse
		if len(res.Res)-1 < int(res.Index) {
			panic("index error")
		} else if len(res.Res)-1 == int(res.Index) {
			if res.Res[res.Index].Err != nil {
				res.Index++
				return nil, res.Res[res.Index].Err
			}
			resp = &res.Res[res.Index].Res
		} else {
			if res.Res[res.Index].Err != nil {
				res.Index++
				return nil, res.Res[res.Index].Err

			}
			resp = &res.Res[res.Index].Res
			res.Index++
		}
		return resp, nil
	}
	panic("map error")
}

func (s *ProtocolServer) WaitSynced(ctx context.Context, req *curppb.WaitSyncedRequest) (*curppb.WaitSyncedResponse, error) {
	fmt.Println("wait synced")
	time.Sleep(500 * time.Millisecond)
	breq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	sreq := string(breq)
	if _, ok := s.WaitSyncedMap[sreq]; ok {

	}
	if res, ok := s.WaitSyncedMap["default"]; ok {
		var resp *curppb.WaitSyncedResponse
		if len(res.Res)-1 < int(res.Index) {
			panic("index error")
		} else if len(res.Res)-1 == int(res.Index) {
			if res.Res[res.Index].Err != nil {
				return nil, res.Res[res.Index].Err
			}
			resp = &res.Res[res.Index].Res
		} else {
			if res.Res[res.Index].Err != nil {
				return nil, res.Res[res.Index].Err
			}
			resp = &res.Res[res.Index].Res
			res.Index++
		}
		return resp, nil
	}
	panic("map error")
}
