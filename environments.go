// This file was auto-generated by Fern from our API Definition.

package api

// Environments defines all of the API environments.
// These values can be used with the ClientWithBaseURL
// ClientOption to override the client's default environment,
// if any.
var Environments = struct {
	Production struct {
		Admin      string
		Auth       string
		Chat       string
		Cloud      string
		Group      string
		Identity   string
		Job        string
		Kv         string
		Matchmaker string
		Module     string
		Portal     string
	}
}{
	Production: struct {
		Admin      string
		Auth       string
		Chat       string
		Cloud      string
		Group      string
		Identity   string
		Job        string
		Kv         string
		Matchmaker string
		Module     string
		Portal     string
	}{
		Admin:      "https://admin.api.rivet.gg/v1",
		Auth:       "https://auth.api.rivet.gg/v1",
		Chat:       "https://chat.api.rivet.gg/v1",
		Cloud:      "https://cloud.api.rivet.gg/v1",
		Group:      "https://group.api.rivet.gg/v1",
		Identity:   "https://identity.api.rivet.gg/v1",
		Job:        "https://job.api.rivet.gg/v1",
		Kv:         "https://kv.api.rivet.gg/v1",
		Matchmaker: "https://matchmaker.api.rivet.gg/v1",
		Module:     "https://module.api.rivet.gg/v1",
		Portal:     "https://portal.api.rivet.gg/v1",
	},
}