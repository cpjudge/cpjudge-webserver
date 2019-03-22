// Code generated by protoc-gen-go. DO NOT EDIT.
// source: evaluator/evaluator.proto

package evaluator

import (
	context "context"
	fmt "fmt"
	math "math"

	submission "github.com/cpjudge/cpjudge_webserver/proto/submission"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// Ref: https://www.quora.com/What-is-WA-RTE-CTE-and-TLE-on-CodeChef
type EvaluationStatus int32

const (
	EvaluationStatus_CORRECT_ANSWER      EvaluationStatus = 0
	EvaluationStatus_WRONG_ANSWER        EvaluationStatus = 1
	EvaluationStatus_TIME_LIMIT_EXCEEDED EvaluationStatus = 2
	EvaluationStatus_COMPILATION_ERROR   EvaluationStatus = 3
	EvaluationStatus_RUNTIME_ERROR       EvaluationStatus = 4
)

var EvaluationStatus_name = map[int32]string{
	0: "CORRECT_ANSWER",
	1: "WRONG_ANSWER",
	2: "TIME_LIMIT_EXCEEDED",
	3: "COMPILATION_ERROR",
	4: "RUNTIME_ERROR",
}

var EvaluationStatus_value = map[string]int32{
	"CORRECT_ANSWER":      0,
	"WRONG_ANSWER":        1,
	"TIME_LIMIT_EXCEEDED": 2,
	"COMPILATION_ERROR":   3,
	"RUNTIME_ERROR":       4,
}

func (x EvaluationStatus) String() string {
	return proto.EnumName(EvaluationStatus_name, int32(x))
}

func (EvaluationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a1f69954773e339, []int{0}
}

type CodeStatus struct {
	CodeStatus           EvaluationStatus `protobuf:"varint,1,opt,name=code_status,json=codeStatus,proto3,enum=evaluator.EvaluationStatus" json:"code_status"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CodeStatus) Reset()         { *m = CodeStatus{} }
func (m *CodeStatus) String() string { return proto.CompactTextString(m) }
func (*CodeStatus) ProtoMessage()    {}
func (*CodeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1f69954773e339, []int{0}
}

func (m *CodeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeStatus.Unmarshal(m, b)
}
func (m *CodeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeStatus.Marshal(b, m, deterministic)
}
func (m *CodeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeStatus.Merge(m, src)
}
func (m *CodeStatus) XXX_Size() int {
	return xxx_messageInfo_CodeStatus.Size(m)
}
func (m *CodeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CodeStatus proto.InternalMessageInfo

func (m *CodeStatus) GetCodeStatus() EvaluationStatus {
	if m != nil {
		return m.CodeStatus
	}
	return EvaluationStatus_CORRECT_ANSWER
}

func init() {
	proto.RegisterEnum("evaluator.EvaluationStatus", EvaluationStatus_name, EvaluationStatus_value)
	proto.RegisterType((*CodeStatus)(nil), "evaluator.CodeStatus")
}

func init() { proto.RegisterFile("evaluator/evaluator.proto", fileDescriptor_9a1f69954773e339) }

var fileDescriptor_9a1f69954773e339 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2d, 0x4b, 0xcc,
	0x29, 0x4d, 0x2c, 0xc9, 0x2f, 0xd2, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0xe1, 0x02, 0x52, 0xd2, 0xc5, 0xa5, 0x49, 0xb9, 0x99, 0xc5, 0xc5, 0x99, 0xf9, 0x79, 0xfa, 0x08,
	0x26, 0x44, 0x9d, 0x92, 0x17, 0x17, 0x97, 0x73, 0x7e, 0x4a, 0x6a, 0x70, 0x49, 0x62, 0x49, 0x69,
	0xb1, 0x90, 0x0d, 0x17, 0x77, 0x72, 0x7e, 0x4a, 0x6a, 0x7c, 0x31, 0x98, 0x2b, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x67, 0x24, 0xad, 0x87, 0x30, 0xdc, 0x15, 0xc2, 0xca, 0xcc, 0xcf, 0x83, 0xe8, 0x08,
	0xe2, 0x4a, 0x86, 0xeb, 0xd6, 0xaa, 0xe6, 0x12, 0x40, 0x97, 0x17, 0x12, 0xe2, 0xe2, 0x73, 0xf6,
	0x0f, 0x0a, 0x72, 0x75, 0x0e, 0x89, 0x77, 0xf4, 0x0b, 0x0e, 0x77, 0x0d, 0x12, 0x60, 0x10, 0x12,
	0xe0, 0xe2, 0x09, 0x0f, 0xf2, 0xf7, 0x73, 0x87, 0x89, 0x30, 0x0a, 0x89, 0x73, 0x09, 0x87, 0x78,
	0xfa, 0xba, 0xc6, 0xfb, 0x78, 0xfa, 0x7a, 0x86, 0xc4, 0xbb, 0x46, 0x38, 0xbb, 0xba, 0xba, 0xb8,
	0xba, 0x08, 0x30, 0x09, 0x89, 0x72, 0x09, 0x3a, 0xfb, 0xfb, 0x06, 0x78, 0xfa, 0x38, 0x86, 0x78,
	0xfa, 0xfb, 0xc5, 0xbb, 0x06, 0x05, 0xf9, 0x07, 0x09, 0x30, 0x0b, 0x09, 0x72, 0xf1, 0x06, 0x85,
	0xfa, 0x81, 0xb5, 0x40, 0x84, 0x58, 0x8c, 0x7c, 0xb8, 0x38, 0x5d, 0x61, 0xce, 0x14, 0xb2, 0xe7,
	0xe2, 0x81, 0x72, 0x52, 0x41, 0xbe, 0x13, 0x12, 0xd3, 0x43, 0xf2, 0x78, 0x30, 0x9c, 0x29, 0x25,
	0x8a, 0xe4, 0x35, 0x44, 0x30, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x43, 0xc7, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x3d, 0x77, 0x1f, 0x62, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EvaluatorClient is the client API for Evaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvaluatorClient interface {
	EvaluateCode(ctx context.Context, in *submission.Submission, opts ...grpc.CallOption) (*CodeStatus, error)
}

type evaluatorClient struct {
	cc *grpc.ClientConn
}

func NewEvaluatorClient(cc *grpc.ClientConn) EvaluatorClient {
	return &evaluatorClient{cc}
}

func (c *evaluatorClient) EvaluateCode(ctx context.Context, in *submission.Submission, opts ...grpc.CallOption) (*CodeStatus, error) {
	out := new(CodeStatus)
	err := c.cc.Invoke(ctx, "/evaluator.Evaluator/EvaluateCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaluatorServer is the server API for Evaluator service.
type EvaluatorServer interface {
	EvaluateCode(context.Context, *submission.Submission) (*CodeStatus, error)
}

func RegisterEvaluatorServer(s *grpc.Server, srv EvaluatorServer) {
	s.RegisterService(&_Evaluator_serviceDesc, srv)
}

func _Evaluator_EvaluateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(submission.Submission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaluatorServer).EvaluateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evaluator.Evaluator/EvaluateCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaluatorServer).EvaluateCode(ctx, req.(*submission.Submission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Evaluator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evaluator.Evaluator",
	HandlerType: (*EvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EvaluateCode",
			Handler:    _Evaluator_EvaluateCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evaluator/evaluator.proto",
}