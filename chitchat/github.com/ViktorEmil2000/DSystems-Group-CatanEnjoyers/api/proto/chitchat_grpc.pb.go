// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: chitchat.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Services_ChatServic_FullMethodName = "/chitchat.Services/ChatServic"
)

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	ChatServic(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FromClient, FromServer], error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) ChatServic(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FromClient, FromServer], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Services_ServiceDesc.Streams[0], Services_ChatServic_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FromClient, FromServer]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Services_ChatServicClient = grpc.BidiStreamingClient[FromClient, FromServer]

// ServicesServer is the server API for Services service.
// All implementations should embed UnimplementedServicesServer
// for forward compatibility.
type ServicesServer interface {
	ChatServic(grpc.BidiStreamingServer[FromClient, FromServer]) error
}

// UnimplementedServicesServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicesServer struct{}

func (UnimplementedServicesServer) ChatServic(grpc.BidiStreamingServer[FromClient, FromServer]) error {
	return status.Errorf(codes.Unimplemented, "method ChatServic not implemented")
}
func (UnimplementedServicesServer) testEmbeddedByValue() {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	// If the following call pancis, it indicates UnimplementedServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_ChatServic_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServicesServer).ChatServic(&grpc.GenericServerStream[FromClient, FromServer]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Services_ChatServicServer = grpc.BidiStreamingServer[FromClient, FromServer]

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chitchat.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatServic",
			Handler:       _Services_ChatServic_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chitchat.proto",
}