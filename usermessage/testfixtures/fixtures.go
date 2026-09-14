// Package testfixtures exposes versioned user-message JSON fixtures for
// compatibility tests in common, lantern-cloud, and Radiance.
package testfixtures

import "embed"

const (
	V1RequestPath           = "v1/request.json"
	V1ResponseEmptyPath     = "v1/response_empty.json"
	V1ResponseMessagePath   = "v1/response_message.json"
	V1ResponseNoActionPath  = "v1/response_no_action.json"
	V1ResponseOpenPlansPath = "v1/response_open_plans.json"
)

// FS contains the versioned JSON fixtures named by the exported path
// constants above.
//
//go:embed v1/*.json
var FS embed.FS
