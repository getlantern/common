// Package usermessage defines the public JSON wire contract used to fetch
// resolved in-app messages. Campaign authoring, targeting, and publication
// types intentionally live in lantern-cloud rather than this package.
package usermessage

import "time"

const (
	// CapabilityUserMessagesV1 identifies both support for user messages and
	// version 1 of this wire contract.
	CapabilityUserMessagesV1 = "user_messages_v1"

	// MaxPollIntervalSeconds is the longest interval a server may recommend
	// between successful user-message requests.
	MaxPollIntervalSeconds = 5 * 60

	// Wire-size limits are measured in UTF-8 bytes, except
	// MaxSeenDisplayIDs, which limits the number of list entries.
	MaxLocaleLength       = 64
	MaxPlatformLength     = 32
	MaxAppVersionLength   = 64
	MaxSeenDisplayIDs     = 128
	MaxDisplayIDLength    = 128
	MaxDiagnosticIDLength = 128
	MaxMessageBodyLength  = 2048
	MaxButtonLabelLength  = 128
	MaxActionURLLength    = 2048
)

// Surface identifies how a resolved message should be presented. Consumers
// must ignore messages with surfaces they do not support.
type Surface string

const (
	SurfaceSnackbar Surface = "snackbar"
)

// ActionType identifies an allowlisted operation initiated by a message
// button. The absence of an Action means the message has no button action.
type ActionType string

const (
	ActionTypeOpenHTTPSURL ActionType = "open_https_url"
	ActionTypeOpenPlans    ActionType = "open_plans"
)

// UserMessageRequest contains only the client context needed to resolve a
// message. Authentication and canonical user identity are supplied by the
// transport/account layer, not asserted in this payload.
type UserMessageRequest struct {
	Locale         string   `json:"locale"`
	Platform       string   `json:"platform"`
	AppVersion     string   `json:"app_version"`
	Capability     string   `json:"capability"`
	SeenDisplayIDs []string `json:"seen_display_ids,omitempty"`
}

// UserMessageResponse contains at most one resolved message. A nil Message
// means no message is currently eligible for this client.
type UserMessageResponse struct {
	Message             *ResolvedUserMessage `json:"message,omitempty"`
	PollIntervalSeconds int                  `json:"poll_interval_seconds"`
}

// ResolvedUserMessage is presentation-ready. It deliberately excludes
// campaign targeting and other backoffice-only state.
type ResolvedUserMessage struct {
	DisplayID   string    `json:"display_id"`
	CampaignID  string    `json:"campaign_id"`
	RevisionID  string    `json:"revision_id"`
	DeliveryID  string    `json:"delivery_id"`
	Surface     Surface   `json:"surface"`
	Locale      string    `json:"locale"`
	Body        string    `json:"body"`
	ButtonLabel string    `json:"button_label,omitempty"`
	Action      *Action   `json:"action,omitempty"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// Action contains the data required for an allowlisted client operation.
// URL is required only for ActionTypeOpenHTTPSURL.
type Action struct {
	Type ActionType `json:"type"`
	URL  string     `json:"url,omitempty"`
}
