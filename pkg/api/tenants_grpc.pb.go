// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: tenants.proto

package __

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

const (
	TenantsService_CreateTenants_FullMethodName = "/api.TenantsService/CreateTenants"
	TenantsService_UpdateTenants_FullMethodName = "/api.TenantsService/UpdateTenants"
	TenantsService_DeleteTenants_FullMethodName = "/api.TenantsService/DeleteTenants"
)

// TenantsServiceClient is the client API for TenantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantsServiceClient interface {
	CreateTenants(ctx context.Context, in *CreateTenantsRequest, opts ...grpc.CallOption) (*Status, error)
	UpdateTenants(ctx context.Context, in *UpdateTenantsRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteTenants(ctx context.Context, in *IdTenantsRequest, opts ...grpc.CallOption) (*Status, error)
}

type tenantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantsServiceClient(cc grpc.ClientConnInterface) TenantsServiceClient {
	return &tenantsServiceClient{cc}
}

func (c *tenantsServiceClient) CreateTenants(ctx context.Context, in *CreateTenantsRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, TenantsService_CreateTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) UpdateTenants(ctx context.Context, in *UpdateTenantsRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, TenantsService_UpdateTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) DeleteTenants(ctx context.Context, in *IdTenantsRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, TenantsService_DeleteTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantsServiceServer is the server API for TenantsService service.
// All implementations must embed UnimplementedTenantsServiceServer
// for forward compatibility
type TenantsServiceServer interface {
	CreateTenants(context.Context, *CreateTenantsRequest) (*Status, error)
	UpdateTenants(context.Context, *UpdateTenantsRequest) (*Status, error)
	DeleteTenants(context.Context, *IdTenantsRequest) (*Status, error)
	mustEmbedUnimplementedTenantsServiceServer()
}

// UnimplementedTenantsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTenantsServiceServer struct {
}

func (UnimplementedTenantsServiceServer) CreateTenants(context.Context, *CreateTenantsRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenants not implemented")
}
func (UnimplementedTenantsServiceServer) UpdateTenants(context.Context, *UpdateTenantsRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenants not implemented")
}
func (UnimplementedTenantsServiceServer) DeleteTenants(context.Context, *IdTenantsRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenants not implemented")
}
func (UnimplementedTenantsServiceServer) mustEmbedUnimplementedTenantsServiceServer() {}

// UnsafeTenantsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantsServiceServer will
// result in compilation errors.
type UnsafeTenantsServiceServer interface {
	mustEmbedUnimplementedTenantsServiceServer()
}

func RegisterTenantsServiceServer(s grpc.ServiceRegistrar, srv TenantsServiceServer) {
	s.RegisterService(&TenantsService_ServiceDesc, srv)
}

func _TenantsService_CreateTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_CreateTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateTenants(ctx, req.(*CreateTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_UpdateTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).UpdateTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_UpdateTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).UpdateTenants(ctx, req.(*UpdateTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_DeleteTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).DeleteTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantsService_DeleteTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).DeleteTenants(ctx, req.(*IdTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantsService_ServiceDesc is the grpc.ServiceDesc for TenantsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TenantsService",
	HandlerType: (*TenantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenants",
			Handler:    _TenantsService_CreateTenants_Handler,
		},
		{
			MethodName: "UpdateTenants",
			Handler:    _TenantsService_UpdateTenants_Handler,
		},
		{
			MethodName: "DeleteTenants",
			Handler:    _TenantsService_DeleteTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenants.proto",
}