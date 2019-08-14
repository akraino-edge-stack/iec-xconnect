/*
 * Copyright 2017-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by protoc-gen-go.
// source: device.proto
// DO NOT EDIT!

/*
Package voltha is a generated protocol buffer package.

It is generated from these files:
	device.proto

It has these top-level messages:
	DeviceType
	DeviceTypes
	PmConfig
	PmGroupConfig
	PmConfigs
	Port
	Ports
	Device
	Devices
*/
package device

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import voltha3 "github.com/opencord/voltha/netconf/translator/voltha/common"
import openflow_13 "github.com/opencord/voltha/netconf/translator/voltha/openflow_13"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PmConfig_PmType int32

const (
	PmConfig_COUNTER PmConfig_PmType = 0
	PmConfig_GUAGE   PmConfig_PmType = 1
	PmConfig_STATE   PmConfig_PmType = 2
)

var PmConfig_PmType_name = map[int32]string{
	0: "COUNTER",
	1: "GUAGE",
	2: "STATE",
}
var PmConfig_PmType_value = map[string]int32{
	"COUNTER": 0,
	"GUAGE":   1,
	"STATE":   2,
}

func (x PmConfig_PmType) String() string {
	return proto.EnumName(PmConfig_PmType_name, int32(x))
}
func (PmConfig_PmType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Port_PortType int32

const (
	Port_UNKNOWN      Port_PortType = 0
	Port_ETHERNET_NNI Port_PortType = 1
	Port_ETHERNET_UNI Port_PortType = 2
	Port_PON_OLT      Port_PortType = 3
	Port_PON_ONU      Port_PortType = 4
)

var Port_PortType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ETHERNET_NNI",
	2: "ETHERNET_UNI",
	3: "PON_OLT",
	4: "PON_ONU",
}
var Port_PortType_value = map[string]int32{
	"UNKNOWN":      0,
	"ETHERNET_NNI": 1,
	"ETHERNET_UNI": 2,
	"PON_OLT":      3,
	"PON_ONU":      4,
}

func (x Port_PortType) String() string {
	return proto.EnumName(Port_PortType_name, int32(x))
}
func (Port_PortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

// A Device Type
type DeviceType struct {
	// Unique name for the device type
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Name of the adapter that handles device type
	Adapter                     string `protobuf:"bytes,2,opt,name=adapter" json:"adapter,omitempty"`
	AcceptsBulkFlowUpdate       bool   `protobuf:"varint,3,opt,name=accepts_bulk_flow_update,json=acceptsBulkFlowUpdate" json:"accepts_bulk_flow_update,omitempty"`
	AcceptsAddRemoveFlowUpdates bool   `protobuf:"varint,4,opt,name=accepts_add_remove_flow_updates,json=acceptsAddRemoveFlowUpdates" json:"accepts_add_remove_flow_updates,omitempty"`
}

func (m *DeviceType) Reset()                    { *m = DeviceType{} }
func (m *DeviceType) String() string            { return proto.CompactTextString(m) }
func (*DeviceType) ProtoMessage()               {}
func (*DeviceType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceType) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceType) GetAdapter() string {
	if m != nil {
		return m.Adapter
	}
	return ""
}

func (m *DeviceType) GetAcceptsBulkFlowUpdate() bool {
	if m != nil {
		return m.AcceptsBulkFlowUpdate
	}
	return false
}

func (m *DeviceType) GetAcceptsAddRemoveFlowUpdates() bool {
	if m != nil {
		return m.AcceptsAddRemoveFlowUpdates
	}
	return false
}

