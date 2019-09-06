// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator.proto

package validator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	v1 "google.golang.org/genproto/googleapis/cloud/asset/v1"
	v11 "google.golang.org/genproto/googleapis/iam/v1"
	grpc "google.golang.org/grpc"
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

// Asset contains GCP resource metadata and additional metadata set on a resource, such as Cloud IAM policy.
// WARNING: these field names are directly used to structure data passed to templates.
// Changes in field names will result in changes to the data provided to the templates.
type Asset struct {
	// GCP resource name as defined by Cloud Asset Inventory.
	// See https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/resource-name-format for the format.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cloud Asset Inventory type (CAI API v1 format). Example: "sqladmin.googleapis.com/Instance" is the type of Cloud SQL instance.
	// This field has a redundant "asset" prefix to be consistent with Cloud Asset Inventory output.
	// See https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/overview#supported_resource_types for the list of types.
	AssetType string `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// Ancestral project/folder/org information in a path-like format.
	// For example, a GCP project that is nested under a folder may have the following path:
	// organization/9999/folder/8888/project/7777
	AncestryPath string `protobuf:"bytes,3,opt,name=ancestry_path,json=ancestryPath,proto3" json:"ancestry_path,omitempty"`
	// GCP resource metadata.
	Resource *v1.Resource `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	// IAM policy associated with the resource.
	IamPolicy            *v11.Policy `protobuf:"bytes,5,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Asset) GetAncestryPath() string {
	if m != nil {
		return m.AncestryPath
	}
	return ""
}

func (m *Asset) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Asset) GetIamPolicy() *v11.Policy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

// Constraint contains the configuration for a constraint.
type Constraint struct {
	// Metadata contains the user-provided constraint metadata
	Metadata             *_struct.Value `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Constraint) Reset()         { *m = Constraint{} }
func (m *Constraint) String() string { return proto.CompactTextString(m) }
func (*Constraint) ProtoMessage()    {}
func (*Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{1}
}

func (m *Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constraint.Unmarshal(m, b)
}
func (m *Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constraint.Marshal(b, m, deterministic)
}
func (m *Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constraint.Merge(m, src)
}
func (m *Constraint) XXX_Size() int {
	return xxx_messageInfo_Constraint.Size(m)
}
func (m *Constraint) XXX_DiscardUnknown() {
	xxx_messageInfo_Constraint.DiscardUnknown(m)
}

var xxx_messageInfo_Constraint proto.InternalMessageInfo

