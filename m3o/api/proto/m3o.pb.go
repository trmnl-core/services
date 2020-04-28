// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/m3o/api/proto/m3o.proto

package go_micro_api_m3o

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ProfilePictureUrl    string   `protobuf:"bytes,5,opt,name=profile_picture_url,json=profilePictureUrl,proto3" json:"profile_picture_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{0}
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

func (m *User) GetProfilePictureUrl() string {
	if m != nil {
		return m.ProfilePictureUrl
	}
	return ""
}

type ReadAccountRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAccountRequest) Reset()         { *m = ReadAccountRequest{} }
func (m *ReadAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAccountRequest) ProtoMessage()    {}
func (*ReadAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{1}
}

func (m *ReadAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountRequest.Unmarshal(m, b)
}
func (m *ReadAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountRequest.Marshal(b, m, deterministic)
}
func (m *ReadAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountRequest.Merge(m, src)
}
func (m *ReadAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAccountRequest.Size(m)
}
func (m *ReadAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountRequest proto.InternalMessageInfo

type ReadAccountResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAccountResponse) Reset()         { *m = ReadAccountResponse{} }
func (m *ReadAccountResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAccountResponse) ProtoMessage()    {}
func (*ReadAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{2}
}

func (m *ReadAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountResponse.Unmarshal(m, b)
}
func (m *ReadAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountResponse.Marshal(b, m, deterministic)
}
func (m *ReadAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountResponse.Merge(m, src)
}
func (m *ReadAccountResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAccountResponse.Size(m)
}
func (m *ReadAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountResponse proto.InternalMessageInfo

func (m *ReadAccountResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Project struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ApiDomain            string   `protobuf:"bytes,4,opt,name=api_domain,json=apiDomain,proto3" json:"api_domain,omitempty"`
	WebDomain            string   `protobuf:"bytes,5,opt,name=web_domain,json=webDomain,proto3" json:"web_domain,omitempty"`
	Repository           string   `protobuf:"bytes,6,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{3}
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

func (m *Project) GetApiDomain() string {
	if m != nil {
		return m.ApiDomain
	}
	return ""
}

func (m *Project) GetWebDomain() string {
	if m != nil {
		return m.WebDomain
	}
	return ""
}

func (m *Project) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

type CreateProjectRequest struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{4}
}

func (m *CreateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectRequest.Unmarshal(m, b)
}
func (m *CreateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectRequest.Merge(m, src)
}
func (m *CreateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProjectRequest.Size(m)
}
func (m *CreateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectRequest proto.InternalMessageInfo

func (m *CreateProjectRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type CreateProjectResponse struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectResponse) Reset()         { *m = CreateProjectResponse{} }
func (m *CreateProjectResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProjectResponse) ProtoMessage()    {}
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{5}
}

