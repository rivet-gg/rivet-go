// This file was auto-generated by Fern from our API Definition.

package client

import (
	adminclient "github.com/rivet-gg/rivet-go/admin/client"
	cloudclient "github.com/rivet-gg/rivet-go/cloud/client"
	core "github.com/rivet-gg/rivet-go/core"
	groupclient "github.com/rivet-gg/rivet-go/group/client"
	identityclient "github.com/rivet-gg/rivet-go/identity/client"
	kvclient "github.com/rivet-gg/rivet-go/kv/client"
	matchmakerclient "github.com/rivet-gg/rivet-go/matchmaker/client"
	moduleclient "github.com/rivet-gg/rivet-go/module/client"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header

	Admin      *adminclient.Client
	Cloud      *cloudclient.Client
	Group      *groupclient.Client
	Identity   *identityclient.Client
	Kv         *kvclient.Client
	Module     *moduleclient.Client
	Matchmaker *matchmakerclient.Client
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
		Admin:      adminclient.NewClient(opts...),
		Cloud:      cloudclient.NewClient(opts...),
		Group:      groupclient.NewClient(opts...),
		Identity:   identityclient.NewClient(opts...),
		Kv:         kvclient.NewClient(opts...),
		Module:     moduleclient.NewClient(opts...),
		Matchmaker: matchmakerclient.NewClient(opts...),
	}
}
