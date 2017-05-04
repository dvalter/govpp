// Package vpe represents the VPP binary API of the 'vpe' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api/vpe.api.json' on Thu, 27 Apr 2017 11:34:11 CEST.
package vpe

import "gerrit.fd.io/r/govpp.git/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xd3bcc288

// IP4FibCounter represents the VPP binary API data type 'ip4_fib_counter'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 3:
//
//        ["ip4_fib_counter",
//            ["u32", "address"],
//            ["u8", "address_length"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xb2739495"}
//        ],
//
type IP4FibCounter struct {
	Address       uint32
	AddressLength uint8
	Packets       uint64
	Bytes         uint64
}

func (*IP4FibCounter) GetTypeName() string {
	return "ip4_fib_counter"
}
func (*IP4FibCounter) GetCrcString() string {
	return "b2739495"
}

// IP4NbrCounter represents the VPP binary API data type 'ip4_nbr_counter'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 10:
//
//        ["ip4_nbr_counter",
//            ["u32", "address"],
//            ["u8", "link_type"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0x487e2e85"}
//        ],
//
type IP4NbrCounter struct {
	Address  uint32
	LinkType uint8
	Packets  uint64
	Bytes    uint64
}

func (*IP4NbrCounter) GetTypeName() string {
	return "ip4_nbr_counter"
}
func (*IP4NbrCounter) GetCrcString() string {
	return "487e2e85"
}

// IP6FibCounter represents the VPP binary API data type 'ip6_fib_counter'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 17:
//
//        ["ip6_fib_counter",
//            ["u64", "address", 2],
//            ["u8", "address_length"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xcf35769b"}
//        ],
//
type IP6FibCounter struct {
	Address       []uint64 `struc:"[2]uint64"`
	AddressLength uint8
	Packets       uint64
	Bytes         uint64
}

func (*IP6FibCounter) GetTypeName() string {
	return "ip6_fib_counter"
}
func (*IP6FibCounter) GetCrcString() string {
	return "cf35769b"
}

// IP6NbrCounter represents the VPP binary API data type 'ip6_nbr_counter'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 24:
//
//        ["ip6_nbr_counter",
//            ["u64", "address", 2],
//            ["u8", "link_type"],
//            ["u64", "packets"],
//            ["u64", "bytes"],
//            {"crc" : "0xefca741e"}
//        ]
//
type IP6NbrCounter struct {
	Address  []uint64 `struc:"[2]uint64"`
	LinkType uint8
	Packets  uint64
	Bytes    uint64
}

func (*IP6NbrCounter) GetTypeName() string {
	return "ip6_nbr_counter"
}
func (*IP6NbrCounter) GetCrcString() string {
	return "efca741e"
}

// CreateVlanSubif represents the VPP binary API message 'create_vlan_subif'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 33:
//
//        ["create_vlan_subif",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "vlan_id"],
//            {"crc" : "0xaf9ae1e9"}
//        ],
//
type CreateVlanSubif struct {
	SwIfIndex uint32
	VlanID    uint32
}

func (*CreateVlanSubif) GetMessageName() string {
	return "create_vlan_subif"
}
func (*CreateVlanSubif) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CreateVlanSubif) GetCrcString() string {
	return "af9ae1e9"
}
func NewCreateVlanSubif() api.Message {
	return &CreateVlanSubif{}
}

// CreateVlanSubifReply represents the VPP binary API message 'create_vlan_subif_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 41:
//
//        ["create_vlan_subif_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x8f36b888"}
//        ],
//
type CreateVlanSubifReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*CreateVlanSubifReply) GetMessageName() string {
	return "create_vlan_subif_reply"
}
func (*CreateVlanSubifReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CreateVlanSubifReply) GetCrcString() string {
	return "8f36b888"
}
func NewCreateVlanSubifReply() api.Message {
	return &CreateVlanSubifReply{}
}

// SwInterfaceSetMplsEnable represents the VPP binary API message 'sw_interface_set_mpls_enable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 48:
//
//        ["sw_interface_set_mpls_enable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "enable"],
//            {"crc" : "0x37f6357e"}
//        ],
//
type SwInterfaceSetMplsEnable struct {
	SwIfIndex uint32
	Enable    uint8
}

func (*SwInterfaceSetMplsEnable) GetMessageName() string {
	return "sw_interface_set_mpls_enable"
}
func (*SwInterfaceSetMplsEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetMplsEnable) GetCrcString() string {
	return "37f6357e"
}
func NewSwInterfaceSetMplsEnable() api.Message {
	return &SwInterfaceSetMplsEnable{}
}

// SwInterfaceSetMplsEnableReply represents the VPP binary API message 'sw_interface_set_mpls_enable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 56:
//
//        ["sw_interface_set_mpls_enable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5ffd3ca9"}
//        ],
//
type SwInterfaceSetMplsEnableReply struct {
	Retval int32
}

func (*SwInterfaceSetMplsEnableReply) GetMessageName() string {
	return "sw_interface_set_mpls_enable_reply"
}
func (*SwInterfaceSetMplsEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetMplsEnableReply) GetCrcString() string {
	return "5ffd3ca9"
}
func NewSwInterfaceSetMplsEnableReply() api.Message {
	return &SwInterfaceSetMplsEnableReply{}
}

// ProxyArpAddDel represents the VPP binary API message 'proxy_arp_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 62:
//
//        ["proxy_arp_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vrf_id"],
//            ["u8", "is_add"],
//            ["u8", "low_address", 4],
//            ["u8", "hi_address", 4],
//            {"crc" : "0x4bef9951"}
//        ],
//
type ProxyArpAddDel struct {
	VrfID      uint32
	IsAdd      uint8
	LowAddress []byte `struc:"[4]byte"`
	HiAddress  []byte `struc:"[4]byte"`
}

func (*ProxyArpAddDel) GetMessageName() string {
	return "proxy_arp_add_del"
}
func (*ProxyArpAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ProxyArpAddDel) GetCrcString() string {
	return "4bef9951"
}
func NewProxyArpAddDel() api.Message {
	return &ProxyArpAddDel{}
}

// ProxyArpAddDelReply represents the VPP binary API message 'proxy_arp_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 72:
//
//        ["proxy_arp_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8e2d621d"}
//        ],
//
type ProxyArpAddDelReply struct {
	Retval int32
}

func (*ProxyArpAddDelReply) GetMessageName() string {
	return "proxy_arp_add_del_reply"
}
func (*ProxyArpAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ProxyArpAddDelReply) GetCrcString() string {
	return "8e2d621d"
}
func NewProxyArpAddDelReply() api.Message {
	return &ProxyArpAddDelReply{}
}

// ProxyArpIntfcEnableDisable represents the VPP binary API message 'proxy_arp_intfc_enable_disable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 78:
//
//        ["proxy_arp_intfc_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "enable_disable"],
//            {"crc" : "0x3ee1998e"}
//        ],
//
type ProxyArpIntfcEnableDisable struct {
	SwIfIndex     uint32
	EnableDisable uint8
}

func (*ProxyArpIntfcEnableDisable) GetMessageName() string {
	return "proxy_arp_intfc_enable_disable"
}
func (*ProxyArpIntfcEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ProxyArpIntfcEnableDisable) GetCrcString() string {
	return "3ee1998e"
}
func NewProxyArpIntfcEnableDisable() api.Message {
	return &ProxyArpIntfcEnableDisable{}
}

// ProxyArpIntfcEnableDisableReply represents the VPP binary API message 'proxy_arp_intfc_enable_disable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 86:
//
//        ["proxy_arp_intfc_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x23d273cd"}
//        ],
//
type ProxyArpIntfcEnableDisableReply struct {
	Retval int32
}

func (*ProxyArpIntfcEnableDisableReply) GetMessageName() string {
	return "proxy_arp_intfc_enable_disable_reply"
}
func (*ProxyArpIntfcEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ProxyArpIntfcEnableDisableReply) GetCrcString() string {
	return "23d273cd"
}
func NewProxyArpIntfcEnableDisableReply() api.Message {
	return &ProxyArpIntfcEnableDisableReply{}
}

// ResetVrf represents the VPP binary API message 'reset_vrf'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 92:
//
//        ["reset_vrf",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u32", "vrf_id"],
//            {"crc" : "0xeb07deb0"}
//        ],
//
type ResetVrf struct {
	IsIpv6 uint8
	VrfID  uint32
}

func (*ResetVrf) GetMessageName() string {
	return "reset_vrf"
}
func (*ResetVrf) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ResetVrf) GetCrcString() string {
	return "eb07deb0"
}
func NewResetVrf() api.Message {
	return &ResetVrf{}
}

