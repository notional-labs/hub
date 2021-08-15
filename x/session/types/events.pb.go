// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sentinel-official/hub/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventModule struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
}

func (m *EventModule) Reset()         { *m = EventModule{} }
func (m *EventModule) String() string { return proto.CompactTextString(m) }
func (*EventModule) ProtoMessage()    {}
func (*EventModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{0}
}
func (m *EventModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventModule.Merge(m, src)
}
func (m *EventModule) XXX_Size() int {
	return m.Size()
}
func (m *EventModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EventModule.DiscardUnknown(m)
}

var xxx_messageInfo_EventModule proto.InternalMessageInfo

type EventStart struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node         string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Subscription uint64 `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty" yaml:"subscription"`
}

func (m *EventStart) Reset()         { *m = EventStart{} }
func (m *EventStart) String() string { return proto.CompactTextString(m) }
func (*EventStart) ProtoMessage()    {}
func (*EventStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{1}
}
func (m *EventStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStart.Merge(m, src)
}
func (m *EventStart) XXX_Size() int {
	return m.Size()
}
func (m *EventStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventStart proto.InternalMessageInfo

type EventUpdate struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node         string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Subscription uint64 `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty" yaml:"subscription"`
}

func (m *EventUpdate) Reset()         { *m = EventUpdate{} }
func (m *EventUpdate) String() string { return proto.CompactTextString(m) }
func (*EventUpdate) ProtoMessage()    {}
func (*EventUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{2}
}
func (m *EventUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdate.Merge(m, src)
}
func (m *EventUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdate proto.InternalMessageInfo

type EventSetStatus struct {
	Id           uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node         string       `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Subscription uint64       `protobuf:"varint,3,opt,name=subscription,proto3" json:"subscription,omitempty" yaml:"subscription"`
	Status       types.Status `protobuf:"varint,4,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventSetStatus) Reset()         { *m = EventSetStatus{} }
func (m *EventSetStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetStatus) ProtoMessage()    {}
func (*EventSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34d43b21999bd7a, []int{3}
}
func (m *EventSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetStatus.Merge(m, src)
}
func (m *EventSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventModule)(nil), "sentinel.session.v1.EventModule")
	proto.RegisterType((*EventStart)(nil), "sentinel.session.v1.EventStart")
	proto.RegisterType((*EventUpdate)(nil), "sentinel.session.v1.EventUpdate")
	proto.RegisterType((*EventSetStatus)(nil), "sentinel.session.v1.EventSetStatus")
}

func init() { proto.RegisterFile("sentinel/session/v1/events.proto", fileDescriptor_e34d43b21999bd7a) }

