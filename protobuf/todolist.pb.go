// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: todolist.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type CreateRequest struct {
	Matter            string `protobuf:"bytes,1,opt,name=matter,proto3" json:"matter,omitempty"`
	EndTime           string `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	FinishedCondition string `protobuf:"bytes,3,opt,name=finishedCondition,proto3" json:"finishedCondition,omitempty"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f86ffb2bbe2fd5, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetMatter() string {
	if m != nil {
		return m.Matter
	}
	return ""
}

func (m *CreateRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *CreateRequest) GetFinishedCondition() string {
	if m != nil {
		return m.FinishedCondition
	}
	return ""
}

type CreateResponse struct {
	ResMessage string `protobuf:"bytes,1,opt,name=resMessage,proto3" json:"resMessage,omitempty"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f86ffb2bbe2fd5, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetResMessage() string {
	if m != nil {
		return m.ResMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "protobuf.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "protobuf.CreateResponse")
}

func init() { proto.RegisterFile("todolist.proto", fileDescriptor_36f86ffb2bbe2fd5) }

var fileDescriptor_36f86ffb2bbe2fd5 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x6f, 0x15, 0xa2, 0x0e, 0x18, 0x71, 0x0b, 0x5d, 0x2c, 0x16, 0xb9, 0xca, 0x42, 0x0e,
	0xd1, 0x37, 0x30, 0xd8, 0x69, 0x13, 0xaf, 0xb0, 0xbd, 0xb8, 0x13, 0x1d, 0x30, 0x3b, 0x71, 0x67,
	0xe2, 0x73, 0xf8, 0x58, 0x96, 0x29, 0x2d, 0xe5, 0xee, 0x45, 0x84, 0xf3, 0x16, 0x0c, 0xa9, 0x96,
	0x7f, 0xff, 0x99, 0xf9, 0xf8, 0x60, 0xac, 0x1c, 0xf8, 0x8d, 0x44, 0xab, 0x65, 0x62, 0x65, 0xbb,
	0xdf, 0x3f, 0xb3, 0xd5, 0xbc, 0x64, 0x38, 0x9c, 0x24, 0x6c, 0x14, 0xa7, 0xf8, 0xbe, 0x42, 0x51,
	0x7b, 0x02, 0xa3, 0x45, 0xa3, 0x8a, 0xc9, 0x99, 0x73, 0x73, 0x71, 0x30, 0x1d, 0x92, 0x75, 0xb0,
	0x87, 0x31, 0xd4, 0xb4, 0x40, 0xb7, 0xd3, 0x17, 0x39, 0xda, 0x4b, 0x38, 0x9e, 0x53, 0x24, 0x79,
	0xc5, 0x30, 0xe1, 0x18, 0x48, 0x89, 0xa3, 0xdb, 0xed, 0x67, 0xb6, 0x8b, 0xf2, 0x0a, 0xc6, 0x19,
	0x28, 0x4b, 0x8e, 0x82, 0xd6, 0x03, 0x24, 0x94, 0x07, 0x14, 0x69, 0x5e, 0x70, 0xa0, 0xfe, 0xfb,
	0xb9, 0x7e, 0x82, 0xa3, 0x9a, 0x03, 0xdf, 0x93, 0xe8, 0x23, 0xa6, 0x0f, 0x7a, 0x46, 0x7b, 0x97,
	0x8f, 0xd4, 0x83, 0x97, 0x3d, 0xad, 0xb2, 0x52, 0xb5, 0xe1, 0x73, 0xe6, 0xb6, 0x8b, 0x3f, 0x6e,
	0x59, 0xdc, 0xba, 0xaf, 0xd6, 0x9b, 0x75, 0xeb, 0xcd, 0x4f, 0xeb, 0xcd, 0x67, 0xe7, 0x8b, 0x75,
	0xe7, 0x8b, 0xef, 0xce, 0x17, 0xb3, 0x51, 0xbf, 0x74, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0x2e, 0xdd, 0x6a, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoListServiceClient is the client API for TodoListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoListServiceClient interface {
	CreateTodolist(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type todoListServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoListServiceClient(cc *grpc.ClientConn) TodoListServiceClient {
	return &todoListServiceClient{cc}
}

func (c *todoListServiceClient) CreateTodolist(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/protobuf.TodoListService/CreateTodolist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServiceServer is the server API for TodoListService service.
type TodoListServiceServer interface {
	CreateTodolist(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedTodoListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoListServiceServer struct {
}

func (*UnimplementedTodoListServiceServer) CreateTodolist(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodolist not implemented")
}

func RegisterTodoListServiceServer(s *grpc.Server, srv TodoListServiceServer) {
	s.RegisterService(&_TodoListService_serviceDesc, srv)
}

func _TodoListService_CreateTodolist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).CreateTodolist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TodoListService/CreateTodolist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).CreateTodolist(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.TodoListService",
	HandlerType: (*TodoListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodolist",
			Handler:    _TodoListService_CreateTodolist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist.proto",
}

func (m *CreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinishedCondition) > 0 {
		i -= len(m.FinishedCondition)
		copy(dAtA[i:], m.FinishedCondition)
		i = encodeVarintTodolist(dAtA, i, uint64(len(m.FinishedCondition)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EndTime) > 0 {
		i -= len(m.EndTime)
		copy(dAtA[i:], m.EndTime)
		i = encodeVarintTodolist(dAtA, i, uint64(len(m.EndTime)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Matter) > 0 {
		i -= len(m.Matter)
		copy(dAtA[i:], m.Matter)
		i = encodeVarintTodolist(dAtA, i, uint64(len(m.Matter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResMessage) > 0 {
		i -= len(m.ResMessage)
		copy(dAtA[i:], m.ResMessage)
		i = encodeVarintTodolist(dAtA, i, uint64(len(m.ResMessage)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTodolist(dAtA []byte, offset int, v uint64) int {
	offset -= sovTodolist(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Matter)
	if l > 0 {
		n += 1 + l + sovTodolist(uint64(l))
	}
	l = len(m.EndTime)
	if l > 0 {
		n += 1 + l + sovTodolist(uint64(l))
	}
	l = len(m.FinishedCondition)
	if l > 0 {
		n += 1 + l + sovTodolist(uint64(l))
	}
	return n
}

func (m *CreateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResMessage)
	if l > 0 {
		n += 1 + l + sovTodolist(uint64(l))
	}
	return n
}

func sovTodolist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTodolist(x uint64) (n int) {
	return sovTodolist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTodolist
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
			return fmt.Errorf("proto: CreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTodolist
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
				return ErrInvalidLengthTodolist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTodolist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTodolist
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
				return ErrInvalidLengthTodolist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTodolist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishedCondition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTodolist
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
				return ErrInvalidLengthTodolist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTodolist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinishedCondition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTodolist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTodolist
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
func (m *CreateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTodolist
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
			return fmt.Errorf("proto: CreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTodolist
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
				return ErrInvalidLengthTodolist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTodolist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTodolist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTodolist
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
func skipTodolist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTodolist
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
					return 0, ErrIntOverflowTodolist
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
					return 0, ErrIntOverflowTodolist
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
				return 0, ErrInvalidLengthTodolist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTodolist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTodolist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTodolist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTodolist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTodolist = fmt.Errorf("proto: unexpected end of group")
)