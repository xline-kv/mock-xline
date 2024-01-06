package server

import (
	"context"
	"time"

	curppb "github.com/xline-kv/mock-xline/api/gen/curppb"
	"github.com/xline-kv/mock-xline/api/gen/mockpb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type protocolServer struct {
	curppb.ProtocolServer

	fetchClusterCnt uint64
	proposeCnt      uint64
	waitSyncedCnt   uint64

	fetchClusterMap map[mockRequest]mockResponses
	proposeMap      map[scmd]mockResponses
	waitSyncedMap   map[scmd]mockResponses
	commandMap      map[spid]bcmd
}

func NewProtocolServer() *protocolServer {
	return &protocolServer{
		fetchClusterMap: map[string][]*mockpb.Response{},
		proposeMap:      map[string][]*mockpb.Response{},
		commandMap:      map[string][]byte{},
		waitSyncedMap:   map[string][]*mockpb.Response{},
	}
}

func (s *protocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
	logrus.Info("rcv fetch cluster request:", req.String())

	resps, ok := s.fetchClusterMap[req.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found, unknown request")
	}

	var res curppb.FetchClusterResponse
	for _, resp := range resps {
		if resp.GetIndex() == s.fetchClusterCnt {
			s.fetchClusterCnt++
			if resp.GetSuccess() != nil {
				if err := proto.Unmarshal(resp.GetSuccess(), &res); err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
				}
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return &res, nil
			} else if resp.GetError() != nil {
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return nil, status2error(resp.GetError())
			} else {
				return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found")
}

func (s *protocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	logrus.Info("rcv propose request:", req.String())

	s.commandMap[req.GetProposeId().String()] = req.GetCommand()

	resps, ok := s.proposeMap[string(req.GetCommand())]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found, unknown request")
	}

	var res curppb.ProposeResponse
	for _, resp := range resps {
		if resp.GetIndex() == s.proposeCnt {
			s.proposeCnt++
			if resp.GetSuccess() != nil {
				if err := proto.Unmarshal(resp.GetSuccess(), &res); err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
				}
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return &res, nil
			} else if resp.GetError() != nil {
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return nil, status2error(resp.GetError())
			} else {
				return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found")
}

func (s *protocolServer) WaitSynced(ctx context.Context, req *curppb.WaitSyncedRequest) (*curppb.WaitSyncedResponse, error) {
	logrus.Info("rcv wait synced request:", req.String())

	bcmd, ok := s.commandMap[req.ProposeId.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "command not found")
	}

	resps, ok := s.waitSyncedMap[string(bcmd)]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "responses not found")
	}

	var res curppb.WaitSyncedResponse
	for _, resp := range resps {
		if resp.GetIndex() == s.waitSyncedCnt {
			s.waitSyncedCnt++
			if resp.GetSuccess() != nil {
				if err := proto.Unmarshal(resp.GetSuccess(), &res); err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
				}
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return &res, nil
			} else if resp.GetError() != nil {
				time.Sleep(timestamp2duration(resp.GetDelay()))
				return nil, status2error(resp.GetError())
			} else {
				return nil, status.Errorf(codes.InvalidArgument, "invalid mock response")
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "data not found")
}
