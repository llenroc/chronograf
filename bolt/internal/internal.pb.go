// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Exploration
	Source
	Server
	Layout
	Cell
	Query
	AlertRule
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Exploration struct {
	ID        int64  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	UserID    int64  `protobuf:"varint,3,opt,name=UserID,json=userID,proto3" json:"UserID,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=CreatedAt,json=createdAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,6,opt,name=UpdatedAt,json=updatedAt,proto3" json:"UpdatedAt,omitempty"`
	Default   bool   `protobuf:"varint,7,opt,name=Default,json=default,proto3" json:"Default,omitempty"`
}

func (m *Exploration) Reset()                    { *m = Exploration{} }
func (m *Exploration) String() string            { return proto.CompactTextString(m) }
func (*Exploration) ProtoMessage()               {}
func (*Exploration) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Source struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,6,opt,name=URL,json=uRL,proto3" json:"URL,omitempty"`
	Default  bool   `protobuf:"varint,7,opt,name=Default,json=default,proto3" json:"Default,omitempty"`
	Telegraf string `protobuf:"bytes,8,opt,name=Telegraf,json=telegraf,proto3" json:"Telegraf,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

type Server struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,5,opt,name=URL,json=uRL,proto3" json:"URL,omitempty"`
	SrcID    int64  `protobuf:"varint,6,opt,name=SrcID,json=srcID,proto3" json:"SrcID,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

type Layout struct {
	ID          string  `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Application string  `protobuf:"bytes,2,opt,name=Application,json=application,proto3" json:"Application,omitempty"`
	Measurement string  `protobuf:"bytes,3,opt,name=Measurement,json=measurement,proto3" json:"Measurement,omitempty"`
	Cells       []*Cell `protobuf:"bytes,4,rep,name=Cells,json=cells" json:"Cells,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32    `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32    `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string   `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64  `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	Command  string   `protobuf:"bytes,1,opt,name=Command,json=command,proto3" json:"Command,omitempty"`
	DB       string   `protobuf:"bytes,2,opt,name=DB,json=dB,proto3" json:"DB,omitempty"`
	RP       string   `protobuf:"bytes,3,opt,name=RP,json=rP,proto3" json:"RP,omitempty"`
	GroupBys []string `protobuf:"bytes,4,rep,name=GroupBys,json=groupBys" json:"GroupBys,omitempty"`
	Wheres   []string `protobuf:"bytes,5,rep,name=Wheres,json=wheres" json:"Wheres,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,json=jSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,json=srcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,json=kapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

