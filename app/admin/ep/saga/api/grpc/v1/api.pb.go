// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/admin/ep/saga/api/grpc/v1/api.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type PushMsgReq struct {
	Username             []string `protobuf:"bytes,1,rep,name=username" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgReq) Reset()         { *m = PushMsgReq{} }
func (m *PushMsgReq) String() string { return proto.CompactTextString(m) }
func (*PushMsgReq) ProtoMessage()    {}
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d78a3f3b47284b33, []int{0}
}
func (m *PushMsgReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushMsgReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PushMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgReq.Merge(dst, src)
}
func (m *PushMsgReq) XXX_Size() int {
	return m.Size()
}
func (m *PushMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgReq proto.InternalMessageInfo

func (m *PushMsgReq) GetUsername() []string {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *PushMsgReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PushMsgReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgReply) Reset()         { *m = PushMsgReply{} }
func (m *PushMsgReply) String() string { return proto.CompactTextString(m) }
func (*PushMsgReply) ProtoMessage()    {}
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d78a3f3b47284b33, []int{1}
}
func (m *PushMsgReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushMsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushMsgReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PushMsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgReply.Merge(dst, src)
}
func (m *PushMsgReply) XXX_Size() int {
	return m.Size()
}
func (m *PushMsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PushMsgReq)(nil), "test.ep.sagaadmin.v1.PushMsgReq")
	proto.RegisterType((*PushMsgReply)(nil), "test.ep.sagaadmin.v1.PushMsgReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SagaAdminClient is the client API for SagaAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SagaAdminClient interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error)
}

type sagaAdminClient struct {
	cc *grpc.ClientConn
}

func NewSagaAdminClient(cc *grpc.ClientConn) SagaAdminClient {
	return &sagaAdminClient{cc}
}

func (c *sagaAdminClient) PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error) {
	out := new(PushMsgReply)
	err := c.cc.Invoke(ctx, "/test.ep.sagaadmin.v1.SagaAdmin/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SagaAdminServer is the server API for SagaAdmin service.
type SagaAdminServer interface {
	PushMsg(context.Context, *PushMsgReq) (*PushMsgReply, error)
}

func RegisterSagaAdminServer(s *grpc.Server, srv SagaAdminServer) {
	s.RegisterService(&_SagaAdmin_serviceDesc, srv)
}

func _SagaAdmin_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SagaAdminServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ep.sagaadmin.v1.SagaAdmin/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SagaAdminServer).PushMsg(ctx, req.(*PushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SagaAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.ep.sagaadmin.v1.SagaAdmin",
	HandlerType: (*SagaAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _SagaAdmin_PushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/admin/ep/saga/api/grpc/v1/api.proto",
}

func (m *PushMsgReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushMsgReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		for _, s := range m.Username {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PushMsgReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushMsgReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PushMsgReq) Size() (n int) {
	var l int
	_ = l
	if len(m.Username) > 0 {
		for _, s := range m.Username {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PushMsgReply) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PushMsgReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: PushMsgReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushMsgReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = append(m.Username, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushMsgReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: PushMsgReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushMsgReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
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
				next, err := skipApi(dAtA[start:])
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
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/admin/ep/saga/api/grpc/v1/api.proto", fileDescriptor_api_d78a3f3b47284b33)
}

var fileDescriptor_api_d78a3f3b47284b33 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2c, 0x28, 0xd0,
	0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x4f, 0x2d, 0xd0, 0x2f, 0x4e, 0x4c, 0x4f, 0xd4, 0x4f, 0x2c,
	0xc8, 0xd4, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x33, 0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x44, 0x4a, 0x52, 0x8b, 0x4b, 0xf4, 0x52, 0x0b, 0xf4, 0x40, 0x6a, 0xc0, 0x1a, 0xf4,
	0xca, 0x0c, 0x95, 0x9c, 0xb8, 0xb8, 0x02, 0x4a, 0x8b, 0x33, 0x7c, 0x8b, 0xd3, 0x83, 0x52, 0x0b,
	0x85, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x98,
	0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc, 0xbc, 0x92, 0xd4, 0xbc, 0x12,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x89, 0x8f, 0x8b, 0x07, 0x6e, 0x46, 0x41,
	0x4e, 0xa5, 0x51, 0x0c, 0x17, 0x67, 0x70, 0x62, 0x7a, 0xa2, 0x23, 0xc8, 0x0e, 0x21, 0x7f, 0x2e,
	0x76, 0xa8, 0xa4, 0x90, 0x82, 0x1e, 0x36, 0x27, 0xe8, 0x21, 0xec, 0x97, 0x52, 0x22, 0xa0, 0xa2,
	0x20, 0xa7, 0xd2, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x8c, 0x62, 0x2a, 0x33, 0x4c, 0x62, 0x03, 0x7b, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xac, 0x58, 0xa2, 0x9e, 0x0b, 0x01, 0x00, 0x00,
}
