// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browserbaseunofficial

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apijson"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/requestconfig"
	"github.com/stainless-sdks/browserbase-unofficial-go/option"
	"github.com/stainless-sdks/browserbase-unofficial-go/packages/param"
	"github.com/stainless-sdks/browserbase-unofficial-go/packages/respjson"
)

// ContextService contains methods and other services that help with interacting
// with the browserbase-unofficial API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextService] method instead.
type ContextService struct {
	Options []option.RequestOption
}

// NewContextService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContextService(opts ...option.RequestOption) (r ContextService) {
	r = ContextService{}
	r.Options = opts
	return
}

// Create a Context
func (r *ContextService) New(ctx context.Context, body ContextNewParams, opts ...option.RequestOption) (res *CreateContextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contexts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Context
func (r *ContextService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ContextGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/contexts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Context
func (r *ContextService) Update(ctx context.Context, id string, opts ...option.RequestOption) (res *CreateContextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/contexts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type CreateContextResponse struct {
	ID string `json:"id,required"`
	// The cipher algorithm used to encrypt the user-data-directory. AES-256-CBC is
	// currently the only supported algorithm.
	CipherAlgorithm string `json:"cipherAlgorithm,required"`
	// The initialization vector size used to encrypt the user-data-directory.
	// [Read more about how to use it](/features/contexts).
	InitializationVectorSize int64 `json:"initializationVectorSize,required"`
	// The public key to encrypt the user-data-directory.
	PublicKey string `json:"publicKey,required"`
	// An upload URL to upload a custom user-data-directory.
	UploadURL string `json:"uploadUrl,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CipherAlgorithm          respjson.Field
		InitializationVectorSize respjson.Field
		PublicKey                respjson.Field
		UploadURL                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateContextResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateContextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextGetResponse struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The Project ID linked to the uploaded Context.
	ProjectID string    `json:"projectId,required"`
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ProjectID   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextNewParams struct {
	// The Project ID. Can be found in
	// [Settings](https://www.browserbase.com/settings).
	ProjectID string `json:"projectId,required"`
	paramObj
}

func (r ContextNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContextNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