func (m *CreateProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectResponse.Unmarshal(m, b)
}
func (m *CreateProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectResponse.Marshal(b, m, deterministic)
}
func (m *CreateProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectResponse.Merge(m, src)
}
func (m *CreateProjectResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProjectResponse.Size(m)
}
func (m *CreateProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectResponse proto.InternalMessageInfo

func (m *CreateProjectResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateProjectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ApiDomain            string   `protobuf:"bytes,3,opt,name=api_domain,json=apiDomain,proto3" json:"api_domain,omitempty"`
	WebDomain            string   `protobuf:"bytes,4,opt,name=web_domain,json=webDomain,proto3" json:"web_domain,omitempty"`
	Repository           string   `protobuf:"bytes,5,opt,name=repository,proto3" json:"repository,omitempty"`
	GithubToken          string   `protobuf:"bytes,6,opt,name=github_token,json=githubToken,proto3" json:"github_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProjectRequest) Reset()         { *m = UpdateProjectRequest{} }
func (m *UpdateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectRequest) ProtoMessage()    {}
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{6}
}

func (m *UpdateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectRequest.Unmarshal(m, b)
}
func (m *UpdateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectRequest.Merge(m, src)
}
func (m *UpdateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectRequest.Size(m)
}
func (m *UpdateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectRequest proto.InternalMessageInfo

func (m *UpdateProjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProjectRequest) GetApiDomain() string {
	if m != nil {
		return m.ApiDomain
	}
	return ""
}

func (m *UpdateProjectRequest) GetWebDomain() string {
	if m != nil {
		return m.WebDomain
	}
	return ""
}

func (m *UpdateProjectRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *UpdateProjectRequest) GetGithubToken() string {
	if m != nil {
		return m.GithubToken
	}
	return ""
}

type UpdateProjectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProjectResponse) Reset()         { *m = UpdateProjectResponse{} }
func (m *UpdateProjectResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectResponse) ProtoMessage()    {}
func (*UpdateProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{7}
}

func (m *UpdateProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectResponse.Unmarshal(m, b)
}
func (m *UpdateProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectResponse.Marshal(b, m, deterministic)
}
func (m *UpdateProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectResponse.Merge(m, src)
}
func (m *UpdateProjectResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectResponse.Size(m)
}
func (m *UpdateProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectResponse proto.InternalMessageInfo

type ListProjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectsRequest) Reset()         { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()    {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{8}
}

func (m *ListProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsRequest.Unmarshal(m, b)
}
func (m *ListProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsRequest.Merge(m, src)
}
func (m *ListProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProjectsRequest.Size(m)
}
func (m *ListProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsRequest proto.InternalMessageInfo

type ListProjectsResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProjectsResponse) Reset()         { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()    {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{9}
}

func (m *ListProjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResponse.Unmarshal(m, b)
}
func (m *ListProjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListProjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResponse.Merge(m, src)
}
func (m *ListProjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResponse.Size(m)
}
func (m *ListProjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResponse proto.InternalMessageInfo

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type VerifyGithubTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyGithubTokenRequest) Reset()         { *m = VerifyGithubTokenRequest{} }
func (m *VerifyGithubTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyGithubTokenRequest) ProtoMessage()    {}
func (*VerifyGithubTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{10}
}

func (m *VerifyGithubTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyGithubTokenRequest.Unmarshal(m, b)
}
func (m *VerifyGithubTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyGithubTokenRequest.Marshal(b, m, deterministic)
}
func (m *VerifyGithubTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyGithubTokenRequest.Merge(m, src)
}
func (m *VerifyGithubTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyGithubTokenRequest.Size(m)
}
func (m *VerifyGithubTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyGithubTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyGithubTokenRequest proto.InternalMessageInfo

func (m *VerifyGithubTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyGithubTokenResponse struct {
	Repos                []string `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyGithubTokenResponse) Reset()         { *m = VerifyGithubTokenResponse{} }
func (m *VerifyGithubTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyGithubTokenResponse) ProtoMessage()    {}
func (*VerifyGithubTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{11}
}

func (m *VerifyGithubTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyGithubTokenResponse.Unmarshal(m, b)
}
func (m *VerifyGithubTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyGithubTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyGithubTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyGithubTokenResponse.Merge(m, src)
}
func (m *VerifyGithubTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyGithubTokenResponse.Size(m)
}
func (m *VerifyGithubTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyGithubTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyGithubTokenResponse proto.InternalMessageInfo

func (m *VerifyGithubTokenResponse) GetRepos() []string {
	if m != nil {
		return m.Repos
	}
	return nil
}

type WebhookAPIKeyRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebhookAPIKeyRequest) Reset()         { *m = WebhookAPIKeyRequest{} }
func (m *WebhookAPIKeyRequest) String() string { return proto.CompactTextString(m) }
func (*WebhookAPIKeyRequest) ProtoMessage()    {}
func (*WebhookAPIKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{12}
}

func (m *WebhookAPIKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookAPIKeyRequest.Unmarshal(m, b)
}
func (m *WebhookAPIKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookAPIKeyRequest.Marshal(b, m, deterministic)
}
func (m *WebhookAPIKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookAPIKeyRequest.Merge(m, src)
}
func (m *WebhookAPIKeyRequest) XXX_Size() int {
	return xxx_messageInfo_WebhookAPIKeyRequest.Size(m)
}
func (m *WebhookAPIKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookAPIKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookAPIKeyRequest proto.InternalMessageInfo

func (m *WebhookAPIKeyRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type WebhookAPIKeyResponse struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebhookAPIKeyResponse) Reset()         { *m = WebhookAPIKeyResponse{} }
func (m *WebhookAPIKeyResponse) String() string { return proto.CompactTextString(m) }
func (*WebhookAPIKeyResponse) ProtoMessage()    {}
func (*WebhookAPIKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee7567aa95c263, []int{13}
}

func (m *WebhookAPIKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookAPIKeyResponse.Unmarshal(m, b)
}
func (m *WebhookAPIKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookAPIKeyResponse.Marshal(b, m, deterministic)
}
func (m *WebhookAPIKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookAPIKeyResponse.Merge(m, src)
}
func (m *WebhookAPIKeyResponse) XXX_Size() int {
	return xxx_messageInfo_WebhookAPIKeyResponse.Size(m)
}
func (m *WebhookAPIKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookAPIKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookAPIKeyResponse proto.InternalMessageInfo

func (m *WebhookAPIKeyResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *WebhookAPIKeyResponse) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.api.m3o.User")
	proto.RegisterType((*ReadAccountRequest)(nil), "go.micro.api.m3o.ReadAccountRequest")
	proto.RegisterType((*ReadAccountResponse)(nil), "go.micro.api.m3o.ReadAccountResponse")
	proto.RegisterType((*Project)(nil), "go.micro.api.m3o.Project")
	proto.RegisterType((*CreateProjectRequest)(nil), "go.micro.api.m3o.CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "go.micro.api.m3o.CreateProjectResponse")
	proto.RegisterType((*UpdateProjectRequest)(nil), "go.micro.api.m3o.UpdateProjectRequest")
	proto.RegisterType((*UpdateProjectResponse)(nil), "go.micro.api.m3o.UpdateProjectResponse")
	proto.RegisterType((*ListProjectsRequest)(nil), "go.micro.api.m3o.ListProjectsRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "go.micro.api.m3o.ListProjectsResponse")
	proto.RegisterType((*VerifyGithubTokenRequest)(nil), "go.micro.api.m3o.VerifyGithubTokenRequest")
	proto.RegisterType((*VerifyGithubTokenResponse)(nil), "go.micro.api.m3o.VerifyGithubTokenResponse")
	proto.RegisterType((*WebhookAPIKeyRequest)(nil), "go.micro.api.m3o.WebhookAPIKeyRequest")
	proto.RegisterType((*WebhookAPIKeyResponse)(nil), "go.micro.api.m3o.WebhookAPIKeyResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/m3o/api/proto/m3o.proto", fileDescriptor_adee7567aa95c263)
}

var fileDescriptor_adee7567aa95c263 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x4e, 0xd4, 0x40,
	0x14, 0x65, 0xd9, 0x2e, 0xd0, 0xcb, 0x47, 0x64, 0xe8, 0x6a, 0x59, 0xc4, 0xe0, 0xa8, 0x68, 0x30,
	0xe9, 0x0a, 0x1b, 0x1e, 0x80, 0x68, 0x62, 0x08, 0x68, 0xc8, 0xe2, 0x4a, 0x8c, 0x3f, 0x36, 0xdd,
	0xf6, 0x02, 0x23, 0xdb, 0x4e, 0x9d, 0x69, 0x25, 0xbc, 0x89, 0xaf, 0xe0, 0x33, 0xf8, 0x60, 0xfe,
	0x35, 0x9d, 0x99, 0x2e, 0xec, 0xb6, 0x02, 0xf1, 0x17, 0x9d, 0x73, 0xee, 0xdc, 0xb9, 0xe7, 0xcc,
	0x19, 0x16, 0xb6, 0xcf, 0x58, 0x7a, 0x9e, 0x0d, 0xbc, 0x80, 0x47, 0xed, 0x88, 0x05, 0x82, 0xb7,
	0x25, 0x8a, 0x1f, 0x2c, 0x40, 0xd9, 0x8e, 0x3a, 0xbc, 0xed, 0x27, 0xac, 0x9d, 0x08, 0x9e, 0xf2,
	0x7c, 0xe5, 0xa9, 0x2f, 0xf2, 0xe0, 0x8c, 0x7b, 0xaa, 0xd4, 0xf3, 0x13, 0xe6, 0x45, 0x1d, 0x4e,
	0x7f, 0xd6, 0xc0, 0xea, 0x49, 0x14, 0x64, 0x09, 0xa6, 0x59, 0xe8, 0xd6, 0x36, 0x6a, 0xaf, 0xec,
	0xee, 0x34, 0x0b, 0xc9, 0x3a, 0xc0, 0x29, 0x13, 0x32, 0xed, 0xc7, 0x7e, 0x84, 0xee, 0xb4, 0xc2,
	0x6d, 0x85, 0x7c, 0xf4, 0x23, 0x24, 0x6b, 0x60, 0x0f, 0xfd, 0x82, 0xad, 0x2b, 0x76, 0x2e, 0x07,
	0x14, 0xe9, 0x40, 0x03, 0x23, 0x9f, 0x0d, 0x5d, 0x4b, 0x11, 0x7a, 0x41, 0x3c, 0x58, 0x49, 0x04,
	0x3f, 0x65, 0x43, 0xec, 0x27, 0x2c, 0x48, 0x33, 0x81, 0xfd, 0x4c, 0x0c, 0xdd, 0x86, 0xaa, 0x59,
	0x36, 0xd4, 0x91, 0x66, 0x7a, 0x62, 0x48, 0x1d, 0x20, 0x5d, 0xf4, 0xc3, 0xbd, 0x20, 0xe0, 0x59,
	0x9c, 0x76, 0xf1, 0x7b, 0x86, 0x32, 0xa5, 0x7b, 0xb0, 0x32, 0x86, 0xca, 0x84, 0xc7, 0x12, 0xc9,
	0x16, 0x58, 0x99, 0x44, 0xa1, 0x04, 0xcc, 0xef, 0x3c, 0xf4, 0x26, 0x85, 0x7a, 0xb9, 0xc8, 0xae,
	0xaa, 0xa1, 0xbf, 0x6a, 0x30, 0x7b, 0x24, 0xf8, 0x37, 0x0c, 0xd2, 0x92, 0x6c, 0x02, 0xd6, 0x0d,
	0xc1, 0xea, 0x9b, 0x3c, 0x06, 0x3b, 0xff, 0x2b, 0x13, 0x3f, 0x28, 0xb4, 0x5e, 0x03, 0xb9, 0x51,
	0x7e, 0xc2, 0xfa, 0x21, 0x8f, 0x7c, 0x16, 0x1b, 0xc5, 0xb6, 0x9f, 0xb0, 0x77, 0x0a, 0xc8, 0xe9,
	0x4b, 0x1c, 0x14, 0xb4, 0x16, 0x6b, 0x5f, 0xe2, 0xc0, 0xd0, 0x4f, 0x00, 0x04, 0x26, 0x5c, 0xb2,
	0x94, 0x8b, 0x2b, 0x77, 0x46, 0xd1, 0x37, 0x10, 0x7a, 0x00, 0xce, 0x5b, 0x81, 0x7e, 0x8a, 0x66,
	0x60, 0x63, 0x03, 0xe9, 0xc0, 0x6c, 0xa2, 0x11, 0x23, 0x79, 0xb5, 0x2c, 0xb9, 0xd8, 0x52, 0x54,
	0xd2, 0x43, 0x68, 0x4e, 0x34, 0x33, 0xee, 0xfd, 0x57, 0xb7, 0xdf, 0x35, 0x70, 0x7a, 0x49, 0x58,
	0x9e, 0xed, 0x3e, 0x9e, 0x8e, 0xbb, 0x56, 0xbf, 0xdd, 0x35, 0xeb, 0x76, 0xd7, 0x1a, 0x93, 0xae,
	0x91, 0xa7, 0xb0, 0xa0, 0x1f, 0x47, 0x3f, 0xe5, 0x17, 0x18, 0x1b, 0x5f, 0xe7, 0x35, 0xf6, 0x29,
	0x87, 0xe8, 0x23, 0x68, 0x4e, 0x0c, 0xaf, 0xbd, 0xa0, 0x4d, 0x58, 0x39, 0x64, 0x32, 0x35, 0xb0,
	0x2c, 0x72, 0xf7, 0x01, 0x9c, 0x71, 0xd8, 0x58, 0xb7, 0x0b, 0x73, 0xc6, 0x10, 0xe9, 0xd6, 0x36,
	0xea, 0xb7, 0x7b, 0x37, 0x2a, 0xa5, 0x6f, 0xc0, 0xfd, 0x8c, 0x82, 0x9d, 0x5e, 0xbd, 0xbf, 0x9e,
	0xa9, 0xf0, 0xcf, 0x81, 0x86, 0x1e, 0x5b, 0x5b, 0xa8, 0x17, 0x74, 0x1b, 0x56, 0x2b, 0x76, 0x98,
	0x29, 0x1c, 0x68, 0x28, 0xf9, 0x6a, 0x04, 0xbb, 0xab, 0x17, 0x74, 0x17, 0x9c, 0x13, 0x1c, 0x9c,
	0x73, 0x7e, 0xb1, 0x77, 0xb4, 0x7f, 0x80, 0x57, 0xc5, 0x01, 0xeb, 0x00, 0x66, 0x90, 0xfe, 0xe8,
	0xa2, 0x6c, 0x83, 0xec, 0x87, 0xf4, 0x0b, 0x34, 0x27, 0xb6, 0x99, 0x53, 0xd6, 0xc0, 0x0e, 0x86,
	0x0c, 0xe3, 0x1b, 0xdb, 0xe6, 0x34, 0xb0, 0x1f, 0x92, 0x67, 0xb0, 0x68, 0x48, 0x89, 0x81, 0xc0,
	0xd4, 0x5c, 0xf7, 0x82, 0x06, 0x8f, 0x15, 0xb6, 0x73, 0x06, 0x4b, 0xe6, 0xe5, 0x1e, 0xeb, 0xff,
	0x55, 0xa4, 0x07, 0x56, 0xfe, 0x9e, 0xc9, 0xf3, 0xb2, 0x6b, 0xe5, 0xd7, 0xdf, 0x7a, 0x71, 0x47,
	0x95, 0xb9, 0xc3, 0xa9, 0x9d, 0x3f, 0x75, 0x58, 0x32, 0xae, 0x17, 0x27, 0x7d, 0x85, 0x19, 0x9d,
	0x7e, 0xb2, 0x59, 0xee, 0x52, 0xf5, 0xc8, 0x5a, 0x2f, 0xef, 0xac, 0x2b, 0xce, 0xcb, 0x9b, 0xeb,
	0x38, 0x55, 0x35, 0xaf, 0x7a, 0x25, 0x55, 0xcd, 0xab, 0x03, 0x39, 0x45, 0x4e, 0xc0, 0xca, 0xb3,
	0x47, 0x2a, 0xd4, 0x57, 0x44, 0xb5, 0xb5, 0x79, 0x57, 0xd9, 0xa8, 0x71, 0x0c, 0xcb, 0xa5, 0x4c,
	0x91, 0xad, 0xf2, 0xf6, 0x7f, 0x45, 0xb5, 0xf5, 0xfa, 0x5e, 0xb5, 0xa3, 0xf3, 0x06, 0xb0, 0x38,
	0x96, 0xac, 0x2a, 0xb3, 0xaa, 0x12, 0x5b, 0x65, 0x56, 0x65, 0x44, 0xe9, 0xd4, 0x60, 0x46, 0xfd,
	0xd4, 0x75, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x22, 0x73, 0x7e, 0xb3, 0x1f, 0x07, 0x00, 0x00,
}