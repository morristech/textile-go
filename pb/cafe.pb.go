// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cafe.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CafeChallenge struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeChallenge) Reset()         { *m = CafeChallenge{} }
func (m *CafeChallenge) String() string { return proto.CompactTextString(m) }
func (*CafeChallenge) ProtoMessage()    {}
func (*CafeChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{0}
}
func (m *CafeChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeChallenge.Unmarshal(m, b)
}
func (m *CafeChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeChallenge.Marshal(b, m, deterministic)
}
func (dst *CafeChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeChallenge.Merge(dst, src)
}
func (m *CafeChallenge) XXX_Size() int {
	return xxx_messageInfo_CafeChallenge.Size(m)
}
func (m *CafeChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CafeChallenge proto.InternalMessageInfo

func (m *CafeChallenge) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CafeNonce struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeNonce) Reset()         { *m = CafeNonce{} }
func (m *CafeNonce) String() string { return proto.CompactTextString(m) }
func (*CafeNonce) ProtoMessage()    {}
func (*CafeNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{1}
}
func (m *CafeNonce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeNonce.Unmarshal(m, b)
}
func (m *CafeNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeNonce.Marshal(b, m, deterministic)
}
func (dst *CafeNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeNonce.Merge(dst, src)
}
func (m *CafeNonce) XXX_Size() int {
	return xxx_messageInfo_CafeNonce.Size(m)
}
func (m *CafeNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeNonce.DiscardUnknown(m)
}

var xxx_messageInfo_CafeNonce proto.InternalMessageInfo

func (m *CafeNonce) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CafeRegistration struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sig                  []byte   `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRegistration) Reset()         { *m = CafeRegistration{} }
func (m *CafeRegistration) String() string { return proto.CompactTextString(m) }
func (*CafeRegistration) ProtoMessage()    {}
func (*CafeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{2}
}
func (m *CafeRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRegistration.Unmarshal(m, b)
}
func (m *CafeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRegistration.Marshal(b, m, deterministic)
}
func (dst *CafeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRegistration.Merge(dst, src)
}
func (m *CafeRegistration) XXX_Size() int {
	return xxx_messageInfo_CafeRegistration.Size(m)
}
func (m *CafeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRegistration proto.InternalMessageInfo

func (m *CafeRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CafeRegistration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CafeRegistration) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *CafeRegistration) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type CafeSession struct {
	Access               string               `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Exp                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=exp,proto3" json:"exp,omitempty"`
	Refresh              string               `protobuf:"bytes,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Rexp                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=rexp,proto3" json:"rexp,omitempty"`
	Subject              string               `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Type                 string               `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	HttpAddr             string               `protobuf:"bytes,7,opt,name=httpAddr,proto3" json:"httpAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeSession) Reset()         { *m = CafeSession{} }
func (m *CafeSession) String() string { return proto.CompactTextString(m) }
func (*CafeSession) ProtoMessage()    {}
func (*CafeSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{3}
}
func (m *CafeSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeSession.Unmarshal(m, b)
}
func (m *CafeSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeSession.Marshal(b, m, deterministic)
}
func (dst *CafeSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeSession.Merge(dst, src)
}
func (m *CafeSession) XXX_Size() int {
	return xxx_messageInfo_CafeSession.Size(m)
}
func (m *CafeSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeSession proto.InternalMessageInfo

func (m *CafeSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeSession) GetExp() *timestamp.Timestamp {
	if m != nil {
		return m.Exp
	}
	return nil
}

func (m *CafeSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

func (m *CafeSession) GetRexp() *timestamp.Timestamp {
	if m != nil {
		return m.Rexp
	}
	return nil
}

func (m *CafeSession) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CafeSession) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CafeSession) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

type CafeRefreshSession struct {
	Access               string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              string   `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRefreshSession) Reset()         { *m = CafeRefreshSession{} }
