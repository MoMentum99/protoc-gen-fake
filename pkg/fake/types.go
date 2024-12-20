package fake

import (
	"context"
	"google.golang.org/protobuf/proto"
)

// Response represents a pre-defined response for a gRPC method
type Response struct {
	Method   string
	Response proto.Message
	Error    error
}

// CustomResponseFunc defines a function type for generating custom responses
type CustomResponseFunc func(ctx context.Context, req interface{}) (proto.Message, error)
