// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeSkewRequest struct {
	SourceBucket         string   `protobuf:"bytes,1,opt,name=source_bucket,json=sourceBucket,proto3" json:"source_bucket,omitempty"`
	SourceKey            string   `protobuf:"bytes,2,opt,name=source_key,json=sourceKey,proto3" json:"source_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeSkewRequest) Reset()         { *m = DeSkewRequest{} }
func (m *DeSkewRequest) String() string { return proto.CompactTextString(m) }
func (*DeSkewRequest) ProtoMessage()    {}
func (*DeSkewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *DeSkewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeSkewRequest.Unmarshal(m, b)
}
func (m *DeSkewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeSkewRequest.Marshal(b, m, deterministic)
}
func (m *DeSkewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeSkewRequest.Merge(m, src)
}
func (m *DeSkewRequest) XXX_Size() int {
	return xxx_messageInfo_DeSkewRequest.Size(m)
}
func (m *DeSkewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeSkewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeSkewRequest proto.InternalMessageInfo

func (m *DeSkewRequest) GetSourceBucket() string {
	if m != nil {
		return m.SourceBucket
	}
	return ""
}

func (m *DeSkewRequest) GetSourceKey() string {
	if m != nil {
		return m.SourceKey
	}
	return ""
}

type DeSkewResponse struct {
	NewBucket            string   `protobuf:"bytes,1,opt,name=new_bucket,json=newBucket,proto3" json:"new_bucket,omitempty"`
	NewKey               string   `protobuf:"bytes,2,opt,name=new_key,json=newKey,proto3" json:"new_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeSkewResponse) Reset()         { *m = DeSkewResponse{} }
func (m *DeSkewResponse) String() string { return proto.CompactTextString(m) }
func (*DeSkewResponse) ProtoMessage()    {}
func (*DeSkewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *DeSkewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeSkewResponse.Unmarshal(m, b)
}
func (m *DeSkewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeSkewResponse.Marshal(b, m, deterministic)
}
func (m *DeSkewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeSkewResponse.Merge(m, src)
}
func (m *DeSkewResponse) XXX_Size() int {
	return xxx_messageInfo_DeSkewResponse.Size(m)
}
func (m *DeSkewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeSkewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeSkewResponse proto.InternalMessageInfo

func (m *DeSkewResponse) GetNewBucket() string {
	if m != nil {
		return m.NewBucket
	}
	return ""
}

func (m *DeSkewResponse) GetNewKey() string {
	if m != nil {
		return m.NewKey
	}
	return ""
}

func init() {
	proto.RegisterType((*DeSkewRequest)(nil), "DeSkewRequest")
	proto.RegisterType((*DeSkewResponse)(nil), "DeSkewResponse")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xf3, 0x94, 0x82, 0xb9, 0x78,
	0x5d, 0x52, 0x83, 0xb3, 0x53, 0xcb, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x94, 0xb9,
	0x78, 0x8b, 0xf3, 0x4b, 0x8b, 0x92, 0x53, 0xe3, 0x93, 0x4a, 0x93, 0xb3, 0x53, 0x4b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0x20, 0x82, 0x4e, 0x60, 0x31, 0x21, 0x59, 0x2e, 0x2e, 0xa8,
	0xa2, 0xec, 0xd4, 0x4a, 0x09, 0x26, 0xb0, 0x0a, 0x4e, 0x88, 0x88, 0x77, 0x6a, 0xa5, 0x92, 0x07,
	0x17, 0x1f, 0xcc, 0xd0, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0x86, 0xbc, 0xd4, 0x72, 0x54,
	0x23, 0x39, 0xf3, 0x52, 0xcb, 0xa1, 0xe6, 0x89, 0x73, 0xb1, 0x83, 0xa4, 0x11, 0x86, 0xb1, 0xe5,
	0xa5, 0x96, 0x7b, 0xa7, 0x56, 0x1a, 0x59, 0x71, 0xb1, 0x41, 0x4c, 0x12, 0x32, 0xe0, 0xe2, 0x76,
	0x49, 0x2d, 0xce, 0x4e, 0x2d, 0xf7, 0xcc, 0x4d, 0x4c, 0x4f, 0x15, 0xe2, 0xd3, 0x43, 0x71, 0xb6,
	0x14, 0xbf, 0x1e, 0xaa, 0x8d, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x1f, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0x03, 0x4b, 0xac, 0xf8, 0x00, 0x00, 0x00,
}
