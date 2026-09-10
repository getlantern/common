package usermessage

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/getlantern/common/usermessage/testfixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV1JSONFixturesRoundTrip(t *testing.T) {
	t.Run("request", func(t *testing.T) {
		fixture := readFixture(t, testfixtures.V1RequestPath)
		var request UserMessageRequest
		require.NoError(t, json.Unmarshal(fixture, &request))
		require.NoError(t, request.Validate())
		assert.Equal(t, CapabilityUserMessagesV1, request.Capabilities.Version)
		assert.Equal(t, []Surface{SurfaceSnackbar}, request.Capabilities.Surfaces)
		assert.Equal(t, []ActionType{
			ActionTypeOpenHTTPSURL,
			ActionTypeOpenPlans,
		}, request.Capabilities.Actions)
		assert.Equal(t, "fa-IR", request.Locale)
		assert.Equal(t, []string{
			"campaign-018f:generation-1",
			"campaign-0190:generation-3",
		}, request.SeenDisplayIDs)
		assertJSONRoundTrip(t, fixture, request)
	})

	t.Run("response with message", func(t *testing.T) {
		fixture := readFixture(t, testfixtures.V1ResponseMessagePath)
		var response UserMessageResponse
		require.NoError(t, json.Unmarshal(fixture, &response))
		require.NoError(t, response.Validate())
		require.NotNil(t, response.Message)
		assert.Equal(t, SurfaceSnackbar, response.Message.Surface)
		require.NotNil(t, response.Message.Action)
		assert.Equal(t, ActionTypeOpenHTTPSURL, response.Message.Action.Type)
		assertJSONRoundTrip(t, fixture, response)
	})

	t.Run("response without message", func(t *testing.T) {
		fixture := readFixture(t, testfixtures.V1ResponseEmptyPath)
		var response UserMessageResponse
		require.NoError(t, json.Unmarshal(fixture, &response))
		require.NoError(t, response.Validate())
		assert.Nil(t, response.Message)
		assertJSONRoundTrip(t, fixture, response)
	})

	t.Run("message without action", func(t *testing.T) {
		fixture := readFixture(t, testfixtures.V1ResponseNoActionPath)
		var response UserMessageResponse
		require.NoError(t, json.Unmarshal(fixture, &response))
		require.NoError(t, response.Validate())
		require.NotNil(t, response.Message)
		assert.Empty(t, response.Message.ButtonLabel)
		assert.Nil(t, response.Message.Action)
		assertJSONRoundTrip(t, fixture, response)
	})

	t.Run("open plans action", func(t *testing.T) {
		fixture := readFixture(t, testfixtures.V1ResponseOpenPlansPath)
		var response UserMessageResponse
		require.NoError(t, json.Unmarshal(fixture, &response))
		require.NoError(t, response.Validate())
		require.NotNil(t, response.Message)
		require.NotNil(t, response.Message.Action)
		assert.Equal(t, ActionTypeOpenPlans, response.Message.Action.Type)
		assert.Empty(t, response.Message.Action.URL)
		assertJSONRoundTrip(t, fixture, response)
	})
}

func TestRequestValidation(t *testing.T) {
	valid := validRequest()
	require.NoError(t, valid.Validate())

	tests := []struct {
		name   string
		mutate func(*UserMessageRequest)
		field  string
	}{
		{
			name: "missing locale",
			mutate: func(r *UserMessageRequest) {
				r.Locale = ""
			},
			field: "locale",
		},
		{
			name: "non BCP 47 locale",
			mutate: func(r *UserMessageRequest) {
				r.Locale = "fa_IR"
			},
			field: "locale",
		},
		{
			name: "private use locale without subtag",
			mutate: func(r *UserMessageRequest) {
				r.Locale = "x"
			},
			field: "locale",
		},
		{
			name: "extension without value",
			mutate: func(r *UserMessageRequest) {
				r.Locale = "en-a"
			},
			field: "locale",
		},
		{
			name: "numeric extension without value",
			mutate: func(r *UserMessageRequest) {
				r.Locale = "en-1"
			},
			field: "locale",
		},
		{
			name: "missing platform",
			mutate: func(r *UserMessageRequest) {
				r.Platform = ""
			},
			field: "platform",
		},
		{
			name: "missing app version",
			mutate: func(r *UserMessageRequest) {
				r.AppVersion = ""
			},
			field: "app_version",
		},
		{
			name: "wrong capability version",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Version = "user_messages_v2"
			},
			field: "capabilities.version",
		},
		{
			name: "missing supported surfaces",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Surfaces = nil
			},
			field: "capabilities.surfaces",
		},
		{
			name: "unknown supported surface",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Surfaces = []Surface{"future_surface"}
			},
			field: "capabilities.surfaces[0]",
		},
		{
			name: "duplicate supported surface",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Surfaces = []Surface{SurfaceSnackbar, SurfaceSnackbar}
			},
			field: "capabilities.surfaces[1]",
		},
		{
			name: "unknown supported action",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Actions = []ActionType{"future_action"}
			},
			field: "capabilities.actions[0]",
		},
		{
			name: "duplicate supported action",
			mutate: func(r *UserMessageRequest) {
				r.Capabilities.Actions = []ActionType{
					ActionTypeOpenPlans,
					ActionTypeOpenPlans,
				}
			},
			field: "capabilities.actions[1]",
		},
		{
			name: "too many seen IDs",
			mutate: func(r *UserMessageRequest) {
				r.SeenDisplayIDs = make([]string, MaxSeenDisplayIDs+1)
			},
			field: "seen_display_ids",
		},
		{
			name: "duplicate seen ID",
			mutate: func(r *UserMessageRequest) {
				r.SeenDisplayIDs = []string{"display-1", "display-1"}
			},
			field: "seen_display_ids[1]",
		},
		{
			name: "unsafe seen ID",
			mutate: func(r *UserMessageRequest) {
				r.SeenDisplayIDs = []string{"display 1"}
			},
			field: "seen_display_ids[0]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := valid
			request.SeenDisplayIDs = append([]string(nil), valid.SeenDisplayIDs...)
			tt.mutate(&request)
			assertValidationField(t, request.Validate(), tt.field)
		})
	}
}

