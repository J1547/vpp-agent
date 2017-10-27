// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interfaces.proto

/*
Package interfaces is a generated protocol buffer package.

It is generated from these files:
	interfaces.proto

It has these top-level messages:
	Interfaces
	InterfacesState
	InterfaceErrors
*/
package interfaces

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InterfaceType int32

const (
	InterfaceType_SOFTWARE_LOOPBACK   InterfaceType = 0
	InterfaceType_ETHERNET_CSMACD     InterfaceType = 1
	InterfaceType_MEMORY_INTERFACE    InterfaceType = 2
	InterfaceType_TAP_INTERFACE       InterfaceType = 3
	InterfaceType_AF_PACKET_INTERFACE InterfaceType = 4
	InterfaceType_VXLAN_TUNNEL        InterfaceType = 5
)

var InterfaceType_name = map[int32]string{
	0: "SOFTWARE_LOOPBACK",
	1: "ETHERNET_CSMACD",
	2: "MEMORY_INTERFACE",
	3: "TAP_INTERFACE",
	4: "AF_PACKET_INTERFACE",
	5: "VXLAN_TUNNEL",
}
var InterfaceType_value = map[string]int32{
	"SOFTWARE_LOOPBACK":   0,
	"ETHERNET_CSMACD":     1,
	"MEMORY_INTERFACE":    2,
	"TAP_INTERFACE":       3,
	"AF_PACKET_INTERFACE": 4,
	"VXLAN_TUNNEL":        5,
}

func (x InterfaceType) String() string {
	return proto.EnumName(InterfaceType_name, int32(x))
}
func (InterfaceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// from vpp/build-root/install-vpp-native/vpp/include/vnet/interface.h
type RxModeType int32

const (
	RxModeType_UNKNOWN   RxModeType = 0
	RxModeType_POLLING   RxModeType = 1
	RxModeType_INTERRUPT RxModeType = 2
	RxModeType_ADAPTIVE  RxModeType = 3
	RxModeType_DEFAULT   RxModeType = 4
)

var RxModeType_name = map[int32]string{
	0: "UNKNOWN",
	1: "POLLING",
	2: "INTERRUPT",
	3: "ADAPTIVE",
	4: "DEFAULT",
}
var RxModeType_value = map[string]int32{
	"UNKNOWN":   0,
	"POLLING":   1,
	"INTERRUPT": 2,
	"ADAPTIVE":  3,
	"DEFAULT":   4,
}

func (x RxModeType) String() string {
	return proto.EnumName(RxModeType_name, int32(x))
}
func (RxModeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Interfaces_Interface_Memif_MemifMode int32

const (
	Interfaces_Interface_Memif_ETHERNET    Interfaces_Interface_Memif_MemifMode = 0
	Interfaces_Interface_Memif_IP          Interfaces_Interface_Memif_MemifMode = 1
	Interfaces_Interface_Memif_PUNT_INJECT Interfaces_Interface_Memif_MemifMode = 2
)

var Interfaces_Interface_Memif_MemifMode_name = map[int32]string{
	0: "ETHERNET",
	1: "IP",
	2: "PUNT_INJECT",
}
var Interfaces_Interface_Memif_MemifMode_value = map[string]int32{
	"ETHERNET":    0,
	"IP":          1,
	"PUNT_INJECT": 2,
}

func (x Interfaces_Interface_Memif_MemifMode) String() string {
	return proto.EnumName(Interfaces_Interface_Memif_MemifMode_name, int32(x))
}
func (Interfaces_Interface_Memif_MemifMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1, 0}
}

type InterfacesState_Interface_Status int32

const (
	InterfacesState_Interface_UNKNOWN_STATUS InterfacesState_Interface_Status = 0
	InterfacesState_Interface_UP             InterfacesState_Interface_Status = 1
	InterfacesState_Interface_DOWN           InterfacesState_Interface_Status = 2
	InterfacesState_Interface_DELETED        InterfacesState_Interface_Status = 3
)

var InterfacesState_Interface_Status_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "UP",
	2: "DOWN",
	3: "DELETED",
}
var InterfacesState_Interface_Status_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"UP":             1,
	"DOWN":           2,
	"DELETED":        3,
}

func (x InterfacesState_Interface_Status) String() string {
	return proto.EnumName(InterfacesState_Interface_Status_name, int32(x))
}
func (InterfacesState_Interface_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

type InterfacesState_Interface_Duplex int32

const (
	InterfacesState_Interface_UNKNOWN_DUPLEX InterfacesState_Interface_Duplex = 0
	InterfacesState_Interface_HALF           InterfacesState_Interface_Duplex = 1
	InterfacesState_Interface_FULL           InterfacesState_Interface_Duplex = 2
)

var InterfacesState_Interface_Duplex_name = map[int32]string{
	0: "UNKNOWN_DUPLEX",
	1: "HALF",
	2: "FULL",
}
var InterfacesState_Interface_Duplex_value = map[string]int32{
	"UNKNOWN_DUPLEX": 0,
	"HALF":           1,
	"FULL":           2,
}

func (x InterfacesState_Interface_Duplex) String() string {
	return proto.EnumName(InterfacesState_Interface_Duplex_name, int32(x))
}
func (InterfacesState_Interface_Duplex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 1}
}

type Interfaces struct {
	Interface []*Interfaces_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *Interfaces) Reset()                    { *m = Interfaces{} }
