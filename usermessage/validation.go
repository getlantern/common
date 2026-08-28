package usermessage

import (
	"fmt"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ValidationError identifies the invalid wire-contract field and why it was
// rejected. Field uses the corresponding JSON field path.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("usermessage: invalid %s: %s", e.Field, e.Reason)
}

// Validate checks whether r is safe and structurally valid for the v1 wire
// contract. Server-side targeting still owns semantic platform and app-version
// validation.
func (r UserMessageRequest) Validate() error {
	if err := validateLocale("locale", r.Locale); err != nil {
		return err
	}
	if err := validateSingleLine("platform", r.Platform, MaxPlatformLength, true); err != nil {
		return err
	}
	if err := validateSingleLine("app_version", r.AppVersion, MaxAppVersionLength, true); err != nil {
		return err
	}
	if err := r.Capabilities.Validate(); err != nil {
		return err
	}
	if len(r.SeenDisplayIDs) > MaxSeenDisplayIDs {
		return invalid("seen_display_ids", fmt.Sprintf("must contain at most %d entries", MaxSeenDisplayIDs))
	}

	seen := make(map[string]struct{}, len(r.SeenDisplayIDs))
	for i, id := range r.SeenDisplayIDs {
		field := fmt.Sprintf("seen_display_ids[%d]", i)
		if err := validateOpaqueID(field, id, MaxDisplayIDLength); err != nil {
			return err
		}
		if _, ok := seen[id]; ok {
			return invalid(field, "must not duplicate another display ID")
		}
		seen[id] = struct{}{}
	}
	return nil
}

// Validate checks whether c describes a supported, internally consistent
// client feature set for this wire version.
func (c ClientCapabilities) Validate() error {
	if c.Version != CapabilityUserMessagesV1 {
		return invalid("capabilities.version", fmt.Sprintf("must be %q", CapabilityUserMessagesV1))
	}
	if len(c.Surfaces) == 0 {
		return invalid("capabilities.surfaces", "must contain at least one surface")
	}
	if len(c.Surfaces) > MaxSupportedSurfaces {
		return invalid(
			"capabilities.surfaces",
			fmt.Sprintf("must contain at most %d entries", MaxSupportedSurfaces),
		)
	}
	seenSurfaces := make(map[Surface]struct{}, len(c.Surfaces))
	for i, surface := range c.Surfaces {
		field := fmt.Sprintf("capabilities.surfaces[%d]", i)
		if !surface.Valid() {
			return invalid(field, fmt.Sprintf("unsupported value %q", surface))
		}
		if _, found := seenSurfaces[surface]; found {
			return invalid(field, "must not duplicate another surface")
		}
		seenSurfaces[surface] = struct{}{}
	}

	if len(c.Actions) > MaxSupportedActions {
		return invalid(
			"capabilities.actions",
			fmt.Sprintf("must contain at most %d entries", MaxSupportedActions),
		)
	}
	seenActions := make(map[ActionType]struct{}, len(c.Actions))
	for i, action := range c.Actions {
		field := fmt.Sprintf("capabilities.actions[%d]", i)
		if !action.Valid() {
			return invalid(field, fmt.Sprintf("unsupported value %q", action))
		}
		if _, found := seenActions[action]; found {
			return invalid(field, "must not duplicate another action")
		}
		seenActions[action] = struct{}{}
	}
	return nil
}

// Validate checks whether r is safe and structurally valid for the v1 wire
// contract.
func (r UserMessageResponse) Validate() error {
	if r.PollIntervalSeconds <= 0 || r.PollIntervalSeconds > MaxPollIntervalSeconds {
		return invalid(
			"poll_interval_seconds",
			fmt.Sprintf("must be between 1 and %d", MaxPollIntervalSeconds),
		)
	}
	if r.Message != nil {
		return r.Message.Validate()
	}
	return nil
}

// Validate checks whether m is presentation-ready and safe for the v1 wire
// contract.
func (m ResolvedUserMessage) Validate() error {
	if err := validateOpaqueID("message.display_id", m.DisplayID, MaxDisplayIDLength); err != nil {
		return err
	}
	if err := validateOpaqueID("message.campaign_id", m.CampaignID, MaxDiagnosticIDLength); err != nil {
		return err
	}
	if err := validateOpaqueID("message.revision_id", m.RevisionID, MaxDiagnosticIDLength); err != nil {
		return err
	}
	if err := validateOpaqueID("message.delivery_id", m.DeliveryID, MaxDiagnosticIDLength); err != nil {
		return err
	}
	if !m.Surface.Valid() {
		return invalid("message.surface", fmt.Sprintf("unsupported value %q", m.Surface))
	}
	if err := validateLocale("message.locale", m.Locale); err != nil {
		return err
	}
	if err := validatePlainText("message.body", m.Body, MaxMessageBodyLength, true, true); err != nil {
		return err
	}
	if m.ExpiresAt.IsZero() {
		return invalid("message.expires_at", "must be set")
	}

	if m.Action == nil {
		if m.ButtonLabel != "" {
			return invalid("message.button_label", "must be empty when action is absent")
		}
		return nil
	}
	if err := validatePlainText("message.button_label", m.ButtonLabel, MaxButtonLabelLength, false, true); err != nil {
		return err
	}
	return m.Action.Validate()
}

