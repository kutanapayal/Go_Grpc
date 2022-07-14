// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: a.proto

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

// AServiceClient is the client API for AService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AServiceClient interface {
	A(ctx context.Context, in *ARequest, opts ...grpc.CallOption) (*AResponse, error)
}

type aServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAServiceClient(cc grpc.ClientConnInterface) AServiceClient {
	return &aServiceClient{cc}
}

func (c *aServiceClient) A(ctx context.Context, in *ARequest, opts ...grpc.CallOption) (*AResponse, error) {
	out := new(AResponse)
	err := c.cc.Invoke(ctx, "/a.aService/a", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AServiceServer is the server API for AService service.
// All implementations must embed UnimplementedAServiceServer
// for forward compatibility
type AServiceServer interface {
	A(context.Context, *ARequest) (*AResponse, error)
	mustEmbedUnimplementedAServiceServer()
}

// UnimplementedAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAServiceServer struct {
}

func (UnimplementedAServiceServer) A(context.Context, *ARequest) (*AResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method A not implemented")
}
func (UnimplementedAServiceServer) mustEmbedUnimplementedAServiceServer() {}

// UnsafeAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AServiceServer will
// result in compilation errors.
type UnsafeAServiceServer interface {
	mustEmbedUnimplementedAServiceServer()
}

func RegisterAServiceServer(s grpc.ServiceRegistrar, srv AServiceServer) {
	s.RegisterService(&AService_ServiceDesc, srv)
}

func _AService_A_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AServiceServer).A(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/a.aService/a",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AServiceServer).A(ctx, req.(*ARequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AService_ServiceDesc is the grpc.ServiceDesc for AService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "a.aService",
	HandlerType: (*AServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "a",
			Handler:    _AService_A_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "a.proto",
}

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeServiceClient interface {
	Primedivions(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_PrimedivionsClient, error)
}

type primeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeServiceClient(cc grpc.ClientConnInterface) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) Primedivions(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_PrimedivionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeService_ServiceDesc.Streams[0], "/a.primeService/primedivions", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServicePrimedivionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeService_PrimedivionsClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type primeServicePrimedivionsClient struct {
	grpc.ClientStream
}

func (x *primeServicePrimedivionsClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeServiceServer is the server API for PrimeService service.
// All implementations must embed UnimplementedPrimeServiceServer
// for forward compatibility
type PrimeServiceServer interface {
	Primedivions(*PrimeRequest, PrimeService_PrimedivionsServer) error
	mustEmbedUnimplementedPrimeServiceServer()
}

// UnimplementedPrimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (UnimplementedPrimeServiceServer) Primedivions(*PrimeRequest, PrimeService_PrimedivionsServer) error {
	return status.Errorf(codes.Unimplemented, "method Primedivions not implemented")
}
func (UnimplementedPrimeServiceServer) mustEmbedUnimplementedPrimeServiceServer() {}

// UnsafePrimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeServiceServer will
// result in compilation errors.
type UnsafePrimeServiceServer interface {
	mustEmbedUnimplementedPrimeServiceServer()
}

func RegisterPrimeServiceServer(s grpc.ServiceRegistrar, srv PrimeServiceServer) {
	s.RegisterService(&PrimeService_ServiceDesc, srv)
}

func _PrimeService_Primedivions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeServiceServer).Primedivions(m, &primeServicePrimedivionsServer{stream})
}

type PrimeService_PrimedivionsServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type primeServicePrimedivionsServer struct {
	grpc.ServerStream
}

func (x *primeServicePrimedivionsServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PrimeService_ServiceDesc is the grpc.ServiceDesc for PrimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "a.primeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "primedivions",
			Handler:       _PrimeService_Primedivions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "a.proto",
}

// AvgServiceClient is the client API for AvgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvgServiceClient interface {
	FindAvg(ctx context.Context, opts ...grpc.CallOption) (AvgService_FindAvgClient, error)
}

type avgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvgServiceClient(cc grpc.ClientConnInterface) AvgServiceClient {
	return &avgServiceClient{cc}
}

func (c *avgServiceClient) FindAvg(ctx context.Context, opts ...grpc.CallOption) (AvgService_FindAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &AvgService_ServiceDesc.Streams[0], "/a.avgService/findAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &avgServiceFindAvgClient{stream}
	return x, nil
}

type AvgService_FindAvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type avgServiceFindAvgClient struct {
	grpc.ClientStream
}

func (x *avgServiceFindAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avgServiceFindAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvgServiceServer is the server API for AvgService service.
// All implementations must embed UnimplementedAvgServiceServer
// for forward compatibility
type AvgServiceServer interface {
	FindAvg(AvgService_FindAvgServer) error
	mustEmbedUnimplementedAvgServiceServer()
}

// UnimplementedAvgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvgServiceServer struct {
}