func (m *Interfaces) String() string            { return proto.CompactTextString(m) }
func (*Interfaces) ProtoMessage()               {}
func (*Interfaces) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Interfaces) GetInterface() []*Interfaces_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type Interfaces_Interface struct {
	Name        string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type        InterfaceType `protobuf:"varint,3,opt,name=type,enum=interfaces.InterfaceType" json:"type,omitempty"`
	Enabled     bool          `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
	PhysAddress string        `protobuf:"bytes,5,opt,name=phys_address,json=physAddress" json:"phys_address,omitempty"`
	Mtu         uint32        `protobuf:"varint,6,opt,name=mtu" json:"mtu,omitempty"`
	Vrf         uint32        `protobuf:"varint,7,opt,name=vrf" json:"vrf,omitempty"`
	RxMode      RxModeType    `protobuf:"varint,8,opt,name=rx_mode,json=rxMode,enum=interfaces.RxModeType" json:"rx_mode,omitempty"`
	// Required format is "ipAddress/ipPrefix"
	IpAddresses    []string                             `protobuf:"bytes,10,rep,name=ip_addresses,json=ipAddresses" json:"ip_addresses,omitempty"`
	RxModeSettings *Interfaces_Interface_RxModeSettings `protobuf:"bytes,11,opt,name=rxModeSettings" json:"rxModeSettings,omitempty"`
	Memif          *Interfaces_Interface_Memif          `protobuf:"bytes,101,opt,name=memif" json:"memif,omitempty"`
	Vxlan          *Interfaces_Interface_Vxlan          `protobuf:"bytes,102,opt,name=vxlan" json:"vxlan,omitempty"`
	Afpacket       *Interfaces_Interface_Afpacket       `protobuf:"bytes,103,opt,name=afpacket" json:"afpacket,omitempty"`
	Tap            *Interfaces_Interface_Tap            `protobuf:"bytes,104,opt,name=tap" json:"tap,omitempty"`
}

func (m *Interfaces_Interface) Reset()                    { *m = Interfaces_Interface{} }
func (m *Interfaces_Interface) String() string            { return proto.CompactTextString(m) }
func (*Interfaces_Interface) ProtoMessage()               {}
func (*Interfaces_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Interfaces_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interfaces_Interface) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Interfaces_Interface) GetType() InterfaceType {
	if m != nil {
		return m.Type
	}
	return InterfaceType_SOFTWARE_LOOPBACK
}

func (m *Interfaces_Interface) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Interfaces_Interface) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *Interfaces_Interface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Interfaces_Interface) GetVrf() uint32 {
	if m != nil {
		return m.Vrf
	}
	return 0
}

func (m *Interfaces_Interface) GetRxMode() RxModeType {
	if m != nil {
		return m.RxMode
	}
	return RxModeType_UNKNOWN
}

func (m *Interfaces_Interface) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *Interfaces_Interface) GetRxModeSettings() *Interfaces_Interface_RxModeSettings {
	if m != nil {
		return m.RxModeSettings
	}
	return nil
}

func (m *Interfaces_Interface) GetMemif() *Interfaces_Interface_Memif {
	if m != nil {
		return m.Memif
	}
	return nil
}

func (m *Interfaces_Interface) GetVxlan() *Interfaces_Interface_Vxlan {
	if m != nil {
		return m.Vxlan
	}
	return nil
}

func (m *Interfaces_Interface) GetAfpacket() *Interfaces_Interface_Afpacket {
	if m != nil {
		return m.Afpacket
	}
	return nil
}

func (m *Interfaces_Interface) GetTap() *Interfaces_Interface_Tap {
	if m != nil {
		return m.Tap
	}
	return nil
}

type Interfaces_Interface_RxModeSettings struct {
	RxMode       RxModeType `protobuf:"varint,1,opt,name=rx_mode,json=rxMode,enum=interfaces.RxModeType" json:"rx_mode,omitempty"`
	QueueID      uint32     `protobuf:"varint,2,opt,name=QueueID,json=queueID" json:"QueueID,omitempty"`
	QueueIDValid uint32     `protobuf:"varint,3,opt,name=QueueIDValid,json=queueIDValid" json:"QueueIDValid,omitempty"`
}

func (m *Interfaces_Interface_RxModeSettings) Reset()         { *m = Interfaces_Interface_RxModeSettings{} }
func (m *Interfaces_Interface_RxModeSettings) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_RxModeSettings) ProtoMessage()    {}
func (*Interfaces_Interface_RxModeSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *Interfaces_Interface_RxModeSettings) GetRxMode() RxModeType {
	if m != nil {
		return m.RxMode
	}
	return RxModeType_UNKNOWN
}

func (m *Interfaces_Interface_RxModeSettings) GetQueueID() uint32 {
	if m != nil {
		return m.QueueID
	}
	return 0
}

func (m *Interfaces_Interface_RxModeSettings) GetQueueIDValid() uint32 {
	if m != nil {
		return m.QueueIDValid
	}
	return 0
}

type Interfaces_Interface_Memif struct {
	Master         bool                                 `protobuf:"varint,1,opt,name=master" json:"master,omitempty"`
	Mode           Interfaces_Interface_Memif_MemifMode `protobuf:"varint,2,opt,name=mode,enum=interfaces.Interfaces_Interface_Memif_MemifMode" json:"mode,omitempty"`
	Id             uint32                               `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	SocketFilename string                               `protobuf:"bytes,4,opt,name=socket_filename,json=socketFilename" json:"socket_filename,omitempty"`
	Secret         string                               `protobuf:"bytes,5,opt,name=secret" json:"secret,omitempty"`
	RingSize       uint32                               `protobuf:"varint,6,opt,name=ring_size,json=ringSize" json:"ring_size,omitempty"`
	BufferSize     uint32                               `protobuf:"varint,7,opt,name=buffer_size,json=bufferSize" json:"buffer_size,omitempty"`
	RxQueues       uint32                               `protobuf:"varint,8,opt,name=rx_queues,json=rxQueues" json:"rx_queues,omitempty"`
	TxQueues       uint32                               `protobuf:"varint,9,opt,name=tx_queues,json=txQueues" json:"tx_queues,omitempty"`
}