func init() {
	proto.RegisterType((*Exploration)(nil), "internal.Exploration")
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0xe5, 0x49, 0x9c, 0xc4, 0x1e, 0x54, 0x90, 0x85, 0x90, 0x85, 0x58, 0x44, 0x11, 0x8b,
	0x61, 0xd3, 0x05, 0x3c, 0x41, 0xdb, 0x20, 0x34, 0x50, 0xda, 0xe2, 0x69, 0xc5, 0x8a, 0x85, 0x3b,
	0xb9, 0x6d, 0x03, 0xf9, 0xc3, 0xb1, 0x99, 0x66, 0xcb, 0x9a, 0xb7, 0xe1, 0x01, 0xe0, 0xd1, 0x90,
	0x1d, 0x87, 0x19, 0x89, 0x1f, 0x75, 0xf9, 0xf9, 0xdc, 0xf1, 0x3d, 0xf7, 0xf8, 0x4e, 0xe8, 0x5e,
	0xd9, 0x68, 0x50, 0x8d, 0xac, 0xf6, 0x3b, 0xd5, 0xea, 0x96, 0x25, 0x13, 0x67, 0xdf, 0x11, 0x9d,
	0xbf, 0xbc, 0xed, 0xaa, 0x56, 0x49, 0x5d, 0xb6, 0x0d, 0xdb, 0xa3, 0xb3, 0x65, 0xce, 0x51, 0x8a,
	0x16, 0x81, 0x98, 0x95, 0x39, 0x63, 0x34, 0x3c, 0x91, 0x35, 0xf0, 0x59, 0x8a, 0x16, 0x44, 0x84,
	0x8d, 0xac, 0x81, 0x3d, 0xa2, 0xd1, 0x45, 0x0f, 0x6a, 0x99, 0xf3, 0xc0, 0xd5, 0x45, 0xc6, 0x91,
	0xad, 0xcd, 0xa5, 0x96, 0x3c, 0x1c, 0x6b, 0x0b, 0xa9, 0x25, 0x7b, 0x42, 0xc9, 0x91, 0x02, 0xa9,
	0xa1, 0x38, 0xd0, 0x1c, 0xbb, 0x72, 0xb2, 0x9e, 0x0e, 0xac, 0x7a, 0xd1, 0x15, 0x5e, 0x8d, 0x46,
	0xd5, 0x4c, 0x07, 0x8c, 0xd3, 0x38, 0x87, 0x2b, 0x69, 0x2a, 0xcd, 0xe3, 0x14, 0x2d, 0x12, 0x11,
	0x17, 0x23, 0x66, 0x3f, 0x11, 0x8d, 0x56, 0xad, 0x51, 0x6b, 0xb8, 0x93, 0x61, 0x46, 0xc3, 0xf3,
	0xa1, 0x03, 0x67, 0x97, 0x88, 0x50, 0x0f, 0x1d, 0xb0, 0xc7, 0x34, 0xb1, 0x43, 0x58, 0xdd, 0x1b,
	0x4e, 0x8c, 0x67, 0xab, 0x9d, 0xc9, 0xbe, 0xdf, 0xb4, 0xaa, 0x70, 0x9e, 0x89, 0x48, 0x3a, 0xcf,
	0xec, 0x01, 0x0d, 0x2e, 0xc4, 0xb1, 0x33, 0x4b, 0x44, 0x60, 0xc4, 0xf1, 0xbf, 0x6d, 0xda, 0x7b,
	0xce, 0xa1, 0x82, 0x6b, 0x25, 0xaf, 0x78, 0x32, 0xde, 0xa3, 0x3d, 0x67, 0xdf, 0xec, 0x08, 0xa0,
	0xbe, 0x80, 0xba, 0xd3, 0x08, 0xbb, 0x76, 0x83, 0xff, 0xd8, 0x0d, 0xff, 0x6e, 0x17, 0x6f, 0xed,
	0x3e, 0xa4, 0x78, 0xa5, 0xd6, 0xcb, 0xdc, 0xe7, 0x8d, 0x7b, 0x0b, 0xd9, 0x57, 0x44, 0xa3, 0x63,
	0x39, 0xb4, 0x46, 0xef, 0xd8, 0x21, 0xce, 0x4e, 0x4a, 0xe7, 0x07, 0x5d, 0x57, 0x95, 0x6b, 0xb7,
	0x21, 0xde, 0xd5, 0x5c, 0x6e, 0x8f, 0x6c, 0xc5, 0x5b, 0x90, 0xbd, 0x51, 0x50, 0x43, 0xa3, 0xbd,
	0xbf, 0x79, 0xbd, 0x3d, 0x62, 0x4f, 0x29, 0x3e, 0x82, 0xaa, 0xea, 0x79, 0x98, 0x06, 0x8b, 0xf9,
	0xf3, 0xbd, 0xfd, 0xdf, 0x0b, 0x69, 0x8f, 0x05, 0x5e, 0x5b, 0x31, 0xfb, 0x81, 0x68, 0x68, 0x99,
	0xdd, 0xa3, 0xe8, 0xd6, 0x39, 0xc0, 0x02, 0xdd, 0x5a, 0x1a, 0x5c, 0x5b, 0x2c, 0xd0, 0x60, 0x69,
	0xe3, 0x5a, 0x60, 0x81, 0x36, 0x96, 0x6e, 0xdc, 0xd0, 0x58, 0xa0, 0x1b, 0xf6, 0x8c, 0xc6, 0x9f,
	0x0d, 0xa8, 0x12, 0x7a, 0x8e, 0x5d, 0xa3, 0xfb, 0xdb, 0x46, 0xef, 0x0c, 0xa8, 0x41, 0x4c, 0xba,
	0xfd, 0x61, 0xe9, 0x5f, 0x11, 0x95, 0x36, 0x72, 0x17, 0x6d, 0xbc, 0x13, 0x39, 0xa7, 0xf1, 0xa0,
	0x64, 0x73, 0x0d, 0x3d, 0x4f, 0xd2, 0x60, 0x11, 0x88, 0x09, 0x9d, 0x52, 0xc9, 0x4b, 0xa8, 0x7a,
	0x4e, 0xd2, 0x60, 0x41, 0xc4, 0x84, 0x99, 0xa1, 0xd8, 0xf5, 0xb1, 0x25, 0x47, 0x6d, 0x5d, 0xcb,
	0xa6, 0xf0, 0x49, 0xc6, 0xeb, 0x11, 0x6d, 0xbc, 0xf9, 0xa1, 0x4f, 0x71, 0x56, 0x1c, 0x5a, 0x16,
	0x67, 0x3e, 0xb3, 0x99, 0x3a, 0xb3, 0xaf, 0xf9, 0x4a, 0xb5, 0xa6, 0x3b, 0x1c, 0xc6, 0xb4, 0x88,
	0x48, 0xae, 0x3d, 0xdb, 0x7f, 0xde, 0xfb, 0x1b, 0x50, 0x7e, 0x3c, 0x22, 0xa2, 0x8d, 0xa3, 0xec,
	0x03, 0x25, 0x07, 0x15, 0x28, 0x2d, 0x4c, 0x05, 0x7f, 0xbc, 0x1f, 0xa3, 0xe1, 0xeb, 0xd5, 0xe9,
	0xc9, 0xb4, 0x4e, 0x1f, 0x57, 0xa7, 0x27, 0xdb, 0x25, 0x08, 0x76, 0x96, 0xc0, 0x5e, 0xff, 0x46,
	0x76, 0x72, 0x99, 0xbb, 0x44, 0x03, 0x11, 0x7d, 0x72, 0x74, 0x19, 0xb9, 0xaf, 0xc6, 0x8b, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xac, 0x44, 0x65, 0x47, 0x04, 0x00, 0x00,
}
