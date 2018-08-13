// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/artist.proto

package no_sysco_middleware_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Artist struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artist) Reset()         { *m = Artist{} }
func (m *Artist) String() string { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()    {}
func (*Artist) Descriptor() ([]byte, []int) {
	return fileDescriptor_artist_4fd6a624dc58df3c, []int{0}
}
func (m *Artist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artist.Unmarshal(m, b)
}
func (m *Artist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artist.Marshal(b, m, deterministic)
}
func (dst *Artist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artist.Merge(dst, src)
}
func (m *Artist) XXX_Size() int {
	return xxx_messageInfo_Artist.Size(m)
}
func (m *Artist) XXX_DiscardUnknown() {
	xxx_messageInfo_Artist.DiscardUnknown(m)
}

var xxx_messageInfo_Artist proto.InternalMessageInfo

func (m *Artist) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artist) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SimpleArtistRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleArtistRequest) Reset()         { *m = SimpleArtistRequest{} }
func (m *SimpleArtistRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleArtistRequest) ProtoMessage()    {}
func (*SimpleArtistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artist_4fd6a624dc58df3c, []int{1}
}
func (m *SimpleArtistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleArtistRequest.Unmarshal(m, b)
}
func (m *SimpleArtistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleArtistRequest.Marshal(b, m, deterministic)
}
func (dst *SimpleArtistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleArtistRequest.Merge(dst, src)
}
func (m *SimpleArtistRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleArtistRequest.Size(m)
}
func (m *SimpleArtistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleArtistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleArtistRequest proto.InternalMessageInfo

func (m *SimpleArtistRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Artist)(nil), "no.sysco.middleware.grpc.Artist")
	proto.RegisterType((*SimpleArtistRequest)(nil), "no.sysco.middleware.grpc.SimpleArtistRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArtistServiceClient interface {
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ArtistService_GetAllClient, error)
	GetArtistByAlbum(ctx context.Context, in *SimpleArtistRequest, opts ...grpc.CallOption) (ArtistService_GetArtistByAlbumClient, error)
	GetArtistByTrack(ctx context.Context, in *SimpleArtistRequest, opts ...grpc.CallOption) (ArtistService_GetArtistByTrackClient, error)
}

type artistServiceClient struct {
	cc *grpc.ClientConn
}

func NewArtistServiceClient(cc *grpc.ClientConn) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ArtistService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArtistService_serviceDesc.Streams[0], "/no.sysco.middleware.grpc.ArtistService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &artistServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArtistService_GetAllClient interface {
	Recv() (*Artist, error)
	grpc.ClientStream
}

type artistServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *artistServiceGetAllClient) Recv() (*Artist, error) {
	m := new(Artist)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *artistServiceClient) GetArtistByAlbum(ctx context.Context, in *SimpleArtistRequest, opts ...grpc.CallOption) (ArtistService_GetArtistByAlbumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArtistService_serviceDesc.Streams[1], "/no.sysco.middleware.grpc.ArtistService/GetArtistByAlbum", opts...)
	if err != nil {
		return nil, err
	}
	x := &artistServiceGetArtistByAlbumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArtistService_GetArtistByAlbumClient interface {
	Recv() (*Artist, error)
	grpc.ClientStream
}

type artistServiceGetArtistByAlbumClient struct {
	grpc.ClientStream
}

func (x *artistServiceGetArtistByAlbumClient) Recv() (*Artist, error) {
	m := new(Artist)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *artistServiceClient) GetArtistByTrack(ctx context.Context, in *SimpleArtistRequest, opts ...grpc.CallOption) (ArtistService_GetArtistByTrackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArtistService_serviceDesc.Streams[2], "/no.sysco.middleware.grpc.ArtistService/GetArtistByTrack", opts...)
	if err != nil {
		return nil, err
	}
	x := &artistServiceGetArtistByTrackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArtistService_GetArtistByTrackClient interface {
	Recv() (*Artist, error)
	grpc.ClientStream
}

type artistServiceGetArtistByTrackClient struct {
	grpc.ClientStream
}

func (x *artistServiceGetArtistByTrackClient) Recv() (*Artist, error) {
	m := new(Artist)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArtistServiceServer is the server API for ArtistService service.
type ArtistServiceServer interface {
	GetAll(*empty.Empty, ArtistService_GetAllServer) error
	GetArtistByAlbum(*SimpleArtistRequest, ArtistService_GetArtistByAlbumServer) error
	GetArtistByTrack(*SimpleArtistRequest, ArtistService_GetArtistByTrackServer) error
}

func RegisterArtistServiceServer(s *grpc.Server, srv ArtistServiceServer) {
	s.RegisterService(&_ArtistService_serviceDesc, srv)
}

func _ArtistService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArtistServiceServer).GetAll(m, &artistServiceGetAllServer{stream})
}

type ArtistService_GetAllServer interface {
	Send(*Artist) error
	grpc.ServerStream
}

type artistServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *artistServiceGetAllServer) Send(m *Artist) error {
	return x.ServerStream.SendMsg(m)
}

