package server

import (
	curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Operation struct {
	Method string `json:"method"`
	Input  struct {
		FetchCluster *curppb.FetchClusterRequest `json:"fetch_cluster"`
		Propose      *curppb.ProposeRequest      `json:"propose"`
		WaitSynced   *curppb.WaitSyncedRequest   `json:"wait_synced"`
	} `json:"input"`
	Output struct {
		FetchCluster []FetchClusterResponse `json:"fetch_cluster"`
		Propose      ProposeResponse        `json:"propose"`
		WaitSynced   WaitSyncedResponse     `json:"wait_synced"`
	} `json:"output"`
}

type FetchClusterResponse struct {
	Data  *curppb.FetchClusterResponse `json:"data"`
	Error *RpcError                    `json:"error"`
}

type ProposeResponse struct {
	Data *struct {
		Type  string                 `json:"type"`
		Range *xlinepb.RangeResponse `json:"range"`
	} `json:"data"`
	Error *RpcError `json:"error"`
}

type WaitSyncedResponse struct {
	Data *struct {
		ExecuteResult *struct {
			Type  string                 `json:"type"`
			Range *xlinepb.RangeResponse `json:"range"`
		} `json:"execute_result"`
		AfterSyncResult *xlinepb.SyncResponse `json:"after_sync_result"`
	} `json:"data"`
	Error *RpcError `json:"error"`
}

type RpcError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func (e *RpcError) Err() error {
	var detail *curppb.CurpError = nil
	switch e.Message {
	case "key conflict":
		detail = &curppb.CurpError{Err: &curppb.CurpError_KeyConflict{}}
	case "duplicated":
		detail = &curppb.CurpError{Err: &curppb.CurpError_Duplicated{}}
	case "shutting down":
		detail = &curppb.CurpError{Err: &curppb.CurpError_ShuttingDown{}}
	}

	st, err := status.New(codes.Code(e.Code), e.Message).WithDetails(detail)
	if err != nil {
		panic(err)
	}
	return st.Err()
}
