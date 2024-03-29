// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd3, 0xcd, 0x4a, 0x2b, 0x41,
	0x10, 0x05, 0xe0, 0xbb, 0xb9, 0xe1, 0xd2, 0x37, 0x89, 0xd8, 0x0b, 0x35, 0x63, 0xe2, 0xdf, 0x03,
	0x64, 0xa1, 0x6b, 0x57, 0x11, 0x67, 0x13, 0x50, 0x02, 0x82, 0xe0, 0x6a, 0x24, 0x85, 0x04, 0xc2,
	0x4c, 0xec, 0x6e, 0xf3, 0xa2, 0xbe, 0x90, 0x24, 0x35, 0xdd, 0xa9, 0xaa, 0xae, 0x71, 0x25, 0x9e,
	0x53, 0xf9, 0x60, 0x0e, 0x33, 0xa6, 0xef, 0xc1, 0x6d, 0xc1, 0x4d, 0x37, 0xae, 0x09, 0x8d, 0xfd,
	0xbb, 0xff, 0x53, 0x0c, 0x1c, 0x7c, 0x7e, 0x81, 0x0f, 0x98, 0x16, 0x43, 0x07, 0x7e, 0xd3, 0xd4,
	0x1e, 0xf0, 0xff, 0xdb, 0xef, 0xbe, 0xe9, 0x3d, 0x37, 0xbe, 0x5a, 0x56, 0x76, 0x66, 0x4c, 0x09,
	0x61, 0xbe, 0xf2, 0x61, 0x55, 0x7f, 0xd8, 0x33, 0x3c, 0x98, 0x1e, 0xa2, 0x05, 0x42, 0xc5, 0x48,
	0x69, 0xd0, 0xbc, 0xf9, 0x63, 0xe7, 0x66, 0xf0, 0xb2, 0x59, 0x56, 0x01, 0xa2, 0x73, 0xde, 0x5e,
	0xb3, 0x34, 0x52, 0x63, 0xbd, 0xa4, 0xda, 0xcc, 0x81, 0xa2, 0xb1, 0x54, 0x6a, 0xa2, 0xa4, 0xda,
	0x03, 0xac, 0x21, 0xd7, 0x58, 0x2a, 0x35, 0x51, 0x26, 0xed, 0xc9, 0x0c, 0x4b, 0x08, 0x0b, 0xd8,
	0x8d, 0x5e, 0x85, 0x55, 0x53, 0xdb, 0xf1, 0x61, 0x18, 0x12, 0x47, 0x6f, 0xd2, 0xd1, 0x26, 0xf0,
	0xd5, 0x1c, 0xe3, 0x0e, 0xd4, 0xbc, 0x64, 0x0b, 0x29, 0xec, 0x55, 0xf7, 0x01, 0x95, 0x71, 0x13,
	0x4d, 0xce, 0x1a, 0x29, 0x2b, 0x07, 0x54, 0xc6, 0x7d, 0x34, 0x39, 0x6b, 0xa4, 0xac, 0x1c, 0x24,
	0xf9, 0xde, 0xfc, 0x2b, 0x21, 0x94, 0xbb, 0x7b, 0x7b, 0x72, 0x98, 0x6e, 0x1f, 0x44, 0xe7, 0x34,
	0xcb, 0xd3, 0xcf, 0x1f, 0xcd, 0x7f, 0x5c, 0x04, 0x85, 0x11, 0x5b, 0x89, 0x21, 0x85, 0x56, 0x51,
	0x07, 0x9f, 0x9f, 0x3b, 0x24, 0x93, 0x0e, 0xab, 0xa8, 0x83, 0x4f, 0xcb, 0x1d, 0x92, 0x49, 0x87,
	0x55, 0xd4, 0x29, 0x21, 0xcc, 0xaa, 0x35, 0xd4, 0xcb, 0xca, 0x59, 0xf2, 0x2d, 0xc6, 0x4c, 0x3a,
	0xac, 0xa2, 0x6f, 0x2f, 0x3e, 0x70, 0xa2, 0xf8, 0xb7, 0x28, 0xb5, 0x49, 0x47, 0x4b, 0x41, 0x7c,
	0xf2, 0x0c, 0xe4, 0xb1, 0x04, 0x65, 0x4b, 0x41, 0x9c, 0x20, 0x03, 0x79, 0x2c, 0x41, 0xd9, 0x26,
	0x70, 0x61, 0x8e, 0x76, 0x5b, 0x34, 0xf5, 0x16, 0x9c, 0xc7, 0x37, 0x95, 0x7c, 0x93, 0x34, 0x8f,
	0xe4, 0x45, 0x57, 0x9d, 0xcc, 0x37, 0x63, 0xdb, 0x45, 0x28, 0xcb, 0xbf, 0x49, 0x4d, 0xbe, 0xfe,
	0xe5, 0x82, 0xe2, 0xed, 0x3a, 0x1a, 0x9e, 0x57, 0x12, 0xd7, 0x2e, 0x28, 0xde, 0x2e, 0xa5, 0xe1,
	0x79, 0x25, 0x71, 0xed, 0x22, 0xe2, 0xef, 0xbd, 0xfd, 0xcd, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x73, 0x0b, 0x8b, 0x92, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PosadaClient is the client API for Posada service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PosadaClient interface {
	// Returns listings based on request parameters.
	GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*GetListingResponse, error)
	// Updates a listing.
	UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*UpdateListingResponse, error)
	// Creates a listing.
	CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*CreateListingResponse, error)
	// Deletes a listing.
	DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*DeleteListingResponse, error)
	// Returns reservations based on request parameters.
	GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error)
	// Updates a Reservation.
	UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error)
	// Creates a Reservation.
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	// Deletes a Reservation.
	DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	// Returns guests based on request parameters.
	GetGuest(ctx context.Context, in *GetGuestRequest, opts ...grpc.CallOption) (*GetGuestResponse, error)
	// Updates a Guest.
	UpdateGuest(ctx context.Context, in *UpdateGuestRequest, opts ...grpc.CallOption) (*UpdateGuestResponse, error)
	// Creates a Guest.
	CreateGuest(ctx context.Context, in *CreateGuestRequest, opts ...grpc.CallOption) (*CreateGuestResponse, error)
	// Deletes a Guest.
	DeleteGuest(ctx context.Context, in *DeleteGuestRequest, opts ...grpc.CallOption) (*DeleteGuestResponse, error)
	// Returns calendars based on request parameters.
	GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error)
	// Updates a Calendar.
	UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*UpdateCalendarResponse, error)
	// Creates a Calendar.
	CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*CreateCalendarResponse, error)
	// Deletes a Calendar.
	DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*DeleteCalendarResponse, error)
	// Returns conversation based on request parameters.
	GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error)
	// Updates a Conversation.
	UpdateConversation(ctx context.Context, in *UpdateConversationRequest, opts ...grpc.CallOption) (*UpdateConversationResponse, error)
	// Creates a Conversation.
	CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*CreateConversationResponse, error)
	// Deletes a Conversation.
	DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error)
}

