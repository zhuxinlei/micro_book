// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

package book

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

type IdReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{0}
}

func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReq.Unmarshal(m, b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return xxx_messageInfo_IdReq.Size(m)
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

func (m *IdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BookInfoReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookName             string   `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	PublishTime          int64    `protobuf:"varint,5,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Desc                 string   `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookInfoReply) Reset()         { *m = BookInfoReply{} }
func (m *BookInfoReply) String() string { return proto.CompactTextString(m) }
func (*BookInfoReply) ProtoMessage()    {}
func (*BookInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{1}
}

func (m *BookInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookInfoReply.Unmarshal(m, b)
}
func (m *BookInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookInfoReply.Marshal(b, m, deterministic)
}
func (m *BookInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfoReply.Merge(m, src)
}
func (m *BookInfoReply) XXX_Size() int {
	return xxx_messageInfo_BookInfoReply.Size(m)
}
func (m *BookInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfoReply proto.InternalMessageInfo

func (m *BookInfoReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BookInfoReply) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

func (m *BookInfoReply) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BookInfoReply) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *BookInfoReply) GetPublishTime() int64 {
	if m != nil {
		return m.PublishTime
	}
	return 0
}

func (m *BookInfoReply) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*IdReq)(nil), "book.IdReq")
	proto.RegisterType((*BookInfoReply)(nil), "book.BookInfoReply")
}

func init() { proto.RegisterFile("book.proto", fileDescriptor_1e89d0eaa98dc5d8) }

var fileDescriptor_1e89d0eaa98dc5d8 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0x34, 0x09, 0xf4, 0x0a, 0x0c, 0x07, 0x02, 0x0b, 0x96, 0xd0, 0xc9, 0x12, 0x52,
	0x07, 0xf8, 0x07, 0x6c, 0x5d, 0x18, 0x2c, 0x56, 0x54, 0x25, 0xf5, 0x41, 0xad, 0xd6, 0x3d, 0x93,
	0xb8, 0x03, 0x3f, 0x87, 0x7f, 0x8a, 0x7c, 0xcd, 0x02, 0x0b, 0xdb, 0x7b, 0x4f, 0xcf, 0xbe, 0xfb,
	0x0e, 0xa0, 0x63, 0xde, 0x2e, 0x62, 0xcf, 0x89, 0xb1, 0xcc, 0x7a, 0x7e, 0x03, 0xd5, 0xd2, 0x59,
	0xfa, 0xc4, 0x0b, 0x28, 0xbc, 0xd3, 0xaa, 0x51, 0x66, 0x62, 0x0b, 0xef, 0xe6, 0xdf, 0x0a, 0xce,
	0x9f, 0x99, 0xb7, 0xcb, 0xfd, 0x3b, 0x5b, 0x8a, 0xbb, 0xaf, 0xbf, 0x0d, 0xbc, 0x83, 0x69, 0xfe,
	0x62, 0xb5, 0x6f, 0x03, 0xe9, 0xa2, 0x51, 0x66, 0x6a, 0x4f, 0x73, 0xf0, 0xd2, 0x06, 0xc2, 0x2b,
	0xa8, 0x62, 0xef, 0xd7, 0xa4, 0x27, 0x8d, 0x32, 0xca, 0x1e, 0x0d, 0x5e, 0x43, 0xdd, 0x1e, 0xd2,
	0x86, 0x7b, 0x5d, 0x4a, 0x7f, 0x74, 0x78, 0x0f, 0x67, 0xf1, 0xd0, 0xed, 0xfc, 0xb0, 0x59, 0x25,
	0x1f, 0x48, 0x57, 0x32, 0x64, 0x36, 0x66, 0xaf, 0x3e, 0x10, 0x22, 0x94, 0x8e, 0x86, 0xb5, 0xae,
	0xe5, 0xa1, 0xe8, 0xc7, 0x37, 0x10, 0x08, 0x7c, 0x80, 0x93, 0x0f, 0x4a, 0x79, 0x5b, 0x9c, 0x2d,
	0x04, 0x51, 0x98, 0x6e, 0x2f, 0x8f, 0xe6, 0x37, 0x86, 0x81, 0x32, 0xd1, 0x90, 0xfe, 0x6f, 0x76,
	0xb5, 0x1c, 0xea, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x43, 0x19, 0xbf, 0x08, 0x36, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookClient interface {
	GetBook(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BookInfoReply, error)
	Test(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BookInfoReply, error)
}

type bookClient struct {
	cc *grpc.ClientConn
}

func NewBookClient(cc *grpc.ClientConn) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetBook(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BookInfoReply, error) {
	out := new(BookInfoReply)
	err := c.cc.Invoke(ctx, "/book.book/getBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Test(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BookInfoReply, error) {
	out := new(BookInfoReply)
	err := c.cc.Invoke(ctx, "/book.book/test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
type BookServer interface {
	GetBook(context.Context, *IdReq) (*BookInfoReply, error)
	Test(context.Context, *IdReq) (*BookInfoReply, error)
}

// UnimplementedBookServer can be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (*UnimplementedBookServer) GetBook(ctx context.Context, req *IdReq) (*BookInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookServer) Test(ctx context.Context, req *IdReq) (*BookInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterBookServer(s *grpc.Server, srv BookServer) {
	s.RegisterService(&_Book_serviceDesc, srv)
}

func _Book_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.book/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBook(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.book/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Test(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Book_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBook",
			Handler:    _Book_GetBook_Handler,
		},
		{
			MethodName: "test",
			Handler:    _Book_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
