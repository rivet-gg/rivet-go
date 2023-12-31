// This file was auto-generated by Fern from our API Definition.

package games

import (
	uuid "github.com/google/uuid"
	rivetgo "github.com/rivet-gg/rivet-go"
	cloud "github.com/rivet-gg/rivet-go/cloud"
	game "github.com/rivet-gg/rivet-go/game"
	group "github.com/rivet-gg/rivet-go/group"
	upload "github.com/rivet-gg/rivet-go/upload"
)

type GetGameByIdRequest struct {
	// A query parameter denoting the requests watch index.
	WatchIndex *string `json:"-"`
}

type GetGamesRequest struct {
	// A query parameter denoting the requests watch index.
	WatchIndex *string `json:"-"`
}

type CreateGameRequest struct {
	// A human readable short identifier used to references resources. Different than a `rivet.common#Uuid` because this is intended to be human readable. Different than `rivet.common#DisplayName` because this should not include special characters and be short.
	NameId string `json:"name_id"`
	// Represent a resource's readable display name.
	DisplayName      string    `json:"display_name"`
	DeveloperGroupId uuid.UUID `json:"developer_group_id"`
}

type CreateGameResponse struct {
	GameId uuid.UUID `json:"game_id"`
}

type GameBannerUploadPrepareRequest struct {
	// The path/filename of the game banner.
	Path string `json:"path"`
	// The MIME type of the game banner.
	Mime *string `json:"mime,omitempty"`
	// Unsigned 64 bit integer.
	ContentLength int64 `json:"content_length"`
}

type GameBannerUploadPrepareResponse struct {
	UploadId         uuid.UUID                `json:"upload_id"`
	PresignedRequest *upload.PresignedRequest `json:"presigned_request,omitempty"`
}

type GameLogoUploadPrepareRequest struct {
	// The path/filename of the game logo.
	Path string `json:"path"`
	// The MIME type of the game logo.
	Mime *string `json:"mime,omitempty"`
	// Unsigned 64 bit integer.
	ContentLength int64 `json:"content_length"`
}

type GameLogoUploadPrepareResponse struct {
	UploadId         uuid.UUID                `json:"upload_id"`
	PresignedRequest *upload.PresignedRequest `json:"presigned_request,omitempty"`
}

type GetGameByIdResponse struct {
	Game  *cloud.GameFull        `json:"game,omitempty"`
	Watch *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type GetGamesResponse struct {
	// A list of game summaries.
	Games []*game.Summary `json:"games,omitempty"`
	// A list of group summaries.
	Groups []*group.Summary       `json:"groups,omitempty"`
	Watch  *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type ValidateGameRequest struct {
	// Represent a resource's readable display name.
	DisplayName string `json:"display_name"`
	// A human readable short identifier used to references resources. Different than a `rivet.common#Uuid` because this is intended to be human readable. Different than `rivet.common#DisplayName` because this should not include special characters and be short.
	NameId string `json:"name_id"`
}

type ValidateGameResponse struct {
	// A list of validation errors.
	Errors []*rivetgo.ValidationError `json:"errors,omitempty"`
}