func (m *Constraint) GetMetadata() *_struct.Value {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Violation contains the relevant information to explain how a constraint is violated.
type Violation struct {
	// The name of the constraint that's violated.
	Constraint string `protobuf:"bytes,1,opt,name=constraint,proto3" json:"constraint,omitempty"`
	// GCP resource name. This is the same name in Asset.
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Human readable error message.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Metadata is optional. It contains the constraint-specific information that can potentially be used for remediation.
	// Example: In a firewall rule constraint violation, Metadata can contain the open port number.
	Metadata *_struct.Value `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The full constraint configuration.
	ConstraintConfig     *Constraint `protobuf:"bytes,5,opt,name=constraint_config,json=constraintConfig,proto3" json:"constraint_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Violation) Reset()         { *m = Violation{} }
func (m *Violation) String() string { return proto.CompactTextString(m) }
func (*Violation) ProtoMessage()    {}
func (*Violation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{2}
}

func (m *Violation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Violation.Unmarshal(m, b)
}
func (m *Violation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Violation.Marshal(b, m, deterministic)
}
func (m *Violation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Violation.Merge(m, src)
}
func (m *Violation) XXX_Size() int {
	return xxx_messageInfo_Violation.Size(m)
}
func (m *Violation) XXX_DiscardUnknown() {
	xxx_messageInfo_Violation.DiscardUnknown(m)
}

var xxx_messageInfo_Violation proto.InternalMessageInfo

func (m *Violation) GetConstraint() string {
	if m != nil {
		return m.Constraint
	}
	return ""
}

func (m *Violation) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Violation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Violation) GetMetadata() *_struct.Value {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Violation) GetConstraintConfig() *Constraint {
	if m != nil {
		return m.ConstraintConfig
	}
	return nil
}

type AddDataRequest struct {
	Assets               []*Asset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDataRequest) Reset()         { *m = AddDataRequest{} }
func (m *AddDataRequest) String() string { return proto.CompactTextString(m) }
func (*AddDataRequest) ProtoMessage()    {}
func (*AddDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{3}
}

func (m *AddDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDataRequest.Unmarshal(m, b)
}
func (m *AddDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDataRequest.Marshal(b, m, deterministic)
}
func (m *AddDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDataRequest.Merge(m, src)
}
func (m *AddDataRequest) XXX_Size() int {
	return xxx_messageInfo_AddDataRequest.Size(m)
}
func (m *AddDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDataRequest proto.InternalMessageInfo

func (m *AddDataRequest) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

type AddDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDataResponse) Reset()         { *m = AddDataResponse{} }
func (m *AddDataResponse) String() string { return proto.CompactTextString(m) }
func (*AddDataResponse) ProtoMessage()    {}
func (*AddDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{4}
}

func (m *AddDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDataResponse.Unmarshal(m, b)
}
func (m *AddDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDataResponse.Marshal(b, m, deterministic)
}
func (m *AddDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDataResponse.Merge(m, src)
}
func (m *AddDataResponse) XXX_Size() int {
	return xxx_messageInfo_AddDataResponse.Size(m)
}
func (m *AddDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddDataResponse proto.InternalMessageInfo

type AuditRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditRequest) Reset()         { *m = AuditRequest{} }
func (m *AuditRequest) String() string { return proto.CompactTextString(m) }
func (*AuditRequest) ProtoMessage()    {}
func (*AuditRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{5}
}

func (m *AuditRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditRequest.Unmarshal(m, b)
}
func (m *AuditRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditRequest.Marshal(b, m, deterministic)
}
func (m *AuditRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditRequest.Merge(m, src)
}
func (m *AuditRequest) XXX_Size() int {
	return xxx_messageInfo_AuditRequest.Size(m)
}
func (m *AuditRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuditRequest proto.InternalMessageInfo

type AuditResponse struct {
	Violations           []*Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AuditResponse) Reset()         { *m = AuditResponse{} }
func (m *AuditResponse) String() string { return proto.CompactTextString(m) }
func (*AuditResponse) ProtoMessage()    {}
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{6}
}

func (m *AuditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditResponse.Unmarshal(m, b)
}
func (m *AuditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditResponse.Marshal(b, m, deterministic)
}
func (m *AuditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditResponse.Merge(m, src)
}
func (m *AuditResponse) XXX_Size() int {
	return xxx_messageInfo_AuditResponse.Size(m)
}
func (m *AuditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditResponse proto.InternalMessageInfo

func (m *AuditResponse) GetViolations() []*Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

type ResetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetRequest) Reset()         { *m = ResetRequest{} }
func (m *ResetRequest) String() string { return proto.CompactTextString(m) }
func (*ResetRequest) ProtoMessage()    {}
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{7}
}

func (m *ResetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetRequest.Unmarshal(m, b)
}
func (m *ResetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetRequest.Marshal(b, m, deterministic)
}
func (m *ResetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetRequest.Merge(m, src)
}
func (m *ResetRequest) XXX_Size() int {
	return xxx_messageInfo_ResetRequest.Size(m)
}
func (m *ResetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetRequest proto.InternalMessageInfo

type ResetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetResponse) Reset()         { *m = ResetResponse{} }
func (m *ResetResponse) String() string { return proto.CompactTextString(m) }
func (*ResetResponse) ProtoMessage()    {}
func (*ResetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{8}
}

func (m *ResetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetResponse.Unmarshal(m, b)
}
func (m *ResetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetResponse.Marshal(b, m, deterministic)
}
func (m *ResetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetResponse.Merge(m, src)
}
func (m *ResetResponse) XXX_Size() int {
	return xxx_messageInfo_ResetResponse.Size(m)
}
func (m *ResetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Asset)(nil), "validator.Asset")
	proto.RegisterType((*Constraint)(nil), "validator.Constraint")
	proto.RegisterType((*Violation)(nil), "validator.Violation")
	proto.RegisterType((*AddDataRequest)(nil), "validator.AddDataRequest")
	proto.RegisterType((*AddDataResponse)(nil), "validator.AddDataResponse")
	proto.RegisterType((*AuditRequest)(nil), "validator.AuditRequest")
	proto.RegisterType((*AuditResponse)(nil), "validator.AuditResponse")
	proto.RegisterType((*ResetRequest)(nil), "validator.ResetRequest")
	proto.RegisterType((*ResetResponse)(nil), "validator.ResetResponse")
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_bf1c6ec7c0d80dd5) }

