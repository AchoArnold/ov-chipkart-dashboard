// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: raw_records.proto

package transactions

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CredentialsRawRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password   string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	CardNumber string               `protobuf:"bytes,3,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	StartDate  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate    *timestamp.Timestamp `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *CredentialsRawRecordsRequest) Reset() {
	*x = CredentialsRawRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_records_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialsRawRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialsRawRecordsRequest) ProtoMessage() {}

func (x *CredentialsRawRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raw_records_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialsRawRecordsRequest.ProtoReflect.Descriptor instead.
func (*CredentialsRawRecordsRequest) Descriptor() ([]byte, []int) {
	return file_raw_records_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialsRawRecordsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CredentialsRawRecordsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CredentialsRawRecordsRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CredentialsRawRecordsRequest) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CredentialsRawRecordsRequest) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type BytesRawRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNumber string               `protobuf:"bytes,1,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	StartDate  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *BytesRawRecordsRequest) Reset() {
	*x = BytesRawRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_records_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesRawRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesRawRecordsRequest) ProtoMessage() {}

func (x *BytesRawRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raw_records_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesRawRecordsRequest.ProtoReflect.Descriptor instead.
func (*BytesRawRecordsRequest) Descriptor() ([]byte, []int) {
	return file_raw_records_proto_rawDescGZIP(), []int{1}
}

func (x *BytesRawRecordsRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *BytesRawRecordsRequest) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *BytesRawRecordsRequest) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type RawRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckInInfo            string                `protobuf:"bytes,1,opt,name=checkInInfo,proto3" json:"checkInInfo,omitempty"`
	CheckInText            string                `protobuf:"bytes,2,opt,name=checkInText,proto3" json:"checkInText,omitempty"`
	Fare                   *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=fare,proto3" json:"fare,omitempty"`
	FareCalculation        string                `protobuf:"bytes,4,opt,name=fareCalculation,proto3" json:"fareCalculation,omitempty"`
	FareText               string                `protobuf:"bytes,5,opt,name=fareText,proto3" json:"fareText,omitempty"`
	ModalType              string                `protobuf:"bytes,6,opt,name=modalType,proto3" json:"modalType,omitempty"`
	ProductInfo            string                `protobuf:"bytes,7,opt,name=productInfo,proto3" json:"productInfo,omitempty"`
	ProductText            string                `protobuf:"bytes,8,opt,name=productText,proto3" json:"productText,omitempty"`
	Pto                    string                `protobuf:"bytes,9,opt,name=pto,proto3" json:"pto,omitempty"`
	TransactionDateTime    *timestamp.Timestamp  `protobuf:"bytes,10,opt,name=transactionDateTime,proto3" json:"transactionDateTime,omitempty"`
	TransactionInfo        string                `protobuf:"bytes,11,opt,name=transactionInfo,proto3" json:"transactionInfo,omitempty"`
	TransactionName        string                `protobuf:"bytes,12,opt,name=transactionName,proto3" json:"transactionName,omitempty"`
	EPurseMut              *wrappers.DoubleValue `protobuf:"bytes,13,opt,name=ePurseMut,proto3" json:"ePurseMut,omitempty"`
	EPurseMutInfo          string                `protobuf:"bytes,14,opt,name=ePurseMutInfo,proto3" json:"ePurseMutInfo,omitempty"`
	TransactionExplanation string                `protobuf:"bytes,15,opt,name=transactionExplanation,proto3" json:"transactionExplanation,omitempty"`
	TransactionPriority    string                `protobuf:"bytes,16,opt,name=transactionPriority,proto3" json:"transactionPriority,omitempty"`
}

func (x *RawRecord) Reset() {
	*x = RawRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_records_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawRecord) ProtoMessage() {}

