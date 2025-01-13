// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: basic/database/protos/database/database.proto

package database

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Database_Uno_CreateUser_FullMethodName     = "/susubot.basic.database.database/Uno_CreateUser"
	Database_Uno_GetUser_FullMethodName        = "/susubot.basic.database.database/Uno_GetUser"
	Database_Uno_VerifyUser_FullMethodName     = "/susubot.basic.database.database/Uno_VerifyUser"
	Database_Uno_UpdateUser_FullMethodName     = "/susubot.basic.database.database/Uno_UpdateUser"
	Database_Uno_ChangePassword_FullMethodName = "/susubot.basic.database.database/Uno_ChangePassword"
)

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	Uno_CreateUser(ctx context.Context, in *Uno_CreateUserRequest, opts ...grpc.CallOption) (*Uno_CreateUserResponse, error)
	Uno_GetUser(ctx context.Context, in *Uno_GetUserRequest, opts ...grpc.CallOption) (*Uno_GetUserResponse, error)
	Uno_VerifyUser(ctx context.Context, in *Uno_VerifyUserRequest, opts ...grpc.CallOption) (*Uno_VerifyUserResponse, error)
	Uno_UpdateUser(ctx context.Context, in *Uno_UpdateUserRequest, opts ...grpc.CallOption) (*Uno_UpdateUserResponse, error)
	Uno_ChangePassword(ctx context.Context, in *Uno_ChangePasswordRequest, opts ...grpc.CallOption) (*Uno_ChangePasswordResponse, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Uno_CreateUser(ctx context.Context, in *Uno_CreateUserRequest, opts ...grpc.CallOption) (*Uno_CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Uno_CreateUserResponse)
	err := c.cc.Invoke(ctx, Database_Uno_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Uno_GetUser(ctx context.Context, in *Uno_GetUserRequest, opts ...grpc.CallOption) (*Uno_GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Uno_GetUserResponse)
	err := c.cc.Invoke(ctx, Database_Uno_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Uno_VerifyUser(ctx context.Context, in *Uno_VerifyUserRequest, opts ...grpc.CallOption) (*Uno_VerifyUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Uno_VerifyUserResponse)
	err := c.cc.Invoke(ctx, Database_Uno_VerifyUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Uno_UpdateUser(ctx context.Context, in *Uno_UpdateUserRequest, opts ...grpc.CallOption) (*Uno_UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Uno_UpdateUserResponse)
	err := c.cc.Invoke(ctx, Database_Uno_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Uno_ChangePassword(ctx context.Context, in *Uno_ChangePasswordRequest, opts ...grpc.CallOption) (*Uno_ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Uno_ChangePasswordResponse)
	err := c.cc.Invoke(ctx, Database_Uno_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility.
type DatabaseServer interface {
	Uno_CreateUser(context.Context, *Uno_CreateUserRequest) (*Uno_CreateUserResponse, error)
	Uno_GetUser(context.Context, *Uno_GetUserRequest) (*Uno_GetUserResponse, error)
	Uno_VerifyUser(context.Context, *Uno_VerifyUserRequest) (*Uno_VerifyUserResponse, error)
	Uno_UpdateUser(context.Context, *Uno_UpdateUserRequest) (*Uno_UpdateUserResponse, error)
	Uno_ChangePassword(context.Context, *Uno_ChangePasswordRequest) (*Uno_ChangePasswordResponse, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatabaseServer struct{}

func (UnimplementedDatabaseServer) Uno_CreateUser(context.Context, *Uno_CreateUserRequest) (*Uno_CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uno_CreateUser not implemented")
}
func (UnimplementedDatabaseServer) Uno_GetUser(context.Context, *Uno_GetUserRequest) (*Uno_GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uno_GetUser not implemented")
}
func (UnimplementedDatabaseServer) Uno_VerifyUser(context.Context, *Uno_VerifyUserRequest) (*Uno_VerifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uno_VerifyUser not implemented")
}
func (UnimplementedDatabaseServer) Uno_UpdateUser(context.Context, *Uno_UpdateUserRequest) (*Uno_UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uno_UpdateUser not implemented")
}
func (UnimplementedDatabaseServer) Uno_ChangePassword(context.Context, *Uno_ChangePasswordRequest) (*Uno_ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uno_ChangePassword not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}
func (UnimplementedDatabaseServer) testEmbeddedByValue()                  {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	// If the following call pancis, it indicates UnimplementedDatabaseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_Uno_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uno_CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Uno_CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Uno_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Uno_CreateUser(ctx, req.(*Uno_CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Uno_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uno_GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Uno_GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Uno_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Uno_GetUser(ctx, req.(*Uno_GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Uno_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uno_VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Uno_VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Uno_VerifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Uno_VerifyUser(ctx, req.(*Uno_VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Uno_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uno_UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Uno_UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Uno_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Uno_UpdateUser(ctx, req.(*Uno_UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Uno_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uno_ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Uno_ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Uno_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Uno_ChangePassword(ctx, req.(*Uno_ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "susubot.basic.database.database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Uno_CreateUser",
			Handler:    _Database_Uno_CreateUser_Handler,
		},
		{
			MethodName: "Uno_GetUser",
			Handler:    _Database_Uno_GetUser_Handler,
		},
		{
			MethodName: "Uno_VerifyUser",
			Handler:    _Database_Uno_VerifyUser_Handler,
		},
		{
			MethodName: "Uno_UpdateUser",
			Handler:    _Database_Uno_UpdateUser_Handler,
		},
		{
			MethodName: "Uno_ChangePassword",
			Handler:    _Database_Uno_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic/database/protos/database/database.proto",
}
