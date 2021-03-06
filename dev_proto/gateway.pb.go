// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package proto

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

type Category int32

const (
	Category_CMD    Category = 0
	Category_CONFIG Category = 1
)

var Category_name = map[int32]string{
	0: "CMD",
	1: "CONFIG",
}

var Category_value = map[string]int32{
	"CMD":    0,
	"CONFIG": 1,
}

func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}

func (Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

type Device int32

const (
	Device_LAMP   Device = 0
	Device_HEATER Device = 1
)

var Device_name = map[int32]string{
	0: "LAMP",
	1: "HEATER",
}

var Device_value = map[string]int32{
	"LAMP":   0,
	"HEATER": 1,
}

func (x Device) String() string {
	return proto.EnumName(Device_name, int32(x))
}

func (Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

type Operation int32

const (
	Operation_OFF Operation = 0
	Operation_ON  Operation = 1
)

var Operation_name = map[int32]string{
	0: "OFF",
	1: "ON",
}

var Operation_value = map[string]int32{
	"OFF": 0,
	"ON":  1,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

//message是固定的。UserInfo是类名，可以随意指定，符合规范即可
type UplinkFrame struct {
	FrameType            int32    `protobuf:"varint,1,opt,name=FrameType,proto3" json:"FrameType,omitempty"`
	DevAddr              []byte   `protobuf:"bytes,2,opt,name=DevAddr,proto3" json:"DevAddr,omitempty"`
	FrameNum             uint32   `protobuf:"varint,3,opt,name=FrameNum,proto3" json:"FrameNum,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	PhyPayload           []byte   `protobuf:"bytes,5,opt,name=PhyPayload,proto3" json:"PhyPayload,omitempty"`
	UplinkId             uint32   `protobuf:"varint,6,opt,name=UplinkId,proto3" json:"UplinkId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkFrame) Reset()         { *m = UplinkFrame{} }
func (m *UplinkFrame) String() string { return proto.CompactTextString(m) }
func (*UplinkFrame) ProtoMessage()    {}
func (*UplinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *UplinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrame.Unmarshal(m, b)
}
func (m *UplinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrame.Marshal(b, m, deterministic)
}
func (m *UplinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrame.Merge(m, src)
}
func (m *UplinkFrame) XXX_Size() int {
	return xxx_messageInfo_UplinkFrame.Size(m)
}
func (m *UplinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrame proto.InternalMessageInfo

func (m *UplinkFrame) GetFrameType() int32 {
	if m != nil {
		return m.FrameType
	}
	return 0
}

func (m *UplinkFrame) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *UplinkFrame) GetFrameNum() uint32 {
	if m != nil {
		return m.FrameNum
	}
	return 0
}

func (m *UplinkFrame) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *UplinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *UplinkFrame) GetUplinkId() uint32 {
	if m != nil {
		return m.UplinkId
	}
	return 0
}

type DownlinkFrame struct {
	FrameType            int32    `protobuf:"varint,1,opt,name=FrameType,proto3" json:"FrameType,omitempty"`
	DevAddr              []byte   `protobuf:"bytes,2,opt,name=DevAddr,proto3" json:"DevAddr,omitempty"`
	FrameNum             uint32   `protobuf:"varint,3,opt,name=FrameNum,proto3" json:"FrameNum,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	PhyPayload           []byte   `protobuf:"bytes,5,opt,name=PhyPayload,proto3" json:"PhyPayload,omitempty"`
	DownlinkId           uint32   `protobuf:"varint,6,opt,name=DownlinkId,proto3" json:"DownlinkId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrame) Reset()         { *m = DownlinkFrame{} }
func (m *DownlinkFrame) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrame) ProtoMessage()    {}
func (*DownlinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *DownlinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrame.Unmarshal(m, b)
}
func (m *DownlinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrame.Marshal(b, m, deterministic)
}
func (m *DownlinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrame.Merge(m, src)
}
func (m *DownlinkFrame) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrame.Size(m)
}
func (m *DownlinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrame proto.InternalMessageInfo

func (m *DownlinkFrame) GetFrameType() int32 {
	if m != nil {
		return m.FrameType
	}
	return 0
}

func (m *DownlinkFrame) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *DownlinkFrame) GetFrameNum() uint32 {
	if m != nil {
		return m.FrameNum
	}
	return 0
}

func (m *DownlinkFrame) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DownlinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *DownlinkFrame) GetDownlinkId() uint32 {
	if m != nil {
		return m.DownlinkId
	}
	return 0
}

type Payload struct {
	Kind                 uint32   `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key                  uint32   `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *Payload) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Payload) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Category", Category_name, Category_value)
	proto.RegisterEnum("proto.Device", Device_name, Device_value)
	proto.RegisterEnum("proto.Operation", Operation_name, Operation_value)
	proto.RegisterType((*UplinkFrame)(nil), "proto.UplinkFrame")
	proto.RegisterType((*DownlinkFrame)(nil), "proto.DownlinkFrame")
	proto.RegisterType((*Payload)(nil), "proto.Payload")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x41, 0x6a, 0xc2, 0x40,
	0x14, 0x86, 0x1d, 0x8d, 0x51, 0x5f, 0x0d, 0x0c, 0x6f, 0x35, 0x14, 0xb1, 0xe2, 0x4a, 0x5c, 0x74,
	0xd3, 0x13, 0x04, 0x63, 0x5a, 0xa1, 0x9a, 0x30, 0xd8, 0x03, 0x4c, 0x9b, 0xc1, 0x86, 0xc4, 0x4c,
	0x18, 0xd2, 0xc8, 0x1c, 0xab, 0xab, 0x5e, 0xaf, 0x64, 0x42, 0xac, 0x47, 0xe8, 0x2a, 0xff, 0xff,
	0xe5, 0xf1, 0xf3, 0x2d, 0x06, 0xbc, 0x93, 0xa8, 0xe4, 0x45, 0x98, 0xc7, 0x52, 0xab, 0x4a, 0xe1,
	0xd0, 0x7e, 0x96, 0xdf, 0x04, 0xee, 0xde, 0xca, 0x3c, 0x2d, 0xb2, 0x50, 0x8b, 0xb3, 0xc4, 0x19,
	0x4c, 0x6c, 0x38, 0x9a, 0x52, 0x32, 0xb2, 0x20, 0xab, 0x21, 0xff, 0x03, 0xc8, 0x60, 0x14, 0xc8,
	0xda, 0x4f, 0x12, 0xcd, 0xfa, 0x0b, 0xb2, 0x9a, 0xf2, 0xae, 0xe2, 0x3d, 0x8c, 0xed, 0xd9, 0xe1,
	0xeb, 0xcc, 0x06, 0x0b, 0xb2, 0xf2, 0xf8, 0xb5, 0x23, 0x82, 0x13, 0x2b, 0x5d, 0x31, 0xc7, 0x72,
	0x9b, 0x71, 0x0e, 0x10, 0x7f, 0x9a, 0x58, 0x98, 0x5c, 0x89, 0x84, 0x0d, 0xed, 0xd8, 0x0d, 0x69,
	0xf6, 0x5a, 0xad, 0x5d, 0xc2, 0xdc, 0x76, 0xaf, 0xeb, 0xcb, 0x1f, 0x02, 0x5e, 0xa0, 0x2e, 0xc5,
	0xff, 0xb3, 0x9e, 0x03, 0x74, 0x62, 0x57, 0xef, 0x1b, 0xb2, 0xf4, 0x61, 0xd4, 0x9d, 0x22, 0x38,
	0x59, 0x5a, 0x24, 0xd6, 0xd6, 0xe3, 0x36, 0x23, 0x85, 0x41, 0x26, 0x8d, 0x95, 0xf4, 0x78, 0x13,
	0x1b, 0x52, 0x8b, 0xdc, 0xba, 0x4d, 0x79, 0x13, 0xd7, 0x0f, 0x30, 0xde, 0x88, 0x4a, 0x9e, 0x94,
	0x36, 0x38, 0x82, 0xc1, 0x66, 0x1f, 0xd0, 0x1e, 0x02, 0xb8, 0x9b, 0xe8, 0x10, 0xee, 0x9e, 0x29,
	0x59, 0xcf, 0xc1, 0x0d, 0x64, 0x9d, 0x7e, 0x48, 0x1c, 0x83, 0xf3, 0xea, 0xef, 0xe3, 0xf6, 0xff,
	0xcb, 0xd6, 0x3f, 0x6e, 0x39, 0x25, 0xeb, 0x19, 0x4c, 0xa2, 0x52, 0x6a, 0x51, 0xa5, 0xaa, 0x68,
	0x16, 0xa2, 0x30, 0xa4, 0x3d, 0x74, 0xa1, 0x1f, 0x1d, 0x28, 0x79, 0x77, 0xed, 0xb3, 0x78, 0xfa,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x2a, 0x04, 0xaa, 0x2e, 0x02, 0x00, 0x00,
}