// Validate checks whether a contains exactly the data required by its type.
func (a Action) Validate() error {
	if !a.Type.Valid() {
		return invalid("message.action.type", fmt.Sprintf("unsupported value %q", a.Type))
	}
	switch a.Type {
	case ActionTypeOpenHTTPSURL:
		if len(a.URL) == 0 {
			return invalid("message.action.url", "must be set for open_https_url")
		}
		if len(a.URL) > MaxActionURLLength {
			return invalid("message.action.url", fmt.Sprintf("must be at most %d bytes", MaxActionURLLength))
		}
		if !utf8.ValidString(a.URL) || containsDisallowedControl(a.URL, false) {
			return invalid("message.action.url", "must be valid UTF-8 without control characters")
		}
		u, err := url.Parse(a.URL)
		if err != nil || u.Scheme != "https" || u.Hostname() == "" {
			return invalid("message.action.url", "must be an absolute HTTPS URL")
		}
		if u.User != nil {
			return invalid("message.action.url", "must not contain user information")
		}
	case ActionTypeOpenPlans:
		if a.URL != "" {
			return invalid("message.action.url", "must be empty for open_plans")
		}
	}
	return nil
}

// Valid reports whether s is supported by this version of the contract.
func (s Surface) Valid() bool {
	return s == SurfaceSnackbar
}

// Valid reports whether t is supported by this version of the contract.
func (t ActionType) Valid() bool {
	switch t {
	case ActionTypeOpenHTTPSURL, ActionTypeOpenPlans:
		return true
	default:
		return false
	}
}

func validateLocale(field, value string) error {
	if err := validateSingleLine(field, value, MaxLocaleLength, true); err != nil {
		return err
	}
	if strings.HasPrefix(value, "-") || strings.HasSuffix(value, "-") {
		return invalid(field, "must be a BCP 47 language tag")
	}
	parts := strings.Split(value, "-")
	if len(parts) == 1 && (strings.EqualFold(parts[0], "x") || strings.EqualFold(parts[0], "i")) {
		return invalid(field, "must include a subtag after the private-use or grandfathered prefix")
	}
	for i, part := range parts {
		if len(part) == 0 || len(part) > 8 {
			return invalid(field, "must be a BCP 47 language tag")
		}
		for _, r := range part {
			if r > unicode.MaxASCII || !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				return invalid(field, "must contain only BCP 47 subtags")
			}
		}
		if i == 0 {
			if len(part) < 2 && !strings.EqualFold(part, "x") && !strings.EqualFold(part, "i") {
				return invalid(field, "must begin with a language subtag")
			}
			for _, r := range part {
				if !unicode.IsLetter(r) {
					return invalid(field, "must begin with a language subtag")
				}
			}
		}
	}
	return nil
}

func validateOpaqueID(field, value string, max int) error {
	if len(value) == 0 {
		return invalid(field, "must be set")
	}
	if len(value) > max {
		return invalid(field, fmt.Sprintf("must be at most %d bytes", max))
	}
	for _, r := range value {
		if r > unicode.MaxASCII || !isOpaqueIDRune(r) {
			return invalid(field, "must contain only ASCII letters, digits, '.', '_', ':', or '-'")
		}
	}
	return nil
}

func isOpaqueIDRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || strings.ContainsRune("._:-", r)
}

func validateSingleLine(field, value string, max int, required bool) error {
	return validatePlainText(field, value, max, false, required)
}

func validatePlainText(field, value string, max int, allowLineBreaks, required bool) error {
	if required && strings.TrimSpace(value) == "" {
		return invalid(field, "must be set")
	}
	if len(value) > max {
		return invalid(field, fmt.Sprintf("must be at most %d bytes", max))
	}
	if !utf8.ValidString(value) {
		return invalid(field, "must be valid UTF-8")
	}
	if containsDisallowedControl(value, allowLineBreaks) {
		return invalid(field, "must not contain control characters")
	}
	return nil
}

func containsDisallowedControl(value string, allowLineBreaks bool) bool {
	for _, r := range value {
		if !unicode.IsControl(r) {
			continue
		}
		if allowLineBreaks && (r == '\n' || r == '\r' || r == '\t') {
			continue
		}
		return true
	}
	return false
}

func invalid(field, reason string) error {
	return &ValidationError{Field: field, Reason: reason}
}
