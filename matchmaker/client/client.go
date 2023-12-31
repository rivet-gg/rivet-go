// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/rivet-gg/rivet-go/core"
	lobbies "github.com/rivet-gg/rivet-go/matchmaker/lobbies"
	players "github.com/rivet-gg/rivet-go/matchmaker/players"
	regions "github.com/rivet-gg/rivet-go/matchmaker/regions"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header

	Lobbies *lobbies.Client
	Players *players.Client
	Regions *regions.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
		Lobbies:    lobbies.NewClient(opts...),
		Players:    players.NewClient(opts...),
		Regions:    regions.NewClient(opts...),
	}
}