// ResetVrfReply represents the VPP binary API message 'reset_vrf_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 100:
//
//        ["reset_vrf_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x5f283863"}
//        ],
//
type ResetVrfReply struct {
	Retval int32
}

func (*ResetVrfReply) GetMessageName() string {
	return "reset_vrf_reply"
}
func (*ResetVrfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ResetVrfReply) GetCrcString() string {
	return "5f283863"
}
func NewResetVrfReply() api.Message {
	return &ResetVrfReply{}
}

// IsAddressReachable represents the VPP binary API message 'is_address_reachable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 106:
//
//        ["is_address_reachable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "next_hop_sw_if_index"],
//            ["u8", "is_known"],
//            ["u8", "is_ipv6"],
//            ["u8", "is_error"],
//            ["u8", "address", 16],
//            {"crc" : "0xa8b6e322"}
//        ],
//
type IsAddressReachable struct {
	NextHopSwIfIndex uint32
	IsKnown          uint8
	IsIpv6           uint8
	IsError          uint8
	Address          []byte `struc:"[16]byte"`
}

func (*IsAddressReachable) GetMessageName() string {
	return "is_address_reachable"
}
func (*IsAddressReachable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IsAddressReachable) GetCrcString() string {
	return "a8b6e322"
}
func NewIsAddressReachable() api.Message {
	return &IsAddressReachable{}
}

// WantStats represents the VPP binary API message 'want_stats'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 117:
//
//        ["want_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x4f2effb4"}
//        ],
//
type WantStats struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantStats) GetMessageName() string {
	return "want_stats"
}
func (*WantStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantStats) GetCrcString() string {
	return "4f2effb4"
}
func NewWantStats() api.Message {
	return &WantStats{}
}

// WantStatsReply represents the VPP binary API message 'want_stats_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 125:
//
//        ["want_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xb36abf5f"}
//        ],
//
type WantStatsReply struct {
	Retval int32
}

func (*WantStatsReply) GetMessageName() string {
	return "want_stats_reply"
}
func (*WantStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantStatsReply) GetCrcString() string {
	return "b36abf5f"
}
func NewWantStatsReply() api.Message {
	return &WantStatsReply{}
}

// VnetIP4FibCounters represents the VPP binary API message 'vnet_ip4_fib_counters'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 131:
//
//        ["vnet_ip4_fib_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "vrf_id"],
//            ["u32", "count"],
//            ["vl_api_ip4_fib_counter_t", "c", 0, "count"],
//            {"crc" : "0x1ab9d6c5"}
//        ],
//
type VnetIP4FibCounters struct {
	VrfID uint32
	Count uint32 `struc:"sizeof=C"`
	C     []IP4FibCounter
}

func (*VnetIP4FibCounters) GetMessageName() string {
	return "vnet_ip4_fib_counters"
}
func (*VnetIP4FibCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP4FibCounters) GetCrcString() string {
	return "1ab9d6c5"
}
func NewVnetIP4FibCounters() api.Message {
	return &VnetIP4FibCounters{}
}

// VnetIP4NbrCounters represents the VPP binary API message 'vnet_ip4_nbr_counters'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 138:
//
//        ["vnet_ip4_nbr_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "count"],
//            ["u32", "sw_if_index"],
//            ["u8", "begin"],
//            ["vl_api_ip4_nbr_counter_t", "c", 0, "count"],
//            {"crc" : "0xfc2b5092"}
//        ],
//
type VnetIP4NbrCounters struct {
	Count     uint32 `struc:"sizeof=C"`
	SwIfIndex uint32
	Begin     uint8
	C         []IP4NbrCounter
}

func (*VnetIP4NbrCounters) GetMessageName() string {
	return "vnet_ip4_nbr_counters"
}
func (*VnetIP4NbrCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP4NbrCounters) GetCrcString() string {
	return "fc2b5092"
}
func NewVnetIP4NbrCounters() api.Message {
	return &VnetIP4NbrCounters{}
}

// VnetIP6FibCounters represents the VPP binary API message 'vnet_ip6_fib_counters'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 146:
//
//        ["vnet_ip6_fib_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "vrf_id"],
//            ["u32", "count"],
//            ["vl_api_ip6_fib_counter_t", "c", 0, "count"],
//            {"crc" : "0x9ab453ae"}
//        ],
//
type VnetIP6FibCounters struct {
	VrfID uint32
	Count uint32 `struc:"sizeof=C"`
	C     []IP6FibCounter
}

func (*VnetIP6FibCounters) GetMessageName() string {
	return "vnet_ip6_fib_counters"
}
func (*VnetIP6FibCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP6FibCounters) GetCrcString() string {
	return "9ab453ae"
}
func NewVnetIP6FibCounters() api.Message {
	return &VnetIP6FibCounters{}
}

// VnetIP6NbrCounters represents the VPP binary API message 'vnet_ip6_nbr_counters'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 153:
//
//        ["vnet_ip6_nbr_counters",
//            ["u16", "_vl_msg_id"],
//            ["u32", "count"],
//            ["u32", "sw_if_index"],
//            ["u8", "begin"],
//            ["vl_api_ip6_nbr_counter_t", "c", 0, "count"],
//            {"crc" : "0x181b673f"}
//        ],
//
type VnetIP6NbrCounters struct {
	Count     uint32 `struc:"sizeof=C"`
	SwIfIndex uint32
	Begin     uint8
	C         []IP6NbrCounter
}

func (*VnetIP6NbrCounters) GetMessageName() string {
	return "vnet_ip6_nbr_counters"
}
func (*VnetIP6NbrCounters) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*VnetIP6NbrCounters) GetCrcString() string {
	return "181b673f"
}
func NewVnetIP6NbrCounters() api.Message {
	return &VnetIP6NbrCounters{}
}

// VnetGetSummaryStats represents the VPP binary API message 'vnet_get_summary_stats'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 161:
//
//        ["vnet_get_summary_stats",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0x16435c20"}
//        ],
//
type VnetGetSummaryStats struct {
}

func (*VnetGetSummaryStats) GetMessageName() string {
	return "vnet_get_summary_stats"
}
func (*VnetGetSummaryStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VnetGetSummaryStats) GetCrcString() string {
	return "16435c20"
}
func NewVnetGetSummaryStats() api.Message {
	return &VnetGetSummaryStats{}
}

// VnetSummaryStatsReply represents the VPP binary API message 'vnet_summary_stats_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 167:
//
//        ["vnet_summary_stats_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u64", "total_pkts", 2],
//            ["u64", "total_bytes", 2],
//            ["f64", "vector_rate"],
//            {"crc" : "0x87a8fa9f"}
//        ],
//
type VnetSummaryStatsReply struct {
	Retval     int32
	TotalPkts  []uint64 `struc:"[2]uint64"`
	TotalBytes []uint64 `struc:"[2]uint64"`
	VectorRate float64
}

func (*VnetSummaryStatsReply) GetMessageName() string {
	return "vnet_summary_stats_reply"
}
func (*VnetSummaryStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VnetSummaryStatsReply) GetCrcString() string {
	return "87a8fa9f"
}
func NewVnetSummaryStatsReply() api.Message {
	return &VnetSummaryStatsReply{}
}

// OamEvent represents the VPP binary API message 'oam_event'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 176:
//
//        ["oam_event",
//            ["u16", "_vl_msg_id"],
//            ["u8", "dst_address", 4],
//            ["u8", "state"],
//            {"crc" : "0x4f285ade"}
//        ],
//
type OamEvent struct {
	DstAddress []byte `struc:"[4]byte"`
	State      uint8
}

func (*OamEvent) GetMessageName() string {
	return "oam_event"
}
func (*OamEvent) GetMessageType() api.MessageType {
	return api.OtherMessage
}
func (*OamEvent) GetCrcString() string {
	return "4f285ade"
}
func NewOamEvent() api.Message {
	return &OamEvent{}
}

// WantOamEvents represents the VPP binary API message 'want_oam_events'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 182:
//
//        ["want_oam_events",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "enable_disable"],
//            ["u32", "pid"],
//            {"crc" : "0x948ef12a"}
//        ],
//
type WantOamEvents struct {
	EnableDisable uint32
	Pid           uint32
}

func (*WantOamEvents) GetMessageName() string {
	return "want_oam_events"
}
func (*WantOamEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantOamEvents) GetCrcString() string {
	return "948ef12a"
}
func NewWantOamEvents() api.Message {
	return &WantOamEvents{}
}

