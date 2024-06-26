// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: volumeGroupSnapshot.proto

package volumegroupsnapshot

import (
	common "github.com/dell/dell-csi-extensions/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateVolumeGroupSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// volumeIDs to be snapped
	SourceVolumeIDs []string `protobuf:"bytes,1,rep,name=SourceVolumeIDs,proto3" json:"SourceVolumeIDs,omitempty"`
	// name of snapshot group
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	// parameters map from VolumeGroupSnapshot instance
	Parameters map[string]string `protobuf:"bytes,5,rep,name=Parameters,proto3" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateVolumeGroupSnapshotRequest) Reset() {
	*x = CreateVolumeGroupSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volumeGroupSnapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeGroupSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeGroupSnapshotRequest) ProtoMessage() {}

func (x *CreateVolumeGroupSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_volumeGroupSnapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeGroupSnapshotRequest.ProtoReflect.Descriptor instead.
func (*CreateVolumeGroupSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_volumeGroupSnapshot_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVolumeGroupSnapshotRequest) GetSourceVolumeIDs() []string {
	if x != nil {
		return x.SourceVolumeIDs
	}
	return nil
}

func (x *CreateVolumeGroupSnapshotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVolumeGroupSnapshotRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateVolumeGroupSnapshotRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type CreateVolumeGroupSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// snapshots in group
	Snapshots []*Snapshot `protobuf:"bytes,1,rep,name=Snapshots,proto3" json:"Snapshots,omitempty"`
	// snapshot group csi id on array
	SnapshotGroupID string `protobuf:"bytes,2,opt,name=SnapshotGroupID,proto3" json:"SnapshotGroupID,omitempty"`
	// time VGS was created
	CreationTime int64 `protobuf:"varint,3,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
}

func (x *CreateVolumeGroupSnapshotResponse) Reset() {
	*x = CreateVolumeGroupSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volumeGroupSnapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeGroupSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeGroupSnapshotResponse) ProtoMessage() {}

func (x *CreateVolumeGroupSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_volumeGroupSnapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeGroupSnapshotResponse.ProtoReflect.Descriptor instead.
func (*CreateVolumeGroupSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_volumeGroupSnapshot_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVolumeGroupSnapshotResponse) GetSnapshots() []*Snapshot {
	if x != nil {
		return x.Snapshots
	}
	return nil
}

func (x *CreateVolumeGroupSnapshotResponse) GetSnapshotGroupID() string {
	if x != nil {
		return x.SnapshotGroupID
	}
	return ""
}

func (x *CreateVolumeGroupSnapshotResponse) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Size of the source volume in bytes
	CapacityBytes int64 `protobuf:"varint,1,opt,name=capacity_bytes,json=capacityBytes,proto3" json:"capacity_bytes,omitempty"`
	// Snapshot ID - CSI snapshot ID. Should uniquely identify the snapshot for the driver.
	SnapId string `protobuf:"bytes,2,opt,name=snap_id,json=snapId,proto3" json:"snap_id,omitempty"`
	// ID of source volume
	SourceId string `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// bool value to determine if this snap is ready for use
	ReadyToUse bool `protobuf:"varint,4,opt,name=readyToUse,proto3" json:"readyToUse,omitempty"`
	// time snapshot was created
	CreationTime int64 `protobuf:"varint,5,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	// name of snapshot found in array
	Name string `protobuf:"bytes,6,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volumeGroupSnapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_volumeGroupSnapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_volumeGroupSnapshot_proto_rawDescGZIP(), []int{2}
}

func (x *Snapshot) GetCapacityBytes() int64 {
	if x != nil {
		return x.CapacityBytes
	}
	return 0
}

func (x *Snapshot) GetSnapId() string {
	if x != nil {
		return x.SnapId
	}
	return ""
}

func (x *Snapshot) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *Snapshot) GetReadyToUse() bool {
	if x != nil {
		return x.ReadyToUse
	}
	return false
}

func (x *Snapshot) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *Snapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_volumeGroupSnapshot_proto protoreflect.FileDescriptor

var file_volumeGroupSnapshot_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x02, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x44, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb1, 0x01, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x09, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6e, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6e, 0x61, 0x70, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x86, 0x02, 0x0a, 0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x5a, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x92, 0x01, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x38, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x39, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x2e, 0x2f, 0x3b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volumeGroupSnapshot_proto_rawDescOnce sync.Once
	file_volumeGroupSnapshot_proto_rawDescData = file_volumeGroupSnapshot_proto_rawDesc
)

func file_volumeGroupSnapshot_proto_rawDescGZIP() []byte {
	file_volumeGroupSnapshot_proto_rawDescOnce.Do(func() {
		file_volumeGroupSnapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_volumeGroupSnapshot_proto_rawDescData)
	})
	return file_volumeGroupSnapshot_proto_rawDescData
}

var file_volumeGroupSnapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_volumeGroupSnapshot_proto_goTypes = []interface{}{
	(*CreateVolumeGroupSnapshotRequest)(nil),  // 0: volumegroupsnapshot.v1.CreateVolumeGroupSnapshotRequest
	(*CreateVolumeGroupSnapshotResponse)(nil), // 1: volumegroupsnapshot.v1.CreateVolumeGroupSnapshotResponse
	(*Snapshot)(nil),                       // 2: volumegroupsnapshot.v1.Snapshot
	nil,                                    // 3: volumegroupsnapshot.v1.CreateVolumeGroupSnapshotRequest.ParametersEntry
	(*common.ProbeControllerRequest)(nil),  // 4: common.v1.ProbeControllerRequest
	(*common.ProbeControllerResponse)(nil), // 5: common.v1.ProbeControllerResponse
}
var file_volumeGroupSnapshot_proto_depIdxs = []int32{
	3, // 0: volumegroupsnapshot.v1.CreateVolumeGroupSnapshotRequest.Parameters:type_name -> volumegroupsnapshot.v1.CreateVolumeGroupSnapshotRequest.ParametersEntry
	2, // 1: volumegroupsnapshot.v1.CreateVolumeGroupSnapshotResponse.Snapshots:type_name -> volumegroupsnapshot.v1.Snapshot
	4, // 2: volumegroupsnapshot.v1.VolumeGroupSnapshot.ProbeController:input_type -> common.v1.ProbeControllerRequest
	0, // 3: volumegroupsnapshot.v1.VolumeGroupSnapshot.CreateVolumeGroupSnapshot:input_type -> volumegroupsnapshot.v1.CreateVolumeGroupSnapshotRequest
	5, // 4: volumegroupsnapshot.v1.VolumeGroupSnapshot.ProbeController:output_type -> common.v1.ProbeControllerResponse
	1, // 5: volumegroupsnapshot.v1.VolumeGroupSnapshot.CreateVolumeGroupSnapshot:output_type -> volumegroupsnapshot.v1.CreateVolumeGroupSnapshotResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_volumeGroupSnapshot_proto_init() }
func file_volumeGroupSnapshot_proto_init() {
	if File_volumeGroupSnapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volumeGroupSnapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeGroupSnapshotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_volumeGroupSnapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeGroupSnapshotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_volumeGroupSnapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_volumeGroupSnapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_volumeGroupSnapshot_proto_goTypes,
		DependencyIndexes: file_volumeGroupSnapshot_proto_depIdxs,
		MessageInfos:      file_volumeGroupSnapshot_proto_msgTypes,
	}.Build()
	File_volumeGroupSnapshot_proto = out.File
	file_volumeGroupSnapshot_proto_rawDesc = nil
	file_volumeGroupSnapshot_proto_goTypes = nil
	file_volumeGroupSnapshot_proto_depIdxs = nil
}
