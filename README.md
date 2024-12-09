# protoc-gen-fake

A protocol buffer compiler plugin that generates fake gRPC client responses for testing and development.

## Installation

```bash
go install github.com/ao-labs/protoc-gen-fake@latest
```

## Usage

### With protoc

```bash
protoc --fake_out=. your.proto
```

### With buf

1. Add the plugin to your `buf.gen.yaml`:

```yaml
version: v1
plugins:
  - plugin: fake
    out: gen/go
    opt: paths=source_relative
```

2. Run buf generate:

```bash
buf generate
```

## Generated Code

The plugin generates fake responses for each service defined in your proto files. For example:

```protobuf
service UserService {
    rpc GetUser (GetUserRequest) returns (GetUserResponse);
}
```

Will generate:

```go
var DefaultUserServiceResponses = map[string]Response{
    "/package.UserService/GetUser": {
        Method: "/package.UserService/GetUser",
        Response: &GetUserResponse{
            // Default values for fields
        },
    },
}
```

## Using Generated Fakes

```go
// Create a fake connection with default responses
conn := NewFakeClientConn(example.DefaultUserServiceResponses)
client := example.NewUserServiceClient(conn)

// Use the client as normal
resp, err := client.GetUser(context.Background(), &example.GetUserRequest{
    Id: "123",
})
```

## Customizing Responses

You can customize the default responses:

```go
responses := example.DefaultUserServiceResponses
responses["/package.UserService/GetUser"] = Response{
    Method: "/package.UserService/GetUser",
    Response: &example.GetUserResponse{
        Id: "custom_id",
        Name: "Custom Name",
    },
}

conn := NewFakeClientConn(responses)
```

## Features

- Automatic generation of fake responses for all service methods
- Support for all basic protobuf types
- Customizable default values
- Compatible with both protoc and buf
- Easy integration with existing gRPC clients

