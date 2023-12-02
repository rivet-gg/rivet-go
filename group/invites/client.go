// This file was auto-generated by Fern from our API Definition.

package invites

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	uuid "github.com/google/uuid"
	rivetgo "github.com/rivet-gg/rivet-go"
	core "github.com/rivet-gg/rivet-go/core"
	group "github.com/rivet-gg/rivet-go/group"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
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
	}
}

// Inspects a group invite returning information about the team that created it.
//
// Provided by `rivet.api.group#CreateInviteResponse$code`.
func (c *Client) GetInvite(ctx context.Context, groupInviteCode string) (*group.GetInviteResponse, error) {
	baseURL := "https://api.rivet.gg"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"group/invites/%v", groupInviteCode)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 500:
			value := new(rivetgo.InternalError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(rivetgo.RateLimitError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(rivetgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 408:
			value := new(rivetgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(rivetgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 400:
			value := new(rivetgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *group.GetInviteResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Consumes a group invite to join a group.
//
// Provided by `rivet.api.group#CreateInviteResponse$code`.
func (c *Client) ConsumeInvite(ctx context.Context, groupInviteCode string) (*group.ConsumeInviteResponse, error) {
	baseURL := "https://api.rivet.gg"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"group/invites/%v/consume", groupInviteCode)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 500:
			value := new(rivetgo.InternalError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(rivetgo.RateLimitError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(rivetgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 408:
			value := new(rivetgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(rivetgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 400:
			value := new(rivetgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *group.ConsumeInviteResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Creates a group invite. Can be shared with other identities to let them join this group.
func (c *Client) CreateInvite(ctx context.Context, groupId uuid.UUID, request *group.CreateInviteRequest) (*group.CreateInviteResponse, error) {
	baseURL := "https://api.rivet.gg"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"group/groups/%v/invites", groupId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 500:
			value := new(rivetgo.InternalError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 429:
			value := new(rivetgo.RateLimitError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(rivetgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 408:
			value := new(rivetgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(rivetgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 400:
			value := new(rivetgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *group.CreateInviteResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
