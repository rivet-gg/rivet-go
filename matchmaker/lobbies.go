// This file was auto-generated by Fern from our API Definition.

package matchmaker

import (
	fmt "fmt"
	captcha "github.com/rivet-gg/rivet-go/captcha"
	core "github.com/rivet-gg/rivet-go/core"
)

type CreateLobbyRequest struct {
	GameMode         string                         `json:"game_mode"`
	Region           *core.Optional[string]         `json:"region,omitempty"`
	Captcha          *core.Optional[captcha.Config] `json:"captcha,omitempty"`
	Publicity        CustomLobbyPublicity           `json:"publicity,omitempty"`
	LobbyConfig      *core.Optional[any]            `json:"lobby_config,omitempty"`
	VerificationData *core.Optional[any]            `json:"verification_data,omitempty"`
}

type FindLobbyRequest struct {
	Origin                 *string                        `json:"-"`
	GameModes              []string                       `json:"game_modes,omitempty"`
	Regions                []string                       `json:"regions,omitempty"`
	PreventAutoCreateLobby *core.Optional[bool]           `json:"prevent_auto_create_lobby,omitempty"`
	Captcha                *core.Optional[captcha.Config] `json:"captcha,omitempty"`
	VerificationData       *core.Optional[any]            `json:"verification_data,omitempty"`
}

type JoinLobbyRequest struct {
	LobbyId          string                         `json:"lobby_id"`
	Captcha          *core.Optional[captcha.Config] `json:"captcha,omitempty"`
	VerificationData *core.Optional[any]            `json:"verification_data,omitempty"`
}

type ListLobbiesRequest struct {
	IncludeState *bool `json:"-"`
}

type SetLobbyClosedRequest struct {
	IsClosed bool `json:"is_closed"`
}

// Methods to verify a captcha
type Config struct {
	Hcaptcha  *ConfigHcaptcha  `json:"hcaptcha,omitempty"`
	Turnstile *ConfigTurnstile `json:"turnstile,omitempty"`
}

type CustomLobbyPublicity string

const (
	CustomLobbyPublicityPublic  CustomLobbyPublicity = "public"
	CustomLobbyPublicityPrivate CustomLobbyPublicity = "private"
)

func NewCustomLobbyPublicityFromString(s string) (CustomLobbyPublicity, error) {
	switch s {
	case "public":
		return CustomLobbyPublicityPublic, nil
	case "private":
		return CustomLobbyPublicityPrivate, nil
	}
	var t CustomLobbyPublicity
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CustomLobbyPublicity) Ptr() *CustomLobbyPublicity {
	return &c
}

type CreateLobbyResponse struct {
	Lobby  *JoinLobby           `json:"lobby,omitempty"`
	Ports  map[string]*JoinPort `json:"ports,omitempty"`
	Player *JoinPlayer          `json:"player,omitempty"`
}

type FindLobbyResponse struct {
	Lobby  *JoinLobby           `json:"lobby,omitempty"`
	Ports  map[string]*JoinPort `json:"ports,omitempty"`
	Player *JoinPlayer          `json:"player,omitempty"`
}

type JoinLobbyResponse struct {
	Lobby  *JoinLobby           `json:"lobby,omitempty"`
	Ports  map[string]*JoinPort `json:"ports,omitempty"`
	Player *JoinPlayer          `json:"player,omitempty"`
}

type ListLobbiesResponse struct {
	GameModes []*GameModeInfo `json:"game_modes,omitempty"`
	Regions   []*RegionInfo   `json:"regions,omitempty"`
	Lobbies   []*LobbyInfo    `json:"lobbies,omitempty"`
}