func TestRequestMaximumLengths(t *testing.T) {
	t.Run("locale", func(t *testing.T) {
		request := validRequest()
		request.Locale = "en-x-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaa"
		require.Len(t, request.Locale, MaxLocaleLength)
		require.NoError(t, request.Validate())

		request.Locale += "a"
		assertValidationField(t, request.Validate(), "locale")
	})

	tests := []struct {
		name      string
		field     string
		max       int
		validChar string
		mutate    func(*UserMessageRequest, string)
	}{
		{
			name:      "platform",
			field:     "platform",
			max:       MaxPlatformLength,
			validChar: "a",
			mutate: func(r *UserMessageRequest, value string) {
				r.Platform = value
			},
		},
		{
			name:      "app version",
			field:     "app_version",
			max:       MaxAppVersionLength,
			validChar: "1",
			mutate: func(r *UserMessageRequest, value string) {
				r.AppVersion = value
			},
		},
		{
			name:      "display ID",
			field:     "seen_display_ids[0]",
			max:       MaxDisplayIDLength,
			validChar: "a",
			mutate: func(r *UserMessageRequest, value string) {
				r.SeenDisplayIDs = []string{value}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := validRequest()
			tt.mutate(&request, strings.Repeat(tt.validChar, tt.max))
			require.NoError(t, request.Validate())

			tt.mutate(&request, strings.Repeat(tt.validChar, tt.max+1))
			assertValidationField(t, request.Validate(), tt.field)
		})
	}
}

func TestResponseValidation(t *testing.T) {
	valid := validResponse()
	require.NoError(t, valid.Validate())
	require.NoError(t, (UserMessageResponse{PollIntervalSeconds: MaxPollIntervalSeconds}).Validate())

	tests := []struct {
		name   string
		mutate func(*UserMessageResponse)
		field  string
	}{
		{
			name: "zero poll interval",
			mutate: func(r *UserMessageResponse) {
				r.PollIntervalSeconds = 0
			},
			field: "poll_interval_seconds",
		},
		{
			name: "poll interval over maximum",
			mutate: func(r *UserMessageResponse) {
				r.PollIntervalSeconds = MaxPollIntervalSeconds + 1
			},
			field: "poll_interval_seconds",
		},
		{
			name: "unknown surface",
			mutate: func(r *UserMessageResponse) {
				r.Message.Surface = Surface("modal")
			},
			field: "message.surface",
		},
		{
			name: "blank body",
			mutate: func(r *UserMessageResponse) {
				r.Message.Body = "  "
			},
			field: "message.body",
		},
		{
			name: "control character in body",
			mutate: func(r *UserMessageResponse) {
				r.Message.Body = "hello\x00world"
			},
			field: "message.body",
		},
		{
			name: "missing expiration",
			mutate: func(r *UserMessageResponse) {
				r.Message.ExpiresAt = time.Time{}
			},
			field: "message.expires_at",
		},
		{
			name: "button without action",
			mutate: func(r *UserMessageResponse) {
				r.Message.Action = nil
			},
			field: "message.button_label",
		},
		{
			name: "action without button",
			mutate: func(r *UserMessageResponse) {
				r.Message.ButtonLabel = ""
			},
			field: "message.button_label",
		},
		{
			name: "unknown action",
			mutate: func(r *UserMessageResponse) {
				r.Message.Action.Type = ActionType("open_deep_link")
			},
			field: "message.action.type",
		},
		{
			name: "non HTTPS URL",
			mutate: func(r *UserMessageResponse) {
				r.Message.Action.URL = "http://example.com"
			},
			field: "message.action.url",
		},
		{
			name: "URL with user information",
			mutate: func(r *UserMessageResponse) {
				r.Message.Action.URL = "https://user:password@example.com"
			},
			field: "message.action.url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := cloneResponse(valid)
			tt.mutate(&response)
			assertValidationField(t, response.Validate(), tt.field)
		})
	}
}

