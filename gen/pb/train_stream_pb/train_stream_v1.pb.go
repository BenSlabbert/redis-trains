// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pkg/stream/train_stream_v1.proto

package train_stream_pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Payload:
	//	*Event_Error
	//	*Event_Starting
	//	*Event_ArrivingAtStation
	//	*Event_LeavingStation
	//	*Event_TravelingToNextStation
	//	*Event_Stopping
	//	*Event_TurningAround
	//	*Event_Stopped
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[0]
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
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetError() *ErrorMessage {
	if x, ok := x.GetPayload().(*Event_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Event) GetStarting() *StartingMessage {
	if x, ok := x.GetPayload().(*Event_Starting); ok {
		return x.Starting
	}
	return nil
}

func (x *Event) GetArrivingAtStation() *ArrivingAtStationMessage {
	if x, ok := x.GetPayload().(*Event_ArrivingAtStation); ok {
		return x.ArrivingAtStation
	}
	return nil
}

func (x *Event) GetLeavingStation() *LeavingStationMessage {
	if x, ok := x.GetPayload().(*Event_LeavingStation); ok {
		return x.LeavingStation
	}
	return nil
}

func (x *Event) GetTravelingToNextStation() *TravelingToNextStationMessage {
	if x, ok := x.GetPayload().(*Event_TravelingToNextStation); ok {
		return x.TravelingToNextStation
	}
	return nil
}

func (x *Event) GetStopping() *StoppingMessage {
	if x, ok := x.GetPayload().(*Event_Stopping); ok {
		return x.Stopping
	}
	return nil
}

func (x *Event) GetTurningAround() *TurningAroundMessage {
	if x, ok := x.GetPayload().(*Event_TurningAround); ok {
		return x.TurningAround
	}
	return nil
}

func (x *Event) GetStopped() *StoppedMessage {
	if x, ok := x.GetPayload().(*Event_Stopped); ok {
		return x.Stopped
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_Error struct {
	Error *ErrorMessage `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type Event_Starting struct {
	Starting *StartingMessage `protobuf:"bytes,3,opt,name=starting,proto3,oneof"`
}

type Event_ArrivingAtStation struct {
	ArrivingAtStation *ArrivingAtStationMessage `protobuf:"bytes,4,opt,name=arriving_at_station,json=arrivingAtStation,proto3,oneof"`
}

type Event_LeavingStation struct {
	LeavingStation *LeavingStationMessage `protobuf:"bytes,5,opt,name=leaving_station,json=leavingStation,proto3,oneof"`
}

type Event_TravelingToNextStation struct {
	TravelingToNextStation *TravelingToNextStationMessage `protobuf:"bytes,6,opt,name=traveling_to_next_station,json=travelingToNextStation,proto3,oneof"`
}

type Event_Stopping struct {
	Stopping *StoppingMessage `protobuf:"bytes,7,opt,name=stopping,proto3,oneof"`
}

type Event_TurningAround struct {
	TurningAround *TurningAroundMessage `protobuf:"bytes,8,opt,name=turning_around,json=turningAround,proto3,oneof"`
}

type Event_Stopped struct {
	Stopped *StoppedMessage `protobuf:"bytes,9,opt,name=stopped,proto3,oneof"`
}

func (*Event_Error) isEvent_Payload() {}

func (*Event_Starting) isEvent_Payload() {}

func (*Event_ArrivingAtStation) isEvent_Payload() {}

func (*Event_LeavingStation) isEvent_Payload() {}

func (*Event_TravelingToNextStation) isEvent_Payload() {}

func (*Event_Stopping) isEvent_Payload() {}

func (*Event_TurningAround) isEvent_Payload() {}

func (*Event_Stopped) isEvent_Payload() {}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{1}
}

type StartingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartingMessage) Reset() {
	*x = StartingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartingMessage) ProtoMessage() {}

func (x *StartingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartingMessage.ProtoReflect.Descriptor instead.
func (*StartingMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{2}
}

type ArrivingAtStationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArrivingAtStationMessage) Reset() {
	*x = ArrivingAtStationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrivingAtStationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrivingAtStationMessage) ProtoMessage() {}

func (x *ArrivingAtStationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrivingAtStationMessage.ProtoReflect.Descriptor instead.
func (*ArrivingAtStationMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{3}
}

type LeavingStationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeavingStationMessage) Reset() {
	*x = LeavingStationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeavingStationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeavingStationMessage) ProtoMessage() {}

func (x *LeavingStationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeavingStationMessage.ProtoReflect.Descriptor instead.
func (*LeavingStationMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{4}
}

type TravelingToNextStationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TravelingToNextStationMessage) Reset() {
	*x = TravelingToNextStationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelingToNextStationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelingToNextStationMessage) ProtoMessage() {}

func (x *TravelingToNextStationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelingToNextStationMessage.ProtoReflect.Descriptor instead.
func (*TravelingToNextStationMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{5}
}

type StoppingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoppingMessage) Reset() {
	*x = StoppingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoppingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoppingMessage) ProtoMessage() {}

func (x *StoppingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoppingMessage.ProtoReflect.Descriptor instead.
func (*StoppingMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{6}
}

type TurningAroundMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TurningAroundMessage) Reset() {
	*x = TurningAroundMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurningAroundMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurningAroundMessage) ProtoMessage() {}

func (x *TurningAroundMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurningAroundMessage.ProtoReflect.Descriptor instead.
func (*TurningAroundMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{7}
}

type StoppedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoppedMessage) Reset() {
	*x = StoppedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoppedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoppedMessage) ProtoMessage() {}

func (x *StoppedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_stream_train_stream_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoppedMessage.ProtoReflect.Descriptor instead.
func (*StoppedMessage) Descriptor() ([]byte, []int) {
	return file_pkg_stream_train_stream_v1_proto_rawDescGZIP(), []int{8}
}

var File_pkg_stream_train_stream_v1_proto protoreflect.FileDescriptor

var file_pkg_stream_train_stream_v1_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x05, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x3e, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x5b, 0x0a, 0x13, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x41,
	0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x11, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x41, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0f,
	0x6c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x6c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x6b, 0x0a, 0x19, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x5f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x16, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x4e, 0x0a, 0x0e,
	0x74, 0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x74,
	0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x72, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x41, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x1d,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x4e, 0x65, 0x78, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x11, 0x0a,
	0x0f, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x16, 0x0a, 0x14, 0x54, 0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f,
	0x3b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_stream_train_stream_v1_proto_rawDescOnce sync.Once
	file_pkg_stream_train_stream_v1_proto_rawDescData = file_pkg_stream_train_stream_v1_proto_rawDesc
)

func file_pkg_stream_train_stream_v1_proto_rawDescGZIP() []byte {
	file_pkg_stream_train_stream_v1_proto_rawDescOnce.Do(func() {
		file_pkg_stream_train_stream_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_stream_train_stream_v1_proto_rawDescData)
	})
	return file_pkg_stream_train_stream_v1_proto_rawDescData
}

var file_pkg_stream_train_stream_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_stream_train_stream_v1_proto_goTypes = []interface{}{
	(*Event)(nil),                         // 0: train_stream_pb.event
	(*ErrorMessage)(nil),                  // 1: train_stream_pb.ErrorMessage
	(*StartingMessage)(nil),               // 2: train_stream_pb.StartingMessage
	(*ArrivingAtStationMessage)(nil),      // 3: train_stream_pb.ArrivingAtStationMessage
	(*LeavingStationMessage)(nil),         // 4: train_stream_pb.LeavingStationMessage
	(*TravelingToNextStationMessage)(nil), // 5: train_stream_pb.TravelingToNextStationMessage
	(*StoppingMessage)(nil),               // 6: train_stream_pb.StoppingMessage
	(*TurningAroundMessage)(nil),          // 7: train_stream_pb.TurningAroundMessage
	(*StoppedMessage)(nil),                // 8: train_stream_pb.StoppedMessage
	(*timestamp.Timestamp)(nil),           // 9: google.protobuf.Timestamp
}
var file_pkg_stream_train_stream_v1_proto_depIdxs = []int32{
	9, // 0: train_stream_pb.event.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: train_stream_pb.event.error:type_name -> train_stream_pb.ErrorMessage
	2, // 2: train_stream_pb.event.starting:type_name -> train_stream_pb.StartingMessage
	3, // 3: train_stream_pb.event.arriving_at_station:type_name -> train_stream_pb.ArrivingAtStationMessage
	4, // 4: train_stream_pb.event.leaving_station:type_name -> train_stream_pb.LeavingStationMessage
	5, // 5: train_stream_pb.event.traveling_to_next_station:type_name -> train_stream_pb.TravelingToNextStationMessage
	6, // 6: train_stream_pb.event.stopping:type_name -> train_stream_pb.StoppingMessage
	7, // 7: train_stream_pb.event.turning_around:type_name -> train_stream_pb.TurningAroundMessage
	8, // 8: train_stream_pb.event.stopped:type_name -> train_stream_pb.StoppedMessage
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_stream_train_stream_v1_proto_init() }
func file_pkg_stream_train_stream_v1_proto_init() {
	if File_pkg_stream_train_stream_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_stream_train_stream_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartingMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrivingAtStationMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeavingStationMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelingToNextStationMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoppingMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurningAroundMessage); i {
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
		file_pkg_stream_train_stream_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoppedMessage); i {
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
	file_pkg_stream_train_stream_v1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_Error)(nil),
		(*Event_Starting)(nil),
		(*Event_ArrivingAtStation)(nil),
		(*Event_LeavingStation)(nil),
		(*Event_TravelingToNextStation)(nil),
		(*Event_Stopping)(nil),
		(*Event_TurningAround)(nil),
		(*Event_Stopped)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_stream_train_stream_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_stream_train_stream_v1_proto_goTypes,
		DependencyIndexes: file_pkg_stream_train_stream_v1_proto_depIdxs,
		MessageInfos:      file_pkg_stream_train_stream_v1_proto_msgTypes,
	}.Build()
	File_pkg_stream_train_stream_v1_proto = out.File
	file_pkg_stream_train_stream_v1_proto_rawDesc = nil
	file_pkg_stream_train_stream_v1_proto_goTypes = nil
	file_pkg_stream_train_stream_v1_proto_depIdxs = nil
}
