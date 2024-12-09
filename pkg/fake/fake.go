package fake

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Response represents a pre-defined response for a gRPC method
type Response struct {
	Method   string
	Response proto.Message
	Error    error
}

// NewFakeClientConn creates a new gRPC client connection that returns fake responses
func NewFakeClientConn(responses map[string]Response) *grpc.ClientConn {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if fakeResp, exists := responses[method]; exists {
			if fakeResp.Error != nil {
				return fakeResp.Error
			}

			if replyMsg, ok := reply.(proto.Message); ok {
				proto.Reset(replyMsg)
				proto.Merge(replyMsg, fakeResp.Response)
				return nil
			}
		}
		return nil
	}

	conn, _ := grpc.Dial(
		"fake:///",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor),
	)

	return conn
}
