// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package property_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PropertyServiceClient is the client API for PropertyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PropertyServiceClient interface {
	//Create method for adding a new property type into the system
	AddPropertyType(ctx context.Context, in *PropertyType, opts ...grpc.CallOption) (*PropertyType, error)
	//List method for showing all property types in the system
	ListType(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (PropertyService_ListTypeClient, error)
	//Create method for adding a new locality into the system
	AddLocality(ctx context.Context, in *Locality, opts ...grpc.CallOption) (*Locality, error)
	//Delete method for removing an existing locality from the system
	DeleteLocality(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Locality, error)
	//Create method for adding a new property into the system
	CreateProperty(ctx context.Context, in *Property, opts ...grpc.CallOption) (*PropertyState, error)
	//Update property request to modify its current form to another
	UpdateProperty(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Property, error)
	//Delete property request to modify its current form to another
	DeleteProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error)
	//State request to determine active state and status of a property
	StateOfProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error)
	//History request returns all the state transitions a property has had over its lifetime in the system
	HistoryOfProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (PropertyService_HistoryOfPropertyClient, error)
	//Search method is for client request to query for properties that match query
	SearchProperty(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (PropertyService_SearchPropertyClient, error)
	//ListSubscriptions for a particular property
	ListSubscriptions(ctx context.Context, in *SubscriptionListRequest, opts ...grpc.CallOption) (PropertyService_ListSubscriptionsClient, error)
	//AddSubscription for a profile to a property
	AddSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	//Delete subscription of profile to a property
	DeleteSubscription(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error)
}

type propertyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPropertyServiceClient(cc grpc.ClientConnInterface) PropertyServiceClient {
	return &propertyServiceClient{cc}
}

