package server

import (
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
)

// mock request configuration
type Configs []*Config

// configuration of a mock request
type Config struct {
	// mock request type, possible fields are `put_request`, `get_request`, `delete_request`, etc.
	Type string `yaml:"type"`
	// request configuration, must match the`type` above.
	ConfigRequest struct {
		// fetch cluster request configuration
		FetchClusterRequest FetchClusterConfigRequest `yaml:"fetch_cluster_request"`
		// put request configuration
		PutRequest PutConfigRequest `yaml:"put_request"`
		// ...
	} `yaml:"config_request"`
	// response configuration, must match the`type` above.
	ConfigResponse struct {
		// fetch cluster response configuration
		FetchClusterResponse []FetchClusterConfigResponse `yaml:"fetch_cluster_response"`
		// put response configuration
		PutResponse []PutConfigResponseWrapper `yaml:"put_response"`
		// ...
	} `yaml:"config_response"`
}

// fetch cluster request configuration
type FetchClusterConfigRequest struct {
	Linearizable bool `yaml:"linearizable"`
}

// put request configuration
type PutConfigRequest struct {
	Key    string `yaml:"key"`
	Value  string `yaml:"value"`
	PrevKv bool   `yaml:"prev_kv"`
}

// put response configuration
type PutConfigResponseWrapper struct {
	Propose    PutConfigResponse `yaml:"propose"`
	WaitSynced PutConfigResponse `yaml:"wait_synced"`
	Index      uint              `yaml:"index"`
}

type FetchClusterConfigResponse struct {
	// whether the response was successful or not
	// possible fields are `success`, `error`
	Type string `yaml:"type"`
	// put response configuration for success
	Success FetchClusterResponse `yaml:"success"`
	// grpc error configuration for error
	Error GrpcError `yaml:"error"`
	// response delay, used to simulate network delay
	Delay uint `yaml:"delay"`
	Index uint `yaml:"index"`
}

type PutConfigResponse struct {
	// whether the response was successful or not
	// possible fields are `success`, `error`
	Type string `yaml:"type"`
	// put response configuration for success
	Success PutResponse `yaml:"success"`
	// grpc error configuration for error
	Error GrpcError `yaml:"error"`
	// response delay, used to simulate network delay
	Delay uint `yaml:"delay"`
}

// fetch cluster response configuration
type FetchClusterResponse struct {
	LeaderId  uint64 `yaml:"leader_id"`
	Term      uint64 `yaml:"term"`
	ClusterId uint64 `yaml:"cluster_id"`
	Members   []struct {
		Id        uint64   `yaml:"id"`
		Name      string   `yaml:"name"`
		Addrs     []string `yaml:"addrs"`
		IsLearner bool     `yaml:"is_learner"`
	} `yaml:"members"`
	ClusterVersion uint64 `yaml:"cluster_version"`
}

// put response configuration
type PutResponse struct {
	// previous key value in put response
	PrevKv struct {
		// key in previous key value
		Key string `yaml:"key"`
		// value in previous key value
		Value string `yaml:"value"`
	} `yaml:"prev_kv"`
}

// grpc error configuration
type GrpcError struct {
	// grpc error type
	Code uint `yaml:"code"`
	// grpc error message
	Message string `json:"message"`
}

func BuildFetchClusterConfigMapFromConfigs(cfgs Configs) map[FetchClusterConfigRequest][]FetchClusterConfigResponse {
	ret := make(map[FetchClusterConfigRequest][]FetchClusterConfigResponse)
	for _, cfg := range cfgs {
		if cfg.Type == "fetch_cluster" {
			ret[cfg.ConfigRequest.FetchClusterRequest] = cfg.ConfigResponse.FetchClusterResponse
		}
	}
	return ret
}

func BuildPutConfigMapFromConfigs(cfgs Configs) map[PutConfigRequest][]PutConfigResponseWrapper {
	ret := make(map[PutConfigRequest][]PutConfigResponseWrapper)
	for _, cfg := range cfgs {
		if cfg.Type == "put" {
			ret[cfg.ConfigRequest.PutRequest] = cfg.ConfigResponse.PutResponse
		}
	}
	return ret
}

func CmpFetchClusterRequest(expect *FetchClusterConfigRequest, actual *curppb.FetchClusterRequest) bool {
	return expect.Linearizable == actual.Linearizable
}

func CmpPutRequest(expect *PutConfigRequest, actual *xlinepb.PutRequest) bool {
	if expect.Key != string(actual.Key) {
		return false
	}
	if expect.Value != string(actual.Value) {
		return false
	}
	if expect.PrevKv != actual.PrevKv {
		return false
	}
	return true
}

func BuildFetchClusterResponseFromFetchClusterConfigResponse(res *FetchClusterConfigResponse) *curppb.FetchClusterResponse {
	ret := curppb.FetchClusterResponse{
		LeaderId:  &res.Success.LeaderId,
		Term:      res.Success.Term,
		ClusterId: res.Success.ClusterId,
		Members: []*curppb.Member{
			{
				Id:    res.Success.Members[0].Id,
				Name:  res.Success.Members[0].Name,
				Addrs: res.Success.Members[0].Addrs,
			},
		},
		ClusterVersion: res.Success.ClusterVersion,
	}
	return &ret
}

func BuildPutResponseFromPutConfigResponse(res *PutConfigResponse) *xlinepb.PutResponse {
	ret := xlinepb.PutResponse{
		PrevKv: &xlinepb.KeyValue{
			Key:   []byte(res.Success.PrevKv.Key),
			Value: []byte(res.Success.PrevKv.Value),
		},
	}
	return &ret
}
