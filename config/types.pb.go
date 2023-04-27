// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: config/types.proto

package config

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

type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                              string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr                              string                     `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Cert                              string                     `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	AuthToken                         string                     `protobuf:"bytes,4,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Trusted                           bool                       `protobuf:"varint,5,opt,name=trusted,proto3" json:"trusted,omitempty"`
	MaxPreconnect                     int32                      `protobuf:"varint,6,opt,name=max_preconnect,json=maxPreconnect,proto3" json:"max_preconnect,omitempty"`
	Bias                              int32                      `protobuf:"varint,7,opt,name=bias,proto3" json:"bias,omitempty"`
	PluggableTransport                string                     `protobuf:"bytes,8,opt,name=pluggable_transport,json=pluggableTransport,proto3" json:"pluggable_transport,omitempty"`
	PluggableTransportSettings        map[string]string          `protobuf:"bytes,9,rep,name=pluggable_transport_settings,json=pluggableTransportSettings,proto3" json:"pluggable_transport_settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ENHTTPURL                         string                     `protobuf:"bytes,10,opt,name=ENHTTPURL,proto3" json:"ENHTTPURL,omitempty"`
	TLSDesktopOrderedCipherSuiteNames []string                   `protobuf:"bytes,11,rep,name=TLSDesktop_ordered_cipher_suite_names,json=TLSDesktopOrderedCipherSuiteNames,proto3" json:"TLSDesktop_ordered_cipher_suite_names,omitempty"`
	TLSMobileOrderedCipherSuiteNames  []string                   `protobuf:"bytes,12,rep,name=TLSMobile_ordered_cipher_suite_names,json=TLSMobileOrderedCipherSuiteNames,proto3" json:"TLSMobile_ordered_cipher_suite_names,omitempty"`
	TLSServerNameIndicator            string                     `protobuf:"bytes,13,opt,name=TLSServer_name_indicator,json=TLSServerNameIndicator,proto3" json:"TLSServer_name_indicator,omitempty"`
	TLSClientSessionCacheSize         int32                      `protobuf:"varint,14,opt,name=TLSClient_session_cache_size,json=TLSClientSessionCacheSize,proto3" json:"TLSClient_session_cache_size,omitempty"`
	TLSClientHelloID                  string                     `protobuf:"bytes,15,opt,name=TLSClient_helloID,json=TLSClientHelloID,proto3" json:"TLSClient_helloID,omitempty"`
	TLSClientHelloSplitting           bool                       `protobuf:"varint,16,opt,name=TLSClient_hello_splitting,json=TLSClientHelloSplitting,proto3" json:"TLSClient_hello_splitting,omitempty"`
	TLSClientSessionState             string                     `protobuf:"bytes,17,opt,name=TLSClient_session_state,json=TLSClientSessionState,proto3" json:"TLSClient_session_state,omitempty"`
	Location                          *ProxyConfig_ProxyLocation `protobuf:"bytes,18,opt,name=location,proto3" json:"location,omitempty"`
	MultipathEndpoint                 string                     `protobuf:"bytes,19,opt,name=multipath_endpoint,json=multipathEndpoint,proto3" json:"multipath_endpoint,omitempty"`
	MultiplexedAddr                   string                     `protobuf:"bytes,20,opt,name=multiplexed_addr,json=multiplexedAddr,proto3" json:"multiplexed_addr,omitempty"`
	MultiplexedPhysicalConns          int32                      `protobuf:"varint,21,opt,name=multiplexed_physical_conns,json=multiplexedPhysicalConns,proto3" json:"multiplexed_physical_conns,omitempty"`
	MultiplexedProtocol               string                     `protobuf:"bytes,22,opt,name=multiplexed_protocol,json=multiplexedProtocol,proto3" json:"multiplexed_protocol,omitempty"`
	MultiplexedSettings               map[string]string          `protobuf:"bytes,23,rep,name=multiplexed_settings,json=multiplexedSettings,proto3" json:"multiplexed_settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Track                             string                     `protobuf:"bytes,24,opt,name=track,proto3" json:"track,omitempty"`
	Region                            string                     `protobuf:"bytes,25,opt,name=region,proto3" json:"region,omitempty"`
	AllowedDomains                    []string                   `protobuf:"bytes,26,rep,name=allowed_domains,json=allowedDomains,proto3" json:"allowed_domains,omitempty"`
	StunServers                       []string                   `protobuf:"bytes,27,rep,name=stun_servers,json=stunServers,proto3" json:"stun_servers,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_config_types_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProxyConfig) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ProxyConfig) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *ProxyConfig) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *ProxyConfig) GetTrusted() bool {
	if x != nil {
		return x.Trusted
	}
	return false
}

