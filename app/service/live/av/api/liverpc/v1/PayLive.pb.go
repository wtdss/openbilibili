// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/PayLive.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PayLiveLiveValidateReq struct {
	// 房间id
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	// 来源：pc：表示PC端；pc_link：表示PC直播姬；ios：表示ios端；ios_link：表示ios_link端；android：安卓端；android_link：安卓link端；ipad：ipad；android_pad:安卓pad端；live_mng：表示live后台;vc_mng：表示vc后台;"
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform"`
}

func (m *PayLiveLiveValidateReq) Reset()         { *m = PayLiveLiveValidateReq{} }
func (m *PayLiveLiveValidateReq) String() string { return proto.CompactTextString(m) }
func (*PayLiveLiveValidateReq) ProtoMessage()    {}
func (*PayLiveLiveValidateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_PayLive_366356d2c332491f, []int{0}
}
func (m *PayLiveLiveValidateReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayLiveLiveValidateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayLiveLiveValidateReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PayLiveLiveValidateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayLiveLiveValidateReq.Merge(dst, src)
}
func (m *PayLiveLiveValidateReq) XXX_Size() int {
	return m.Size()
}
func (m *PayLiveLiveValidateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayLiveLiveValidateReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayLiveLiveValidateReq proto.InternalMessageInfo

func (m *PayLiveLiveValidateReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PayLiveLiveValidateReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type PayLiveLiveValidateResp struct {
	//
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	//
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data *PayLiveLiveValidateResp_Data `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *PayLiveLiveValidateResp) Reset()         { *m = PayLiveLiveValidateResp{} }
func (m *PayLiveLiveValidateResp) String() string { return proto.CompactTextString(m) }
func (*PayLiveLiveValidateResp) ProtoMessage()    {}
func (*PayLiveLiveValidateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_PayLive_366356d2c332491f, []int{1}
}
func (m *PayLiveLiveValidateResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayLiveLiveValidateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayLiveLiveValidateResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PayLiveLiveValidateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayLiveLiveValidateResp.Merge(dst, src)
}
func (m *PayLiveLiveValidateResp) XXX_Size() int {
	return m.Size()
}
func (m *PayLiveLiveValidateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayLiveLiveValidateResp.DiscardUnknown(m)
}

var xxx_messageInfo_PayLiveLiveValidateResp proto.InternalMessageInfo

func (m *PayLiveLiveValidateResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PayLiveLiveValidateResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *PayLiveLiveValidateResp) GetData() *PayLiveLiveValidateResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type PayLiveLiveValidateResp_Data struct {
	// 1 允许 ; 0 不允许
	Permission int64 `protobuf:"varint,1,opt,name=permission,proto3" json:"permission"`
}

func (m *PayLiveLiveValidateResp_Data) Reset()         { *m = PayLiveLiveValidateResp_Data{} }
func (m *PayLiveLiveValidateResp_Data) String() string { return proto.CompactTextString(m) }
func (*PayLiveLiveValidateResp_Data) ProtoMessage()    {}
func (*PayLiveLiveValidateResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_PayLive_366356d2c332491f, []int{1, 0}
}
func (m *PayLiveLiveValidateResp_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PayLiveLiveValidateResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PayLiveLiveValidateResp_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PayLiveLiveValidateResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayLiveLiveValidateResp_Data.Merge(dst, src)
}
func (m *PayLiveLiveValidateResp_Data) XXX_Size() int {
	return m.Size()
}
func (m *PayLiveLiveValidateResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_PayLiveLiveValidateResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_PayLiveLiveValidateResp_Data proto.InternalMessageInfo

func (m *PayLiveLiveValidateResp_Data) GetPermission() int64 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func init() {
	proto.RegisterType((*PayLiveLiveValidateReq)(nil), "av.v1.PayLiveLiveValidateReq")
	proto.RegisterType((*PayLiveLiveValidateResp)(nil), "av.v1.PayLiveLiveValidateResp")
	proto.RegisterType((*PayLiveLiveValidateResp_Data)(nil), "av.v1.PayLiveLiveValidateResp.Data")
}
func (m *PayLiveLiveValidateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayLiveLiveValidateReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RoomId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(m.RoomId))
	}
	if len(m.Platform) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(len(m.Platform)))
		i += copy(dAtA[i:], m.Platform)
	}
	return i, nil
}

func (m *PayLiveLiveValidateResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayLiveLiveValidateResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PayLiveLiveValidateResp_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PayLiveLiveValidateResp_Data) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Permission != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayLive(dAtA, i, uint64(m.Permission))
	}
	return i, nil
}

func encodeVarintPayLive(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PayLiveLiveValidateReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovPayLive(uint64(m.RoomId))
	}
	l = len(m.Platform)
	if l > 0 {
		n += 1 + l + sovPayLive(uint64(l))
	}
	return n
}

func (m *PayLiveLiveValidateResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovPayLive(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPayLive(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovPayLive(uint64(l))
	}
	return n
}

func (m *PayLiveLiveValidateResp_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Permission != 0 {
		n += 1 + sovPayLive(uint64(m.Permission))
	}
	return n
}

func sovPayLive(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayLive(x uint64) (n int) {
	return sovPayLive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PayLiveLiveValidateReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayLive
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
			return fmt.Errorf("proto: PayLiveLiveValidateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayLiveLiveValidateReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
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
				return ErrInvalidLengthPayLive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayLive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayLive
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
func (m *PayLiveLiveValidateResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayLive
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
			return fmt.Errorf("proto: PayLiveLiveValidateResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PayLiveLiveValidateResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
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
				return ErrInvalidLengthPayLive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
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
				return ErrInvalidLengthPayLive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &PayLiveLiveValidateResp_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayLive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayLive
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
func (m *PayLiveLiveValidateResp_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayLive
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			m.Permission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayLive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permission |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPayLive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayLive
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
func skipPayLive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayLive
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
					return 0, ErrIntOverflowPayLive
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
					return 0, ErrIntOverflowPayLive
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
				return 0, ErrInvalidLengthPayLive
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayLive
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
				next, err := skipPayLive(dAtA[start:])
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
	ErrInvalidLengthPayLive = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayLive   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/PayLive.proto", fileDescriptor_PayLive_366356d2c332491f) }

var fileDescriptor_PayLive_366356d2c332491f = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0x02, 0x41,
	0x18, 0x75, 0xd4, 0xd4, 0x3e, 0x25, 0x62, 0x0e, 0x65, 0x62, 0xb3, 0x62, 0x1d, 0xf6, 0xd2, 0x88,
	0x06, 0xdd, 0x93, 0x2e, 0x41, 0x41, 0xcc, 0x21, 0xa2, 0x4b, 0x8c, 0xee, 0xba, 0x0e, 0xb8, 0xcd,
	0xba, 0x3b, 0x2e, 0xf4, 0x2f, 0xfa, 0x59, 0x5d, 0x02, 0x8f, 0x9d, 0x96, 0xd0, 0xdb, 0xfe, 0x8a,
	0x98, 0x71, 0x93, 0x3d, 0x94, 0x87, 0xf9, 0x78, 0xef, 0x0d, 0xdf, 0x7b, 0x8f, 0x19, 0x38, 0x8c,
	0xfb, 0xbd, 0x07, 0xfe, 0x76, 0x27, 0x62, 0x97, 0x06, 0xa1, 0x54, 0x12, 0xef, 0xf1, 0x98, 0xc6,
	0xfd, 0xd6, 0x85, 0x27, 0xd4, 0x74, 0x31, 0xa2, 0x63, 0xe9, 0xf7, 0x3c, 0xe9, 0xc9, 0x9e, 0xb9,
	0x1d, 0x2d, 0x26, 0x86, 0x19, 0x62, 0xd0, 0x66, 0xab, 0x3b, 0x85, 0xa3, 0xcc, 0x46, 0x9f, 0x47,
	0x3e, 0x13, 0x0e, 0x57, 0x2e, 0x73, 0xe7, 0xf8, 0x1c, 0xaa, 0xa1, 0x94, 0xfe, 0x8b, 0x70, 0x9a,
	0xa8, 0x83, 0xec, 0xd2, 0xb0, 0x9e, 0x26, 0xd6, 0xaf, 0xc4, 0x2a, 0x1a, 0xdc, 0x3a, 0xd8, 0x86,
	0x5a, 0x30, 0xe3, 0x6a, 0x22, 0x43, 0xbf, 0x59, 0xec, 0x20, 0x7b, 0x7f, 0xd8, 0x48, 0x13, 0x6b,
	0xab, 0xb1, 0x2d, 0xea, 0x7e, 0x22, 0x38, 0xfe, 0x33, 0x2a, 0x0a, 0x70, 0x1b, 0xca, 0x63, 0xe9,
	0xb8, 0x59, 0x50, 0x2d, 0x4d, 0x2c, 0xc3, 0x99, 0x99, 0xf8, 0x04, 0x4a, 0x7e, 0xe4, 0x65, 0xf6,
	0xd5, 0x34, 0xb1, 0x34, 0x65, 0x7a, 0xe0, 0x6b, 0x28, 0x3b, 0x5c, 0xf1, 0x66, 0xa9, 0x83, 0xec,
	0xfa, 0xe0, 0x8c, 0x9a, 0x37, 0xa0, 0xff, 0xc4, 0xd0, 0x1b, 0xae, 0xf8, 0xc6, 0x5d, 0x2f, 0x31,
	0x33, 0x5b, 0x57, 0x50, 0xd6, 0x3a, 0xa6, 0x00, 0x81, 0x1b, 0xfa, 0x22, 0x8a, 0x84, 0x7c, 0xcd,
	0x9a, 0x1c, 0xa4, 0x89, 0x95, 0x53, 0x59, 0x0e, 0x0f, 0x9e, 0xa0, 0x9a, 0xe5, 0xe0, 0x7b, 0x68,
	0xcc, 0x72, 0x59, 0xf8, 0x74, 0x57, 0x8f, 0x79, 0x8b, 0xec, 0xae, 0x39, 0x6c, 0x7f, 0xac, 0x08,
	0x5a, 0xae, 0x08, 0xfa, 0x5e, 0x11, 0xf4, 0xbe, 0x26, 0x85, 0xe5, 0x9a, 0x14, 0xbe, 0xd6, 0xa4,
	0xf0, 0x5c, 0x8c, 0xfb, 0xa3, 0x8a, 0xf9, 0xb8, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6,
	0xd8, 0x2e, 0x24, 0x02, 0x02, 0x00, 0x00,
}
