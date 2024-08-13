// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: modules/nodesale/protobuf/nodesale.proto

// protoc modules/nodesale/protobuf/nodesale.proto --go_out=. --go_opt=module=github.com/gaze-network/indexer-network

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_ACTION_DEPLOY   Action = 0
	Action_ACTION_PURCHASE Action = 1
	Action_ACTION_DELEGATE Action = 2
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_DEPLOY",
		1: "ACTION_PURCHASE",
		2: "ACTION_DELEGATE",
	}
	Action_value = map[string]int32{
		"ACTION_DEPLOY":   0,
		"ACTION_PURCHASE": 1,
		"ACTION_DELEGATE": 2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_modules_nodesale_protobuf_nodesale_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_modules_nodesale_protobuf_nodesale_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{0}
}

type NodeSaleEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   Action          `protobuf:"varint,1,opt,name=action,proto3,enum=nodesale.Action" json:"action,omitempty"`
	Deploy   *ActionDeploy   `protobuf:"bytes,2,opt,name=deploy,proto3,oneof" json:"deploy,omitempty"`
	Purchase *ActionPurchase `protobuf:"bytes,3,opt,name=purchase,proto3,oneof" json:"purchase,omitempty"`
	Delegate *ActionDelegate `protobuf:"bytes,4,opt,name=delegate,proto3,oneof" json:"delegate,omitempty"`
}

func (x *NodeSaleEvent) Reset() {
	*x = NodeSaleEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSaleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSaleEvent) ProtoMessage() {}

func (x *NodeSaleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSaleEvent.ProtoReflect.Descriptor instead.
func (*NodeSaleEvent) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{0}
}

func (x *NodeSaleEvent) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_DEPLOY
}

func (x *NodeSaleEvent) GetDeploy() *ActionDeploy {
	if x != nil {
		return x.Deploy
	}
	return nil
}

func (x *NodeSaleEvent) GetPurchase() *ActionPurchase {
	if x != nil {
		return x.Purchase
	}
	return nil
}

func (x *NodeSaleEvent) GetDelegate() *ActionDelegate {
	if x != nil {
		return x.Delegate
	}
	return nil
}

type ActionDeploy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartsAt              uint32  `protobuf:"varint,2,opt,name=startsAt,proto3" json:"startsAt,omitempty"`
	EndsAt                uint32  `protobuf:"varint,3,opt,name=endsAt,proto3" json:"endsAt,omitempty"`
	Tiers                 []*Tier `protobuf:"bytes,4,rep,name=tiers,proto3" json:"tiers,omitempty"`
	SellerPublicKey       string  `protobuf:"bytes,5,opt,name=sellerPublicKey,proto3" json:"sellerPublicKey,omitempty"`
	MaxPerAddress         uint32  `protobuf:"varint,6,opt,name=maxPerAddress,proto3" json:"maxPerAddress,omitempty"`
	MaxDiscountPercentage uint32  `protobuf:"varint,7,opt,name=maxDiscountPercentage,proto3" json:"maxDiscountPercentage,omitempty"`
	SellerWallet          string  `protobuf:"bytes,8,opt,name=sellerWallet,proto3" json:"sellerWallet,omitempty"`
}

func (x *ActionDeploy) Reset() {
	*x = ActionDeploy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionDeploy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionDeploy) ProtoMessage() {}

func (x *ActionDeploy) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionDeploy.ProtoReflect.Descriptor instead.
func (*ActionDeploy) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{1}
}

func (x *ActionDeploy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionDeploy) GetStartsAt() uint32 {
	if x != nil {
		return x.StartsAt
	}
	return 0
}

func (x *ActionDeploy) GetEndsAt() uint32 {
	if x != nil {
		return x.EndsAt
	}
	return 0
}

func (x *ActionDeploy) GetTiers() []*Tier {
	if x != nil {
		return x.Tiers
	}
	return nil
}

func (x *ActionDeploy) GetSellerPublicKey() string {
	if x != nil {
		return x.SellerPublicKey
	}
	return ""
}

func (x *ActionDeploy) GetMaxPerAddress() uint32 {
	if x != nil {
		return x.MaxPerAddress
	}
	return 0
}

func (x *ActionDeploy) GetMaxDiscountPercentage() uint32 {
	if x != nil {
		return x.MaxDiscountPercentage
	}
	return 0
}

func (x *ActionDeploy) GetSellerWallet() string {
	if x != nil {
		return x.SellerWallet
	}
	return ""
}

