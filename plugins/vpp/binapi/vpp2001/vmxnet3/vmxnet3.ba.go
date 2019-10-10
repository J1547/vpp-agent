// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/vmxnet3.api.json

/*
Package vmxnet3 is a generated VPP binary API for 'vmxnet3' module.

It consists of:
	  2 types
	  6 messages
	  3 services
*/
package vmxnet3

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "vmxnet3"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd126c788
)

// Vmxnet3RxList represents VPP binary API type 'vmxnet3_rx_list'.
type Vmxnet3RxList struct {
	RxQsize   uint16
	RxFill    []uint16 `struc:"[2]uint16"`
	RxNext    uint16
	RxProduce []uint16 `struc:"[2]uint16"`
	RxConsume []uint16 `struc:"[2]uint16"`
}

func (*Vmxnet3RxList) GetTypeName() string {
	return "vmxnet3_rx_list"
}

// Vmxnet3TxList represents VPP binary API type 'vmxnet3_tx_list'.
type Vmxnet3TxList struct {
	TxQsize   uint16
	TxNext    uint16
	TxProduce uint16
	TxConsume uint16
}

func (*Vmxnet3TxList) GetTypeName() string {
	return "vmxnet3_tx_list"
}

// Vmxnet3Create represents VPP binary API message 'vmxnet3_create'.
type Vmxnet3Create struct {
	PciAddr    uint32
	EnableElog int32
	RxqSize    uint16
	RxqNum     uint16
	TxqSize    uint16
	TxqNum     uint16
	Bind       uint8
	EnableGso  uint8
}

func (*Vmxnet3Create) GetMessageName() string {
	return "vmxnet3_create"
}
func (*Vmxnet3Create) GetCrcString() string {
	return "80cc3559"
}
func (*Vmxnet3Create) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// Vmxnet3CreateReply represents VPP binary API message 'vmxnet3_create_reply'.
type Vmxnet3CreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*Vmxnet3CreateReply) GetMessageName() string {
	return "vmxnet3_create_reply"
}
func (*Vmxnet3CreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*Vmxnet3CreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// Vmxnet3Delete represents VPP binary API message 'vmxnet3_delete'.
type Vmxnet3Delete struct {
	SwIfIndex uint32
}

func (*Vmxnet3Delete) GetMessageName() string {
	return "vmxnet3_delete"
}
func (*Vmxnet3Delete) GetCrcString() string {
	return "529cb13f"
}
func (*Vmxnet3Delete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// Vmxnet3DeleteReply represents VPP binary API message 'vmxnet3_delete_reply'.
type Vmxnet3DeleteReply struct {
	Retval int32
}

func (*Vmxnet3DeleteReply) GetMessageName() string {
	return "vmxnet3_delete_reply"
}
func (*Vmxnet3DeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*Vmxnet3DeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// Vmxnet3Details represents VPP binary API message 'vmxnet3_details'.
type Vmxnet3Details struct {
	SwIfIndex   uint32
	IfName      []byte `struc:"[64]byte"`
	HwAddr      []byte `struc:"[6]byte"`
	PciAddr     uint32
	Version     uint8
	AdminUpDown uint8
	RxCount     uint8
	RxList      []Vmxnet3RxList `struc:"[16]Vmxnet3RxList"`
	TxCount     uint8
	TxList      []Vmxnet3TxList `struc:"[8]Vmxnet3TxList"`
}

func (*Vmxnet3Details) GetMessageName() string {
	return "vmxnet3_details"
}
func (*Vmxnet3Details) GetCrcString() string {
	return "25f4412f"
}
func (*Vmxnet3Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// Vmxnet3Dump represents VPP binary API message 'vmxnet3_dump'.
type Vmxnet3Dump struct{}

func (*Vmxnet3Dump) GetMessageName() string {
	return "vmxnet3_dump"
}
func (*Vmxnet3Dump) GetCrcString() string {
	return "51077d14"
}
func (*Vmxnet3Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*Vmxnet3Create)(nil), "vmxnet3.Vmxnet3Create")
	api.RegisterMessage((*Vmxnet3CreateReply)(nil), "vmxnet3.Vmxnet3CreateReply")
	api.RegisterMessage((*Vmxnet3Delete)(nil), "vmxnet3.Vmxnet3Delete")
	api.RegisterMessage((*Vmxnet3DeleteReply)(nil), "vmxnet3.Vmxnet3DeleteReply")
	api.RegisterMessage((*Vmxnet3Details)(nil), "vmxnet3.Vmxnet3Details")
	api.RegisterMessage((*Vmxnet3Dump)(nil), "vmxnet3.Vmxnet3Dump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Vmxnet3Create)(nil),
		(*Vmxnet3CreateReply)(nil),
		(*Vmxnet3Delete)(nil),
		(*Vmxnet3DeleteReply)(nil),
		(*Vmxnet3Details)(nil),
		(*Vmxnet3Dump)(nil),
	}
}

// RPCService represents RPC service API for vmxnet3 module.
type RPCService interface {
	DumpVmxnet3(ctx context.Context, in *Vmxnet3Dump) (RPCService_DumpVmxnet3Client, error)
	Vmxnet3Create(ctx context.Context, in *Vmxnet3Create) (*Vmxnet3CreateReply, error)
	Vmxnet3Delete(ctx context.Context, in *Vmxnet3Delete) (*Vmxnet3DeleteReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpVmxnet3(ctx context.Context, in *Vmxnet3Dump) (RPCService_DumpVmxnet3Client, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpVmxnet3Client{stream}
	return x, nil
}

type RPCService_DumpVmxnet3Client interface {
	Recv() (*Vmxnet3Details, error)
}

type serviceClient_DumpVmxnet3Client struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpVmxnet3Client) Recv() (*Vmxnet3Details, error) {
	m := new(Vmxnet3Details)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) Vmxnet3Create(ctx context.Context, in *Vmxnet3Create) (*Vmxnet3CreateReply, error) {
	out := new(Vmxnet3CreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Vmxnet3Delete(ctx context.Context, in *Vmxnet3Delete) (*Vmxnet3DeleteReply, error) {
	out := new(Vmxnet3DeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
