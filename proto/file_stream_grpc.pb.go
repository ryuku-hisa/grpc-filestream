// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: file_stream.proto

package proto

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
	DataStreamHandler_DataStream_FullMethodName = "/DataStreamHandler/DataStream"
)

// DataStreamHandlerClient is the client API for DataStreamHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataStreamHandlerClient interface {
	DataStream(ctx context.Context, opts ...grpc.CallOption) (DataStreamHandler_DataStreamClient, error)
}

type dataStreamHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataStreamHandlerClient(cc grpc.ClientConnInterface) DataStreamHandlerClient {
	return &dataStreamHandlerClient{cc}
}

func (c *dataStreamHandlerClient) DataStream(ctx context.Context, opts ...grpc.CallOption) (DataStreamHandler_DataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataStreamHandler_ServiceDesc.Streams[0], DataStreamHandler_DataStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataStreamHandlerDataStreamClient{stream}
	return x, nil
}

type DataStreamHandler_DataStreamClient interface {
	Send(*DataStreamRequest) error
	CloseAndRecv() (*DataStreamResponse, error)
	grpc.ClientStream
}

type dataStreamHandlerDataStreamClient struct {
	grpc.ClientStream
}

func (x *dataStreamHandlerDataStreamClient) Send(m *DataStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataStreamHandlerDataStreamClient) CloseAndRecv() (*DataStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DataStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataStreamHandlerServer is the server API for DataStreamHandler service.
// All implementations must embed UnimplementedDataStreamHandlerServer
// for forward compatibility
type DataStreamHandlerServer interface {
	DataStream(DataStreamHandler_DataStreamServer) error
	mustEmbedUnimplementedDataStreamHandlerServer()
}

// UnimplementedDataStreamHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedDataStreamHandlerServer struct {
}

func (UnimplementedDataStreamHandlerServer) DataStream(DataStreamHandler_DataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DataStream not implemented")
}
func (UnimplementedDataStreamHandlerServer) mustEmbedUnimplementedDataStreamHandlerServer() {}

// UnsafeDataStreamHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataStreamHandlerServer will
// result in compilation errors.
type UnsafeDataStreamHandlerServer interface {
	mustEmbedUnimplementedDataStreamHandlerServer()
}

func RegisterDataStreamHandlerServer(s grpc.ServiceRegistrar, srv DataStreamHandlerServer) {
	s.RegisterService(&DataStreamHandler_ServiceDesc, srv)
}

func _DataStreamHandler_DataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataStreamHandlerServer).DataStream(&dataStreamHandlerDataStreamServer{stream})
}

type DataStreamHandler_DataStreamServer interface {
	SendAndClose(*DataStreamResponse) error
	Recv() (*DataStreamRequest, error)
	grpc.ServerStream
}

type dataStreamHandlerDataStreamServer struct {
	grpc.ServerStream
}

func (x *dataStreamHandlerDataStreamServer) SendAndClose(m *DataStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataStreamHandlerDataStreamServer) Recv() (*DataStreamRequest, error) {
	m := new(DataStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataStreamHandler_ServiceDesc is the grpc.ServiceDesc for DataStreamHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataStreamHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataStreamHandler",
	HandlerType: (*DataStreamHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataStream",
			Handler:       _DataStreamHandler_DataStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file_stream.proto",
}
