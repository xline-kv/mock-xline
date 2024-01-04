package server

import (
	"context"
	"fmt"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	"github.com/xline-kv/mock-xline/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type protocolServer struct {
	curppb.ProtocolServer

	fetchClusterSample sample.FetchClusterMap
	proposeSample      sample.ProposeMap
	waitSyncSample     sample.WaitSyncedMap
	commandSample      sample.CommandMap
}

func NewProtocolServer(
	fetchClusterSample sample.FetchClusterMap,
	proposeSample sample.ProposeMap,
	waitSyncSample sample.WaitSyncedMap,
) *protocolServer {
	return &protocolServer{
		fetchClusterSample: fetchClusterSample,
		proposeSample:      proposeSample,
		waitSyncSample:     waitSyncSample,
		commandSample:      sample.CommandMap{},
	}
}

func (s *protocolServer) FetchCluster(ctx context.Context, req *curppb.FetchClusterRequest) (*curppb.FetchClusterResponse, error) {
	fmt.Println("fetch cluster", req)

	res, ok := s.fetchClusterSample[req.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found")
	}

	return res, nil
}

func (s *protocolServer) Propose(ctx context.Context, req *curppb.ProposeRequest) (*curppb.ProposeResponse, error) {
	fmt.Println("propose:", req)

	var cmd xlinepb.Command
	bcmd := req.Command
	err := proto.Unmarshal(bcmd, &cmd)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid command")
	}

	s.commandSample[req.GetProposeId().String()] = &cmd

	res, ok := s.proposeSample[cmd.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found")
	}

	return res, nil
}

func (s *protocolServer) WaitSynced(ctx context.Context, req *curppb.WaitSyncedRequest) (*curppb.WaitSyncedResponse, error) {
	fmt.Println("wait synced:", req)

	cmd, ok := s.commandSample[req.GetProposeId().String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found")
	}

	res, ok := s.waitSyncSample[cmd.String()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "data not found")
	}

	return res, nil
}