func _ArtistService_GetArtistByAlbum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleArtistRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArtistServiceServer).GetArtistByAlbum(m, &artistServiceGetArtistByAlbumServer{stream})
}

type ArtistService_GetArtistByAlbumServer interface {
	Send(*Artist) error
	grpc.ServerStream
}

type artistServiceGetArtistByAlbumServer struct {
	grpc.ServerStream
}

func (x *artistServiceGetArtistByAlbumServer) Send(m *Artist) error {
	return x.ServerStream.SendMsg(m)
}

func _ArtistService_GetArtistByTrack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleArtistRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArtistServiceServer).GetArtistByTrack(m, &artistServiceGetArtistByTrackServer{stream})
}

type ArtistService_GetArtistByTrackServer interface {
	Send(*Artist) error
	grpc.ServerStream
}

type artistServiceGetArtistByTrackServer struct {
	grpc.ServerStream
}

func (x *artistServiceGetArtistByTrackServer) Send(m *Artist) error {
	return x.ServerStream.SendMsg(m)
}

var _ArtistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "no.sysco.middleware.grpc.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _ArtistService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetArtistByAlbum",
			Handler:       _ArtistService_GetArtistByAlbum_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetArtistByTrack",
			Handler:       _ArtistService_GetArtistByTrack_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema/artist.proto",
}

func init() { proto.RegisterFile("schema/artist.proto", fileDescriptor_artist_4fd6a624dc58df3c) }

var fileDescriptor_artist_4fd6a624dc58df3c = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xc1, 0x4a, 0xfc, 0x30,
	0x10, 0xc6, 0x69, 0xff, 0x7f, 0x0a, 0x8e, 0x28, 0x92, 0x15, 0x29, 0xeb, 0xa5, 0x2c, 0x08, 0x5e,
	0x4c, 0x17, 0x7d, 0x82, 0x96, 0x15, 0x6f, 0x1e, 0x76, 0x7d, 0x81, 0x34, 0x9d, 0xad, 0x83, 0x49,
	0x53, 0x93, 0x54, 0xe9, 0xf3, 0xf8, 0xa2, 0xd2, 0x54, 0x41, 0xd0, 0x5e, 0x04, 0x6f, 0xc9, 0x64,
	0xe6, 0xf7, 0x7d, 0xf9, 0x06, 0x16, 0x4e, 0x3e, 0xa2, 0x16, 0xb9, 0xb0, 0x9e, 0x9c, 0xe7, 0x9d,
	0x35, 0xde, 0xb0, 0xb4, 0x35, 0xdc, 0x0d, 0x4e, 0x1a, 0xae, 0xa9, 0xae, 0x15, 0xbe, 0x0a, 0x8b,
	0xbc, 0xb1, 0x9d, 0x5c, 0x9e, 0x37, 0xc6, 0x34, 0x0a, 0xf3, 0xd0, 0x57, 0xf5, 0xfb, 0x1c, 0x75,
	0xe7, 0x87, 0x69, 0x6c, 0x75, 0x0f, 0x49, 0x11, 0x30, 0xec, 0x18, 0x62, 0xaa, 0xd3, 0x28, 0x8b,
	0x2e, 0x0f, 0xb6, 0x31, 0xd5, 0x8c, 0xc1, 0xff, 0x56, 0x68, 0x4c, 0xe3, 0x50, 0x09, 0x67, 0x96,
	0xc1, 0x61, 0x8d, 0x4e, 0x5a, 0xea, 0x3c, 0x99, 0x36, 0xfd, 0x17, 0x9e, 0xbe, 0x96, 0x56, 0x17,
	0xb0, 0xd8, 0x91, 0xee, 0x14, 0x4e, 0xd4, 0x2d, 0x3e, 0xf7, 0xf8, 0x1d, 0x7e, 0xfd, 0x16, 0xc3,
	0xd1, 0xd4, 0xb1, 0x43, 0xfb, 0x42, 0x12, 0xd9, 0x06, 0x92, 0x3b, 0xf4, 0x85, 0x52, 0xec, 0x8c,
	0x4f, 0x86, 0xf9, 0xa7, 0x61, 0x7e, 0x3b, 0x1a, 0x5e, 0x66, 0x7c, 0xee, 0x8b, 0x7c, 0x42, 0xad,
	0x23, 0x86, 0x70, 0x32, 0x52, 0xc2, 0xb5, 0x1c, 0x0a, 0x55, 0xf5, 0x9a, 0x5d, 0xcd, 0xcf, 0xfd,
	0x60, 0xf5, 0x17, 0x32, 0x0f, 0x56, 0xc8, 0xa7, 0x3f, 0x90, 0x29, 0xd7, 0x30, 0xbb, 0xd5, 0xf2,
	0xf4, 0x43, 0x5d, 0x38, 0xdc, 0xe0, 0x9e, 0x5a, 0x1a, 0xe3, 0xaf, 0x92, 0x90, 0xd9, 0xcd, 0x7b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x84, 0xf7, 0x49, 0x23, 0x02, 0x00, 0x00,
}