func (x *RawRecord) ProtoReflect() protoreflect.Message {
	mi := &file_raw_records_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawRecord.ProtoReflect.Descriptor instead.
func (*RawRecord) Descriptor() ([]byte, []int) {
	return file_raw_records_proto_rawDescGZIP(), []int{2}
}

func (x *RawRecord) GetCheckInInfo() string {
	if x != nil {
		return x.CheckInInfo
	}
	return ""
}

func (x *RawRecord) GetCheckInText() string {
	if x != nil {
		return x.CheckInText
	}
	return ""
}

func (x *RawRecord) GetFare() *wrappers.DoubleValue {
	if x != nil {
		return x.Fare
	}
	return nil
}

func (x *RawRecord) GetFareCalculation() string {
	if x != nil {
		return x.FareCalculation
	}
	return ""
}

func (x *RawRecord) GetFareText() string {
	if x != nil {
		return x.FareText
	}
	return ""
}

func (x *RawRecord) GetModalType() string {
	if x != nil {
		return x.ModalType
	}
	return ""
}

func (x *RawRecord) GetProductInfo() string {
	if x != nil {
		return x.ProductInfo
	}
	return ""
}

func (x *RawRecord) GetProductText() string {
	if x != nil {
		return x.ProductText
	}
	return ""
}

func (x *RawRecord) GetPto() string {
	if x != nil {
		return x.Pto
	}
	return ""
}

func (x *RawRecord) GetTransactionDateTime() *timestamp.Timestamp {
	if x != nil {
		return x.TransactionDateTime
	}
	return nil
}

func (x *RawRecord) GetTransactionInfo() string {
	if x != nil {
		return x.TransactionInfo
	}
	return ""
}

func (x *RawRecord) GetTransactionName() string {
	if x != nil {
		return x.TransactionName
	}
	return ""
}

func (x *RawRecord) GetEPurseMut() *wrappers.DoubleValue {
	if x != nil {
		return x.EPurseMut
	}
	return nil
}

func (x *RawRecord) GetEPurseMutInfo() string {
	if x != nil {
		return x.EPurseMutInfo
	}
	return ""
}

func (x *RawRecord) GetTransactionExplanation() string {
	if x != nil {
		return x.TransactionExplanation
	}
	return ""
}

func (x *RawRecord) GetTransactionPriority() string {
	if x != nil {
		return x.TransactionPriority
	}
	return ""
}

type RawRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawRecords []*RawRecord `protobuf:"bytes,1,rep,name=rawRecords,proto3" json:"rawRecords,omitempty"`
}

func (x *RawRecordsResponse) Reset() {
	*x = RawRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_records_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawRecordsResponse) ProtoMessage() {}

func (x *RawRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raw_records_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawRecordsResponse.ProtoReflect.Descriptor instead.
func (*RawRecordsResponse) Descriptor() ([]byte, []int) {
	return file_raw_records_proto_rawDescGZIP(), []int{3}
}

func (x *RawRecordsResponse) GetRawRecords() []*RawRecord {
	if x != nil {
		return x.RawRecords
	}
	return nil
}

var File_raw_records_proto protoreflect.FileDescriptor

var file_raw_records_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x61, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x16,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa9, 0x05, 0x0a, 0x09, 0x52, 0x61, 0x77, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x66, 0x61, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x66, 0x61, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x61,
	0x72, 0x65, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x61, 0x72, 0x65, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x72, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x74, 0x6f, 0x12, 0x4c, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x65, 0x50, 0x75, 0x72, 0x73, 0x65,
	0x4d, 0x75, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x65, 0x50, 0x75, 0x72, 0x73, 0x65, 0x4d,
	0x75, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x50, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x75, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x50, 0x75, 0x72, 0x73,
	0x65, 0x4d, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x4d, 0x0a, 0x12, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x32, 0xe3, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x19, 0x52, 0x61, 0x77,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x13, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raw_records_proto_rawDescOnce sync.Once
	file_raw_records_proto_rawDescData = file_raw_records_proto_rawDesc
)

func file_raw_records_proto_rawDescGZIP() []byte {
	file_raw_records_proto_rawDescOnce.Do(func() {
		file_raw_records_proto_rawDescData = protoimpl.X.CompressGZIP(file_raw_records_proto_rawDescData)
	})
	return file_raw_records_proto_rawDescData
}

