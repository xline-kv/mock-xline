package data

import (
	// curppb "github.com/xline-kv/mock-xline/gen/curppb"
	xlinepb "github.com/xline-kv/mock-xline/gen/xlinepb"
	// "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewProtocolDataMap() map[Request]Response {

	cmd := &xlinepb.Command{
		Keys: []*xlinepb.KeyRange{
			{
				Key:      []byte("key"),
				RangeEnd: []byte("key"),
			},
		},
		Request: &xlinepb.RequestWithToken{
			Token: nil,
			RequestWrapper: &xlinepb.RequestWithToken_PutRequest{
				PutRequest: &xlinepb.PutRequest{
					Key:   []byte("key"),
					Value: []byte("value"),
				},
			},
		},
		ProposeId: &xlinepb.ProposeId{
			ClientId: 1,
			SeqNum:   1,
		},
	}
	req, err := NewProposeRequest(cmd, 1)
	if err != nil {
		panic(err)
	}

	rpcErr := status.Errorf(codes.FailedPrecondition, "wrong-cluster-version")
	res := NewResponse(nil, rpcErr)

	return map[Request]Response{
		*req: res,
	}
}
