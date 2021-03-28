// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scanner/scanner.proto

package scanner

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type RangeMsg struct {
	Start                int32    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeMsg) Reset()         { *m = RangeMsg{} }
func (m *RangeMsg) String() string { return proto.CompactTextString(m) }
func (*RangeMsg) ProtoMessage()    {}
func (*RangeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{0}
}

func (m *RangeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangeMsg.Unmarshal(m, b)
}
func (m *RangeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangeMsg.Marshal(b, m, deterministic)
}
func (m *RangeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeMsg.Merge(m, src)
}
func (m *RangeMsg) XXX_Size() int {
	return xxx_messageInfo_RangeMsg.Size(m)
}
func (m *RangeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RangeMsg proto.InternalMessageInfo

func (m *RangeMsg) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *RangeMsg) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

type HexameterMsg struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Para                 bool     `protobuf:"varint,2,opt,name=para,proto3" json:"para,omitempty"`
	Txt                  string   `protobuf:"bytes,3,opt,name=txt,proto3" json:"txt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HexameterMsg) Reset()         { *m = HexameterMsg{} }
func (m *HexameterMsg) String() string { return proto.CompactTextString(m) }
func (*HexameterMsg) ProtoMessage()    {}
func (*HexameterMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{1}
}

func (m *HexameterMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HexameterMsg.Unmarshal(m, b)
}
func (m *HexameterMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HexameterMsg.Marshal(b, m, deterministic)
}
func (m *HexameterMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HexameterMsg.Merge(m, src)
}
func (m *HexameterMsg) XXX_Size() int {
	return xxx_messageInfo_HexameterMsg.Size(m)
}
func (m *HexameterMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HexameterMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HexameterMsg proto.InternalMessageInfo

func (m *HexameterMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HexameterMsg) GetPara() bool {
	if m != nil {
		return m.Para
	}
	return false
}

func (m *HexameterMsg) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

type HexMsg struct {
	Hexes                []*HexameterMsg `protobuf:"bytes,1,rep,name=hexes,proto3" json:"hexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HexMsg) Reset()         { *m = HexMsg{} }
func (m *HexMsg) String() string { return proto.CompactTextString(m) }
func (*HexMsg) ProtoMessage()    {}
func (*HexMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{2}
}