func (c *propertyServiceClient) AddPropertyType(ctx context.Context, in *PropertyType, opts ...grpc.CallOption) (*PropertyType, error) {
	out := new(PropertyType)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/AddPropertyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) ListType(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (PropertyService_ListTypeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PropertyService_ServiceDesc.Streams[0], "/apis.PropertyService/ListType", opts...)
	if err != nil {
		return nil, err
	}
	x := &propertyServiceListTypeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PropertyService_ListTypeClient interface {
	Recv() (*PropertyType, error)
	grpc.ClientStream
}

type propertyServiceListTypeClient struct {
	grpc.ClientStream
}

func (x *propertyServiceListTypeClient) Recv() (*PropertyType, error) {
	m := new(PropertyType)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *propertyServiceClient) AddLocality(ctx context.Context, in *Locality, opts ...grpc.CallOption) (*Locality, error) {
	out := new(Locality)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/AddLocality", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteLocality(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Locality, error) {
	out := new(Locality)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/DeleteLocality", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) CreateProperty(ctx context.Context, in *Property, opts ...grpc.CallOption) (*PropertyState, error) {
	out := new(PropertyState)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/CreateProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) UpdateProperty(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Property, error) {
	out := new(Property)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/UpdateProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error) {
	out := new(PropertyState)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/DeleteProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) StateOfProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error) {
	out := new(PropertyState)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/StateOfProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) HistoryOfProperty(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (PropertyService_HistoryOfPropertyClient, error) {
	stream, err := c.cc.NewStream(ctx, &PropertyService_ServiceDesc.Streams[1], "/apis.PropertyService/HistoryOfProperty", opts...)
	if err != nil {
		return nil, err
	}
	x := &propertyServiceHistoryOfPropertyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PropertyService_HistoryOfPropertyClient interface {
	Recv() (*PropertyState, error)
	grpc.ClientStream
}

type propertyServiceHistoryOfPropertyClient struct {
	grpc.ClientStream
}

func (x *propertyServiceHistoryOfPropertyClient) Recv() (*PropertyState, error) {
	m := new(PropertyState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *propertyServiceClient) SearchProperty(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (PropertyService_SearchPropertyClient, error) {
	stream, err := c.cc.NewStream(ctx, &PropertyService_ServiceDesc.Streams[2], "/apis.PropertyService/SearchProperty", opts...)
	if err != nil {
		return nil, err
	}
	x := &propertyServiceSearchPropertyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PropertyService_SearchPropertyClient interface {
	Recv() (*Property, error)
	grpc.ClientStream
}

type propertyServiceSearchPropertyClient struct {
	grpc.ClientStream
}

func (x *propertyServiceSearchPropertyClient) Recv() (*Property, error) {
	m := new(Property)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *propertyServiceClient) ListSubscriptions(ctx context.Context, in *SubscriptionListRequest, opts ...grpc.CallOption) (PropertyService_ListSubscriptionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PropertyService_ServiceDesc.Streams[3], "/apis.PropertyService/ListSubscriptions", opts...)
	if err != nil {
		return nil, err
	}
	x := &propertyServiceListSubscriptionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PropertyService_ListSubscriptionsClient interface {
	Recv() (*Subscription, error)
	grpc.ClientStream
}

type propertyServiceListSubscriptionsClient struct {
	grpc.ClientStream
}

func (x *propertyServiceListSubscriptionsClient) Recv() (*Subscription, error) {
	m := new(Subscription)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *propertyServiceClient) AddSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	out := new(Subscription)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) DeleteSubscription(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*PropertyState, error) {
	out := new(PropertyState)
	err := c.cc.Invoke(ctx, "/apis.PropertyService/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyServiceServer is the server API for PropertyService service.
// All implementations must embed UnimplementedPropertyServiceServer
// for forward compatibility
type PropertyServiceServer interface {
	//Create method for adding a new property type into the system
	AddPropertyType(context.Context, *PropertyType) (*PropertyType, error)
	//List method for showing all property types in the system
	ListType(*SearchRequest, PropertyService_ListTypeServer) error
	//Create method for adding a new locality into the system
	AddLocality(context.Context, *Locality) (*Locality, error)
	//Delete method for removing an existing locality from the system
	DeleteLocality(context.Context, *RequestID) (*Locality, error)
	//Create method for adding a new property into the system
	CreateProperty(context.Context, *Property) (*PropertyState, error)
	//Update property request to modify its current form to another
	UpdateProperty(context.Context, *UpdateRequest) (*Property, error)
	//Delete property request to modify its current form to another
	DeleteProperty(context.Context, *RequestID) (*PropertyState, error)
	//State request to determine active state and status of a property
	StateOfProperty(context.Context, *RequestID) (*PropertyState, error)
	//History request returns all the state transitions a property has had over its lifetime in the system
	HistoryOfProperty(*RequestID, PropertyService_HistoryOfPropertyServer) error
	//Search method is for client request to query for properties that match query
	SearchProperty(*SearchRequest, PropertyService_SearchPropertyServer) error
	//ListSubscriptions for a particular property
	ListSubscriptions(*SubscriptionListRequest, PropertyService_ListSubscriptionsServer) error
	//AddSubscription for a profile to a property
	AddSubscription(context.Context, *Subscription) (*Subscription, error)
	//Delete subscription of profile to a property
	DeleteSubscription(context.Context, *RequestID) (*PropertyState, error)
	mustEmbedUnimplementedPropertyServiceServer()
}

// UnimplementedPropertyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPropertyServiceServer struct {
}

func (UnimplementedPropertyServiceServer) AddPropertyType(context.Context, *PropertyType) (*PropertyType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPropertyType not implemented")
}
func (UnimplementedPropertyServiceServer) ListType(*SearchRequest, PropertyService_ListTypeServer) error {
	return status.Errorf(codes.Unimplemented, "method ListType not implemented")
}
func (UnimplementedPropertyServiceServer) AddLocality(context.Context, *Locality) (*Locality, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLocality not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteLocality(context.Context, *RequestID) (*Locality, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocality not implemented")
}
func (UnimplementedPropertyServiceServer) CreateProperty(context.Context, *Property) (*PropertyState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProperty not implemented")
}
func (UnimplementedPropertyServiceServer) UpdateProperty(context.Context, *UpdateRequest) (*Property, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProperty not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteProperty(context.Context, *RequestID) (*PropertyState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProperty not implemented")
}
func (UnimplementedPropertyServiceServer) StateOfProperty(context.Context, *RequestID) (*PropertyState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StateOfProperty not implemented")
}
func (UnimplementedPropertyServiceServer) HistoryOfProperty(*RequestID, PropertyService_HistoryOfPropertyServer) error {
	return status.Errorf(codes.Unimplemented, "method HistoryOfProperty not implemented")
}
func (UnimplementedPropertyServiceServer) SearchProperty(*SearchRequest, PropertyService_SearchPropertyServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchProperty not implemented")
}
func (UnimplementedPropertyServiceServer) ListSubscriptions(*SubscriptionListRequest, PropertyService_ListSubscriptionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}
func (UnimplementedPropertyServiceServer) AddSubscription(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (UnimplementedPropertyServiceServer) DeleteSubscription(context.Context, *RequestID) (*PropertyState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedPropertyServiceServer) mustEmbedUnimplementedPropertyServiceServer() {}

// UnsafePropertyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PropertyServiceServer will
// result in compilation errors.
type UnsafePropertyServiceServer interface {
	mustEmbedUnimplementedPropertyServiceServer()
}

func RegisterPropertyServiceServer(s grpc.ServiceRegistrar, srv PropertyServiceServer) {
	s.RegisterService(&PropertyService_ServiceDesc, srv)
}

func _PropertyService_AddPropertyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertyType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).AddPropertyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/AddPropertyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).AddPropertyType(ctx, req.(*PropertyType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_ListType_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PropertyServiceServer).ListType(m, &propertyServiceListTypeServer{stream})
}

type PropertyService_ListTypeServer interface {
	Send(*PropertyType) error
	grpc.ServerStream
}

type propertyServiceListTypeServer struct {
	grpc.ServerStream
}

func (x *propertyServiceListTypeServer) Send(m *PropertyType) error {
	return x.ServerStream.SendMsg(m)
}

func _PropertyService_AddLocality_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Locality)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).AddLocality(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/AddLocality",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).AddLocality(ctx, req.(*Locality))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteLocality_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteLocality(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/DeleteLocality",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteLocality(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_CreateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Property)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).CreateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/CreateProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).CreateProperty(ctx, req.(*Property))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_UpdateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).UpdateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/UpdateProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).UpdateProperty(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/DeleteProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteProperty(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_StateOfProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).StateOfProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/StateOfProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).StateOfProperty(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_HistoryOfProperty_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PropertyServiceServer).HistoryOfProperty(m, &propertyServiceHistoryOfPropertyServer{stream})
}