func TestResponseMaximumLengths(t *testing.T) {
	tests := []struct {
		name   string
		field  string
		max    int
		mutate func(*ResolvedUserMessage, string)
	}{
		{
			name:  "display ID",
			field: "message.display_id",
			max:   MaxDisplayIDLength,
			mutate: func(m *ResolvedUserMessage, value string) {
				m.DisplayID = value
			},
		},
		{
			name:  "diagnostic ID",
			field: "message.campaign_id",
			max:   MaxDiagnosticIDLength,
			mutate: func(m *ResolvedUserMessage, value string) {
				m.CampaignID = value
			},
		},
		{
			name:  "message body",
			field: "message.body",
			max:   MaxMessageBodyLength,
			mutate: func(m *ResolvedUserMessage, value string) {
				m.Body = value
			},
		},
		{
			name:  "button label",
			field: "message.button_label",
			max:   MaxButtonLabelLength,
			mutate: func(m *ResolvedUserMessage, value string) {
				m.ButtonLabel = value
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := cloneResponse(validResponse())
			tt.mutate(response.Message, strings.Repeat("a", tt.max))
			require.NoError(t, response.Validate())

			tt.mutate(response.Message, strings.Repeat("a", tt.max+1))
			assertValidationField(t, response.Validate(), tt.field)
		})
	}

	t.Run("action URL", func(t *testing.T) {
		response := cloneResponse(validResponse())
		prefix := "https://example.com/"
		response.Message.Action.URL = prefix + strings.Repeat("a", MaxActionURLLength-len(prefix))
		require.NoError(t, response.Validate())

		response.Message.Action.URL += "a"
		assertValidationField(t, response.Validate(), "message.action.url")
	})
}

func TestActionValidation(t *testing.T) {
	require.NoError(t, (Action{Type: ActionTypeOpenPlans}).Validate())

	err := (Action{Type: ActionTypeOpenPlans, URL: "https://example.com/plans"}).Validate()
	assertValidationField(t, err, "message.action.url")
}

func TestValidationErrorSupportsErrorsAs(t *testing.T) {
	request := validRequest()
	request.Locale = ""
	validationErr := request.Validate()

	var target *ValidationError
	require.True(t, errors.As(validationErr, &target))
	assert.Equal(t, "locale", target.Field)
}

func validRequest() UserMessageRequest {
	return UserMessageRequest{
		Locale:     "en-US",
		Platform:   "android",
		AppVersion: "9.2.1",
		Capabilities: ClientCapabilities{
			Version:  CapabilityUserMessagesV1,
			Surfaces: []Surface{SurfaceSnackbar},
			Actions:  []ActionType{ActionTypeOpenHTTPSURL, ActionTypeOpenPlans},
		},
		SeenDisplayIDs: []string{"campaign-1:generation-1"},
	}
}

func validResponse() UserMessageResponse {
	return UserMessageResponse{
		Message: &ResolvedUserMessage{
			DisplayID:   "campaign-2:generation-1",
			CampaignID:  "campaign-2",
			RevisionID:  "revision-3",
			DeliveryID:  "delivery-4",
			Surface:     SurfaceSnackbar,
			Locale:      "en-US",
			Body:        "Tell us what you think about Lantern.",
			ButtonLabel: "Take survey",
			Action: &Action{
				Type: ActionTypeOpenHTTPSURL,
				URL:  "https://example.com/survey",
			},
			ExpiresAt: time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC),
		},
		PollIntervalSeconds: MaxPollIntervalSeconds,
	}
}

func cloneResponse(response UserMessageResponse) UserMessageResponse {
	cloned := response
	if response.Message != nil {
		message := *response.Message
		cloned.Message = &message
		if response.Message.Action != nil {
			action := *response.Message.Action
			cloned.Message.Action = &action
		}
	}
	return cloned
}

func readFixture(t *testing.T, path string) []byte {
	t.Helper()
	data, err := testfixtures.FS.ReadFile(path)
	require.NoError(t, err)
	return data
}

func assertJSONRoundTrip(t *testing.T, fixture []byte, value any) {
	t.Helper()
	encoded, err := json.Marshal(value)
	require.NoError(t, err)
	assert.JSONEq(t, string(fixture), string(encoded))

	// Standard JSON decoding deliberately ignores unknown fields, allowing a
	// newer producer to add optional data without breaking a v1 consumer.
	withUnknown := strings.TrimSuffix(string(encoded), "}") + `,"future_field":true}`
	decoded := valueForRoundTrip(value)
	require.NoError(t, json.Unmarshal([]byte(withUnknown), decoded))
}

func valueForRoundTrip(value any) any {
	switch value.(type) {
	case UserMessageRequest:
		return &UserMessageRequest{}
	case UserMessageResponse:
		return &UserMessageResponse{}
	default:
		panic("unsupported fixture type")
	}
}

func assertValidationField(t *testing.T, err error, field string) {
	t.Helper()
	require.Error(t, err)
	var validationErr *ValidationError
	require.True(t, errors.As(err, &validationErr))
	assert.Equal(t, field, validationErr.Field)
}