// A plurality of device types
type DeviceTypes struct {
	Items []*DeviceType `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *DeviceTypes) Reset()                    { *m = DeviceTypes{} }
func (m *DeviceTypes) String() string            { return proto.CompactTextString(m) }
func (*DeviceTypes) ProtoMessage()               {}
func (*DeviceTypes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceTypes) GetItems() []*DeviceType {
	if m != nil {
		return m.Items
	}
	return nil
}

type PmConfig struct {
	Name       string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type       PmConfig_PmType `protobuf:"varint,2,opt,name=type,enum=voltha.PmConfig_PmType" json:"type,omitempty"`
	Enabled    bool            `protobuf:"varint,3,opt,name=enabled" json:"enabled,omitempty"`
	SampleFreq uint32          `protobuf:"varint,4,opt,name=sample_freq,json=sampleFreq" json:"sample_freq,omitempty"`
}

func (m *PmConfig) Reset()                    { *m = PmConfig{} }
func (m *PmConfig) String() string            { return proto.CompactTextString(m) }
func (*PmConfig) ProtoMessage()               {}
func (*PmConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PmConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PmConfig) GetType() PmConfig_PmType {
	if m != nil {
		return m.Type
	}
	return PmConfig_COUNTER
}

func (m *PmConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PmConfig) GetSampleFreq() uint32 {
	if m != nil {
		return m.SampleFreq
	}
	return 0
}

type PmGroupConfig struct {
	GroupName string      `protobuf:"bytes,1,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	GroupFreq uint32      `protobuf:"varint,2,opt,name=group_freq,json=groupFreq" json:"group_freq,omitempty"`
	Enabled   bool        `protobuf:"varint,3,opt,name=enabled" json:"enabled,omitempty"`
	Metrics   []*PmConfig `protobuf:"bytes,4,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *PmGroupConfig) Reset()                    { *m = PmGroupConfig{} }
func (m *PmGroupConfig) String() string            { return proto.CompactTextString(m) }
func (*PmGroupConfig) ProtoMessage()               {}
func (*PmGroupConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PmGroupConfig) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *PmGroupConfig) GetGroupFreq() uint32 {
	if m != nil {
		return m.GroupFreq
	}
	return 0
}

func (m *PmGroupConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PmGroupConfig) GetMetrics() []*PmConfig {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type PmConfigs struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DefaultFreq uint32 `protobuf:"varint,2,opt,name=default_freq,json=defaultFreq" json:"default_freq,omitempty"`
	// Forces group names and group semantics
	Grouped bool `protobuf:"varint,3,opt,name=grouped" json:"grouped,omitempty"`
	// Allows Pm to set an individual sample frequency
	FreqOverride bool             `protobuf:"varint,4,opt,name=freq_override,json=freqOverride" json:"freq_override,omitempty"`
	Groups       []*PmGroupConfig `protobuf:"bytes,5,rep,name=groups" json:"groups,omitempty"`
	Metrics      []*PmConfig      `protobuf:"bytes,6,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *PmConfigs) Reset()                    { *m = PmConfigs{} }
func (m *PmConfigs) String() string            { return proto.CompactTextString(m) }
func (*PmConfigs) ProtoMessage()               {}
func (*PmConfigs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PmConfigs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PmConfigs) GetDefaultFreq() uint32 {
	if m != nil {
		return m.DefaultFreq
	}
	return 0
}

func (m *PmConfigs) GetGrouped() bool {
	if m != nil {
		return m.Grouped
	}
	return false
}

func (m *PmConfigs) GetFreqOverride() bool {
	if m != nil {
		return m.FreqOverride
	}
	return false
}

func (m *PmConfigs) GetGroups() []*PmGroupConfig {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *PmConfigs) GetMetrics() []*PmConfig {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Port struct {
	PortNo     uint32                        `protobuf:"varint,1,opt,name=port_no,json=portNo" json:"port_no,omitempty"`
	Label      string                        `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Type       Port_PortType                 `protobuf:"varint,3,opt,name=type,enum=voltha.Port_PortType" json:"type,omitempty"`
	AdminState voltha3.AdminState_AdminState `protobuf:"varint,5,opt,name=admin_state,json=adminState,enum=voltha.AdminState_AdminState" json:"admin_state,omitempty"`
	OperStatus voltha3.OperStatus_OperStatus `protobuf:"varint,6,opt,name=oper_status,json=operStatus,enum=voltha.OperStatus_OperStatus" json:"oper_status,omitempty"`
	DeviceId   string                        `protobuf:"bytes,7,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Peers      []*Port_PeerPort              `protobuf:"bytes,8,rep,name=peers" json:"peers,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Port) GetPortNo() uint32 {
	if m != nil {
		return m.PortNo
	}
	return 0
}

func (m *Port) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Port) GetType() Port_PortType {
	if m != nil {
		return m.Type
	}
	return Port_UNKNOWN
}

func (m *Port) GetAdminState() voltha3.AdminState_AdminState {
	if m != nil {
		return m.AdminState
	}
	return voltha3.AdminState_UNKNOWN
}

func (m *Port) GetOperStatus() voltha3.OperStatus_OperStatus {
	if m != nil {
		return m.OperStatus
	}
	return voltha3.OperStatus_UNKNOWN
}

func (m *Port) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Port) GetPeers() []*Port_PeerPort {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Port_PeerPort struct {
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	PortNo   uint32 `protobuf:"varint,2,opt,name=port_no,json=portNo" json:"port_no,omitempty"`
}

func (m *Port_PeerPort) Reset()                    { *m = Port_PeerPort{} }
func (m *Port_PeerPort) String() string            { return proto.CompactTextString(m) }
func (*Port_PeerPort) ProtoMessage()               {}
func (*Port_PeerPort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *Port_PeerPort) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Port_PeerPort) GetPortNo() uint32 {
	if m != nil {
		return m.PortNo
	}
	return 0
}

