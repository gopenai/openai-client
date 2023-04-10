// Code generated by ogen, DO NOT EDIT.

package openaiclient

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CancelFineTune implements cancelFineTune operation.
	//
	// Immediately cancel a fine-tune job.
	//
	// POST /fine-tunes/{fine_tune_id}/cancel
	CancelFineTune(ctx context.Context, params CancelFineTuneParams) (FineTune, error)
	// CreateChatCompletion implements createChatCompletion operation.
	//
	// Creates a completion for the chat message.
	//
	// POST /chat/completions
	CreateChatCompletion(ctx context.Context, req *CreateChatCompletionRequest) (*CreateChatCompletionResponse, error)
	// CreateEdit implements createEdit operation.
	//
	// Creates a new edit for the provided input, instruction, and parameters.
	//
	// POST /edits
	CreateEdit(ctx context.Context, req *CreateEditRequest) (*CreateEditResponse, error)
	// CreateFineTune implements createFineTune operation.
	//
	// Creates a job that fine-tunes a specified model from a given dataset.
	// Response includes details of the enqueued job including job status and the name of the fine-tuned
	// models once complete.
	// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
	//
	// POST /fine-tunes
	CreateFineTune(ctx context.Context, req *CreateFineTuneRequest) (FineTune, error)
	// CreateImage implements createImage operation.
	//
	// Creates an image given a prompt.
	//
	// POST /images/generations
	CreateImage(ctx context.Context, req *CreateImageRequest) (ImagesResponse, error)
	// CreateImageEdit implements createImageEdit operation.
	//
	// Creates an edited or extended image given an original image and a prompt.
	//
	// POST /images/edits
	CreateImageEdit(ctx context.Context, req *CreateImageEditRequestForm) (ImagesResponse, error)
	// CreateImageVariation implements createImageVariation operation.
	//
	// Creates a variation of a given image.
	//
	// POST /images/variations
	CreateImageVariation(ctx context.Context, req *CreateImageVariationRequestForm) (ImagesResponse, error)
	// CreateModeration implements createModeration operation.
	//
	// Classifies if text violates OpenAI's Content Policy.
	//
	// POST /moderations
	CreateModeration(ctx context.Context, req *CreateModerationRequest) (*CreateModerationResponse, error)
	// DeleteFile implements deleteFile operation.
	//
	// Delete a file.
	//
	// DELETE /files/{file_id}
	DeleteFile(ctx context.Context, params DeleteFileParams) (*DeleteFileResponse, error)
	// DeleteModel implements deleteModel operation.
	//
	// Delete a fine-tuned model. You must have the Owner role in your organization.
	//
	// DELETE /models/{model}
	DeleteModel(ctx context.Context, params DeleteModelParams) (*DeleteModelResponse, error)
	// DownloadFile implements downloadFile operation.
	//
	// Returns the contents of the specified file.
	//
	// GET /files/{file_id}/content
	DownloadFile(ctx context.Context, params DownloadFileParams) (string, error)
	// ListFiles implements listFiles operation.
	//
	// Returns a list of files that belong to the user's organization.
	//
	// GET /files
	ListFiles(ctx context.Context) (*ListFilesResponse, error)
	// ListFineTuneEvents implements listFineTuneEvents operation.
	//
	// Get fine-grained status updates for a fine-tune job.
	//
	// GET /fine-tunes/{fine_tune_id}/events
	ListFineTuneEvents(ctx context.Context, params ListFineTuneEventsParams) (*ListFineTuneEventsResponse, error)
	// ListFineTunes implements listFineTunes operation.
	//
	// List your organization's fine-tuning jobs.
	//
	// GET /fine-tunes
	ListFineTunes(ctx context.Context) (*ListFineTunesResponse, error)
	// ListModels implements listModels operation.
	//
	// Lists the currently available models, and provides basic information about each one such as the
	// owner and availability.
	//
	// GET /models
	ListModels(ctx context.Context) (*ListModelsResponse, error)
	// RetrieveFile implements retrieveFile operation.
	//
	// Returns information about a specific file.
	//
	// GET /files/{file_id}
	RetrieveFile(ctx context.Context, params RetrieveFileParams) (OpenAIFile, error)
	// RetrieveFineTune implements retrieveFineTune operation.
	//
	// Gets info about the fine-tune job.
	// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
	//
	// GET /fine-tunes/{fine_tune_id}
	RetrieveFineTune(ctx context.Context, params RetrieveFineTuneParams) (FineTune, error)
	// RetrieveModel implements retrieveModel operation.
	//
	// Retrieves a model instance, providing basic information about the model such as the owner and
	// permissioning.
	//
	// GET /models/{model}
	RetrieveModel(ctx context.Context, params RetrieveModelParams) (Model, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
