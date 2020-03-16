// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users.proto

package go_micro_srv_users

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

type EventType int32

const (
	EventType_Unknown     EventType = 0
	EventType_UserCreated EventType = 1
	EventType_UserUpdated EventType = 2
	EventType_UserDeleted EventType = 3
)

var EventType_name = map[int32]string{
	0: "Unknown",
	1: "UserCreated",
	2: "UserUpdated",
	3: "UserDeleted",
}

var EventType_value = map[string]int32{
	"Unknown":     0,
	"UserCreated": 1,
	"UserUpdated": 2,
	"UserDeleted": 3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{0}
}

type Event struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=go.micro.srv.users.EventType" json:"type,omitempty"`
	User                 *User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *Event) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64             `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	FirstName            string            `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string            `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string            `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Username             string            `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ValidateOnly         bool     `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

type CreateResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{4}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{6}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{7}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SearchRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{10}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type SearchResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{11}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.srv.users.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "go.micro.srv.users.Event")
	proto.RegisterType((*User)(nil), "go.micro.srv.users.User")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.srv.users.User.MetadataEntry")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.users.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.users.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.users.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.users.DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.users.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.users.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.users.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.users.UpdateResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.users.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.users.SearchResponse")
}

func init() { proto.RegisterFile("proto/users.proto", fileDescriptor_b1c161a4c7514913) }