func (x *ProxyConfig) GetMaxPreconnect() int32 {
	if x != nil {
		return x.MaxPreconnect
	}
	return 0
}

func (x *ProxyConfig) GetBias() int32 {
	if x != nil {
		return x.Bias
	}
	return 0
}

func (x *ProxyConfig) GetPluggableTransport() string {
	if x != nil {
		return x.PluggableTransport
	}
	return ""
}

func (x *ProxyConfig) GetPluggableTransportSettings() map[string]string {
	if x != nil {
		return x.PluggableTransportSettings
	}
	return nil
}

func (x *ProxyConfig) GetENHTTPURL() string {
	if x != nil {
		return x.ENHTTPURL
	}
	return ""
}

func (x *ProxyConfig) GetTLSDesktopOrderedCipherSuiteNames() []string {
	if x != nil {
		return x.TLSDesktopOrderedCipherSuiteNames
	}
	return nil
}

func (x *ProxyConfig) GetTLSMobileOrderedCipherSuiteNames() []string {
	if x != nil {
		return x.TLSMobileOrderedCipherSuiteNames
	}
	return nil
}

func (x *ProxyConfig) GetTLSServerNameIndicator() string {
	if x != nil {
		return x.TLSServerNameIndicator
	}
	return ""
}

func (x *ProxyConfig) GetTLSClientSessionCacheSize() int32 {
	if x != nil {
		return x.TLSClientSessionCacheSize
	}
	return 0
}

func (x *ProxyConfig) GetTLSClientHelloID() string {
	if x != nil {
		return x.TLSClientHelloID
	}
	return ""
}

func (x *ProxyConfig) GetTLSClientHelloSplitting() bool {
	if x != nil {
		return x.TLSClientHelloSplitting
	}
	return false
}

func (x *ProxyConfig) GetTLSClientSessionState() string {
	if x != nil {
		return x.TLSClientSessionState
	}
	return ""
}

func (x *ProxyConfig) GetLocation() *ProxyConfig_ProxyLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ProxyConfig) GetMultipathEndpoint() string {
	if x != nil {
		return x.MultipathEndpoint
	}
	return ""
}

func (x *ProxyConfig) GetMultiplexedAddr() string {
	if x != nil {
		return x.MultiplexedAddr
	}
	return ""
}

func (x *ProxyConfig) GetMultiplexedPhysicalConns() int32 {
	if x != nil {
		return x.MultiplexedPhysicalConns
	}
	return 0
}

func (x *ProxyConfig) GetMultiplexedProtocol() string {
	if x != nil {
		return x.MultiplexedProtocol
	}
	return ""
}

func (x *ProxyConfig) GetMultiplexedSettings() map[string]string {
	if x != nil {
		return x.MultiplexedSettings
	}
	return nil
}

func (x *ProxyConfig) GetTrack() string {
	if x != nil {
		return x.Track
	}
	return ""
}

func (x *ProxyConfig) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ProxyConfig) GetAllowedDomains() []string {
	if x != nil {
		return x.AllowedDomains
	}
	return nil
}

func (x *ProxyConfig) GetStunServers() []string {
	if x != nil {
		return x.StunServers
	}
	return nil
}