// WantOamEventsReply represents the VPP binary API message 'want_oam_events_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 190:
//
//        ["want_oam_events_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x266a677d"}
//        ],
//
type WantOamEventsReply struct {
	Retval int32
}

func (*WantOamEventsReply) GetMessageName() string {
	return "want_oam_events_reply"
}
func (*WantOamEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantOamEventsReply) GetCrcString() string {
	return "266a677d"
}
func NewWantOamEventsReply() api.Message {
	return &WantOamEventsReply{}
}

// OamAddDel represents the VPP binary API message 'oam_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 196:
//
//        ["oam_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vrf_id"],
//            ["u8", "src_address", 4],
//            ["u8", "dst_address", 4],
//            ["u8", "is_add"],
//            {"crc" : "0xb14bc7df"}
//        ],
//
type OamAddDel struct {
	VrfID      uint32
	SrcAddress []byte `struc:"[4]byte"`
	DstAddress []byte `struc:"[4]byte"`
	IsAdd      uint8
}

func (*OamAddDel) GetMessageName() string {
	return "oam_add_del"
}
func (*OamAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*OamAddDel) GetCrcString() string {
	return "b14bc7df"
}
func NewOamAddDel() api.Message {
	return &OamAddDel{}
}

// OamAddDelReply represents the VPP binary API message 'oam_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 206:
//
//        ["oam_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xc5594eec"}
//        ],
//
type OamAddDelReply struct {
	Retval int32
}

func (*OamAddDelReply) GetMessageName() string {
	return "oam_add_del_reply"
}
func (*OamAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*OamAddDelReply) GetCrcString() string {
	return "c5594eec"
}
func NewOamAddDelReply() api.Message {
	return &OamAddDelReply{}
}

// ResetFib represents the VPP binary API message 'reset_fib'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 212:
//
//        ["reset_fib",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "vrf_id"],
//            ["u8", "is_ipv6"],
//            {"crc" : "0x6f17106b"}
//        ],
//
type ResetFib struct {
	VrfID  uint32
	IsIpv6 uint8
}

func (*ResetFib) GetMessageName() string {
	return "reset_fib"
}
func (*ResetFib) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ResetFib) GetCrcString() string {
	return "6f17106b"
}
func NewResetFib() api.Message {
	return &ResetFib{}
}

// ResetFibReply represents the VPP binary API message 'reset_fib_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 220:
//
//        ["reset_fib_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x990dcbf8"}
//        ],
//
type ResetFibReply struct {
	Retval int32
}

func (*ResetFibReply) GetMessageName() string {
	return "reset_fib_reply"
}
func (*ResetFibReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ResetFibReply) GetCrcString() string {
	return "990dcbf8"
}
func NewResetFibReply() api.Message {
	return &ResetFibReply{}
}

// CreateLoopback represents the VPP binary API message 'create_loopback'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 226:
//
//        ["create_loopback",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "mac_address", 6],
//            {"crc" : "0xb2602de5"}
//        ],
//
type CreateLoopback struct {
	MacAddress []byte `struc:"[6]byte"`
}

func (*CreateLoopback) GetMessageName() string {
	return "create_loopback"
}
func (*CreateLoopback) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CreateLoopback) GetCrcString() string {
	return "b2602de5"
}
func NewCreateLoopback() api.Message {
	return &CreateLoopback{}
}

// CreateLoopbackReply represents the VPP binary API message 'create_loopback_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 233:
//
//        ["create_loopback_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x9520f804"}
//        ],
//
type CreateLoopbackReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*CreateLoopbackReply) GetMessageName() string {
	return "create_loopback_reply"
}
func (*CreateLoopbackReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CreateLoopbackReply) GetCrcString() string {
	return "9520f804"
}
func NewCreateLoopbackReply() api.Message {
	return &CreateLoopbackReply{}
}

// CreateLoopbackInstance represents the VPP binary API message 'create_loopback_instance'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 240:
//
//        ["create_loopback_instance",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "mac_address", 6],
//            ["u8", "is_specified"],
//            ["u32", "user_instance"],
//            {"crc" : "0x967694f1"}
//        ],
//
type CreateLoopbackInstance struct {
	MacAddress   []byte `struc:"[6]byte"`
	IsSpecified  uint8
	UserInstance uint32
}

func (*CreateLoopbackInstance) GetMessageName() string {
	return "create_loopback_instance"
}
func (*CreateLoopbackInstance) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CreateLoopbackInstance) GetCrcString() string {
	return "967694f1"
}
func NewCreateLoopbackInstance() api.Message {
	return &CreateLoopbackInstance{}
}

// CreateLoopbackInstanceReply represents the VPP binary API message 'create_loopback_instance_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 249:
//
//        ["create_loopback_instance_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xd52c63b6"}
//        ],
//
type CreateLoopbackInstanceReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*CreateLoopbackInstanceReply) GetMessageName() string {
	return "create_loopback_instance_reply"
}
func (*CreateLoopbackInstanceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CreateLoopbackInstanceReply) GetCrcString() string {
	return "d52c63b6"
}
func NewCreateLoopbackInstanceReply() api.Message {
	return &CreateLoopbackInstanceReply{}
}

// DeleteLoopback represents the VPP binary API message 'delete_loopback'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 256:
//
//        ["delete_loopback",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0xded428b0"}
//        ],
//
type DeleteLoopback struct {
	SwIfIndex uint32
}

func (*DeleteLoopback) GetMessageName() string {
	return "delete_loopback"
}
func (*DeleteLoopback) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DeleteLoopback) GetCrcString() string {
	return "ded428b0"
}
func NewDeleteLoopback() api.Message {
	return &DeleteLoopback{}
}

// DeleteLoopbackReply represents the VPP binary API message 'delete_loopback_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 263:
//
//        ["delete_loopback_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xc91dafa5"}
//        ],
//
type DeleteLoopbackReply struct {
	Retval int32
}

func (*DeleteLoopbackReply) GetMessageName() string {
	return "delete_loopback_reply"
}
func (*DeleteLoopbackReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DeleteLoopbackReply) GetCrcString() string {
	return "c91dafa5"
}
func NewDeleteLoopbackReply() api.Message {
	return &DeleteLoopbackReply{}
}

// ControlPing represents the VPP binary API message 'control_ping'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 269:
//
//        ["control_ping",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xea1bf4f7"}
//        ],
//
type ControlPing struct {
}

func (*ControlPing) GetMessageName() string {
	return "control_ping"
}
func (*ControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ControlPing) GetCrcString() string {
	return "ea1bf4f7"
}
func NewControlPing() api.Message {
	return &ControlPing{}
}

// ControlPingReply represents the VPP binary API message 'control_ping_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 275:
//
//        ["control_ping_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "client_index"],
//            ["u32", "vpe_pid"],
//            {"crc" : "0xaa016e7b"}
//        ],
//
type ControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePid      uint32
}

func (*ControlPingReply) GetMessageName() string {
	return "control_ping_reply"
}
func (*ControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ControlPingReply) GetCrcString() string {
	return "aa016e7b"
}
func NewControlPingReply() api.Message {
	return &ControlPingReply{}
}

// CliRequest represents the VPP binary API message 'cli_request'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 283:
//
//        ["cli_request",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u64", "cmd_in_shmem"],
//            {"crc" : "0xfef056d0"}
//        ],
//
type CliRequest struct {
	CmdInShmem uint64
}

func (*CliRequest) GetMessageName() string {
	return "cli_request"
}
func (*CliRequest) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CliRequest) GetCrcString() string {
	return "fef056d0"
}
func NewCliRequest() api.Message {
	return &CliRequest{}
}

// CliInband represents the VPP binary API message 'cli_inband'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 290:
//
//        ["cli_inband",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "length"],
//            ["u8", "cmd", 0, "length"],
//            {"crc" : "0x22345937"}
//        ],
//
type CliInband struct {
	Length uint32 `struc:"sizeof=Cmd"`
	Cmd    []byte
}

func (*CliInband) GetMessageName() string {
	return "cli_inband"
}
func (*CliInband) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CliInband) GetCrcString() string {
	return "22345937"
}
func NewCliInband() api.Message {
	return &CliInband{}
}

// CliReply represents the VPP binary API message 'cli_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 298:
//
//        ["cli_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u64", "reply_in_shmem"],
//            {"crc" : "0x594a0b2e"}
//        ],
//
type CliReply struct {
	Retval       int32
	ReplyInShmem uint64
}

