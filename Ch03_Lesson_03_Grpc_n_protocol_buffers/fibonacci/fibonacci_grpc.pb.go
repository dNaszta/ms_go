// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FibonacciClient is the client API for Fibonacci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibonacciClient interface {
	Calculate(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciReply, error)
}

type fibonacciClient struct {
	cc grpc.ClientConnInterface
}

func NewFibonacciClient(cc grpc.ClientConnInterface) FibonacciClient {
	return &fibonacciClient{cc}
}

func (c *fibonacciClient) Calculate(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciReply, error) {
	out := new(FibonacciReply)
	err := c.cc.Invoke(ctx, "/fibonacci.Fibonacci/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibonacciServer is the server API for Fibonacci service.
// All implementations must embed UnimplementedFibonacciServer
// for forward compatibility
type FibonacciServer interface {
	Calculate(context.Context, *FibonacciRequest) (*FibonacciReply, error)
	mustEmbedUnimplementedFibonacciServer()
}

// UnimplementedFibonacciServer must be embedded to have forward compatible implementations.
type UnimplementedFibonacciServer struct {
}

func (UnimplementedFibonacciServer) Calculate(context.Context, *FibonacciRequest) (*FibonacciReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedFibonacciServer) mustEmbedUnimplementedFibonacciServer() {}

// UnsafeFibonacciServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibonacciServer will
// result in compilation errors.
type UnsafeFibonacciServer interface {
	mustEmbedUnimplementedFibonacciServer()
}

func RegisterFibonacciServer(s grpc.ServiceRegistrar, srv FibonacciServer) {
	s.RegisterService(&_Fibonacci_serviceDesc, srv)
}

func _Fibonacci_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibonacciRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibonacci.Fibonacci/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServer).Calculate(ctx, req.(*FibonacciRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fibonacci_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fibonacci.Fibonacci",
	HandlerType: (*FibonacciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Fibonacci_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fibonacci/fibonacci.proto",
}
