// This file was auto-generated by Fern from our API Definition.

package games

type CreateCloudTokenResponse struct {
	// A JSON Web Token.
	// Slightly modified to include a description prefix and use Protobufs of
	// JSON.
	Token string `json:"token"`
}