// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fatal_fruit/ns/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgReserve represents a message to purchase a namespace.
type MsgReserve struct {
	// name defines the human readable address
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// resolveAddress defines the bech32 address to resolve to
	ResolveAddress string `protobuf:"bytes,2,opt,name=resolveAddress,proto3" json:"resolveAddress,omitempty"`
	// owner is the address of the owner of listing
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	//   price is the last price paid for listing
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MsgReserve) Reset()         { *m = MsgReserve{} }
func (m *MsgReserve) String() string { return proto.CompactTextString(m) }
func (*MsgReserve) ProtoMessage()    {}
func (*MsgReserve) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0075a375a07b0e2, []int{0}
}
func (m *MsgReserve) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReserve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReserve.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReserve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReserve.Merge(m, src)
}
func (m *MsgReserve) XXX_Size() int {
	return m.Size()
}
func (m *MsgReserve) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReserve.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReserve proto.InternalMessageInfo

func (m *MsgReserve) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgReserve) GetResolveAddress() string {
	if m != nil {
		return m.ResolveAddress
	}
	return ""
}

func (m *MsgReserve) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgReserve) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// MsgSendResponse defines the Msg/Reserve response type.
type MsgReserveResponse struct {
}

func (m *MsgReserveResponse) Reset()         { *m = MsgReserveResponse{} }
func (m *MsgReserveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReserveResponse) ProtoMessage()    {}
func (*MsgReserveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0075a375a07b0e2, []int{1}
}
func (m *MsgReserveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReserveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReserveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReserveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReserveResponse.Merge(m, src)
}
func (m *MsgReserveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReserveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReserveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReserveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgReserve)(nil), "fatal_fruit.ns.v1.MsgReserve")
	proto.RegisterType((*MsgReserveResponse)(nil), "fatal_fruit.ns.v1.MsgReserveResponse")
}

func init() { proto.RegisterFile("fatal_fruit/ns/v1/tx.proto", fileDescriptor_c0075a375a07b0e2) }