var fileDescriptor_b1c161a4c7514913 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xed, 0x38, 0x1f, 0x93, 0xda, 0x98, 0x55, 0x0f, 0x56, 0x50, 0xd4, 0xb0, 0x48, 0x28,
	0x02, 0x64, 0x44, 0xb8, 0x20, 0xbe, 0x84, 0x80, 0x9e, 0x2a, 0x40, 0x35, 0xcd, 0xb9, 0xda, 0xc6,
	0x03, 0xb5, 0xe2, 0xd8, 0xc1, 0x76, 0x82, 0xfc, 0x6f, 0xf8, 0x1b, 0xfc, 0x3b, 0xb4, 0xb3, 0xd9,
	0x50, 0x17, 0x07, 0x41, 0x6e, 0x9e, 0x79, 0x6f, 0xde, 0xcc, 0xce, 0xbc, 0x04, 0x6e, 0x2f, 0xf3,
	0xac, 0xcc, 0x1e, 0xaf, 0x0a, 0xcc, 0x8b, 0x80, 0xbe, 0x19, 0xfb, 0x9a, 0x05, 0x8b, 0x78, 0x96,
	0x67, 0x41, 0x91, 0xaf, 0x03, 0x42, 0xf8, 0x15, 0xd8, 0x27, 0x6b, 0x4c, 0x4b, 0xf6, 0x04, 0x5a,
	0x65, 0xb5, 0x44, 0xdf, 0x18, 0x19, 0x63, 0x77, 0x32, 0x0c, 0xfe, 0xe4, 0x06, 0x44, 0x3c, 0xaf,
	0x96, 0x18, 0x12, 0x95, 0x3d, 0x82, 0x96, 0x04, 0x7c, 0x73, 0x64, 0x8c, 0xfb, 0x13, 0xbf, 0xa9,
	0x64, 0x5a, 0x60, 0x1e, 0x12, 0x8b, 0xff, 0x34, 0xa1, 0x25, 0x43, 0xe6, 0x82, 0x19, 0x47, 0xd4,
	0xa7, 0x17, 0x9a, 0x71, 0xc4, 0x7c, 0xe8, 0xcc, 0x72, 0x14, 0x25, 0x46, 0xa4, 0x64, 0x85, 0x3a,
	0x94, 0xc8, 0x6a, 0x19, 0x11, 0x62, 0x29, 0x64, 0x13, 0xb2, 0x21, 0xc0, 0x97, 0x38, 0x2f, 0xca,
	0x8b, 0x54, 0x2c, 0xd0, 0x6f, 0x91, 0x56, 0x8f, 0x32, 0x1f, 0xc5, 0x02, 0xd9, 0x1d, 0xe8, 0x25,
	0x42, 0xa3, 0x36, 0xa1, 0x5d, 0x99, 0x20, 0xf0, 0x08, 0x6c, 0x5c, 0x88, 0x38, 0xf1, 0xdb, 0x04,
	0xa8, 0x80, 0x0d, 0xa0, 0x2b, 0xc7, 0xa4, 0x8a, 0x8e, 0xaa, 0xd0, 0x31, 0x7b, 0x0b, 0xdd, 0x05,
	0x96, 0x22, 0x12, 0xa5, 0xf0, 0xbb, 0x23, 0x6b, 0xdc, 0x9f, 0xdc, 0xdf, 0xf5, 0xd8, 0xe0, 0xc3,
	0x86, 0x78, 0x92, 0x96, 0x79, 0x15, 0x6e, 0xeb, 0x06, 0x2f, 0xc0, 0xa9, 0x41, 0xcc, 0x03, 0x6b,
	0x8e, 0xd5, 0x66, 0x0f, 0xf2, 0x53, 0x0e, 0xb6, 0x16, 0xc9, 0x0a, 0x69, 0x0d, 0xbd, 0x50, 0x05,
	0xcf, 0xcd, 0x67, 0x06, 0xbf, 0x04, 0xe7, 0x1d, 0xed, 0x24, 0xc4, 0x6f, 0x2b, 0x2c, 0xca, 0xed,
	0xea, 0x8d, 0x7f, 0x59, 0x3d, 0xbb, 0x07, 0xce, 0x5a, 0x24, 0xb1, 0x5c, 0xdd, 0x45, 0x96, 0x26,
	0x15, 0x35, 0xe8, 0x86, 0x87, 0x3a, 0xf9, 0x29, 0x4d, 0x2a, 0x7e, 0x0e, 0xae, 0xee, 0x51, 0x2c,
	0xb3, 0xb4, 0xc0, 0xff, 0x6c, 0x72, 0x04, 0x76, 0x99, 0xcd, 0x31, 0xd5, 0xd3, 0x53, 0xc0, 0x8f,
	0xc1, 0x79, 0x8f, 0x09, 0xfe, 0x9e, 0xfc, 0xc6, 0xf5, 0xb9, 0x07, 0xae, 0x26, 0xa8, 0xb6, 0x7c,
	0x08, 0xfd, 0x10, 0x45, 0xb4, 0xab, 0xe0, 0x25, 0x1c, 0x2a, 0x78, 0x9f, 0x29, 0xf9, 0x2b, 0x70,
	0xa6, 0xe4, 0xa1, 0xbd, 0x36, 0xc9, 0x5f, 0x83, 0xab, 0xcb, 0xf7, 0x6a, 0xff, 0x10, 0x9c, 0xcf,
	0x28, 0xf2, 0xd9, 0x95, 0x6e, 0x7f, 0xdd, 0x76, 0x46, 0xdd, 0x76, 0xfc, 0x0d, 0xb8, 0x9a, 0xbc,
	0x69, 0x16, 0x80, 0x4d, 0x92, 0xbe, 0x41, 0x2e, 0xdc, 0xdd, 0x4d, 0xd1, 0x1e, 0x9c, 0x42, 0x6f,
	0xfb, 0xa3, 0x65, 0x7d, 0xe8, 0x4c, 0xd3, 0x79, 0x9a, 0x7d, 0x4f, 0xbd, 0x03, 0x76, 0x0b, 0xfa,
	0x92, 0xa8, 0x2e, 0x1e, 0x79, 0x86, 0x4e, 0xa8, 0xd7, 0x45, 0x9e, 0xa9, 0x13, 0xea, 0x38, 0x91,
	0x67, 0x4d, 0x7e, 0x58, 0x60, 0xcb, 0x4c, 0xc1, 0xce, 0xa0, 0xad, 0x0a, 0xd9, 0xdd, 0xa6, 0x09,
	0x6a, 0x56, 0x1d, 0xf0, 0xbf, 0x51, 0x36, 0x27, 0x3f, 0x60, 0xa7, 0xd0, 0x92, 0x57, 0x65, 0xc7,
	0x4d, 0xec, 0x6b, 0x76, 0x18, 0x8c, 0x76, 0x13, 0xb6, 0x62, 0x67, 0xd0, 0x56, 0xef, 0x68, 0x9e,
	0xaf, 0x66, 0x80, 0xe6, 0xf9, 0xea, 0x47, 0x56, 0x92, 0x6a, 0x13, 0xcd, 0x92, 0x35, 0x8f, 0x37,
	0x4b, 0xde, 0x70, 0x39, 0x49, 0xaa, 0xf3, 0x36, 0x4b, 0xd6, 0x7c, 0xd2, 0x2c, 0x59, 0x77, 0x07,
	0x3f, 0xb8, 0x6c, 0xd3, 0x1f, 0xfd, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x07, 0x41,
	0x2d, 0xfd, 0x05, 0x00, 0x00,
}