// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api.proto

package pb

import (
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type EventKind int32

const (
	EventKind_Empty EventKind = 0
)

// Enum value maps for EventKind.
var (
	EventKind_name = map[int32]string{
		0: "Empty",
	}
	EventKind_value = map[string]int32{
		"Empty": 0,
	}
)

func (x EventKind) Enum() *EventKind {
	p := new(EventKind)
	*p = x
	return p
}

func (x EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (EventKind) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventKind.Descriptor instead.
func (EventKind) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type PushEventOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer []byte `protobuf:"bytes,1,opt,name=Signer,proto3" json:"Signer,omitempty"`
	Sig    []byte `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
}

func (x *PushEventOutput) Reset() {
	*x = PushEventOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventOutput) ProtoMessage() {}

func (x *PushEventOutput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventOutput.ProtoReflect.Descriptor instead.
func (*PushEventOutput) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *PushEventOutput) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *PushEventOutput) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type PushEventInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (x *PushEventInput) Reset() {
	*x = PushEventInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventInput) ProtoMessage() {}

func (x *PushEventInput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventInput.ProtoReflect.Descriptor instead.
func (*PushEventInput) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *PushEventInput) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type GetEventOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   *Event `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
	Confirm bool   `protobuf:"varint,2,opt,name=Confirm,proto3" json:"Confirm,omitempty"`
}

func (x *GetEventOutput) Reset() {
	*x = GetEventOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventOutput) ProtoMessage() {}

func (x *GetEventOutput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventOutput.ProtoReflect.Descriptor instead.
func (*GetEventOutput) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventOutput) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *GetEventOutput) GetConfirm() bool {
	if x != nil {
		return x.Confirm
	}
	return false
}

type GetEventInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer []byte `protobuf:"bytes,1,opt,name=Signer,proto3" json:"Signer,omitempty"`
	Index  uint64 `protobuf:"fixed64,2,opt,name=Index,proto3" json:"Index,omitempty"`
}

func (x *GetEventInput) Reset() {
	*x = GetEventInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventInput) ProtoMessage() {}

func (x *GetEventInput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventInput.ProtoReflect.Descriptor instead.
func (*GetEventInput) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventInput) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *GetEventInput) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type EventPtr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer []byte `protobuf:"bytes,1,opt,name=Signer,proto3" json:"Signer,omitempty"`
	Index  uint64 `protobuf:"fixed64,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Digest []byte `protobuf:"bytes,3,opt,name=Digest,proto3" json:"Digest,omitempty"`
}

func (x *EventPtr) Reset() {
	*x = EventPtr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPtr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPtr) ProtoMessage() {}

func (x *EventPtr) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPtr.ProtoReflect.Descriptor instead.
func (*EventPtr) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *EventPtr) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *EventPtr) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *EventPtr) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig     []byte      `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Signer  []byte      `protobuf:"bytes,2,opt,name=Signer,proto3" json:"Signer,omitempty"`
	Index   uint64      `protobuf:"fixed64,3,opt,name=Index,proto3" json:"Index,omitempty"`
	Kind    EventKind   `protobuf:"varint,5,opt,name=Kind,proto3,enum=api.EventKind" json:"Kind,omitempty"`
	Events  []*EventPtr `protobuf:"bytes,6,rep,name=Events,proto3" json:"Events,omitempty"`
	Payload []byte      `protobuf:"bytes,7,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Event) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *Event) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *Event) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Event) GetKind() EventKind {
	if x != nil {
		return x.Kind
	}
	return EventKind_Empty
}

func (x *Event) GetEvents() []*EventPtr {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *Event) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x22, 0x32,
	0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x20, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x50, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x74, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x22, 0xac, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x04, 0x4b,
	0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x25, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x74, 0x72, 0x52, 0x06,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2a, 0x16, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x10, 0x00, 0x32, 0x72, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12,
	0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(EventKind)(0),          // 0: api.EventKind
	(*PushEventOutput)(nil), // 1: api.PushEventOutput
	(*PushEventInput)(nil),  // 2: api.PushEventInput
	(*GetEventOutput)(nil),  // 3: api.GetEventOutput
	(*GetEventInput)(nil),   // 4: api.GetEventInput
	(*EventPtr)(nil),        // 5: api.EventPtr
	(*Event)(nil),           // 6: api.Event
}
var file_api_proto_depIdxs = []int32{
	6, // 0: api.PushEventInput.Event:type_name -> api.Event
	6, // 1: api.GetEventOutput.Event:type_name -> api.Event
	0, // 2: api.Event.Kind:type_name -> api.EventKind
	5, // 3: api.Event.Events:type_name -> api.EventPtr
	4, // 4: api.API.GetEvent:input_type -> api.GetEventInput
	2, // 5: api.API.PushEvent:input_type -> api.PushEventInput
	3, // 6: api.API.GetEvent:output_type -> api.GetEventOutput
	1, // 7: api.API.PushEvent:output_type -> api.PushEventOutput
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushEventOutput); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushEventInput); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventOutput); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventInput); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPtr); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
