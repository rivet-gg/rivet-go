// This file was auto-generated by Fern from our API Definition.

package namespaces

import (
	cloud "github.com/rivet-gg/rivet-go/cloud"
)

type GetAnalyticsMatchmakerLiveResponse struct {
	// A list of analytics lobby summaries.
	Lobbies []*cloud.LobbySummaryAnalytics `json:"lobbies,omitempty"`
}