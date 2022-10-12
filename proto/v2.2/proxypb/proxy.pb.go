// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

package proxypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/birdwatcher/proto/v2.2/commonpb"
	internalpb "github.com/milvus-io/birdwatcher/proto/v2.2/internalpb"
	milvuspb "github.com/milvus-io/birdwatcher/proto/v2.2/milvuspb"
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

type InvalidateCollMetaCacheRequest struct {
	// MsgType:
	//  DropCollection    ->  {meta cache, dml channels}
	//  Other             ->  {meta cache}
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	CollectionID         int64             `protobuf:"varint,4,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCollMetaCacheRequest) Reset()         { *m = InvalidateCollMetaCacheRequest{} }
func (m *InvalidateCollMetaCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCollMetaCacheRequest) ProtoMessage()    {}
func (*InvalidateCollMetaCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *InvalidateCollMetaCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.Merge(m, src)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Size(m)
}
func (m *InvalidateCollMetaCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCollMetaCacheRequest proto.InternalMessageInfo

func (m *InvalidateCollMetaCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCollMetaCacheRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

type InvalidateCredCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username             string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCredCacheRequest) Reset()         { *m = InvalidateCredCacheRequest{} }
func (m *InvalidateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCredCacheRequest) ProtoMessage()    {}
func (*InvalidateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{1}
}

func (m *InvalidateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCredCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCredCacheRequest.Merge(m, src)
}
func (m *InvalidateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCredCacheRequest.Size(m)
}
func (m *InvalidateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCredCacheRequest proto.InternalMessageInfo

func (m *InvalidateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpdateCredCacheRequest struct {
	Base     *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// password stored in cache
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCredCacheRequest) Reset()         { *m = UpdateCredCacheRequest{} }
func (m *UpdateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCredCacheRequest) ProtoMessage()    {}
func (*UpdateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{2}
}

func (m *UpdateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCredCacheRequest.Unmarshal(m, b)
}
func (m *UpdateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCredCacheRequest.Merge(m, src)
}
func (m *UpdateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCredCacheRequest.Size(m)
}
func (m *UpdateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCredCacheRequest proto.InternalMessageInfo

func (m *UpdateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *UpdateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateCredCacheRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RefreshPolicyInfoCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	OpType               int32             `protobuf:"varint,2,opt,name=opType,proto3" json:"opType,omitempty"`
	OpKey                string            `protobuf:"bytes,3,opt,name=opKey,proto3" json:"opKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RefreshPolicyInfoCacheRequest) Reset()         { *m = RefreshPolicyInfoCacheRequest{} }
func (m *RefreshPolicyInfoCacheRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshPolicyInfoCacheRequest) ProtoMessage()    {}
func (*RefreshPolicyInfoCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{3}
}

func (m *RefreshPolicyInfoCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Unmarshal(m, b)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Marshal(b, m, deterministic)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.Merge(m, src)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Size(m)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshPolicyInfoCacheRequest proto.InternalMessageInfo

func (m *RefreshPolicyInfoCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RefreshPolicyInfoCacheRequest) GetOpType() int32 {
	if m != nil {
		return m.OpType
	}
	return 0
}

func (m *RefreshPolicyInfoCacheRequest) GetOpKey() string {
	if m != nil {
		return m.OpKey
	}
	return ""
}

type SetRatesRequest struct {
	Base                 *commonpb.MsgBase  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Rates                []*internalpb.Rate `protobuf:"bytes,2,rep,name=rates,proto3" json:"rates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SetRatesRequest) Reset()         { *m = SetRatesRequest{} }
func (m *SetRatesRequest) String() string { return proto.CompactTextString(m) }
func (*SetRatesRequest) ProtoMessage()    {}
func (*SetRatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{4}
}

func (m *SetRatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRatesRequest.Unmarshal(m, b)
}
func (m *SetRatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRatesRequest.Marshal(b, m, deterministic)
}
func (m *SetRatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRatesRequest.Merge(m, src)
}
func (m *SetRatesRequest) XXX_Size() int {
	return xxx_messageInfo_SetRatesRequest.Size(m)
}
func (m *SetRatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRatesRequest proto.InternalMessageInfo

func (m *SetRatesRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *SetRatesRequest) GetRates() []*internalpb.Rate {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*InvalidateCollMetaCacheRequest)(nil), "milvus.protov2.proxy.InvalidateCollMetaCacheRequest")
	proto.RegisterType((*InvalidateCredCacheRequest)(nil), "milvus.protov2.proxy.InvalidateCredCacheRequest")
	proto.RegisterType((*UpdateCredCacheRequest)(nil), "milvus.protov2.proxy.UpdateCredCacheRequest")
	proto.RegisterType((*RefreshPolicyInfoCacheRequest)(nil), "milvus.protov2.proxy.RefreshPolicyInfoCacheRequest")
	proto.RegisterType((*SetRatesRequest)(nil), "milvus.protov2.proxy.SetRatesRequest")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6a, 0xdb, 0x30,
	0x18, 0xad, 0xdb, 0xa6, 0xed, 0xbe, 0x86, 0x16, 0x44, 0xd7, 0x65, 0x86, 0x94, 0x60, 0xe8, 0x96,
	0x41, 0x97, 0x14, 0xa7, 0x4f, 0xd0, 0x14, 0x42, 0x18, 0x19, 0xc5, 0xd9, 0x6e, 0x76, 0x33, 0x64,
	0xfb, 0x4b, 0x22, 0xb0, 0x25, 0xcf, 0x92, 0xb3, 0x65, 0x77, 0x83, 0x3d, 0xd5, 0xde, 0x63, 0xef,
	0x33, 0xfc, 0x13, 0x27, 0x71, 0x9d, 0x64, 0x23, 0xb0, 0x3b, 0x1f, 0xe9, 0x7c, 0x3a, 0xe7, 0xd8,
	0x3a, 0x86, 0xd3, 0x20, 0x14, 0xdf, 0x66, 0xad, 0x20, 0x14, 0x4a, 0x90, 0x0b, 0x9f, 0x79, 0xd3,
	0x48, 0xa6, 0x68, 0x6a, 0xb6, 0x92, 0x3d, 0xbd, 0xea, 0x08, 0xdf, 0x17, 0x3c, 0x5d, 0xd5, 0xcf,
	0x18, 0x57, 0x18, 0x72, 0xea, 0x65, 0xb8, 0xba, 0x3c, 0x63, 0xfc, 0xd2, 0xe0, 0xaa, 0xcf, 0xa7,
	0xd4, 0x63, 0x2e, 0x55, 0xd8, 0x15, 0x9e, 0x37, 0x40, 0x45, 0xbb, 0xd4, 0x99, 0xa0, 0x85, 0x5f,
	0x22, 0x94, 0x8a, 0x98, 0x70, 0x68, 0x53, 0x89, 0x35, 0xad, 0xa1, 0x35, 0x4f, 0xcd, 0xab, 0x56,
	0x41, 0x33, 0x13, 0x1b, 0xc8, 0xf1, 0x3d, 0x95, 0x68, 0x25, 0x5c, 0xf2, 0x02, 0x8e, 0x5d, 0xfb,
	0x33, 0xa7, 0x3e, 0xd6, 0xf6, 0x1b, 0x5a, 0xf3, 0x99, 0x75, 0xe4, 0xda, 0xef, 0xa9, 0x8f, 0xe4,
	0x35, 0x9c, 0x3b, 0xc2, 0xf3, 0xd0, 0x51, 0x4c, 0xf0, 0x94, 0x70, 0x90, 0x10, 0xce, 0x16, 0xcb,
	0x09, 0xd1, 0x80, 0xea, 0x62, 0xa5, 0xff, 0x50, 0x3b, 0x6c, 0x68, 0xcd, 0x03, 0x6b, 0x65, 0xcd,
	0xf0, 0x40, 0x5f, 0xf2, 0x1e, 0xa2, 0xbb, 0xb3, 0x6f, 0x1d, 0x4e, 0x22, 0x19, 0xbf, 0xad, 0xdc,
	0x78, 0x8e, 0x8d, 0x9f, 0x1a, 0x5c, 0x7e, 0x0c, 0xfe, 0x87, 0x54, 0xbc, 0x17, 0x50, 0x29, 0xbf,
	0x8a, 0xd0, 0xcd, 0x5e, 0x4f, 0x8e, 0x8d, 0x1f, 0x1a, 0xd4, 0x2d, 0x1c, 0x85, 0x28, 0x27, 0x8f,
	0xc2, 0x63, 0xce, 0xac, 0xcf, 0x47, 0x62, 0x67, 0x37, 0x97, 0x70, 0x24, 0x82, 0x0f, 0xb3, 0x20,
	0xf5, 0x52, 0xb1, 0x32, 0x44, 0x2e, 0xa0, 0x22, 0x82, 0x77, 0x38, 0xcb, 0x6c, 0xa4, 0xc0, 0xf8,
	0x0e, 0xe7, 0x43, 0x54, 0x16, 0x55, 0x28, 0x77, 0x11, 0xed, 0x40, 0x25, 0x8c, 0xcf, 0xa8, 0xed,
	0x37, 0x0e, 0x9a, 0xa7, 0x66, 0xbd, 0x38, 0x94, 0xdf, 0xdc, 0x58, 0xc9, 0x4a, 0xb9, 0xe6, 0xef,
	0x63, 0xa8, 0x3c, 0xc6, 0xf7, 0x9c, 0x04, 0x40, 0x7a, 0xa8, 0xba, 0xc2, 0x0f, 0x04, 0x47, 0xae,
	0x86, 0x2a, 0xde, 0x27, 0xb7, 0xc5, 0x53, 0x32, 0xf8, 0x94, 0x9a, 0x59, 0xd7, 0x5f, 0xad, 0x99,
	0x28, 0xd0, 0x8d, 0x3d, 0x12, 0xc1, 0x45, 0x0f, 0x13, 0xc8, 0xa4, 0x62, 0x8e, 0xec, 0x4e, 0x28,
	0xe7, 0xe8, 0x91, 0xbb, 0xb5, 0xce, 0xcb, 0xe8, 0x73, 0xdd, 0xeb, 0x35, 0xba, 0x43, 0x15, 0x32,
	0x3e, 0xb6, 0x50, 0x06, 0x82, 0x4b, 0x34, 0xf6, 0xc8, 0x14, 0xea, 0xab, 0x1d, 0x4d, 0x1b, 0x90,
	0x37, 0xf5, 0xa9, 0x7e, 0xfa, 0x93, 0xd8, 0x5c, 0x6c, 0xbd, 0xbe, 0xe6, 0x23, 0xc5, 0x86, 0xa3,
	0x38, 0x2e, 0x42, 0xb5, 0x87, 0xea, 0xc1, 0x9d, 0xc7, 0xbc, 0xd9, 0x14, 0x33, 0xa7, 0xfd, 0x73,
	0x3c, 0x0e, 0x2f, 0x57, 0x6b, 0x8c, 0x5c, 0x31, 0xea, 0xa5, 0xd1, 0x6e, 0xb7, 0x46, 0x2b, 0x94,
	0x71, 0x7b, 0xac, 0x11, 0x3c, 0x5f, 0xf4, 0x78, 0x59, 0xeb, 0xa6, 0x5c, 0xab, 0xbc, 0xf4, 0xdb,
	0x75, 0x38, 0x5c, 0x96, 0x17, 0x95, 0x74, 0xca, 0x85, 0x36, 0xd6, 0x7a, 0xbb, 0xde, 0x04, 0xce,
	0x7b, 0xa8, 0x92, 0x6e, 0x0c, 0x50, 0x85, 0xcc, 0x91, 0xa4, 0xb9, 0xbe, 0x0c, 0x19, 0x65, 0x7e,
	0xfa, 0x9b, 0xbf, 0x60, 0xe6, 0x5f, 0xcc, 0x82, 0x93, 0x79, 0xff, 0xc9, 0x75, 0x79, 0x96, 0xc2,
	0xff, 0x61, 0xab, 0xfb, 0xfb, 0xbb, 0x4f, 0xe6, 0x98, 0xa9, 0x49, 0x64, 0xc7, 0x3b, 0xed, 0x94,
	0xfc, 0x96, 0x89, 0xec, 0xa9, 0x3d, 0xbf, 0x6a, 0xed, 0x64, 0xbe, 0x9d, 0x88, 0x04, 0xb6, 0x7d,
	0x94, 0xc0, 0xce, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xa4, 0x20, 0xb3, 0x17, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	GetComponentStates(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption) (*milvuspb.ComponentStates, error)
	GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	GetProxyMetrics(ctx context.Context, in *milvuspb.GetMetricsRequest, opts ...grpc.CallOption) (*milvuspb.GetMetricsResponse, error)
	SetRates(ctx context.Context, in *SetRatesRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetComponentStates(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption) (*milvuspb.ComponentStates, error) {
	out := new(milvuspb.ComponentStates)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/GetComponentStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/GetStatisticsChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/InvalidateCollectionMetaCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/GetDdChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/InvalidateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/UpdateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/RefreshPolicyInfoCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetProxyMetrics(ctx context.Context, in *milvuspb.GetMetricsRequest, opts ...grpc.CallOption) (*milvuspb.GetMetricsResponse, error) {
	out := new(milvuspb.GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/GetProxyMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) SetRates(ctx context.Context, in *SetRatesRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.protov2.proxy.Proxy/SetRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	GetComponentStates(context.Context, *milvuspb.GetComponentStatesRequest) (*milvuspb.ComponentStates, error)
	GetStatisticsChannel(context.Context, *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(context.Context, *InvalidateCollMetaCacheRequest) (*commonpb.Status, error)
	GetDdChannel(context.Context, *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error)
	InvalidateCredentialCache(context.Context, *InvalidateCredCacheRequest) (*commonpb.Status, error)
	UpdateCredentialCache(context.Context, *UpdateCredCacheRequest) (*commonpb.Status, error)
	RefreshPolicyInfoCache(context.Context, *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error)
	GetProxyMetrics(context.Context, *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error)
	SetRates(context.Context, *SetRatesRequest) (*commonpb.Status, error)
}

// UnimplementedProxyServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (*UnimplementedProxyServer) GetComponentStates(ctx context.Context, req *milvuspb.GetComponentStatesRequest) (*milvuspb.ComponentStates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentStates not implemented")
}
func (*UnimplementedProxyServer) GetStatisticsChannel(ctx context.Context, req *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatisticsChannel not implemented")
}
func (*UnimplementedProxyServer) InvalidateCollectionMetaCache(ctx context.Context, req *InvalidateCollMetaCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCollectionMetaCache not implemented")
}
func (*UnimplementedProxyServer) GetDdChannel(ctx context.Context, req *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDdChannel not implemented")
}
func (*UnimplementedProxyServer) InvalidateCredentialCache(ctx context.Context, req *InvalidateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) UpdateCredentialCache(ctx context.Context, req *UpdateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) RefreshPolicyInfoCache(ctx context.Context, req *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshPolicyInfoCache not implemented")
}
func (*UnimplementedProxyServer) GetProxyMetrics(ctx context.Context, req *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProxyMetrics not implemented")
}
func (*UnimplementedProxyServer) SetRates(ctx context.Context, req *SetRatesRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRates not implemented")
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_GetComponentStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.GetComponentStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetComponentStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/GetComponentStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetComponentStates(ctx, req.(*milvuspb.GetComponentStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetStatisticsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetStatisticsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/GetStatisticsChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, req.(*internalpb.GetStatisticsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCollectionMetaCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCollMetaCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/InvalidateCollectionMetaCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, req.(*InvalidateCollMetaCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetDdChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetDdChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetDdChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/GetDdChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetDdChannel(ctx, req.(*internalpb.GetDdChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/InvalidateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, req.(*InvalidateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_UpdateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/UpdateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, req.(*UpdateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_RefreshPolicyInfoCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshPolicyInfoCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/RefreshPolicyInfoCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, req.(*RefreshPolicyInfoCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetProxyMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetProxyMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/GetProxyMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetProxyMetrics(ctx, req.(*milvuspb.GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_SetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).SetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.protov2.proxy.Proxy/SetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).SetRates(ctx, req.(*SetRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.protov2.proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComponentStates",
			Handler:    _Proxy_GetComponentStates_Handler,
		},
		{
			MethodName: "GetStatisticsChannel",
			Handler:    _Proxy_GetStatisticsChannel_Handler,
		},
		{
			MethodName: "InvalidateCollectionMetaCache",
			Handler:    _Proxy_InvalidateCollectionMetaCache_Handler,
		},
		{
			MethodName: "GetDdChannel",
			Handler:    _Proxy_GetDdChannel_Handler,
		},
		{
			MethodName: "InvalidateCredentialCache",
			Handler:    _Proxy_InvalidateCredentialCache_Handler,
		},
		{
			MethodName: "UpdateCredentialCache",
			Handler:    _Proxy_UpdateCredentialCache_Handler,
		},
		{
			MethodName: "RefreshPolicyInfoCache",
			Handler:    _Proxy_RefreshPolicyInfoCache_Handler,
		},
		{
			MethodName: "GetProxyMetrics",
			Handler:    _Proxy_GetProxyMetrics_Handler,
		},
		{
			MethodName: "SetRates",
			Handler:    _Proxy_SetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}