func (*CliReply) GetMessageName() string {
	return "cli_reply"
}
func (*CliReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CliReply) GetCrcString() string {
	return "594a0b2e"
}
func NewCliReply() api.Message {
	return &CliReply{}
}

// CliInbandReply represents the VPP binary API message 'cli_inband_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 305:
//
//        ["cli_inband_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "length"],
//            ["u8", "reply", 0, "length"],
//            {"crc" : "0xc1835761"}
//        ],
//
type CliInbandReply struct {
	Retval int32
	Length uint32 `struc:"sizeof=Reply"`
	Reply  []byte
}

func (*CliInbandReply) GetMessageName() string {
	return "cli_inband_reply"
}
func (*CliInbandReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CliInbandReply) GetCrcString() string {
	return "c1835761"
}
func NewCliInbandReply() api.Message {
	return &CliInbandReply{}
}

// SetArpNeighborLimit represents the VPP binary API message 'set_arp_neighbor_limit'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 313:
//
//        ["set_arp_neighbor_limit",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u32", "arp_neighbor_limit"],
//            {"crc" : "0xc1690cb4"}
//        ],
//
type SetArpNeighborLimit struct {
	IsIpv6           uint8
	ArpNeighborLimit uint32
}

func (*SetArpNeighborLimit) GetMessageName() string {
	return "set_arp_neighbor_limit"
}
func (*SetArpNeighborLimit) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SetArpNeighborLimit) GetCrcString() string {
	return "c1690cb4"
}
func NewSetArpNeighborLimit() api.Message {
	return &SetArpNeighborLimit{}
}

// SetArpNeighborLimitReply represents the VPP binary API message 'set_arp_neighbor_limit_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 321:
//
//        ["set_arp_neighbor_limit_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa6b30518"}
//        ],
//
type SetArpNeighborLimitReply struct {
	Retval int32
}

func (*SetArpNeighborLimitReply) GetMessageName() string {
	return "set_arp_neighbor_limit_reply"
}
func (*SetArpNeighborLimitReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SetArpNeighborLimitReply) GetCrcString() string {
	return "a6b30518"
}
func NewSetArpNeighborLimitReply() api.Message {
	return &SetArpNeighborLimitReply{}
}

// L2PatchAddDel represents the VPP binary API message 'l2_patch_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 327:
//
//        ["l2_patch_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "rx_sw_if_index"],
//            ["u32", "tx_sw_if_index"],
//            ["u8", "is_add"],
//            {"crc" : "0x9b10029a"}
//        ],
//
type L2PatchAddDel struct {
	RxSwIfIndex uint32
	TxSwIfIndex uint32
	IsAdd       uint8
}

func (*L2PatchAddDel) GetMessageName() string {
	return "l2_patch_add_del"
}
func (*L2PatchAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2PatchAddDel) GetCrcString() string {
	return "9b10029a"
}
func NewL2PatchAddDel() api.Message {
	return &L2PatchAddDel{}
}

// L2PatchAddDelReply represents the VPP binary API message 'l2_patch_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 336:
//
//        ["l2_patch_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xa85e37be"}
//        ],
//
type L2PatchAddDelReply struct {
	Retval int32
}

func (*L2PatchAddDelReply) GetMessageName() string {
	return "l2_patch_add_del_reply"
}
func (*L2PatchAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2PatchAddDelReply) GetCrcString() string {
	return "a85e37be"
}
func NewL2PatchAddDelReply() api.Message {
	return &L2PatchAddDelReply{}
}

// SwInterfaceSetVpath represents the VPP binary API message 'sw_interface_set_vpath'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 342:
//
//        ["sw_interface_set_vpath",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "enable"],
//            {"crc" : "0x1bc2fd5e"}
//        ],
//
type SwInterfaceSetVpath struct {
	SwIfIndex uint32
	Enable    uint8
}

func (*SwInterfaceSetVpath) GetMessageName() string {
	return "sw_interface_set_vpath"
}
func (*SwInterfaceSetVpath) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetVpath) GetCrcString() string {
	return "1bc2fd5e"
}
func NewSwInterfaceSetVpath() api.Message {
	return &SwInterfaceSetVpath{}
}

// SwInterfaceSetVpathReply represents the VPP binary API message 'sw_interface_set_vpath_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 350:
//
//        ["sw_interface_set_vpath_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x828dbe62"}
//        ],
//
type SwInterfaceSetVpathReply struct {
	Retval int32
}

func (*SwInterfaceSetVpathReply) GetMessageName() string {
	return "sw_interface_set_vpath_reply"
}
func (*SwInterfaceSetVpathReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetVpathReply) GetCrcString() string {
	return "828dbe62"
}
func NewSwInterfaceSetVpathReply() api.Message {
	return &SwInterfaceSetVpathReply{}
}

// SwInterfaceSetL2Xconnect represents the VPP binary API message 'sw_interface_set_l2_xconnect'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 356:
//
//        ["sw_interface_set_l2_xconnect",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "rx_sw_if_index"],
//            ["u32", "tx_sw_if_index"],
//            ["u8", "enable"],
//            {"crc" : "0x48a4c4c8"}
//        ],
//
type SwInterfaceSetL2Xconnect struct {
	RxSwIfIndex uint32
	TxSwIfIndex uint32
	Enable      uint8
}

func (*SwInterfaceSetL2Xconnect) GetMessageName() string {
	return "sw_interface_set_l2_xconnect"
}
func (*SwInterfaceSetL2Xconnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetL2Xconnect) GetCrcString() string {
	return "48a4c4c8"
}
func NewSwInterfaceSetL2Xconnect() api.Message {
	return &SwInterfaceSetL2Xconnect{}
}

// SwInterfaceSetL2XconnectReply represents the VPP binary API message 'sw_interface_set_l2_xconnect_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 365:
//
//        ["sw_interface_set_l2_xconnect_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x6e45eed4"}
//        ],
//
type SwInterfaceSetL2XconnectReply struct {
	Retval int32
}

func (*SwInterfaceSetL2XconnectReply) GetMessageName() string {
	return "sw_interface_set_l2_xconnect_reply"
}
func (*SwInterfaceSetL2XconnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetL2XconnectReply) GetCrcString() string {
	return "6e45eed4"
}
func NewSwInterfaceSetL2XconnectReply() api.Message {
	return &SwInterfaceSetL2XconnectReply{}
}

// SwInterfaceSetL2Bridge represents the VPP binary API message 'sw_interface_set_l2_bridge'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 371:
//
//        ["sw_interface_set_l2_bridge",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "rx_sw_if_index"],
//            ["u32", "bd_id"],
//            ["u8", "shg"],
//            ["u8", "bvi"],
//            ["u8", "enable"],
//            {"crc" : "0x36c739e8"}
//        ],
//
type SwInterfaceSetL2Bridge struct {
	RxSwIfIndex uint32
	BdID        uint32
	Shg         uint8
	Bvi         uint8
	Enable      uint8
}

func (*SwInterfaceSetL2Bridge) GetMessageName() string {
	return "sw_interface_set_l2_bridge"
}
func (*SwInterfaceSetL2Bridge) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetL2Bridge) GetCrcString() string {
	return "36c739e8"
}
func NewSwInterfaceSetL2Bridge() api.Message {
	return &SwInterfaceSetL2Bridge{}
}

// SwInterfaceSetL2BridgeReply represents the VPP binary API message 'sw_interface_set_l2_bridge_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 382:
//
//        ["sw_interface_set_l2_bridge_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x347e08d9"}
//        ],
//
type SwInterfaceSetL2BridgeReply struct {
	Retval int32
}

func (*SwInterfaceSetL2BridgeReply) GetMessageName() string {
	return "sw_interface_set_l2_bridge_reply"
}
func (*SwInterfaceSetL2BridgeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetL2BridgeReply) GetCrcString() string {
	return "347e08d9"
}
func NewSwInterfaceSetL2BridgeReply() api.Message {
	return &SwInterfaceSetL2BridgeReply{}
}

// BdIPMacAddDel represents the VPP binary API message 'bd_ip_mac_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 388:
//
//        ["bd_ip_mac_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "bd_id"],
//            ["u8", "is_add"],
//            ["u8", "is_ipv6"],
//            ["u8", "ip_address", 16],
//            ["u8", "mac_address", 6],
//            {"crc" : "0xad819817"}
//        ],
//
type BdIPMacAddDel struct {
	BdID       uint32
	IsAdd      uint8
	IsIpv6     uint8
	IPAddress  []byte `struc:"[16]byte"`
	MacAddress []byte `struc:"[6]byte"`
}

