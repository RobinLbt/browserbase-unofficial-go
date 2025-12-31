// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browserbaseunofficial

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apiform"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apijson"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/requestconfig"
	"github.com/stainless-sdks/browserbase-unofficial-go/option"
	"github.com/stainless-sdks/browserbase-unofficial-go/packages/respjson"
)

// ExtensionService contains methods and other services that help with interacting
// with the browserbase-unofficial API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionService] method instead.
type ExtensionService struct {
	Options []option.RequestOption
}

// NewExtensionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtensionService(opts ...option.RequestOption) (r ExtensionService) {
	r = ExtensionService{}
	r.Options = opts
	return
}

// Upload an Extension
func (r *ExtensionService) New(ctx context.Context, body ExtensionNewParams, opts ...option.RequestOption) (res *Extension, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/extensions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extension
func (r *ExtensionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Extension, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/extensions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Extension
func (r *ExtensionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/extensions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Extension struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	FileName  string    `json:"fileName,required"`
	// The Project ID linked to the uploaded Extension.
	ProjectID string    `json:"projectId,required"`
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		FileName    respjson.Field
		ProjectID   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Extension) RawJSON() string { return r.JSON.raw }
func (r *Extension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionNewParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r ExtensionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
