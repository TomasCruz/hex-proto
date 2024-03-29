// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: translator/translator.proto

package translator

import (
	context "context"
	scanner "github.com/TomasCruz/hex-proto/scanner"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TranslatorClient is the client API for Translator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslatorClient interface {
	Languages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArrLangMsg, error)
	SetLanguage(ctx context.Context, in *LangMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	HWTrans(ctx context.Context, in *HWLangMsg, opts ...grpc.CallOption) (*HWTransMsg, error)
	SetHWTrans(ctx context.Context, in *HWTransMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	W2WTrans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error)
	SetW2WTrans(ctx context.Context, in *TransMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	Trans(ctx context.Context, in *HexLangMsg, opts ...grpc.CallOption) (*TransMsg, error)
	SetTrans(ctx context.Context, in *TransMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	HexWordCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountMsg, error)
	SetTranslatable(ctx context.Context, in *scanner.ArrTranslatableMsg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type translatorClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorClient(cc grpc.ClientConnInterface) TranslatorClient {
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

func (c *translatorClient) SetW2WTrans(ctx context.Context, in *TransMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/translator.Translator/SetW2WTrans", in, out, opts...)
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

func (c *translatorClient) HexWordCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountMsg, error) {
	out := new(CountMsg)
	err := c.cc.Invoke(ctx, "/translator.Translator/HexWordCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorClient) SetTranslatable(ctx context.Context, in *scanner.ArrTranslatableMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/translator.Translator/SetTranslatable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServer is the server API for Translator service.
// All implementations must embed UnimplementedTranslatorServer
// for forward compatibility
type TranslatorServer interface {
	Languages(context.Context, *empty.Empty) (*ArrLangMsg, error)
	SetLanguage(context.Context, *LangMsg) (*empty.Empty, error)
	HWTrans(context.Context, *HWLangMsg) (*HWTransMsg, error)
	SetHWTrans(context.Context, *HWTransMsg) (*empty.Empty, error)
	W2WTrans(context.Context, *HexLangMsg) (*TransMsg, error)
	SetW2WTrans(context.Context, *TransMsg) (*empty.Empty, error)
	Trans(context.Context, *HexLangMsg) (*TransMsg, error)
	SetTrans(context.Context, *TransMsg) (*empty.Empty, error)
	HexWordCount(context.Context, *empty.Empty) (*CountMsg, error)
	SetTranslatable(context.Context, *scanner.ArrTranslatableMsg) (*empty.Empty, error)
	mustEmbedUnimplementedTranslatorServer()
}

// UnimplementedTranslatorServer must be embedded to have forward compatible implementations.
type UnimplementedTranslatorServer struct {
}

func (UnimplementedTranslatorServer) Languages(context.Context, *empty.Empty) (*ArrLangMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Languages not implemented")
}
func (UnimplementedTranslatorServer) SetLanguage(context.Context, *LangMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLanguage not implemented")
}
func (UnimplementedTranslatorServer) HWTrans(context.Context, *HWLangMsg) (*HWTransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HWTrans not implemented")
}
func (UnimplementedTranslatorServer) SetHWTrans(context.Context, *HWTransMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHWTrans not implemented")
}
func (UnimplementedTranslatorServer) W2WTrans(context.Context, *HexLangMsg) (*TransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method W2WTrans not implemented")
}
func (UnimplementedTranslatorServer) SetW2WTrans(context.Context, *TransMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetW2WTrans not implemented")
}
func (UnimplementedTranslatorServer) Trans(context.Context, *HexLangMsg) (*TransMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trans not implemented")
}
func (UnimplementedTranslatorServer) SetTrans(context.Context, *TransMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTrans not implemented")
}
func (UnimplementedTranslatorServer) HexWordCount(context.Context, *empty.Empty) (*CountMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HexWordCount not implemented")
}
func (UnimplementedTranslatorServer) SetTranslatable(context.Context, *scanner.ArrTranslatableMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTranslatable not implemented")
}
func (UnimplementedTranslatorServer) mustEmbedUnimplementedTranslatorServer() {}

// UnsafeTranslatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslatorServer will
// result in compilation errors.
type UnsafeTranslatorServer interface {
	mustEmbedUnimplementedTranslatorServer()
}

func RegisterTranslatorServer(s grpc.ServiceRegistrar, srv TranslatorServer) {
	s.RegisterService(&Translator_ServiceDesc, srv)
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

func _Translator_SetW2WTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).SetW2WTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/SetW2WTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).SetW2WTrans(ctx, req.(*TransMsg))
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

func _Translator_HexWordCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).HexWordCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/HexWordCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).HexWordCount(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Translator_SetTranslatable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scanner.ArrTranslatableMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).SetTranslatable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.Translator/SetTranslatable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).SetTranslatable(ctx, req.(*scanner.ArrTranslatableMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Translator_ServiceDesc is the grpc.ServiceDesc for Translator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Translator_ServiceDesc = grpc.ServiceDesc{
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
			MethodName: "SetW2WTrans",
			Handler:    _Translator_SetW2WTrans_Handler,
		},
		{
			MethodName: "Trans",
			Handler:    _Translator_Trans_Handler,
		},
		{
			MethodName: "SetTrans",
			Handler:    _Translator_SetTrans_Handler,
		},
		{
			MethodName: "HexWordCount",
			Handler:    _Translator_HexWordCount_Handler,
		},
		{
			MethodName: "SetTranslatable",
			Handler:    _Translator_SetTranslatable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translator/translator.proto",
}
