// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: ledger.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionType int32

const (
	TransactionType_REVERSAL     TransactionType = 0
	TransactionType_CONFIRMATION TransactionType = 1
	TransactionType_CANCELLATION TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "REVERSAL",
		1: "CONFIRMATION",
		2: "CANCELLATION",
	}
	TransactionType_value = map[string]int32{
		"REVERSAL":     0,
		"CONFIRMATION": 1,
		"CANCELLATION": 2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_ledger_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_ledger_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{0}
}

type Result int32

const (
	Result_OK      Result = 0
	Result_FAILURE Result = 1
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "OK",
		1: "FAILURE",
	}
	Result_value = map[string]int32{
		"OK":      0,
		"FAILURE": 1,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_ledger_proto_enumTypes[1].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_ledger_proto_enumTypes[1]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{1}
}

type RequestConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *ConfirmationTransactionData `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Type        TransactionType              `protobuf:"varint,2,opt,name=type,proto3,enum=TransactionType" json:"type,omitempty"`
}

func (x *RequestConfirmation) Reset() {
	*x = RequestConfirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestConfirmation) ProtoMessage() {}

func (x *RequestConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestConfirmation.ProtoReflect.Descriptor instead.
func (*RequestConfirmation) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *RequestConfirmation) GetTransaction() *ConfirmationTransactionData {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *RequestConfirmation) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_REVERSAL
}

type RequestCancellation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *CancellationTransactionData `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Type        TransactionType              `protobuf:"varint,2,opt,name=type,proto3,enum=TransactionType" json:"type,omitempty"`
}

func (x *RequestCancellation) Reset() {
	*x = RequestCancellation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCancellation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCancellation) ProtoMessage() {}

func (x *RequestCancellation) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCancellation.ProtoReflect.Descriptor instead.
func (*RequestCancellation) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *RequestCancellation) GetTransaction() *CancellationTransactionData {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *RequestCancellation) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_REVERSAL
}

type RequestReversal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *ReversalTransactionData `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Type        TransactionType          `protobuf:"varint,2,opt,name=type,proto3,enum=TransactionType" json:"type,omitempty"`
}

func (x *RequestReversal) Reset() {
	*x = RequestReversal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReversal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReversal) ProtoMessage() {}

func (x *RequestReversal) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReversal.ProtoReflect.Descriptor instead.
func (*RequestReversal) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{2}
}

func (x *RequestReversal) GetTransaction() *ReversalTransactionData {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *RequestReversal) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_REVERSAL
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegisterId string `protobuf:"bytes,1,opt,name=register_id,json=registerId,proto3" json:"register_id,omitempty"`
	Result     Result `protobuf:"varint,2,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionResponse) GetRegisterId() string {
	if x != nil {
		return x.RegisterId
	}
	return ""
}

func (x *TransactionResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_OK
}

type ConfirmationTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalId        string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Value             float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	LocalTime         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
	MerchantId        string                 `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	CardId            string                 `protobuf:"bytes,6,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	CurrencyCode      string                 `protobuf:"bytes,7,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	AuthorizationCode string                 `protobuf:"bytes,8,opt,name=authorization_code,json=authorizationCode,proto3" json:"authorization_code,omitempty"`
	CountryCode       string                 `protobuf:"bytes,9,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *ConfirmationTransactionData) Reset() {
	*x = ConfirmationTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationTransactionData) ProtoMessage() {}

func (x *ConfirmationTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationTransactionData.ProtoReflect.Descriptor instead.
func (*ConfirmationTransactionData) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmationTransactionData) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *ConfirmationTransactionData) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ConfirmationTransactionData) GetLocalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LocalTime
	}
	return nil
}

func (x *ConfirmationTransactionData) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ConfirmationTransactionData) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *ConfirmationTransactionData) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *ConfirmationTransactionData) GetAuthorizationCode() string {
	if x != nil {
		return x.AuthorizationCode
	}
	return ""
}

func (x *ConfirmationTransactionData) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type CancellationTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalId        string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Value             float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	LocalTime         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
	MerchantId        string                 `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	CardId            string                 `protobuf:"bytes,5,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	CurrencyCode      string                 `protobuf:"bytes,7,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	AuthorizationCode string                 `protobuf:"bytes,8,opt,name=authorization_code,json=authorizationCode,proto3" json:"authorization_code,omitempty"`
	CountryCode       string                 `protobuf:"bytes,9,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *CancellationTransactionData) Reset() {
	*x = CancellationTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancellationTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancellationTransactionData) ProtoMessage() {}

func (x *CancellationTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancellationTransactionData.ProtoReflect.Descriptor instead.
func (*CancellationTransactionData) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{5}
}

func (x *CancellationTransactionData) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *CancellationTransactionData) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CancellationTransactionData) GetLocalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LocalTime
	}
	return nil
}

func (x *CancellationTransactionData) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *CancellationTransactionData) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *CancellationTransactionData) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *CancellationTransactionData) GetAuthorizationCode() string {
	if x != nil {
		return x.AuthorizationCode
	}
	return ""
}

func (x *CancellationTransactionData) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type ReversalTransactionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalId        string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Value             float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	LocalTime         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
	MerchantId        string                 `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	CardId            string                 `protobuf:"bytes,6,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	CurrencyCode      string                 `protobuf:"bytes,7,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	AuthorizationCode string                 `protobuf:"bytes,8,opt,name=authorization_code,json=authorizationCode,proto3" json:"authorization_code,omitempty"`
	CountryCode       string                 `protobuf:"bytes,9,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *ReversalTransactionData) Reset() {
	*x = ReversalTransactionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReversalTransactionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReversalTransactionData) ProtoMessage() {}

func (x *ReversalTransactionData) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReversalTransactionData.ProtoReflect.Descriptor instead.
func (*ReversalTransactionData) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{6}
}

func (x *ReversalTransactionData) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *ReversalTransactionData) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ReversalTransactionData) GetLocalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LocalTime
	}
	return nil
}

func (x *ReversalTransactionData) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ReversalTransactionData) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *ReversalTransactionData) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *ReversalTransactionData) GetAuthorizationCode() string {
	if x != nil {
		return x.AuthorizationCode
	}
	return ""
}

func (x *ReversalTransactionData) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

var File_ledger_proto protoreflect.FileDescriptor

var file_ledger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7b, 0x0a, 0x13,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x57,
	0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xc0, 0x02, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x1b, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xbc, 0x02,
	0x0a, 0x17, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x43, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x02, 0x2a, 0x1d, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01,
	0x32, 0xc1, 0x01, 0x0a, 0x0d, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x08, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x1a, 0x14, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2d, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledger_proto_rawDescOnce sync.Once
	file_ledger_proto_rawDescData = file_ledger_proto_rawDesc
)

func file_ledger_proto_rawDescGZIP() []byte {
	file_ledger_proto_rawDescOnce.Do(func() {
		file_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_proto_rawDescData)
	})
	return file_ledger_proto_rawDescData
}

var file_ledger_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ledger_proto_goTypes = []interface{}{
	(TransactionType)(0),                // 0: TransactionType
	(Result)(0),                         // 1: Result
	(*RequestConfirmation)(nil),         // 2: RequestConfirmation
	(*RequestCancellation)(nil),         // 3: RequestCancellation
	(*RequestReversal)(nil),             // 4: RequestReversal
	(*TransactionResponse)(nil),         // 5: TransactionResponse
	(*ConfirmationTransactionData)(nil), // 6: ConfirmationTransactionData
	(*CancellationTransactionData)(nil), // 7: CancellationTransactionData
	(*ReversalTransactionData)(nil),     // 8: ReversalTransactionData
	(*timestamppb.Timestamp)(nil),       // 9: google.protobuf.Timestamp
}
var file_ledger_proto_depIdxs = []int32{
	6,  // 0: RequestConfirmation.transaction:type_name -> ConfirmationTransactionData
	0,  // 1: RequestConfirmation.type:type_name -> TransactionType
	7,  // 2: RequestCancellation.transaction:type_name -> CancellationTransactionData
	0,  // 3: RequestCancellation.type:type_name -> TransactionType
	8,  // 4: RequestReversal.transaction:type_name -> ReversalTransactionData
	0,  // 5: RequestReversal.type:type_name -> TransactionType
	1,  // 6: TransactionResponse.result:type_name -> Result
	9,  // 7: ConfirmationTransactionData.local_time:type_name -> google.protobuf.Timestamp
	9,  // 8: CancellationTransactionData.local_time:type_name -> google.protobuf.Timestamp
	9,  // 9: ReversalTransactionData.local_time:type_name -> google.protobuf.Timestamp
	2,  // 10: LedgerService.Confirmation:input_type -> RequestConfirmation
	3,  // 11: LedgerService.Cancellation:input_type -> RequestCancellation
	4,  // 12: LedgerService.Reversal:input_type -> RequestReversal
	5,  // 13: LedgerService.Confirmation:output_type -> TransactionResponse
	5,  // 14: LedgerService.Cancellation:output_type -> TransactionResponse
	5,  // 15: LedgerService.Reversal:output_type -> TransactionResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_ledger_proto_init() }
func file_ledger_proto_init() {
	if File_ledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestConfirmation); i {
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
		file_ledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCancellation); i {
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
		file_ledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReversal); i {
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
		file_ledger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
		file_ledger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationTransactionData); i {
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
		file_ledger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancellationTransactionData); i {
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
		file_ledger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReversalTransactionData); i {
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
			RawDescriptor: file_ledger_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ledger_proto_goTypes,
		DependencyIndexes: file_ledger_proto_depIdxs,
		EnumInfos:         file_ledger_proto_enumTypes,
		MessageInfos:      file_ledger_proto_msgTypes,
	}.Build()
	File_ledger_proto = out.File
	file_ledger_proto_rawDesc = nil
	file_ledger_proto_goTypes = nil
	file_ledger_proto_depIdxs = nil
}