func (m *Interfaces_Interface_Memif) Reset()         { *m = Interfaces_Interface_Memif{} }
func (m *Interfaces_Interface_Memif) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Memif) ProtoMessage()    {}
func (*Interfaces_Interface_Memif) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

func (m *Interfaces_Interface_Memif) GetMaster() bool {
	if m != nil {
		return m.Master
	}
	return false
}

func (m *Interfaces_Interface_Memif) GetMode() Interfaces_Interface_Memif_MemifMode {
	if m != nil {
		return m.Mode
	}
	return Interfaces_Interface_Memif_ETHERNET
}

func (m *Interfaces_Interface_Memif) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Interfaces_Interface_Memif) GetSocketFilename() string {
	if m != nil {
		return m.SocketFilename
	}
	return ""
}

func (m *Interfaces_Interface_Memif) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Interfaces_Interface_Memif) GetRingSize() uint32 {
	if m != nil {
		return m.RingSize
	}
	return 0
}

func (m *Interfaces_Interface_Memif) GetBufferSize() uint32 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

func (m *Interfaces_Interface_Memif) GetRxQueues() uint32 {
	if m != nil {
		return m.RxQueues
	}
	return 0
}

func (m *Interfaces_Interface_Memif) GetTxQueues() uint32 {
	if m != nil {
		return m.TxQueues
	}
	return 0
}

type Interfaces_Interface_Vxlan struct {
	SrcAddress string `protobuf:"bytes,1,opt,name=src_address,json=srcAddress" json:"src_address,omitempty"`
	DstAddress string `protobuf:"bytes,2,opt,name=dst_address,json=dstAddress" json:"dst_address,omitempty"`
	Vni        uint32 `protobuf:"varint,3,opt,name=vni" json:"vni,omitempty"`
}

func (m *Interfaces_Interface_Vxlan) Reset()         { *m = Interfaces_Interface_Vxlan{} }
func (m *Interfaces_Interface_Vxlan) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Vxlan) ProtoMessage()    {}
func (*Interfaces_Interface_Vxlan) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 2}
}

func (m *Interfaces_Interface_Vxlan) GetSrcAddress() string {
	if m != nil {
		return m.SrcAddress
	}
	return ""
}

func (m *Interfaces_Interface_Vxlan) GetDstAddress() string {
	if m != nil {
		return m.DstAddress
	}
	return ""
}

func (m *Interfaces_Interface_Vxlan) GetVni() uint32 {
	if m != nil {
		return m.Vni
	}
	return 0
}

type Interfaces_Interface_Afpacket struct {
	HostIfName string `protobuf:"bytes,1,opt,name=host_if_name,json=hostIfName" json:"host_if_name,omitempty"`
}

func (m *Interfaces_Interface_Afpacket) Reset()         { *m = Interfaces_Interface_Afpacket{} }
func (m *Interfaces_Interface_Afpacket) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Afpacket) ProtoMessage()    {}
func (*Interfaces_Interface_Afpacket) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 3}
}

func (m *Interfaces_Interface_Afpacket) GetHostIfName() string {
	if m != nil {
		return m.HostIfName
	}
	return ""
}

type Interfaces_Interface_Tap struct {
	HostIfName string `protobuf:"bytes,1,opt,name=host_if_name,json=hostIfName" json:"host_if_name,omitempty"`
}

func (m *Interfaces_Interface_Tap) Reset()                    { *m = Interfaces_Interface_Tap{} }
func (m *Interfaces_Interface_Tap) String() string            { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Tap) ProtoMessage()               {}
func (*Interfaces_Interface_Tap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0, 4} }

func (m *Interfaces_Interface_Tap) GetHostIfName() string {
	if m != nil {
		return m.HostIfName
	}
	return ""
}

type InterfacesState struct {
	Interface []*InterfacesState_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *InterfacesState) Reset()                    { *m = InterfacesState{} }
