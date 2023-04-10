// Code generated by ogen, DO NOT EDIT.

package openaiclient

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CancelFineTune implements cancelFineTune operation.
//
// Immediately cancel a fine-tune job.
//
// POST /fine-tunes/{fine_tune_id}/cancel
func (UnimplementedHandler) CancelFineTune(ctx context.Context, params CancelFineTuneParams) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChatCompletion implements createChatCompletion operation.
//
// Creates a completion for the chat message.
//
// POST /chat/completions
func (UnimplementedHandler) CreateChatCompletion(ctx context.Context, req *CreateChatCompletionRequest) (r *CreateChatCompletionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateEdit implements createEdit operation.
//
// Creates a new edit for the provided input, instruction, and parameters.
//
// POST /edits
func (UnimplementedHandler) CreateEdit(ctx context.Context, req *CreateEditRequest) (r *CreateEditResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateFineTune implements createFineTune operation.
//
// Creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned
// models once complete.
// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
//
// POST /fine-tunes
func (UnimplementedHandler) CreateFineTune(ctx context.Context, req *CreateFineTuneRequest) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImage implements createImage operation.
//
// Creates an image given a prompt.
//
// POST /images/generations
func (UnimplementedHandler) CreateImage(ctx context.Context, req *CreateImageRequest) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImageEdit implements createImageEdit operation.
//
// Creates an edited or extended image given an original image and a prompt.
//
// POST /images/edits
func (UnimplementedHandler) CreateImageEdit(ctx context.Context, req *CreateImageEditRequestForm) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImageVariation implements createImageVariation operation.
//
// Creates a variation of a given image.
//
// POST /images/variations
func (UnimplementedHandler) CreateImageVariation(ctx context.Context, req *CreateImageVariationRequestForm) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateModeration implements createModeration operation.
//
// Classifies if text violates OpenAI's Content Policy.
//
// POST /moderations
func (UnimplementedHandler) CreateModeration(ctx context.Context, req *CreateModerationRequest) (r *CreateModerationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteFile implements deleteFile operation.
//
// Delete a file.
//
// DELETE /files/{file_id}
func (UnimplementedHandler) DeleteFile(ctx context.Context, params DeleteFileParams) (r *DeleteFileResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteModel implements deleteModel operation.
//
// Delete a fine-tuned model. You must have the Owner role in your organization.
//
// DELETE /models/{model}
func (UnimplementedHandler) DeleteModel(ctx context.Context, params DeleteModelParams) (r *DeleteModelResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DownloadFile implements downloadFile operation.
//
// Returns the contents of the specified file.
//
// GET /files/{file_id}/content
func (UnimplementedHandler) DownloadFile(ctx context.Context, params DownloadFileParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFiles implements listFiles operation.
//
// Returns a list of files that belong to the user's organization.
//
// GET /files
func (UnimplementedHandler) ListFiles(ctx context.Context) (r *ListFilesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFineTuneEvents implements listFineTuneEvents operation.
//
// Get fine-grained status updates for a fine-tune job.
//
// GET /fine-tunes/{fine_tune_id}/events
func (UnimplementedHandler) ListFineTuneEvents(ctx context.Context, params ListFineTuneEventsParams) (r *ListFineTuneEventsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFineTunes implements listFineTunes operation.
//
// List your organization's fine-tuning jobs.
//
// GET /fine-tunes
func (UnimplementedHandler) ListFineTunes(ctx context.Context) (r *ListFineTunesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListModels implements listModels operation.
//
// Lists the currently available models, and provides basic information about each one such as the
// owner and availability.
//
// GET /models
func (UnimplementedHandler) ListModels(ctx context.Context) (r *ListModelsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveFile implements retrieveFile operation.
//
// Returns information about a specific file.
//
// GET /files/{file_id}
func (UnimplementedHandler) RetrieveFile(ctx context.Context, params RetrieveFileParams) (r OpenAIFile, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveFineTune implements retrieveFineTune operation.
//
// Gets info about the fine-tune job.
// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
//
// GET /fine-tunes/{fine_tune_id}
func (UnimplementedHandler) RetrieveFineTune(ctx context.Context, params RetrieveFineTuneParams) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveModel implements retrieveModel operation.
//
// Retrieves a model instance, providing basic information about the model such as the owner and
// permissioning.
//
// GET /models/{model}
func (UnimplementedHandler) RetrieveModel(ctx context.Context, params RetrieveModelParams) (r Model, _ error) {
	return r, ht.ErrNotImplemented
}
