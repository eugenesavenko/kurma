// Code generated by protoc-gen-go.
// source: stage1/client/init.proto
// DO NOT EDIT!

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	stage1/client/init.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	ContainerRequest
	ListResponse
	ByteChunk
	Container
	None
*/
package client

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Container_State int32

const (
	Container_NEW      Container_State = 0
	Container_STARTING Container_State = 1
	Container_RUNNING  Container_State = 2
	Container_STOPPING Container_State = 3
	Container_STOPPED  Container_State = 4
	Container_EXITED   Container_State = 5
)

var Container_State_name = map[int32]string{
	0: "NEW",
	1: "STARTING",
	2: "RUNNING",
	3: "STOPPING",
	4: "STOPPED",
	5: "EXITED",
}
var Container_State_value = map[string]int32{
	"NEW":      0,
	"STARTING": 1,
	"RUNNING":  2,
	"STOPPING": 3,
	"STOPPED":  4,
	"EXITED":   5,
}

func (x Container_State) String() string {
	return proto.EnumName(Container_State_name, int32(x))
}

type CreateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Manifest []byte `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}

type CreateResponse struct {
	ImageUploadId string     `protobuf:"bytes,1,opt,name=image_upload_id" json:"image_upload_id,omitempty"`
	Container     *Container `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}

func (m *CreateResponse) GetContainer() *Container {
	if m != nil {
		return m.Container
	}
	return nil
}

type ContainerRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *ContainerRequest) Reset()         { *m = ContainerRequest{} }
func (m *ContainerRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerRequest) ProtoMessage()    {}

type ListResponse struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}

func (m *ListResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type ByteChunk struct {
	StreamId string `protobuf:"bytes,1,opt,name=stream_id" json:"stream_id,omitempty"`
	Bytes    []byte `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (m *ByteChunk) Reset()         { *m = ByteChunk{} }
func (m *ByteChunk) String() string { return proto.CompactTextString(m) }
func (*ByteChunk) ProtoMessage()    {}

type Container struct {
	Uuid     string          `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Manifest []byte          `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
	State    Container_State `protobuf:"varint,3,opt,name=state,enum=client.Container_State" json:"state,omitempty"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}

type None struct {
}

func (m *None) Reset()         { *m = None{} }
func (m *None) String() string { return proto.CompactTextString(m) }
func (*None) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("client.Container_State", Container_State_name, Container_State_value)
}

// Client API for Kurma service

type KurmaClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (Kurma_UploadImageClient, error)
	Destroy(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*None, error)
	List(ctx context.Context, in *None, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*Container, error)
}

type kurmaClient struct {
	cc *grpc.ClientConn
}

func NewKurmaClient(cc *grpc.ClientConn) KurmaClient {
	return &kurmaClient{cc}
}

func (c *kurmaClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/client.Kurma/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurmaClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (Kurma_UploadImageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Kurma_serviceDesc.Streams[0], c.cc, "/client.Kurma/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &kurmaUploadImageClient{stream}
	return x, nil
}

type Kurma_UploadImageClient interface {
	Send(*ByteChunk) error
	CloseAndRecv() (*None, error)
	grpc.ClientStream
}

type kurmaUploadImageClient struct {
	grpc.ClientStream
}

func (x *kurmaUploadImageClient) Send(m *ByteChunk) error {
	return x.ClientStream.SendProto(m)
}

func (x *kurmaUploadImageClient) CloseAndRecv() (*None, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(None)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kurmaClient) Destroy(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := grpc.Invoke(ctx, "/client.Kurma/Destroy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurmaClient) List(ctx context.Context, in *None, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/client.Kurma/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurmaClient) Get(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := grpc.Invoke(ctx, "/client.Kurma/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Kurma service

type KurmaServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	UploadImage(Kurma_UploadImageServer) error
	Destroy(context.Context, *ContainerRequest) (*None, error)
	List(context.Context, *None) (*ListResponse, error)
	Get(context.Context, *ContainerRequest) (*Container, error)
}

func RegisterKurmaServer(s *grpc.Server, srv KurmaServer) {
	s.RegisterService(&_Kurma_serviceDesc, srv)
}

func _Kurma_Create_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(CreateRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(KurmaServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Kurma_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KurmaServer).UploadImage(&kurmaUploadImageServer{stream})
}

type Kurma_UploadImageServer interface {
	SendAndClose(*None) error
	Recv() (*ByteChunk, error)
	grpc.ServerStream
}

type kurmaUploadImageServer struct {
	grpc.ServerStream
}

func (x *kurmaUploadImageServer) SendAndClose(m *None) error {
	return x.ServerStream.SendProto(m)
}

func (x *kurmaUploadImageServer) Recv() (*ByteChunk, error) {
	m := new(ByteChunk)
	if err := x.ServerStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Kurma_Destroy_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(ContainerRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(KurmaServer).Destroy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Kurma_List_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(None)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(KurmaServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Kurma_Get_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(ContainerRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(KurmaServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Kurma_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.Kurma",
	HandlerType: (*KurmaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Kurma_Create_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Kurma_Destroy_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Kurma_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Kurma_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _Kurma_UploadImage_Handler,
			ClientStreams: true,
		},
	},
}
