package common

import (
	"log/slog"
	"testing"

	"github.com/sagernet/sing/common/json"
	"github.com/stretchr/testify/assert"
)

func TestConfigRequestSerialization(t *testing.T) {
	original := ConfigRequest{
		DeviceID:       "device123",
		SingboxVersion: "2.0.0",
		Platform:       "linux",
		AppName:        "testApp",
		PreferredLocation: &ServerLocation{
			Country:     "USA",
			City:        "New York",
			Latitude:    40.7128,
			Longitude:   -74.0060,
			CountryCode: "US",
		},
		UserID:      "user123",
		WGPublicKey: "publicKey123",
	}

	// Serialize to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to serialize ConfigRequest: %v", err)
	}

	slog.Info("Serialized ConfigRequest", slog.String("data", string(data)))
	// Deserialize back to struct
	var deserialized ConfigRequest
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
	}
	// Check deep equality of structs.
	assert.ObjectsAreEqual(original, deserialized)
}

func TestConfigRequestDefaultValues(t *testing.T) {
	req := ConfigRequest{}

	if req.DeviceID != "" {
		t.Errorf("Expected default DeviceID to be empty, got: %s", req.DeviceID)
	}
	// PreferredLocation is a pointer, so its zero value is nil — dereferencing it
	// to read .Country panics rather than testing anything.
	if req.PreferredLocation != nil {
		t.Errorf("Expected default PreferredLocation to be nil, got: %+v", req.PreferredLocation)
	}
}

// A client holding no modules must serialize to exactly what it did before the
// field existed. The declaration is an optimization; it must not become a way to
// tell an older client from a newer one, nor a reason to treat them differently.
func TestConfigRequestModulesOmittedWhenEmpty(t *testing.T) {
	// Decode and look for the key rather than substring-searching the JSON: a search for "modules"
	// would also match the word appearing inside some other field's value, so it could pass for the
	// wrong reason. The property under test is "no such key".
	keys := func(t *testing.T, req ConfigRequest) map[string]json.RawMessage {
		t.Helper()
		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("Failed to serialize ConfigRequest: %v", err)
		}
		var out map[string]json.RawMessage
		if err := json.Unmarshal(data, &out); err != nil {
			t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
		}
		return out
	}

	// Both empties, because they are different values that must behave the same. `omitempty` drops a
	// map of length zero, so nil and an initialized-but-empty map are equivalent *today* — asserting
	// only the nil case would let a later change to the tag or the field's type break the other
	// silently, which is exactly the kind of gap this test exists to close.
	for name, modules := range map[string]map[string]uint32{
		"nil":   nil,
		"empty": {},
	} {
		t.Run(name, func(t *testing.T) {
			assert.NotContains(t, keys(t, ConfigRequest{DeviceID: "d", Modules: modules}), "modules",
				"a client holding no modules must send no `modules` key at all")
		})
	}

	held := ConfigRequest{
		DeviceID: "d",
		Modules:  map[string]uint32{"bip324": 3, "obfs-xor": 1},
	}
	assert.JSONEq(t, `{"bip324":3,"obfs-xor":1}`, string(keys(t, held)["modules"]))

	data, err := json.Marshal(held)
	if err != nil {
		t.Fatalf("Failed to serialize ConfigRequest: %v", err)
	}
	var back ConfigRequest
	if err := json.Unmarshal(data, &back); err != nil {
		t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
	}
	assert.Equal(t, map[string]uint32{"bip324": 3, "obfs-xor": 1}, back.Modules)
}

// The capability and the inventory answer different questions, and the server
// needs both: "can this client run a delivered module at all" is not the same as
// "which ones does it already have". A client that supports modules but holds
// none is the case that would otherwise be indistinguishable from one that
// cannot use them, since Modules is omitted when empty.
func TestTransportModulesCapabilityIsSeparateFromTheInventory(t *testing.T) {
	supportsButHoldsNone := ConfigRequest{
		DeviceID:     "d",
		Capabilities: []string{CapabilityTransportModules},
	}
	data, err := json.Marshal(supportsButHoldsNone)
	if err != nil {
		t.Fatalf("Failed to serialize ConfigRequest: %v", err)
	}
	var out map[string]json.RawMessage
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
	}
	assert.Contains(t, out, "capabilities", "the capability is what says the client can run a module")
	assert.NotContains(t, out, "modules", "holding none must still omit the inventory")

	// And a client that cannot use modules sends neither, so the server can tell
	// the two apart — which is the whole point of the capability.
	data, err = json.Marshal(ConfigRequest{DeviceID: "d"})
	if err != nil {
		t.Fatalf("Failed to serialize ConfigRequest: %v", err)
	}
	out = nil
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
	}
	assert.NotContains(t, out, "capabilities")
	assert.NotContains(t, out, "modules")
}

