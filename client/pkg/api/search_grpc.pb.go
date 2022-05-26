// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: search.proto

package api

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

// SearchingClient is the client API for Searching service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchingClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchingClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchingClient(cc grpc.ClientConnInterface) SearchingClient {
	return &searchingClient{cc}
}

func (c *searchingClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/api.Searching/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchingServer is the server API for Searching service.
// All implementations should embed UnimplementedSearchingServer
// for forward compatibility
type SearchingServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchingServer should be embedded to have forward compatible implementations.
type UnimplementedSearchingServer struct {
}

func (UnimplementedSearchingServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeSearchingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchingServer will
// result in compilation errors.
type UnsafeSearchingServer interface {
	mustEmbedUnimplementedSearchingServer()
}

func RegisterSearchingServer(s grpc.ServiceRegistrar, srv SearchingServer) {
	s.RegisterService(&Searching_ServiceDesc, srv)
}

func _Searching_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchingServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Searching/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchingServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Searching_ServiceDesc is the grpc.ServiceDesc for Searching service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Searching_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Searching",
	HandlerType: (*SearchingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Searching_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}