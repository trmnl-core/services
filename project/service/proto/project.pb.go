// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/project/service/proto/project.proto

package go_micro_service_project

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

type Project struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string    `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WebDomain            string    `protobuf:"bytes,4,opt,name=web_domain,json=webDomain,proto3" json:"web_domain,omitempty"`
	ApiDomain            string    `protobuf:"bytes,5,opt,name=api_domain,json=apiDomain,proto3" json:"api_domain,omitempty"`
	Members              []*Member `protobuf:"bytes,6,rep,name=members,proto3" json:"members,omitempty"`
	Repository           string    `protobuf:"bytes,7,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Project) GetWebDomain() string {
	if m != nil {
		return m.WebDomain
	}
	return ""
}

func (m *Project) GetApiDomain() string {
	if m != nil {
		return m.ApiDomain
	}
	return ""
}

func (m *Project) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Project) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

type Member struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{1}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{2}
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

func (m *ReadRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ReadResponse struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{3}
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

func (m *ReadResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type CreateRequest struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{4}
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

func (m *CreateRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type CreateResponse struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{5}
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

func (m *CreateResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateRequest struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{6}
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

func (m *UpdateRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{7}
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

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type AddMemberRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	MemberId             string   `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMemberRequest) Reset()         { *m = AddMemberRequest{} }
func (m *AddMemberRequest) String() string { return proto.CompactTextString(m) }
func (*AddMemberRequest) ProtoMessage()    {}
func (*AddMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{10}
}

func (m *AddMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMemberRequest.Unmarshal(m, b)
}
func (m *AddMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMemberRequest.Marshal(b, m, deterministic)
}
func (m *AddMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMemberRequest.Merge(m, src)
}
func (m *AddMemberRequest) XXX_Size() int {
	return xxx_messageInfo_AddMemberRequest.Size(m)
}
func (m *AddMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMemberRequest proto.InternalMessageInfo

func (m *AddMemberRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *AddMemberRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type AddMemberResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMemberResponse) Reset()         { *m = AddMemberResponse{} }
func (m *AddMemberResponse) String() string { return proto.CompactTextString(m) }
func (*AddMemberResponse) ProtoMessage()    {}
func (*AddMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{11}
}

func (m *AddMemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMemberResponse.Unmarshal(m, b)
}
func (m *AddMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMemberResponse.Marshal(b, m, deterministic)
}
func (m *AddMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMemberResponse.Merge(m, src)
}
func (m *AddMemberResponse) XXX_Size() int {
	return xxx_messageInfo_AddMemberResponse.Size(m)
}
func (m *AddMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMemberResponse proto.InternalMessageInfo

type RemoveMemberRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	MemberId             string   `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveMemberRequest) Reset()         { *m = RemoveMemberRequest{} }
func (m *RemoveMemberRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveMemberRequest) ProtoMessage()    {}
func (*RemoveMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{12}
}

func (m *RemoveMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveMemberRequest.Unmarshal(m, b)
}
func (m *RemoveMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveMemberRequest.Marshal(b, m, deterministic)
}
func (m *RemoveMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMemberRequest.Merge(m, src)
}
func (m *RemoveMemberRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveMemberRequest.Size(m)
}
func (m *RemoveMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMemberRequest proto.InternalMessageInfo

func (m *RemoveMemberRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *RemoveMemberRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type RemoveMemberResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveMemberResponse) Reset()         { *m = RemoveMemberResponse{} }
func (m *RemoveMemberResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveMemberResponse) ProtoMessage()    {}
func (*RemoveMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{13}
}

func (m *RemoveMemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveMemberResponse.Unmarshal(m, b)
}
func (m *RemoveMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveMemberResponse.Marshal(b, m, deterministic)
}
func (m *RemoveMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMemberResponse.Merge(m, src)
}
func (m *RemoveMemberResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveMemberResponse.Size(m)
}
func (m *RemoveMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMemberResponse proto.InternalMessageInfo

type ListMembershipsRequest struct {
	MemberId             string   `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMembershipsRequest) Reset()         { *m = ListMembershipsRequest{} }
func (m *ListMembershipsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMembershipsRequest) ProtoMessage()    {}
func (*ListMembershipsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{14}
}

func (m *ListMembershipsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembershipsRequest.Unmarshal(m, b)
}
func (m *ListMembershipsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembershipsRequest.Marshal(b, m, deterministic)
}
func (m *ListMembershipsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembershipsRequest.Merge(m, src)
}
func (m *ListMembershipsRequest) XXX_Size() int {
	return xxx_messageInfo_ListMembershipsRequest.Size(m)
}
func (m *ListMembershipsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembershipsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembershipsRequest proto.InternalMessageInfo

func (m *ListMembershipsRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type ListMembershipsResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMembershipsResponse) Reset()         { *m = ListMembershipsResponse{} }
func (m *ListMembershipsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMembershipsResponse) ProtoMessage()    {}
func (*ListMembershipsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea461ee6def96e0, []int{15}
}

func (m *ListMembershipsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembershipsResponse.Unmarshal(m, b)
}
func (m *ListMembershipsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembershipsResponse.Marshal(b, m, deterministic)
}
func (m *ListMembershipsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembershipsResponse.Merge(m, src)
}
func (m *ListMembershipsResponse) XXX_Size() int {
	return xxx_messageInfo_ListMembershipsResponse.Size(m)
}
func (m *ListMembershipsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembershipsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembershipsResponse proto.InternalMessageInfo

func (m *ListMembershipsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "go.micro.service.project.Project")
	proto.RegisterType((*Member)(nil), "go.micro.service.project.Member")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.service.project.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.service.project.ReadResponse")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.project.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.project.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.service.project.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.service.project.UpdateResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.service.project.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.service.project.ListResponse")
	proto.RegisterType((*AddMemberRequest)(nil), "go.micro.service.project.AddMemberRequest")
	proto.RegisterType((*AddMemberResponse)(nil), "go.micro.service.project.AddMemberResponse")
	proto.RegisterType((*RemoveMemberRequest)(nil), "go.micro.service.project.RemoveMemberRequest")
	proto.RegisterType((*RemoveMemberResponse)(nil), "go.micro.service.project.RemoveMemberResponse")
	proto.RegisterType((*ListMembershipsRequest)(nil), "go.micro.service.project.ListMembershipsRequest")
	proto.RegisterType((*ListMembershipsResponse)(nil), "go.micro.service.project.ListMembershipsResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/project/service/proto/project.proto", fileDescriptor_7ea461ee6def96e0)
}

var fileDescriptor_7ea461ee6def96e0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x4d, 0xd2, 0x34, 0xa9, 0x27, 0x6d, 0x28, 0x5b, 0x54, 0x2c, 0x43, 0x51, 0x58, 0x09, 0x88,
	0x40, 0x38, 0x10, 0xc4, 0x0b, 0x55, 0x1f, 0x10, 0xbc, 0x54, 0x34, 0x08, 0x8c, 0x10, 0xbc, 0xa0,
	0xca, 0x97, 0xa5, 0x5d, 0x24, 0x67, 0x8d, 0xd7, 0x6d, 0xe1, 0x3b, 0xf9, 0x06, 0xfe, 0x03, 0xed,
	0xcd, 0xb5, 0x1d, 0xd9, 0x0d, 0x6a, 0x9e, 0xb2, 0x7b, 0x66, 0xce, 0x99, 0xcb, 0xce, 0xc4, 0x70,
	0x70, 0x42, 0xb3, 0xd3, 0xb3, 0xc0, 0x0d, 0x59, 0x3c, 0x89, 0x69, 0x98, 0xb2, 0x09, 0x27, 0xe9,
	0x39, 0x0d, 0x09, 0x9f, 0x24, 0x29, 0xfb, 0x41, 0xc2, 0xcc, 0x00, 0xe2, 0x9e, 0x31, 0x83, 0xba,
	0xf2, 0x86, 0xec, 0x13, 0xe6, 0x4a, 0x9a, 0xab, 0xbd, 0x5c, 0x6d, 0xc7, 0x7f, 0xdb, 0xd0, 0xff,
	0xa0, 0xce, 0x68, 0x08, 0x1d, 0x1a, 0xd9, 0xed, 0x51, 0x7b, 0x6c, 0x79, 0x1d, 0x1a, 0x21, 0x04,
	0xdd, 0xb9, 0x1f, 0x13, 0xbb, 0x23, 0x11, 0x79, 0x46, 0x77, 0xc1, 0x12, 0xbf, 0x3c, 0xf1, 0x43,
	0x62, 0xaf, 0x49, 0xc3, 0x25, 0x80, 0xf6, 0x00, 0x2e, 0x48, 0x70, 0x1c, 0xb1, 0xd8, 0xa7, 0x73,
	0xbb, 0xab, 0xcc, 0x17, 0x24, 0x78, 0x2b, 0x01, 0x61, 0xf6, 0x13, 0x6a, 0xcc, 0xeb, 0xca, 0xec,
	0x27, 0x54, 0x9b, 0x5f, 0x41, 0x3f, 0x26, 0x71, 0x40, 0x52, 0x6e, 0xf7, 0x46, 0x6b, 0xe3, 0xc1,
	0x74, 0xe4, 0xd6, 0xe5, 0xed, 0xce, 0xa4, 0xa3, 0x67, 0x08, 0xe8, 0x1e, 0x40, 0x4a, 0x12, 0xc6,
	0x69, 0xc6, 0xd2, 0xdf, 0x76, 0x5f, 0x4a, 0x17, 0x10, 0x6c, 0x43, 0x4f, 0x51, 0xaa, 0x55, 0xe2,
	0x7d, 0x18, 0x78, 0xc4, 0x8f, 0x3c, 0xf2, 0xf3, 0x8c, 0xf0, 0xc5, 0x26, 0x94, 0x0a, 0xee, 0x54,
	0x0a, 0xc6, 0xef, 0x60, 0x53, 0x91, 0x79, 0xc2, 0xe6, 0x9c, 0xa0, 0x7d, 0xe8, 0xeb, 0x0c, 0xa5,
	0xc4, 0x60, 0x7a, 0xbf, 0xbe, 0x04, 0xdd, 0x76, 0xcf, 0x30, 0xf0, 0x11, 0x6c, 0xbd, 0x49, 0x89,
	0x9f, 0x11, 0x93, 0xcb, 0xb5, 0xd4, 0x66, 0x30, 0x34, 0x6a, 0x2b, 0x4a, 0xee, 0x73, 0x12, 0xad,
	0x2a, 0xb9, 0x6d, 0x18, 0x1a, 0x35, 0x95, 0x1c, 0xde, 0x82, 0xc1, 0x11, 0xe5, 0x99, 0x56, 0xc7,
	0x33, 0xd8, 0x54, 0x57, 0x9d, 0xfb, 0x01, 0x6c, 0x68, 0x2e, 0xb7, 0xdb, 0x72, 0x38, 0x96, 0x08,
	0x97, 0x53, 0xf0, 0x7b, 0xd8, 0x7e, 0x1d, 0x45, 0x7a, 0x68, 0x74, 0x01, 0x7b, 0x00, 0xda, 0x7e,
	0x9c, 0xbf, 0xb8, 0xa5, 0x91, 0xc3, 0x08, 0xdd, 0x01, 0x4b, 0x0d, 0x97, 0xb0, 0xaa, 0x87, 0xdf,
	0x50, 0xc0, 0x61, 0x84, 0x77, 0xe0, 0x66, 0x41, 0x4f, 0x97, 0xf0, 0x11, 0x76, 0x3c, 0x12, 0xb3,
	0x73, 0xb2, 0xba, 0x38, 0xbb, 0x70, 0xab, 0x2c, 0xa9, 0x43, 0xbd, 0x84, 0x5d, 0xd1, 0x1e, 0x85,
	0xf2, 0x53, 0x9a, 0x70, 0x13, 0xad, 0x24, 0xd7, 0xae, 0xc8, 0x7d, 0x85, 0xdb, 0x0b, 0xb4, 0x95,
	0x34, 0x78, 0xfa, 0x67, 0x1d, 0x86, 0x1a, 0xfd, 0xa4, 0x9c, 0xd1, 0x17, 0xe8, 0x8a, 0xdd, 0x40,
	0x0f, 0xea, 0x75, 0x0a, 0x8b, 0xe7, 0x3c, 0xbc, 0xca, 0x4d, 0x97, 0xde, 0x42, 0xdf, 0xa0, 0xa7,
	0x26, 0x1b, 0x3d, 0xaa, 0xe7, 0x94, 0x36, 0xc9, 0x19, 0x5f, 0xed, 0x58, 0x94, 0x57, 0xb3, 0xd9,
	0x24, 0x5f, 0xda, 0x85, 0x26, 0xf9, 0xca, 0x98, 0xb7, 0x44, 0x5b, 0xc4, 0x1b, 0x34, 0xb5, 0xa5,
	0xb0, 0x08, 0x4d, 0x6d, 0x29, 0x2e, 0x08, 0x6e, 0xa1, 0xef, 0x60, 0xe5, 0x33, 0x89, 0x1e, 0xd7,
	0xd3, 0xaa, 0x8b, 0xe0, 0x3c, 0x59, 0xca, 0x37, 0x8f, 0xc3, 0xc4, 0x7f, 0xde, 0xe5, 0x4c, 0xa2,
	0xa7, 0x4d, 0x0f, 0xb7, 0xb0, 0x0e, 0x8e, 0xbb, 0xac, 0x7b, 0x1e, 0xf0, 0x17, 0xdc, 0xa8, 0x4c,
	0x2d, 0x7a, 0xd6, 0xdc, 0x95, 0xc5, 0xbd, 0x70, 0x9e, 0xff, 0x07, 0xc3, 0x44, 0x0e, 0x7a, 0xf2,
	0xf3, 0xf9, 0xe2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x37, 0x8c, 0xc1, 0x7f, 0x07, 0x00,
	0x00,
}