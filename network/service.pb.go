// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

// Create
type CreateRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode                 uint32   `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
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

func (m *CreateRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

// Remove
type RemoveRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RemoveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Rename
type RenameRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Oldname              string   `protobuf:"bytes,2,opt,name=oldname,proto3" json:"oldname,omitempty"`
	Newname              string   `protobuf:"bytes,3,opt,name=newname,proto3" json:"newname,omitempty"`
	Newdirpath           string   `protobuf:"bytes,4,opt,name=newdirpath,proto3" json:"newdirpath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameRequest) Reset()         { *m = RenameRequest{} }
func (m *RenameRequest) String() string { return proto.CompactTextString(m) }
func (*RenameRequest) ProtoMessage()    {}
func (*RenameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RenameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameRequest.Unmarshal(m, b)
}
func (m *RenameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameRequest.Marshal(b, m, deterministic)
}
func (m *RenameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameRequest.Merge(m, src)
}
func (m *RenameRequest) XXX_Size() int {
	return xxx_messageInfo_RenameRequest.Size(m)
}
func (m *RenameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameRequest proto.InternalMessageInfo

func (m *RenameRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RenameRequest) GetOldname() string {
	if m != nil {
		return m.Oldname
	}
	return ""
}

func (m *RenameRequest) GetNewname() string {
	if m != nil {
		return m.Newname
	}
	return ""
}

func (m *RenameRequest) GetNewdirpath() string {
	if m != nil {
		return m.Newdirpath
	}
	return ""
}

// Mkdir
type MkdirRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode                 uint32   `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MkdirRequest) Reset()         { *m = MkdirRequest{} }
func (m *MkdirRequest) String() string { return proto.CompactTextString(m) }
func (*MkdirRequest) ProtoMessage()    {}
func (*MkdirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *MkdirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkdirRequest.Unmarshal(m, b)
}
func (m *MkdirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkdirRequest.Marshal(b, m, deterministic)
}
func (m *MkdirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkdirRequest.Merge(m, src)
}
func (m *MkdirRequest) XXX_Size() int {
	return xxx_messageInfo_MkdirRequest.Size(m)
}
func (m *MkdirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MkdirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MkdirRequest proto.InternalMessageInfo

func (m *MkdirRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MkdirRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MkdirRequest) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

// Write
type WriteRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WriteRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WriteRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Link
type LinkRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Newname              string   `protobuf:"bytes,2,opt,name=newname,proto3" json:"newname,omitempty"`
	Old                  string   `protobuf:"bytes,3,opt,name=old,proto3" json:"old,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *LinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkRequest.Unmarshal(m, b)
}
func (m *LinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkRequest.Marshal(b, m, deterministic)
}
func (m *LinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkRequest.Merge(m, src)
}
func (m *LinkRequest) XXX_Size() int {
	return xxx_messageInfo_LinkRequest.Size(m)
}
func (m *LinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinkRequest proto.InternalMessageInfo

func (m *LinkRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *LinkRequest) GetNewname() string {
	if m != nil {
		return m.Newname
	}
	return ""
}

func (m *LinkRequest) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

// Symlink
type SymlinkRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Newname              string   `protobuf:"bytes,3,opt,name=newname,proto3" json:"newname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymlinkRequest) Reset()         { *m = SymlinkRequest{} }
func (m *SymlinkRequest) String() string { return proto.CompactTextString(m) }
func (*SymlinkRequest) ProtoMessage()    {}
func (*SymlinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *SymlinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymlinkRequest.Unmarshal(m, b)
}
func (m *SymlinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymlinkRequest.Marshal(b, m, deterministic)
}
func (m *SymlinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymlinkRequest.Merge(m, src)
}
func (m *SymlinkRequest) XXX_Size() int {
	return xxx_messageInfo_SymlinkRequest.Size(m)
}
func (m *SymlinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymlinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymlinkRequest proto.InternalMessageInfo

func (m *SymlinkRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SymlinkRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *SymlinkRequest) GetNewname() string {
	if m != nil {
		return m.Newname
	}
	return ""
}

// Setattr
type SetattrRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode                 uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Atime                int64    `protobuf:"varint,3,opt,name=atime,proto3" json:"atime,omitempty"`
	Mtime                int64    `protobuf:"varint,4,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetattrRequest) Reset()         { *m = SetattrRequest{} }
func (m *SetattrRequest) String() string { return proto.CompactTextString(m) }
func (*SetattrRequest) ProtoMessage()    {}
func (*SetattrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *SetattrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetattrRequest.Unmarshal(m, b)
}
func (m *SetattrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetattrRequest.Marshal(b, m, deterministic)
}
func (m *SetattrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetattrRequest.Merge(m, src)
}
func (m *SetattrRequest) XXX_Size() int {
	return xxx_messageInfo_SetattrRequest.Size(m)
}
func (m *SetattrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetattrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetattrRequest proto.InternalMessageInfo

func (m *SetattrRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SetattrRequest) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *SetattrRequest) GetAtime() int64 {
	if m != nil {
		return m.Atime
	}
	return 0
}

func (m *SetattrRequest) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "network.Void")
	proto.RegisterType((*CreateRequest)(nil), "network.CreateRequest")
	proto.RegisterType((*RemoveRequest)(nil), "network.RemoveRequest")
	proto.RegisterType((*RenameRequest)(nil), "network.RenameRequest")
	proto.RegisterType((*MkdirRequest)(nil), "network.MkdirRequest")
	proto.RegisterType((*WriteRequest)(nil), "network.WriteRequest")
	proto.RegisterType((*LinkRequest)(nil), "network.LinkRequest")
	proto.RegisterType((*SymlinkRequest)(nil), "network.SymlinkRequest")
	proto.RegisterType((*SetattrRequest)(nil), "network.SetattrRequest")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xa5, 0x4d, 0x36, 0x15, 0xb3, 0xcd, 0x0a, 0x59, 0x4b, 0x89, 0x38, 0xa0, 0x95, 0x4f, 0x7b,
	0xa1, 0x68, 0xe9, 0x81, 0x0f, 0x40, 0xe2, 0x00, 0x94, 0x43, 0x2a, 0x95, 0xb3, 0xc1, 0xd3, 0x62,
	0x35, 0x89, 0x8b, 0x6b, 0x1a, 0xf1, 0x39, 0xfc, 0x29, 0xf2, 0xc4, 0x0d, 0x8e, 0xd2, 0x12, 0x21,
	0x71, 0x9b, 0x79, 0xd3, 0xf7, 0x46, 0x7e, 0xf3, 0x1a, 0x48, 0x0f, 0x68, 0x8e, 0xea, 0x2b, 0xce,
	0xf7, 0x46, 0x5b, 0xcd, 0x26, 0x15, 0xda, 0x5a, 0x9b, 0x1d, 0x4f, 0x20, 0x5e, 0x6b, 0x25, 0xf9,
	0x07, 0x48, 0xdf, 0x1a, 0x14, 0x16, 0x73, 0xfc, 0xfe, 0x03, 0x0f, 0x96, 0x31, 0x88, 0xf7, 0xc2,
	0x7e, 0xcb, 0x46, 0x77, 0xa3, 0xfb, 0xc7, 0x39, 0xd5, 0x0e, 0xab, 0x44, 0x89, 0xd9, 0xb8, 0xc1,
	0x5c, 0xed, 0xb0, 0x52, 0x4b, 0xcc, 0xa2, 0xbb, 0xd1, 0x7d, 0x9a, 0x53, 0xcd, 0xdf, 0x40, 0x9a,
	0x63, 0xa9, 0x8f, 0xff, 0x2a, 0xc6, 0x6b, 0x47, 0x74, 0xd5, 0xdf, 0x88, 0x19, 0x4c, 0x74, 0x21,
	0x03, 0xee, 0xa9, 0x75, 0x93, 0x0a, 0x6b, 0x9a, 0x44, 0xcd, 0xc4, 0xb7, 0xec, 0x05, 0x40, 0x85,
	0xb5, 0x54, 0x86, 0xd4, 0x62, 0x1a, 0x06, 0x08, 0x7f, 0x0f, 0xd3, 0xe5, 0x4e, 0x2a, 0xf3, 0x3f,
	0x5e, 0xff, 0x09, 0xa6, 0x9f, 0x8d, 0x1a, 0x74, 0x52, 0x0a, 0x2b, 0x48, 0x6b, 0x9a, 0x53, 0xcd,
	0x66, 0x90, 0xe8, 0xcd, 0xe6, 0x80, 0x96, 0xd4, 0xa2, 0xdc, 0x77, 0x7c, 0x09, 0xd7, 0x1f, 0x55,
	0xb5, 0x1b, 0xb0, 0xe4, 0xf4, 0xf0, 0x71, 0xf7, 0xe1, 0x4f, 0x20, 0xd2, 0x85, 0xf4, 0x76, 0xb8,
	0x92, 0xaf, 0xe1, 0x66, 0xf5, 0xb3, 0x2c, 0x06, 0x14, 0x67, 0x90, 0x58, 0x61, 0xb6, 0x68, 0xbd,
	0xa0, 0xef, 0x2e, 0x5b, 0xcc, 0x25, 0xdc, 0xac, 0xd0, 0x0a, 0x6b, 0x87, 0x4c, 0x24, 0xc3, 0xc6,
	0x7f, 0x0c, 0x63, 0xb7, 0x70, 0x25, 0xac, 0xf2, 0x8a, 0x51, 0xde, 0x34, 0x0e, 0x2d, 0x09, 0x8d,
	0x1b, 0x94, 0x9a, 0xd7, 0xbf, 0x22, 0xb8, 0x7e, 0xa7, 0x0a, 0x5c, 0x8a, 0x4a, 0x6c, 0xd1, 0xb0,
	0x07, 0x48, 0x9a, 0xdc, 0xb2, 0xd9, 0xdc, 0x67, 0x7a, 0xde, 0x09, 0xf2, 0xf3, 0xb4, 0xc5, 0x29,
	0xe8, 0x8f, 0xd8, 0x2b, 0xb8, 0xa2, 0xfb, 0xb0, 0xa7, 0xed, 0x24, 0xbc, 0x57, 0x9f, 0xf0, 0x00,
	0x49, 0x13, 0xe7, 0x60, 0x47, 0x27, 0xdf, 0x67, 0x77, 0x50, 0x9e, 0x82, 0x1d, 0x61, 0xbe, 0x2e,
	0xec, 0xa0, 0x8b, 0x85, 0x3b, 0x82, 0xbf, 0x42, 0x9f, 0xf2, 0x12, 0x62, 0x97, 0x0b, 0x76, 0xdb,
	0x0e, 0x82, 0x98, 0xf4, 0x7f, 0xbe, 0x80, 0x89, 0xbf, 0x3b, 0x7b, 0xd6, 0xce, 0xba, 0x49, 0x38,
	0x4f, 0x6a, 0x8e, 0x1a, 0x92, 0x3a, 0x67, 0xee, 0x91, 0xbe, 0x24, 0xf4, 0x8d, 0x59, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0x14, 0x68, 0x5c, 0x74, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileManagerClient is the client API for FileManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileManagerClient interface {
	// Sends a greeting
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Void, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*Void, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Void, error)
	Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*Void, error)
	Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*Void, error)
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*Void, error)
	Symlink(ctx context.Context, in *SymlinkRequest, opts ...grpc.CallOption) (*Void, error)
	Setattr(ctx context.Context, in *SetattrRequest, opts ...grpc.CallOption) (*Void, error)
}

type fileManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerClient(cc grpc.ClientConnInterface) FileManagerClient {
	return &fileManagerClient{cc}
}

func (c *fileManagerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Symlink(ctx context.Context, in *SymlinkRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Symlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Setattr(ctx context.Context, in *SetattrRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/network.FileManager/Setattr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileManagerServer is the server API for FileManager service.
type FileManagerServer interface {
	// Sends a greeting
	Create(context.Context, *CreateRequest) (*Void, error)
	Write(context.Context, *WriteRequest) (*Void, error)
	Remove(context.Context, *RemoveRequest) (*Void, error)
	Mkdir(context.Context, *MkdirRequest) (*Void, error)
	Rename(context.Context, *RenameRequest) (*Void, error)
	Link(context.Context, *LinkRequest) (*Void, error)
	Symlink(context.Context, *SymlinkRequest) (*Void, error)
	Setattr(context.Context, *SetattrRequest) (*Void, error)
}

// UnimplementedFileManagerServer can be embedded to have forward compatible implementations.
type UnimplementedFileManagerServer struct {
}

func (*UnimplementedFileManagerServer) Create(ctx context.Context, req *CreateRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedFileManagerServer) Write(ctx context.Context, req *WriteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedFileManagerServer) Remove(ctx context.Context, req *RemoveRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedFileManagerServer) Mkdir(ctx context.Context, req *MkdirRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (*UnimplementedFileManagerServer) Rename(ctx context.Context, req *RenameRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (*UnimplementedFileManagerServer) Link(ctx context.Context, req *LinkRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedFileManagerServer) Symlink(ctx context.Context, req *SymlinkRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Symlink not implemented")
}
func (*UnimplementedFileManagerServer) Setattr(ctx context.Context, req *SetattrRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setattr not implemented")
}

func RegisterFileManagerServer(s *grpc.Server, srv FileManagerServer) {
	s.RegisterService(&_FileManager_serviceDesc, srv)
}

func _FileManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Mkdir(ctx, req.(*MkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Rename(ctx, req.(*RenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Symlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymlinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Symlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Symlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Symlink(ctx, req.(*SymlinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Setattr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetattrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Setattr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.FileManager/Setattr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Setattr(ctx, req.(*SetattrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "network.FileManager",
	HandlerType: (*FileManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FileManager_Create_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _FileManager_Write_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _FileManager_Remove_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _FileManager_Mkdir_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FileManager_Rename_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _FileManager_Link_Handler,
		},
		{
			MethodName: "Symlink",
			Handler:    _FileManager_Symlink_Handler,
		},
		{
			MethodName: "Setattr",
			Handler:    _FileManager_Setattr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
