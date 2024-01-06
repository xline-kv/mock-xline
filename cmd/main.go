package main

import (
	"flag"
	"net"

	curppb "github.com/xline-kv/mock-xline/api/gen/curppb"
	"github.com/xline-kv/mock-xline/api/gen/mockpb"

	"github.com/sirupsen/logrus"
	"github.com/xline-kv/mock-xline/server"
	"google.golang.org/grpc"
)

func main() {
	name := flag.String("name", "node1", "node name")
	addr := flag.String("addr", "127.0.0.1:5555", "node address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protocolServer := server.NewProtocolServer()
	mockServer := server.NewMockServer(protocolServer)
	curppb.RegisterProtocolServer(s, protocolServer)
	mockpb.RegisterMockServer(s, mockServer)
	logrus.Infof("%v starts member change service listening at %v\n", *name, lis.Addr())

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