var fileDescriptor_bf1c6ec7c0d80dd5 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x6d, 0xe8, 0x6e, 0xdb, 0x4c, 0xbb, 0xdd, 0xd6, 0xe2, 0x23, 0x44, 0x7c, 0xac, 0xc2, 0x65,
	0x4f, 0x89, 0xba, 0x70, 0x02, 0x0e, 0x6c, 0x0b, 0xf7, 0xca, 0x42, 0xbd, 0xae, 0xa6, 0x89, 0xbb,
	0xb5, 0x94, 0xc4, 0x21, 0x76, 0x56, 0xda, 0xdf, 0xc8, 0x8d, 0x1b, 0xff, 0x06, 0xc5, 0x76, 0x12,
	0x47, 0x70, 0xe0, 0x66, 0xcf, 0xbc, 0xf7, 0xe6, 0xbd, 0xb1, 0x61, 0xbe, 0xc3, 0x9c, 0x67, 0xa8,
	0x44, 0x1d, 0x57, 0xb5, 0x50, 0x82, 0xf8, 0x7d, 0x21, 0x0c, 0xb7, 0x42, 0x6c, 0x73, 0x96, 0x70,
	0x2c, 0x92, 0xdd, 0x55, 0x52, 0x89, 0x9c, 0xa7, 0x7b, 0x03, 0x0b, 0x5f, 0xd9, 0x9e, 0xbe, 0xdd,
	0x37, 0x0f, 0x89, 0x54, 0x75, 0x93, 0x2a, 0xdb, 0x8d, 0x6c, 0x37, 0xcd, 0x45, 0x93, 0x25, 0x28,
	0x25, 0x53, 0xad, 0x82, 0x3e, 0x48, 0x83, 0x89, 0x7e, 0x79, 0x30, 0x5d, 0xb7, 0x05, 0x42, 0x60,
	0x52, 0x62, 0xc1, 0x02, 0x6f, 0xe1, 0x2d, 0x7d, 0xaa, 0xcf, 0xe4, 0x35, 0x80, 0x46, 0x6f, 0xd4,
	0xbe, 0x62, 0xc1, 0x13, 0xdd, 0xf1, 0x75, 0xe5, 0xfb, 0xbe, 0x62, 0xe4, 0x1d, 0xcc, 0xb0, 0x4c,
	0x99, 0x54, 0xf5, 0x7e, 0x53, 0xa1, 0x7a, 0x0c, 0x0e, 0x35, 0xe2, 0xac, 0x2b, 0xde, 0xa2, 0x7a,
	0x24, 0x9f, 0xe0, 0xa4, 0x66, 0x52, 0x34, 0x75, 0xca, 0x82, 0xc9, 0xc2, 0x5b, 0x9e, 0xae, 0xde,
	0xc6, 0xc6, 0x58, 0xac, 0x8d, 0xc5, 0x5a, 0x2f, 0xde, 0x5d, 0xc5, 0xd4, 0xc2, 0x68, 0x4f, 0x20,
	0x1f, 0x00, 0x38, 0x16, 0x1b, 0x13, 0x3a, 0x98, 0x6a, 0xfa, 0xb3, 0x8e, 0xce, 0xb1, 0x68, 0x69,
	0xb7, 0xba, 0x49, 0x7d, 0x8e, 0x85, 0x39, 0x46, 0x5f, 0x00, 0x6e, 0x44, 0x29, 0x55, 0x8d, 0xbc,
	0x54, 0x64, 0x05, 0x27, 0x05, 0x53, 0x98, 0xa1, 0x42, 0xab, 0xf0, 0xbc, 0x53, 0xe8, 0xf6, 0x16,
	0xdf, 0x61, 0xde, 0x30, 0xda, 0xe3, 0xa2, 0xdf, 0x1e, 0xf8, 0x77, 0x5c, 0xe4, 0xa8, 0xb8, 0x28,
	0xc9, 0x1b, 0x80, 0xb4, 0xd7, 0xb3, 0x0b, 0x72, 0x2a, 0x24, 0x74, 0x22, 0x9a, 0x25, 0x0d, 0x09,
	0x02, 0x38, 0x2e, 0x98, 0x94, 0xb8, 0x65, 0x76, 0x3b, 0xdd, 0x75, 0xe4, 0x6b, 0xf2, 0x7f, 0xbe,
	0xc8, 0x35, 0x5c, 0x0e, 0x73, 0x37, 0xa9, 0x28, 0x1f, 0xf8, 0xb6, 0x5f, 0xcb, 0xf0, 0x89, 0x86,
	0xf4, 0xf4, 0x62, 0xc0, 0xdf, 0x68, 0x78, 0xf4, 0x11, 0xce, 0xd7, 0x59, 0xf6, 0x15, 0x15, 0x52,
	0xf6, 0xa3, 0x61, 0x52, 0x91, 0x25, 0x1c, 0x99, 0x4f, 0x11, 0x78, 0x8b, 0xc3, 0xe5, 0xe9, 0xea,
	0xc2, 0x91, 0xd2, 0x9f, 0x83, 0xda, 0x7e, 0x74, 0x09, 0xf3, 0x9e, 0x2b, 0x2b, 0x51, 0x4a, 0x16,
	0x9d, 0xc3, 0xd9, 0xba, 0xc9, 0xb8, 0xb2, 0x62, 0xd1, 0x37, 0x98, 0xd9, 0xbb, 0x01, 0xb4, 0x6f,
	0xb8, 0xeb, 0x56, 0xd9, 0x4d, 0x78, 0xea, 0x4c, 0xe8, 0xf7, 0x4c, 0x1d, 0x5c, 0x2b, 0x4b, 0x59,
	0x3b, 0xda, 0xca, 0xce, 0x61, 0x66, 0xef, 0x46, 0x76, 0xf5, 0xb3, 0x7d, 0xa2, 0x4e, 0x84, 0x5c,
	0xc3, 0xb1, 0x35, 0x46, 0x5e, 0xba, 0xee, 0x47, 0x41, 0xc3, 0xf0, 0x5f, 0x2d, 0x9b, 0xe3, 0x80,
	0x7c, 0x86, 0xa9, 0x76, 0x4e, 0x5e, 0xb8, 0x30, 0x27, 0x5b, 0x18, 0xfc, 0xdd, 0x70, 0xd9, 0xda,
	0xe0, 0x88, 0xed, 0x46, 0x18, 0xb1, 0x47, 0x59, 0xa2, 0x83, 0xfb, 0x23, 0xfd, 0xe4, 0xef, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x82, 0x69, 0x59, 0x0a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValidatorClient is the client API for Validator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatorClient interface {
	// AddData adds GCP resource metadata to be audited later.
	AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error)
	// Audit checks the GCP resource metadata that has been added via AddData to determine if any of the constraint is violated.
	Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error)
	// Reset clears previously added data from the underlying query evaluation engine.
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
}

