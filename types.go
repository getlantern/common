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

	// Whether or not client-side logs should be enabled.
	LOGS = "otel.logs"

	// Whether or not users should have the option to launch private servers on GCP.
	GCP = "private.gcp"

	// Whether or not the client should run the unbounded widget proxy.
	UNBOUNDED = "unbounded"
)

// Client capabilities, advertised in ConfigRequest.Capabilities, let the server
// enable optional behavior per-client — capability negotiation instead of
// version sniffing.
const (
	// CapabilityNonSelectableOutbounds: the client honors
	// ConfigResponse.NonSelectableOutbounds (merges those outbounds into its box
	// config but keeps them out of the proxy-selection groups). The server gates
	// infrastructure outbounds like the proxyless download_detour on this, so an
	// older client can't surface one as a selectable proxy and route traffic
	// through it.
	CapabilityNonSelectableOutbounds = "non_selectable_outbounds"

	// CapabilityTransportModules: the client can install and run a signed
	// transport-module bundle delivered in its config, and verifies it against a
	// compiled-in key.
	//
	// Needed because ConfigRequest.Modules cannot carry this by itself: it is
	// omitted when empty, so a client that supports modules but holds none looks
	// on the wire exactly like a client that cannot use them at all. Without the
	// capability the server would have to guess, and guessing wrong means either
	// never bootstrapping a client's first module or sending every client a
	// module-bearing outbound it will silently skip — with the module's bytes,
	// which are the expensive part, attached.
	CapabilityTransportModules = "transport_modules"
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

type SmartRoutingRules []SmartRoutingRule

type AdBlockRules RuleSets

type RuleSets []RuleSet

type SmartRoutingRule struct {
	Category  string   `json:"category,omitempty"`
	RuleSets  RuleSets `json:"rule_sets,omitempty"`
	Outbounds []string `json:"outbounds,omitempty"`
}

type RuleSet struct {
	Tag string `json:"tag,omitempty"`
	URL string `json:"url,omitempty"`
	// ruleset format: sing-box/constant.RuleSetFormatBinary (SRS) or sing-box/constant.RuleSetFormatSource (JSON)
	// Defaults to SRS if omitted
	Format string `json:"format,omitempty"`
	// outbound to use for downloading the ruleset. If omitted, uses the "direct" outbound.
	DownloadDetour string `json:"download_detour,omitempty"`
}

type UnboundedConfig struct {
	DiscoverySrv      string   `json:"discovery_srv,omitempty"`
	DiscoveryEndpoint string   `json:"discovery_endpoint,omitempty"`
	EgressAddr        string   `json:"egress_addr,omitempty"`
	EgressEndpoint    string   `json:"egress_endpoint,omitempty"`
	CTableSize        int      `json:"ctable_size,omitempty"`
	PTableSize        int      `json:"ptable_size,omitempty"`
	STUNServers       []string `json:"stun_servers,omitempty"`
}

// DefaultDonorSTUNServers returns a fresh fallback pool that requires no DNS lookup.
func DefaultDonorSTUNServers() []string {
	return []string{
		"stun:5.39.72.109:3478",
		"stun:176.9.24.184:3478",
		"stun:20.93.239.169:3478",
		"stun:46.225.95.169:3478",
		"stun:136.243.59.79:3478",
		"stun:199.4.110.11:3478",
		"stun:203.56.114.226:3478",
		"stun:35.158.233.7:3478",
	}
}

type ConfigResponse struct {
	Country           string           `json:"country,omitempty"`
	IP                string           `json:"ip,omitempty"`
	Servers           []ServerLocation `json:"servers,omitempty"`
	OutboundLocations `json:"outbound_locations,omitempty"`
	OTEL              `json:"otel,omitempty"`
	Features          map[string]bool   `json:"features,omitempty"`
	Options           O.Options         `json:"options,omitempty"`
	SmartRouting      SmartRoutingRules `json:"smart_routing,omitempty"`
	AdBlock           AdBlockRules      `json:"ad_block,omitempty"`

	// NonSelectableOutbounds lists the tags of outbounds in Options that are
	// infrastructure (e.g. a proxyless download_detour for rule-set fetches). The
	// client merges them into its box config so references (download_detour, route
	// rules) resolve, but keeps them out of BOTH proxy-selection groups:
	//   - they are NOT added to the auto (URLTest) group, so auto-selection never
	//     routes user traffic through them, and
	//   - they are NOT offered in the manual selector, so the user can't pick them.
	// This lets the server introduce new optional/infra outbounds without a client
	// release — the client honors whatever tags appear here.
	NonSelectableOutbounds []string `json:"non_selectable_outbounds,omitempty"`

	// PollIntervalSeconds tells the client how long to wait before fetching
	// a new config. The server adjusts this based on bandit confidence —
	// shorter intervals when learning (new ASN, high entropy), longer when
	// the assignment is stable. Zero means use the client's default.
	PollIntervalSeconds int `json:"poll_interval_seconds,omitempty"`

	// BanditURLOverrides maps outbound tags to per-proxy callback URLs for
	// the bandit Thompson sampling system. When set, these override the
	// default MutableURLTest URL for each specific outbound, allowing the
	// server to detect which proxies successfully connected.
	BanditURLOverrides  map[string]string `json:"bandit_url_overrides,omitempty"`
	BanditThroughputURL string            `json:"bandit_throughput_url,omitempty"`

	// RouteSelectionReportIntervalSeconds tells the client how often to POST
	// its per-route selection history (data-plane stall/reset counters) back
	// to the server's /v1/bandit/report endpoint. Omitted or zero means the
	// client does not report — the interval doubles as the enable switch.
	RouteSelectionReportIntervalSeconds int `json:"route_selection_report_interval_seconds,omitempty"`

	// BanditReportTokens maps outbound tags to HMAC-signed opaque report
	// tokens, parallel to BanditURLOverrides. The client echoes each tag's
	// token back in its selection-history report so the server can attribute
	// the counters to a route without trusting any client-asserted identity.
	BanditReportTokens map[string]string `json:"bandit_report_tokens,omitempty"`

	Unbounded *UnboundedConfig `json:"unbounded,omitempty"`
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
	// Capabilities advertises optional client behaviors the server can gate on
	// (see the Capability* consts), e.g. honoring NonSelectableOutbounds.
	Capabilities   []string `json:"capabilities,omitempty"`
	MetricsOptedIn bool     `json:"metrics_opted_in,omitempty"`
	Version        string   `json:"version,omitempty"`

	// Modules names the signed transport-module bundles the client already
	// holds, as engine name -> bundle version, so the server can omit bytes it
	// would otherwise re-send.
	//
	// Distinct from Capabilities, which is a set of boolean tokens: this is an
	// inventory, and its values are what the server compares against. The two
	// are complementary — CapabilityTransportModules says the client can load a
	// delivered module at all, and this says which ones it already has. A client
	// that supports modules but holds none sends the capability and no Modules,
	// which is exactly the case the server must be able to tell apart from a
	// client that cannot use them.
	//
	// A transport module can be delivered inline in the config itself, which
	// means it rides every fetch that offers it. The response ETag does not
	// help: the body is regenerated per request (the bandit re-picks routes),
	// so it never repeats and never yields a 304. Without this field an inline
	// module is re-sent on every poll, indefinitely.
	//
	// The version rather than a content hash: the client's bundle store already
	// persists exactly this, and the artifact's own Ed25519 signature is what
	// authenticates its bytes, so a hash would be a second identity for one
	// thing.
	//
	// This is a hint and never authorization. A client claiming an engine it
	// cannot actually load simply skips that outbound, so a wrong or dishonest
	// declaration only degrades that client. Nothing here may gate access, and
	// omitted bytes must be the only difference it makes.
	//
	// Omitted when empty, so a client holding nothing — or one built without
	// module support — sends exactly what it always sent.
	Modules map[string]uint32 `json:"modules,omitempty"`
}
