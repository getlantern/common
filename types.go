package config

import O "github.com/sagernet/sing-box/option"

type UserInfo struct {
	ProToken string `json:"pro_token,omitempty"`
	Country  string `json:"country,omitempty"`
	IP       string `json:"ip,omitempty"`
}

type ServerLocation struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

type AvailableServerLocations struct {
	Servers []ServerLocation `json:"servers,omitempty"`
}

// Map of outbound tag strings to server locations
type OutboundLocations map[string]*ServerLocation

type ConfigResponse struct {
	UserInfo
	AvailableServerLocations
	OutboundLocations
	O.Options
}

type ConfigRequest struct {
	ProToken          string         `json:"pro_token,omitempty"`
	ClientVersion     string         `json:"client_version,omitempty"`
	DeviceID          string         `json:"device_id,omitempty"`
	SingboxVersion    string         `json:"singbox_version,omitempty"`
	OS                string         `json:"os,omitempty"`
	AppName           string         `json:"app_name,omitempty"`
	PreferredLocation ServerLocation `json:"preferred_location,omitempty"`
	UserID            string         `json:"user_id,omitempty"`
}