func (*BdIPMacAddDel) GetMessageName() string {
	return "bd_ip_mac_add_del"
}
func (*BdIPMacAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*BdIPMacAddDel) GetCrcString() string {
	return "ad819817"
}
func NewBdIPMacAddDel() api.Message {
	return &BdIPMacAddDel{}
}

// BdIPMacAddDelReply represents the VPP binary API message 'bd_ip_mac_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 399:
//
//        ["bd_ip_mac_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x55bab3b4"}
//        ],
//
type BdIPMacAddDelReply struct {
	Retval int32
}

func (*BdIPMacAddDelReply) GetMessageName() string {
	return "bd_ip_mac_add_del_reply"
}
func (*BdIPMacAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*BdIPMacAddDelReply) GetCrcString() string {
	return "55bab3b4"
}
func NewBdIPMacAddDelReply() api.Message {
	return &BdIPMacAddDelReply{}
}

// ClassifySetInterfaceIPTable represents the VPP binary API message 'classify_set_interface_ip_table'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 405:
//
//        ["classify_set_interface_ip_table",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u32", "sw_if_index"],
//            ["u32", "table_index"],
//            {"crc" : "0x0dc45308"}
//        ],
//
type ClassifySetInterfaceIPTable struct {
	IsIpv6     uint8
	SwIfIndex  uint32
	TableIndex uint32
}

func (*ClassifySetInterfaceIPTable) GetMessageName() string {
	return "classify_set_interface_ip_table"
}
func (*ClassifySetInterfaceIPTable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifySetInterfaceIPTable) GetCrcString() string {
	return "0dc45308"
}
func NewClassifySetInterfaceIPTable() api.Message {
	return &ClassifySetInterfaceIPTable{}
}

// ClassifySetInterfaceIPTableReply represents the VPP binary API message 'classify_set_interface_ip_table_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 414:
//
//        ["classify_set_interface_ip_table_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xdc391c34"}
//        ],
//
type ClassifySetInterfaceIPTableReply struct {
	Retval int32
}

func (*ClassifySetInterfaceIPTableReply) GetMessageName() string {
	return "classify_set_interface_ip_table_reply"
}
func (*ClassifySetInterfaceIPTableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifySetInterfaceIPTableReply) GetCrcString() string {
	return "dc391c34"
}
func NewClassifySetInterfaceIPTableReply() api.Message {
	return &ClassifySetInterfaceIPTableReply{}
}

// ClassifySetInterfaceL2Tables represents the VPP binary API message 'classify_set_interface_l2_tables'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 420:
//
//        ["classify_set_interface_l2_tables",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "ip4_table_index"],
//            ["u32", "ip6_table_index"],
//            ["u32", "other_table_index"],
//            ["u8", "is_input"],
//            {"crc" : "0xed9ccf0d"}
//        ],
//
type ClassifySetInterfaceL2Tables struct {
	SwIfIndex       uint32
	IP4TableIndex   uint32
	IP6TableIndex   uint32
	OtherTableIndex uint32
	IsInput         uint8
}

func (*ClassifySetInterfaceL2Tables) GetMessageName() string {
	return "classify_set_interface_l2_tables"
}
func (*ClassifySetInterfaceL2Tables) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ClassifySetInterfaceL2Tables) GetCrcString() string {
	return "ed9ccf0d"
}
func NewClassifySetInterfaceL2Tables() api.Message {
	return &ClassifySetInterfaceL2Tables{}
}

// ClassifySetInterfaceL2TablesReply represents the VPP binary API message 'classify_set_interface_l2_tables_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 431:
//
//        ["classify_set_interface_l2_tables_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x8df20579"}
//        ],
//
type ClassifySetInterfaceL2TablesReply struct {
	Retval int32
}

func (*ClassifySetInterfaceL2TablesReply) GetMessageName() string {
	return "classify_set_interface_l2_tables_reply"
}
func (*ClassifySetInterfaceL2TablesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ClassifySetInterfaceL2TablesReply) GetCrcString() string {
	return "8df20579"
}
func NewClassifySetInterfaceL2TablesReply() api.Message {
	return &ClassifySetInterfaceL2TablesReply{}
}

// GetNodeIndex represents the VPP binary API message 'get_node_index'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 437:
//
//        ["get_node_index",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "node_name", 64],
//            {"crc" : "0x226d3f8c"}
//        ],
//
type GetNodeIndex struct {
	NodeName []byte `struc:"[64]byte"`
}

func (*GetNodeIndex) GetMessageName() string {
	return "get_node_index"
}
func (*GetNodeIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GetNodeIndex) GetCrcString() string {
	return "226d3f8c"
}
func NewGetNodeIndex() api.Message {
	return &GetNodeIndex{}
}

// GetNodeIndexReply represents the VPP binary API message 'get_node_index_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 444:
//
//        ["get_node_index_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "node_index"],
//            {"crc" : "0x29116865"}
//        ],
//
type GetNodeIndexReply struct {
	Retval    int32
	NodeIndex uint32
}

func (*GetNodeIndexReply) GetMessageName() string {
	return "get_node_index_reply"
}
func (*GetNodeIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GetNodeIndexReply) GetCrcString() string {
	return "29116865"
}
func NewGetNodeIndexReply() api.Message {
	return &GetNodeIndexReply{}
}

// AddNodeNext represents the VPP binary API message 'add_node_next'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 451:
//
//        ["add_node_next",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "node_name", 64],
//            ["u8", "next_name", 64],
//            {"crc" : "0xe4202993"}
//        ],
//
type AddNodeNext struct {
	NodeName []byte `struc:"[64]byte"`
	NextName []byte `struc:"[64]byte"`
}

func (*AddNodeNext) GetMessageName() string {
	return "add_node_next"
}
func (*AddNodeNext) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*AddNodeNext) GetCrcString() string {
	return "e4202993"
}
func NewAddNodeNext() api.Message {
	return &AddNodeNext{}
}

// AddNodeNextReply represents the VPP binary API message 'add_node_next_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 459:
//
//        ["add_node_next_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "next_index"],
//            {"crc" : "0xe89d6eed"}
//        ],
//
type AddNodeNextReply struct {
	Retval    int32
	NextIndex uint32
}

func (*AddNodeNextReply) GetMessageName() string {
	return "add_node_next_reply"
}
func (*AddNodeNextReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*AddNodeNextReply) GetCrcString() string {
	return "e89d6eed"
}
func NewAddNodeNextReply() api.Message {
	return &AddNodeNextReply{}
}

// L2InterfaceEfpFilter represents the VPP binary API message 'l2_interface_efp_filter'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 466:
//
//        ["l2_interface_efp_filter",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "enable_disable"],
//            {"crc" : "0x07c9d601"}
//        ],
//
type L2InterfaceEfpFilter struct {
	SwIfIndex     uint32
	EnableDisable uint32
}

func (*L2InterfaceEfpFilter) GetMessageName() string {
	return "l2_interface_efp_filter"
}
func (*L2InterfaceEfpFilter) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*L2InterfaceEfpFilter) GetCrcString() string {
	return "07c9d601"
}
func NewL2InterfaceEfpFilter() api.Message {
	return &L2InterfaceEfpFilter{}
}

// L2InterfaceEfpFilterReply represents the VPP binary API message 'l2_interface_efp_filter_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 474:
//
//        ["l2_interface_efp_filter_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x0f4bb0c0"}
//        ],
//
type L2InterfaceEfpFilterReply struct {
	Retval int32
}

func (*L2InterfaceEfpFilterReply) GetMessageName() string {
	return "l2_interface_efp_filter_reply"
}
func (*L2InterfaceEfpFilterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*L2InterfaceEfpFilterReply) GetCrcString() string {
	return "0f4bb0c0"
}
func NewL2InterfaceEfpFilterReply() api.Message {
	return &L2InterfaceEfpFilterReply{}
}

// CreateSubif represents the VPP binary API message 'create_subif'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 480:
//
//        ["create_subif",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "sub_id"],
//            ["u8", "no_tags"],
//            ["u8", "one_tag"],
//            ["u8", "two_tags"],
//            ["u8", "dot1ad"],
//            ["u8", "exact_match"],
//            ["u8", "default_sub"],
//            ["u8", "outer_vlan_id_any"],
//            ["u8", "inner_vlan_id_any"],
//            ["u16", "outer_vlan_id"],
//            ["u16", "inner_vlan_id"],
//            {"crc" : "0x150e6757"}
//        ],
//
type CreateSubif struct {
	SwIfIndex      uint32
	SubID          uint32
	NoTags         uint8
	OneTag         uint8
	TwoTags        uint8
	Dot1ad         uint8
	ExactMatch     uint8
	DefaultSub     uint8
	OuterVlanIDAny uint8
	InnerVlanIDAny uint8
	OuterVlanID    uint16
	InnerVlanID    uint16
}

