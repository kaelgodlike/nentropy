// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal_raft.proto

package multiraftbase

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RaftTombstone contains information about a replica that has been deleted.
type RaftTombstone struct {
	NextReplicaID ReplicaID `protobuf:"varint,1,opt,name=next_replica_id,json=nextReplicaId,proto3,casttype=ReplicaID" json:"next_replica_id,omitempty"`
}

func (m *RaftTombstone) Reset()                    { *m = RaftTombstone{} }
func (m *RaftTombstone) String() string            { return proto.CompactTextString(m) }
func (*RaftTombstone) ProtoMessage()               {}
func (*RaftTombstone) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{0} }

func (m *RaftTombstone) GetNextReplicaID() ReplicaID {
	if m != nil {
		return m.NextReplicaID
	}
	return 0
}

// RaftSnapshotData is the payload of a raftpb.Snapshot. It contains a raw copy of
// all of the range's data and metadata, including the raft log, sequence cache, etc.
type RaftSnapshotData struct {
	// The latest RangeDescriptor
	PgDescriptor *PgDescriptor                `protobuf:"bytes,1,opt,name=pg_descriptor,json=pgDescriptor" json:"pg_descriptor,omitempty"`
	KV           []*RaftSnapshotData_KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
	// These are really raftpb.Entry, but we model them as raw bytes to avoid
	// roundtripping through memory.
	LogEntries [][]byte `protobuf:"bytes,3,rep,name=log_entries,json=logEntries" json:"log_entries,omitempty"`
}

func (m *RaftSnapshotData) Reset()                    { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string            { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()               {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) { return fileDescriptorInternalRaft, []int{1} }

func (m *RaftSnapshotData) GetPgDescriptor() *PgDescriptor {
	if m != nil {
		return m.PgDescriptor
	}
	return nil
}

func (m *RaftSnapshotData) GetKV() []*RaftSnapshotData_KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *RaftSnapshotData) GetLogEntries() [][]byte {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

type RaftSnapshotData_KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RaftSnapshotData_KeyValue) Reset()         { *m = RaftSnapshotData_KeyValue{} }
func (m *RaftSnapshotData_KeyValue) String() string { return proto.CompactTextString(m) }
func (*RaftSnapshotData_KeyValue) ProtoMessage()    {}
func (*RaftSnapshotData_KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptorInternalRaft, []int{1, 0}
}

func (m *RaftSnapshotData_KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RaftSnapshotData_KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftTombstone)(nil), "multiraftbase.RaftTombstone")
	proto.RegisterType((*RaftSnapshotData)(nil), "multiraftbase.RaftSnapshotData")
	proto.RegisterType((*RaftSnapshotData_KeyValue)(nil), "multiraftbase.RaftSnapshotData.KeyValue")
}
func (m *RaftTombstone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftTombstone) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NextReplicaID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInternalRaft(dAtA, i, uint64(m.NextReplicaID))
	}
	return i, nil
}