type validatorClient struct {
	cc *grpc.ClientConn
}

func NewValidatorClient(cc *grpc.ClientConn) ValidatorClient {
	return &validatorClient{cc}
}

func (c *validatorClient) AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error) {
	out := new(AddDataResponse)
	err := c.cc.Invoke(ctx, "/validator.Validator/AddData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorClient) Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error) {
	out := new(AuditResponse)
	err := c.cc.Invoke(ctx, "/validator.Validator/Audit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/validator.Validator/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServer is the server API for Validator service.
type ValidatorServer interface {
	// AddData adds GCP resource metadata to be audited later.
	AddData(context.Context, *AddDataRequest) (*AddDataResponse, error)
	// Audit checks the GCP resource metadata that has been added via AddData to determine if any of the constraint is violated.
	Audit(context.Context, *AuditRequest) (*AuditResponse, error)
	// Reset clears previously added data from the underlying query evaluation engine.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
}

func RegisterValidatorServer(s *grpc.Server, srv ValidatorServer) {
	s.RegisterService(&_Validator_serviceDesc, srv)
}

func _Validator_AddData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServer).AddData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validator.Validator/AddData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServer).AddData(ctx, req.(*AddDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Validator_Audit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServer).Audit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validator.Validator/Audit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServer).Audit(ctx, req.(*AuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Validator_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validator.Validator/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Validator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validator.Validator",
	HandlerType: (*ValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddData",
			Handler:    _Validator_AddData_Handler,
		},
		{
			MethodName: "Audit",
			Handler:    _Validator_Audit_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Validator_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validator.proto",
}