func (m *CafeRefreshSession) String() string { return proto.CompactTextString(m) }
func (*CafeRefreshSession) ProtoMessage()    {}
func (*CafeRefreshSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{4}
}
func (m *CafeRefreshSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRefreshSession.Unmarshal(m, b)
}
func (m *CafeRefreshSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRefreshSession.Marshal(b, m, deterministic)
}
func (dst *CafeRefreshSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRefreshSession.Merge(dst, src)
}
func (m *CafeRefreshSession) XXX_Size() int {
	return xxx_messageInfo_CafeRefreshSession.Size(m)
}
func (m *CafeRefreshSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRefreshSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRefreshSession proto.InternalMessageInfo

func (m *CafeRefreshSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeRefreshSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

type CafeStore struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStore) Reset()         { *m = CafeStore{} }
func (m *CafeStore) String() string { return proto.CompactTextString(m) }
func (*CafeStore) ProtoMessage()    {}
func (*CafeStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{5}
}
func (m *CafeStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStore.Unmarshal(m, b)
}
func (m *CafeStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStore.Marshal(b, m, deterministic)
}
func (dst *CafeStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStore.Merge(dst, src)
}
func (m *CafeStore) XXX_Size() int {
	return xxx_messageInfo_CafeStore.Size(m)
}
func (m *CafeStore) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStore.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStore proto.InternalMessageInfo

func (m *CafeStore) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStore) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeBlockList struct {
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeBlockList) Reset()         { *m = CafeBlockList{} }
func (m *CafeBlockList) String() string { return proto.CompactTextString(m) }
func (*CafeBlockList) ProtoMessage()    {}
func (*CafeBlockList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{6}
}
func (m *CafeBlockList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeBlockList.Unmarshal(m, b)
}
func (m *CafeBlockList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeBlockList.Marshal(b, m, deterministic)
}
func (dst *CafeBlockList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeBlockList.Merge(dst, src)
}
func (m *CafeBlockList) XXX_Size() int {
	return xxx_messageInfo_CafeBlockList.Size(m)
}
func (m *CafeBlockList) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeBlockList.DiscardUnknown(m)
}

var xxx_messageInfo_CafeBlockList proto.InternalMessageInfo

func (m *CafeBlockList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeBlock struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=rawData,proto3" json:"rawData,omitempty"`
	Cid                  string   `protobuf:"bytes,3,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeBlock) Reset()         { *m = CafeBlock{} }
func (m *CafeBlock) String() string { return proto.CompactTextString(m) }
func (*CafeBlock) ProtoMessage()    {}
func (*CafeBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{7}
}
func (m *CafeBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeBlock.Unmarshal(m, b)
}
func (m *CafeBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeBlock.Marshal(b, m, deterministic)
}
func (dst *CafeBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeBlock.Merge(dst, src)
}
func (m *CafeBlock) XXX_Size() int {
	return xxx_messageInfo_CafeBlock.Size(m)
}
func (m *CafeBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CafeBlock proto.InternalMessageInfo

func (m *CafeBlock) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeBlock) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *CafeBlock) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type CafeStoreThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	SkCipher             []byte   `protobuf:"bytes,3,opt,name=skCipher,proto3" json:"skCipher,omitempty"`
	HeadCipher           []byte   `protobuf:"bytes,4,opt,name=headCipher,proto3" json:"headCipher,omitempty"`
	NameCipher           []byte   `protobuf:"bytes,5,opt,name=nameCipher,proto3" json:"nameCipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreThread) Reset()         { *m = CafeStoreThread{} }
func (m *CafeStoreThread) String() string { return proto.CompactTextString(m) }
func (*CafeStoreThread) ProtoMessage()    {}
func (*CafeStoreThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{8}
}
func (m *CafeStoreThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreThread.Unmarshal(m, b)
}
func (m *CafeStoreThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreThread.Marshal(b, m, deterministic)
}
func (dst *CafeStoreThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreThread.Merge(dst, src)
}
func (m *CafeStoreThread) XXX_Size() int {
	return xxx_messageInfo_CafeStoreThread.Size(m)
}
func (m *CafeStoreThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreThread proto.InternalMessageInfo

func (m *CafeStoreThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStoreThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeStoreThread) GetSkCipher() []byte {
	if m != nil {
		return m.SkCipher
	}
	return nil
}

func (m *CafeStoreThread) GetHeadCipher() []byte {
	if m != nil {
		return m.HeadCipher
	}
	return nil
}

func (m *CafeStoreThread) GetNameCipher() []byte {
	if m != nil {
		return m.NameCipher
	}
	return nil
}

type CafeStored struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStored) Reset()         { *m = CafeStored{} }
func (m *CafeStored) String() string { return proto.CompactTextString(m) }
func (*CafeStored) ProtoMessage()    {}
func (*CafeStored) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{9}
}
func (m *CafeStored) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStored.Unmarshal(m, b)
}
func (m *CafeStored) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStored.Marshal(b, m, deterministic)
}
func (dst *CafeStored) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStored.Merge(dst, src)
}
func (m *CafeStored) XXX_Size() int {
	return xxx_messageInfo_CafeStored.Size(m)
}
func (m *CafeStored) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStored.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStored proto.InternalMessageInfo

