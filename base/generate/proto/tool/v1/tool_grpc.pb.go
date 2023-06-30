// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/tool/v1/tool.proto

package toolv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ToolServiceClient is the client API for ToolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToolServiceClient interface {
	// Sends a greeting
	Timer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*TimerResponse, error)
}

type toolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToolServiceClient(cc grpc.ClientConnInterface) ToolServiceClient {
	return &toolServiceClient{cc}
}

func (c *toolServiceClient) Timer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*TimerResponse, error) {
	out := new(TimerResponse)
	err := c.cc.Invoke(ctx, "/proto.tool.v1.ToolService/Timer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToolServiceServer is the server API for ToolService service.
// All implementations should embed UnimplementedToolServiceServer
// for forward compatibility
type ToolServiceServer interface {
	// Sends a greeting
	Timer(context.Context, *TimerRequest) (*TimerResponse, error)
}

// UnimplementedToolServiceServer should be embedded to have forward compatible implementations.
type UnimplementedToolServiceServer struct {
}

func (UnimplementedToolServiceServer) Timer(context.Context, *TimerRequest) (*TimerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timer not implemented")
}

// UnsafeToolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToolServiceServer will
// result in compilation errors.
type UnsafeToolServiceServer interface {
	mustEmbedUnimplementedToolServiceServer()
}

func RegisterToolServiceServer(s grpc.ServiceRegistrar, srv ToolServiceServer) {
	s.RegisterService(&ToolService_ServiceDesc, srv)
}

func _ToolService_Timer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).Timer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.tool.v1.ToolService/Timer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).Timer(ctx, req.(*TimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToolService_ServiceDesc is the grpc.ServiceDesc for ToolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.tool.v1.ToolService",
	HandlerType: (*ToolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Timer",
			Handler:    _ToolService_Timer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tool/v1/tool.proto",
}
