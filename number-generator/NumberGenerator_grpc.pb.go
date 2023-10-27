// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: NumberGenerator.proto

package number_generator

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
	NumberGenerator_Generate_FullMethodName = "/NumberGenerator/Generate"
)

// NumberGeneratorClient is the client API for NumberGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NumberGeneratorClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type numberGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberGeneratorClient(cc grpc.ClientConnInterface) NumberGeneratorClient {
	return &numberGeneratorClient{cc}
}

func (c *numberGeneratorClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, NumberGenerator_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberGeneratorServer is the server API for NumberGenerator service.
// All implementations must embed UnimplementedNumberGeneratorServer
// for forward compatibility
type NumberGeneratorServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedNumberGeneratorServer()
}

// UnimplementedNumberGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedNumberGeneratorServer struct {
}

func (UnimplementedNumberGeneratorServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedNumberGeneratorServer) mustEmbedUnimplementedNumberGeneratorServer() {}

// UnsafeNumberGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberGeneratorServer will
// result in compilation errors.
type UnsafeNumberGeneratorServer interface {
	mustEmbedUnimplementedNumberGeneratorServer()
}

func RegisterNumberGeneratorServer(s grpc.ServiceRegistrar, srv NumberGeneratorServer) {
	s.RegisterService(&NumberGenerator_ServiceDesc, srv)
}

func _NumberGenerator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberGeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberGenerator_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberGeneratorServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NumberGenerator_ServiceDesc is the grpc.ServiceDesc for NumberGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NumberGenerator",
	HandlerType: (*NumberGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _NumberGenerator_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "NumberGenerator.proto",
}
