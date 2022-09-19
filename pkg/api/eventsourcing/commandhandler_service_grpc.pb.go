// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: api/eventsourcing/commandhandler_service.proto

package eventsourcing

import (
	context "context"
	commands "github.com/finleap-connect/monoskope/pkg/api/eventsourcing/commands"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommandHandlerClient is the client API for CommandHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandHandlerClient interface {
	// Execute executes a command.
	Execute(ctx context.Context, in *commands.Command, opts ...grpc.CallOption) (*CommandReply, error)
}

type commandHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandHandlerClient(cc grpc.ClientConnInterface) CommandHandlerClient {
	return &commandHandlerClient{cc}
}

func (c *commandHandlerClient) Execute(ctx context.Context, in *commands.Command, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := c.cc.Invoke(ctx, "/eventsourcing.CommandHandler/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandHandlerServer is the server API for CommandHandler service.
// All implementations must embed UnimplementedCommandHandlerServer
// for forward compatibility
type CommandHandlerServer interface {
	// Execute executes a command.
	Execute(context.Context, *commands.Command) (*CommandReply, error)
	mustEmbedUnimplementedCommandHandlerServer()
}

// UnimplementedCommandHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedCommandHandlerServer struct {
}

func (UnimplementedCommandHandlerServer) Execute(context.Context, *commands.Command) (*CommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedCommandHandlerServer) mustEmbedUnimplementedCommandHandlerServer() {}

// UnsafeCommandHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandHandlerServer will
// result in compilation errors.
type UnsafeCommandHandlerServer interface {
	mustEmbedUnimplementedCommandHandlerServer()
}

func RegisterCommandHandlerServer(s grpc.ServiceRegistrar, srv CommandHandlerServer) {
	s.RegisterService(&CommandHandler_ServiceDesc, srv)
}

func _CommandHandler_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandHandlerServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventsourcing.CommandHandler/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandHandlerServer).Execute(ctx, req.(*commands.Command))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandHandler_ServiceDesc is the grpc.ServiceDesc for CommandHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventsourcing.CommandHandler",
	HandlerType: (*CommandHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _CommandHandler_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/eventsourcing/commandhandler_service.proto",
}
