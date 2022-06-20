// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// FleetDashServiceClient is the client API for FleetDashService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FleetDashServiceClient interface {
	PostEveLogEventBatch(ctx context.Context, in *EveLogEventBatch, opts ...grpc.CallOption) (*BatchedEveLogEventResponse, error)
}

type fleetDashServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFleetDashServiceClient(cc grpc.ClientConnInterface) FleetDashServiceClient {
	return &fleetDashServiceClient{cc}
}

func (c *fleetDashServiceClient) PostEveLogEventBatch(ctx context.Context, in *EveLogEventBatch, opts ...grpc.CallOption) (*BatchedEveLogEventResponse, error) {
	out := new(BatchedEveLogEventResponse)
	err := c.cc.Invoke(ctx, "/FleetDashService/PostEveLogEventBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FleetDashServiceServer is the server API for FleetDashService service.
// All implementations must embed UnimplementedFleetDashServiceServer
// for forward compatibility
type FleetDashServiceServer interface {
	PostEveLogEventBatch(context.Context, *EveLogEventBatch) (*BatchedEveLogEventResponse, error)
	mustEmbedUnimplementedFleetDashServiceServer()
}

// UnimplementedFleetDashServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFleetDashServiceServer struct {
}

func (UnimplementedFleetDashServiceServer) PostEveLogEventBatch(context.Context, *EveLogEventBatch) (*BatchedEveLogEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostEveLogEventBatch not implemented")
}
func (UnimplementedFleetDashServiceServer) mustEmbedUnimplementedFleetDashServiceServer() {}

// UnsafeFleetDashServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FleetDashServiceServer will
// result in compilation errors.
type UnsafeFleetDashServiceServer interface {
	mustEmbedUnimplementedFleetDashServiceServer()
}

func RegisterFleetDashServiceServer(s grpc.ServiceRegistrar, srv FleetDashServiceServer) {
	s.RegisterService(&FleetDashService_ServiceDesc, srv)
}

func _FleetDashService_PostEveLogEventBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EveLogEventBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FleetDashServiceServer).PostEveLogEventBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FleetDashService/PostEveLogEventBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FleetDashServiceServer).PostEveLogEventBatch(ctx, req.(*EveLogEventBatch))
	}
	return interceptor(ctx, in, info, handler)
}

// FleetDashService_ServiceDesc is the grpc.ServiceDesc for FleetDashService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FleetDashService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FleetDashService",
	HandlerType: (*FleetDashServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostEveLogEventBatch",
			Handler:    _FleetDashService_PostEveLogEventBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fleetdash.proto",
}