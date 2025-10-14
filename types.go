package common

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	O "github.com/sagernet/sing-box/option"
)

const SINGBOX = "sing-box"

type UserLevel string

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

	UserLevelPlatinum UserLevel = "platinum"
	UserLevelPro      UserLevel = "pro"
	UserLevelFree     UserLevel = ""

	StatusExpired = "expired"
	StatusActive  = "active"
)

type Duration struct {
	time.Duration
}

func (d Duration) Years() int {
	yearDuration := 365 * 24 * time.Hour
	years := int(d.Duration / yearDuration)
	return years
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	var p DurationParts
	if err := json.Unmarshal(b, &p); err != nil {
		return fmt.Errorf("unmarshaling duration parts: %w", err)
	}

	d.Duration = time.Duration(p.Days)*24*time.Hour + time.Duration(p.Months)*30*24*time.Hour + time.Duration(p.Years)*365*24*time.Hour
	return
}

type DurationParts struct {
	Days   int `json:"days"`
	Months int `json:"months"`
	Years  int `json:"years,omitempty"`
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	yearDuration := 365 * 24 * time.Hour
	years := int(d.Duration / yearDuration)
	monthDuration := 30 * 24 * time.Hour
	months := int((d.Duration - yearDuration*time.Duration(years)) / monthDuration)
	days := int((d.Duration - yearDuration*time.Duration(years) - monthDuration*time.Duration(months)) / (24 * time.Hour))
	dp := DurationParts{
		Days:   days,
		Months: months,
		Years:  years,
	}
	data, err := json.Marshal(dp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal duration plan: %w", err)
	}
	return data, nil
}

type PriceMap map[string]int

func (p PriceMap) String() string {
	var s string
	for k, v := range p.sortedKeys() {
		s += fmt.Sprintf("%v => %v\n", k, v)
	}
	return s
}

func (p PriceMap) sortedKeys() []string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (p PriceMap) FirstValue() int {
	return p[p.sortedKeys()[0]]
}

func (p PriceMap) First() (string, int) {
	key := p.sortedKeys()[0]
	return key, p[key]
}

type Plan struct {
	Id                   string    `json:"id,omitempty"`
	Description          string    `json:"description,omitempty"`
	Duration             Duration  `json:"duration,omitempty"`
	Price                PriceMap  `json:"price,omitempty"`
	ExpectedMonthlyPrice PriceMap  `json:"expectedMonthlyPrice,omitempty"`
	UsdPrice             int       `json:"usdPrice,omitempty"`
	UsdPrice1Y           int       `json:"usdPrice1Y,omitempty"`
	UsdPrice2Y           int       `json:"usdPrice2Y,omitempty"`
	Subscription         bool      `json:"subscription,omitempty"`
	RedeemFor            Duration  `json:"redeemFor,omitempty"`
	RenewalBonus         Duration  `json:"renewalBonus,omitempty"`
	RenewalBonusExpired  Duration  `json:"renewalBonusExpired,omitempty"`
	RenewalBonusExpected Duration  `json:"renewalBonusExpected,omitempty"`
	Discount             float64   `json:"discount"`
	BestValue            bool      `json:"bestValue"`
	Level                UserLevel `json:"level"` // either empty, "pro", or "platinum"
}

func (plan *Plan) DurationMonths() int {
	return int(plan.Duration.Hours() / 24 / 30)
}

func (plan *Plan) DurationSeconds() int {
	return int(plan.Duration.Seconds())
}

type Device struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
}

type PlanDetails struct {
	Provider      string `json:"provider"`
	PurchaseToken string `json:"purchase_token"`
	Time          int64  `json:"time"`
	Plan          string `json:"plan"`
	Expiration    int64  `json:"expiration"`
}

type PlanDefinition struct {
	Level               UserLevel
	Price1Year          []int
	Price2Year          []int
	CurrencyPrice1Year  map[string][]int
	CurrencyPrice2Year  map[string][]int
	Price1Month         []int
	CurrencyPrice1Month map[string][]int
}

type UserData struct {
	UserId           int64               `json:"userId,omitempty"`
	Code             string              `json:"code,omitempty"`
	Token            string              `json:"token,omitempty"`
	Referral         string              `json:"referral,omitempty"`
	Email            string              `json:"email,omitempty"`
	Status           string              `json:"userStatus,omitempty"`
	Level            UserLevel           `json:"userLevel,omitempty"`
	Locale           string              `json:"locale,omitempty"`
	Expiration       int64               `json:"expiration,omitempty"`
	Subscription     string              `json:"subscription,omitempty"`
	Purchases        []map[string]string `json:"purchases,omitempty"`
	BonusDays        string              `json:"bonusDays,omitempty"`
	BonusMonths      string              `json:"bonusMonths,omitempty"`
	Inviters         []string            `json:"inviters"`
	Invitees         []string            `json:"invitees"`
	Devices          []Device            `json:"devices"`
	PurchasedPlans   []PlanDetails       `json:"purchased_plans"`
	SubscriptionData *SubscriptionData   `json:"subscriptionData"`
	IP               string              `json:"ip,omitempty"`
	Country          string              `json:"country,omitempty"`
	EmailConfirmed   bool                `json:"emailConfirmed,omitempty"`
}

type SubscriptionData struct {
	SubscriptionID     string     `json:"subscriptionID"`
	PlanID             string     `json:"planID"`
	StripeCustomerID   string     `json:"stripeCustomerID"`
	Status             string     `json:"status"`
	CancellationReason string     `json:"cancellationReason,omitempty"`
	Provider           string     `json:"provider"`
	CreatedAt          time.Time  `json:"createdAt"`
	StartAt            time.Time  `json:"startAt"`
	EndAt              time.Time  `json:"endAt"`
	CancelledAt        *time.Time `json:"cancelledAt"`
	AutoRenew          bool       `json:"autoRenew"`
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
	TracesSampleRate float64           `json:"sample_rate,omitempty"`
	MetricsInterval  int               `json:"metrics_interval,omitempty"`
}

// Map of outbound tag strings to server locations
type OutboundLocations map[string]*ServerLocation

type ConfigResponse struct {
	UserData          `json:"user_data,omitempty"`
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
