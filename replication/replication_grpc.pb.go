// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: replication.proto

package replication

import (
	context "context"
	common "github.com/dell/dell-csi-extensions/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Replication_ProbeController_FullMethodName                 = "/replication.v1.Replication/ProbeController"
	Replication_GetReplicationCapabilities_FullMethodName      = "/replication.v1.Replication/GetReplicationCapabilities"
	Replication_CreateStorageProtectionGroup_FullMethodName    = "/replication.v1.Replication/CreateStorageProtectionGroup"
	Replication_CreateRemoteVolume_FullMethodName              = "/replication.v1.Replication/CreateRemoteVolume"
	Replication_DeleteStorageProtectionGroup_FullMethodName    = "/replication.v1.Replication/DeleteStorageProtectionGroup"
	Replication_DeleteLocalVolume_FullMethodName               = "/replication.v1.Replication/DeleteLocalVolume"
	Replication_ExecuteAction_FullMethodName                   = "/replication.v1.Replication/ExecuteAction"
	Replication_GetStorageProtectionGroupStatus_FullMethodName = "/replication.v1.Replication/GetStorageProtectionGroupStatus"
)

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationClient interface {
	// ProbeController is used to verify if the CSI driver controller plugin is ready to service Replication APIs
	ProbeController(ctx context.Context, in *common.ProbeControllerRequest, opts ...grpc.CallOption) (*common.ProbeControllerResponse, error)
	// GetReplicationCapabilities is used to query CSI drivers for their supported replication capabilities
	GetReplicationCapabilities(ctx context.Context, in *GetReplicationCapabilityRequest, opts ...grpc.CallOption) (*GetReplicationCapabilityResponse, error)
	// CreateStorageProtectionGroup is used to create Storage Protection Group on array
	CreateStorageProtectionGroup(ctx context.Context, in *CreateStorageProtectionGroupRequest, opts ...grpc.CallOption) (*CreateStorageProtectionGroupResponse, error)
	// CreateRemoteVolume is used to create volume on remote array
	CreateRemoteVolume(ctx context.Context, in *CreateRemoteVolumeRequest, opts ...grpc.CallOption) (*CreateRemoteVolumeResponse, error)
	// DeleteStorageProtectionGroup is used to delete a Storage Protection group
	DeleteStorageProtectionGroup(ctx context.Context, in *DeleteStorageProtectionGroupRequest, opts ...grpc.CallOption) (*DeleteStorageProtectionGroupResponse, error)
	// DeleteLocalVolume is used to delete a replicated volume on the local array upon request from the remote replication controller
	DeleteLocalVolume(ctx context.Context, in *DeleteLocalVolumeRequest, opts ...grpc.CallOption) (*DeleteLocalVolumeResponse, error)
	// ExecuteAction is used to update the action to failover, testfailover, failback, suspend, etc.
	ExecuteAction(ctx context.Context, in *ExecuteActionRequest, opts ...grpc.CallOption) (*ExecuteActionResponse, error)
	// GetStorageProtectionGroupStatus is used to monitor the status of Storage Protection groups
	GetStorageProtectionGroupStatus(ctx context.Context, in *GetStorageProtectionGroupStatusRequest, opts ...grpc.CallOption) (*GetStorageProtectionGroupStatusResponse, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) ProbeController(ctx context.Context, in *common.ProbeControllerRequest, opts ...grpc.CallOption) (*common.ProbeControllerResponse, error) {
	out := new(common.ProbeControllerResponse)
	err := c.cc.Invoke(ctx, Replication_ProbeController_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) GetReplicationCapabilities(ctx context.Context, in *GetReplicationCapabilityRequest, opts ...grpc.CallOption) (*GetReplicationCapabilityResponse, error) {
	out := new(GetReplicationCapabilityResponse)
	err := c.cc.Invoke(ctx, Replication_GetReplicationCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) CreateStorageProtectionGroup(ctx context.Context, in *CreateStorageProtectionGroupRequest, opts ...grpc.CallOption) (*CreateStorageProtectionGroupResponse, error) {
	out := new(CreateStorageProtectionGroupResponse)
	err := c.cc.Invoke(ctx, Replication_CreateStorageProtectionGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) CreateRemoteVolume(ctx context.Context, in *CreateRemoteVolumeRequest, opts ...grpc.CallOption) (*CreateRemoteVolumeResponse, error) {
	out := new(CreateRemoteVolumeResponse)
	err := c.cc.Invoke(ctx, Replication_CreateRemoteVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) DeleteStorageProtectionGroup(ctx context.Context, in *DeleteStorageProtectionGroupRequest, opts ...grpc.CallOption) (*DeleteStorageProtectionGroupResponse, error) {
	out := new(DeleteStorageProtectionGroupResponse)
	err := c.cc.Invoke(ctx, Replication_DeleteStorageProtectionGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) DeleteLocalVolume(ctx context.Context, in *DeleteLocalVolumeRequest, opts ...grpc.CallOption) (*DeleteLocalVolumeResponse, error) {
	out := new(DeleteLocalVolumeResponse)
	err := c.cc.Invoke(ctx, Replication_DeleteLocalVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) ExecuteAction(ctx context.Context, in *ExecuteActionRequest, opts ...grpc.CallOption) (*ExecuteActionResponse, error) {
	out := new(ExecuteActionResponse)
	err := c.cc.Invoke(ctx, Replication_ExecuteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) GetStorageProtectionGroupStatus(ctx context.Context, in *GetStorageProtectionGroupStatusRequest, opts ...grpc.CallOption) (*GetStorageProtectionGroupStatusResponse, error) {
	out := new(GetStorageProtectionGroupStatusResponse)
	err := c.cc.Invoke(ctx, Replication_GetStorageProtectionGroupStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations should embed UnimplementedReplicationServer
// for forward compatibility
type ReplicationServer interface {
	// ProbeController is used to verify if the CSI driver controller plugin is ready to service Replication APIs
	ProbeController(context.Context, *common.ProbeControllerRequest) (*common.ProbeControllerResponse, error)
	// GetReplicationCapabilities is used to query CSI drivers for their supported replication capabilities
	GetReplicationCapabilities(context.Context, *GetReplicationCapabilityRequest) (*GetReplicationCapabilityResponse, error)
	// CreateStorageProtectionGroup is used to create Storage Protection Group on array
	CreateStorageProtectionGroup(context.Context, *CreateStorageProtectionGroupRequest) (*CreateStorageProtectionGroupResponse, error)
	// CreateRemoteVolume is used to create volume on remote array
	CreateRemoteVolume(context.Context, *CreateRemoteVolumeRequest) (*CreateRemoteVolumeResponse, error)
	// DeleteStorageProtectionGroup is used to delete a Storage Protection group
	DeleteStorageProtectionGroup(context.Context, *DeleteStorageProtectionGroupRequest) (*DeleteStorageProtectionGroupResponse, error)
	// DeleteLocalVolume is used to delete a replicated volume on the local array upon request from the remote replication controller
	DeleteLocalVolume(context.Context, *DeleteLocalVolumeRequest) (*DeleteLocalVolumeResponse, error)
	// ExecuteAction is used to update the action to failover, testfailover, failback, suspend, etc.
	ExecuteAction(context.Context, *ExecuteActionRequest) (*ExecuteActionResponse, error)
	// GetStorageProtectionGroupStatus is used to monitor the status of Storage Protection groups
	GetStorageProtectionGroupStatus(context.Context, *GetStorageProtectionGroupStatusRequest) (*GetStorageProtectionGroupStatusResponse, error)
}

// UnimplementedReplicationServer should be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (UnimplementedReplicationServer) ProbeController(context.Context, *common.ProbeControllerRequest) (*common.ProbeControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbeController not implemented")
}
func (UnimplementedReplicationServer) GetReplicationCapabilities(context.Context, *GetReplicationCapabilityRequest) (*GetReplicationCapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplicationCapabilities not implemented")
}
func (UnimplementedReplicationServer) CreateStorageProtectionGroup(context.Context, *CreateStorageProtectionGroupRequest) (*CreateStorageProtectionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorageProtectionGroup not implemented")
}
func (UnimplementedReplicationServer) CreateRemoteVolume(context.Context, *CreateRemoteVolumeRequest) (*CreateRemoteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRemoteVolume not implemented")
}
func (UnimplementedReplicationServer) DeleteStorageProtectionGroup(context.Context, *DeleteStorageProtectionGroupRequest) (*DeleteStorageProtectionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorageProtectionGroup not implemented")
}
func (UnimplementedReplicationServer) DeleteLocalVolume(context.Context, *DeleteLocalVolumeRequest) (*DeleteLocalVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocalVolume not implemented")
}
func (UnimplementedReplicationServer) ExecuteAction(context.Context, *ExecuteActionRequest) (*ExecuteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteAction not implemented")
}
func (UnimplementedReplicationServer) GetStorageProtectionGroupStatus(context.Context, *GetStorageProtectionGroupStatusRequest) (*GetStorageProtectionGroupStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageProtectionGroupStatus not implemented")
}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	s.RegisterService(&Replication_ServiceDesc, srv)
}

func _Replication_ProbeController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ProbeControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).ProbeController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_ProbeController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).ProbeController(ctx, req.(*common.ProbeControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_GetReplicationCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplicationCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).GetReplicationCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_GetReplicationCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).GetReplicationCapabilities(ctx, req.(*GetReplicationCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_CreateStorageProtectionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageProtectionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).CreateStorageProtectionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_CreateStorageProtectionGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).CreateStorageProtectionGroup(ctx, req.(*CreateStorageProtectionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_CreateRemoteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRemoteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).CreateRemoteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_CreateRemoteVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).CreateRemoteVolume(ctx, req.(*CreateRemoteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_DeleteStorageProtectionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageProtectionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).DeleteStorageProtectionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_DeleteStorageProtectionGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).DeleteStorageProtectionGroup(ctx, req.(*DeleteStorageProtectionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_DeleteLocalVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocalVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).DeleteLocalVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_DeleteLocalVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).DeleteLocalVolume(ctx, req.(*DeleteLocalVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_ExecuteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).ExecuteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_ExecuteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).ExecuteAction(ctx, req.(*ExecuteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_GetStorageProtectionGroupStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageProtectionGroupStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).GetStorageProtectionGroupStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_GetStorageProtectionGroupStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).GetStorageProtectionGroupStatus(ctx, req.(*GetStorageProtectionGroupStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Replication_ServiceDesc is the grpc.ServiceDesc for Replication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replication.v1.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProbeController",
			Handler:    _Replication_ProbeController_Handler,
		},
		{
			MethodName: "GetReplicationCapabilities",
			Handler:    _Replication_GetReplicationCapabilities_Handler,
		},
		{
			MethodName: "CreateStorageProtectionGroup",
			Handler:    _Replication_CreateStorageProtectionGroup_Handler,
		},
		{
			MethodName: "CreateRemoteVolume",
			Handler:    _Replication_CreateRemoteVolume_Handler,
		},
		{
			MethodName: "DeleteStorageProtectionGroup",
			Handler:    _Replication_DeleteStorageProtectionGroup_Handler,
		},
		{
			MethodName: "DeleteLocalVolume",
			Handler:    _Replication_DeleteLocalVolume_Handler,
		},
		{
			MethodName: "ExecuteAction",
			Handler:    _Replication_ExecuteAction_Handler,
		},
		{
			MethodName: "GetStorageProtectionGroupStatus",
			Handler:    _Replication_GetStorageProtectionGroupStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication.proto",
}
