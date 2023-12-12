package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	"github.com/xline-kv/mock-xline/server"

)

func main()  {
	lis, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	curppb.RegisterProtocolServer(s, &server.ProtocolServer{})
	//...
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}