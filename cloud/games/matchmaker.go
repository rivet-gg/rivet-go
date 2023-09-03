// This file was auto-generated by Fern from our API Definition.

package games

import (
	fmt "fmt"
	rivetgo "github.com/rivet-gg/rivet-go"
)

type GetLobbyLogsRequest struct {
	Stream LogStream `json:"-"`
	// A query parameter denoting the requests watch index.
	WatchIndex *string `json:"-"`
}

type DeleteMatchmakerLobbyResponse struct {
	// Whether or not the lobby was successfully stopped.
	DidRemove *bool `json:"did_remove,omitempty"`
}

type ExportLobbyLogsRequest struct {
	Stream LogStream `json:"stream,omitempty"`
}

type ExportLobbyLogsResponse struct {
	// The URL to a CSV file for the given lobby history.
	Url string `json:"url"`
}

type ExportMatchmakerLobbyHistoryRequest struct {
	// Unsigned 64 bit integer.
	QueryStart *int64 `json:"query_start,omitempty"`
	// Unsigned 64 bit integer.
	QueryEnd *int64 `json:"query_end,omitempty"`
}

type ExportMatchmakerLobbyHistoryResponse struct {
	// The URL to a CSV file for the given lobby history.
	Url string `json:"url"`
}

type GetLobbyLogsResponse struct {
	// Sorted old to new.
	Lines []string `json:"lines,omitempty"`
	// Sorted old to new.
	Timestamps []string               `json:"timestamps,omitempty"`
	Watch      *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type LogStream string

const (
	LogStreamStdOut LogStream = "std_out"
	LogStreamStdErr LogStream = "std_err"
)

func NewLogStreamFromString(s string) (LogStream, error) {
	switch s {
	case "std_out":
		return LogStreamStdOut, nil
	case "std_err":
		return LogStreamStdErr, nil
	}
	var t LogStream
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LogStream) Ptr() *LogStream {
	return &l
}