type Ports struct {
	Items []*Port `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Ports) Reset()                    { *m = Ports{} }
func (m *Ports) String() string            { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()               {}
func (*Ports) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Ports) GetItems() []*Port {
	if m != nil {
		return m.Items
	}
	return nil
}

// A Physical Device instance
type Device struct {
	// Voltha's device identifier
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Device type, refers to one of the registered device types
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// Is this device a root device. Each logical switch has one root
	// device that is associated with the logical flow switch.
	Root bool `protobuf:"varint,3,opt,name=root" json:"root,omitempty"`
	// Parent device id, in the device tree (for a root device, the parent_id
	// is the logical_device.id)
	ParentId     string `protobuf:"bytes,4,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	ParentPortNo uint32 `protobuf:"varint,20,opt,name=parent_port_no,json=parentPortNo" json:"parent_port_no,omitempty"`
	// Vendor, version, serial number, etc.
	Vendor          string `protobuf:"bytes,5,opt,name=vendor" json:"vendor,omitempty"`
	Model           string `protobuf:"bytes,6,opt,name=model" json:"model,omitempty"`
	HardwareVersion string `protobuf:"bytes,7,opt,name=hardware_version,json=hardwareVersion" json:"hardware_version,omitempty"`
	FirmwareVersion string `protobuf:"bytes,8,opt,name=firmware_version,json=firmwareVersion" json:"firmware_version,omitempty"`
	SoftwareVersion string `protobuf:"bytes,9,opt,name=software_version,json=softwareVersion" json:"software_version,omitempty"`
	SerialNumber    string `protobuf:"bytes,10,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	// Addapter that takes care of device
	Adapter string `protobuf:"bytes,11,opt,name=adapter" json:"adapter,omitempty"`
	// Device contact on vlan (if 0, no vlan)
	Vlan uint32 `protobuf:"varint,12,opt,name=vlan" json:"vlan,omitempty"`
	// Types that are valid to be assigned to Address:
	//	*Device_MacAddress
	//	*Device_Ipv4Address
	//	*Device_Ipv6Address
	//	*Device_HostAndPort
	Address       isDevice_Address                    `protobuf_oneof:"address"`
	ProxyAddress  *Device_ProxyAddress                `protobuf:"bytes,19,opt,name=proxy_address,json=proxyAddress" json:"proxy_address,omitempty"`
	AdminState    voltha3.AdminState_AdminState       `protobuf:"varint,16,opt,name=admin_state,json=adminState,enum=voltha.AdminState_AdminState" json:"admin_state,omitempty"`
	OperStatus    voltha3.OperStatus_OperStatus       `protobuf:"varint,17,opt,name=oper_status,json=operStatus,enum=voltha.OperStatus_OperStatus" json:"oper_status,omitempty"`
	Reason        string                              `protobuf:"bytes,22,opt,name=reason" json:"reason,omitempty"`
	ConnectStatus voltha3.ConnectStatus_ConnectStatus `protobuf:"varint,18,opt,name=connect_status,json=connectStatus,enum=voltha.ConnectStatus_ConnectStatus" json:"connect_status,omitempty"`
	// Device type specific attributes
	Custom     *google_protobuf1.Any   `protobuf:"bytes,64,opt,name=custom" json:"custom,omitempty"`
	Ports      []*Port                 `protobuf:"bytes,128,rep,name=ports" json:"ports,omitempty"`
	Flows      *openflow_13.Flows      `protobuf:"bytes,129,opt,name=flows" json:"flows,omitempty"`
	FlowGroups *openflow_13.FlowGroups `protobuf:"bytes,130,opt,name=flow_groups,json=flowGroups" json:"flow_groups,omitempty"`
	// PmConfigs will eventually converted to a child node of the
	// device to falicitata callbacks and to simplify manipulation.
	PmConfigs *PmConfigs `protobuf:"bytes,131,opt,name=pm_configs,json=pmConfigs" json:"pm_configs,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isDevice_Address interface {
	isDevice_Address()
}

type Device_MacAddress struct {
	MacAddress string `protobuf:"bytes,13,opt,name=mac_address,json=macAddress,oneof"`
}
type Device_Ipv4Address struct {
	Ipv4Address string `protobuf:"bytes,14,opt,name=ipv4_address,json=ipv4Address,oneof"`
}
type Device_Ipv6Address struct {
	Ipv6Address string `protobuf:"bytes,15,opt,name=ipv6_address,json=ipv6Address,oneof"`
}
type Device_HostAndPort struct {
	HostAndPort string `protobuf:"bytes,21,opt,name=host_and_port,json=hostAndPort,oneof"`
}

func (*Device_MacAddress) isDevice_Address()  {}
func (*Device_Ipv4Address) isDevice_Address() {}
func (*Device_Ipv6Address) isDevice_Address() {}
func (*Device_HostAndPort) isDevice_Address() {}

func (m *Device) GetAddress() isDevice_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Device) GetRoot() bool {
	if m != nil {
		return m.Root
	}
	return false
}

func (m *Device) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *Device) GetParentPortNo() uint32 {
	if m != nil {
		return m.ParentPortNo
	}
	return 0
}