func (m *CafeStored) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeDeliverMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeliverMessage) Reset()         { *m = CafeDeliverMessage{} }
func (m *CafeDeliverMessage) String() string { return proto.CompactTextString(m) }
func (*CafeDeliverMessage) ProtoMessage()    {}
func (*CafeDeliverMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{10}
}
func (m *CafeDeliverMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeliverMessage.Unmarshal(m, b)
}
func (m *CafeDeliverMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeliverMessage.Marshal(b, m, deterministic)
}
func (dst *CafeDeliverMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeliverMessage.Merge(dst, src)
}
func (m *CafeDeliverMessage) XXX_Size() int {
	return xxx_messageInfo_CafeDeliverMessage.Size(m)
}
func (m *CafeDeliverMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeliverMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeliverMessage proto.InternalMessageInfo

func (m *CafeDeliverMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeDeliverMessage) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type CafeCheckMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeCheckMessages) Reset()         { *m = CafeCheckMessages{} }
func (m *CafeCheckMessages) String() string { return proto.CompactTextString(m) }
func (*CafeCheckMessages) ProtoMessage()    {}
func (*CafeCheckMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{11}
}
func (m *CafeCheckMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeCheckMessages.Unmarshal(m, b)
}
func (m *CafeCheckMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeCheckMessages.Marshal(b, m, deterministic)
}
func (dst *CafeCheckMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeCheckMessages.Merge(dst, src)
}
func (m *CafeCheckMessages) XXX_Size() int {
	return xxx_messageInfo_CafeCheckMessages.Size(m)
}
func (m *CafeCheckMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeCheckMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeCheckMessages proto.InternalMessageInfo

func (m *CafeCheckMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeMessage struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId               string               `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeMessage) Reset()         { *m = CafeMessage{} }
func (m *CafeMessage) String() string { return proto.CompactTextString(m) }
func (*CafeMessage) ProtoMessage()    {}
func (*CafeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{12}
}
func (m *CafeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessage.Unmarshal(m, b)
}
func (m *CafeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessage.Marshal(b, m, deterministic)
}
func (dst *CafeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessage.Merge(dst, src)
}
func (m *CafeMessage) XXX_Size() int {
	return xxx_messageInfo_CafeMessage.Size(m)
}
func (m *CafeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessage proto.InternalMessageInfo

func (m *CafeMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeMessage) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *CafeMessage) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type CafeMessages struct {
	Messages             []*CafeMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CafeMessages) Reset()         { *m = CafeMessages{} }
func (m *CafeMessages) String() string { return proto.CompactTextString(m) }
func (*CafeMessages) ProtoMessage()    {}
func (*CafeMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{13}
}
func (m *CafeMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessages.Unmarshal(m, b)
}
func (m *CafeMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessages.Marshal(b, m, deterministic)
}
func (dst *CafeMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessages.Merge(dst, src)
}
func (m *CafeMessages) XXX_Size() int {
	return xxx_messageInfo_CafeMessages.Size(m)
}
func (m *CafeMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessages proto.InternalMessageInfo

func (m *CafeMessages) GetMessages() []*CafeMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CafeDeleteMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessages) Reset()         { *m = CafeDeleteMessages{} }
func (m *CafeDeleteMessages) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessages) ProtoMessage()    {}
func (*CafeDeleteMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{14}
}
func (m *CafeDeleteMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessages.Unmarshal(m, b)
}
func (m *CafeDeleteMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessages.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessages.Merge(dst, src)
}
func (m *CafeDeleteMessages) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessages.Size(m)
}
func (m *CafeDeleteMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessages proto.InternalMessageInfo

func (m *CafeDeleteMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeleteMessagesAck struct {
	More                 bool     `protobuf:"varint,2,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessagesAck) Reset()         { *m = CafeDeleteMessagesAck{} }
func (m *CafeDeleteMessagesAck) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessagesAck) ProtoMessage()    {}
func (*CafeDeleteMessagesAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_1b9ac35a00b1c35a, []int{15}
}
func (m *CafeDeleteMessagesAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessagesAck.Unmarshal(m, b)
}
func (m *CafeDeleteMessagesAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessagesAck.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessagesAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessagesAck.Merge(dst, src)
}
func (m *CafeDeleteMessagesAck) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessagesAck.Size(m)
}
func (m *CafeDeleteMessagesAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessagesAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessagesAck proto.InternalMessageInfo

func (m *CafeDeleteMessagesAck) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func init() {
	proto.RegisterType((*CafeChallenge)(nil), "CafeChallenge")
	proto.RegisterType((*CafeNonce)(nil), "CafeNonce")
	proto.RegisterType((*CafeRegistration)(nil), "CafeRegistration")
	proto.RegisterType((*CafeSession)(nil), "CafeSession")
	proto.RegisterType((*CafeRefreshSession)(nil), "CafeRefreshSession")
	proto.RegisterType((*CafeStore)(nil), "CafeStore")
	proto.RegisterType((*CafeBlockList)(nil), "CafeBlockList")
	proto.RegisterType((*CafeBlock)(nil), "CafeBlock")
	proto.RegisterType((*CafeStoreThread)(nil), "CafeStoreThread")
	proto.RegisterType((*CafeStored)(nil), "CafeStored")
	proto.RegisterType((*CafeDeliverMessage)(nil), "CafeDeliverMessage")
	proto.RegisterType((*CafeCheckMessages)(nil), "CafeCheckMessages")
	proto.RegisterType((*CafeMessage)(nil), "CafeMessage")
	proto.RegisterType((*CafeMessages)(nil), "CafeMessages")
	proto.RegisterType((*CafeDeleteMessages)(nil), "CafeDeleteMessages")
	proto.RegisterType((*CafeDeleteMessagesAck)(nil), "CafeDeleteMessagesAck")
}

func init() { proto.RegisterFile("cafe.proto", fileDescriptor_cafe_1b9ac35a00b1c35a) }

var fileDescriptor_cafe_1b9ac35a00b1c35a = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xd2, 0xae, 0xdb, 0xee, 0x0a, 0x8c, 0x08, 0xa6, 0x68, 0x42, 0x50, 0xcc, 0x4b, 0x07,
	0x28, 0x93, 0x86, 0x90, 0x78, 0x64, 0xeb, 0x84, 0x84, 0xc4, 0x78, 0xc8, 0xf6, 0xc4, 0x9b, 0x1b,
	0xdf, 0x26, 0x5e, 0xd3, 0x24, 0xb2, 0xdd, 0x01, 0xdf, 0xc1, 0x47, 0xf2, 0x1b, 0xe8, 0xda, 0x4e,
	0x5a, 0x60, 0x55, 0xdf, 0x7c, 0xae, 0x4f, 0xcf, 0x39, 0xf7, 0xfa, 0xa6, 0x00, 0x19, 0x9f, 0x61,
	0xd2, 0xa8, 0xda, 0xd4, 0xc7, 0x2f, 0xf2, 0xba, 0xce, 0x4b, 0x3c, 0xb5, 0x68, 0xba, 0x9c, 0x9d,
	0x1a, 0xb9, 0x40, 0x6d, 0xf8, 0xa2, 0x71, 0x04, 0x76, 0x02, 0x0f, 0x26, 0x7c, 0x86, 0x93, 0x82,
	0x97, 0x25, 0x56, 0x39, 0x46, 0x31, 0xec, 0x72, 0x21, 0x14, 0x6a, 0x1d, 0x07, 0xa3, 0x60, 0xbc,
	0x9f, 0xb6, 0x90, 0xbd, 0x84, 0x7d, 0xa2, 0x7e, 0xad, 0xab, 0x0c, 0xa3, 0x27, 0xb0, 0x73, 0xc7,
	0xcb, 0x25, 0x7a, 0x92, 0x03, 0xec, 0x16, 0x0e, 0x89, 0x92, 0x62, 0x2e, 0xb5, 0x51, 0xdc, 0xc8,
	0xba, 0xda, 0x2c, 0xb8, 0xd2, 0x08, 0xd7, 0x34, 0xa8, 0x5a, 0x91, 0x45, 0xdc, 0x73, 0x55, 0x0b,
	0xa2, 0x43, 0xe8, 0x69, 0x99, 0xc7, 0xfd, 0x51, 0x30, 0x1e, 0xa6, 0x74, 0x64, 0xbf, 0x03, 0x38,
	0x20, 0xb3, 0x6b, 0xd4, 0x9a, 0x7c, 0x8e, 0x60, 0xc0, 0xb3, 0x6c, 0x65, 0xe3, 0x51, 0xf4, 0x16,
	0x7a, 0xf8, 0xa3, 0xb1, 0x1e, 0x07, 0x67, 0xc7, 0x89, 0x1b, 0x48, 0xd2, 0x0e, 0x24, 0xb9, 0x69,
	0x07, 0x92, 0x12, 0x8d, 0xd2, 0x2a, 0x9c, 0x29, 0xd4, 0x85, 0xf7, 0x6f, 0x61, 0x94, 0x40, 0x5f,
	0x91, 0x50, 0x7f, 0xab, 0x90, 0xe5, 0x91, 0x92, 0x5e, 0x4e, 0x6f, 0x31, 0x33, 0xf1, 0x8e, 0x53,
	0xf2, 0x30, 0x8a, 0xa0, 0x6f, 0x7e, 0x36, 0x18, 0x0f, 0x6c, 0xd9, 0x9e, 0xa3, 0x63, 0xd8, 0x2b,
	0x8c, 0x69, 0xce, 0x85, 0x50, 0xf1, 0xae, 0xad, 0x77, 0x98, 0x7d, 0x82, 0xc8, 0x4d, 0xd5, 0x06,
	0xd9, 0xd6, 0xef, 0x5a, 0x07, 0xe1, 0x5f, 0x1d, 0xb0, 0xf7, 0xee, 0x01, 0xaf, 0x4d, 0xad, 0xec,
	0x98, 0x4d, 0x3d, 0xc7, 0xaa, 0x7d, 0x40, 0x0b, 0x28, 0x5a, 0x26, 0x85, 0x8e, 0xc3, 0x51, 0x8f,
	0xa2, 0xd1, 0x99, 0xbd, 0x72, 0x2b, 0x72, 0x51, 0xd6, 0xd9, 0xfc, 0x8b, 0xd4, 0xe6, 0x5e, 0xd2,
	0x95, 0xd3, 0xb6, 0xa4, 0x0d, 0xda, 0x14, 0x8c, 0x7f, 0xbf, 0xe4, 0x86, 0xdb, 0x60, 0xc3, 0xb4,
	0x85, 0xf4, 0xb8, 0x99, 0x14, 0x7e, 0xe0, 0x74, 0x64, 0xbf, 0x02, 0x78, 0xd4, 0x65, 0xbd, 0x29,
	0x14, 0x72, 0xb1, 0x41, 0xf5, 0x21, 0x84, 0x52, 0xf8, 0x4e, 0x43, 0x29, 0x68, 0x90, 0x7a, 0x3e,
	0x91, 0x4d, 0x81, 0xca, 0x0a, 0x0e, 0xd3, 0x0e, 0x47, 0xcf, 0x01, 0x0a, 0xe4, 0xc2, 0xdf, 0xba,
	0x5d, 0x5a, 0xab, 0xd0, 0x7d, 0xc5, 0x17, 0xe8, 0xef, 0x77, 0xdc, 0xfd, 0xaa, 0xc2, 0x9e, 0x01,
	0x74, 0xa1, 0x84, 0x77, 0x0e, 0x5a, 0x67, 0xf6, 0xd1, 0x3d, 0xd3, 0x25, 0x96, 0xf2, 0x0e, 0xd5,
	0x15, 0x6a, 0xcd, 0x73, 0xfc, 0x97, 0x45, 0xf9, 0xb2, 0x52, 0x62, 0x65, 0x3e, 0xb7, 0xa9, 0x3b,
	0xcc, 0x4e, 0xe0, 0xb1, 0xfb, 0x18, 0x31, 0x9b, 0xfb, 0xdf, 0xeb, 0xfb, 0xdb, 0x66, 0xe8, 0x96,
	0x7f, 0x93, 0xcb, 0x11, 0x0c, 0x1a, 0x44, 0xd5, 0x79, 0x78, 0x44, 0x4b, 0x2c, 0xb8, 0x71, 0xdf,
	0xd6, 0x96, 0x25, 0x26, 0x1e, 0xfb, 0x00, 0xc3, 0x35, 0x1b, 0x1d, 0x8d, 0x61, 0x6f, 0xe1, 0xcf,
	0x71, 0x30, 0xea, 0x8d, 0x0f, 0xce, 0x86, 0xc9, 0x1a, 0x21, 0xed, 0x6e, 0xd9, 0xeb, 0x6e, 0x1a,
	0x68, 0x70, 0x4b, 0x33, 0x6f, 0xe0, 0xe9, 0xff, 0xdc, 0xf3, 0x6c, 0x4e, 0x9b, 0xb6, 0xa8, 0x95,
	0xfb, 0x83, 0xd8, 0x4b, 0xed, 0xf9, 0xa2, 0xff, 0x2d, 0x6c, 0xa6, 0xd3, 0x81, 0x8d, 0xfc, 0xee,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x01, 0x97, 0x8e, 0xed, 0x04, 0x00, 0x00,
}
