package common

import O "github.com/sagernet/sing-box/option"

const SINGBOX = "sing-box"

type UserInfo struct {
	ProToken string `json:"pro_token,omitempty"`
	Country  string `json:"country,omitempty"`
	IP       string `json:"ip,omitempty"`
	ID       string `json:"id,omitempty"`
}

type ServerLocation struct {
	Country     string  `json:"country,omitempty"`
	City        string  `json:"city,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
}

type OTEL struct {
	Endpoint         string            `json:"endpoint,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
	TracesEnabled    bool              `json:"traces_enabled,omitempty"`
	TracesSampleRate float64           `json:"sample_rate,omitempty"`
	MetricsEnabled   bool              `json:"metrics_enabled,omitempty"`
	MetricsInterval  int               `json:"metrics_interval,omitempty"`
}

// Map of outbound tag strings to server locations
type OutboundLocations map[string]*ServerLocation

type ConfigResponse struct {
	Date              string `json:"date,omitempty"`
	UserInfo          `json:"user_info,omitempty"`
	Servers           []ServerLocation `json:"servers,omitempty"`
	OutboundLocations `json:"outbound_locations,omitempty"`
	OTEL              `json:"otel,omitempty"`
	Options           O.Options `json:"options,omitempty"`
}

type ConfigRequest struct {
	DeviceID          string          `json:"device_id,omitempty"`
	SingboxVersion    string          `json:"singbox_version,omitempty"`
	OS                string          `json:"os,omitempty"`
	AppName           string          `json:"app_name,omitempty"`
	PreferredLocation *ServerLocation `json:"preferred_location,omitempty"`
	UserID            string          `json:"user_id,omitempty"`
	ProToken          string          `json:"pro_token,omitempty"`
	WGPublicKey       string          `json:"wg_public_key,omitempty"`
	Backend           string          `json:"backend,omitempty"`
	Locale            string          `json:"locale,omitempty"`
	Protocols         []string        `json:"protocols,omitempty"`
}