func (m *Device) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Device) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Device) GetHardwareVersion() string {
	if m != nil {
		return m.HardwareVersion
	}
	return ""
}

func (m *Device) GetFirmwareVersion() string {
	if m != nil {
		return m.FirmwareVersion
	}
	return ""
}

func (m *Device) GetSoftwareVersion() string {
	if m != nil {
		return m.SoftwareVersion
	}
	return ""
}

func (m *Device) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Device) GetAdapter() string {
	if m != nil {
		return m.Adapter
	}
	return ""
}

func (m *Device) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *Device) GetMacAddress() string {
	if x, ok := m.GetAddress().(*Device_MacAddress); ok {
		return x.MacAddress
	}
	return ""
}

func (m *Device) GetIpv4Address() string {
	if x, ok := m.GetAddress().(*Device_Ipv4Address); ok {
		return x.Ipv4Address
	}
	return ""
}

func (m *Device) GetIpv6Address() string {
	if x, ok := m.GetAddress().(*Device_Ipv6Address); ok {
		return x.Ipv6Address
	}
	return ""
}

func (m *Device) GetHostAndPort() string {
	if x, ok := m.GetAddress().(*Device_HostAndPort); ok {
		return x.HostAndPort
	}
	return ""
}

func (m *Device) GetProxyAddress() *Device_ProxyAddress {
	if m != nil {
		return m.ProxyAddress
	}
	return nil
}

func (m *Device) GetAdminState() voltha3.AdminState_AdminState {
	if m != nil {
		return m.AdminState
	}
	return voltha3.AdminState_UNKNOWN
}

func (m *Device) GetOperStatus() voltha3.OperStatus_OperStatus {
	if m != nil {
		return m.OperStatus
	}
	return voltha3.OperStatus_UNKNOWN
}

func (m *Device) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Device) GetConnectStatus() voltha3.ConnectStatus_ConnectStatus {
	if m != nil {
		return m.ConnectStatus
	}
	return voltha3.ConnectStatus_UNKNOWN
}

func (m *Device) GetCustom() *google_protobuf1.Any {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *Device) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Device) GetFlows() *openflow_13.Flows {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *Device) GetFlowGroups() *openflow_13.FlowGroups {
	if m != nil {
		return m.FlowGroups
	}
	return nil
}

func (m *Device) GetPmConfigs() *PmConfigs {
	if m != nil {
		return m.PmConfigs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Device) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Device_OneofMarshaler, _Device_OneofUnmarshaler, _Device_OneofSizer, []interface{}{
		(*Device_MacAddress)(nil),
		(*Device_Ipv4Address)(nil),
		(*Device_Ipv6Address)(nil),
		(*Device_HostAndPort)(nil),
	}
}

func _Device_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Device)
	// address
	switch x := m.Address.(type) {
	case *Device_MacAddress:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.MacAddress)
	case *Device_Ipv4Address:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv4Address)
	case *Device_Ipv6Address:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv6Address)
	case *Device_HostAndPort:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.HostAndPort)
	case nil:
	default:
		return fmt.Errorf("Device.Address has unexpected type %T", x)
	}
	return nil
}

func _Device_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Device)
	switch tag {
	case 13: // address.mac_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &Device_MacAddress{x}
		return true, err
	case 14: // address.ipv4_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &Device_Ipv4Address{x}
		return true, err
	case 15: // address.ipv6_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &Device_Ipv6Address{x}
		return true, err
	case 21: // address.host_and_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &Device_HostAndPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _Device_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Device)
	// address
	switch x := m.Address.(type) {
	case *Device_MacAddress:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.MacAddress)))
		n += len(x.MacAddress)
	case *Device_Ipv4Address:
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv4Address)))
		n += len(x.Ipv4Address)
	case *Device_Ipv6Address:
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv6Address)))
		n += len(x.Ipv6Address)
	case *Device_HostAndPort:
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.HostAndPort)))
		n += len(x.HostAndPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Device_ProxyAddress struct {
	DeviceId     string `protobuf:"bytes,1,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	ChannelId    uint32 `protobuf:"varint,2,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	OnuId        uint32 `protobuf:"varint,3,opt,name=onu_id,json=onuId" json:"onu_id,omitempty"`
	OnuSessionId uint32 `protobuf:"varint,4,opt,name=onu_session_id,json=onuSessionId" json:"onu_session_id,omitempty"`
}

func (m *Device_ProxyAddress) Reset()                    { *m = Device_ProxyAddress{} }
func (m *Device_ProxyAddress) String() string            { return proto.CompactTextString(m) }
func (*Device_ProxyAddress) ProtoMessage()               {}
func (*Device_ProxyAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *Device_ProxyAddress) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Device_ProxyAddress) GetChannelId() uint32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *Device_ProxyAddress) GetOnuId() uint32 {
	if m != nil {
		return m.OnuId
	}
	return 0
}

