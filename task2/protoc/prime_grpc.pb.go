// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoc

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

// NumServiceClient is the client API for NumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NumServiceClient interface {
	Number(ctx context.Context, opts ...grpc.CallOption) (NumService_NumberClient, error)
}

type numServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNumServiceClient(cc grpc.ClientConnInterface) NumServiceClient {
	return &numServiceClient{cc}
}

func (c *numServiceClient) Number(ctx context.Context, opts ...grpc.CallOption) (NumService_NumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &NumService_ServiceDesc.Streams[0], "/protoc.NumService/Number", opts...)
	if err != nil {
		return nil, err
	}
	x := &numServiceNumberClient{stream}
	return x, nil
}

type NumService_NumberClient interface {
	Send(*NumberRequest) error
	Recv() (*NumberResponse, error)
	grpc.ClientStream
}

type numServiceNumberClient struct {
	grpc.ClientStream
}

func (x *numServiceNumberClient) Send(m *NumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *numServiceNumberClient) Recv() (*NumberResponse, error) {
	m := new(NumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumServiceServer is the server API for NumService service.
// All implementations must embed UnimplementedNumServiceServer
// for forward compatibility
type NumServiceServer interface {
	Number(NumService_NumberServer) error
	mustEmbedUnimplementedNumServiceServer()
}

// UnimplementedNumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNumServiceServer struct {
}

func (UnimplementedNumServiceServer) Number(NumService_NumberServer) error {
	return status.Errorf(codes.Unimplemented, "method Number not implemented")
}
func (UnimplementedNumServiceServer) mustEmbedUnimplementedNumServiceServer() {}

// UnsafeNumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumServiceServer will
// result in compilation errors.
type UnsafeNumServiceServer interface {
	mustEmbedUnimplementedNumServiceServer()
}

func RegisterNumServiceServer(s grpc.ServiceRegistrar, srv NumServiceServer) {
	s.RegisterService(&NumService_ServiceDesc, srv)
}

func _NumService_Number_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NumServiceServer).Number(&numServiceNumberServer{stream})
}

type NumService_NumberServer interface {
	Send(*NumberResponse) error
	Recv() (*NumberRequest, error)
	grpc.ServerStream
}

type numServiceNumberServer struct {
	grpc.ServerStream
}

func (x *numServiceNumberServer) Send(m *NumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *numServiceNumberServer) Recv() (*NumberRequest, error) {
	m := new(NumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumService_ServiceDesc is the grpc.ServiceDesc for NumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoc.NumService",
	HandlerType: (*NumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Number",
			Handler:       _NumService_Number_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "task2/protoc/prime.proto",
}
