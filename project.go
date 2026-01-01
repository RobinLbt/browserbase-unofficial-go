// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browserbaseunofficial

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/RobinLbt/browserbase-unofficial-go/internal/apijson"
	"github.com/RobinLbt/browserbase-unofficial-go/internal/requestconfig"
	"github.com/RobinLbt/browserbase-unofficial-go/option"
	"github.com/RobinLbt/browserbase-unofficial-go/packages/respjson"
)

// ProjectService contains methods and other services that help with interacting
// with the browserbase-unofficial API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// Project
func (r *ProjectService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/projects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all projects
func (r *ProjectService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Project, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Project Usage
func (r *ProjectService) Usage(ctx context.Context, id string, opts ...option.RequestOption) (res *ProjectUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/projects/%s/usage", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Project struct {
	ID             string    `json:"id,required"`
	CreatedAt      time.Time `json:"createdAt,required" format:"date-time"`
	DefaultTimeout int64     `json:"defaultTimeout,required"`
	Name           string    `json:"name,required"`
	OwnerID        string    `json:"ownerId,required"`
	UpdatedAt      time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DefaultTimeout respjson.Field
		Name           respjson.Field
		OwnerID        respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUsageResponse struct {
	BrowserMinutes int64 `json:"browserMinutes,required"`
	ProxyBytes     int64 `json:"proxyBytes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserMinutes respjson.Field
		ProxyBytes     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
