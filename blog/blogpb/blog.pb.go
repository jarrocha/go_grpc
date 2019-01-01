// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
}

func init() { proto.RegisterFile("blogpb/blog.proto", fileDescriptor_1cd072c3eda6f7ba) }

var fileDescriptor_1cd072c3eda6f7ba = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4b, 0x87, 0x40,
	0x10, 0xc5, 0xf9, 0x9a, 0x7d, 0xd3, 0x11, 0x02, 0x87, 0xa0, 0xa5, 0x20, 0xc2, 0x53, 0x27, 0x03,
	0xed, 0x1e, 0xd8, 0xa9, 0x4b, 0x07, 0xbb, 0x75, 0x09, 0x75, 0x07, 0x5b, 0x10, 0xd7, 0xd6, 0xb1,
	0xbf, 0x3f, 0x76, 0x96, 0x28, 0xf0, 0xd0, 0x69, 0xf7, 0xbd, 0xf9, 0xf1, 0x3e, 0x0c, 0xe4, 0xfd,
	0x64, 0xc7, 0xa5, 0xbf, 0xf7, 0x4f, 0xb9, 0x38, 0xcb, 0x16, 0x63, 0xff, 0x2f, 0x06, 0x88, 0x9b,
	0xc9, 0x8e, 0x78, 0x0e, 0x91, 0xd1, 0xea, 0x70, 0x7b, 0xb8, 0x4b, 0xdb, 0xc8, 0x68, 0xbc, 0x86,
	0xb4, 0xdb, 0xf8, 0xc3, 0xba, 0x77, 0xa3, 0x55, 0x24, 0x76, 0x12, 0x8c, 0x67, 0x8d, 0x17, 0x70,
	0xca, 0x86, 0x27, 0x52, 0x27, 0x52, 0x08, 0x02, 0x15, 0x9c, 0x0d, 0x76, 0x66, 0x9a, 0x59, 0xc5,
	0xe2, 0xff, 0xc8, 0xa2, 0x86, 0xfc, 0xc9, 0x51, 0xc7, 0xe4, 0xa3, 0x5a, 0xfa, 0xdc, 0x68, 0x65,
	0xbc, 0x01, 0x21, 0x90, 0xcc, 0xac, 0x82, 0x52, 0xd0, 0xa4, 0x21, 0x90, 0x3d, 0x00, 0xfe, 0x1d,
	0x5a, 0x17, 0x3b, 0xaf, 0xf4, 0xdf, 0x54, 0xf5, 0x02, 0x99, 0x57, 0xaf, 0xe4, 0xbe, 0xcc, 0x40,
	0xf8, 0x08, 0xf0, 0xbb, 0x04, 0x2f, 0x43, 0xfb, 0x8e, 0xe5, 0x4a, 0xed, 0x0b, 0x21, 0xaf, 0x49,
	0xde, 0x8e, 0xe1, 0x74, 0xfd, 0x51, 0xce, 0x56, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x58, 0x20,
	0x1e, 0x8c, 0x4b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogpb/blog.proto",
}