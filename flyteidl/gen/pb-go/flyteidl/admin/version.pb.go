// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/version.proto

package admin

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

// Response for the GetVersion API
type GetVersionResponse struct {
	// The control plane version information. FlyteAdmin and related components
	// form the control plane of Flyte
	ControlPlaneVersion  *Version `protobuf:"bytes,1,opt,name=control_plane_version,json=controlPlaneVersion,proto3" json:"control_plane_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionResponse) Reset()         { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()    {}
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a025621cd13402e3, []int{0}
}

func (m *GetVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionResponse.Unmarshal(m, b)
}
func (m *GetVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionResponse.Merge(m, src)
}
func (m *GetVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetVersionResponse.Size(m)
}
func (m *GetVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionResponse proto.InternalMessageInfo

func (m *GetVersionResponse) GetControlPlaneVersion() *Version {
	if m != nil {
		return m.ControlPlaneVersion
	}
	return nil
}

// Provides Version information for a component
type Version struct {
	// Specifies the GIT sha of the build
	Build string `protobuf:"bytes,1,opt,name=Build,proto3" json:"Build,omitempty"`
	// Version for the build, should follow a semver
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	// Build timestamp
	BuildTime            string   `protobuf:"bytes,3,opt,name=BuildTime,proto3" json:"BuildTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_a025621cd13402e3, []int{1}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

// Empty request for GetVersion
type GetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a025621cd13402e3, []int{2}
}

func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetVersionResponse)(nil), "flyteidl.admin.GetVersionResponse")
	proto.RegisterType((*Version)(nil), "flyteidl.admin.Version")
	proto.RegisterType((*GetVersionRequest)(nil), "flyteidl.admin.GetVersionRequest")
}

func init() { proto.RegisterFile("flyteidl/admin/version.proto", fileDescriptor_a025621cd13402e3) }

var fileDescriptor_a025621cd13402e3 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x4b, 0x2d, 0x2a, 0xce,
	0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xc9, 0xea, 0x81, 0x65, 0x95,
	0x12, 0xb9, 0x84, 0xdc, 0x53, 0x4b, 0xc2, 0x20, 0x6a, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0xbc, 0xb9, 0x44, 0x93, 0xf3, 0xf3, 0x4a, 0x8a, 0xf2, 0x73, 0xe2, 0x0b, 0x72, 0x12,
	0xf3, 0x52, 0xe3, 0xa1, 0x86, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0xa1, 0x9a,
	0xa2, 0x07, 0xd3, 0x2f, 0x0c, 0xd5, 0x15, 0x00, 0xd2, 0x04, 0x15, 0x54, 0x0a, 0xe7, 0x62, 0x87,
	0x32, 0x85, 0x44, 0xb8, 0x58, 0x9d, 0x4a, 0x33, 0x73, 0x52, 0xc0, 0xe6, 0x70, 0x06, 0x41, 0x38,
	0x42, 0x12, 0x70, 0x05, 0x12, 0x4c, 0x60, 0x71, 0xb8, 0x7a, 0x19, 0x2e, 0x4e, 0xb0, 0x92, 0x90,
	0xcc, 0xdc, 0x54, 0x09, 0x66, 0xb0, 0x1c, 0x42, 0x40, 0x49, 0x98, 0x4b, 0x10, 0xd9, 0xed, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x4e, 0xe6, 0x51, 0xa6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x60, 0x77, 0xe6, 0x17, 0xa5, 0xeb, 0xc3, 0x03, 0x25, 0x3d, 0x35, 0x4f, 0xbf,
	0x20, 0x49, 0x37, 0x3d, 0x5f, 0x1f, 0x35, 0x9c, 0x92, 0xd8, 0xc0, 0x01, 0x64, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x77, 0x9b, 0x1b, 0x7c, 0x40, 0x01, 0x00, 0x00,
}