func (m *HexMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HexMsg.Unmarshal(m, b)
}
func (m *HexMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HexMsg.Marshal(b, m, deterministic)
}
func (m *HexMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HexMsg.Merge(m, src)
}
func (m *HexMsg) XXX_Size() int {
	return xxx_messageInfo_HexMsg.Size(m)
}
func (m *HexMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HexMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HexMsg proto.InternalMessageInfo

func (m *HexMsg) GetHexes() []*HexameterMsg {
	if m != nil {
		return m.Hexes
	}
	return nil
}

type HexWordsMsg struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Order                int32    `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	WordId               int32    `protobuf:"varint,3,opt,name=wordId,proto3" json:"wordId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HexWordsMsg) Reset()         { *m = HexWordsMsg{} }
func (m *HexWordsMsg) String() string { return proto.CompactTextString(m) }
func (*HexWordsMsg) ProtoMessage()    {}
func (*HexWordsMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{3}
}

func (m *HexWordsMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HexWordsMsg.Unmarshal(m, b)
}
func (m *HexWordsMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HexWordsMsg.Marshal(b, m, deterministic)
}
func (m *HexWordsMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HexWordsMsg.Merge(m, src)
}
func (m *HexWordsMsg) XXX_Size() int {
	return xxx_messageInfo_HexWordsMsg.Size(m)
}
func (m *HexWordsMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HexWordsMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HexWordsMsg proto.InternalMessageInfo

func (m *HexWordsMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HexWordsMsg) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *HexWordsMsg) GetWordId() int32 {
	if m != nil {
		return m.WordId
	}
	return 0
}

type ParsedMsg struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Para                 bool           `protobuf:"varint,2,opt,name=para,proto3" json:"para,omitempty"`
	Txt                  string         `protobuf:"bytes,3,opt,name=txt,proto3" json:"txt,omitempty"`
	HexWords             []*HexWordsMsg `protobuf:"bytes,4,rep,name=hexWords,proto3" json:"hexWords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ParsedMsg) Reset()         { *m = ParsedMsg{} }
func (m *ParsedMsg) String() string { return proto.CompactTextString(m) }
func (*ParsedMsg) ProtoMessage()    {}
func (*ParsedMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{4}
}

func (m *ParsedMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParsedMsg.Unmarshal(m, b)
}
func (m *ParsedMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParsedMsg.Marshal(b, m, deterministic)
}
func (m *ParsedMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParsedMsg.Merge(m, src)
}
func (m *ParsedMsg) XXX_Size() int {
	return xxx_messageInfo_ParsedMsg.Size(m)
}
func (m *ParsedMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ParsedMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ParsedMsg proto.InternalMessageInfo

func (m *ParsedMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ParsedMsg) GetPara() bool {
	if m != nil {
		return m.Para
	}
	return false
}

func (m *ParsedMsg) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

func (m *ParsedMsg) GetHexWords() []*HexWordsMsg {
	if m != nil {
		return m.HexWords
	}
	return nil
}

type WordMsg struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Txt                  string   `protobuf:"bytes,2,opt,name=txt,proto3" json:"txt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordMsg) Reset()         { *m = WordMsg{} }
func (m *WordMsg) String() string { return proto.CompactTextString(m) }
func (*WordMsg) ProtoMessage()    {}
func (*WordMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{5}
}

func (m *WordMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordMsg.Unmarshal(m, b)
}
func (m *WordMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordMsg.Marshal(b, m, deterministic)
}
func (m *WordMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordMsg.Merge(m, src)
}
func (m *WordMsg) XXX_Size() int {
	return xxx_messageInfo_WordMsg.Size(m)
}
func (m *WordMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WordMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WordMsg proto.InternalMessageInfo

func (m *WordMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WordMsg) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

type TranslatableMsg struct {
	Words                []*WordMsg   `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	Parsed               []*ParsedMsg `protobuf:"bytes,3,rep,name=parsed,proto3" json:"parsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TranslatableMsg) Reset()         { *m = TranslatableMsg{} }
func (m *TranslatableMsg) String() string { return proto.CompactTextString(m) }
func (*TranslatableMsg) ProtoMessage()    {}
func (*TranslatableMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85bb07bffd332d8, []int{6}
}

func (m *TranslatableMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslatableMsg.Unmarshal(m, b)
}
func (m *TranslatableMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslatableMsg.Marshal(b, m, deterministic)
}
func (m *TranslatableMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslatableMsg.Merge(m, src)
}
func (m *TranslatableMsg) XXX_Size() int {
	return xxx_messageInfo_TranslatableMsg.Size(m)
}
func (m *TranslatableMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslatableMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TranslatableMsg proto.InternalMessageInfo

func (m *TranslatableMsg) GetWords() []*WordMsg {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *TranslatableMsg) GetParsed() []*ParsedMsg {
	if m != nil {
		return m.Parsed
	}
	return nil
}

func init() {
	proto.RegisterType((*RangeMsg)(nil), "scanner.RangeMsg")
	proto.RegisterType((*HexameterMsg)(nil), "scanner.HexameterMsg")
	proto.RegisterType((*HexMsg)(nil), "scanner.HexMsg")
	proto.RegisterType((*HexWordsMsg)(nil), "scanner.HexWordsMsg")
	proto.RegisterType((*ParsedMsg)(nil), "scanner.ParsedMsg")
	proto.RegisterType((*WordMsg)(nil), "scanner.WordMsg")
	proto.RegisterType((*TranslatableMsg)(nil), "scanner.TranslatableMsg")
}

func init() { proto.RegisterFile("scanner/scanner.proto", fileDescriptor_c85bb07bffd332d8) }

var fileDescriptor_c85bb07bffd332d8 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0xcb, 0xda, 0x40,
	0x10, 0x35, 0x89, 0x89, 0x3a, 0x96, 0x6a, 0x17, 0x95, 0x60, 0x2f, 0xb2, 0x87, 0x22, 0x15, 0x62,
	0xb1, 0xf4, 0x1f, 0x54, 0x4c, 0x29, 0x42, 0x59, 0x0a, 0x3d, 0xaf, 0xdd, 0x69, 0x14, 0x34, 0x09,
	0xbb, 0x5b, 0x4c, 0x2f, 0xfd, 0xed, 0x1f, 0xbb, 0xd9, 0xe4, 0x0b, 0x1f, 0x9e, 0xbe, 0x93, 0x33,
	0xeb, 0x7b, 0xf3, 0xde, 0x9b, 0x0c, 0xcc, 0xd5, 0x6f, 0x9e, 0xe7, 0x28, 0xb7, 0xee, 0x37, 0x29,
	0x65, 0xa1, 0x0b, 0x32, 0x70, 0xed, 0xf2, 0x7d, 0x56, 0x14, 0xd9, 0x15, 0xb7, 0xf6, 0xf9, 0xf4,
	0xf7, 0xcf, 0x16, 0x6f, 0xa5, 0xfe, 0x57, 0xa3, 0xe8, 0x0e, 0x86, 0x8c, 0xe7, 0x19, 0x1e, 0x55,
	0x46, 0x66, 0x10, 0x2a, 0xcd, 0xa5, 0x8e, 0xbd, 0x95, 0xb7, 0x0e, 0x59, 0xdd, 0x90, 0x29, 0x04,
	0x98, 0x8b, 0xd8, 0xb7, 0x6f, 0xa6, 0xa4, 0x5f, 0xe1, 0x4d, 0x8a, 0x15, 0xbf, 0xa1, 0x46, 0x69,
	0x78, 0x6f, 0xc1, 0xbf, 0x08, 0x47, 0xf2, 0x2f, 0x82, 0x10, 0xe8, 0x97, 0x5c, 0x72, 0x4b, 0x19,
	0x32, 0x5b, 0x9b, 0x29, 0xba, 0xd2, 0x71, 0xb0, 0xf2, 0xd6, 0x23, 0x66, 0x4a, 0xfa, 0x05, 0xa2,
	0x14, 0x2b, 0xc3, 0xdf, 0x40, 0x78, 0xc6, 0x0a, 0x55, 0xec, 0xad, 0x82, 0xf5, 0x78, 0x37, 0x4f,
	0x9a, 0x20, 0x5d, 0x15, 0x56, 0x63, 0xe8, 0x77, 0x18, 0xa7, 0x58, 0xfd, 0x2a, 0xa4, 0x50, 0x8f,
	0xb4, 0x67, 0x10, 0x16, 0x52, 0xa0, 0x74, 0x7e, 0xeb, 0x86, 0x2c, 0x20, 0xba, 0x17, 0x52, 0x7c,
	0x13, 0xd6, 0x40, 0xc8, 0x5c, 0x47, 0x15, 0x8c, 0x7e, 0x70, 0xa9, 0x50, 0xbc, 0x3a, 0x06, 0xf9,
	0x04, 0xc3, 0xb3, 0xf3, 0x13, 0xf7, 0xad, 0xff, 0x59, 0xd7, 0x7f, 0x63, 0x94, 0xb5, 0x28, 0xba,
	0x81, 0x81, 0x29, 0x1e, 0x49, 0xba, 0xf1, 0xfe, 0xf3, 0x96, 0x10, 0x26, 0x3f, 0x25, 0xcf, 0xd5,
	0x95, 0x6b, 0x7e, 0xba, 0xda, 0xcf, 0xf4, 0x01, 0xc2, 0xbb, 0x95, 0xab, 0xd7, 0x35, 0x6d, 0xe5,
	0xdc, 0x54, 0x56, 0xff, 0x4d, 0x3e, 0x42, 0x54, 0xda, 0x70, 0x71, 0x60, 0x81, 0xa4, 0x05, 0xb6,
	0x99, 0x99, 0x43, 0xec, 0xfe, 0x43, 0x90, 0x62, 0x65, 0xc2, 0x1c, 0x50, 0xa7, 0x66, 0xd1, 0xe4,
	0x5d, 0x0b, 0x6f, 0x0e, 0x64, 0x39, 0xe9, 0x26, 0x3b, 0xaa, 0x8c, 0xf6, 0xc8, 0x1e, 0x26, 0x07,
	0xd4, 0x5d, 0x8b, 0x64, 0x91, 0xd4, 0x07, 0x97, 0x34, 0x07, 0x97, 0xec, 0xcd, 0xc1, 0x2d, 0xe3,
	0x96, 0xfd, 0x22, 0x11, 0xed, 0x9d, 0x22, 0x8b, 0xfd, 0xfc, 0x14, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0x4c, 0x91, 0x92, 0xcc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HexClient is the client API for Hex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HexClient interface {
	GetHexes(ctx context.Context, in *RangeMsg, opts ...grpc.CallOption) (*HexMsg, error)
	GetTranslatable(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TranslatableMsg, error)
}

type hexClient struct {
	cc *grpc.ClientConn
}

func NewHexClient(cc *grpc.ClientConn) HexClient {
	return &hexClient{cc}
}

func (c *hexClient) GetHexes(ctx context.Context, in *RangeMsg, opts ...grpc.CallOption) (*HexMsg, error) {
	out := new(HexMsg)
	err := c.cc.Invoke(ctx, "/scanner.Hex/GetHexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexClient) GetTranslatable(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TranslatableMsg, error) {
	out := new(TranslatableMsg)
	err := c.cc.Invoke(ctx, "/scanner.Hex/GetTranslatable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HexServer is the server API for Hex service.
type HexServer interface {
	GetHexes(context.Context, *RangeMsg) (*HexMsg, error)
	GetTranslatable(context.Context, *empty.Empty) (*TranslatableMsg, error)
}

// UnimplementedHexServer can be embedded to have forward compatible implementations.
type UnimplementedHexServer struct {
}

func (*UnimplementedHexServer) GetHexes(ctx context.Context, req *RangeMsg) (*HexMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHexes not implemented")
}
func (*UnimplementedHexServer) GetTranslatable(ctx context.Context, req *empty.Empty) (*TranslatableMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslatable not implemented")
}

func RegisterHexServer(s *grpc.Server, srv HexServer) {
	s.RegisterService(&_Hex_serviceDesc, srv)
}

func _Hex_GetHexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexServer).GetHexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanner.Hex/GetHexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexServer).GetHexes(ctx, req.(*RangeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hex_GetTranslatable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexServer).GetTranslatable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanner.Hex/GetTranslatable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexServer).GetTranslatable(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scanner.Hex",
	HandlerType: (*HexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHexes",
			Handler:    _Hex_GetHexes_Handler,
		},
		{
			MethodName: "GetTranslatable",
			Handler:    _Hex_GetTranslatable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanner/scanner.proto",
}