func (m *RaftSnapshotData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftSnapshotData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PgDescriptor != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternalRaft(dAtA, i, uint64(m.PgDescriptor.Size()))
		n1, err := m.PgDescriptor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.KV) > 0 {
		for _, msg := range m.KV {
			dAtA[i] = 0x12
			i++
			i = encodeVarintInternalRaft(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.LogEntries) > 0 {
		for _, b := range m.LogEntries {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintInternalRaft(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *RaftSnapshotData_KeyValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftSnapshotData_KeyValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternalRaft(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInternalRaft(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeFixed64InternalRaft(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32InternalRaft(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInternalRaft(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RaftTombstone) Size() (n int) {
	var l int
	_ = l
	if m.NextReplicaID != 0 {
		n += 1 + sovInternalRaft(uint64(m.NextReplicaID))
	}
	return n
}

func (m *RaftSnapshotData) Size() (n int) {
	var l int
	_ = l
	if m.PgDescriptor != nil {
		l = m.PgDescriptor.Size()
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	if len(m.KV) > 0 {
		for _, e := range m.KV {
			l = e.Size()
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	if len(m.LogEntries) > 0 {
		for _, b := range m.LogEntries {
			l = len(b)
			n += 1 + l + sovInternalRaft(uint64(l))
		}
	}
	return n
}

func (m *RaftSnapshotData_KeyValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInternalRaft(uint64(l))
	}
	return n
}

func sovInternalRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternalRaft(x uint64) (n int) {
	return sovInternalRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftTombstone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftTombstone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftTombstone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextReplicaID", wireType)
			}
			m.NextReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextReplicaID |= (ReplicaID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: RaftSnapshotData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftSnapshotData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PgDescriptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PgDescriptor == nil {
				m.PgDescriptor = &PgDescriptor{}
			}
			if err := m.PgDescriptor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KV", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
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
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KV = append(m.KV, &RaftSnapshotData_KeyValue{})
			if err := m.KV[len(m.KV)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogEntries", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogEntries = append(m.LogEntries, make([]byte, postIndex-iNdEx))
			copy(m.LogEntries[len(m.LogEntries)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func (m *RaftSnapshotData_KeyValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalRaft
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
			return fmt.Errorf("proto: KeyValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternalRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternalRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalRaft
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
func skipInternalRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
					return 0, ErrIntOverflowInternalRaft
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
				return 0, ErrInvalidLengthInternalRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternalRaft
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
				next, err := skipInternalRaft(dAtA[start:])
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
	ErrInvalidLengthInternalRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternalRaft   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("internal_raft.proto", fileDescriptorInternalRaft) }

var fileDescriptorInternalRaft = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x4d, 0x42, 0x8b, 0x4e, 0x1b, 0x2d, 0x63, 0x17, 0xa1, 0x42, 0x52, 0xba, 0xca, 0x2a,
	0x42, 0xbd, 0x40, 0x29, 0x55, 0x90, 0x80, 0xc8, 0x28, 0x71, 0x19, 0xa6, 0xcd, 0x6b, 0x0c, 0xa6,
	0x33, 0x61, 0xf2, 0x2a, 0xed, 0x4d, 0x3c, 0x92, 0x4b, 0x4f, 0x10, 0x24, 0x2e, 0xbd, 0x81, 0x2b,
	0x49, 0x82, 0xa5, 0x76, 0xf7, 0xff, 0xff, 0x37, 0xef, 0xcd, 0xcf, 0x23, 0xe7, 0x89, 0x40, 0x50,
	0x82, 0xa7, 0xa1, 0xe2, 0x4b, 0xf4, 0x32, 0x25, 0x51, 0x52, 0x73, 0xb5, 0x4e, 0x31, 0xa9, 0x82,
	0x39, 0xcf, 0x61, 0x70, 0xba, 0x02, 0xe4, 0x11, 0x47, 0xde, 0xe0, 0x41, 0x3f, 0x96, 0xb1, 0xac,
	0xe5, 0x65, 0xa5, 0x9a, 0x74, 0xf4, 0x44, 0x4c, 0xc6, 0x97, 0xf8, 0x28, 0x57, 0xf3, 0x1c, 0xa5,
	0x00, 0x7a, 0x43, 0xce, 0x04, 0x6c, 0x30, 0x54, 0x90, 0xa5, 0xc9, 0x82, 0x87, 0x49, 0x64, 0x69,
	0x43, 0xcd, 0x6d, 0x4d, 0xed, 0xb2, 0x70, 0xcc, 0x3b, 0xd8, 0x20, 0x6b, 0xc8, 0xed, 0xec, 0xa7,
	0x70, 0x4e, 0x76, 0x86, 0x99, 0x62, 0x8f, 0x45, 0xa3, 0x6f, 0x8d, 0xf4, 0xaa, 0xcd, 0x0f, 0x82,
	0x67, 0xf9, 0xb3, 0xc4, 0x19, 0x47, 0x4e, 0x27, 0xc4, 0xcc, 0xe2, 0x30, 0x82, 0x7c, 0xa1, 0x92,
	0x0c, 0xa5, 0xaa, 0x57, 0x77, 0xc6, 0x17, 0xde, 0xbf, 0xea, 0xde, 0x7d, 0x3c, 0xdb, 0x3d, 0x61,
	0xdd, 0x6c, 0xcf, 0xd1, 0x09, 0xd1, 0xfd, 0xc0, 0xd2, 0x87, 0x86, 0xdb, 0x19, 0xbb, 0x07, 0x63,
	0x87, 0xdf, 0x79, 0x3e, 0x6c, 0x03, 0x9e, 0xae, 0x61, 0xda, 0x2e, 0x0b, 0x47, 0xf7, 0x03, 0xa6,
	0xfb, 0x01, 0x75, 0x48, 0x27, 0x95, 0x71, 0x08, 0x02, 0x55, 0x02, 0xb9, 0x65, 0x0c, 0x0d, 0xb7,
	0xcb, 0x48, 0x2a, 0xe3, 0xeb, 0x26, 0x19, 0x8c, 0xc9, 0xf1, 0xdf, 0x20, 0xed, 0x11, 0xe3, 0x05,
	0xb6, 0x75, 0xcd, 0x2e, 0xab, 0x24, 0xed, 0x93, 0xd6, 0x6b, 0x85, 0x2c, 0xbd, 0xce, 0x1a, 0x33,
	0xed, 0xbd, 0x97, 0xb6, 0xf6, 0x51, 0xda, 0xda, 0x67, 0x69, 0x6b, 0x6f, 0x5f, 0xf6, 0xd1, 0xbc,
	0x5d, 0xdf, 0xf7, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xdc, 0xb5, 0x87, 0xab, 0x01, 0x00,
	0x00,
}