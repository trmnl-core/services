// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/explore/proto/explore/explore.proto

package go_micro_srv_explore

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a718096c7c133a, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.explore.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.explore.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.explore.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.explore.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.explore.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.explore.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.explore.Pong")
}

func init() {
	proto.RegisterFile("github.com/micro/services/explore/proto/explore/explore.proto", fileDescriptor_01a718096c7c133a)
}

var fileDescriptor_01a718096c7c133a = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0xbb, 0xb4, 0x5f, 0xdb, 0x6f, 0x4e, 0x75, 0x29, 0x22, 0xd1, 0x16, 0xd9, 0x83, 0xd6,
	0xcb, 0xa6, 0xe8, 0xd9, 0x93, 0x88, 0x5e, 0x04, 0x89, 0x67, 0x0f, 0x69, 0x18, 0xd6, 0x60, 0x77,
	0xb7, 0xee, 0x6c, 0x8a, 0xfe, 0xed, 0x5e, 0x24, 0x9b, 0x8d, 0x88, 0x24, 0x78, 0xca, 0x4c, 0x7e,
	0xef, 0x0d, 0xf3, 0x86, 0x85, 0x6b, 0x55, 0xfa, 0x97, 0x6a, 0x23, 0x0b, 0xab, 0x53, 0x5d, 0x16,
	0xce, 0xa6, 0x84, 0x6e, 0x5f, 0x16, 0x48, 0x29, 0xbe, 0xef, 0xb6, 0xd6, 0x61, 0xba, 0x73, 0xd6,
	0xdb, 0xef, 0x2e, 0x7e, 0x65, 0xf8, 0xcb, 0xe7, 0xca, 0xca, 0x60, 0x93, 0xe4, 0xf6, 0x32, 0x32,
	0x71, 0x0c, 0x93, 0x07, 0x24, 0xca, 0x15, 0xf2, 0x19, 0x0c, 0x29, 0xff, 0x38, 0x62, 0xa7, 0x6c,
	0xf5, 0x3f, 0xab, 0x4b, 0xb1, 0x80, 0x49, 0x86, 0x6f, 0x15, 0x92, 0xe7, 0x1c, 0x46, 0x26, 0xd7,
	0x18, 0x69, 0xa8, 0xc5, 0x09, 0x4c, 0x33, 0xa4, 0x9d, 0x35, 0x14, 0xcc, 0x9a, 0x54, 0x6b, 0xd6,
	0xa4, 0xc4, 0x0a, 0x66, 0x4f, 0xde, 0x61, 0xae, 0x4b, 0xa3, 0xda, 0x29, 0x73, 0xf8, 0x57, 0xd8,
	0xca, 0xf8, 0xa0, 0x1b, 0x66, 0x4d, 0x23, 0x2e, 0xe0, 0xe0, 0x87, 0x32, 0x0e, 0xec, 0x96, 0x2e,
	0x61, 0xf4, 0x58, 0x1a, 0xc5, 0x0f, 0x61, 0x4c, 0xde, 0xd9, 0x57, 0x8c, 0x38, 0x76, 0x81, 0xdb,
	0x7e, 0x7e, 0xf9, 0xc9, 0x60, 0x72, 0xdb, 0x44, 0xe7, 0x77, 0x30, 0xba, 0xc9, 0xb7, 0x5b, 0xbe,
	0x90, 0x5d, 0x97, 0x91, 0x71, 0xe7, 0x64, 0xd9, 0x87, 0x9b, 0x45, 0xc5, 0x80, 0x3f, 0xc3, 0xb8,
	0xd9, 0x9f, 0x9f, 0x75, 0x6b, 0x7f, 0xdf, 0x21, 0x39, 0xff, 0x53, 0xd7, 0x0e, 0x5f, 0x33, 0x7e,
	0x0f, 0xd3, 0x3a, 0x73, 0xc8, 0x95, 0x74, 0x1b, 0x6b, 0x9e, 0xf4, 0x31, 0x6b, 0x94, 0x18, 0xac,
	0xd8, 0x9a, 0x6d, 0xc6, 0xe1, 0x25, 0x5c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x82, 0x1e,
	0x17, 0x4a, 0x02, 0x00, 0x00,
}