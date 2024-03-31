// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: rq/v1/service.proto

package v1

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
	RQ_Query_FullMethodName         = "/v1.RQ/Query"
	RQ_UploadSnippet_FullMethodName = "/v1.RQ/UploadSnippet"
	RQ_RenameSnippet_FullMethodName = "/v1.RQ/RenameSnippet"
	RQ_ListSnippets_FullMethodName  = "/v1.RQ/ListSnippets"
	RQ_GetSnippet_FullMethodName    = "/v1.RQ/GetSnippet"
	RQ_GetVersion_FullMethodName    = "/v1.RQ/GetVersion"
)

// RQClient is the client API for RQ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RQClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	UploadSnippet(ctx context.Context, in *UploadSnippetRequest, opts ...grpc.CallOption) (*UploadSnippetResponse, error)
	RenameSnippet(ctx context.Context, in *RenameSnippetRequest, opts ...grpc.CallOption) (*RenameSnippetResponse, error)
	ListSnippets(ctx context.Context, in *ListStippetsRequest, opts ...grpc.CallOption) (*SnippetInfo_List, error)
	GetSnippet(ctx context.Context, in *GetSnippetRequest, opts ...grpc.CallOption) (*GetSnippetResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type rQClient struct {
	cc grpc.ClientConnInterface
}

func NewRQClient(cc grpc.ClientConnInterface) RQClient {
	return &rQClient{cc}
}

func (c *rQClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, RQ_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rQClient) UploadSnippet(ctx context.Context, in *UploadSnippetRequest, opts ...grpc.CallOption) (*UploadSnippetResponse, error) {
	out := new(UploadSnippetResponse)
	err := c.cc.Invoke(ctx, RQ_UploadSnippet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rQClient) RenameSnippet(ctx context.Context, in *RenameSnippetRequest, opts ...grpc.CallOption) (*RenameSnippetResponse, error) {
	out := new(RenameSnippetResponse)
	err := c.cc.Invoke(ctx, RQ_RenameSnippet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rQClient) ListSnippets(ctx context.Context, in *ListStippetsRequest, opts ...grpc.CallOption) (*SnippetInfo_List, error) {
	out := new(SnippetInfo_List)
	err := c.cc.Invoke(ctx, RQ_ListSnippets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rQClient) GetSnippet(ctx context.Context, in *GetSnippetRequest, opts ...grpc.CallOption) (*GetSnippetResponse, error) {
	out := new(GetSnippetResponse)
	err := c.cc.Invoke(ctx, RQ_GetSnippet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rQClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, RQ_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RQServer is the server API for RQ service.
// All implementations must embed UnimplementedRQServer
// for forward compatibility
type RQServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	UploadSnippet(context.Context, *UploadSnippetRequest) (*UploadSnippetResponse, error)
	RenameSnippet(context.Context, *RenameSnippetRequest) (*RenameSnippetResponse, error)
	ListSnippets(context.Context, *ListStippetsRequest) (*SnippetInfo_List, error)
	GetSnippet(context.Context, *GetSnippetRequest) (*GetSnippetResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*Version, error)
	mustEmbedUnimplementedRQServer()
}

// UnimplementedRQServer must be embedded to have forward compatible implementations.
type UnimplementedRQServer struct {
}

func (UnimplementedRQServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedRQServer) UploadSnippet(context.Context, *UploadSnippetRequest) (*UploadSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadSnippet not implemented")
}
func (UnimplementedRQServer) RenameSnippet(context.Context, *RenameSnippetRequest) (*RenameSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameSnippet not implemented")
}
func (UnimplementedRQServer) ListSnippets(context.Context, *ListStippetsRequest) (*SnippetInfo_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnippets not implemented")
}
func (UnimplementedRQServer) GetSnippet(context.Context, *GetSnippetRequest) (*GetSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnippet not implemented")
}
func (UnimplementedRQServer) GetVersion(context.Context, *GetVersionRequest) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedRQServer) mustEmbedUnimplementedRQServer() {}

// UnsafeRQServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RQServer will
// result in compilation errors.
type UnsafeRQServer interface {
	mustEmbedUnimplementedRQServer()
}

func RegisterRQServer(s grpc.ServiceRegistrar, srv RQServer) {
	s.RegisterService(&RQ_ServiceDesc, srv)
}

func _RQ_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RQ_UploadSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).UploadSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_UploadSnippet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).UploadSnippet(ctx, req.(*UploadSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RQ_RenameSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).RenameSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_RenameSnippet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).RenameSnippet(ctx, req.(*RenameSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RQ_ListSnippets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStippetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).ListSnippets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_ListSnippets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).ListSnippets(ctx, req.(*ListStippetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RQ_GetSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).GetSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_GetSnippet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).GetSnippet(ctx, req.(*GetSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RQ_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RQServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RQ_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RQServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RQ_ServiceDesc is the grpc.ServiceDesc for RQ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RQ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RQ",
	HandlerType: (*RQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _RQ_Query_Handler,
		},
		{
			MethodName: "UploadSnippet",
			Handler:    _RQ_UploadSnippet_Handler,
		},
		{
			MethodName: "RenameSnippet",
			Handler:    _RQ_RenameSnippet_Handler,
		},
		{
			MethodName: "ListSnippets",
			Handler:    _RQ_ListSnippets_Handler,
		},
		{
			MethodName: "GetSnippet",
			Handler:    _RQ_GetSnippet_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _RQ_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rq/v1/service.proto",
}
