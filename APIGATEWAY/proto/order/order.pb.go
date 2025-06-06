// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: order.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemIds       []string               `protobuf:"bytes,3,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

func (x *Order) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemIds       []string               `protobuf:"bytes,2,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemIds       []string               `protobuf:"bytes,3,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateOrderRequest) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

func (x *UpdateOrderRequest) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *UpdateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PatchOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchOrderStatusRequest) Reset() {
	*x = PatchOrderStatusRequest{}
	mi := &file_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchOrderStatusRequest) ProtoMessage() {}

func (x *PatchOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*PatchOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *PatchOrderStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatchOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PatchOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchOrderStatusResponse) Reset() {
	*x = PatchOrderStatusResponse{}
	mi := &file_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchOrderStatusResponse) ProtoMessage() {}

func (x *PatchOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*PatchOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *PatchOrderStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	mi := &file_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrderResponse) Reset() {
	*x = DeleteOrderResponse{}
	mi := &file_order_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderResponse) ProtoMessage() {}

func (x *DeleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int64                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Skip          int64                  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_order_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{11}
}

func (x *ListOrdersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOrdersRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_order_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{12}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ListOrdersByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Skip          int64                  `protobuf:"varint,3,opt,name=skip,proto3" json:"skip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersByUserRequest) Reset() {
	*x = ListOrdersByUserRequest{}
	mi := &file_order_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersByUserRequest) ProtoMessage() {}

func (x *ListOrdersByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersByUserRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersByUserRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{13}
}

func (x *ListOrdersByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListOrdersByUserRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOrdersByUserRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

type ListOrdersByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersByUserResponse) Reset() {
	*x = ListOrdersByUserResponse{}
	mi := &file_order_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersByUserResponse) ProtoMessage() {}

func (x *ListOrdersByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersByUserResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersByUserResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{14}
}

func (x *ListOrdersByUserResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x05order\"\xa3\x01\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x19\n" +
	"\bitem_ids\x18\x03 \x03(\tR\aitemIds\x12\x1f\n" +
	"\vtotal_price\x18\x04 \x01(\x01R\n" +
	"totalPrice\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\"H\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x19\n" +
	"\bitem_ids\x18\x02 \x03(\tR\aitemIds\"%\n" +
	"\x13CreateOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"!\n" +
	"\x0fGetOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"6\n" +
	"\x10GetOrderResponse\x12\"\n" +
	"\x05order\x18\x01 \x01(\v2\f.order.OrderR\x05order\"\x91\x01\n" +
	"\x12UpdateOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x19\n" +
	"\bitem_ids\x18\x03 \x03(\tR\aitemIds\x12\x1f\n" +
	"\vtotal_price\x18\x04 \x01(\x01R\n" +
	"totalPrice\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\"/\n" +
	"\x13UpdateOrderResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"A\n" +
	"\x17PatchOrderStatusRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"4\n" +
	"\x18PatchOrderStatusResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"$\n" +
	"\x12DeleteOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"/\n" +
	"\x13DeleteOrderResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"=\n" +
	"\x11ListOrdersRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x03R\x05limit\x12\x12\n" +
	"\x04skip\x18\x02 \x01(\x03R\x04skip\":\n" +
	"\x12ListOrdersResponse\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders\"\\\n" +
	"\x17ListOrdersByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x03R\x05limit\x12\x12\n" +
	"\x04skip\x18\x03 \x01(\x03R\x04skip\"@\n" +
	"\x18ListOrdersByUserResponse\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders2\x8a\x04\n" +
	"\fOrderService\x12D\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\x1a.order.CreateOrderResponse\x12;\n" +
	"\bGetOrder\x12\x16.order.GetOrderRequest\x1a\x17.order.GetOrderResponse\x12D\n" +
	"\vUpdateOrder\x12\x19.order.UpdateOrderRequest\x1a\x1a.order.UpdateOrderResponse\x12D\n" +
	"\vDeleteOrder\x12\x19.order.DeleteOrderRequest\x1a\x1a.order.DeleteOrderResponse\x12A\n" +
	"\n" +
	"ListOrders\x12\x18.order.ListOrdersRequest\x1a\x19.order.ListOrdersResponse\x12S\n" +
	"\x10PatchOrderStatus\x12\x1e.order.PatchOrderStatusRequest\x1a\x1f.order.PatchOrderStatusResponse\x12S\n" +
	"\x10ListOrdersByUser\x12\x1e.order.ListOrdersByUserRequest\x1a\x1f.order.ListOrdersByUserResponseB\x1bZ\x19order_service/proto;protob\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_order_proto_goTypes = []any{
	(*Order)(nil),                    // 0: order.Order
	(*CreateOrderRequest)(nil),       // 1: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),      // 2: order.CreateOrderResponse
	(*GetOrderRequest)(nil),          // 3: order.GetOrderRequest
	(*GetOrderResponse)(nil),         // 4: order.GetOrderResponse
	(*UpdateOrderRequest)(nil),       // 5: order.UpdateOrderRequest
	(*UpdateOrderResponse)(nil),      // 6: order.UpdateOrderResponse
	(*PatchOrderStatusRequest)(nil),  // 7: order.PatchOrderStatusRequest
	(*PatchOrderStatusResponse)(nil), // 8: order.PatchOrderStatusResponse
	(*DeleteOrderRequest)(nil),       // 9: order.DeleteOrderRequest
	(*DeleteOrderResponse)(nil),      // 10: order.DeleteOrderResponse
	(*ListOrdersRequest)(nil),        // 11: order.ListOrdersRequest
	(*ListOrdersResponse)(nil),       // 12: order.ListOrdersResponse
	(*ListOrdersByUserRequest)(nil),  // 13: order.ListOrdersByUserRequest
	(*ListOrdersByUserResponse)(nil), // 14: order.ListOrdersByUserResponse
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: order.GetOrderResponse.order:type_name -> order.Order
	0,  // 1: order.ListOrdersResponse.orders:type_name -> order.Order
	0,  // 2: order.ListOrdersByUserResponse.orders:type_name -> order.Order
	1,  // 3: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	3,  // 4: order.OrderService.GetOrder:input_type -> order.GetOrderRequest
	5,  // 5: order.OrderService.UpdateOrder:input_type -> order.UpdateOrderRequest
	9,  // 6: order.OrderService.DeleteOrder:input_type -> order.DeleteOrderRequest
	11, // 7: order.OrderService.ListOrders:input_type -> order.ListOrdersRequest
	7,  // 8: order.OrderService.PatchOrderStatus:input_type -> order.PatchOrderStatusRequest
	13, // 9: order.OrderService.ListOrdersByUser:input_type -> order.ListOrdersByUserRequest
	2,  // 10: order.OrderService.CreateOrder:output_type -> order.CreateOrderResponse
	4,  // 11: order.OrderService.GetOrder:output_type -> order.GetOrderResponse
	6,  // 12: order.OrderService.UpdateOrder:output_type -> order.UpdateOrderResponse
	10, // 13: order.OrderService.DeleteOrder:output_type -> order.DeleteOrderResponse
	12, // 14: order.OrderService.ListOrders:output_type -> order.ListOrdersResponse
	8,  // 15: order.OrderService.PatchOrderStatus:output_type -> order.PatchOrderStatusResponse
	14, // 16: order.OrderService.ListOrdersByUser:output_type -> order.ListOrdersByUserResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}