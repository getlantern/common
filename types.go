package config

import O "github.com/sagernet/sing-box/option"

type UserInfo struct {
	ProToken string `json:"pro_token,omitempty"`
	Country  string `json:"country,omitempty"`
	IP       string `json:"ip,omitempty"`
}

type ServerLocation struct {
	Country     string  `json:"country,omitempty"`
	City        string  `json:"city,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
}

// Map of outbound tag strings to server locations
type OutboundLocations map[string]*ServerLocation

type ConfigResponse struct {
	Date              string `json:"date,omitempty"`
	UserInfo          `json:"user_info,omitempty"`
	Servers           []ServerLocation `json:"servers,omitempty"`
	OutboundLocations `json:"outbound_locations,omitempty"`
	Options           O.Options `json:"options,omitempty"`
}

type ConfigRequest struct {
	ClientVersion     string         `json:"client_version,omitempty"`
	DeviceID          string         `json:"device_id,omitempty"`
	SingboxVersion    string         `json:"singbox_version,omitempty"`
	OS                string         `json:"os,omitempty"`
	AppName           string         `json:"app_name,omitempty"`
	PreferredLocation ServerLocation `json:"preferred_location,omitempty"`
	UserID            string         `json:"user_id,omitempty"`
	WGPublicKey       string         `json:"wg_public_key,omitempty"`
}