func (m *Device_ProxyAddress) GetOnuSessionId() uint32 {
	if m != nil {
		return m.OnuSessionId
	}
	return 0
}

type Devices struct {
	Items []*Device `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Devices) Reset()                    { *m = Devices{} }
func (m *Devices) String() string            { return proto.CompactTextString(m) }
func (*Devices) ProtoMessage()               {}
func (*Devices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Devices) GetItems() []*Device {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceType)(nil), "voltha.DeviceType")
	proto.RegisterType((*DeviceTypes)(nil), "voltha.DeviceTypes")
	proto.RegisterType((*PmConfig)(nil), "voltha.PmConfig")
	proto.RegisterType((*PmGroupConfig)(nil), "voltha.PmGroupConfig")
	proto.RegisterType((*PmConfigs)(nil), "voltha.PmConfigs")
	proto.RegisterType((*Port)(nil), "voltha.Port")
	proto.RegisterType((*Port_PeerPort)(nil), "voltha.Port.PeerPort")
	proto.RegisterType((*Ports)(nil), "voltha.Ports")
	proto.RegisterType((*Device)(nil), "voltha.Device")
	proto.RegisterType((*Device_ProxyAddress)(nil), "voltha.Device.ProxyAddress")
	proto.RegisterType((*Devices)(nil), "voltha.Devices")
	proto.RegisterEnum("voltha.PmConfig_PmType", PmConfig_PmType_name, PmConfig_PmType_value)
	proto.RegisterEnum("voltha.Port_PortType", Port_PortType_name, Port_PortType_value)
}

func init() { proto.RegisterFile("device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4d, 0x6f, 0xdb, 0xb6,
	0x1b, 0x8f, 0x1c, 0xcb, 0x2f, 0x8f, 0x5f, 0xea, 0xf2, 0xdf, 0xfc, 0xab, 0x26, 0x30, 0x9a, 0xba,
	0x3d, 0x64, 0xcd, 0xe6, 0x76, 0xed, 0xd0, 0x0e, 0x3b, 0x0c, 0x71, 0x5a, 0xb7, 0x0d, 0x36, 0x38,
	0x9e, 0x62, 0x6f, 0x47, 0x81, 0x91, 0xe8, 0x44, 0xa8, 0x44, 0xaa, 0xa4, 0xec, 0x2e, 0xb7, 0xbd,
	0x1c, 0xf6, 0x01, 0xf6, 0x51, 0x86, 0x7d, 0x89, 0x01, 0xbd, 0xed, 0x13, 0xec, 0x30, 0xec, 0xb8,
	0x53, 0xce, 0x03, 0x49, 0x51, 0x91, 0xd2, 0xa1, 0xdb, 0x2e, 0x06, 0xf9, 0x7b, 0x21, 0xf9, 0x3c,
	0x7e, 0xf8, 0x50, 0xd0, 0x0e, 0xc8, 0x2a, 0xf4, 0xc9, 0x30, 0xe1, 0x2c, 0x65, 0xa8, 0xb6, 0x62,
	0x51, 0x7a, 0x8a, 0x37, 0x21, 0x26, 0x29, 0xd6, 0xd8, 0xe6, 0x8d, 0x13, 0xc6, 0x4e, 0x22, 0x72,
	0x4f, 0xcd, 0x8e, 0x97, 0x8b, 0x7b, 0x98, 0x9e, 0x65, 0x54, 0xdb, 0x67, 0x71, 0xcc, 0x68, 0x36,
	0xbb, 0xca, 0x12, 0x42, 0x17, 0x11, 0x7b, 0xed, 0x7d, 0xf8, 0x30, 0x83, 0xd0, 0x19, 0xa6, 0x27,
	0x1e, 0x4b, 0xd2, 0x90, 0x51, 0xa1, 0xb1, 0xc1, 0xcf, 0x16, 0xc0, 0x53, 0xb5, 0xe9, 0xec, 0x2c,
	0x21, 0xa8, 0x0b, 0x95, 0x30, 0x70, 0xac, 0x6d, 0x6b, 0xa7, 0xe9, 0x56, 0xc2, 0x00, 0x39, 0x50,
	0xc7, 0x01, 0x4e, 0x52, 0xc2, 0x9d, 0x8a, 0x02, 0xcd, 0x14, 0x3d, 0x06, 0x07, 0xfb, 0x3e, 0x49,
	0x52, 0xe1, 0x1d, 0x2f, 0xa3, 0x97, 0x9e, 0xda, 0x6a, 0x99, 0x04, 0x38, 0x25, 0xce, 0xfa, 0xb6,
	0xb5, 0xd3, 0x70, 0x37, 0x32, 0x7e, 0x7f, 0x19, 0xbd, 0x7c, 0x16, 0xb1, 0xd7, 0x73, 0x45, 0xa2,
	0xa7, 0x70, 0xd3, 0x18, 0x71, 0x10, 0x78, 0x9c, 0xc4, 0x6c, 0x45, 0x8a, 0x76, 0xe1, 0x54, 0x95,
	0x7f, 0x2b, 0x93, 0x8d, 0x82, 0xc0, 0x55, 0xa2, 0x8b, 0x45, 0xc4, 0xe0, 0x31, 0xb4, 0x2e, 0x8e,
	0x2d, 0xd0, 0x0e, 0xd8, 0x61, 0x4a, 0x62, 0xe1, 0x58, 0xdb, 0xeb, 0x3b, 0xad, 0x07, 0x68, 0xa8,
	0x53, 0x37, 0xbc, 0xd0, 0xb8, 0x5a, 0x30, 0xf8, 0xc9, 0x82, 0xc6, 0x34, 0x7e, 0xc2, 0xe8, 0x22,
	0x3c, 0x41, 0x08, 0xaa, 0x14, 0xc7, 0x24, 0x0b, 0x58, 0x8d, 0xd1, 0x2e, 0x54, 0xd3, 0xb3, 0x84,
	0xa8, 0x78, 0xbb, 0x0f, 0xae, 0x9b, 0x95, 0x8c, 0x67, 0x38, 0x8d, 0xd5, 0x72, 0x4a, 0x24, 0xf3,
	0x43, 0x28, 0x3e, 0x8e, 0x48, 0x90, 0x05, 0x6d, 0xa6, 0xe8, 0x26, 0xb4, 0x04, 0x8e, 0x93, 0x88,
	0x78, 0x0b, 0x4e, 0x5e, 0xa9, 0x90, 0x3a, 0x2e, 0x68, 0xe8, 0x19, 0x27, 0xaf, 0x06, 0xbb, 0x50,
	0xd3, 0x4b, 0xa1, 0x16, 0xd4, 0x9f, 0x1c, 0xce, 0x27, 0xb3, 0xb1, 0xdb, 0x5b, 0x43, 0x4d, 0xb0,
	0x9f, 0xcf, 0x47, 0xcf, 0xc7, 0x3d, 0x4b, 0x0e, 0x8f, 0x66, 0xa3, 0xd9, 0xb8, 0x57, 0x19, 0xfc,
	0x68, 0x41, 0x67, 0x1a, 0x3f, 0xe7, 0x6c, 0x99, 0x64, 0x47, 0xef, 0x03, 0x9c, 0xc8, 0xa9, 0x57,
	0x08, 0xa0, 0xa9, 0x90, 0x89, 0x8c, 0x22, 0xa7, 0xd5, 0xee, 0x15, 0xb5, 0xbb, 0xa6, 0xe5, 0xe6,
	0xef, 0x38, 0xf7, 0x5d, 0xa8, 0xc7, 0x24, 0xe5, 0xa1, 0x2f, 0xff, 0x06, 0x99, 0xcb, 0xde, 0xe5,
	0x0c, 0xb8, 0x46, 0x30, 0xf8, 0xdd, 0x82, 0xa6, 0x41, 0xc5, 0x5b, 0xb5, 0x73, 0x4b, 0x96, 0xf3,
	0x02, 0x2f, 0xa3, 0xb4, 0x78, 0x88, 0x56, 0x86, 0xa9, 0x63, 0xdc, 0x84, 0xba, 0x3a, 0x93, 0x39,
	0xc6, 0xbe, 0xfd, 0xc7, 0xf9, 0x9b, 0xbe, 0xe5, 0x1a, 0x14, 0xdd, 0x85, 0x8e, 0xf4, 0x7a, 0x6c,
	0x45, 0x38, 0x0f, 0x03, 0xa2, 0x4b, 0xc3, 0xc8, 0xda, 0x92, 0x3b, 0xcc, 0x28, 0xf4, 0x01, 0xd4,
	0x94, 0x4d, 0x38, 0xb6, 0x3a, 0xf8, 0xc6, 0xc5, 0xc1, 0x0b, 0x89, 0x73, 0x33, 0x51, 0x31, 0xd0,
	0xda, 0x3f, 0x05, 0xfa, 0xcb, 0x3a, 0x54, 0xa7, 0x8c, 0xa7, 0xe8, 0x3a, 0xd4, 0x13, 0xc6, 0x53,
	0x8f, 0x32, 0x15, 0x68, 0xc7, 0xad, 0xc9, 0xe9, 0x84, 0xa1, 0x6b, 0x60, 0x47, 0xf8, 0x98, 0x44,
	0xd9, 0x35, 0xd1, 0x13, 0xf4, 0x5e, 0x56, 0x4b, 0xeb, 0xaa, 0x96, 0x2e, 0x0e, 0xc4, 0x78, 0xaa,
	0x7e, 0x0a, 0x95, 0xf4, 0x29, 0xb4, 0x70, 0x10, 0x87, 0xd4, 0x13, 0xa9, 0xbc, 0x42, 0xb6, 0x72,
	0xf4, 0x8d, 0x63, 0x24, 0xa9, 0x23, 0xc9, 0x14, 0x86, 0x2e, 0xe0, 0x7c, 0x2c, 0xfd, 0x2c, 0x21,
	0x5c, 0xd9, 0x97, 0x32, 0xa4, 0x92, 0xff, 0x30, 0x21, 0xfc, 0x48, 0x31, 0x85, 0xa1, 0x0b, 0x2c,
	0x1f, 0xa3, 0x2d, 0x68, 0xea, 0xe6, 0xe3, 0x85, 0x81, 0x53, 0x57, 0x41, 0x34, 0x34, 0x70, 0x10,
	0xa0, 0x5d, 0xb0, 0x13, 0x42, 0xb8, 0x70, 0x1a, 0x97, 0x32, 0xab, 0x02, 0x21, 0x84, 0xcb, 0x81,
	0xab, 0x35, 0x9b, 0x7b, 0xd0, 0x30, 0x50, 0x79, 0x55, 0xeb, 0xd2, 0xaa, 0x85, 0x64, 0x56, 0x8a,
	0xc9, 0x1c, 0xcc, 0xa1, 0x61, 0xb2, 0x23, 0x2f, 0xc7, 0x7c, 0xf2, 0xd9, 0xe4, 0xf0, 0xab, 0x49,
	0x6f, 0x0d, 0xf5, 0xa0, 0x3d, 0x9e, 0xbd, 0x18, 0xbb, 0x93, 0xf1, 0xcc, 0x9b, 0x4c, 0x0e, 0x7a,
	0x56, 0x09, 0x99, 0x4f, 0x0e, 0x7a, 0x15, 0x69, 0x98, 0x1e, 0x4e, 0xbc, 0xc3, 0xcf, 0x67, 0xbd,
	0xf5, 0x7c, 0x32, 0x99, 0xf7, 0xaa, 0x9f, 0xd8, 0x7f, 0x9e, 0xbf, 0xe9, 0xaf, 0x0d, 0x76, 0xc1,
	0x96, 0xab, 0x0b, 0x34, 0x28, 0x37, 0x8d, 0x76, 0x31, 0x2a, 0xd3, 0x2e, 0x7e, 0x05, 0xa8, 0xe9,
	0x26, 0x82, 0x36, 0x2e, 0xea, 0xdb, 0x14, 0xa0, 0x2c, 0xf3, 0x1b, 0x85, 0x7e, 0x91, 0x13, 0xfa,
	0x3f, 0xbd, 0x01, 0x55, 0xce, 0x58, 0x5a, 0xae, 0x6d, 0x05, 0xa1, 0x01, 0x34, 0x13, 0xcc, 0x09,
	0x4d, 0x65, 0x62, 0xaa, 0x45, 0x6b, 0x43, 0xe3, 0x2a, 0xeb, 0xdd, 0x4c, 0x63, 0xd2, 0x74, 0x4d,
	0xa6, 0x29, 0xaf, 0x7e, 0x4d, 0x4e, 0x75, 0x01, 0xf6, 0xa1, 0xb6, 0x22, 0x34, 0x60, 0x5c, 0x95,
	0x4e, 0xbe, 0x5a, 0x06, 0xa2, 0x2d, 0xb0, 0x63, 0x16, 0x90, 0x48, 0x15, 0x46, 0xce, 0x6a, 0x0c,
	0xdd, 0x87, 0xde, 0x29, 0xe6, 0xc1, 0x6b, 0xcc, 0x89, 0xb7, 0x22, 0x5c, 0x84, 0x8c, 0xea, 0x12,
	0x30, 0xba, 0x2b, 0x86, 0xfe, 0x52, 0xb3, 0xd2, 0xb1, 0x08, 0x79, 0x5c, 0x72, 0x34, 0x4a, 0x0e,
	0x43, 0x17, 0x1c, 0x82, 0x2d, 0xd2, 0x92, 0xa3, 0x59, 0x72, 0x18, 0xda, 0x38, 0xee, 0x42, 0x47,
	0x10, 0x1e, 0xe2, 0xc8, 0xa3, 0xcb, 0xf8, 0x98, 0x70, 0x07, 0x8a, 0xf2, 0xb6, 0xe6, 0x26, 0x8a,
	0x92, 0x8d, 0xc4, 0xbc, 0x53, 0xad, 0xa2, 0x2a, 0x7f, 0xae, 0x10, 0x54, 0x57, 0x11, 0xa6, 0x4e,
	0x5b, 0x15, 0x9a, 0x1a, 0xa3, 0x5b, 0xd0, 0x8a, 0xb1, 0x2f, 0x5f, 0x21, 0x4e, 0x84, 0x70, 0x3a,
	0xd2, 0xf8, 0x62, 0xcd, 0x85, 0x18, 0xfb, 0x23, 0x8d, 0xa1, 0xdb, 0xd0, 0x0e, 0x93, 0xd5, 0x47,
	0xb9, 0xa6, 0x9b, 0x69, 0x5a, 0x12, 0x2d, 0x8b, 0x1e, 0xe5, 0xa2, 0x2b, 0x05, 0xd1, 0x23, 0x23,
	0xba, 0x03, 0x9d, 0x53, 0x26, 0x52, 0x0f, 0xd3, 0x40, 0xfd, 0x9d, 0xce, 0x86, 0x51, 0x49, 0x78,
	0x44, 0x03, 0x75, 0x5f, 0xf6, 0xa0, 0x93, 0x70, 0xf6, 0xf5, 0x59, 0xbe, 0xd6, 0xff, 0xb6, 0xad,
	0x9d, 0xd6, 0x83, 0xad, 0xf2, 0x7b, 0x36, 0x9c, 0x4a, 0x4d, 0xb6, 0xb2, 0xdb, 0x4e, 0x0a, 0xb3,
	0xcb, 0x7d, 0xa4, 0xf7, 0x5f, 0xfb, 0xc8, 0xb8, 0xdc, 0x47, 0xae, 0xfe, 0x8b, 0x3e, 0x62, 0x92,
	0x5d, 0x6c, 0x27, 0x7d, 0xa8, 0x71, 0x82, 0x05, 0xa3, 0xce, 0xff, 0x4b, 0xe5, 0xa8, 0x41, 0xf4,
	0x05, 0x74, 0x7d, 0x46, 0x29, 0xf1, 0x53, 0xb3, 0x11, 0x52, 0x1b, 0xdd, 0x36, 0x1b, 0x3d, 0xd1,
	0x6c, 0xb6, 0x57, 0x69, 0x66, 0xd6, 0xea, 0xf8, 0x45, 0x14, 0xbd, 0x0f, 0x35, 0x7f, 0x29, 0x52,
	0x16, 0x3b, 0x7b, 0x2a, 0x67, 0xd7, 0x86, 0xfa, 0x53, 0x69, 0x68, 0x3e, 0x95, 0x86, 0x23, 0x7a,
	0xe6, 0x66, 0x1a, 0xf4, 0x10, 0x6c, 0xf9, 0x2f, 0x08, 0xe7, 0x9b, 0xbf, 0xb9, 0xfc, 0xfb, 0xdd,
	0xdf, 0xce, 0xdf, 0xf4, 0x9b, 0x79, 0x7b, 0x72, 0xb5, 0x16, 0xdd, 0x07, 0x5b, 0x7e, 0xa7, 0x08,
	0xe7, 0x5b, 0x4b, 0x6d, 0x81, 0x86, 0xc5, 0x8f, 0x2c, 0xf9, 0x79, 0x22, 0xf6, 0x6d, 0x69, 0x5d,
	0x73, 0xb5, 0x10, 0xed, 0x41, 0x4b, 0xd1, 0xd9, 0xc3, 0xf4, 0x9d, 0xf6, 0x5d, 0x7f, 0xcb, 0xa7,
	0x1e, 0xa8, 0xdc, 0x0c, 0x8b, 0x1c, 0x42, 0x1f, 0x03, 0x24, 0xb1, 0xe7, 0xeb, 0x37, 0xd6, 0xf9,
	0x5e, 0x2f, 0x70, 0xf5, 0xf2, 0x53, 0x95, 0x5b, 0x9b, 0x89, 0x41, 0x36, 0x7f, 0xb0, 0xa0, 0x5d,
	0x2c, 0x94, 0x77, 0x37, 0xe3, 0x3e, 0x80, 0x7f, 0x8a, 0x29, 0x25, 0x91, 0x64, 0xb3, 0x0f, 0x86,
	0x0c, 0x39, 0x08, 0xd0, 0x06, 0xd4, 0x18, 0x5d, 0x4a, 0x6a, 0x5d, 0x51, 0x36, 0xa3, 0xcb, 0x83,
	0x00, 0xdd, 0x81, 0xae, 0x84, 0x05, 0x11, 0xf2, 0xca, 0x9a, 0x5e, 0xd6, 0x71, 0xdb, 0x8c, 0x2e,
	0x8f, 0x34, 0x78, 0x10, 0x64, 0x8d, 0x77, 0xbf, 0x29, 0x2f, 0xa9, 0x3a, 0xca, 0xe0, 0x1e, 0xd4,
	0x75, 0x29, 0xcb, 0x8b, 0x51, 0xea, 0xc2, 0xdd, 0x72, 0xa9, 0x67, 0x7d, 0xf8, 0xb8, 0xa6, 0xfe,
	0xc5, 0x87, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x33, 0xa2, 0x3e, 0x22, 0x0b, 0x00, 0x00,
}