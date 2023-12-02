// This file was auto-generated by Fern from our API Definition.

package games

import (
	uuid "github.com/google/uuid"
	cloud "github.com/rivet-gg/rivet-go/cloud"
	upload "github.com/rivet-gg/rivet-go/upload"
)

type CreateGameBuildRequest struct {
	// Represent a resource's readable display name.
	DisplayName string `json:"display_name"`
	// A tag given to the game build.
	ImageTag        string              `json:"image_tag"`
	ImageFile       *upload.PrepareFile `json:"image_file,omitempty"`
	MultipartUpload *bool               `json:"multipart_upload,omitempty"`
	Kind            *BuildKind          `json:"kind,omitempty"`
	Compression     *BuildCompression   `json:"compression,omitempty"`
}

type CreateGameBuildResponse struct {
	BuildId                uuid.UUID                  `json:"build_id"`
	UploadId               uuid.UUID                  `json:"upload_id"`
	ImagePresignedRequest  *upload.PresignedRequest   `json:"image_presigned_request,omitempty"`
	ImagePresignedRequests []*upload.PresignedRequest `json:"image_presigned_requests,omitempty"`
}

type ListGameBuildsResponse struct {
	// A list of build summaries.
	Builds []*cloud.BuildSummary `json:"builds,omitempty"`
}