var file_raw_records_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_raw_records_proto_goTypes = []interface{}{
	(*CredentialsRawRecordsRequest)(nil), // 0: transactions.CredentialsRawRecordsRequest
	(*BytesRawRecordsRequest)(nil),       // 1: transactions.BytesRawRecordsRequest
	(*RawRecord)(nil),                    // 2: transactions.RawRecord
	(*RawRecordsResponse)(nil),           // 3: transactions.RawRecordsResponse
	(*timestamp.Timestamp)(nil),          // 4: google.protobuf.Timestamp
	(*wrappers.DoubleValue)(nil),         // 5: google.protobuf.DoubleValue
}
var file_raw_records_proto_depIdxs = []int32{
	4,  // 0: transactions.CredentialsRawRecordsRequest.startDate:type_name -> google.protobuf.Timestamp
	4,  // 1: transactions.CredentialsRawRecordsRequest.endDate:type_name -> google.protobuf.Timestamp
	4,  // 2: transactions.BytesRawRecordsRequest.startDate:type_name -> google.protobuf.Timestamp
	4,  // 3: transactions.BytesRawRecordsRequest.endDate:type_name -> google.protobuf.Timestamp
	5,  // 4: transactions.RawRecord.fare:type_name -> google.protobuf.DoubleValue
	4,  // 5: transactions.RawRecord.transactionDateTime:type_name -> google.protobuf.Timestamp
	5,  // 6: transactions.RawRecord.ePurseMut:type_name -> google.protobuf.DoubleValue
	2,  // 7: transactions.RawRecordsResponse.rawRecords:type_name -> transactions.RawRecord
	0,  // 8: transactions.TransactionsService.RawRecordsWithCredentials:input_type -> transactions.CredentialsRawRecordsRequest
	1,  // 9: transactions.TransactionsService.RawRecordsFromBytes:input_type -> transactions.BytesRawRecordsRequest
	3,  // 10: transactions.TransactionsService.RawRecordsWithCredentials:output_type -> transactions.RawRecordsResponse
	3,  // 11: transactions.TransactionsService.RawRecordsFromBytes:output_type -> transactions.RawRecordsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_raw_records_proto_init() }
func file_raw_records_proto_init() {
	if File_raw_records_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raw_records_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialsRawRecordsRequest); i {
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
		file_raw_records_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesRawRecordsRequest); i {
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
		file_raw_records_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawRecord); i {
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
		file_raw_records_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawRecordsResponse); i {
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
			RawDescriptor: file_raw_records_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raw_records_proto_goTypes,
		DependencyIndexes: file_raw_records_proto_depIdxs,
		MessageInfos:      file_raw_records_proto_msgTypes,
	}.Build()
	File_raw_records_proto = out.File
	file_raw_records_proto_rawDesc = nil
	file_raw_records_proto_goTypes = nil
	file_raw_records_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionsServiceClient interface {
	RawRecordsWithCredentials(ctx context.Context, in *CredentialsRawRecordsRequest, opts ...grpc.CallOption) (*RawRecordsResponse, error)
	RawRecordsFromBytes(ctx context.Context, in *BytesRawRecordsRequest, opts ...grpc.CallOption) (*RawRecordsResponse, error)
}

type transactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsServiceClient(cc grpc.ClientConnInterface) TransactionsServiceClient {
	return &transactionsServiceClient{cc}
}

func (c *transactionsServiceClient) RawRecordsWithCredentials(ctx context.Context, in *CredentialsRawRecordsRequest, opts ...grpc.CallOption) (*RawRecordsResponse, error) {
	out := new(RawRecordsResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/RawRecordsWithCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) RawRecordsFromBytes(ctx context.Context, in *BytesRawRecordsRequest, opts ...grpc.CallOption) (*RawRecordsResponse, error) {
	out := new(RawRecordsResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/RawRecordsFromBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
type TransactionsServiceServer interface {
	RawRecordsWithCredentials(context.Context, *CredentialsRawRecordsRequest) (*RawRecordsResponse, error)
	RawRecordsFromBytes(context.Context, *BytesRawRecordsRequest) (*RawRecordsResponse, error)
}

// UnimplementedTransactionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionsServiceServer struct {
}

func (*UnimplementedTransactionsServiceServer) RawRecordsWithCredentials(context.Context, *CredentialsRawRecordsRequest) (*RawRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawRecordsWithCredentials not implemented")
}
func (*UnimplementedTransactionsServiceServer) RawRecordsFromBytes(context.Context, *BytesRawRecordsRequest) (*RawRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawRecordsFromBytes not implemented")
}

func RegisterTransactionsServiceServer(s *grpc.Server, srv TransactionsServiceServer) {
	s.RegisterService(&_TransactionsService_serviceDesc, srv)
}

func _TransactionsService_RawRecordsWithCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsRawRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).RawRecordsWithCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/RawRecordsWithCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).RawRecordsWithCredentials(ctx, req.(*CredentialsRawRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_RawRecordsFromBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytesRawRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).RawRecordsFromBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/RawRecordsFromBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).RawRecordsFromBytes(ctx, req.(*BytesRawRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionsService",
	HandlerType: (*TransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RawRecordsWithCredentials",
			Handler:    _TransactionsService_RawRecordsWithCredentials_Handler,
		},
		{
			MethodName: "RawRecordsFromBytes",
			Handler:    _TransactionsService_RawRecordsFromBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raw_records.proto",
}