type ProxyConfig_ProxyLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City        string  `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country     string  `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	CountryCode string  `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Latitude    float32 `protobuf:"fixed32,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float32 `protobuf:"fixed32,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *ProxyConfig_ProxyLocation) Reset() {
	*x = ProxyConfig_ProxyLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig_ProxyLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig_ProxyLocation) ProtoMessage() {}

func (x *ProxyConfig_ProxyLocation) ProtoReflect() protoreflect.Message {
	mi := &file_config_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig_ProxyLocation.ProtoReflect.Descriptor instead.
func (*ProxyConfig_ProxyLocation) Descriptor() ([]byte, []int) {
	return file_config_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProxyConfig_ProxyLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ProxyConfig_ProxyLocation) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ProxyConfig_ProxyLocation) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *ProxyConfig_ProxyLocation) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ProxyConfig_ProxyLocation) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_config_types_proto protoreflect.FileDescriptor

var file_config_types_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x0c, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x62, 0x69, 0x61, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x6e, 0x0a, 0x1c, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x67, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1a, 0x70, 0x6c, 0x75, 0x67, 0x67,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x4e, 0x48, 0x54, 0x54, 0x50, 0x55,
	0x52, 0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x4e, 0x48, 0x54, 0x54, 0x50,
	0x55, 0x52, 0x4c, 0x12, 0x50, 0x0a, 0x25, 0x54, 0x4c, 0x53, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f,
	0x70, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x21, 0x54, 0x4c, 0x53, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x24, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x20, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x54, 0x4c, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x54, 0x4c, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x3f, 0x0a, 0x1c, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x54, 0x4c, 0x53,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x49, 0x44, 0x12, 0x3a, 0x0a,
	0x19, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x17, 0x54, 0x4c, 0x53,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x54, 0x4c, 0x53, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x74, 0x68,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x3c, 0x0a, 0x1a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x64, 0x5f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x78, 0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e,
	0x73, 0x12, 0x31, 0x0a, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x58, 0x0a, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x78, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x17, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x78, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18,
	0x1a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x75,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x1a, 0x9a, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x1a, 0x4d, 0x0a, 0x1f, 0x50, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x78, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x6c, 0x61,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_types_proto_rawDescOnce sync.Once
	file_config_types_proto_rawDescData = file_config_types_proto_rawDesc
)

func file_config_types_proto_rawDescGZIP() []byte {
	file_config_types_proto_rawDescOnce.Do(func() {
		file_config_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_types_proto_rawDescData)
	})
	return file_config_types_proto_rawDescData
}

var file_config_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_types_proto_goTypes = []interface{}{
	(*ProxyConfig)(nil),               // 0: ProxyConfig
	(*ProxyConfig_ProxyLocation)(nil), // 1: ProxyConfig.ProxyLocation
	nil,                               // 2: ProxyConfig.PluggableTransportSettingsEntry
	nil,                               // 3: ProxyConfig.MultiplexedSettingsEntry
}
var file_config_types_proto_depIdxs = []int32{
	2, // 0: ProxyConfig.pluggable_transport_settings:type_name -> ProxyConfig.PluggableTransportSettingsEntry
	1, // 1: ProxyConfig.location:type_name -> ProxyConfig.ProxyLocation
	3, // 2: ProxyConfig.multiplexed_settings:type_name -> ProxyConfig.MultiplexedSettingsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_types_proto_init() }
func file_config_types_proto_init() {
	if File_config_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
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
		file_config_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig_ProxyLocation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_types_proto_goTypes,
		DependencyIndexes: file_config_types_proto_depIdxs,
		MessageInfos:      file_config_types_proto_msgTypes,
	}.Build()
	File_config_types_proto = out.File
	file_config_types_proto_rawDesc = nil
	file_config_types_proto_goTypes = nil
	file_config_types_proto_depIdxs = nil
}