func TestConfigResponseSerialization(t *testing.T) {
	original := ConfigResponse{
		Servers: []ServerLocation{
			{
				Country:     "USA",
				City:        "New York",
				Latitude:    40.7128,
				Longitude:   -74.0060,
				CountryCode: "US",
			},
			{
				Country:     "Canada",
				City:        "Toronto",
				Latitude:    43.65107,
				Longitude:   -79.347015,
				CountryCode: "CA",
			},
		},
		OutboundLocations: OutboundLocations{
			"tag1": {
				Country:     "Germany",
				City:        "Berlin",
				Latitude:    52.52,
				Longitude:   13.405,
				CountryCode: "DE",
			},
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to serialize ConfigResponse: %v", err)
	}

	slog.Info("Serialized ConfigResponse", slog.String("data", string(data)))
	// Deserialize back to struct
	var deserialized ConfigResponse
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize ConfigResponse: %v", err)
	}

	// Compare original and deserialized structs
	if len(original.Servers) != len(deserialized.Servers) ||
		len(original.OutboundLocations) != len(deserialized.OutboundLocations) {
		t.Errorf("Deserialized ConfigResponse does not match original.\nOriginal: %+v\nDeserialized: %+v", original, deserialized)
	}
}

func TestPollIntervalSecondsRoundTrip(t *testing.T) {
	original := ConfigResponse{
		PollIntervalSeconds: 30,
	}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	// Non-zero value should appear in JSON
	assert.Contains(t, string(data), `"poll_interval_seconds":30`)

	var deserialized ConfigResponse
	err = json.Unmarshal(data, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, 30, deserialized.PollIntervalSeconds)

	// Zero value should be omitted from JSON
	zeroResp := ConfigResponse{}
	data, err = json.Marshal(zeroResp)
	assert.NoError(t, err)
	assert.NotContains(t, string(data), "poll_interval_seconds")
}

func TestNonSelectableOutboundsRoundTrip(t *testing.T) {
	original := ConfigResponse{
		NonSelectableOutbounds: []string{"proxyless"},
	}

	data, err := json.Marshal(original)
	assert.NoError(t, err)
	// Verify the JSON tag name is present (not an exact substring, which would be
	// brittle to encoder formatting); the round-trip below verifies the value.
	assert.Contains(t, string(data), `"non_selectable_outbounds"`)

	var deserialized ConfigResponse
	err = json.Unmarshal(data, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, []string{"proxyless"}, deserialized.NonSelectableOutbounds)

	// Both nil and a non-nil empty slice should be omitted from JSON.
	for _, empty := range []ConfigResponse{{}, {NonSelectableOutbounds: []string{}}} {
		data, err = json.Marshal(empty)
		assert.NoError(t, err)
		assert.NotContains(t, string(data), "non_selectable_outbounds")
	}
}

func TestCapabilitiesRoundTrip(t *testing.T) {
	original := ConfigRequest{
		Capabilities: []string{CapabilityNonSelectableOutbounds},
	}

	data, err := json.Marshal(original)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"capabilities"`)

	var deserialized ConfigRequest
	err = json.Unmarshal(data, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, []string{CapabilityNonSelectableOutbounds}, deserialized.Capabilities)

	// Both nil and a non-nil empty slice should be omitted from JSON.
	for _, empty := range []ConfigRequest{{}, {Capabilities: []string{}}} {
		data, err = json.Marshal(empty)
		assert.NoError(t, err)
		assert.NotContains(t, string(data), `"capabilities"`)
	}
}

func TestConfigResponseDefaultValues(t *testing.T) {
	resp := ConfigResponse{}
	if len(resp.Servers) != 0 {
		t.Errorf("Expected default Servers to be empty, got: %+v", resp.Servers)
	}
	if len(resp.OutboundLocations) != 0 {
		t.Errorf("Expected default OutboundLocations to be empty, got: %+v", resp.OutboundLocations)
	}
}

func TestDonorSTUNConfigRoundTrip(t *testing.T) {
	original := UnboundedConfig{STUNServers: DefaultDonorSTUNServers()}
	encoded, err := json.Marshal(original)
	assert.NoError(t, err)
	var wire map[string]json.RawMessage
	assert.NoError(t, json.Unmarshal(encoded, &wire))
	assert.Contains(t, wire, "stun_servers")
	var decoded UnboundedConfig
	assert.NoError(t, json.Unmarshal(encoded, &decoded))
	assert.Equal(t, original.STUNServers, decoded.STUNServers)
	original.STUNServers[0] = "changed"
	assert.NotContains(t, DefaultDonorSTUNServers(), "changed")
}

func TestDonorSTUNConfigOmission(t *testing.T) {
	for _, servers := range [][]string{nil, {}} {
		encoded, err := json.Marshal(UnboundedConfig{STUNServers: servers})
		assert.NoError(t, err)
		assert.JSONEq(t, "{}", string(encoded))
	}
}
