// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pilot.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendTaskRequest struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=directive" json:"directive,omitempty"`
}

func (m *SendTaskRequest) Reset()                    { *m = SendTaskRequest{} }
func (m *SendTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTaskRequest) ProtoMessage()               {}
func (*SendTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SendTaskRequest) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *SendTaskRequest) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *SendTaskRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type SendTaskResponse struct {
	TaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *SendTaskResponse) Reset()                    { *m = SendTaskResponse{} }
func (m *SendTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*SendTaskResponse) ProtoMessage()               {}
func (*SendTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SendTaskResponse) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type GetTaskStatusRequest struct {
	TaskId []string `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *GetTaskStatusRequest) Reset()                    { *m = GetTaskStatusRequest{} }
func (m *GetTaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTaskStatusRequest) ProtoMessage()               {}
func (*GetTaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *GetTaskStatusRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type TaskStatus struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TaskStatus *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=task_status,json=taskStatus" json:"task_status,omitempty"`
}

func (m *TaskStatus) Reset()                    { *m = TaskStatus{} }
func (m *TaskStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()               {}
func (*TaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TaskStatus) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskStatus) GetTaskStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskStatus
	}
	return nil
}

type GetTaskStatusResponse struct {
	TotalCount    uint32        `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskStatusSet []*TaskStatus `protobuf:"bytes,2,rep,name=task_status_set,json=taskStatusSet" json:"task_status_set,omitempty"`
}

func (m *GetTaskStatusResponse) Reset()                    { *m = GetTaskStatusResponse{} }
func (m *GetTaskStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTaskStatusResponse) ProtoMessage()               {}
func (*GetTaskStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *GetTaskStatusResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *GetTaskStatusResponse) GetTaskStatusSet() []*TaskStatus {
	if m != nil {
		return m.TaskStatusSet
	}
	return nil
}

func init() {
	proto.RegisterType((*SendTaskRequest)(nil), "openpitrix.SendTaskRequest")
	proto.RegisterType((*SendTaskResponse)(nil), "openpitrix.SendTaskResponse")
	proto.RegisterType((*GetTaskStatusRequest)(nil), "openpitrix.GetTaskStatusRequest")
	proto.RegisterType((*TaskStatus)(nil), "openpitrix.TaskStatus")
	proto.RegisterType((*GetTaskStatusResponse)(nil), "openpitrix.GetTaskStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PilotManager service

type PilotManagerClient interface {
	SendTask(ctx context.Context, in *SendTaskRequest, opts ...grpc.CallOption) (*SendTaskResponse, error)
	GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error)
}

type pilotManagerClient struct {
	cc *grpc.ClientConn
}

func NewPilotManagerClient(cc *grpc.ClientConn) PilotManagerClient {
	return &pilotManagerClient{cc}
}

func (c *pilotManagerClient) SendTask(ctx context.Context, in *SendTaskRequest, opts ...grpc.CallOption) (*SendTaskResponse, error) {
	out := new(SendTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotManager/SendTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotManagerClient) GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error) {
	out := new(GetTaskStatusResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotManager/GetTaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PilotManager service

type PilotManagerServer interface {
	SendTask(context.Context, *SendTaskRequest) (*SendTaskResponse, error)
	GetTaskStatus(context.Context, *GetTaskStatusRequest) (*GetTaskStatusResponse, error)
}

func RegisterPilotManagerServer(s *grpc.Server, srv PilotManagerServer) {
	s.RegisterService(&_PilotManager_serviceDesc, srv)
}

func _PilotManager_SendTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotManagerServer).SendTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotManager/SendTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotManagerServer).SendTask(ctx, req.(*SendTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotManager_GetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotManagerServer).GetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotManager/GetTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotManagerServer).GetTaskStatus(ctx, req.(*GetTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PilotManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.PilotManager",
	HandlerType: (*PilotManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTask",
			Handler:    _PilotManager_SendTask_Handler,
		},
		{
			MethodName: "GetTaskStatus",
			Handler:    _PilotManager_GetTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pilot.proto",
}

func init() { proto.RegisterFile("pilot.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0x51, 0xc8, 0xb8, 0x51, 0xca, 0x8a, 0x42, 0x14, 0x22, 0x30, 0x46, 0x42, 0x51,
	0xa1, 0x76, 0x1b, 0xc4, 0x25, 0x12, 0x48, 0x81, 0x03, 0xf4, 0x80, 0x88, 0x12, 0xc4, 0x81, 0x4b,
	0xb4, 0x71, 0x06, 0xb3, 0x69, 0xb4, 0xbb, 0xdd, 0x9d, 0xb4, 0x39, 0x70, 0xea, 0x27, 0x84, 0x4f,
	0xe2, 0x13, 0xf8, 0x05, 0x3e, 0x82, 0x23, 0xf2, 0x26, 0xc1, 0xa6, 0x8d, 0x50, 0x04, 0x27, 0x4b,
	0x33, 0x6f, 0xde, 0x7b, 0xf3, 0x3c, 0x0b, 0xbe, 0x16, 0x13, 0x45, 0x91, 0x36, 0x8a, 0x14, 0x03,
	0xa5, 0x51, 0x6a, 0x41, 0x46, 0xcc, 0xea, 0xf7, 0x52, 0xa5, 0xd2, 0x09, 0xc6, 0xae, 0x33, 0x9c,
	0x7e, 0x8a, 0xcf, 0x0d, 0xd7, 0x1a, 0x8d, 0x5d, 0x60, 0xeb, 0x8d, 0x65, 0x9f, 0x6b, 0x11, 0x73,
	0x29, 0x15, 0x71, 0x12, 0x4a, 0xae, 0xba, 0x4f, 0xdc, 0x27, 0x39, 0x48, 0x51, 0x1e, 0xd8, 0x73,
	0x9e, 0xa6, 0x68, 0x62, 0xa5, 0x1d, 0xe2, 0x2a, 0x3a, 0xfc, 0xe6, 0x41, 0xb5, 0x8f, 0x72, 0xf4,
	0x9e, 0xdb, 0x93, 0x1e, 0x9e, 0x4e, 0xd1, 0x12, 0x7b, 0x06, 0xd7, 0x89, 0xdb, 0x93, 0x81, 0x18,
	0xd5, 0xbc, 0xc0, 0x6b, 0xfa, 0xad, 0x46, 0xb4, 0x50, 0x8c, 0x56, 0x8e, 0xa2, 0x3e, 0x19, 0x21,
	0xd3, 0x0f, 0x7c, 0x32, 0xc5, 0xde, 0x76, 0x06, 0x3e, 0x1e, 0xb1, 0xe7, 0xe0, 0xbb, 0x31, 0x9e,
	0x64, 0x02, 0xb5, 0xd2, 0x06, 0xa3, 0x90, 0x0d, 0x74, 0x1c, 0x9e, 0xb5, 0xa1, 0x3c, 0x12, 0x06,
	0x13, 0x12, 0x67, 0x58, 0xdb, 0xda, 0x60, 0x38, 0x87, 0x87, 0xc7, 0xb0, 0x9b, 0x2f, 0x61, 0xb5,
	0x92, 0x16, 0xff, 0x71, 0x8b, 0x30, 0x86, 0x5b, 0xaf, 0x91, 0x32, 0xa6, 0x3e, 0x71, 0x9a, 0xda,
	0x55, 0x28, 0x77, 0x8a, 0x74, 0x5b, 0xcd, 0xf2, 0xef, 0x81, 0x0b, 0x0f, 0x20, 0x87, 0xff, 0x6f,
	0x78, 0xd6, 0xb1, 0x6c, 0x1e, 0xde, 0x42, 0x35, 0x9c, 0xc1, 0xde, 0x25, 0xd7, 0xcb, 0x14, 0xee,
	0x83, 0x4f, 0x8a, 0xf8, 0x64, 0x90, 0xa8, 0xa9, 0x24, 0x67, 0xa9, 0xd2, 0x03, 0x57, 0x7a, 0x95,
	0x55, 0xd8, 0x0b, 0xa8, 0x16, 0x84, 0x07, 0x16, 0xa9, 0x56, 0x0a, 0xb6, 0x9a, 0x7e, 0xeb, 0x76,
	0x94, 0x9f, 0x64, 0x54, 0x60, 0xae, 0xe4, 0xb2, 0x7d, 0xa4, 0xd6, 0x4f, 0x0f, 0x76, 0xba, 0xd9,
	0x21, 0xbf, 0xe5, 0x92, 0xa7, 0x68, 0xd8, 0x18, 0x6e, 0xac, 0xfe, 0x05, 0xbb, 0x5b, 0xe4, 0xb8,
	0x74, 0x66, 0xf5, 0xc6, 0xfa, 0xe6, 0xc2, 0x78, 0xf8, 0x70, 0xde, 0xf1, 0x59, 0xd9, 0xa2, 0x1c,
	0x05, 0x99, 0xe0, 0xc5, 0xf7, 0x1f, 0x5f, 0x4b, 0xd5, 0x10, 0xe2, 0xb3, 0xa3, 0xd8, 0xbd, 0x1b,
	0xdb, 0xf6, 0xf6, 0xd9, 0x17, 0xa8, 0xfc, 0xb1, 0x36, 0x0b, 0x8a, 0x9c, 0xeb, 0xfe, 0x63, 0xfd,
	0xc1, 0x5f, 0x10, 0x4b, 0xe9, 0x47, 0xf3, 0xce, 0x4d, 0x56, 0x4d, 0x91, 0x9c, 0x72, 0xb0, 0x48,
	0xc6, 0x19, 0xd8, 0x61, 0x05, 0x03, 0x2f, 0x67, 0xf3, 0xce, 0x29, 0x7b, 0x03, 0xec, 0x9d, 0x46,
	0xd9, 0x15, 0x06, 0xc5, 0x2c, 0xe8, 0x1a, 0x35, 0xc6, 0x84, 0xc2, 0xc7, 0xeb, 0xaa, 0x6c, 0xef,
	0x33, 0x91, 0xb6, 0xed, 0x38, 0x2e, 0x58, 0x10, 0xaa, 0x75, 0xed, 0x30, 0x3a, 0x8c, 0x8e, 0xf6,
	0x3d, 0xaf, 0xb5, 0xcb, 0xb5, 0x9e, 0x88, 0xc4, 0xbd, 0xd0, 0x78, 0x6c, 0x95, 0x6c, 0x5f, 0xa9,
	0x7c, 0x2c, 0xe9, 0xe1, 0x70, 0xdb, 0x1d, 0xc4, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47,
	0x15, 0x53, 0x58, 0x43, 0x04, 0x00, 0x00,
}