type posadaClient struct {
	cc *grpc.ClientConn
}

func NewPosadaClient(cc *grpc.ClientConn) PosadaClient {
	return &posadaClient{cc}
}

func (c *posadaClient) GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*GetListingResponse, error) {
	out := new(GetListingResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/GetListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*UpdateListingResponse, error) {
	out := new(UpdateListingResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/UpdateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*CreateListingResponse, error) {
	out := new(CreateListingResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/CreateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*DeleteListingResponse, error) {
	out := new(DeleteListingResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/DeleteListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error) {
	out := new(GetReservationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/GetReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error) {
	out := new(UpdateReservationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) GetGuest(ctx context.Context, in *GetGuestRequest, opts ...grpc.CallOption) (*GetGuestResponse, error) {
	out := new(GetGuestResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/GetGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) UpdateGuest(ctx context.Context, in *UpdateGuestRequest, opts ...grpc.CallOption) (*UpdateGuestResponse, error) {
	out := new(UpdateGuestResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/UpdateGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) CreateGuest(ctx context.Context, in *CreateGuestRequest, opts ...grpc.CallOption) (*CreateGuestResponse, error) {
	out := new(CreateGuestResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/CreateGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) DeleteGuest(ctx context.Context, in *DeleteGuestRequest, opts ...grpc.CallOption) (*DeleteGuestResponse, error) {
	out := new(DeleteGuestResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/DeleteGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) GetCalendar(ctx context.Context, in *GetCalendarRequest, opts ...grpc.CallOption) (*GetCalendarResponse, error) {
	out := new(GetCalendarResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/GetCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*UpdateCalendarResponse, error) {
	out := new(UpdateCalendarResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/UpdateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*CreateCalendarResponse, error) {
	out := new(CreateCalendarResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/CreateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*DeleteCalendarResponse, error) {
	out := new(DeleteCalendarResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/DeleteCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error) {
	out := new(GetConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/GetConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) UpdateConversation(ctx context.Context, in *UpdateConversationRequest, opts ...grpc.CallOption) (*UpdateConversationResponse, error) {
	out := new(UpdateConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/UpdateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*CreateConversationResponse, error) {
	out := new(CreateConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/CreateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posadaClient) DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error) {
	out := new(DeleteConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.Posada/DeleteConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosadaServer is the server API for Posada service.
type PosadaServer interface {
	// Returns listings based on request parameters.
	GetListing(context.Context, *GetListingRequest) (*GetListingResponse, error)
	// Updates a listing.
	UpdateListing(context.Context, *UpdateListingRequest) (*UpdateListingResponse, error)
	// Creates a listing.
	CreateListing(context.Context, *CreateListingRequest) (*CreateListingResponse, error)
	// Deletes a listing.
	DeleteListing(context.Context, *DeleteListingRequest) (*DeleteListingResponse, error)
	// Returns reservations based on request parameters.
	GetReservation(context.Context, *GetReservationRequest) (*GetReservationResponse, error)
	// Updates a Reservation.
	UpdateReservation(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error)
	// Creates a Reservation.
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	// Deletes a Reservation.
	DeleteReservation(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	// Returns guests based on request parameters.
	GetGuest(context.Context, *GetGuestRequest) (*GetGuestResponse, error)
	// Updates a Guest.
	UpdateGuest(context.Context, *UpdateGuestRequest) (*UpdateGuestResponse, error)
	// Creates a Guest.
	CreateGuest(context.Context, *CreateGuestRequest) (*CreateGuestResponse, error)
	// Deletes a Guest.
	DeleteGuest(context.Context, *DeleteGuestRequest) (*DeleteGuestResponse, error)
	// Returns calendars based on request parameters.
	GetCalendar(context.Context, *GetCalendarRequest) (*GetCalendarResponse, error)
	// Updates a Calendar.
	UpdateCalendar(context.Context, *UpdateCalendarRequest) (*UpdateCalendarResponse, error)
	// Creates a Calendar.
	CreateCalendar(context.Context, *CreateCalendarRequest) (*CreateCalendarResponse, error)
	// Deletes a Calendar.
	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*DeleteCalendarResponse, error)
	// Returns conversation based on request parameters.
	GetConversation(context.Context, *GetConversationRequest) (*GetConversationResponse, error)
	// Updates a Conversation.
	UpdateConversation(context.Context, *UpdateConversationRequest) (*UpdateConversationResponse, error)
	// Creates a Conversation.
	CreateConversation(context.Context, *CreateConversationRequest) (*CreateConversationResponse, error)
	// Deletes a Conversation.
	DeleteConversation(context.Context, *DeleteConversationRequest) (*DeleteConversationResponse, error)
}

// UnimplementedPosadaServer can be embedded to have forward compatible implementations.
type UnimplementedPosadaServer struct {
}

func (*UnimplementedPosadaServer) GetListing(ctx context.Context, req *GetListingRequest) (*GetListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListing not implemented")
}
func (*UnimplementedPosadaServer) UpdateListing(ctx context.Context, req *UpdateListingRequest) (*UpdateListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListing not implemented")
}
func (*UnimplementedPosadaServer) CreateListing(ctx context.Context, req *CreateListingRequest) (*CreateListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateListing not implemented")
}
func (*UnimplementedPosadaServer) DeleteListing(ctx context.Context, req *DeleteListingRequest) (*DeleteListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteListing not implemented")
}
func (*UnimplementedPosadaServer) GetReservation(ctx context.Context, req *GetReservationRequest) (*GetReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (*UnimplementedPosadaServer) UpdateReservation(ctx context.Context, req *UpdateReservationRequest) (*UpdateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (*UnimplementedPosadaServer) CreateReservation(ctx context.Context, req *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (*UnimplementedPosadaServer) DeleteReservation(ctx context.Context, req *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (*UnimplementedPosadaServer) GetGuest(ctx context.Context, req *GetGuestRequest) (*GetGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuest not implemented")
}
func (*UnimplementedPosadaServer) UpdateGuest(ctx context.Context, req *UpdateGuestRequest) (*UpdateGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuest not implemented")
}
func (*UnimplementedPosadaServer) CreateGuest(ctx context.Context, req *CreateGuestRequest) (*CreateGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGuest not implemented")
}
func (*UnimplementedPosadaServer) DeleteGuest(ctx context.Context, req *DeleteGuestRequest) (*DeleteGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGuest not implemented")
}
func (*UnimplementedPosadaServer) GetCalendar(ctx context.Context, req *GetCalendarRequest) (*GetCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCalendar not implemented")
}
func (*UnimplementedPosadaServer) UpdateCalendar(ctx context.Context, req *UpdateCalendarRequest) (*UpdateCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCalendar not implemented")
}
func (*UnimplementedPosadaServer) CreateCalendar(ctx context.Context, req *CreateCalendarRequest) (*CreateCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCalendar not implemented")
}
func (*UnimplementedPosadaServer) DeleteCalendar(ctx context.Context, req *DeleteCalendarRequest) (*DeleteCalendarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCalendar not implemented")
}
func (*UnimplementedPosadaServer) GetConversation(ctx context.Context, req *GetConversationRequest) (*GetConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversation not implemented")
}
func (*UnimplementedPosadaServer) UpdateConversation(ctx context.Context, req *UpdateConversationRequest) (*UpdateConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConversation not implemented")
}
func (*UnimplementedPosadaServer) CreateConversation(ctx context.Context, req *CreateConversationRequest) (*CreateConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConversation not implemented")
}
func (*UnimplementedPosadaServer) DeleteConversation(ctx context.Context, req *DeleteConversationRequest) (*DeleteConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConversation not implemented")
}

func RegisterPosadaServer(s *grpc.Server, srv PosadaServer) {
	s.RegisterService(&_Posada_serviceDesc, srv)
}

func _Posada_GetListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).GetListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/GetListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).GetListing(ctx, req.(*GetListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_UpdateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).UpdateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/UpdateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).UpdateListing(ctx, req.(*UpdateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_CreateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).CreateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/CreateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).CreateListing(ctx, req.(*CreateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_DeleteListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).DeleteListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/DeleteListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).DeleteListing(ctx, req.(*DeleteListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/GetReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).GetReservation(ctx, req.(*GetReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).UpdateReservation(ctx, req.(*UpdateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).DeleteReservation(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_GetGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).GetGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/GetGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).GetGuest(ctx, req.(*GetGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_UpdateGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).UpdateGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/UpdateGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).UpdateGuest(ctx, req.(*UpdateGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_CreateGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).CreateGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/CreateGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).CreateGuest(ctx, req.(*CreateGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_DeleteGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).DeleteGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/DeleteGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).DeleteGuest(ctx, req.(*DeleteGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_GetCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).GetCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/GetCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).GetCalendar(ctx, req.(*GetCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_UpdateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).UpdateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/UpdateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).UpdateCalendar(ctx, req.(*UpdateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_CreateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).CreateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/CreateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).CreateCalendar(ctx, req.(*CreateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_DeleteCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).DeleteCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/DeleteCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).DeleteCalendar(ctx, req.(*DeleteCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_GetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).GetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/GetConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).GetConversation(ctx, req.(*GetConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_UpdateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).UpdateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/UpdateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).UpdateConversation(ctx, req.(*UpdateConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).CreateConversation(ctx, req.(*CreateConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posada_DeleteConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosadaServer).DeleteConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Posada/DeleteConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosadaServer).DeleteConversation(ctx, req.(*DeleteConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Posada_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Posada",
	HandlerType: (*PosadaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListing",
			Handler:    _Posada_GetListing_Handler,
		},
		{
			MethodName: "UpdateListing",
			Handler:    _Posada_UpdateListing_Handler,
		},
		{
			MethodName: "CreateListing",
			Handler:    _Posada_CreateListing_Handler,
		},
		{
			MethodName: "DeleteListing",
			Handler:    _Posada_DeleteListing_Handler,
		},
		{
			MethodName: "GetReservation",
			Handler:    _Posada_GetReservation_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _Posada_UpdateReservation_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _Posada_CreateReservation_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _Posada_DeleteReservation_Handler,
		},
		{
			MethodName: "GetGuest",
			Handler:    _Posada_GetGuest_Handler,
		},
		{
			MethodName: "UpdateGuest",
			Handler:    _Posada_UpdateGuest_Handler,
		},
		{
			MethodName: "CreateGuest",
			Handler:    _Posada_CreateGuest_Handler,
		},
		{
			MethodName: "DeleteGuest",
			Handler:    _Posada_DeleteGuest_Handler,
		},
		{
			MethodName: "GetCalendar",
			Handler:    _Posada_GetCalendar_Handler,
		},
		{
			MethodName: "UpdateCalendar",
			Handler:    _Posada_UpdateCalendar_Handler,
		},
		{
			MethodName: "CreateCalendar",
			Handler:    _Posada_CreateCalendar_Handler,
		},
		{
			MethodName: "DeleteCalendar",
			Handler:    _Posada_DeleteCalendar_Handler,
		},
		{
			MethodName: "GetConversation",
			Handler:    _Posada_GetConversation_Handler,
		},
		{
			MethodName: "UpdateConversation",
			Handler:    _Posada_UpdateConversation_Handler,
		},
		{
			MethodName: "CreateConversation",
			Handler:    _Posada_CreateConversation_Handler,
		},
		{
			MethodName: "DeleteConversation",
			Handler:    _Posada_DeleteConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
