// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: add.proto

package grpcpb

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
	DocumentValidation_ValidateDocument_FullMethodName = "/grpc.DocumentValidation/ValidateDocument"
)

// DocumentValidationClient is the client API for DocumentValidation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentValidationClient interface {
	ValidateDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentResponse, error)
}

type documentValidationClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentValidationClient(cc grpc.ClientConnInterface) DocumentValidationClient {
	return &documentValidationClient{cc}
}

func (c *documentValidationClient) ValidateDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentResponse, error) {
	out := new(DocumentResponse)
	err := c.cc.Invoke(ctx, DocumentValidation_ValidateDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentValidationServer is the server API for DocumentValidation service.
// All implementations must embed UnimplementedDocumentValidationServer
// for forward compatibility
type DocumentValidationServer interface {
	ValidateDocument(context.Context, *DocumentRequest) (*DocumentResponse, error)
}

// UnimplementedDocumentValidationServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentValidationServer struct {
}

func (UnimplementedDocumentValidationServer) ValidateDocument(context.Context, *DocumentRequest) (*DocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDocument not implemented")
}
func (UnimplementedDocumentValidationServer) mustEmbedUnimplementedDocumentValidationServer() {}

// UnsafeDocumentValidationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentValidationServer will
// result in compilation errors.

func RegisterDocumentValidationServer(s grpc.ServiceRegistrar, srv DocumentValidationServer) {
	s.RegisterService(&DocumentValidation_ServiceDesc, srv)
}

func _DocumentValidation_ValidateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentValidationServer).ValidateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentValidation_ValidateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentValidationServer).ValidateDocument(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentValidation_ServiceDesc is the grpc.ServiceDesc for DocumentValidation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentValidation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DocumentValidation",
	HandlerType: (*DocumentValidationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateDocument",
			Handler:    _DocumentValidation_ValidateDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add.proto",
}