func (m *InterfacesState) String() string            { return proto.CompactTextString(m) }
func (*InterfacesState) ProtoMessage()               {}
func (*InterfacesState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InterfacesState) GetInterface() []*InterfacesState_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfacesState_Interface struct {
	Name         string                                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	InternalName string                                `protobuf:"bytes,2,opt,name=internal_name,json=internalName" json:"internal_name,omitempty"`
	Type         InterfaceType                         `protobuf:"varint,3,opt,name=type,enum=interfaces.InterfaceType" json:"type,omitempty"`
	IfIndex      uint32                                `protobuf:"varint,4,opt,name=if_index,json=ifIndex" json:"if_index,omitempty"`
	AdminStatus  InterfacesState_Interface_Status      `protobuf:"varint,5,opt,name=admin_status,json=adminStatus,enum=interfaces.InterfacesState_Interface_Status" json:"admin_status,omitempty"`
	OperStatus   InterfacesState_Interface_Status      `protobuf:"varint,6,opt,name=oper_status,json=operStatus,enum=interfaces.InterfacesState_Interface_Status" json:"oper_status,omitempty"`
	LastChange   int64                                 `protobuf:"varint,7,opt,name=last_change,json=lastChange" json:"last_change,omitempty"`
	PhysAddress  string                                `protobuf:"bytes,8,opt,name=phys_address,json=physAddress" json:"phys_address,omitempty"`
	Speed        uint64                                `protobuf:"varint,9,opt,name=speed" json:"speed,omitempty"`
	Mtu          uint32                                `protobuf:"varint,10,opt,name=mtu" json:"mtu,omitempty"`
	Duplex       InterfacesState_Interface_Duplex      `protobuf:"varint,11,opt,name=duplex,enum=interfaces.InterfacesState_Interface_Duplex" json:"duplex,omitempty"`
	Statistics   *InterfacesState_Interface_Statistics `protobuf:"bytes,100,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *InterfacesState_Interface) Reset()                    { *m = InterfacesState_Interface{} }
func (m *InterfacesState_Interface) String() string            { return proto.CompactTextString(m) }
func (*InterfacesState_Interface) ProtoMessage()               {}
func (*InterfacesState_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *InterfacesState_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfacesState_Interface) GetInternalName() string {
	if m != nil {
		return m.InternalName
	}
	return ""
}

func (m *InterfacesState_Interface) GetType() InterfaceType {
	if m != nil {
		return m.Type
	}
	return InterfaceType_SOFTWARE_LOOPBACK
}

func (m *InterfacesState_Interface) GetIfIndex() uint32 {
	if m != nil {
		return m.IfIndex
	}
	return 0
}

func (m *InterfacesState_Interface) GetAdminStatus() InterfacesState_Interface_Status {
	if m != nil {
		return m.AdminStatus
	}
	return InterfacesState_Interface_UNKNOWN_STATUS
}

func (m *InterfacesState_Interface) GetOperStatus() InterfacesState_Interface_Status {
	if m != nil {
		return m.OperStatus
	}
	return InterfacesState_Interface_UNKNOWN_STATUS
}

func (m *InterfacesState_Interface) GetLastChange() int64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *InterfacesState_Interface) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *InterfacesState_Interface) GetSpeed() uint64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *InterfacesState_Interface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfacesState_Interface) GetDuplex() InterfacesState_Interface_Duplex {
	if m != nil {
		return m.Duplex
	}
	return InterfacesState_Interface_UNKNOWN_DUPLEX
}

func (m *InterfacesState_Interface) GetStatistics() *InterfacesState_Interface_Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type InterfacesState_Interface_Statistics struct {
	InPackets       uint64 `protobuf:"varint,1,opt,name=in_packets,json=inPackets" json:"in_packets,omitempty"`
	InBytes         uint64 `protobuf:"varint,2,opt,name=in_bytes,json=inBytes" json:"in_bytes,omitempty"`
	OutPackets      uint64 `protobuf:"varint,3,opt,name=out_packets,json=outPackets" json:"out_packets,omitempty"`
	OutBytes        uint64 `protobuf:"varint,4,opt,name=out_bytes,json=outBytes" json:"out_bytes,omitempty"`
	DropPackets     uint64 `protobuf:"varint,5,opt,name=drop_packets,json=dropPackets" json:"drop_packets,omitempty"`
	PuntPackets     uint64 `protobuf:"varint,6,opt,name=punt_packets,json=puntPackets" json:"punt_packets,omitempty"`
	Ipv4Packets     uint64 `protobuf:"varint,7,opt,name=ipv4_packets,json=ipv4Packets" json:"ipv4_packets,omitempty"`
	Ipv6Packets     uint64 `protobuf:"varint,8,opt,name=ipv6_packets,json=ipv6Packets" json:"ipv6_packets,omitempty"`
	InNobufPackets  uint64 `protobuf:"varint,9,opt,name=in_nobuf_packets,json=inNobufPackets" json:"in_nobuf_packets,omitempty"`
	InMissPackets   uint64 `protobuf:"varint,10,opt,name=in_miss_packets,json=inMissPackets" json:"in_miss_packets,omitempty"`
	InErrorPackets  uint64 `protobuf:"varint,11,opt,name=in_error_packets,json=inErrorPackets" json:"in_error_packets,omitempty"`
	OutErrorPackets uint64 `protobuf:"varint,12,opt,name=out_error_packets,json=outErrorPackets" json:"out_error_packets,omitempty"`
}

func (m *InterfacesState_Interface_Statistics) Reset()         { *m = InterfacesState_Interface_Statistics{} }
func (m *InterfacesState_Interface_Statistics) String() string { return proto.CompactTextString(m) }
func (*InterfacesState_Interface_Statistics) ProtoMessage()    {}
func (*InterfacesState_Interface_Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

func (m *InterfacesState_Interface_Statistics) GetInPackets() uint64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetInBytes() uint64 {
	if m != nil {
		return m.InBytes
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetOutPackets() uint64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetOutBytes() uint64 {
	if m != nil {
		return m.OutBytes
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetDropPackets() uint64 {
	if m != nil {
		return m.DropPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetPuntPackets() uint64 {
	if m != nil {
		return m.PuntPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetIpv4Packets() uint64 {
	if m != nil {
		return m.Ipv4Packets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetIpv6Packets() uint64 {
	if m != nil {
		return m.Ipv6Packets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetInNobufPackets() uint64 {
	if m != nil {
		return m.InNobufPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetInMissPackets() uint64 {
	if m != nil {
		return m.InMissPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetInErrorPackets() uint64 {
	if m != nil {
		return m.InErrorPackets
	}
	return 0
}

func (m *InterfacesState_Interface_Statistics) GetOutErrorPackets() uint64 {
	if m != nil {
		return m.OutErrorPackets
	}
	return 0
}

type InterfaceErrors struct {
	Interface []*InterfaceErrors_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *InterfaceErrors) Reset()                    { *m = InterfaceErrors{} }
func (m *InterfaceErrors) String() string            { return proto.CompactTextString(m) }
func (*InterfaceErrors) ProtoMessage()               {}
func (*InterfaceErrors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InterfaceErrors) GetInterface() []*InterfaceErrors_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfaceErrors_Interface struct {
	InterfaceName string                                 `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	ErrorData     []*InterfaceErrors_Interface_ErrorData `protobuf:"bytes,2,rep,name=error_data,json=errorData" json:"error_data,omitempty"`
}

