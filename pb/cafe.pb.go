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
	return fileDescriptor_cafe_993e84935f545c77, []int{0}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{1}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{2}
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
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeSession) Reset()         { *m = CafeSession{} }
func (m *CafeSession) String() string { return proto.CompactTextString(m) }
func (*CafeSession) ProtoMessage()    {}
func (*CafeSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_993e84935f545c77, []int{3}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{4}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{5}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{6}
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
	return fileDescriptor_cafe_993e84935f545c77, []int{7}
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

type CafeStored struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStored) Reset()         { *m = CafeStored{} }
func (m *CafeStored) String() string { return proto.CompactTextString(m) }
func (*CafeStored) ProtoMessage()    {}
func (*CafeStored) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_993e84935f545c77, []int{8}
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

func (m *CafeStored) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type CafeAddThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	SkCipher             []byte   `protobuf:"bytes,3,opt,name=skCipher,proto3" json:"skCipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeAddThread) Reset()         { *m = CafeAddThread{} }
func (m *CafeAddThread) String() string { return proto.CompactTextString(m) }
func (*CafeAddThread) ProtoMessage()    {}
func (*CafeAddThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_993e84935f545c77, []int{9}
}
func (m *CafeAddThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeAddThread.Unmarshal(m, b)
}
func (m *CafeAddThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeAddThread.Marshal(b, m, deterministic)
}
func (dst *CafeAddThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeAddThread.Merge(dst, src)
}
func (m *CafeAddThread) XXX_Size() int {
	return xxx_messageInfo_CafeAddThread.Size(m)
}
func (m *CafeAddThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeAddThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeAddThread proto.InternalMessageInfo

func (m *CafeAddThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeAddThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeAddThread) GetSkCipher() []byte {
	if m != nil {
		return m.SkCipher
	}
	return nil
}

type CafeRemoveThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRemoveThread) Reset()         { *m = CafeRemoveThread{} }
func (m *CafeRemoveThread) String() string { return proto.CompactTextString(m) }
func (*CafeRemoveThread) ProtoMessage()    {}
func (*CafeRemoveThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_993e84935f545c77, []int{10}
}
func (m *CafeRemoveThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRemoveThread.Unmarshal(m, b)
}
func (m *CafeRemoveThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRemoveThread.Marshal(b, m, deterministic)
}
func (dst *CafeRemoveThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRemoveThread.Merge(dst, src)
}
func (m *CafeRemoveThread) XXX_Size() int {
	return xxx_messageInfo_CafeRemoveThread.Size(m)
}
func (m *CafeRemoveThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRemoveThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRemoveThread proto.InternalMessageInfo

func (m *CafeRemoveThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeRemoveThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeUpdateThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Head                 string   `protobuf:"bytes,3,opt,name=head,proto3" json:"head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeUpdateThread) Reset()         { *m = CafeUpdateThread{} }
func (m *CafeUpdateThread) String() string { return proto.CompactTextString(m) }
func (*CafeUpdateThread) ProtoMessage()    {}
func (*CafeUpdateThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_993e84935f545c77, []int{11}
}
func (m *CafeUpdateThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeUpdateThread.Unmarshal(m, b)
}
func (m *CafeUpdateThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeUpdateThread.Marshal(b, m, deterministic)
}
func (dst *CafeUpdateThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeUpdateThread.Merge(dst, src)
}
func (m *CafeUpdateThread) XXX_Size() int {
	return xxx_messageInfo_CafeUpdateThread.Size(m)
}
func (m *CafeUpdateThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeUpdateThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeUpdateThread proto.InternalMessageInfo

func (m *CafeUpdateThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeUpdateThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeUpdateThread) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
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
	proto.RegisterType((*CafeStored)(nil), "CafeStored")
	proto.RegisterType((*CafeAddThread)(nil), "CafeAddThread")
	proto.RegisterType((*CafeRemoveThread)(nil), "CafeRemoveThread")
	proto.RegisterType((*CafeUpdateThread)(nil), "CafeUpdateThread")
}

func init() { proto.RegisterFile("cafe.proto", fileDescriptor_cafe_993e84935f545c77) }

