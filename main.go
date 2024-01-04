package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	"github.com/xline-kv/mock-xline/sample"
	"github.com/xline-kv/mock-xline/server"
)

func main() {
	name := flag.String("name", "node1", "node name")
	addr := flag.String("addr", "127.0.0.1:2379", "node address")
	test := flag.String("test", "protocol_client", "test case")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	switch *test {
	case "protocol_client":
		s := grpc.NewServer()
		curppb.RegisterProtocolServer(
			s, server.NewProtocolServer(
				sample.NewFetchClusterMap(*test),
				sample.NewProposeMap(*test),
				sample.NewWaitSyncedMap(*test),
			),
		)
		fmt.Printf("%v starts member change service listening at %v\n", *name, lis.Addr())
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	default:
		panic("unknown test case")
	}
}