var fileDescriptor_c0075a375a07b0e2 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0x4e, 0xfc, 0x55, 0x3a, 0x85, 0x82, 0x83, 0xd0, 0x18, 0x30, 0x8a, 0x50, 0x10, 0xc1, 0x19,
	0x62, 0xe9, 0xa1, 0x3d, 0xb5, 0xf6, 0x2c, 0x85, 0xf4, 0x56, 0x28, 0x92, 0xe8, 0x38, 0x86, 0x9a,
	0x19, 0xc9, 0x1b, 0xd3, 0xf6, 0x56, 0x7a, 0xea, 0xb1, 0x7f, 0x46, 0xe9, 0xc9, 0x43, 0xff, 0x08,
	0x8f, 0xb2, 0xa7, 0x3d, 0xed, 0x2e, 0x7a, 0xf0, 0xdf, 0x58, 0x92, 0x8c, 0xab, 0xec, 0x2e, 0xbb,
	0x97, 0xe4, 0xcd, 0xfb, 0xbe, 0xf7, 0xbd, 0xf7, 0xbe, 0x87, 0xec, 0xa9, 0xaf, 0xfc, 0xf9, 0x68,
	0x1a, 0x2f, 0x43, 0x45, 0x05, 0xd0, 0xc4, 0xa5, 0xea, 0x3b, 0x59, 0xc4, 0x52, 0x49, 0x5c, 0x3d,
	0xc1, 0x88, 0x00, 0x92, 0xb8, 0xf6, 0x8b, 0xb1, 0x84, 0x48, 0x02, 0x8d, 0x80, 0xa7, 0xd4, 0x08,
	0x78, 0xce, 0xb5, 0x6b, 0x5c, 0x72, 0x99, 0x85, 0x34, 0x8d, 0x74, 0xd6, 0xd1, 0xf4, 0xc0, 0x07,
	0x46, 0x13, 0x37, 0x60, 0xca, 0x77, 0xe9, 0x58, 0x86, 0x42, 0xe3, 0x55, 0x3f, 0x0a, 0x85, 0xa4,
	0xd9, 0x57, 0xa7, 0x1a, 0xf7, 0x0c, 0xf4, 0x63, 0xc1, 0x40, 0xc3, 0xf5, 0x5c, 0x71, 0x94, 0xb7,
	0xca, 0x1f, 0x39, 0xd4, 0xfe, 0x5d, 0x40, 0x68, 0x08, 0xdc, 0x63, 0xc0, 0xe2, 0x84, 0x61, 0x8c,
	0x4a, 0xc2, 0x8f, 0x98, 0x65, 0xb6, 0xcc, 0xce, 0x53, 0x2f, 0x8b, 0xf1, 0x3b, 0xf4, 0x3c, 0x66,
	0x20, 0xe7, 0x09, 0x7b, 0x3f, 0x99, 0xc4, 0x0c, 0xc0, 0x2a, 0xa4, 0xe8, 0xc0, 0x3a, 0xfb, 0xdf,
	0xab, 0x69, 0x31, 0x8d, 0x7c, 0x52, 0x71, 0x28, 0xb8, 0x77, 0x8b, 0x8f, 0x09, 0x2a, 0xcb, 0x6f,
	0x82, 0xc5, 0x56, 0xf1, 0x91, 0xc2, 0x9c, 0x86, 0x67, 0xa8, 0xe2, 0x47, 0x72, 0x29, 0x94, 0x55,
	0x6a, 0x15, 0x3b, 0xcf, 0xfa, 0x75, 0xa2, 0xd9, 0xa9, 0x25, 0x44, 0x5b, 0x42, 0x3e, 0xc8, 0x50,
	0x0c, 0x5e, 0xaf, 0x2f, 0x9a, 0xc6, 0xbf, 0xcb, 0x66, 0x87, 0x87, 0x6a, 0xb6, 0x0c, 0xc8, 0x58,
	0x46, 0x7a, 0x41, 0xfd, 0xeb, 0xc1, 0xe4, 0xab, 0x36, 0x23, 0x2d, 0x80, 0xbf, 0xfb, 0x55, 0xd7,
	0xf4, 0xb4, 0xfe, 0x5b, 0xf4, 0x6b, 0xbf, 0xea, 0xe6, 0x5d, 0xdb, 0x35, 0x84, 0x8f, 0x4e, 0x78,
	0x0c, 0x16, 0x52, 0x00, 0xeb, 0x7f, 0x41, 0xc5, 0x21, 0x70, 0xfc, 0x11, 0x3d, 0x39, 0x78, 0xd4,
	0x20, 0x77, 0x4e, 0x4c, 0x8e, 0x85, 0xf6, 0xcb, 0x07, 0xe1, 0x83, 0xae, 0x5d, 0xfe, 0x99, 0x0e,
	0x32, 0x78, 0xb3, 0xde, 0x3a, 0xe6, 0x66, 0xeb, 0x98, 0x57, 0x5b, 0xc7, 0xfc, 0xb3, 0x73, 0x8c,
	0xcd, 0xce, 0x31, 0xce, 0x77, 0x8e, 0xf1, 0xb9, 0x79, 0xb2, 0x51, 0xa6, 0xd8, 0xbb, 0x39, 0x6f,
	0xb6, 0x4e, 0x50, 0xc9, 0x2e, 0xf8, 0xea, 0x3a, 0x00, 0x00, 0xff, 0xff, 0x97, 0xf8, 0x73, 0x90,
	0x8e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Reserve defines a method to buy a namespace with an associated bech32
	// address to resolve to.
	Reserve(ctx context.Context, in *MsgReserve, opts ...grpc.CallOption) (*MsgReserveResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Reserve(ctx context.Context, in *MsgReserve, opts ...grpc.CallOption) (*MsgReserveResponse, error) {
	out := new(MsgReserveResponse)
	err := c.cc.Invoke(ctx, "/fatal_fruit.ns.v1.Msg/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Reserve defines a method to buy a namespace with an associated bech32
	// address to resolve to.
	Reserve(context.Context, *MsgReserve) (*MsgReserveResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Reserve(ctx context.Context, req *MsgReserve) (*MsgReserveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReserve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fatal_fruit.ns.v1.Msg/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Reserve(ctx, req.(*MsgReserve))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fatal_fruit.ns.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reserve",
			Handler:    _Msg_Reserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fatal_fruit/ns/v1/tx.proto",
}

func (m *MsgReserve) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReserve) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReserve) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ResolveAddress) > 0 {
		i -= len(m.ResolveAddress)
		copy(dAtA[i:], m.ResolveAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ResolveAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReserveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReserveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReserveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgReserve) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ResolveAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgReserveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgReserve) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReserve: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReserve: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResolveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgReserveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgReserveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReserveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)