// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osd.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import multiraftbase "github.com/journeymidnight/nentropy/multiraft/multiraftbase"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SyncMapType int32

const (
	// CONSISTENT reads are guaranteed to read committed data; the
	// mechanism relies on clocks to determine lease expirations.
	PGMAP SyncMapType = 0
	// CONSENSUS requires that reads must achieve consensus. This is a
	// stronger guarantee of consistency than CONSISTENT.
	//
	// TODO(spencer): current unimplemented.
	OSDMAP SyncMapType = 1
	// INCONSISTENT reads return the latest available, committed values.
	// They are more efficient, but may read stale values as pending
	// intents are ignored.
	POOLMAP SyncMapType = 2
)

var SyncMapType_name = map[int32]string{
	0: "PGMAP",
	1: "OSDMAP",
	2: "POOLMAP",
}
var SyncMapType_value = map[string]int32{
	"PGMAP":   0,
	"OSDMAP":  1,
	"POOLMAP": 2,
}

func (x SyncMapType) String() string {
	return proto.EnumName(SyncMapType_name, int32(x))
}
func (SyncMapType) EnumDescriptor() ([]byte, []int) { return fileDescriptorOsd, []int{0} }

type UnionMap struct {
	Pgmap   *PgMaps  `protobuf:"bytes,1,opt,name=pgmap" json:"pgmap,omitempty"`
	Poolmap *PoolMap `protobuf:"bytes,2,opt,name=poolmap" json:"poolmap,omitempty"`
	Osdmap  *OsdMap  `protobuf:"bytes,3,opt,name=osdmap" json:"osdmap,omitempty"`
}

func (m *UnionMap) Reset()                    { *m = UnionMap{} }
func (m *UnionMap) String() string            { return proto.CompactTextString(m) }
func (*UnionMap) ProtoMessage()               {}
func (*UnionMap) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{0} }

func (m *UnionMap) GetPgmap() *PgMaps {
	if m != nil {
		return m.Pgmap
	}
	return nil
}

func (m *UnionMap) GetPoolmap() *PoolMap {
	if m != nil {
		return m.Poolmap
	}
	return nil
}

func (m *UnionMap) GetOsdmap() *OsdMap {
	if m != nil {
		return m.Osdmap
	}
	return nil
}

type SyncMapRequest struct {
	MapType  SyncMapType `protobuf:"varint,1,opt,name=map_type,json=mapType,proto3,enum=protos.SyncMapType" json:"map_type,omitempty"`
	UnionMap UnionMap    `protobuf:"bytes,2,opt,name=union_map,json=unionMap" json:"union_map"`
}

func (m *SyncMapRequest) Reset()                    { *m = SyncMapRequest{} }
func (m *SyncMapRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncMapRequest) ProtoMessage()               {}
func (*SyncMapRequest) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{1} }

func (m *SyncMapRequest) GetMapType() SyncMapType {
	if m != nil {
		return m.MapType
	}
	return PGMAP
}

func (m *SyncMapRequest) GetUnionMap() UnionMap {
	if m != nil {
		return m.UnionMap
	}
	return UnionMap{}
}

type SyncMapReply struct {
	RetCode int32 `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
}

func (m *SyncMapReply) Reset()                    { *m = SyncMapReply{} }
func (m *SyncMapReply) String() string            { return proto.CompactTextString(m) }
func (*SyncMapReply) ProtoMessage()               {}
func (*SyncMapReply) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{2} }

func (m *SyncMapReply) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

type CreatePgRequest struct {
	PgName          string                         `protobuf:"bytes,1,opt,name=pg_name,json=pgName,proto3" json:"pg_name,omitempty"`
	GroupDescriptor *multiraftbase.GroupDescriptor `protobuf:"bytes,2,opt,name=group_descriptor,json=groupDescriptor" json:"group_descriptor,omitempty"`
}

func (m *CreatePgRequest) Reset()                    { *m = CreatePgRequest{} }
func (m *CreatePgRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePgRequest) ProtoMessage()               {}
func (*CreatePgRequest) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{3} }

func (m *CreatePgRequest) GetPgName() string {
	if m != nil {
		return m.PgName
	}
	return ""
}

func (m *CreatePgRequest) GetGroupDescriptor() *multiraftbase.GroupDescriptor {
	if m != nil {
		return m.GroupDescriptor
	}
	return nil
}

type CreatePgReply struct {
	RetCode int32 `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
}