func (*CreateSubif) GetMessageName() string {
	return "create_subif"
}
func (*CreateSubif) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*CreateSubif) GetCrcString() string {
	return "150e6757"
}
func NewCreateSubif() api.Message {
	return &CreateSubif{}
}

// CreateSubifReply represents the VPP binary API message 'create_subif_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 498:
//
//        ["create_subif_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x92272bcb"}
//        ],
//
type CreateSubifReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*CreateSubifReply) GetMessageName() string {
	return "create_subif_reply"
}
func (*CreateSubifReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*CreateSubifReply) GetCrcString() string {
	return "92272bcb"
}
func NewCreateSubifReply() api.Message {
	return &CreateSubifReply{}
}

// ShowVersion represents the VPP binary API message 'show_version'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 505:
//
//        ["show_version",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xf18f9480"}
//        ],
//
type ShowVersion struct {
}

func (*ShowVersion) GetMessageName() string {
	return "show_version"
}
func (*ShowVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ShowVersion) GetCrcString() string {
	return "f18f9480"
}
func NewShowVersion() api.Message {
	return &ShowVersion{}
}

// ShowVersionReply represents the VPP binary API message 'show_version_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 511:
//
//        ["show_version_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u8", "program", 32],
//            ["u8", "version", 32],
//            ["u8", "build_date", 32],
//            ["u8", "build_directory", 256],
//            {"crc" : "0x83186d9e"}
//        ],
//
type ShowVersionReply struct {
	Retval         int32
	Program        []byte `struc:"[32]byte"`
	Version        []byte `struc:"[32]byte"`
	BuildDate      []byte `struc:"[32]byte"`
	BuildDirectory []byte `struc:"[256]byte"`
}

func (*ShowVersionReply) GetMessageName() string {
	return "show_version_reply"
}
func (*ShowVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ShowVersionReply) GetCrcString() string {
	return "83186d9e"
}
func NewShowVersionReply() api.Message {
	return &ShowVersionReply{}
}

// InterfaceNameRenumber represents the VPP binary API message 'interface_name_renumber'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 521:
//
//        ["interface_name_renumber",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "new_show_dev_instance"],
//            {"crc" : "0x11b7bcec"}
//        ],
//
type InterfaceNameRenumber struct {
	SwIfIndex          uint32
	NewShowDevInstance uint32
}

func (*InterfaceNameRenumber) GetMessageName() string {
	return "interface_name_renumber"
}
func (*InterfaceNameRenumber) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*InterfaceNameRenumber) GetCrcString() string {
	return "11b7bcec"
}
func NewInterfaceNameRenumber() api.Message {
	return &InterfaceNameRenumber{}
}

// InterfaceNameRenumberReply represents the VPP binary API message 'interface_name_renumber_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 529:
//
//        ["interface_name_renumber_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x31594963"}
//        ],
//
type InterfaceNameRenumberReply struct {
	Retval int32
}

func (*InterfaceNameRenumberReply) GetMessageName() string {
	return "interface_name_renumber_reply"
}
func (*InterfaceNameRenumberReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*InterfaceNameRenumberReply) GetCrcString() string {
	return "31594963"
}
func NewInterfaceNameRenumberReply() api.Message {
	return &InterfaceNameRenumberReply{}
}

// WantIP4ArpEvents represents the VPP binary API message 'want_ip4_arp_events'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 535:
//
//        ["want_ip4_arp_events",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "enable_disable"],
//            ["u32", "pid"],
//            ["u32", "address"],
//            {"crc" : "0x5ae044c2"}
//        ],
//
type WantIP4ArpEvents struct {
	EnableDisable uint8
	Pid           uint32
	Address       uint32
}

func (*WantIP4ArpEvents) GetMessageName() string {
	return "want_ip4_arp_events"
}
func (*WantIP4ArpEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP4ArpEvents) GetCrcString() string {
	return "5ae044c2"
}
func NewWantIP4ArpEvents() api.Message {
	return &WantIP4ArpEvents{}
}

// WantIP4ArpEventsReply represents the VPP binary API message 'want_ip4_arp_events_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 544:
//
//        ["want_ip4_arp_events_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xe1c0b59e"}
//        ],
//
type WantIP4ArpEventsReply struct {
	Retval int32
}

func (*WantIP4ArpEventsReply) GetMessageName() string {
	return "want_ip4_arp_events_reply"
}
func (*WantIP4ArpEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP4ArpEventsReply) GetCrcString() string {
	return "e1c0b59e"
}
func NewWantIP4ArpEventsReply() api.Message {
	return &WantIP4ArpEventsReply{}
}

// IP4ArpEvent represents the VPP binary API message 'ip4_arp_event'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 550:
//
//        ["ip4_arp_event",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "address"],
//            ["u32", "pid"],
//            ["u32", "sw_if_index"],
//            ["u8", "new_mac", 6],
//            ["u8", "mac_ip"],
//            {"crc" : "0x7de1837b"}
//        ],
//
type IP4ArpEvent struct {
	Address   uint32
	Pid       uint32
	SwIfIndex uint32
	NewMac    []byte `struc:"[6]byte"`
	MacIP     uint8
}

func (*IP4ArpEvent) GetMessageName() string {
	return "ip4_arp_event"
}
func (*IP4ArpEvent) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IP4ArpEvent) GetCrcString() string {
	return "7de1837b"
}
func NewIP4ArpEvent() api.Message {
	return &IP4ArpEvent{}
}

// WantIP6NdEvents represents the VPP binary API message 'want_ip6_nd_events'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 561:
//
//        ["want_ip6_nd_events",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "enable_disable"],
//            ["u32", "pid"],
//            ["u8", "address", 16],
//            {"crc" : "0x9586ba55"}
//        ],
//
type WantIP6NdEvents struct {
	EnableDisable uint8
	Pid           uint32
	Address       []byte `struc:"[16]byte"`
}

func (*WantIP6NdEvents) GetMessageName() string {
	return "want_ip6_nd_events"
}
func (*WantIP6NdEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*WantIP6NdEvents) GetCrcString() string {
	return "9586ba55"
}
func NewWantIP6NdEvents() api.Message {
	return &WantIP6NdEvents{}
}

// WantIP6NdEventsReply represents the VPP binary API message 'want_ip6_nd_events_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 570:
//
//        ["want_ip6_nd_events_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x95458aad"}
//        ],
//
type WantIP6NdEventsReply struct {
	Retval int32
}

func (*WantIP6NdEventsReply) GetMessageName() string {
	return "want_ip6_nd_events_reply"
}
func (*WantIP6NdEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*WantIP6NdEventsReply) GetCrcString() string {
	return "95458aad"
}
func NewWantIP6NdEventsReply() api.Message {
	return &WantIP6NdEventsReply{}
}

// IP6NdEvent represents the VPP binary API message 'ip6_nd_event'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 576:
//
//        ["ip6_nd_event",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "pid"],
//            ["u32", "sw_if_index"],
//            ["u8", "address", 16],
//            ["u8", "new_mac", 6],
//            ["u8", "mac_ip"],
//            {"crc" : "0x777bb71c"}
//        ],
//
type IP6NdEvent struct {
	Pid       uint32
	SwIfIndex uint32
	Address   []byte `struc:"[16]byte"`
	NewMac    []byte `struc:"[6]byte"`
	MacIP     uint8
}

func (*IP6NdEvent) GetMessageName() string {
	return "ip6_nd_event"
}
func (*IP6NdEvent) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IP6NdEvent) GetCrcString() string {
	return "777bb71c"
}
func NewIP6NdEvent() api.Message {
	return &IP6NdEvent{}
}

// InputACLSetInterface represents the VPP binary API message 'input_acl_set_interface'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 587:
//
//        ["input_acl_set_interface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u32", "ip4_table_index"],
//            ["u32", "ip6_table_index"],
//            ["u32", "l2_table_index"],
//            ["u8", "is_add"],
//            {"crc" : "0x34d2fc33"}
//        ],
//
type InputACLSetInterface struct {
	SwIfIndex     uint32
	IP4TableIndex uint32
	IP6TableIndex uint32
	L2TableIndex  uint32
	IsAdd         uint8
}

func (*InputACLSetInterface) GetMessageName() string {
	return "input_acl_set_interface"
}
func (*InputACLSetInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*InputACLSetInterface) GetCrcString() string {
	return "34d2fc33"
}
func NewInputACLSetInterface() api.Message {
	return &InputACLSetInterface{}
}

