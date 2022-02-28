// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/pb/protobuf-spec/geo_hash.proto

package pb

import (
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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []byte  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Distance  uint64  `protobuf:"varint,4,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Point) GetDistance() uint64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type GHCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Precision int32  `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *GHCreateRequest) Reset() {
	*x = GHCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHCreateRequest) ProtoMessage() {}

func (x *GHCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHCreateRequest.ProtoReflect.Descriptor instead.
func (*GHCreateRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{1}
}

func (x *GHCreateRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GHCreateRequest) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

type GHCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GHCreateResponse) Reset() {
	*x = GHCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHCreateResponse) ProtoMessage() {}

func (x *GHCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHCreateResponse.ProtoReflect.Descriptor instead.
func (*GHCreateResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{2}
}

func (x *GHCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GHDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GHDelRequest) Reset() {
	*x = GHDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHDelRequest) ProtoMessage() {}

func (x *GHDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHDelRequest.ProtoReflect.Descriptor instead.
func (*GHDelRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{3}
}

func (x *GHDelRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GHDelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GHDelResponse) Reset() {
	*x = GHDelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHDelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHDelResponse) ProtoMessage() {}

func (x *GHDelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHDelResponse.ProtoReflect.Descriptor instead.
func (*GHDelResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{4}
}

func (x *GHDelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GHAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Points []*Point `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *GHAddRequest) Reset() {
	*x = GHAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHAddRequest) ProtoMessage() {}

func (x *GHAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHAddRequest.ProtoReflect.Descriptor instead.
func (*GHAddRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{5}
}

func (x *GHAddRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GHAddRequest) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type GHAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GHAddResponse) Reset() {
	*x = GHAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHAddResponse) ProtoMessage() {}

func (x *GHAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHAddResponse.ProtoReflect.Descriptor instead.
func (*GHAddResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{6}
}

func (x *GHAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GHRemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ids [][]byte `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GHRemRequest) Reset() {
	*x = GHRemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHRemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHRemRequest) ProtoMessage() {}

func (x *GHRemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHRemRequest.ProtoReflect.Descriptor instead.
func (*GHRemRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{7}
}

func (x *GHRemRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GHRemRequest) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GHRemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GHRemResponse) Reset() {
	*x = GHRemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHRemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHRemResponse) ProtoMessage() {}

func (x *GHRemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHRemResponse.ProtoReflect.Descriptor instead.
func (*GHRemResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{8}
}

func (x *GHRemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GHGetBoxesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       []byte  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GHGetBoxesRequest) Reset() {
	*x = GHGetBoxesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHGetBoxesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHGetBoxesRequest) ProtoMessage() {}

func (x *GHGetBoxesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHGetBoxesRequest.ProtoReflect.Descriptor instead.
func (*GHGetBoxesRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{9}
}

func (x *GHGetBoxesRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GHGetBoxesRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GHGetBoxesRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GHGetBoxesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *GHGetBoxesResponse) Reset() {
	*x = GHGetBoxesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHGetBoxesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHGetBoxesResponse) ProtoMessage() {}

func (x *GHGetBoxesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHGetBoxesResponse.ProtoReflect.Descriptor instead.
func (*GHGetBoxesResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{10}
}

func (x *GHGetBoxesResponse) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type GHGetNeighborsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       []byte  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GHGetNeighborsRequest) Reset() {
	*x = GHGetNeighborsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHGetNeighborsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHGetNeighborsRequest) ProtoMessage() {}

func (x *GHGetNeighborsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHGetNeighborsRequest.ProtoReflect.Descriptor instead.
func (*GHGetNeighborsRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{11}
}

func (x *GHGetNeighborsRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GHGetNeighborsRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GHGetNeighborsRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GHGetNeighborsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *GHGetNeighborsResponse) Reset() {
	*x = GHGetNeighborsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHGetNeighborsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHGetNeighborsResponse) ProtoMessage() {}

func (x *GHGetNeighborsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHGetNeighborsResponse.ProtoReflect.Descriptor instead.
func (*GHGetNeighborsResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{12}
}

func (x *GHGetNeighborsResponse) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type GHCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GHCountRequest) Reset() {
	*x = GHCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHCountRequest) ProtoMessage() {}

func (x *GHCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHCountRequest.ProtoReflect.Descriptor instead.
func (*GHCountRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{13}
}

func (x *GHCountRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GHCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GHCountResponse) Reset() {
	*x = GHCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHCountResponse) ProtoMessage() {}

func (x *GHCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHCountResponse.ProtoReflect.Descriptor instead.
func (*GHCountResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{14}
}

func (x *GHCountResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GHMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GHMembersRequest) Reset() {
	*x = GHMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHMembersRequest) ProtoMessage() {}

func (x *GHMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHMembersRequest.ProtoReflect.Descriptor instead.
func (*GHMembersRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{15}
}

func (x *GHMembersRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GHMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *GHMembersResponse) Reset() {
	*x = GHMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GHMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GHMembersResponse) ProtoMessage() {}

func (x *GHMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GHMembersResponse.ProtoReflect.Descriptor instead.
func (*GHMembersResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP(), []int{16}
}

func (x *GHMembersResponse) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_internal_pb_protobuf_spec_geo_hash_proto protoreflect.FileDescriptor

var file_internal_pb_protobuf_spec_geo_hash_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x67, 0x65, 0x6f, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x41, 0x0a, 0x0f, 0x47, 0x48, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x48, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x20, 0x0a, 0x0c, 0x47, 0x48, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x48, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46,
	0x0a, 0x0c, 0x47, 0x48, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x24, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x48, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x32, 0x0a, 0x0c, 0x47, 0x48, 0x52, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x48, 0x52, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x5f, 0x0a, 0x11, 0x47, 0x48, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x48, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x63, 0x0a,
	0x15, 0x47, 0x48, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x48, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x48, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x48, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x24, 0x0a, 0x10, 0x47, 0x48, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x39, 0x0a, 0x11, 0x47, 0x48, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_protobuf_spec_geo_hash_proto_rawDescOnce sync.Once
	file_internal_pb_protobuf_spec_geo_hash_proto_rawDescData = file_internal_pb_protobuf_spec_geo_hash_proto_rawDesc
)

func file_internal_pb_protobuf_spec_geo_hash_proto_rawDescGZIP() []byte {
	file_internal_pb_protobuf_spec_geo_hash_proto_rawDescOnce.Do(func() {
		file_internal_pb_protobuf_spec_geo_hash_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_protobuf_spec_geo_hash_proto_rawDescData)
	})
	return file_internal_pb_protobuf_spec_geo_hash_proto_rawDescData
}

var file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_internal_pb_protobuf_spec_geo_hash_proto_goTypes = []interface{}{
	(*Point)(nil),                  // 0: proto.Point
	(*GHCreateRequest)(nil),        // 1: proto.GHCreateRequest
	(*GHCreateResponse)(nil),       // 2: proto.GHCreateResponse
	(*GHDelRequest)(nil),           // 3: proto.GHDelRequest
	(*GHDelResponse)(nil),          // 4: proto.GHDelResponse
	(*GHAddRequest)(nil),           // 5: proto.GHAddRequest
	(*GHAddResponse)(nil),          // 6: proto.GHAddResponse
	(*GHRemRequest)(nil),           // 7: proto.GHRemRequest
	(*GHRemResponse)(nil),          // 8: proto.GHRemResponse
	(*GHGetBoxesRequest)(nil),      // 9: proto.GHGetBoxesRequest
	(*GHGetBoxesResponse)(nil),     // 10: proto.GHGetBoxesResponse
	(*GHGetNeighborsRequest)(nil),  // 11: proto.GHGetNeighborsRequest
	(*GHGetNeighborsResponse)(nil), // 12: proto.GHGetNeighborsResponse
	(*GHCountRequest)(nil),         // 13: proto.GHCountRequest
	(*GHCountResponse)(nil),        // 14: proto.GHCountResponse
	(*GHMembersRequest)(nil),       // 15: proto.GHMembersRequest
	(*GHMembersResponse)(nil),      // 16: proto.GHMembersResponse
}
var file_internal_pb_protobuf_spec_geo_hash_proto_depIdxs = []int32{
	0, // 0: proto.GHAddRequest.points:type_name -> proto.Point
	0, // 1: proto.GHGetBoxesResponse.points:type_name -> proto.Point
	0, // 2: proto.GHGetNeighborsResponse.points:type_name -> proto.Point
	0, // 3: proto.GHMembersResponse.points:type_name -> proto.Point
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_pb_protobuf_spec_geo_hash_proto_init() }
func file_internal_pb_protobuf_spec_geo_hash_proto_init() {
	if File_internal_pb_protobuf_spec_geo_hash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHCreateRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHCreateResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHDelRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHDelResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHAddRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHAddResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHRemRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHRemResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHGetBoxesRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHGetBoxesResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHGetNeighborsRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHGetNeighborsResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHCountRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHCountResponse); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHMembersRequest); i {
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
		file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GHMembersResponse); i {
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
			RawDescriptor: file_internal_pb_protobuf_spec_geo_hash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_protobuf_spec_geo_hash_proto_goTypes,
		DependencyIndexes: file_internal_pb_protobuf_spec_geo_hash_proto_depIdxs,
		MessageInfos:      file_internal_pb_protobuf_spec_geo_hash_proto_msgTypes,
	}.Build()
	File_internal_pb_protobuf_spec_geo_hash_proto = out.File
	file_internal_pb_protobuf_spec_geo_hash_proto_rawDesc = nil
	file_internal_pb_protobuf_spec_geo_hash_proto_goTypes = nil
	file_internal_pb_protobuf_spec_geo_hash_proto_depIdxs = nil
}
