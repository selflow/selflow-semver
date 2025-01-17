// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: docs/selflow-plugin-protocol/sfplugin-0.1.proto

package selflow_plugin

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

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	GetPluginSchema(ctx context.Context, in *GetPluginSchema_Request, opts ...grpc.CallOption) (*GetPluginSchema_Response, error)
	ValidatePluginConfigSchema(ctx context.Context, in *ValidatePluginConfigSchema_Request, opts ...grpc.CallOption) (*ValidatePluginConfigSchema_Response, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) GetPluginSchema(ctx context.Context, in *GetPluginSchema_Request, opts ...grpc.CallOption) (*GetPluginSchema_Response, error) {
	out := new(GetPluginSchema_Response)
	err := c.cc.Invoke(ctx, "/selflow_plugin.Plugin/GetPluginSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ValidatePluginConfigSchema(ctx context.Context, in *ValidatePluginConfigSchema_Request, opts ...grpc.CallOption) (*ValidatePluginConfigSchema_Response, error) {
	out := new(ValidatePluginConfigSchema_Response)
	err := c.cc.Invoke(ctx, "/selflow_plugin.Plugin/ValidatePluginConfigSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	GetPluginSchema(context.Context, *GetPluginSchema_Request) (*GetPluginSchema_Response, error)
	ValidatePluginConfigSchema(context.Context, *ValidatePluginConfigSchema_Request) (*ValidatePluginConfigSchema_Response, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) GetPluginSchema(context.Context, *GetPluginSchema_Request) (*GetPluginSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginSchema not implemented")
}
func (UnimplementedPluginServer) ValidatePluginConfigSchema(context.Context, *ValidatePluginConfigSchema_Request) (*ValidatePluginConfigSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePluginConfigSchema not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_GetPluginSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetPluginSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selflow_plugin.Plugin/GetPluginSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetPluginSchema(ctx, req.(*GetPluginSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ValidatePluginConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePluginConfigSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ValidatePluginConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selflow_plugin.Plugin/ValidatePluginConfigSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ValidatePluginConfigSchema(ctx, req.(*ValidatePluginConfigSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "selflow_plugin.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginSchema",
			Handler:    _Plugin_GetPluginSchema_Handler,
		},
		{
			MethodName: "ValidatePluginConfigSchema",
			Handler:    _Plugin_ValidatePluginConfigSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "docs/selflow-plugin-protocol/sfplugin-0.1.proto",
}

// ArchitectClient is the client API for Architect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchitectClient interface {
	ValidateStepConfigSchema(ctx context.Context, in *ValidateStepConfigSchema_Request, opts ...grpc.CallOption) (*ValidateStepConfigSchema_Response, error)
	RunStep(ctx context.Context, in *RunStep_Request, opts ...grpc.CallOption) (*RunStep_Response, error)
}

type architectClient struct {
	cc grpc.ClientConnInterface
}

func NewArchitectClient(cc grpc.ClientConnInterface) ArchitectClient {
	return &architectClient{cc}
}

func (c *architectClient) ValidateStepConfigSchema(ctx context.Context, in *ValidateStepConfigSchema_Request, opts ...grpc.CallOption) (*ValidateStepConfigSchema_Response, error) {
	out := new(ValidateStepConfigSchema_Response)
	err := c.cc.Invoke(ctx, "/selflow_plugin.Architect/ValidateStepConfigSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *architectClient) RunStep(ctx context.Context, in *RunStep_Request, opts ...grpc.CallOption) (*RunStep_Response, error) {
	out := new(RunStep_Response)
	err := c.cc.Invoke(ctx, "/selflow_plugin.Architect/RunStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchitectServer is the server API for Architect service.
// All implementations must embed UnimplementedArchitectServer
// for forward compatibility
type ArchitectServer interface {
	ValidateStepConfigSchema(context.Context, *ValidateStepConfigSchema_Request) (*ValidateStepConfigSchema_Response, error)
	RunStep(context.Context, *RunStep_Request) (*RunStep_Response, error)
	mustEmbedUnimplementedArchitectServer()
}

// UnimplementedArchitectServer must be embedded to have forward compatible implementations.
type UnimplementedArchitectServer struct {
}

func (UnimplementedArchitectServer) ValidateStepConfigSchema(context.Context, *ValidateStepConfigSchema_Request) (*ValidateStepConfigSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateStepConfigSchema not implemented")
}
func (UnimplementedArchitectServer) RunStep(context.Context, *RunStep_Request) (*RunStep_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunStep not implemented")
}
func (UnimplementedArchitectServer) mustEmbedUnimplementedArchitectServer() {}

// UnsafeArchitectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchitectServer will
// result in compilation errors.
type UnsafeArchitectServer interface {
	mustEmbedUnimplementedArchitectServer()
}

func RegisterArchitectServer(s grpc.ServiceRegistrar, srv ArchitectServer) {
	s.RegisterService(&Architect_ServiceDesc, srv)
}

func _Architect_ValidateStepConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateStepConfigSchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchitectServer).ValidateStepConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selflow_plugin.Architect/ValidateStepConfigSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchitectServer).ValidateStepConfigSchema(ctx, req.(*ValidateStepConfigSchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Architect_RunStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunStep_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchitectServer).RunStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selflow_plugin.Architect/RunStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchitectServer).RunStep(ctx, req.(*RunStep_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Architect_ServiceDesc is the grpc.ServiceDesc for Architect service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Architect_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "selflow_plugin.Architect",
	HandlerType: (*ArchitectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateStepConfigSchema",
			Handler:    _Architect_ValidateStepConfigSchema_Handler,
		},
		{
			MethodName: "RunStep",
			Handler:    _Architect_RunStep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "docs/selflow-plugin-protocol/sfplugin-0.1.proto",
}
