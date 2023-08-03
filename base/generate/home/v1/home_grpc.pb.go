// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: home/v1/home.proto

package homev1

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

// HomeServiceClient is the client API for HomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeServiceClient interface {
	// Sends a greeting
	GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error)
}

type homeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeServiceClient(cc grpc.ClientConnInterface) HomeServiceClient {
	return &homeServiceClient{cc}
}

func (c *homeServiceClient) GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error) {
	out := new(GetHomeResponse)
	err := c.cc.Invoke(ctx, "/home.v1.HomeService/GetHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeServiceServer is the server API for HomeService service.
// All implementations should embed UnimplementedHomeServiceServer
// for forward compatibility
type HomeServiceServer interface {
	// Sends a greeting
	GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error)
}

// UnimplementedHomeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHomeServiceServer struct {
}

func (UnimplementedHomeServiceServer) GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHome not implemented")
}

// UnsafeHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeServiceServer will
// result in compilation errors.
type UnsafeHomeServiceServer interface {
	mustEmbedUnimplementedHomeServiceServer()
}

func RegisterHomeServiceServer(s grpc.ServiceRegistrar, srv HomeServiceServer) {
	s.RegisterService(&HomeService_ServiceDesc, srv)
}

func _HomeService_GetHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeServiceServer).GetHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/home.v1.HomeService/GetHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeServiceServer).GetHome(ctx, req.(*GetHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeService_ServiceDesc is the grpc.ServiceDesc for HomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "home.v1.HomeService",
	HandlerType: (*HomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHome",
			Handler:    _HomeService_GetHome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home/v1/home.proto",
}
