// Code generated by protoc-gen-go. DO NOT EDIT.
// source: translator/translator.proto

package translator

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

type LangMsg struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LangMsg) Reset()         { *m = LangMsg{} }
func (m *LangMsg) String() string { return proto.CompactTextString(m) }
func (*LangMsg) ProtoMessage()    {}
func (*LangMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{0}
}

func (m *LangMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LangMsg.Unmarshal(m, b)
}
func (m *LangMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LangMsg.Marshal(b, m, deterministic)
}
func (m *LangMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LangMsg.Merge(m, src)
}
func (m *LangMsg) XXX_Size() int {
	return xxx_messageInfo_LangMsg.Size(m)
}
func (m *LangMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_LangMsg.DiscardUnknown(m)
}

var xxx_messageInfo_LangMsg proto.InternalMessageInfo

func (m *LangMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LangMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ArrLangMsg struct {
	Languages            []*LangMsg `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArrLangMsg) Reset()         { *m = ArrLangMsg{} }
func (m *ArrLangMsg) String() string { return proto.CompactTextString(m) }
func (*ArrLangMsg) ProtoMessage()    {}
func (*ArrLangMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{1}
}

func (m *ArrLangMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrLangMsg.Unmarshal(m, b)
}
func (m *ArrLangMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrLangMsg.Marshal(b, m, deterministic)
}
func (m *ArrLangMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrLangMsg.Merge(m, src)
}
func (m *ArrLangMsg) XXX_Size() int {
	return xxx_messageInfo_ArrLangMsg.Size(m)
}
func (m *ArrLangMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrLangMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ArrLangMsg proto.InternalMessageInfo

func (m *ArrLangMsg) GetLanguages() []*LangMsg {
	if m != nil {
		return m.Languages
	}
	return nil
}

type HexLangMsg struct {
	HexID                int32    `protobuf:"varint,1,opt,name=hexID,proto3" json:"hexID,omitempty"`
	LangID               int32    `protobuf:"varint,2,opt,name=langID,proto3" json:"langID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HexLangMsg) Reset()         { *m = HexLangMsg{} }
func (m *HexLangMsg) String() string { return proto.CompactTextString(m) }
func (*HexLangMsg) ProtoMessage()    {}
func (*HexLangMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{2}
}

func (m *HexLangMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HexLangMsg.Unmarshal(m, b)
}
func (m *HexLangMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HexLangMsg.Marshal(b, m, deterministic)
}
func (m *HexLangMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HexLangMsg.Merge(m, src)
}
func (m *HexLangMsg) XXX_Size() int {
	return xxx_messageInfo_HexLangMsg.Size(m)
}
func (m *HexLangMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HexLangMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HexLangMsg proto.InternalMessageInfo

func (m *HexLangMsg) GetHexID() int32 {
	if m != nil {
		return m.HexID
	}
	return 0
}

func (m *HexLangMsg) GetLangID() int32 {
	if m != nil {
		return m.LangID
	}
	return 0
}

type HWLangMsg struct {
	HexLang              *HexLangMsg `protobuf:"bytes,1,opt,name=hexLang,proto3" json:"hexLang,omitempty"`
	Order                int32       `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HWLangMsg) Reset()         { *m = HWLangMsg{} }
func (m *HWLangMsg) String() string { return proto.CompactTextString(m) }
func (*HWLangMsg) ProtoMessage()    {}
func (*HWLangMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{3}
}

func (m *HWLangMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWLangMsg.Unmarshal(m, b)
}
func (m *HWLangMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWLangMsg.Marshal(b, m, deterministic)
}
func (m *HWLangMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWLangMsg.Merge(m, src)
}
func (m *HWLangMsg) XXX_Size() int {
	return xxx_messageInfo_HWLangMsg.Size(m)
}
func (m *HWLangMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HWLangMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HWLangMsg proto.InternalMessageInfo

func (m *HWLangMsg) GetHexLang() *HexLangMsg {
	if m != nil {
		return m.HexLang
	}
	return nil
}

func (m *HWLangMsg) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

type HWTransMsg struct {
	HwLang               *HWLangMsg `protobuf:"bytes,1,opt,name=hwLang,proto3" json:"hwLang,omitempty"`
	Grammar              string     `protobuf:"bytes,2,opt,name=grammar,proto3" json:"grammar,omitempty"`
	Trans                string     `protobuf:"bytes,3,opt,name=trans,proto3" json:"trans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HWTransMsg) Reset()         { *m = HWTransMsg{} }
func (m *HWTransMsg) String() string { return proto.CompactTextString(m) }
func (*HWTransMsg) ProtoMessage()    {}
func (*HWTransMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{4}
}

func (m *HWTransMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWTransMsg.Unmarshal(m, b)
}
func (m *HWTransMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWTransMsg.Marshal(b, m, deterministic)
}
func (m *HWTransMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWTransMsg.Merge(m, src)
}
func (m *HWTransMsg) XXX_Size() int {
	return xxx_messageInfo_HWTransMsg.Size(m)
}
func (m *HWTransMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HWTransMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HWTransMsg proto.InternalMessageInfo

func (m *HWTransMsg) GetHwLang() *HWLangMsg {
	if m != nil {
		return m.HwLang
	}
	return nil
}

func (m *HWTransMsg) GetGrammar() string {
	if m != nil {
		return m.Grammar
	}
	return ""
}

func (m *HWTransMsg) GetTrans() string {
	if m != nil {
		return m.Trans
	}
	return ""
}

type TransMsg struct {
	HexID                int32    `protobuf:"varint,1,opt,name=hexID,proto3" json:"hexID,omitempty"`
	Trans                string   `protobuf:"bytes,2,opt,name=trans,proto3" json:"trans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransMsg) Reset()         { *m = TransMsg{} }
func (m *TransMsg) String() string { return proto.CompactTextString(m) }
func (*TransMsg) ProtoMessage()    {}
func (*TransMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_356ab01a6fa627f4, []int{5}
}

func (m *TransMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransMsg.Unmarshal(m, b)
}
func (m *TransMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransMsg.Marshal(b, m, deterministic)
}
func (m *TransMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransMsg.Merge(m, src)
}
func (m *TransMsg) XXX_Size() int {
	return xxx_messageInfo_TransMsg.Size(m)
}
func (m *TransMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TransMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TransMsg proto.InternalMessageInfo

func (m *TransMsg) GetHexID() int32 {
	if m != nil {
		return m.HexID
	}
	return 0
}

func (m *TransMsg) GetTrans() string {
	if m != nil {
		return m.Trans
	}
	return ""
}

func init() {
	proto.RegisterType((*LangMsg)(nil), "translator.LangMsg")
	proto.RegisterType((*ArrLangMsg)(nil), "translator.ArrLangMsg")
	proto.RegisterType((*HexLangMsg)(nil), "translator.HexLangMsg")
	proto.RegisterType((*HWLangMsg)(nil), "translator.HWLangMsg")
	proto.RegisterType((*HWTransMsg)(nil), "translator.HWTransMsg")
	proto.RegisterType((*TransMsg)(nil), "translator.TransMsg")
}

func init() { proto.RegisterFile("translator/translator.proto", fileDescriptor_356ab01a6fa627f4) }

var fileDescriptor_356ab01a6fa627f4 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x4d, 0x6c, 0x8c, 0x19, 0xa1, 0x87, 0xad, 0x95, 0x10, 0x2f, 0xb2, 0x27, 0x2f, 0xc6,
	0x36, 0x85, 0x16, 0xa4, 0x7f, 0x28, 0x58, 0x50, 0xb0, 0x97, 0x44, 0xf0, 0xbc, 0xe2, 0x36, 0x4a,
	0x4d, 0x22, 0x6b, 0xa4, 0xf6, 0x8b, 0xf5, 0xf3, 0x95, 0xec, 0x9f, 0x24, 0x42, 0xec, 0xa1, 0xb7,
	0x9d, 0xc9, 0xfb, 0xbd, 0x37, 0x99, 0x81, 0x6e, 0xca, 0x48, 0xbc, 0xdf, 0x92, 0x34, 0x61, 0xc3,
	0xe2, 0xe9, 0xee, 0x58, 0x92, 0x26, 0x08, 0x8a, 0x8e, 0xd3, 0x0d, 0x93, 0x24, 0xdc, 0xd2, 0x21,
	0xff, 0xb2, 0x3c, 0x7c, 0x0c, 0x69, 0xb4, 0x4b, 0xbf, 0x85, 0x10, 0x0f, 0xc0, 0x9c, 0x91, 0x38,
	0x7c, 0xdf, 0x87, 0xe8, 0x12, 0xf4, 0xcd, 0xca, 0xd6, 0x7a, 0x5a, 0xdf, 0xf0, 0xf5, 0xcd, 0x0a,
	0x21, 0xb8, 0x88, 0x49, 0x44, 0x6d, 0xbd, 0xa7, 0xf5, 0x2d, 0x9f, 0xbf, 0xf1, 0x0b, 0xc0, 0x2b,
	0x63, 0x8a, 0xb8, 0x05, 0x6b, 0x4b, 0xe2, 0xf0, 0x40, 0x42, 0xba, 0xb7, 0xb5, 0x5e, 0xbd, 0xdf,
	0xf2, 0xae, 0xdc, 0xd2, 0x2c, 0x52, 0xe7, 0x17, 0x2a, 0x3c, 0x02, 0x98, 0xd0, 0xa3, 0x32, 0x68,
	0x83, 0xb1, 0xa6, 0xc7, 0xe9, 0x58, 0xa6, 0x8a, 0x02, 0x75, 0xa0, 0x91, 0x01, 0xd3, 0x31, 0x8f,
	0x36, 0x7c, 0x59, 0xe1, 0x00, 0xac, 0xc9, 0x42, 0xa1, 0x37, 0x60, 0xae, 0x85, 0x11, 0x87, 0x5b,
	0x5e, 0xa7, 0x9c, 0x5c, 0x64, 0xf8, 0x4a, 0x96, 0x85, 0x25, 0x6c, 0x45, 0x99, 0x74, 0x15, 0x05,
	0xfe, 0x04, 0x98, 0x2c, 0xe6, 0x19, 0x99, 0xb9, 0x0e, 0xa0, 0xb1, 0xfe, 0x2a, 0x99, 0x5e, 0x9f,
	0x98, 0xaa, 0x70, 0x5f, 0x8a, 0x90, 0x0d, 0x66, 0xc8, 0x48, 0x14, 0x11, 0x26, 0xb7, 0xa4, 0xca,
	0x2c, 0x8c, 0x93, 0x76, 0x9d, 0xf7, 0x45, 0x81, 0xef, 0xa1, 0x99, 0x47, 0x55, 0xff, 0x7b, 0xce,
	0xe9, 0x25, 0xce, 0xfb, 0xa9, 0x03, 0xcc, 0xf3, 0x41, 0xd0, 0x13, 0x58, 0x33, 0xb5, 0x51, 0xd4,
	0x71, 0xc5, 0x7d, 0x5d, 0x75, 0x5f, 0xf7, 0x2d, 0xbb, 0xaf, 0x73, 0xb2, 0x8f, 0xe2, 0x68, 0xb8,
	0x86, 0x1e, 0xa1, 0x15, 0xd0, 0x54, 0x39, 0xa0, 0xaa, 0x93, 0x39, 0x67, 0x5c, 0x71, 0x0d, 0x8d,
	0xc0, 0x94, 0x0b, 0x43, 0xd5, 0xdb, 0x39, 0x4d, 0x2e, 0x96, 0x8b, 0x6b, 0xe8, 0x19, 0x20, 0xa0,
	0xa9, 0xc2, 0xcf, 0xe8, 0xfe, 0xcc, 0x6e, 0x2e, 0xbc, 0x4a, 0x3a, 0xbf, 0xb7, 0xd3, 0x2e, 0xf7,
	0x4b, 0xd9, 0x0f, 0x60, 0xfc, 0x0f, 0x1c, 0x41, 0x33, 0xa0, 0xa9, 0x60, 0x2b, 0x35, 0xe7, 0x07,
	0x5e, 0x36, 0x78, 0xe7, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x72, 0xc4, 0xd1, 0x4f, 0xad, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslatorClient interface {
	Languages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArrLangMsg, error)
	SetLanguage(ctx context.Context, in *LangMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	HWTrans(ctx context.Context, in *HWLangMsg, opts ...grpc.CallOption) (*HWTransMsg, error)
	SetHWTrans(ctx context.Context, in *HWTransMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	W2WTrans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error)
	Trans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error)
	SetTrans(ctx context.Context, in *TransMsg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type translatorClient struct {
	cc *grpc.ClientConn
}

func NewTranslatorClient(cc *grpc.ClientConn) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) Languages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArrLangMsg, error) {
	out := new(ArrLangMsg)
	err := c.cc.Invoke(ctx, "/translator.Translator/Languages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) SetLanguage(ctx context.Context, in *LangMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/translator.Translator/SetLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) HWTrans(ctx context.Context, in *HWLangMsg, opts ...grpc.CallOption) (*HWTransMsg, error) {
	out := new(HWTransMsg)
	err := c.cc.Invoke(ctx, "/translator.Translator/HWTrans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) SetHWTrans(ctx context.Context, in *HWTransMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/translator.Translator/SetHWTrans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) W2WTrans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error) {
	out := new(TransMsg)
	err := c.cc.Invoke(ctx, "/translator.Translator/W2WTrans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) Trans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error) {
	out := new(TransMsg)
	err := c.cc.Invoke(ctx, "/translator.Translator/Trans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) SetTrans(ctx context.Context, in *TransMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/translator.Translator/SetTrans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
type TranslatorServer interface {
	Languages(context.Context, *empty.Empty) (*ArrLangMsg, error)
	SetLanguage(context.Context, *LangMsg) (*empty.Empty, error)
	HWTrans(context.Context, *HWLangMsg) (*HWTransMsg, error)
	SetHWTrans(context.Context, *HWTransMsg) (*empty.Empty, error)
	W2WTrans(context.Context, *HexLangMsg) (*TransMsg, error)
	Trans(context.Context, *HexLangMsg) (*TransMsg, error)
	SetTrans(context.Context, *TransMsg) (*empty.Empty, error)
}

// UnimplementedTranslatorServer can be embedded to have forward compatible implementations.
type UnimplementedTranslatorServer struct {
}

func (*UnimplementedTranslatorServer) Languages(ctx context.Context, req *empty.Empty) (*ArrLangMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Languages not implemented")
}
func (*UnimplementedTranslatorServer) SetLanguage(ctx context.Context, req *LangMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLanguage not implemented")
}
func (*UnimplementedTranslatorServer) HWTrans(ctx context.Context, req *HWLangMsg) (*HWTransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HWTrans not implemented")
}
func (*UnimplementedTranslatorServer) SetHWTrans(ctx context.Context, req *HWTransMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHWTrans not implemented")
}
func (*UnimplementedTranslatorServer) W2WTrans(ctx context.Context, req *HexLangMsg) (*TransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method W2WTrans not implemented")
}
func (*UnimplementedTranslatorServer) Trans(ctx context.Context, req *HexLangMsg) (*TransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trans not implemented")
}
func (*UnimplementedTranslatorServer) SetTrans(ctx context.Context, req *TransMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrans not implemented")
}

func RegisterTranslatorServer(s *grpc.Server, srv TranslatorServer) {
	s.RegisterService(&_Translator_serviceDesc, srv)
}

func _Translator_Languages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Languages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/Languages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Languages(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_SetLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LangMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).SetLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/SetLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).SetLanguage(ctx, req.(*LangMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_HWTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HWLangMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).HWTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/HWTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).HWTrans(ctx, req.(*HWLangMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_SetHWTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HWTransMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).SetHWTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/SetHWTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).SetHWTrans(ctx, req.(*HWTransMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_W2WTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLangMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).W2WTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/W2WTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).W2WTrans(ctx, req.(*HexLangMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_Trans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLangMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Trans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/Trans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Trans(ctx, req.(*HexLangMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_SetTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).SetTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/SetTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).SetTrans(ctx, req.(*TransMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Translator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "translator.Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Languages",
			Handler:    _Translator_Languages_Handler,
		},
		{
			MethodName: "SetLanguage",
			Handler:    _Translator_SetLanguage_Handler,
		},
		{
			MethodName: "HWTrans",
			Handler:    _Translator_HWTrans_Handler,
		},
		{
			MethodName: "SetHWTrans",
			Handler:    _Translator_SetHWTrans_Handler,
		},
		{
			MethodName: "W2WTrans",
			Handler:    _Translator_W2WTrans_Handler,
		},
		{
			MethodName: "Trans",
			Handler:    _Translator_Trans_Handler,
		},
		{
			MethodName: "SetTrans",
			Handler:    _Translator_SetTrans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translator/translator.proto",
}