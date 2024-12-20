package fake

import (
	"google.golang.org/grpc"
)

// NewFakeClientConn creates a new gRPC client connection that returns fake responses
func NewFakeClientConn(responses map[string]Response) *grpc.ClientConn {
	return newClientConn(responses, nil)
}

// NewCustomFakeClientConn creates fake gRPC client connection with default responses and custom handlers
func NewCustomFakeClientConn(
	defaultResponses map[string]Response,
	customHandlers map[string]CustomResponseFunc,
) *grpc.ClientConn {
	return newClientConn(defaultResponses, customHandlers)
}

// newClientConn creates a new gRPC client connection with the given configuration
func newClientConn(
	defaultResponses map[string]Response,
	customHandlers map[string]CustomResponseFunc,
) *grpc.ClientConn {
	conn, _ := grpc.Dial(
		"fake:///",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(createInterceptor(defaultResponses, customHandlers)),
	)
	return conn
}
