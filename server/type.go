package server

import (
	"time"

	"github.com/xline-kv/mock-xline/api/gen/mockpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type mockRequest = string
type spid = string
type bcmd = []byte
type scmd = string
type mockResponses = []*mockpb.Response

func status2error(sts *mockpb.Status) error {
	return status.Errorf(codes.Code(sts.GetCode()), sts.GetMessage())
}

func timestamp2duration(timestamp *timestamppb.Timestamp) time.Duration {
	nanos := timestamp.GetSeconds()*int64(time.Second) + int64(timestamp.GetNanos())
	return time.Duration(nanos)
}