var fileDescriptor_e34d43b21999bd7a = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x41, 0x4e, 0xfa, 0x40,
	0x14, 0xc6, 0x3b, 0xfc, 0x09, 0x09, 0xc3, 0x1f, 0x8c, 0xc5, 0xc4, 0x4a, 0xe2, 0x94, 0x8c, 0x1b,
	0x36, 0xb6, 0x82, 0x3b, 0xdd, 0x35, 0xba, 0x74, 0x53, 0xe3, 0xc6, 0x5d, 0xcb, 0x0c, 0x30, 0x49,
	0xe9, 0x34, 0xcc, 0xb4, 0x91, 0x43, 0x68, 0x3c, 0x86, 0x47, 0x61, 0xc9, 0x4e, 0x57, 0x8d, 0x96,
	0x1b, 0xf4, 0x04, 0xa6, 0xd3, 0x42, 0x34, 0x1e, 0x80, 0x5d, 0xf3, 0xde, 0xf7, 0xfd, 0xde, 0xeb,
	0x9b, 0x0f, 0xf6, 0x05, 0x0d, 0x25, 0x0b, 0x69, 0x60, 0x0b, 0x2a, 0x04, 0xe3, 0xa1, 0x9d, 0x0c,
	0x6d, 0x9a, 0xd0, 0x50, 0x0a, 0x2b, 0x5a, 0x70, 0xc9, 0xf5, 0xee, 0x56, 0x61, 0x55, 0x0a, 0x2b,
	0x19, 0xf6, 0x8e, 0xa6, 0x7c, 0xca, 0x55, 0xdf, 0x2e, 0xbe, 0x4a, 0x69, 0x0f, 0xed, 0x60, 0x72,
	0x19, 0x51, 0x51, 0xa0, 0x84, 0xf4, 0x64, 0x5c, 0xa1, 0xf0, 0x08, 0xb6, 0x6e, 0x0b, 0xf4, 0x1d,
	0x27, 0x71, 0x40, 0xf5, 0x33, 0x58, 0x0f, 0xbd, 0x39, 0x35, 0x40, 0x1f, 0x0c, 0x9a, 0xce, 0x41,
	0x9e, 0x9a, 0xad, 0xa5, 0x37, 0x0f, 0xae, 0x70, 0x51, 0xc5, 0xae, 0x6a, 0xe2, 0x67, 0x00, 0xa1,
	0x32, 0xdd, 0x4b, 0x6f, 0x21, 0xf5, 0x53, 0x58, 0x63, 0x44, 0x39, 0xea, 0x4e, 0x3b, 0x4f, 0xcd,
	0x66, 0xe9, 0x60, 0x04, 0xbb, 0x35, 0x46, 0x14, 0x92, 0x13, 0x6a, 0xd4, 0xfe, 0x20, 0x39, 0x51,
	0x48, 0x4e, 0xa8, 0x7e, 0x0d, 0xff, 0x8b, 0xd8, 0x17, 0xe3, 0x05, 0x8b, 0x24, 0xe3, 0xa1, 0xf1,
	0x4f, 0xd1, 0x8e, 0xf3, 0xd4, 0xec, 0x96, 0xe2, 0x9f, 0x5d, 0xec, 0xfe, 0x12, 0xe3, 0x17, 0x50,
	0xfd, 0xc4, 0x43, 0x44, 0x3c, 0x49, 0xf7, 0xbf, 0xd0, 0x3b, 0x80, 0x9d, 0xf2, 0x40, 0xb4, 0xb8,
	0x91, 0x8c, 0xc5, 0xde, 0x77, 0xd2, 0x6f, 0x60, 0xa3, 0x7c, 0x78, 0xa3, 0xde, 0x07, 0x83, 0xce,
	0xe8, 0xc4, 0xda, 0x85, 0x48, 0x25, 0xc3, 0x4a, 0x86, 0x56, 0xb9, 0xab, 0x73, 0x98, 0xa7, 0x66,
	0xbb, 0x22, 0xaa, 0x0a, 0x76, 0x2b, 0xaf, 0xe3, 0xae, 0xbe, 0x90, 0xf6, 0x96, 0x21, 0x6d, 0x95,
	0x21, 0xb0, 0xce, 0x10, 0xf8, 0xcc, 0x10, 0x78, 0xdd, 0x20, 0x6d, 0xbd, 0x41, 0xda, 0xc7, 0x06,
	0x69, 0x8f, 0x17, 0x53, 0x26, 0x67, 0xb1, 0x6f, 0x8d, 0xf9, 0xdc, 0xde, 0x4e, 0x38, 0xe7, 0x93,
	0x09, 0x1b, 0x33, 0x2f, 0xb0, 0x67, 0xb1, 0x6f, 0x3f, 0xed, 0x72, 0xad, 0x06, 0xfb, 0x0d, 0x95,
	0xc4, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x7e, 0x3c, 0x7d, 0xf8, 0x02, 0x00, 0x00,
}

func (m *EventModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Subscription != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Subscription))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventStart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	return n
}

func (m *EventUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	return n
}

func (m *EventSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Subscription != 0 {
		n += 1 + sovEvents(uint64(m.Subscription))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			m.Subscription = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Subscription |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
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
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
