package common

import (
	"log/slog"
	"testing"

	"github.com/sagernet/sing/common/json"
)

func TestConfigRequestSerialization(t *testing.T) {
	original := ConfigRequest{
		DeviceID:       "device123",
		SingboxVersion: "2.0.0",
		OS:             "linux",
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

	// Compare original and deserialized structs
	if original != deserialized {
		t.Errorf("Deserialized ConfigRequest does not match original.\nOriginal: %+v\nDeserialized: %+v", original, deserialized)
	}
}

func TestConfigRequestDefaultValues(t *testing.T) {
	req := ConfigRequest{}

	if req.DeviceID != "" {
		t.Errorf("Expected default DeviceID to be empty, got: %s", req.DeviceID)
	}
	if req.PreferredLocation.Country != "" {
		t.Errorf("Expected default PreferredLocation.Country to be empty, got: %s", req.PreferredLocation.Country)
	}
}
func TestConfigResponseSerialization(t *testing.T) {
	original := ConfigResponse{
		Date: "2023-01-01",
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
	if original.Date != deserialized.Date ||
		original.UserInfo != deserialized.UserInfo ||
		len(original.Servers) != len(deserialized.Servers) ||
		len(original.OutboundLocations) != len(deserialized.OutboundLocations) {
		t.Errorf("Deserialized ConfigResponse does not match original.\nOriginal: %+v\nDeserialized: %+v", original, deserialized)
	}
}

func TestConfigResponseDefaultValues(t *testing.T) {
	resp := ConfigResponse{}

	if resp.Date != "" {
		t.Errorf("Expected default Date to be empty, got: %s", resp.Date)
	}
	if resp.UserInfo.ProToken != "" {
		t.Errorf("Expected default UserInfo.ProToken to be empty, got: %s", resp.UserInfo.ProToken)
	}
	if len(resp.Servers) != 0 {
		t.Errorf("Expected default Servers to be empty, got: %+v", resp.Servers)
	}
	if len(resp.OutboundLocations) != 0 {
		t.Errorf("Expected default OutboundLocations to be empty, got: %+v", resp.OutboundLocations)
	}
}
