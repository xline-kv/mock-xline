package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/ghodss/yaml"
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	"github.com/xline-kv/mock-xline/server"
	"google.golang.org/grpc"
)

func main() {
	name := flag.String("name", "node1", "node name")
	addr := flag.String("addr", "127.0.0.1:8080", "node address")
	config := flag.String("config", "node1", "config name")
	flag.Parse()

	cfg, err := os.ReadFile(fmt.Sprintf("config/%s.yml", *config))
	if err != nil {
		panic("read config file fail")
	}

	ops := []server.Operation{}
	err = yaml.Unmarshal(cfg, &ops)
	if err != nil {
		panic(err)
	}

	fetchClusterOpMap := map[string][]server.FetchClusterResponse{}
	proposeOpMap := map[string]server.ProposeResponse{}
	waitSyncedOpMap := map[string]server.WaitSyncedResponse{}
	for _, op := range ops {
		switch op.Method {
		case "fetch_cluster":
			fetchClusterOpMap[op.Input.FetchCluster.String()] = op.Output.FetchCluster
		case "propose":
			proposeOpMap[op.Input.Propose.String()] = op.Output.Propose
		case "wait_synced":
			waitSyncedOpMap[op.Input.WaitSynced.String()] = op.Output.WaitSynced
		default:
			panic(fmt.Sprintf("unknown method: %s", op.Method))
		}
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	curppb.RegisterProtocolServer(s, server.NewProtocolServer(
		fetchClusterOpMap,
		proposeOpMap,
		waitSyncedOpMap,
	))
	fmt.Printf("%s starts member change service listening at %s\n", *name, lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