type PropertyService_HistoryOfPropertyServer interface {
	Send(*PropertyState) error
	grpc.ServerStream
}

type propertyServiceHistoryOfPropertyServer struct {
	grpc.ServerStream
}

func (x *propertyServiceHistoryOfPropertyServer) Send(m *PropertyState) error {
	return x.ServerStream.SendMsg(m)
}

func _PropertyService_SearchProperty_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PropertyServiceServer).SearchProperty(m, &propertyServiceSearchPropertyServer{stream})
}

type PropertyService_SearchPropertyServer interface {
	Send(*Property) error
	grpc.ServerStream
}

type propertyServiceSearchPropertyServer struct {
	grpc.ServerStream
}

func (x *propertyServiceSearchPropertyServer) Send(m *Property) error {
	return x.ServerStream.SendMsg(m)
}

func _PropertyService_ListSubscriptions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriptionListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PropertyServiceServer).ListSubscriptions(m, &propertyServiceListSubscriptionsServer{stream})
}

type PropertyService_ListSubscriptionsServer interface {
	Send(*Subscription) error
	grpc.ServerStream
}

type propertyServiceListSubscriptionsServer struct {
	grpc.ServerStream
}

func (x *propertyServiceListSubscriptionsServer) Send(m *Subscription) error {
	return x.ServerStream.SendMsg(m)
}

func _PropertyService_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).AddSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.PropertyService/DeleteSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).DeleteSubscription(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

// PropertyService_ServiceDesc is the grpc.ServiceDesc for PropertyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PropertyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.PropertyService",
	HandlerType: (*PropertyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPropertyType",
			Handler:    _PropertyService_AddPropertyType_Handler,
		},
		{
			MethodName: "AddLocality",
			Handler:    _PropertyService_AddLocality_Handler,
		},
		{
			MethodName: "DeleteLocality",
			Handler:    _PropertyService_DeleteLocality_Handler,
		},
		{
			MethodName: "CreateProperty",
			Handler:    _PropertyService_CreateProperty_Handler,
		},
		{
			MethodName: "UpdateProperty",
			Handler:    _PropertyService_UpdateProperty_Handler,
		},
		{
			MethodName: "DeleteProperty",
			Handler:    _PropertyService_DeleteProperty_Handler,
		},
		{
			MethodName: "StateOfProperty",
			Handler:    _PropertyService_StateOfProperty_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _PropertyService_AddSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _PropertyService_DeleteSubscription_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListType",
			Handler:       _PropertyService_ListType_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HistoryOfProperty",
			Handler:       _PropertyService_HistoryOfProperty_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchProperty",
			Handler:       _PropertyService_SearchProperty_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListSubscriptions",
			Handler:       _PropertyService_ListSubscriptions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "property.proto",
}