// InputACLSetInterfaceReply represents the VPP binary API message 'input_acl_set_interface_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 598:
//
//        ["input_acl_set_interface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xba0110e3"}
//        ],
//
type InputACLSetInterfaceReply struct {
	Retval int32
}

func (*InputACLSetInterfaceReply) GetMessageName() string {
	return "input_acl_set_interface_reply"
}
func (*InputACLSetInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*InputACLSetInterfaceReply) GetCrcString() string {
	return "ba0110e3"
}
func NewInputACLSetInterfaceReply() api.Message {
	return &InputACLSetInterfaceReply{}
}

// GetNodeGraph represents the VPP binary API message 'get_node_graph'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 604:
//
//        ["get_node_graph",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            {"crc" : "0xf8636a76"}
//        ],
//
type GetNodeGraph struct {
}

func (*GetNodeGraph) GetMessageName() string {
	return "get_node_graph"
}
func (*GetNodeGraph) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GetNodeGraph) GetCrcString() string {
	return "f8636a76"
}
func NewGetNodeGraph() api.Message {
	return &GetNodeGraph{}
}

// GetNodeGraphReply represents the VPP binary API message 'get_node_graph_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 610:
//
//        ["get_node_graph_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u64", "reply_in_shmem"],
//            {"crc" : "0x816d91b6"}
//        ],
//
type GetNodeGraphReply struct {
	Retval       int32
	ReplyInShmem uint64
}

func (*GetNodeGraphReply) GetMessageName() string {
	return "get_node_graph_reply"
}
func (*GetNodeGraphReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GetNodeGraphReply) GetCrcString() string {
	return "816d91b6"
}
func NewGetNodeGraphReply() api.Message {
	return &GetNodeGraphReply{}
}

// IoamEnable represents the VPP binary API message 'ioam_enable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 617:
//
//        ["ioam_enable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u16", "id"],
//            ["u8", "seqno"],
//            ["u8", "analyse"],
//            ["u8", "pot_enable"],
//            ["u8", "trace_enable"],
//            ["u32", "node_id"],
//            {"crc" : "0x7bd4abf9"}
//        ],
//
type IoamEnable struct {
	ID          uint16
	Seqno       uint8
	Analyse     uint8
	PotEnable   uint8
	TraceEnable uint8
	NodeID      uint32
}

func (*IoamEnable) GetMessageName() string {
	return "ioam_enable"
}
func (*IoamEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IoamEnable) GetCrcString() string {
	return "7bd4abf9"
}
func NewIoamEnable() api.Message {
	return &IoamEnable{}
}

// IoamEnableReply represents the VPP binary API message 'ioam_enable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 629:
//
//        ["ioam_enable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x58a8fedc"}
//        ],
//
type IoamEnableReply struct {
	Retval int32
}

func (*IoamEnableReply) GetMessageName() string {
	return "ioam_enable_reply"
}
func (*IoamEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IoamEnableReply) GetCrcString() string {
	return "58a8fedc"
}
func NewIoamEnableReply() api.Message {
	return &IoamEnableReply{}
}

// IoamDisable represents the VPP binary API message 'ioam_disable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 635:
//
//        ["ioam_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u16", "id"],
//            {"crc" : "0xaff26d33"}
//        ],
//
type IoamDisable struct {
	ID uint16
}

func (*IoamDisable) GetMessageName() string {
	return "ioam_disable"
}
func (*IoamDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IoamDisable) GetCrcString() string {
	return "aff26d33"
}
func NewIoamDisable() api.Message {
	return &IoamDisable{}
}

// IoamDisableReply represents the VPP binary API message 'ioam_disable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 642:
//
//        ["ioam_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xef118a9d"}
//        ],
//
type IoamDisableReply struct {
	Retval int32
}

func (*IoamDisableReply) GetMessageName() string {
	return "ioam_disable_reply"
}
func (*IoamDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IoamDisableReply) GetCrcString() string {
	return "ef118a9d"
}
func NewIoamDisableReply() api.Message {
	return &IoamDisableReply{}
}

// GetNextIndex represents the VPP binary API message 'get_next_index'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 648:
//
//        ["get_next_index",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "node_name", 64],
//            ["u8", "next_name", 64],
//            {"crc" : "0x52f0e416"}
//        ],
//
type GetNextIndex struct {
	NodeName []byte `struc:"[64]byte"`
	NextName []byte `struc:"[64]byte"`
}

func (*GetNextIndex) GetMessageName() string {
	return "get_next_index"
}
func (*GetNextIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*GetNextIndex) GetCrcString() string {
	return "52f0e416"
}
func NewGetNextIndex() api.Message {
	return &GetNextIndex{}
}

// GetNextIndexReply represents the VPP binary API message 'get_next_index_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 656:
//
//        ["get_next_index_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "next_index"],
//            {"crc" : "0x671fbdb1"}
//        ],
//
type GetNextIndexReply struct {
	Retval    int32
	NextIndex uint32
}

func (*GetNextIndexReply) GetMessageName() string {
	return "get_next_index_reply"
}
func (*GetNextIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*GetNextIndexReply) GetCrcString() string {
	return "671fbdb1"
}
func NewGetNextIndexReply() api.Message {
	return &GetNextIndexReply{}
}

// PgCreateInterface represents the VPP binary API message 'pg_create_interface'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 663:
//
//        ["pg_create_interface",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "interface_id"],
//            {"crc" : "0x253c5959"}
//        ],
//
type PgCreateInterface struct {
	InterfaceID uint32
}

func (*PgCreateInterface) GetMessageName() string {
	return "pg_create_interface"
}
func (*PgCreateInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PgCreateInterface) GetCrcString() string {
	return "253c5959"
}
func NewPgCreateInterface() api.Message {
	return &PgCreateInterface{}
}

// PgCreateInterfaceReply represents the VPP binary API message 'pg_create_interface_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 670:
//
//        ["pg_create_interface_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x21b4f949"}
//        ],
//
type PgCreateInterfaceReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*PgCreateInterfaceReply) GetMessageName() string {
	return "pg_create_interface_reply"
}
func (*PgCreateInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PgCreateInterfaceReply) GetCrcString() string {
	return "21b4f949"
}
func NewPgCreateInterfaceReply() api.Message {
	return &PgCreateInterfaceReply{}
}

// PgCapture represents the VPP binary API message 'pg_capture'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 677:
//
//        ["pg_capture",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "interface_id"],
//            ["u8", "is_enabled"],
//            ["u32", "count"],
//            ["u32", "pcap_name_length"],
//            ["u8", "pcap_file_name", 0, "pcap_name_length"],
//            {"crc" : "0x6ac7fe78"}
//        ],
//
type PgCapture struct {
	InterfaceID    uint32
	IsEnabled      uint8
	Count          uint32
	PcapNameLength uint32 `struc:"sizeof=PcapFileName"`
	PcapFileName   []byte
}

func (*PgCapture) GetMessageName() string {
	return "pg_capture"
}
func (*PgCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PgCapture) GetCrcString() string {
	return "6ac7fe78"
}
func NewPgCapture() api.Message {
	return &PgCapture{}
}

// PgCaptureReply represents the VPP binary API message 'pg_capture_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 688:
//
//        ["pg_capture_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xf403693b"}
//        ],
//
type PgCaptureReply struct {
	Retval int32
}

func (*PgCaptureReply) GetMessageName() string {
	return "pg_capture_reply"
}
func (*PgCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PgCaptureReply) GetCrcString() string {
	return "f403693b"
}
func NewPgCaptureReply() api.Message {
	return &PgCaptureReply{}
}

// PgEnableDisable represents the VPP binary API message 'pg_enable_disable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 694:
//
//        ["pg_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_enabled"],
//            ["u32", "stream_name_length"],
//            ["u8", "stream_name", 0, "stream_name_length"],
//            {"crc" : "0x7d0b90ff"}
//        ],
//
type PgEnableDisable struct {
	IsEnabled        uint8
	StreamNameLength uint32 `struc:"sizeof=StreamName"`
	StreamName       []byte
}

func (*PgEnableDisable) GetMessageName() string {
	return "pg_enable_disable"
}
func (*PgEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*PgEnableDisable) GetCrcString() string {
	return "7d0b90ff"
}
func NewPgEnableDisable() api.Message {
	return &PgEnableDisable{}
}