var fileDescriptor_cafe_993e84935f545c77 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xd2, 0xac, 0xb0, 0x6f, 0x05, 0x4d, 0x16, 0x42, 0x51, 0x0f, 0x50, 0xcc, 0x65, 0x48,
	0x28, 0x93, 0x40, 0x48, 0x5c, 0x59, 0x11, 0xa7, 0x81, 0x44, 0x36, 0x2e, 0xdc, 0xdc, 0xf8, 0x4b,
	0xe2, 0x35, 0x8d, 0x23, 0xdb, 0x1d, 0xf0, 0x1b, 0xf9, 0x53, 0xe8, 0xb3, 0xe3, 0x75, 0x48, 0xab,
	0xa6, 0xdd, 0xbe, 0x67, 0x3f, 0xbf, 0xf7, 0xfc, 0x6c, 0x80, 0x4a, 0xd4, 0x58, 0x0c, 0x46, 0x3b,
	0x3d, 0x7f, 0xd9, 0x68, 0xdd, 0x74, 0x78, 0xea, 0xd1, 0x6a, 0x5b, 0x9f, 0x3a, 0xb5, 0x41, 0xeb,
	0xc4, 0x66, 0x08, 0x04, 0xfe, 0x06, 0x9e, 0x2c, 0x45, 0x8d, 0xcb, 0x56, 0x74, 0x1d, 0xf6, 0x0d,
	0xb2, 0x1c, 0x1e, 0x09, 0x29, 0x0d, 0x5a, 0x9b, 0x27, 0x8b, 0xe4, 0xe4, 0xb0, 0x8c, 0x90, 0xbf,
	0x82, 0x43, 0xa2, 0x7e, 0xd3, 0x7d, 0x85, 0xec, 0x19, 0x1c, 0x5c, 0x8b, 0x6e, 0x8b, 0x23, 0x29,
	0x00, 0x7e, 0x05, 0xc7, 0x44, 0x29, 0xb1, 0x51, 0xd6, 0x19, 0xe1, 0x94, 0xee, 0xf7, 0x0b, 0xee,
	0x34, 0xd2, 0x5b, 0x1a, 0xb4, 0xda, 0x93, 0x45, 0x3e, 0x09, 0xab, 0x1e, 0xb0, 0x63, 0x98, 0x58,
	0xd5, 0xe4, 0xd9, 0x22, 0x39, 0x99, 0x95, 0x34, 0xf2, 0xbf, 0x09, 0x1c, 0x91, 0xd9, 0x05, 0x5a,
	0x4b, 0x3e, 0xcf, 0x61, 0x2a, 0xaa, 0x6a, 0x67, 0x33, 0x22, 0xf6, 0x16, 0x26, 0xf8, 0x7b, 0xf0,
	0x1e, 0x47, 0xef, 0xe6, 0x45, 0x28, 0xa4, 0x88, 0x85, 0x14, 0x97, 0xb1, 0x90, 0x92, 0x68, 0x94,
	0xd6, 0x60, 0x6d, 0xd0, 0xb6, 0xa3, 0x7f, 0x84, 0xac, 0x80, 0xcc, 0x90, 0x50, 0x76, 0xaf, 0x90,
	0xe7, 0x91, 0x92, 0xdd, 0xae, 0xae, 0xb0, 0x72, 0xf9, 0x41, 0x50, 0x1a, 0x21, 0x63, 0x90, 0xb9,
	0x3f, 0x03, 0xe6, 0x53, 0xbf, 0xec, 0x67, 0xfe, 0x05, 0x58, 0x68, 0xce, 0x9b, 0xdd, 0x77, 0xa7,
	0x5b, 0x29, 0xd3, 0xff, 0x52, 0xf2, 0x0f, 0xe1, 0x91, 0x2e, 0x9c, 0x36, 0xbe, 0x4a, 0xa7, 0xd7,
	0xd8, 0xc7, 0x47, 0xf2, 0x80, 0xec, 0x2b, 0x25, 0x6d, 0x9e, 0x2e, 0x26, 0x64, 0x4f, 0x33, 0x7f,
	0x1d, 0xbe, 0xc1, 0x59, 0xa7, 0xab, 0xf5, 0xb9, 0xb2, 0xee, 0x4e, 0xd2, 0xd7, 0xa0, 0xed, 0x49,
	0x7b, 0xb4, 0x29, 0x98, 0xf8, 0xf5, 0x59, 0x38, 0xe1, 0x83, 0xcd, 0xca, 0x08, 0xe9, 0x01, 0x2b,
	0x25, 0xc7, 0x52, 0x69, 0xe4, 0x2f, 0x00, 0x6e, 0xa2, 0xca, 0xb8, 0x9f, 0xec, 0xf6, 0xbf, 0x87,
	0x4c, 0x9f, 0xa4, 0xbc, 0x6c, 0x0d, 0x0a, 0xb9, 0xc7, 0xf2, 0x29, 0xa4, 0x4a, 0x8e, 0x35, 0xa4,
	0x4a, 0xb2, 0x39, 0x3c, 0xb6, 0xeb, 0xa5, 0x1a, 0x5a, 0x34, 0xde, 0x6d, 0x56, 0xde, 0x60, 0xfe,
	0x31, 0xfe, 0xcf, 0x8d, 0xbe, 0xc6, 0x87, 0xa8, 0xf2, 0xf3, 0x70, 0xf2, 0xc7, 0x20, 0x85, 0x7b,
	0xd0, 0x49, 0x6a, 0xb2, 0x45, 0x11, 0x6f, 0xee, 0xe7, 0xb3, 0xec, 0x67, 0x3a, 0xac, 0x56, 0x53,
	0xff, 0x77, 0xde, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x12, 0x99, 0x2b, 0x68, 0xb1, 0x03, 0x00,
	0x00,
}