func (m *InterfaceErrors_Interface) Reset()                    { *m = InterfaceErrors_Interface{} }
func (m *InterfaceErrors_Interface) String() string            { return proto.CompactTextString(m) }
func (*InterfaceErrors_Interface) ProtoMessage()               {}
func (*InterfaceErrors_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *InterfaceErrors_Interface) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *InterfaceErrors_Interface) GetErrorData() []*InterfaceErrors_Interface_ErrorData {
	if m != nil {
		return m.ErrorData
	}
	return nil
}

type InterfaceErrors_Interface_ErrorData struct {
	ChangeType   string `protobuf:"bytes,1,opt,name=change_type,json=changeType" json:"change_type,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	LastChange   int64  `protobuf:"varint,3,opt,name=last_change,json=lastChange" json:"last_change,omitempty"`
}

func (m *InterfaceErrors_Interface_ErrorData) Reset()         { *m = InterfaceErrors_Interface_ErrorData{} }
func (m *InterfaceErrors_Interface_ErrorData) String() string { return proto.CompactTextString(m) }
func (*InterfaceErrors_Interface_ErrorData) ProtoMessage()    {}
func (*InterfaceErrors_Interface_ErrorData) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0, 0}
}

func (m *InterfaceErrors_Interface_ErrorData) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *InterfaceErrors_Interface_ErrorData) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *InterfaceErrors_Interface_ErrorData) GetLastChange() int64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func init() {
	proto.RegisterType((*Interfaces)(nil), "interfaces.Interfaces")
	proto.RegisterType((*Interfaces_Interface)(nil), "interfaces.Interfaces.Interface")
	proto.RegisterType((*Interfaces_Interface_RxModeSettings)(nil), "interfaces.Interfaces.Interface.RxModeSettings")
	proto.RegisterType((*Interfaces_Interface_Memif)(nil), "interfaces.Interfaces.Interface.Memif")
	proto.RegisterType((*Interfaces_Interface_Vxlan)(nil), "interfaces.Interfaces.Interface.Vxlan")
	proto.RegisterType((*Interfaces_Interface_Afpacket)(nil), "interfaces.Interfaces.Interface.Afpacket")
	proto.RegisterType((*Interfaces_Interface_Tap)(nil), "interfaces.Interfaces.Interface.Tap")
	proto.RegisterType((*InterfacesState)(nil), "interfaces.InterfacesState")
	proto.RegisterType((*InterfacesState_Interface)(nil), "interfaces.InterfacesState.Interface")
	proto.RegisterType((*InterfacesState_Interface_Statistics)(nil), "interfaces.InterfacesState.Interface.Statistics")
	proto.RegisterType((*InterfaceErrors)(nil), "interfaces.InterfaceErrors")
	proto.RegisterType((*InterfaceErrors_Interface)(nil), "interfaces.InterfaceErrors.Interface")
	proto.RegisterType((*InterfaceErrors_Interface_ErrorData)(nil), "interfaces.InterfaceErrors.Interface.ErrorData")
	proto.RegisterEnum("interfaces.InterfaceType", InterfaceType_name, InterfaceType_value)
	proto.RegisterEnum("interfaces.RxModeType", RxModeType_name, RxModeType_value)
	proto.RegisterEnum("interfaces.Interfaces_Interface_Memif_MemifMode", Interfaces_Interface_Memif_MemifMode_name, Interfaces_Interface_Memif_MemifMode_value)
	proto.RegisterEnum("interfaces.InterfacesState_Interface_Status", InterfacesState_Interface_Status_name, InterfacesState_Interface_Status_value)
	proto.RegisterEnum("interfaces.InterfacesState_Interface_Duplex", InterfacesState_Interface_Duplex_name, InterfacesState_Interface_Duplex_value)
}

func init() { proto.RegisterFile("interfaces.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xcd, 0x6e, 0xdb, 0xc6,
	0x16, 0xb6, 0x7e, 0xac, 0x9f, 0xa3, 0x1f, 0x33, 0x93, 0xdc, 0x5c, 0x46, 0xc1, 0x45, 0x14, 0xdf,
	0x9b, 0xc4, 0x37, 0x48, 0x9d, 0xc0, 0x2d, 0xb2, 0x69, 0x51, 0x80, 0x91, 0xe8, 0x46, 0x8d, 0x44,
	0xa9, 0x23, 0xca, 0x4e, 0xd1, 0x05, 0x41, 0x8b, 0x43, 0x7b, 0x50, 0x8b, 0x64, 0x38, 0x23, 0x43,
	0xce, 0xa6, 0x8f, 0xd0, 0x45, 0x9f, 0xa5, 0x40, 0xdf, 0xa2, 0x2f, 0xd0, 0x55, 0x1e, 0xa2, 0xeb,
	0x62, 0x66, 0x48, 0x8a, 0x4e, 0x0c, 0x58, 0xed, 0xc6, 0xe0, 0x7c, 0xe7, 0xfb, 0xce, 0x39, 0x9c,
	0x39, 0xf3, 0x51, 0x06, 0x8d, 0x06, 0x9c, 0xc4, 0xbe, 0x3b, 0x27, 0x6c, 0x3f, 0x8a, 0x43, 0x1e,
	0x22, 0x58, 0x23, 0xbb, 0x1f, 0x00, 0x60, 0x90, 0x2d, 0xd1, 0xd7, 0x50, 0xcf, 0x82, 0x7a, 0xa1,
	0x5b, 0xda, 0x6b, 0x1c, 0x74, 0xf7, 0x73, 0x09, 0x06, 0xd7, 0x3c, 0xe2, 0xb5, 0xa4, 0xf3, 0x0b,
	0x40, 0x3d, 0x0b, 0x20, 0x04, 0xe5, 0xc0, 0x5d, 0x88, 0x44, 0x85, 0xbd, 0x3a, 0x96, 0xcf, 0xa8,
	0x0b, 0x0d, 0x8f, 0xb0, 0x79, 0x4c, 0x23, 0x4e, 0xc3, 0x40, 0x2f, 0xca, 0x50, 0x1e, 0x42, 0x9f,
	0x41, 0x99, 0x5f, 0x46, 0x44, 0x2f, 0x75, 0x0b, 0x7b, 0xed, 0x83, 0x7b, 0xd7, 0x96, 0xb7, 0x2f,
	0x23, 0x82, 0x25, 0x0d, 0xe9, 0x50, 0x25, 0x81, 0x7b, 0x72, 0x4e, 0x3c, 0xbd, 0xdc, 0x2d, 0xec,
	0xd5, 0x70, 0xba, 0x44, 0x0f, 0xa1, 0x19, 0x9d, 0x5d, 0x32, 0xc7, 0xf5, 0xbc, 0x98, 0x30, 0xa6,
	0x6f, 0xab, 0x5a, 0x02, 0x33, 0x14, 0x84, 0x34, 0x28, 0x2d, 0xf8, 0x52, 0xaf, 0x74, 0x0b, 0x7b,
	0x2d, 0x2c, 0x1e, 0x05, 0x72, 0x11, 0xfb, 0x7a, 0x55, 0x21, 0x17, 0xb1, 0x8f, 0x9e, 0x43, 0x35,
	0x5e, 0x39, 0x8b, 0xd0, 0x23, 0x7a, 0x4d, 0xb6, 0x74, 0x37, 0xdf, 0x12, 0x5e, 0x8d, 0x42, 0x4f,
	0xf5, 0x53, 0x89, 0xe5, 0xb3, 0xa8, 0x4b, 0xa3, 0xb4, 0x2a, 0x61, 0x3a, 0x74, 0x4b, 0xa2, 0x2e,
	0x8d, 0x8c, 0x14, 0x42, 0xc7, 0xd0, 0x56, 0xe4, 0x29, 0xe1, 0x9c, 0x06, 0xa7, 0x4c, 0x6f, 0x74,
	0x0b, 0x7b, 0x8d, 0x83, 0xe7, 0x37, 0x6d, 0x76, 0x52, 0x2f, 0x95, 0xe1, 0x8f, 0xd2, 0xa0, 0xaf,
	0x60, 0x7b, 0x41, 0x16, 0xd4, 0xd7, 0x89, 0xcc, 0xf7, 0xf8, 0xc6, 0x7c, 0x23, 0xc1, 0xc6, 0x4a,
	0x24, 0xd4, 0x17, 0xab, 0x73, 0x37, 0xd0, 0xfd, 0x0d, 0xd5, 0x47, 0x82, 0x8d, 0x95, 0x08, 0x99,
	0x50, 0x73, 0xfd, 0xc8, 0x9d, 0xff, 0x48, 0xb8, 0x7e, 0x2a, 0x13, 0xfc, 0xff, 0xc6, 0x04, 0x46,
	0x22, 0xc0, 0x99, 0x14, 0xbd, 0x84, 0x12, 0x77, 0x23, 0xfd, 0x4c, 0x66, 0xf8, 0xdf, 0x8d, 0x19,
	0x6c, 0x37, 0xc2, 0x42, 0xd0, 0xf9, 0x09, 0xda, 0x57, 0x37, 0x27, 0x7f, 0x72, 0x85, 0x8d, 0x4e,
	0x4e, 0x87, 0xea, 0x77, 0x4b, 0xb2, 0x24, 0x83, 0xbe, 0x1c, 0xcc, 0x16, 0xae, 0xbe, 0x53, 0x4b,
	0xb4, 0x0b, 0xcd, 0x24, 0x72, 0xe4, 0x9e, 0x53, 0x4f, 0x0e, 0x67, 0x0b, 0x37, 0xdf, 0xe5, 0xb0,
	0xce, 0x87, 0x22, 0x6c, 0xcb, 0xed, 0x44, 0x77, 0xa1, 0xb2, 0x70, 0x19, 0x27, 0xb1, 0xac, 0x5b,
	0xc3, 0xc9, 0x0a, 0xf5, 0xa1, 0x2c, 0xbb, 0x29, 0xca, 0x6e, 0x5e, 0x6c, 0x76, 0x38, 0xea, 0xaf,
	0xe8, 0x0f, 0x4b, 0x35, 0x6a, 0x43, 0x31, 0xeb, 0xa0, 0x48, 0x3d, 0xf4, 0x04, 0x76, 0x58, 0x28,
	0xb6, 0xce, 0xf1, 0xe9, 0x39, 0x91, 0x37, 0xae, 0x2c, 0x47, 0xbd, 0xad, 0xe0, 0xc3, 0x04, 0x15,
	0x6d, 0x31, 0x32, 0x8f, 0x09, 0x4f, 0xae, 0x42, 0xb2, 0x42, 0xf7, 0xa1, 0x1e, 0xd3, 0xe0, 0xd4,
	0x61, 0xf4, 0x3d, 0x49, 0xee, 0x42, 0x4d, 0x00, 0x53, 0xfa, 0x9e, 0xa0, 0x07, 0xd0, 0x38, 0x59,
	0xfa, 0x3e, 0x89, 0x55, 0x58, 0x5d, 0x0c, 0x50, 0x90, 0x24, 0x08, 0xf5, 0xca, 0x91, 0x3b, 0xc1,
	0xe4, 0x0d, 0x11, 0xea, 0x95, 0xdc, 0x2d, 0x26, 0x82, 0x3c, 0x0b, 0xd6, 0x55, 0x90, 0x27, 0xc1,
	0xdd, 0x03, 0xa8, 0x67, 0xef, 0x86, 0x9a, 0x50, 0x33, 0xed, 0xd7, 0x26, 0xb6, 0x4c, 0x5b, 0xdb,
	0x42, 0x15, 0x28, 0x0e, 0x26, 0x5a, 0x01, 0xed, 0x40, 0x63, 0x32, 0xb3, 0x6c, 0x67, 0x60, 0x7d,
	0x6b, 0xf6, 0x6c, 0xad, 0xd8, 0xf9, 0x01, 0xb6, 0xe5, 0xd0, 0x89, 0xbe, 0x58, 0x3c, 0xcf, 0x2e,
	0xb7, 0xf2, 0x18, 0x60, 0xf1, 0x3c, 0xbd, 0xdb, 0x0f, 0xa0, 0xe1, 0x31, 0x9e, 0x11, 0x94, 0xd3,
	0x80, 0xc7, 0x78, 0xee, 0xf2, 0x5f, 0x04, 0x34, 0xd9, 0x48, 0xf1, 0xd8, 0x79, 0x06, 0xb5, 0x74,
	0x20, 0x51, 0x17, 0x9a, 0x67, 0x21, 0xe3, 0x0e, 0xf5, 0x9d, 0x9c, 0x89, 0x81, 0xc0, 0x06, 0xbe,
	0xe5, 0x2e, 0x48, 0xe7, 0x09, 0x94, 0x6c, 0x37, 0xba, 0x99, 0xb8, 0xfb, 0x6b, 0x1d, 0x76, 0xd6,
	0xe7, 0x3b, 0xe5, 0x2e, 0x27, 0xa8, 0xf7, 0xa9, 0xd3, 0x3e, 0xba, 0x7e, 0x1e, 0x24, 0xff, 0x7a,
	0xbb, 0xfd, 0xa3, 0x76, 0x93, 0xdd, 0xfe, 0x17, 0x5a, 0x92, 0x1e, 0xb8, 0xe7, 0xaa, 0x3b, 0xb5,
	0x0d, 0xcd, 0x14, 0x14, 0xfd, 0xfd, 0x5d, 0xc7, 0xbd, 0x07, 0x35, 0xea, 0x3b, 0x34, 0xf0, 0xc8,
	0x4a, 0x0e, 0x5a, 0x0b, 0x57, 0xa9, 0x3f, 0x10, 0x4b, 0x34, 0x86, 0xa6, 0xeb, 0x2d, 0x68, 0xe0,
	0x30, 0xee, 0xf2, 0xa5, 0xb2, 0xdc, 0xf6, 0xc1, 0xb3, 0x8d, 0x5e, 0x6c, 0x7f, 0x2a, 0x35, 0xb8,
	0x21, 0x33, 0xa8, 0x05, 0x1a, 0x41, 0x23, 0x8c, 0xc4, 0xec, 0xa9, 0x7c, 0x95, 0x7f, 0x90, 0x0f,
	0x44, 0x82, 0x24, 0xdd, 0x03, 0x68, 0x9c, 0xbb, 0x8c, 0x3b, 0xf3, 0x33, 0x37, 0x38, 0x55, 0xc3,
	0x5c, 0xc2, 0x20, 0xa0, 0x9e, 0x44, 0x3e, 0xf9, 0x66, 0xd4, 0x3e, 0xfd, 0x66, 0xdc, 0x81, 0x6d,
	0x16, 0x11, 0xe2, 0xc9, 0x71, 0x2e, 0x63, 0xb5, 0x48, 0xbf, 0x24, 0xb0, 0xfe, 0x92, 0xf4, 0xa1,
	0xe2, 0x2d, 0xa3, 0x73, 0xb2, 0x92, 0xde, 0xbe, 0x71, 0xd7, 0x7d, 0xa9, 0xc1, 0x89, 0x16, 0x4d,
	0x00, 0xc4, 0xbb, 0x53, 0xc6, 0xe9, 0x9c, 0xe9, 0x9e, 0x34, 0xc5, 0x17, 0x9b, 0xbf, 0xbf, 0xd2,
	0xe1, 0x5c, 0x8e, 0xce, 0x6f, 0x25, 0x80, 0x75, 0x08, 0xfd, 0x07, 0x80, 0x06, 0x8e, 0x1a, 0x7a,
	0x75, 0x8d, 0xca, 0x62, 0xc4, 0x26, 0x0a, 0x90, 0x87, 0x1d, 0x38, 0x27, 0x97, 0x9c, 0xa8, 0x2b,
	0x54, 0xc6, 0x55, 0x1a, 0xbc, 0x12, 0x4b, 0xb1, 0x99, 0xe1, 0x92, 0x67, 0xd2, 0x92, 0x8c, 0x42,
	0xb8, 0xe4, 0xa9, 0xf6, 0x3e, 0xd4, 0x05, 0x41, 0x89, 0xcb, 0x32, 0x5c, 0x0b, 0x97, 0x5c, 0xa9,
	0x1f, 0x42, 0xd3, 0x8b, 0xc3, 0x28, 0x93, 0x6f, 0xcb, 0x78, 0x43, 0x60, 0xa9, 0x5e, 0x1c, 0xc6,
	0x32, 0x58, 0x57, 0xa8, 0x28, 0x8a, 0xc0, 0x72, 0x14, 0x1a, 0x5d, 0x7c, 0x91, 0x51, 0xaa, 0x8a,
	0x22, 0xb0, 0xab, 0x94, 0x97, 0x19, 0xa5, 0x96, 0x51, 0x5e, 0xa6, 0x94, 0x3d, 0xf1, 0x2b, 0xc9,
	0x09, 0xc2, 0x93, 0xa5, 0x9f, 0xd1, 0xd4, 0xe9, 0xb6, 0x69, 0x60, 0x09, 0x38, 0x65, 0x3e, 0x86,
	0x1d, 0x1a, 0x38, 0x0b, 0xca, 0x58, 0x46, 0x04, 0x49, 0x6c, 0xd1, 0x60, 0x44, 0x19, 0xbb, 0x9a,
	0x91, 0xc4, 0x71, 0x18, 0x67, 0xc4, 0x46, 0x9a, 0xd1, 0x14, 0x70, 0xca, 0x7c, 0x0a, 0xb7, 0xc4,
	0x26, 0x5d, 0xa5, 0x36, 0x25, 0x75, 0x27, 0x5c, 0xf2, 0x3c, 0x77, 0xf7, 0x4b, 0xa8, 0x24, 0x83,
	0x8c, 0xa0, 0x3d, 0xb3, 0xde, 0x58, 0xe3, 0x63, 0xcb, 0x99, 0xda, 0x86, 0x3d, 0x9b, 0x2a, 0xcf,
	0x9c, 0x09, 0xcf, 0xac, 0x41, 0xb9, 0x3f, 0x3e, 0xb6, 0xb4, 0x22, 0x6a, 0x40, 0xb5, 0x6f, 0x0e,
	0x4d, 0xdb, 0xec, 0x6b, 0xa5, 0xdd, 0x17, 0x50, 0x51, 0xb3, 0x95, 0x17, 0xf7, 0x67, 0x93, 0xa1,
	0xf9, 0x56, 0xdb, 0x12, 0xa2, 0xd7, 0xc6, 0xf0, 0x50, 0xc9, 0x0f, 0x67, 0xc3, 0xa1, 0x56, 0xdc,
	0xfd, 0xbd, 0x98, 0xf3, 0x2d, 0xd9, 0x08, 0xdb, 0xdc, 0xb7, 0x14, 0xff, 0x7a, 0xdf, 0xfa, 0xb3,
	0x90, 0xf7, 0xad, 0x47, 0xd0, 0xce, 0x42, 0x79, 0x0b, 0x6d, 0x65, 0xa8, 0x74, 0x29, 0x0b, 0x40,
	0x6d, 0x92, 0xe7, 0x72, 0x57, 0x2f, 0xca, 0xd2, 0xcf, 0x37, 0x2a, 0xbd, 0x2f, 0x81, 0xbe, 0xcb,
	0x5d, 0x5c, 0x27, 0xe9, 0x63, 0x27, 0x86, 0x7a, 0x86, 0x8b, 0x59, 0x56, 0x9e, 0xe0, 0x48, 0x27,
	0x4c, 0x3c, 0x5c, 0x41, 0xc2, 0xfa, 0x84, 0x91, 0xaa, 0xea, 0x0b, 0xc2, 0x98, 0x7b, 0x9a, 0x19,
	0xa9, 0x04, 0x47, 0x0a, 0xfb, 0xd8, 0x5e, 0x4a, 0x1f, 0xdb, 0xcb, 0xd3, 0x9f, 0x0b, 0xd0, 0xba,
	0x62, 0xa9, 0xe8, 0x5f, 0x70, 0x6b, 0x3a, 0x3e, 0xb4, 0x8f, 0x0d, 0x6c, 0x3a, 0xc3, 0xf1, 0x78,
	0xf2, 0xca, 0xe8, 0xbd, 0xd1, 0xb6, 0xd0, 0x6d, 0xd8, 0x49, 0xbf, 0x86, 0x4e, 0x6f, 0x3a, 0x32,
	0x7a, 0x7d, 0xad, 0x80, 0xee, 0x80, 0x36, 0x32, 0x47, 0x63, 0xfc, 0xbd, 0x33, 0xb0, 0x6c, 0x13,
	0x1f, 0x1a, 0x3d, 0x53, 0x2b, 0xa2, 0x5b, 0xd0, 0xb2, 0x8d, 0x49, 0x0e, 0x2a, 0xa1, 0x7f, 0xc3,
	0x6d, 0xe3, 0xd0, 0x99, 0x18, 0xbd, 0x37, 0xa6, 0x9d, 0x0b, 0x94, 0x91, 0x06, 0xcd, 0xa3, 0xb7,
	0x43, 0xc3, 0x72, 0xec, 0x99, 0x65, 0x99, 0x43, 0x6d, 0xfb, 0xe9, 0x04, 0x60, 0xfd, 0x43, 0x48,
	0x0c, 0x4c, 0x32, 0x19, 0xda, 0x96, 0x58, 0x4c, 0xc6, 0xc3, 0xe1, 0xc0, 0xfa, 0x46, 0x2b, 0xa0,
	0x16, 0xd4, 0x65, 0x22, 0x3c, 0x9b, 0xd8, 0x5a, 0x51, 0x7c, 0xad, 0x8d, 0xbe, 0x31, 0xb1, 0x07,
	0x47, 0xa2, 0x9e, 0x9c, 0xb3, 0x43, 0x63, 0x36, 0xb4, 0xb5, 0xf2, 0x49, 0x45, 0xfe, 0x97, 0xf1,
	0xf9, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0xf9, 0x8e, 0x5b, 0x79, 0x0c, 0x00, 0x00,
}