func (UnimplementedAvgServiceServer) FindAvg(AvgService_FindAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAvg not implemented")
}
func (UnimplementedAvgServiceServer) mustEmbedUnimplementedAvgServiceServer() {}

// UnsafeAvgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvgServiceServer will
// result in compilation errors.
type UnsafeAvgServiceServer interface {
	mustEmbedUnimplementedAvgServiceServer()
}

func RegisterAvgServiceServer(s grpc.ServiceRegistrar, srv AvgServiceServer) {
	s.RegisterService(&AvgService_ServiceDesc, srv)
}

func _AvgService_FindAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvgServiceServer).FindAvg(&avgServiceFindAvgServer{stream})
}

type AvgService_FindAvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type avgServiceFindAvgServer struct {
	grpc.ServerStream
}

func (x *avgServiceFindAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avgServiceFindAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvgService_ServiceDesc is the grpc.ServiceDesc for AvgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "a.avgService",
	HandlerType: (*AvgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "findAvg",
			Handler:       _AvgService_FindAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "a.proto",
}

// MaxServiceClient is the client API for MaxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaxServiceClient interface {
	FindMax(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaxClient, error)
}

type maxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaxServiceClient(cc grpc.ClientConnInterface) MaxServiceClient {
	return &maxServiceClient{cc}
}

func (c *maxServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &MaxService_ServiceDesc.Streams[0], "/a.maxService/findMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxServiceFindMaxClient{stream}
	return x, nil
}

type MaxService_FindMaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type maxServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *maxServiceFindMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxServiceFindMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxServiceServer is the server API for MaxService service.
// All implementations must embed UnimplementedMaxServiceServer
// for forward compatibility
type MaxServiceServer interface {
	FindMax(MaxService_FindMaxServer) error
	mustEmbedUnimplementedMaxServiceServer()
}

// UnimplementedMaxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMaxServiceServer struct {
}

func (UnimplementedMaxServiceServer) FindMax(MaxService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (UnimplementedMaxServiceServer) mustEmbedUnimplementedMaxServiceServer() {}

// UnsafeMaxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaxServiceServer will
// result in compilation errors.
type UnsafeMaxServiceServer interface {
	mustEmbedUnimplementedMaxServiceServer()
}

func RegisterMaxServiceServer(s grpc.ServiceRegistrar, srv MaxServiceServer) {
	s.RegisterService(&MaxService_ServiceDesc, srv)
}

func _MaxService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxServiceServer).FindMax(&maxServiceFindMaxServer{stream})
}

type MaxService_FindMaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type maxServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *maxServiceFindMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxServiceFindMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxService_ServiceDesc is the grpc.ServiceDesc for MaxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "a.maxService",
	HandlerType: (*MaxServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "findMax",
			Handler:       _MaxService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "a.proto",
}

// SqrtServiceClient is the client API for SqrtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqrtServiceClient interface {
	FindSqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
	Findabs(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type sqrtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqrtServiceClient(cc grpc.ClientConnInterface) SqrtServiceClient {
	return &sqrtServiceClient{cc}
}

func (c *sqrtServiceClient) FindSqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/a.sqrtService/findSqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqrtServiceClient) Findabs(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/a.sqrtService/findabs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqrtServiceServer is the server API for SqrtService service.
// All implementations must embed UnimplementedSqrtServiceServer
// for forward compatibility
type SqrtServiceServer interface {
	FindSqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	Findabs(context.Context, *SqrtRequest) (*SqrtResponse, error)
	mustEmbedUnimplementedSqrtServiceServer()
}

// UnimplementedSqrtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSqrtServiceServer struct {
}

func (UnimplementedSqrtServiceServer) FindSqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSqrt not implemented")
}
func (UnimplementedSqrtServiceServer) Findabs(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Findabs not implemented")
}
func (UnimplementedSqrtServiceServer) mustEmbedUnimplementedSqrtServiceServer() {}

// UnsafeSqrtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqrtServiceServer will
// result in compilation errors.
type UnsafeSqrtServiceServer interface {
	mustEmbedUnimplementedSqrtServiceServer()
}

func RegisterSqrtServiceServer(s grpc.ServiceRegistrar, srv SqrtServiceServer) {
	s.RegisterService(&SqrtService_ServiceDesc, srv)
}

func _SqrtService_FindSqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqrtServiceServer).FindSqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/a.sqrtService/findSqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqrtServiceServer).FindSqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SqrtService_Findabs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqrtServiceServer).Findabs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/a.sqrtService/findabs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqrtServiceServer).Findabs(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SqrtService_ServiceDesc is the grpc.ServiceDesc for SqrtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SqrtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "a.sqrtService",
	HandlerType: (*SqrtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findSqrt",
			Handler:    _SqrtService_FindSqrt_Handler,
		},
		{
			MethodName: "findabs",
			Handler:    _SqrtService_Findabs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "a.proto",
}