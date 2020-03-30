// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Event struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Duration             int32                `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Typeduration         int32                `protobuf:"varint,5,opt,name=typeduration,proto3" json:"typeduration,omitempty"`
	Title                string               `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Note                 string               `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Event) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetTypeduration() int32 {
	if m != nil {
		return m.Typeduration
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

type Events struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Events.Unmarshal(m, b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Events.Marshal(b, m, deterministic)
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return xxx_messageInfo_Events.Size(m)
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Count struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterType((*Events)(nil), "proto.Events")
	proto.RegisterType((*Id)(nil), "proto.Id")
	proto.RegisterType((*Count)(nil), "proto.Count")
}

func init() {
	proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2)
}

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xdd, 0x8a, 0x9b, 0x40,
	0x18, 0x55, 0x37, 0xba, 0xfa, 0x99, 0x2e, 0x65, 0x58, 0x8a, 0x58, 0x4a, 0x65, 0x68, 0xc1, 0x2b,
	0x17, 0xdc, 0x96, 0xf6, 0x76, 0xd9, 0x86, 0x65, 0x7b, 0x29, 0x79, 0x01, 0xe3, 0x7c, 0x4d, 0x04,
	0x75, 0x44, 0x3f, 0x1b, 0x72, 0xdf, 0xe7, 0xea, 0xb3, 0x95, 0xcc, 0x68, 0xda, 0x24, 0x84, 0xb2,
	0x57, 0xce, 0x99, 0xf3, 0xe3, 0xe1, 0x0c, 0xdc, 0x14, 0x79, 0x85, 0x8d, 0xc8, 0xbb, 0xa4, 0xed,
	0x24, 0x49, 0x66, 0xab, 0x4f, 0xf8, 0x76, 0x2d, 0xe5, 0xba, 0xc2, 0x3b, 0x85, 0x56, 0xc3, 0x8f,
	0x3b, 0xac, 0x5b, 0xda, 0x69, 0x4d, 0xf8, 0xfe, 0x94, 0xa4, 0xb2, 0xc6, 0x9e, 0xf2, 0xba, 0xd5,
	0x02, 0xfe, 0xcb, 0x02, 0x7b, 0xf1, 0x13, 0x1b, 0x62, 0x37, 0x60, 0x95, 0x22, 0x30, 0x23, 0x33,
	0xb6, 0x33, 0xab, 0x14, 0xec, 0x2b, 0x78, 0x3d, 0xe5, 0x1d, 0x2d, 0xcb, 0x1a, 0x03, 0x2b, 0x32,
	0x63, 0x3f, 0x0d, 0x13, 0x1d, 0x97, 0x4c, 0x71, 0xc9, 0x72, 0x8a, 0xcb, 0xfe, 0x8a, 0xd9, 0x27,
	0xb8, 0xc6, 0x46, 0x28, 0xdf, 0xd5, 0x7f, 0x7d, 0x93, 0x94, 0x85, 0xe0, 0x8a, 0xa1, 0xcb, 0xa9,
	0x94, 0x4d, 0x30, 0x53, 0x2d, 0x0e, 0x98, 0x71, 0x98, 0xd3, 0xae, 0xc5, 0x03, 0x6f, 0x2b, 0xfe,
	0xe8, 0x8e, 0xdd, 0x82, 0x4d, 0x25, 0x55, 0x18, 0x38, 0x91, 0x19, 0x7b, 0x99, 0x06, 0x8c, 0xc1,
	0xac, 0x91, 0x84, 0xc1, 0xb5, 0xba, 0x54, 0xe7, 0xef, 0x33, 0xd7, 0x7d, 0xed, 0x65, 0xfe, 0x46,
	0x6e, 0xeb, 0xa1, 0xd8, 0x6c, 0xf3, 0xae, 0xe6, 0x09, 0x38, 0x6a, 0x85, 0x9e, 0x7d, 0x00, 0x07,
	0xd5, 0x29, 0x30, 0xa3, 0xab, 0xd8, 0x4f, 0xe7, 0xba, 0x74, 0xa2, 0xe8, 0x6c, 0xe4, 0xf8, 0x2d,
	0x58, 0xcf, 0xe2, 0x74, 0x32, 0xfe, 0x0e, 0xec, 0x47, 0x39, 0x34, 0xb4, 0xef, 0x52, 0xec, 0x0f,
	0x23, 0xa7, 0x41, 0xfa, 0xdb, 0x02, 0xf7, 0x71, 0x7c, 0x43, 0x96, 0x82, 0xfb, 0x20, 0x84, 0x9e,
	0xfe, 0xe8, 0x1f, 0xe1, 0x9b, 0xb3, 0xb5, 0x16, 0xfb, 0x17, 0xe5, 0x06, 0xfb, 0x08, 0xee, 0x13,
	0x92, 0xf6, 0x78, 0xa3, 0xe7, 0x59, 0x84, 0x47, 0x76, 0x6e, 0xb0, 0x14, 0xfc, 0x6f, 0x58, 0x21,
	0xe1, 0x99, 0xf2, 0x72, 0xf4, 0x3d, 0x78, 0x0b, 0x51, 0xd2, 0xcb, 0xfa, 0x7c, 0x81, 0xf9, 0x13,
	0xd2, 0x43, 0x55, 0x8d, 0xdb, 0x5d, 0x50, 0x86, 0xaf, 0xfe, 0xcd, 0xeb, 0xb9, 0xc1, 0x3e, 0x83,
	0xaf, 0x86, 0xca, 0xb0, 0x90, 0x9d, 0xb8, 0xe8, 0x9b, 0x7a, 0x28, 0x2d, 0x37, 0x56, 0x8e, 0x82,
	0xf7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x48, 0xac, 0x4c, 0x0a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
	GetEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error)
	DeleteEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error)
	EditEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAllEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Events, error)
	CountRecord(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Count, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Calendar/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/proto.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) EditEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Calendar/EditEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetAllEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/proto.Calendar/GetAllEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) CountRecord(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/proto.Calendar/CountRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	AddEvent(context.Context, *Event) (*empty.Empty, error)
	GetEvent(context.Context, *Id) (*Event, error)
	DeleteEvent(context.Context, *Id) (*empty.Empty, error)
	EditEvent(context.Context, *Event) (*empty.Empty, error)
	GetAllEvents(context.Context, *empty.Empty) (*Events, error)
	CountRecord(context.Context, *empty.Empty) (*Count, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) AddEvent(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *Id) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) DeleteEvent(ctx context.Context, req *Id) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedCalendarServer) EditEvent(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEvent not implemented")
}
func (*UnimplementedCalendarServer) GetAllEvents(ctx context.Context, req *empty.Empty) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (*UnimplementedCalendarServer) CountRecord(ctx context.Context, req *empty.Empty) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountRecord not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_EditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).EditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/EditEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).EditEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/GetAllEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetAllEvents(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_CountRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CountRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calendar/CountRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CountRecord(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
		{
			MethodName: "EditEvent",
			Handler:    _Calendar_EditEvent_Handler,
		},
		{
			MethodName: "GetAllEvents",
			Handler:    _Calendar_GetAllEvents_Handler,
		},
		{
			MethodName: "CountRecord",
			Handler:    _Calendar_CountRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}