type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriceSat      uint32 `protobuf:"varint,1,opt,name=priceSat,proto3" json:"priceSat,omitempty"`
	Limit         uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	MaxPerAddress uint32 `protobuf:"varint,3,opt,name=maxPerAddress,proto3" json:"maxPerAddress,omitempty"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{2}
}

func (x *Tier) GetPriceSat() uint32 {
	if x != nil {
		return x.PriceSat
	}
	return 0
}

func (x *Tier) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Tier) GetMaxPerAddress() uint32 {
	if x != nil {
		return x.MaxPerAddress
	}
	return 0
}

type ActionPurchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload         *PurchasePayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	SellerSignature string           `protobuf:"bytes,2,opt,name=sellerSignature,proto3" json:"sellerSignature,omitempty"`
}

func (x *ActionPurchase) Reset() {
	*x = ActionPurchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionPurchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionPurchase) ProtoMessage() {}

func (x *ActionPurchase) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionPurchase.ProtoReflect.Descriptor instead.
func (*ActionPurchase) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{3}
}

func (x *ActionPurchase) GetPayload() *PurchasePayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ActionPurchase) GetSellerSignature() string {
	if x != nil {
		return x.SellerSignature
	}
	return ""
}

type PurchasePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeployID       *ActionID `protobuf:"bytes,1,opt,name=deployID,proto3" json:"deployID,omitempty"`
	BuyerPublicKey string    `protobuf:"bytes,2,opt,name=buyerPublicKey,proto3" json:"buyerPublicKey,omitempty"`
	NodeIDs        []uint32  `protobuf:"varint,3,rep,packed,name=nodeIDs,proto3" json:"nodeIDs,omitempty"`
	TotalAmountSat int64     `protobuf:"varint,4,opt,name=totalAmountSat,proto3" json:"totalAmountSat,omitempty"`
	TimeOutBlock   uint64    `protobuf:"varint,5,opt,name=timeOutBlock,proto3" json:"timeOutBlock,omitempty"`
}

func (x *PurchasePayload) Reset() {
	*x = PurchasePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchasePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchasePayload) ProtoMessage() {}

func (x *PurchasePayload) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchasePayload.ProtoReflect.Descriptor instead.
func (*PurchasePayload) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{4}
}

func (x *PurchasePayload) GetDeployID() *ActionID {
	if x != nil {
		return x.DeployID
	}
	return nil
}

func (x *PurchasePayload) GetBuyerPublicKey() string {
	if x != nil {
		return x.BuyerPublicKey
	}
	return ""
}

func (x *PurchasePayload) GetNodeIDs() []uint32 {
	if x != nil {
		return x.NodeIDs
	}
	return nil
}

func (x *PurchasePayload) GetTotalAmountSat() int64 {
	if x != nil {
		return x.TotalAmountSat
	}
	return 0
}

func (x *PurchasePayload) GetTimeOutBlock() uint64 {
	if x != nil {
		return x.TimeOutBlock
	}
	return 0
}

type ActionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block   uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	TxIndex uint32 `protobuf:"varint,2,opt,name=txIndex,proto3" json:"txIndex,omitempty"`
}

func (x *ActionID) Reset() {
	*x = ActionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionID) ProtoMessage() {}

func (x *ActionID) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionID.ProtoReflect.Descriptor instead.
func (*ActionID) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{5}
}

func (x *ActionID) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *ActionID) GetTxIndex() uint32 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

type ActionDelegate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegateePublicKey string    `protobuf:"bytes,1,opt,name=delegateePublicKey,proto3" json:"delegateePublicKey,omitempty"`
	NodeIDs            []uint32  `protobuf:"varint,2,rep,packed,name=nodeIDs,proto3" json:"nodeIDs,omitempty"`
	DeployID           *ActionID `protobuf:"bytes,3,opt,name=deployID,proto3" json:"deployID,omitempty"`
}

func (x *ActionDelegate) Reset() {
	*x = ActionDelegate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionDelegate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionDelegate) ProtoMessage() {}

func (x *ActionDelegate) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nodesale_protobuf_nodesale_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionDelegate.ProtoReflect.Descriptor instead.
func (*ActionDelegate) Descriptor() ([]byte, []int) {
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP(), []int{6}
}

func (x *ActionDelegate) GetDelegateePublicKey() string {
	if x != nil {
		return x.DelegateePublicKey
	}
	return ""
}

func (x *ActionDelegate) GetNodeIDs() []uint32 {
	if x != nil {
		return x.NodeIDs
	}
	return nil
}

func (x *ActionDelegate) GetDeployID() *ActionID {
	if x != nil {
		return x.DeployID
	}
	return nil
}

var File_modules_nodesale_protobuf_nodesale_proto protoreflect.FileDescriptor

var file_modules_nodesale_protobuf_nodesale_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x61, 0x6c, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61,
	0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x48, 0x01, 0x52, 0x08, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x39, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x48, 0x02, 0x52, 0x08,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x22, 0xa6, 0x02, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x69, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52, 0x05, 0x74, 0x69, 0x65, 0x72, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x78,
	0x50, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x34, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15,
	0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x5e, 0x0a, 0x04, 0x54, 0x69, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x53, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x53, 0x61, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x0e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0f, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e,
	0x0a, 0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x44, 0x12, 0x26,
	0x0a, 0x0e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65,
	0x4f, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x3a, 0x0a, 0x08,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61,
	0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x08, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x49, 0x44, 0x2a, 0x45, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x55, 0x52,
	0x43, 0x48, 0x41, 0x53, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x47, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x7a, 0x65, 0x2d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_nodesale_protobuf_nodesale_proto_rawDescOnce sync.Once
	file_modules_nodesale_protobuf_nodesale_proto_rawDescData = file_modules_nodesale_protobuf_nodesale_proto_rawDesc
)

func file_modules_nodesale_protobuf_nodesale_proto_rawDescGZIP() []byte {
	file_modules_nodesale_protobuf_nodesale_proto_rawDescOnce.Do(func() {
		file_modules_nodesale_protobuf_nodesale_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_nodesale_protobuf_nodesale_proto_rawDescData)
	})
	return file_modules_nodesale_protobuf_nodesale_proto_rawDescData
}

var file_modules_nodesale_protobuf_nodesale_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_modules_nodesale_protobuf_nodesale_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_modules_nodesale_protobuf_nodesale_proto_goTypes = []interface{}{
	(Action)(0),             // 0: nodesale.Action
	(*NodeSaleEvent)(nil),   // 1: nodesale.NodeSaleEvent
	(*ActionDeploy)(nil),    // 2: nodesale.ActionDeploy
	(*Tier)(nil),            // 3: nodesale.Tier
	(*ActionPurchase)(nil),  // 4: nodesale.ActionPurchase
	(*PurchasePayload)(nil), // 5: nodesale.PurchasePayload
	(*ActionID)(nil),        // 6: nodesale.ActionID
	(*ActionDelegate)(nil),  // 7: nodesale.ActionDelegate
}
var file_modules_nodesale_protobuf_nodesale_proto_depIdxs = []int32{
	0, // 0: nodesale.NodeSaleEvent.action:type_name -> nodesale.Action
	2, // 1: nodesale.NodeSaleEvent.deploy:type_name -> nodesale.ActionDeploy
	4, // 2: nodesale.NodeSaleEvent.purchase:type_name -> nodesale.ActionPurchase
	7, // 3: nodesale.NodeSaleEvent.delegate:type_name -> nodesale.ActionDelegate
	3, // 4: nodesale.ActionDeploy.tiers:type_name -> nodesale.Tier
	5, // 5: nodesale.ActionPurchase.payload:type_name -> nodesale.PurchasePayload
	6, // 6: nodesale.PurchasePayload.deployID:type_name -> nodesale.ActionID
	6, // 7: nodesale.ActionDelegate.deployID:type_name -> nodesale.ActionID
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_modules_nodesale_protobuf_nodesale_proto_init() }
func file_modules_nodesale_protobuf_nodesale_proto_init() {
	if File_modules_nodesale_protobuf_nodesale_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSaleEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionDeploy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionPurchase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchasePayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_nodesale_protobuf_nodesale_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionDelegate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_modules_nodesale_protobuf_nodesale_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_nodesale_protobuf_nodesale_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_nodesale_protobuf_nodesale_proto_goTypes,
		DependencyIndexes: file_modules_nodesale_protobuf_nodesale_proto_depIdxs,
		EnumInfos:         file_modules_nodesale_protobuf_nodesale_proto_enumTypes,
		MessageInfos:      file_modules_nodesale_protobuf_nodesale_proto_msgTypes,
	}.Build()
	File_modules_nodesale_protobuf_nodesale_proto = out.File
	file_modules_nodesale_protobuf_nodesale_proto_rawDesc = nil
	file_modules_nodesale_protobuf_nodesale_proto_goTypes = nil
	file_modules_nodesale_protobuf_nodesale_proto_depIdxs = nil
}