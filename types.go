package common

import O "github.com/sagernet/sing-box/option"

const SINGBOX = "sing-box"

// Feature flags for Lantern clients. These should correspond directly with
// the feature names on the Unleash server at https://unleash.lantr.net/projects/default/features
// These are also exposed to the frontend/Flutter side via the same keys.
const (
	// Whether or not client-side traces should be enabled.
	TRACES = "otel.traces"

	// Whether or not client-side metrics should be enabled.
	METRICS = "otel.metrics"

	// Whether or not users should have the option to launch private servers on GCP.
	GCP = "private.gcp"
)

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
	TracesSampleRate float64           `json:"sample_rate,omitempty"`
	MetricsInterval  int               `json:"metrics_interval,omitempty"`
}

// Map of outbound tag strings to server locations
type OutboundLocations map[string]*ServerLocation

type ConfigResponse struct {
	Country           string           `json:"country,omitempty"`
	IP                string           `json:"ip,omitempty"`
	Pro               bool             `json:"pro,omitempty"`
	Servers           []ServerLocation `json:"servers,omitempty"`
	OutboundLocations `json:"outbound_locations,omitempty"`
	OTEL              `json:"otel,omitempty"`
	Features          map[string]bool `json:"features,omitempty"`
	Options           O.Options       `json:"options,omitempty"`
}

type ConfigRequest struct {
	DeviceID          string          `json:"device_id,omitempty"`
	SingboxVersion    string          `json:"singbox_version,omitempty"`
	Platform          string          `json:"platform,omitempty"`
	AppName           string          `json:"app_name,omitempty"`
	PreferredLocation *ServerLocation `json:"preferred_location,omitempty"`
	UserID            string          `json:"user_id,omitempty"`
	ProToken          string          `json:"pro_token,omitempty"`
	WGPublicKey       string          `json:"wg_public_key,omitempty"`
	Backend           string          `json:"backend,omitempty"`
	Locale            string          `json:"locale,omitempty"`
	Protocols         []string        `json:"protocols,omitempty"`
}