// PgEnableDisableReply represents the VPP binary API message 'pg_enable_disable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 703:
//
//        ["pg_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x02253bd6"}
//        ],
//
type PgEnableDisableReply struct {
	Retval int32
}

func (*PgEnableDisableReply) GetMessageName() string {
	return "pg_enable_disable_reply"
}
func (*PgEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PgEnableDisableReply) GetCrcString() string {
	return "02253bd6"
}
func NewPgEnableDisableReply() api.Message {
	return &PgEnableDisableReply{}
}

// IPSourceAndPortRangeCheckAddDel represents the VPP binary API message 'ip_source_and_port_range_check_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 709:
//
//        ["ip_source_and_port_range_check_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_ipv6"],
//            ["u8", "is_add"],
//            ["u8", "mask_length"],
//            ["u8", "address", 16],
//            ["u8", "number_of_ranges"],
//            ["u16", "low_ports", 32],
//            ["u16", "high_ports", 32],
//            ["u32", "vrf_id"],
//            {"crc" : "0x0f8c6ba0"}
//        ],
//
type IPSourceAndPortRangeCheckAddDel struct {
	IsIpv6         uint8
	IsAdd          uint8
	MaskLength     uint8
	Address        []byte `struc:"[16]byte"`
	NumberOfRanges uint8
	LowPorts       []uint16 `struc:"[32]uint16"`
	HighPorts      []uint16 `struc:"[32]uint16"`
	VrfID          uint32
}

func (*IPSourceAndPortRangeCheckAddDel) GetMessageName() string {
	return "ip_source_and_port_range_check_add_del"
}
func (*IPSourceAndPortRangeCheckAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IPSourceAndPortRangeCheckAddDel) GetCrcString() string {
	return "0f8c6ba0"
}
func NewIPSourceAndPortRangeCheckAddDel() api.Message {
	return &IPSourceAndPortRangeCheckAddDel{}
}

// IPSourceAndPortRangeCheckAddDelReply represents the VPP binary API message 'ip_source_and_port_range_check_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 723:
//
//        ["ip_source_and_port_range_check_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x35df8160"}
//        ],
//
type IPSourceAndPortRangeCheckAddDelReply struct {
	Retval int32
}

func (*IPSourceAndPortRangeCheckAddDelReply) GetMessageName() string {
	return "ip_source_and_port_range_check_add_del_reply"
}
func (*IPSourceAndPortRangeCheckAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IPSourceAndPortRangeCheckAddDelReply) GetCrcString() string {
	return "35df8160"
}
func NewIPSourceAndPortRangeCheckAddDelReply() api.Message {
	return &IPSourceAndPortRangeCheckAddDelReply{}
}

// IPSourceAndPortRangeCheckInterfaceAddDel represents the VPP binary API message 'ip_source_and_port_range_check_interface_add_del'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 729:
//
//        ["ip_source_and_port_range_check_interface_add_del",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u32", "sw_if_index"],
//            ["u32", "tcp_in_vrf_id"],
//            ["u32", "tcp_out_vrf_id"],
//            ["u32", "udp_in_vrf_id"],
//            ["u32", "udp_out_vrf_id"],
//            {"crc" : "0x4a6438f1"}
//        ],
//
type IPSourceAndPortRangeCheckInterfaceAddDel struct {
	IsAdd       uint8
	SwIfIndex   uint32
	TCPInVrfID  uint32
	TCPOutVrfID uint32
	UDPInVrfID  uint32
	UDPOutVrfID uint32
}

func (*IPSourceAndPortRangeCheckInterfaceAddDel) GetMessageName() string {
	return "ip_source_and_port_range_check_interface_add_del"
}
func (*IPSourceAndPortRangeCheckInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*IPSourceAndPortRangeCheckInterfaceAddDel) GetCrcString() string {
	return "4a6438f1"
}
func NewIPSourceAndPortRangeCheckInterfaceAddDel() api.Message {
	return &IPSourceAndPortRangeCheckInterfaceAddDel{}
}

// IPSourceAndPortRangeCheckInterfaceAddDelReply represents the VPP binary API message 'ip_source_and_port_range_check_interface_add_del_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 741:
//
//        ["ip_source_and_port_range_check_interface_add_del_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x6b940f04"}
//        ],
//
type IPSourceAndPortRangeCheckInterfaceAddDelReply struct {
	Retval int32
}

func (*IPSourceAndPortRangeCheckInterfaceAddDelReply) GetMessageName() string {
	return "ip_source_and_port_range_check_interface_add_del_reply"
}
func (*IPSourceAndPortRangeCheckInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*IPSourceAndPortRangeCheckInterfaceAddDelReply) GetCrcString() string {
	return "6b940f04"
}
func NewIPSourceAndPortRangeCheckInterfaceAddDelReply() api.Message {
	return &IPSourceAndPortRangeCheckInterfaceAddDelReply{}
}

// DeleteSubif represents the VPP binary API message 'delete_subif'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 747:
//
//        ["delete_subif",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            {"crc" : "0x6038f848"}
//        ],
//
type DeleteSubif struct {
	SwIfIndex uint32
}

func (*DeleteSubif) GetMessageName() string {
	return "delete_subif"
}
func (*DeleteSubif) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DeleteSubif) GetCrcString() string {
	return "6038f848"
}
func NewDeleteSubif() api.Message {
	return &DeleteSubif{}
}

// DeleteSubifReply represents the VPP binary API message 'delete_subif_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 754:
//
//        ["delete_subif_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0x9d6015dc"}
//        ],
//
type DeleteSubifReply struct {
	Retval int32
}

func (*DeleteSubifReply) GetMessageName() string {
	return "delete_subif_reply"
}
func (*DeleteSubifReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DeleteSubifReply) GetCrcString() string {
	return "9d6015dc"
}
func NewDeleteSubifReply() api.Message {
	return &DeleteSubifReply{}
}

// Punt represents the VPP binary API message 'punt'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 760:
//
//        ["punt",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u8", "is_add"],
//            ["u8", "ipv"],
//            ["u8", "l4_protocol"],
//            ["u16", "l4_port"],
//            {"crc" : "0x4559c976"}
//        ],
//
type Punt struct {
	IsAdd      uint8
	Ipv        uint8
	L4Protocol uint8
	L4Port     uint16
}

func (*Punt) GetMessageName() string {
	return "punt"
}
func (*Punt) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*Punt) GetCrcString() string {
	return "4559c976"
}
func NewPunt() api.Message {
	return &Punt{}
}

// PuntReply represents the VPP binary API message 'punt_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 770:
//
//        ["punt_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xcca27fbe"}
//        ],
//
type PuntReply struct {
	Retval int32
}

func (*PuntReply) GetMessageName() string {
	return "punt_reply"
}
func (*PuntReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*PuntReply) GetCrcString() string {
	return "cca27fbe"
}
func NewPuntReply() api.Message {
	return &PuntReply{}
}

// FeatureEnableDisable represents the VPP binary API message 'feature_enable_disable'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 776:
//
//        ["feature_enable_disable",
//            ["u16", "_vl_msg_id"],
//            ["u32", "client_index"],
//            ["u32", "context"],
//            ["u32", "sw_if_index"],
//            ["u8", "enable"],
//            ["u8", "arc_name", 64],
//            ["u8", "feature_name", 64],
//            {"crc" : "0xbc86393b"}
//        ],
//
type FeatureEnableDisable struct {
	SwIfIndex   uint32
	Enable      uint8
	ArcName     []byte `struc:"[64]byte"`
	FeatureName []byte `struc:"[64]byte"`
}

func (*FeatureEnableDisable) GetMessageName() string {
	return "feature_enable_disable"
}
func (*FeatureEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*FeatureEnableDisable) GetCrcString() string {
	return "bc86393b"
}
func NewFeatureEnableDisable() api.Message {
	return &FeatureEnableDisable{}
}

// FeatureEnableDisableReply represents the VPP binary API message 'feature_enable_disable_reply'.
// Generated from '/usr/share/vpp/api/vpe.api.json', line 786:
//
//        ["feature_enable_disable_reply",
//            ["u16", "_vl_msg_id"],
//            ["u32", "context"],
//            ["i32", "retval"],
//            {"crc" : "0xf6e14373"}
//        ]
//
type FeatureEnableDisableReply struct {
	Retval int32
}

func (*FeatureEnableDisableReply) GetMessageName() string {
	return "feature_enable_disable_reply"
}
func (*FeatureEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*FeatureEnableDisableReply) GetCrcString() string {
	return "f6e14373"
}
func NewFeatureEnableDisableReply() api.Message {
	return &FeatureEnableDisableReply{}
}