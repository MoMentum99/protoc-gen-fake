package fake

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// createInterceptor creates a gRPC interceptor with given responses and custom handlers
func createInterceptor(
	defaultResponses map[string]Response,
	customHandlers map[string]CustomResponseFunc,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Check for custom handler first
		if customHandler, exists := customHandlers[method]; exists {
			response, err := customHandler(ctx, req)
			if err != nil {
				return err
			}
			return mergeResponse(reply, response)
		}

		// Fall back to default response if no custom handler exists
		if defaultResp, exists := defaultResponses[method]; exists {
			if defaultResp.Error != nil {
				return defaultResp.Error
			}
			return mergeResponse(reply, defaultResp.Response)
		}

		return nil
	}
}

// mergeResponse merges the response into the reply message
func mergeResponse(reply interface{}, response proto.Message) error {
	if replyMsg, ok := reply.(proto.Message); ok {
		proto.Reset(replyMsg)
		proto.Merge(replyMsg, response)
		return nil
	}
	return nil
}
