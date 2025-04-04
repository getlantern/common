package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigRequestSerialization(t *testing.T) {
	original := ConfigRequest{
		ClientVersion:  "1.0.0",
		DeviceID:       "device123",
		SingboxVersion: "2.0.0",
		OS:             "linux",
		AppName:        "testApp",
		PreferredLocation: ServerLocation{
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

	// Deserialize back to struct
	var deserialized ConfigRequest
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize ConfigRequest: %v", err)
	}

	// Compare original and deserialized structs
	if original != deserialized {
		t.Errorf("Deserialized ConfigRequest does not match original.\nOriginal: %+v\nDeserialized: %+v", original, deserialized)
	}
}

func TestConfigRequestDefaultValues(t *testing.T) {
	req := ConfigRequest{}

	if req.ClientVersion != "" {
		t.Errorf("Expected default ClientVersion to be empty, got: %s", req.ClientVersion)
	}
	if req.DeviceID != "" {
		t.Errorf("Expected default DeviceID to be empty, got: %s", req.DeviceID)
	}
	if req.PreferredLocation.Country != "" {
		t.Errorf("Expected default PreferredLocation.Country to be empty, got: %s", req.PreferredLocation.Country)
	}
}
func TestConfigResponseSerialization(t *testing.T) {
	original := ConfigResponse{
		UserInfo: UserInfo{
			ProToken: "token123",
			Country:  "USA",
			IP:       "192.168.1.1",
		},
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
			"outbound1": &ServerLocation{
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

	// Deserialize back to struct
	var deserialized ConfigResponse
	err = json.Unmarshal(data, &deserialized)
	if err != nil {
		t.Fatalf("Failed to deserialize ConfigResponse: %v", err)
	}

	// Compare original and deserialized structs
	if original.UserInfo != deserialized.UserInfo {
		t.Errorf("Deserialized UserInfo does not match original.\nOriginal: %+v\nDeserialized: %+v", original.UserInfo, deserialized.UserInfo)
	}

	if len(original.Servers) != len(deserialized.Servers) {
		t.Fatalf("Expected %d servers, got %d", len(original.Servers), len(deserialized.Servers))
	}
	for i, server := range original.Servers {
		if server != deserialized.Servers[i] {
			t.Errorf("Server at index %d does not match.\nOriginal: %+v\nDeserialized: %+v", i, server, deserialized.Servers[i])
		}
	}

	if len(original.OutboundLocations) != len(deserialized.OutboundLocations) {
		t.Fatalf("Expected %d outbound locations, got %d", len(original.OutboundLocations), len(deserialized.OutboundLocations))
	}
	for key, location := range original.OutboundLocations {
		if deserialized.OutboundLocations[key] == nil || *location != *deserialized.OutboundLocations[key] {
			t.Errorf("Outbound location for key %s does not match.\nOriginal: %+v\nDeserialized: %+v", key, location, deserialized.OutboundLocations[key])
		}
	}

	data = []byte("{\"pro_token\":\"mock-token\",\"OutboundLocations\":null}")
	// Deserialize back to struct
	var cr ConfigResponse
	err = json.Unmarshal(data, &cr)
	if err != nil {
		t.Fatalf("Failed to deserialize ConfigResponse: %v", err)
	}

	assert.Equal(t, cr.UserInfo.ProToken, "mock-token")
}
