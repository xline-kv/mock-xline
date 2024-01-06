package server

import (
	"context"

	curppb "github.com/xline-kv/mock-xline/api/gen/curppb"
	"github.com/xline-kv/mock-xline/api/gen/mockpb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type mockServer struct {
	mockpb.MockServer

	*protocolServer
}

func NewMockServer(protocol *protocolServer) *mockServer {
	return &mockServer{
		protocolServer: protocol,
	}
}

func (s mockServer) MockFetchCluster(ctx context.Context, req *mockpb.MockRequest) (*mockpb.MockResponse, error) {
	logrus.Info("rcv mock fetch cluster request:", req.String())

	var r curppb.FetchClusterRequest
	if err := proto.Unmarshal(req.GetRequest(), &r); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid mock request")
	}

	s.fetchClusterMap[r.String()] = req.GetResponse()

	return &mockpb.MockResponse{}, nil
}

func (s mockServer) MockPropose(ctx context.Context, req *mockpb.MockRequest) (*mockpb.MockResponse, error) {
	logrus.Info("rcv mock propose request:", req.String())

	s.proposeMap[string(req.GetRequest())] = req.GetResponse()

	return &mockpb.MockResponse{}, nil
}

func (s mockServer) MockWaitSynced(ctx context.Context, req *mockpb.MockRequest) (*mockpb.MockResponse, error) {
	logrus.Info("rcv mock wait synced request:", req.String())

	s.waitSyncedMap[string(req.GetRequest())] = req.GetResponse()

	return &mockpb.MockResponse{}, nil
}
