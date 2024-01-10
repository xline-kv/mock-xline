package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	curppb "github.com/xline-kv/mock-xline/gen/curppb"

	"github.com/sirupsen/logrus"
	"github.com/xline-kv/mock-xline/server"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

func main() {
	name := flag.String("name", "node1", "node name")
	addr := flag.String("addr", "127.0.0.1:5555", "node address")
	config := flag.String("test", "protocol", "config name")
	flag.Parse()

	data, err := os.ReadFile(fmt.Sprintf("./config/%s.yml", *config))
	if err != nil {
		panic(err)
	}

	cfgs := server.Configs{}
	err = yaml.Unmarshal(data, &cfgs)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	curppb.RegisterProtocolServer(s, server.NewProtocolServer(cfgs))
	logrus.Infof("%s starts %s test service listening at %s\n", *name, *config, lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
