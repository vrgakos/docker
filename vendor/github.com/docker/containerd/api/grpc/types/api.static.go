package types

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	GetServerVersion(ctx context.Context, in *GetServerVersionRequest, opts ...grpc.CallOption) (*GetServerVersionResponse, error)
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*UpdateContainerResponse, error)
	Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error)
	UpdateProcess(ctx context.Context, in *UpdateProcessRequest, opts ...grpc.CallOption) (*UpdateProcessResponse, error)
	AddProcess(ctx context.Context, in *AddProcessRequest, opts ...grpc.CallOption) (*AddProcessResponse, error)
	CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error)
	DeleteCheckpoint(ctx context.Context, in *DeleteCheckpointRequest, opts ...grpc.CallOption) (*DeleteCheckpointResponse, error)
	ListCheckpoint(ctx context.Context, in *ListCheckpointRequest, opts ...grpc.CallOption) (*ListCheckpointResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (API_EventsClient, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetServerVersion(ctx context.Context, in *GetServerVersionRequest, opts ...grpc.CallOption) (*GetServerVersionResponse, error) {
	out := new(GetServerVersionResponse)
	err := grpc.Invoke(ctx, "/types.API/GetServerVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error) {
	out := new(CreateContainerResponse)
	err := grpc.Invoke(ctx, "/types.API/CreateContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (*UpdateContainerResponse, error) {
	out := new(UpdateContainerResponse)
	err := grpc.Invoke(ctx, "/types.API/UpdateContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (*SignalResponse, error) {
	out := new(SignalResponse)
	err := grpc.Invoke(ctx, "/types.API/Signal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateProcess(ctx context.Context, in *UpdateProcessRequest, opts ...grpc.CallOption) (*UpdateProcessResponse, error) {
	out := new(UpdateProcessResponse)
	err := grpc.Invoke(ctx, "/types.API/UpdateProcess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) AddProcess(ctx context.Context, in *AddProcessRequest, opts ...grpc.CallOption) (*AddProcessResponse, error) {
	out := new(AddProcessResponse)
	err := grpc.Invoke(ctx, "/types.API/AddProcess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error) {
	out := new(CreateCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/CreateCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteCheckpoint(ctx context.Context, in *DeleteCheckpointRequest, opts ...grpc.CallOption) (*DeleteCheckpointResponse, error) {
	out := new(DeleteCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/DeleteCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCheckpoint(ctx context.Context, in *ListCheckpointRequest, opts ...grpc.CallOption) (*ListCheckpointResponse, error) {
	out := new(ListCheckpointResponse)
	err := grpc.Invoke(ctx, "/types.API/ListCheckpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := grpc.Invoke(ctx, "/types.API/State", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (API_EventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/types.API/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_EventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type aPIEventsClient struct {
	grpc.ClientStream
}

func (x *aPIEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := grpc.Invoke(ctx, "/types.API/Stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	GetServerVersion(context.Context, *GetServerVersionRequest) (*GetServerVersionResponse, error)
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
	UpdateContainer(context.Context, *UpdateContainerRequest) (*UpdateContainerResponse, error)
	Signal(context.Context, *SignalRequest) (*SignalResponse, error)
	UpdateProcess(context.Context, *UpdateProcessRequest) (*UpdateProcessResponse, error)
	AddProcess(context.Context, *AddProcessRequest) (*AddProcessResponse, error)
	CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CreateCheckpointResponse, error)
	DeleteCheckpoint(context.Context, *DeleteCheckpointRequest) (*DeleteCheckpointResponse, error)
	ListCheckpoint(context.Context, *ListCheckpointRequest) (*ListCheckpointResponse, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
	Events(*EventsRequest, API_EventsServer) error
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetServerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetServerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/GetServerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetServerVersion(ctx, req.(*GetServerVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateContainer(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/UpdateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateContainer(ctx, req.(*UpdateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Signal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Signal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/Signal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Signal(ctx, req.(*SignalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/UpdateProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateProcess(ctx, req.(*UpdateProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_AddProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).AddProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/AddProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).AddProcess(ctx, req.(*AddProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/CreateCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateCheckpoint(ctx, req.(*CreateCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/DeleteCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteCheckpoint(ctx, req.(*DeleteCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/ListCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListCheckpoint(ctx, req.(*ListCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Events(m, &aPIEventsServer{stream})
}

type API_EventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type aPIEventsServer struct {
	grpc.ServerStream
}

func (x *aPIEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _API_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.API/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerVersion",
			Handler:    _API_GetServerVersion_Handler,
		},
		{
			MethodName: "CreateContainer",
			Handler:    _API_CreateContainer_Handler,
		},
		{
			MethodName: "UpdateContainer",
			Handler:    _API_UpdateContainer_Handler,
		},
		{
			MethodName: "Signal",
			Handler:    _API_Signal_Handler,
		},
		{
			MethodName: "UpdateProcess",
			Handler:    _API_UpdateProcess_Handler,
		},
		{
			MethodName: "AddProcess",
			Handler:    _API_AddProcess_Handler,
		},
		{
			MethodName: "CreateCheckpoint",
			Handler:    _API_CreateCheckpoint_Handler,
		},
		{
			MethodName: "DeleteCheckpoint",
			Handler:    _API_DeleteCheckpoint_Handler,
		},
		{
			MethodName: "ListCheckpoint",
			Handler:    _API_ListCheckpoint_Handler,
		},
		{
			MethodName: "State",
			Handler:    _API_State_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _API_Stats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _API_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}