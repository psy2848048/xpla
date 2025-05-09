// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xpla/volunteer/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryVolunteerValidatorsRequest
type QueryVolunteerValidatorsRequest struct {
}

func (m *QueryVolunteerValidatorsRequest) Reset()         { *m = QueryVolunteerValidatorsRequest{} }
func (m *QueryVolunteerValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVolunteerValidatorsRequest) ProtoMessage()    {}
func (*QueryVolunteerValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66f35a978c82749b, []int{0}
}
func (m *QueryVolunteerValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVolunteerValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVolunteerValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVolunteerValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVolunteerValidatorsRequest.Merge(m, src)
}
func (m *QueryVolunteerValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVolunteerValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVolunteerValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVolunteerValidatorsRequest proto.InternalMessageInfo

// QueryVolunteerValidatorsResponse
type QueryVolunteerValidatorsResponse struct {
	VolunteerValidators []string `protobuf:"bytes,1,rep,name=volunteer_validators,json=volunteerValidators,proto3" json:"volunteer_validators,omitempty"`
}

func (m *QueryVolunteerValidatorsResponse) Reset()         { *m = QueryVolunteerValidatorsResponse{} }
func (m *QueryVolunteerValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVolunteerValidatorsResponse) ProtoMessage()    {}
func (*QueryVolunteerValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66f35a978c82749b, []int{1}
}
func (m *QueryVolunteerValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVolunteerValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVolunteerValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVolunteerValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVolunteerValidatorsResponse.Merge(m, src)
}
func (m *QueryVolunteerValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVolunteerValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVolunteerValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVolunteerValidatorsResponse proto.InternalMessageInfo

func (m *QueryVolunteerValidatorsResponse) GetVolunteerValidators() []string {
	if m != nil {
		return m.VolunteerValidators
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryVolunteerValidatorsRequest)(nil), "xpla.volunteer.v1beta1.QueryVolunteerValidatorsRequest")
	proto.RegisterType((*QueryVolunteerValidatorsResponse)(nil), "xpla.volunteer.v1beta1.QueryVolunteerValidatorsResponse")
}

func init() {
	proto.RegisterFile("xpla/volunteer/v1beta1/query.proto", fileDescriptor_66f35a978c82749b)
}

var fileDescriptor_66f35a978c82749b = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xaa, 0x28, 0xc8, 0x49,
	0xd4, 0x2f, 0xcb, 0xcf, 0x29, 0xcd, 0x2b, 0x49, 0x4d, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x03, 0xa9, 0xd1, 0x83, 0xab, 0xd1, 0x83, 0xaa, 0x91, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0xe8, 0x52, 0x52, 0xe4, 0x92, 0x0f, 0x04, 0x19, 0x12, 0x06, 0xd3, 0x17, 0x96, 0x98, 0x93,
	0x99, 0x92, 0x58, 0x92, 0x5f, 0x54, 0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x14, 0xca,
	0xa5, 0x80, 0x5b, 0x49, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x21, 0x97, 0x08, 0xdc, 0xe6,
	0xf8, 0x32, 0xb8, 0xbc, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x67, 0x90, 0x70, 0x19, 0xa6, 0x56, 0xa3,
	0x7d, 0x8c, 0x5c, 0xac, 0x60, 0x73, 0x85, 0xb6, 0x30, 0x72, 0x09, 0x63, 0x31, 0x5c, 0xc8, 0x5c,
	0x0f, 0xbb, 0x97, 0xf4, 0x08, 0xb8, 0x58, 0xca, 0x82, 0x74, 0x8d, 0x10, 0x7f, 0x28, 0x69, 0x35,
	0x5d, 0x7e, 0x32, 0x99, 0x49, 0x45, 0x48, 0x49, 0x1f, 0x47, 0x88, 0x23, 0xfc, 0xe6, 0xe4, 0x7c,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x60, 0x73, 0x52, 0x52, 0xcb, 0x20, 0xe6, 0x55, 0x20, 0x99, 0x58,
	0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x2c, 0x22, 0xf6, 0xe2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// VolunteerValidators
	VolunteerValidators(ctx context.Context, in *QueryVolunteerValidatorsRequest, opts ...grpc.CallOption) (*QueryVolunteerValidatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) VolunteerValidators(ctx context.Context, in *QueryVolunteerValidatorsRequest, opts ...grpc.CallOption) (*QueryVolunteerValidatorsResponse, error) {
	out := new(QueryVolunteerValidatorsResponse)
	err := c.cc.Invoke(ctx, "/xpla.volunteer.v1beta1.Query/VolunteerValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// VolunteerValidators
	VolunteerValidators(context.Context, *QueryVolunteerValidatorsRequest) (*QueryVolunteerValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) VolunteerValidators(ctx context.Context, req *QueryVolunteerValidatorsRequest) (*QueryVolunteerValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolunteerValidators not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_VolunteerValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVolunteerValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VolunteerValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xpla.volunteer.v1beta1.Query/VolunteerValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VolunteerValidators(ctx, req.(*QueryVolunteerValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xpla.volunteer.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VolunteerValidators",
			Handler:    _Query_VolunteerValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xpla/volunteer/v1beta1/query.proto",
}

func (m *QueryVolunteerValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVolunteerValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVolunteerValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryVolunteerValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVolunteerValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVolunteerValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VolunteerValidators) > 0 {
		for iNdEx := len(m.VolunteerValidators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VolunteerValidators[iNdEx])
			copy(dAtA[i:], m.VolunteerValidators[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.VolunteerValidators[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryVolunteerValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryVolunteerValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VolunteerValidators) > 0 {
		for _, s := range m.VolunteerValidators {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryVolunteerValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVolunteerValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVolunteerValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryVolunteerValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryVolunteerValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVolunteerValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolunteerValidators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolunteerValidators = append(m.VolunteerValidators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