func (m *CreatePgReply) Reset()                    { *m = CreatePgReply{} }
func (m *CreatePgReply) String() string            { return proto.CompactTextString(m) }
func (*CreatePgReply) ProtoMessage()               {}
func (*CreatePgReply) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{4} }

func (m *CreatePgReply) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

type DeletePgRequest struct {
	PgName string `protobuf:"bytes,1,opt,name=pg_name,json=pgName,proto3" json:"pg_name,omitempty"`
}

func (m *DeletePgRequest) Reset()                    { *m = DeletePgRequest{} }
func (m *DeletePgRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePgRequest) ProtoMessage()               {}
func (*DeletePgRequest) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{5} }

func (m *DeletePgRequest) GetPgName() string {
	if m != nil {
		return m.PgName
	}
	return ""
}

type DeletePgReply struct {
	RetCode int32 `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
}

func (m *DeletePgReply) Reset()                    { *m = DeletePgReply{} }
func (m *DeletePgReply) String() string            { return proto.CompactTextString(m) }
func (*DeletePgReply) ProtoMessage()               {}
func (*DeletePgReply) Descriptor() ([]byte, []int) { return fileDescriptorOsd, []int{6} }

func (m *DeletePgReply) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func init() {
	proto.RegisterType((*UnionMap)(nil), "protos.UnionMap")
	proto.RegisterType((*SyncMapRequest)(nil), "protos.SyncMapRequest")
	proto.RegisterType((*SyncMapReply)(nil), "protos.SyncMapReply")
	proto.RegisterType((*CreatePgRequest)(nil), "protos.CreatePgRequest")
	proto.RegisterType((*CreatePgReply)(nil), "protos.CreatePgReply")
	proto.RegisterType((*DeletePgRequest)(nil), "protos.DeletePgRequest")
	proto.RegisterType((*DeletePgReply)(nil), "protos.DeletePgReply")
	proto.RegisterEnum("protos.SyncMapType", SyncMapType_name, SyncMapType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OsdRpc service

type OsdRpcClient interface {
	CreatePg(ctx context.Context, in *CreatePgRequest, opts ...grpc.CallOption) (*CreatePgReply, error)
	DeletePg(ctx context.Context, in *DeletePgRequest, opts ...grpc.CallOption) (*DeletePgReply, error)
	SyncMap(ctx context.Context, in *SyncMapRequest, opts ...grpc.CallOption) (*SyncMapReply, error)
}

type osdRpcClient struct {
	cc *grpc.ClientConn
}

func NewOsdRpcClient(cc *grpc.ClientConn) OsdRpcClient {
	return &osdRpcClient{cc}
}

func (c *osdRpcClient) CreatePg(ctx context.Context, in *CreatePgRequest, opts ...grpc.CallOption) (*CreatePgReply, error) {
	out := new(CreatePgReply)
	err := grpc.Invoke(ctx, "/protos.OsdRpc/CreatePg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osdRpcClient) DeletePg(ctx context.Context, in *DeletePgRequest, opts ...grpc.CallOption) (*DeletePgReply, error) {
	out := new(DeletePgReply)
	err := grpc.Invoke(ctx, "/protos.OsdRpc/DeletePg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osdRpcClient) SyncMap(ctx context.Context, in *SyncMapRequest, opts ...grpc.CallOption) (*SyncMapReply, error) {
	out := new(SyncMapReply)
	err := grpc.Invoke(ctx, "/protos.OsdRpc/SyncMap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OsdRpc service

type OsdRpcServer interface {
	CreatePg(context.Context, *CreatePgRequest) (*CreatePgReply, error)
	DeletePg(context.Context, *DeletePgRequest) (*DeletePgReply, error)
	SyncMap(context.Context, *SyncMapRequest) (*SyncMapReply, error)
}

func RegisterOsdRpcServer(s *grpc.Server, srv OsdRpcServer) {
	s.RegisterService(&_OsdRpc_serviceDesc, srv)
}

func _OsdRpc_CreatePg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsdRpcServer).CreatePg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OsdRpc/CreatePg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsdRpcServer).CreatePg(ctx, req.(*CreatePgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsdRpc_DeletePg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsdRpcServer).DeletePg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OsdRpc/DeletePg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsdRpcServer).DeletePg(ctx, req.(*DeletePgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsdRpc_SyncMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsdRpcServer).SyncMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OsdRpc/SyncMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsdRpcServer).SyncMap(ctx, req.(*SyncMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OsdRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.OsdRpc",
	HandlerType: (*OsdRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePg",
			Handler:    _OsdRpc_CreatePg_Handler,
		},
		{
			MethodName: "DeletePg",
			Handler:    _OsdRpc_DeletePg_Handler,
		},
		{
			MethodName: "SyncMap",
			Handler:    _OsdRpc_SyncMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osd.proto",
}

func (m *UnionMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnionMap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pgmap != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.Pgmap.Size()))
		n1, err := m.Pgmap.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Poolmap != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.Poolmap.Size()))
		n2, err := m.Poolmap.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Osdmap != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.Osdmap.Size()))
		n3, err := m.Osdmap.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SyncMapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncMapRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MapType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.MapType))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintOsd(dAtA, i, uint64(m.UnionMap.Size()))
	n4, err := m.UnionMap.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *SyncMapReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncMapReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RetCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.RetCode))
	}
	return i, nil
}

func (m *CreatePgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePgRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PgName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOsd(dAtA, i, uint64(len(m.PgName)))
		i += copy(dAtA[i:], m.PgName)
	}
	if m.GroupDescriptor != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.GroupDescriptor.Size()))
		n5, err := m.GroupDescriptor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *CreatePgReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePgReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RetCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.RetCode))
	}
	return i, nil
}

func (m *DeletePgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletePgRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PgName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOsd(dAtA, i, uint64(len(m.PgName)))
		i += copy(dAtA[i:], m.PgName)
	}
	return i, nil
}

func (m *DeletePgReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletePgReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RetCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOsd(dAtA, i, uint64(m.RetCode))
	}
	return i, nil
}

func encodeFixed64Osd(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Osd(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintOsd(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UnionMap) Size() (n int) {
	var l int
	_ = l
	if m.Pgmap != nil {
		l = m.Pgmap.Size()
		n += 1 + l + sovOsd(uint64(l))
	}
	if m.Poolmap != nil {
		l = m.Poolmap.Size()
		n += 1 + l + sovOsd(uint64(l))
	}
	if m.Osdmap != nil {
		l = m.Osdmap.Size()
		n += 1 + l + sovOsd(uint64(l))
	}
	return n
}

func (m *SyncMapRequest) Size() (n int) {
	var l int
	_ = l
	if m.MapType != 0 {
		n += 1 + sovOsd(uint64(m.MapType))
	}
	l = m.UnionMap.Size()
	n += 1 + l + sovOsd(uint64(l))
	return n
}

func (m *SyncMapReply) Size() (n int) {
	var l int
	_ = l
	if m.RetCode != 0 {
		n += 1 + sovOsd(uint64(m.RetCode))
	}
	return n
}

func (m *CreatePgRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.PgName)
	if l > 0 {
		n += 1 + l + sovOsd(uint64(l))
	}
	if m.GroupDescriptor != nil {
		l = m.GroupDescriptor.Size()
		n += 1 + l + sovOsd(uint64(l))
	}
	return n
}

func (m *CreatePgReply) Size() (n int) {
	var l int
	_ = l
	if m.RetCode != 0 {
		n += 1 + sovOsd(uint64(m.RetCode))
	}
	return n
}

func (m *DeletePgRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.PgName)
	if l > 0 {
		n += 1 + l + sovOsd(uint64(l))
	}
	return n
}

func (m *DeletePgReply) Size() (n int) {
	var l int
	_ = l
	if m.RetCode != 0 {
		n += 1 + sovOsd(uint64(m.RetCode))
	}
	return n
}

func sovOsd(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOsd(x uint64) (n int) {
	return sovOsd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UnionMap) GetValue() interface{} {
	if this.Pgmap != nil {
		return this.Pgmap
	}
	if this.Poolmap != nil {
		return this.Poolmap
	}
	if this.Osdmap != nil {
		return this.Osdmap
	}
	return nil
}

func (this *UnionMap) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *PgMaps:
		this.Pgmap = vt
	case *PoolMap:
		this.Poolmap = vt
	case *OsdMap:
		this.Osdmap = vt
	default:
		return false
	}
	return true
}
func (m *UnionMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnionMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnionMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pgmap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pgmap == nil {
				m.Pgmap = &PgMaps{}
			}
			if err := m.Pgmap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Poolmap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Poolmap == nil {
				m.Poolmap = &PoolMap{}
			}
			if err := m.Poolmap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Osdmap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Osdmap == nil {
				m.Osdmap = &OsdMap{}
			}
			if err := m.Osdmap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SyncMapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncMapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncMapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapType", wireType)
			}
			m.MapType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MapType |= (SyncMapType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnionMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnionMap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SyncMapReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncMapReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncMapReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetCode", wireType)
			}
			m.RetCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreatePgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupDescriptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GroupDescriptor == nil {
				m.GroupDescriptor = &multiraftbase.GroupDescriptor{}
			}
			if err := m.GroupDescriptor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreatePgReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePgReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePgReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetCode", wireType)
			}
			m.RetCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeletePgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeletePgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletePgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOsd
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeletePgReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeletePgReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletePgReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetCode", wireType)
			}
			m.RetCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOsd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOsd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOsd(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOsd
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOsd
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthOsd
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOsd
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOsd(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOsd = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOsd   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("osd.proto", fileDescriptorOsd) }

var fileDescriptorOsd = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xd3, 0x3c,
	0x14, 0x85, 0x93, 0xf9, 0xdb, 0xa6, 0xbd, 0xfd, 0x69, 0x2b, 0x33, 0x30, 0xd0, 0x45, 0x06, 0x45,
	0x08, 0x31, 0x5d, 0xa4, 0x52, 0x67, 0x81, 0x40, 0x6c, 0x98, 0xa9, 0x34, 0x02, 0x51, 0x5a, 0x65,
	0x60, 0x5d, 0xb9, 0x8d, 0xf1, 0x04, 0xc5, 0xb1, 0xb1, 0x9d, 0x45, 0xde, 0x00, 0xde, 0x80, 0x25,
	0x12, 0x0f, 0xc2, 0xb6, 0x4b, 0x9e, 0x00, 0xa1, 0xf2, 0x22, 0xc8, 0x4d, 0xd2, 0xd0, 0x8e, 0x04,
	0xac, 0x7a, 0xec, 0xfb, 0x9d, 0xeb, 0x73, 0xed, 0x06, 0x5a, 0x5c, 0x85, 0xbe, 0x90, 0x5c, 0x73,
	0xd4, 0xd8, 0xfc, 0xa8, 0xfe, 0x0b, 0x1a, 0xe9, 0xab, 0x74, 0xe1, 0x2f, 0x39, 0x1b, 0xbe, 0xe3,
	0xa9, 0x4c, 0x48, 0xc6, 0xa2, 0x30, 0x89, 0xe8, 0x95, 0x1e, 0x26, 0x24, 0xd1, 0x92, 0x8b, 0x6c,
	0xc8, 0xd2, 0x58, 0x47, 0x12, 0xbf, 0xd5, 0x95, 0x5a, 0x60, 0x45, 0x86, 0x8c, 0x68, 0x1c, 0x62,
	0x8d, 0xf3, 0x9e, 0xfd, 0x16, 0xe3, 0x49, 0x21, 0x0f, 0x29, 0xa7, 0x7c, 0x23, 0x87, 0x46, 0xe5,
	0xbb, 0xde, 0x47, 0x1b, 0x9a, 0x6f, 0x92, 0x88, 0x27, 0x13, 0x2c, 0xd0, 0x7d, 0xa8, 0x0b, 0xca,
	0xb0, 0xb8, 0x63, 0xdf, 0xb3, 0x1f, 0xb6, 0x47, 0x9d, 0x9c, 0x51, 0xfe, 0x8c, 0x4e, 0xb0, 0x50,
	0x41, 0x5e, 0x44, 0x27, 0xe0, 0x08, 0xce, 0x63, 0xc3, 0x1d, 0x6c, 0xb8, 0xee, 0x96, 0xe3, 0x3c,
	0x9e, 0x60, 0x11, 0x94, 0x75, 0xf4, 0x00, 0x1a, 0x5c, 0x85, 0x86, 0xfc, 0x6f, 0xb7, 0xe3, 0x54,
	0x85, 0x06, 0x2c, 0xaa, 0x4f, 0x6a, 0xab, 0xcf, 0xc7, 0xb6, 0x97, 0x42, 0xe7, 0x32, 0x4b, 0x96,
	0xa6, 0x40, 0xde, 0xa7, 0x44, 0x69, 0xe4, 0x43, 0x93, 0x61, 0x31, 0xd7, 0x99, 0x20, 0x9b, 0x4c,
	0x9d, 0xd1, 0xcd, 0xb2, 0x43, 0x41, 0xbe, 0xce, 0x04, 0x09, 0x1c, 0x96, 0x0b, 0x74, 0x0a, 0xad,
	0xd4, 0x0c, 0x33, 0xaf, 0xc2, 0xf5, 0x4a, 0x43, 0x39, 0xe5, 0x59, 0x6d, 0xf5, 0xfd, 0xd8, 0x0a,
	0x9a, 0x69, 0xb1, 0xf6, 0x4e, 0xe0, 0xff, 0xed, 0xb1, 0x22, 0xce, 0xd0, 0x5d, 0x68, 0x4a, 0xa2,
	0xe7, 0x4b, 0x1e, 0xe6, 0x87, 0xd6, 0x03, 0x47, 0x12, 0x7d, 0xce, 0x43, 0xe2, 0xa5, 0xd0, 0x3d,
	0x97, 0x04, 0x6b, 0x32, 0xa3, 0x65, 0xc4, 0x23, 0x70, 0x04, 0x9d, 0x27, 0x98, 0xe5, 0x70, 0x2b,
	0x68, 0x08, 0xfa, 0x0a, 0x33, 0x82, 0x9e, 0x43, 0x8f, 0x4a, 0x9e, 0x8a, 0x79, 0x48, 0xd4, 0x52,
	0x46, 0x42, 0x73, 0x59, 0x44, 0x72, 0xfd, 0x9d, 0x37, 0xf3, 0x2f, 0x0c, 0x36, 0xde, 0x52, 0x41,
	0x97, 0xee, 0x6e, 0x78, 0x03, 0xb8, 0x51, 0x1d, 0xfb, 0x97, 0x88, 0x03, 0xe8, 0x8e, 0x49, 0x4c,
	0xfe, 0x25, 0xa2, 0xe9, 0x5b, 0xb1, 0x7f, 0xee, 0x3b, 0x78, 0x04, 0xed, 0xdf, 0xae, 0x1c, 0xb5,
	0xa0, 0x3e, 0xbb, 0x98, 0x3c, 0x9b, 0xf5, 0x2c, 0x04, 0xd0, 0x98, 0x5e, 0x8e, 0x8d, 0xb6, 0x51,
	0x1b, 0x9c, 0xd9, 0x74, 0xfa, 0xd2, 0x2c, 0x0e, 0xfa, 0xb5, 0x0f, 0x5f, 0x5c, 0x6b, 0xf4, 0xd5,
	0x86, 0xc6, 0x54, 0x85, 0x81, 0x58, 0xa2, 0xa7, 0xd0, 0x2c, 0xe7, 0x40, 0x47, 0xe5, 0xbb, 0xec,
	0x5d, 0x68, 0xff, 0xd6, 0xf5, 0x82, 0x88, 0x33, 0xcf, 0x32, 0xee, 0x32, 0x6d, 0xe5, 0xde, 0x9b,
	0xb5, 0x72, 0xef, 0x0c, 0xe6, 0x59, 0xe8, 0x31, 0x38, 0x45, 0x7e, 0x74, 0x7b, 0xef, 0x3f, 0x54,
	0x7a, 0x0f, 0xaf, 0xed, 0x6f, 0xac, 0x67, 0xbd, 0xd5, 0xda, 0xb5, 0xbf, 0xad, 0x5d, 0xfb, 0xc7,
	0xda, 0xb5, 0x3f, 0xfd, 0x74, 0xad, 0x45, 0xfe, 0xa9, 0x9e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x27, 0xf4, 0xf9, 0xe3, 0xbe, 0x03, 0x00, 0x00,
}
