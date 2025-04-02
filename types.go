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

type Config struct {
	UserInfo
	AvailableServerLocations
	OutboundLocations
	O.Options
}
