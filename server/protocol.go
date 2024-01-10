package server

import (
	"context"
	"time"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type protocolServer struct {
	curppb.ProtocolServer

	fetchClusterCnt  uint
	putProposeCnt    uint
	putWaitSyncedCnt uint

	fetchClusterConfigMap map[FetchClusterConfigRequest][]FetchClusterConfigResponse
	putConfigMap          map[PutConfigRequest][]PutConfigResponseWrapper
	waitSyncedConfigMap   map[string]PutConfigRequest
}

func NewProtocolServer(cfgs Configs) *protocolServer {
	return &protocolServer{
		fetchClusterConfigMap: BuildFetchClusterConfigMapFromConfigs(cfgs),
		putConfigMap:          BuildPutConfigMapFromConfigs(cfgs),
		waitSyncedConfigMap:   map[string]PutConfigRequest{},
	}
}

func (s *protocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
	logrus.Infof("fetch cluster: %+v\n", req)

	s.fetchClusterCnt++

	for fcReq, fcReses := range s.fetchClusterConfigMap {
		if CmpFetchClusterRequest(&fcReq, req) {
			for _, fcCfgRes := range fcReses {
				if fcCfgRes.Index == 0 {
					if fcCfgRes.Type == "success" {
						fcRes := BuildFetchClusterResponseFromFetchClusterConfigResponse(&fcCfgRes)
						return fcRes, nil
					}
					if fcCfgRes.Type == "error" {
						return nil, status.Errorf(codes.Code(fcCfgRes.Error.Code), fcCfgRes.Error.Message)
					}
				}
				if fcCfgRes.Index == s.fetchClusterCnt {
					if fcCfgRes.Type == "success" {
						fcRes := BuildFetchClusterResponseFromFetchClusterConfigResponse(&fcCfgRes)
						return fcRes, nil
					}
					if fcCfgRes.Type == "error" {
						return nil, status.Errorf(codes.Code(fcCfgRes.Error.Code), fcCfgRes.Error.Message)
					}
				}
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found")
}

func (s *protocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	logrus.Infof("propose: %+v\n", req)

	s.putProposeCnt++

	cmd := xlinepb.Command{}
	err := proto.Unmarshal(req.Command, &cmd)
	if err != nil {
		return nil, err
	}

	if cmd.Request.GetPutRequest() != nil {
		putReq := cmd.Request.GetPutRequest()
		for cfgReq, cfgRess := range s.putConfigMap {
			s.waitSyncedConfigMap[req.ProposeId.String()] = cfgReq
			if CmpPutRequest(&cfgReq, putReq) {
				for _, cfgRes := range cfgRess {
					if cfgRes.Index == 0 {
						ppsRes := cfgRes.Propose
						if ppsRes.Type == "success" {
							putRes := BuildPutResponseFromPutConfigResponse(&ppsRes)
							bPutRes, _ := proto.Marshal(putRes)
							return &curppb.ProposeResponse{Result: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}}, nil
						}
						if ppsRes.Type == "error" {
							return nil, status.Errorf(codes.Code(ppsRes.Error.Code), ppsRes.Error.Message)
						}
					}
					if cfgRes.Index == s.putProposeCnt {
						ppsRes := cfgRes.Propose
						if ppsRes.Type == "success" {
							putRes := BuildPutResponseFromPutConfigResponse(&ppsRes)
							bPutRes, _ := proto.Marshal(putRes)
							return &curppb.ProposeResponse{Result: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}}, nil
						}
						if ppsRes.Type == "error" {
							return nil, status.Errorf(codes.Code(ppsRes.Error.Code), ppsRes.Error.Message)
						}
					}
				}
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found in put config map")
}

func (s *protocolServer) WaitSynced(ctx context.Context, req *curppb.WaitSyncedRequest) (*curppb.WaitSyncedResponse, error) {
	logrus.Infof("wait synced: %+v\n", req)
	time.Sleep(100 * time.Microsecond)

	s.putWaitSyncedCnt++

	data, ok := s.waitSyncedConfigMap[req.ProposeId.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found in cmd map")
	}

	for cfgReq, cfgRess := range s.putConfigMap {
		if data.Key == cfgReq.Key && data.Value == cfgReq.Value && data.PrevKv == cfgReq.PrevKv {
			for _, cfgRes := range cfgRess {
				if cfgRes.Index == 0 {
					wsRes := cfgRes.WaitSynced
					if wsRes.Type == "success" {
						putRes := BuildPutResponseFromPutConfigResponse(&wsRes)
						bPutRes, _ := proto.Marshal(putRes)
						return &curppb.WaitSyncedResponse{AfterSyncResult: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}, ExeResult: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}}, nil
					}
					if wsRes.Type == "error" {
						return nil, status.Errorf(codes.Code(wsRes.Error.Code), wsRes.Error.Message)
					}
				}
				if cfgRes.Index == s.putProposeCnt {
					wsRes := cfgRes.WaitSynced
					if wsRes.Type == "success" {
						putRes := BuildPutResponseFromPutConfigResponse(&wsRes)
						bPutRes, _ := proto.Marshal(putRes)
						return &curppb.WaitSyncedResponse{AfterSyncResult: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}, ExeResult: &curppb.CmdResult{Result: &curppb.CmdResult_Ok{Ok: bPutRes}}}, nil
					}
					if wsRes.Type == "error" {
						return nil, status.Errorf(codes.Code(wsRes.Error.Code), wsRes.Error.Message)
					}
				}
			}

		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found in put config map")
}
