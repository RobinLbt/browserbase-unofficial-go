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
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apiform"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apijson"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/apiquery"
	"github.com/stainless-sdks/browserbase-unofficial-go/internal/requestconfig"
	"github.com/stainless-sdks/browserbase-unofficial-go/option"
	"github.com/stainless-sdks/browserbase-unofficial-go/packages/param"
	"github.com/stainless-sdks/browserbase-unofficial-go/packages/respjson"
)

// SessionService contains methods and other services that help with interacting
// with the browserbase-unofficial API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r SessionService) {
	r = SessionService{}
	r.Options = opts
	return
}

// Create a Session
func (r *SessionService) New(ctx context.Context, body SessionNewParams, opts ...option.RequestOption) (res *SessionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Session
func (r *SessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Session
func (r *SessionService) Update(ctx context.Context, id string, body SessionUpdateParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Sessions
func (r *SessionService) List(ctx context.Context, query SessionListParams, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create Session Uploads
func (r *SessionService) NewUploads(ctx context.Context, id string, body SessionNewUploadsParams, opts ...option.RequestOption) (res *SessionNewUploadsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/uploads", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Session Live URLs
func (r *SessionService) Debug(ctx context.Context, id string, opts ...option.RequestOption) (res *SessionDebugResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/debug", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Session Downloads
func (r *SessionService) Downloads(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/downloads", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Session Logs
func (r *SessionService) Logs(ctx context.Context, id string, opts ...option.RequestOption) (res *[]SessionLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Session Recording
func (r *SessionService) Recording(ctx context.Context, id string, opts ...option.RequestOption) (res *[]SessionRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/recording", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Region string

const (
	RegionUsWest2      Region = "us-west-2"
	RegionUsEast1      Region = "us-east-1"
	RegionEuCentral1   Region = "eu-central-1"
	RegionApSoutheast1 Region = "ap-southeast-1"
)

type Session struct {
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	ExpiresAt time.Time `json:"expiresAt,required" format:"date-time"`
	// Indicates if the Session was created to be kept alive upon disconnections
	KeepAlive bool `json:"keepAlive,required"`
	// The Project ID linked to the Session.
	ProjectID string `json:"projectId,required"`
	// Bytes used via the [Proxy](/features/stealth-mode#proxies-and-residential-ips)
	ProxyBytes int64 `json:"proxyBytes,required"`
	// The region where the Session is running.
	//
	// Any of "us-west-2", "us-east-1", "eu-central-1", "ap-southeast-1".
	Region    Region    `json:"region,required"`
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// Any of "RUNNING", "ERROR", "TIMED_OUT", "COMPLETED".
	Status    SessionStatus `json:"status,required"`
	UpdatedAt time.Time     `json:"updatedAt,required" format:"date-time"`
	// CPU used by the Session
	AvgCPUUsage int64 `json:"avgCpuUsage"`
	// Optional. The Context linked to the Session.
	ContextID string    `json:"contextId"`
	EndedAt   time.Time `json:"endedAt" format:"date-time"`
	// Memory used by the Session
	MemoryUsage int64 `json:"memoryUsage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		KeepAlive   respjson.Field
		ProjectID   respjson.Field
		ProxyBytes  respjson.Field
		Region      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		AvgCPUUsage respjson.Field
		ContextID   respjson.Field
		EndedAt     respjson.Field
		MemoryUsage respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Session) RawJSON() string { return r.JSON.raw }
func (r *Session) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStatus string

const (
	SessionStatusRunning   SessionStatus = "RUNNING"
	SessionStatusError     SessionStatus = "ERROR"
	SessionStatusTimedOut  SessionStatus = "TIMED_OUT"
	SessionStatusCompleted SessionStatus = "COMPLETED"
)

type SessionNewResponse struct {
	ID string `json:"id,required"`
	// WebSocket URL to connect to the Session.
	ConnectURL string    `json:"connectUrl,required" format:"uri"`
	CreatedAt  time.Time `json:"createdAt,required" format:"date-time"`
	ExpiresAt  time.Time `json:"expiresAt,required" format:"date-time"`
	// Indicates if the Session was created to be kept alive upon disconnections
	KeepAlive bool `json:"keepAlive,required"`
	// The Project ID linked to the Session.
	ProjectID string `json:"projectId,required"`
	// Bytes used via the [Proxy](/features/stealth-mode#proxies-and-residential-ips)
	ProxyBytes int64 `json:"proxyBytes,required"`
	// The region where the Session is running.
	//
	// Any of "us-west-2", "us-east-1", "eu-central-1", "ap-southeast-1".
	Region Region `json:"region,required"`
	// HTTP URL to connect to the Session.
	SeleniumRemoteURL string `json:"seleniumRemoteUrl,required" format:"uri"`
	// Signing key to use when connecting to the Session via HTTP.
	SigningKey string    `json:"signingKey,required"`
	StartedAt  time.Time `json:"startedAt,required" format:"date-time"`
	// Any of "RUNNING", "ERROR", "TIMED_OUT", "COMPLETED".
	Status    SessionStatus `json:"status,required"`
	UpdatedAt time.Time     `json:"updatedAt,required" format:"date-time"`
	// CPU used by the Session
	AvgCPUUsage int64 `json:"avgCpuUsage"`
	// Optional. The Context linked to the Session.
	ContextID string    `json:"contextId"`
	EndedAt   time.Time `json:"endedAt" format:"date-time"`
	// Memory used by the Session
	MemoryUsage int64 `json:"memoryUsage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ConnectURL        respjson.Field
		CreatedAt         respjson.Field
		ExpiresAt         respjson.Field
		KeepAlive         respjson.Field
		ProjectID         respjson.Field
		ProxyBytes        respjson.Field
		Region            respjson.Field
		SeleniumRemoteURL respjson.Field
		SigningKey        respjson.Field
		StartedAt         respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		AvgCPUUsage       respjson.Field
		ContextID         respjson.Field
		EndedAt           respjson.Field
		MemoryUsage       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNewUploadsResponse struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionNewUploadsResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionNewUploadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionDebugResponse struct {
	DebuggerFullscreenURL string                     `json:"debuggerFullscreenUrl,required" format:"uri"`
	DebuggerURL           string                     `json:"debuggerUrl,required" format:"uri"`
	Pages                 []SessionDebugResponsePage `json:"pages,required"`
	WsURL                 string                     `json:"wsUrl,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DebuggerFullscreenURL respjson.Field
		DebuggerURL           respjson.Field
		Pages                 respjson.Field
		WsURL                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionDebugResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionDebugResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionDebugResponsePage struct {
	ID                    string `json:"id,required"`
	DebuggerFullscreenURL string `json:"debuggerFullscreenUrl,required" format:"uri"`
	DebuggerURL           string `json:"debuggerUrl,required" format:"uri"`
	FaviconURL            string `json:"faviconUrl,required" format:"uri"`
	Title                 string `json:"title,required"`
	URL                   string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		DebuggerFullscreenURL respjson.Field
		DebuggerURL           respjson.Field
		FaviconURL            respjson.Field
		Title                 respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionDebugResponsePage) RawJSON() string { return r.JSON.raw }
func (r *SessionDebugResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionLogsResponse struct {
	EventID   string `json:"eventId,required"`
	Method    string `json:"method,required"`
	PageID    int64  `json:"pageId,required"`
	SessionID string `json:"sessionId,required"`
	// milliseconds that have elapsed since the UNIX epoch
	Timestamp int64                       `json:"timestamp,required"`
	FrameID   string                      `json:"frameId"`
	LoaderID  string                      `json:"loaderId"`
	Request   SessionLogsResponseRequest  `json:"request"`
	Response  SessionLogsResponseResponse `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID     respjson.Field
		Method      respjson.Field
		PageID      respjson.Field
		SessionID   respjson.Field
		Timestamp   respjson.Field
		FrameID     respjson.Field
		LoaderID    respjson.Field
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionLogsResponseRequest struct {
	Params  map[string]any `json:"params,required"`
	RawBody string         `json:"rawBody,required"`
	// milliseconds that have elapsed since the UNIX epoch
	Timestamp int64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Params      respjson.Field
		RawBody     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionLogsResponseRequest) RawJSON() string { return r.JSON.raw }
func (r *SessionLogsResponseRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionLogsResponseResponse struct {
	RawBody string         `json:"rawBody,required"`
	Result  map[string]any `json:"result,required"`
	// milliseconds that have elapsed since the UNIX epoch
	Timestamp int64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RawBody     respjson.Field
		Result      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionLogsResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionLogsResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionRecordingResponse struct {
	ID string `json:"id,required"`
	// See
	// [rrweb documentation](https://github.com/rrweb-io/rrweb/blob/master/docs/recipes/dive-into-event.md).
	Data      map[string]any `json:"data,required"`
	SessionID string         `json:"sessionId,required"`
	// milliseconds that have elapsed since the UNIX epoch
	Timestamp int64 `json:"timestamp,required"`
	Type      int64 `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		SessionID   respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNewParams struct {
	// The Project ID. Can be found in
	// [Settings](https://www.browserbase.com/settings).
	ProjectID string `json:"projectId,required"`
	// The uploaded Extension ID. See
	// [Upload Extension](/reference/api/upload-an-extension).
	ExtensionID param.Opt[string] `json:"extensionId,omitzero"`
	// Set to true to keep the session alive even after disconnections. This is
	// available on the Startup plan only.
	KeepAlive param.Opt[bool] `json:"keepAlive,omitzero"`
	// Duration in seconds after which the session will automatically end. Defaults to
	// the Project's `defaultTimeout`.
	Timeout         param.Opt[int64]                `json:"timeout,omitzero"`
	BrowserSettings SessionNewParamsBrowserSettings `json:"browserSettings,omitzero"`
	// Proxy configuration. Can be true for default proxy, or an array of proxy
	// configurations.
	Proxies any `json:"proxies,omitzero"`
	// The region where the Session should run.
	//
	// Any of "us-west-2", "us-east-1", "eu-central-1", "ap-southeast-1".
	Region Region `json:"region,omitzero"`
	paramObj
}

func (r SessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNewParamsBrowserSettings struct {
	// Enable or disable ad blocking in the browser. Defaults to `false`.
	BlockAds param.Opt[bool] `json:"blockAds,omitzero"`
	// The uploaded Extension ID. See
	// [Upload Extension](/reference/api/upload-an-extension).
	ExtensionID param.Opt[string] `json:"extensionId,omitzero"`
	// Enable or disable session logging. Defaults to `true`.
	LogSession param.Opt[bool] `json:"logSession,omitzero"`
	// Enable or disable session recording. Defaults to `true`.
	RecordSession param.Opt[bool] `json:"recordSession,omitzero"`
	// Enable or disable captcha solving in the browser. Defaults to `true`.
	SolveCaptchas param.Opt[bool]                        `json:"solveCaptchas,omitzero"`
	Context       SessionNewParamsBrowserSettingsContext `json:"context,omitzero"`
	// See usage examples
	// [in the Stealth Mode page](/features/stealth-mode#fingerprinting).
	Fingerprint SessionNewParamsBrowserSettingsFingerprint `json:"fingerprint,omitzero"`
	Viewport    SessionNewParamsBrowserSettingsViewport    `json:"viewport,omitzero"`
	paramObj
}

func (r SessionNewParamsBrowserSettings) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParamsBrowserSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParamsBrowserSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type SessionNewParamsBrowserSettingsContext struct {
	// The Context ID.
	ID string `json:"id,required"`
	// Whether or not to persist the context after browsing. Defaults to `false`.
	Persist param.Opt[bool] `json:"persist,omitzero"`
	paramObj
}

func (r SessionNewParamsBrowserSettingsContext) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParamsBrowserSettingsContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParamsBrowserSettingsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// See usage examples
// [in the Stealth Mode page](/features/stealth-mode#fingerprinting).
type SessionNewParamsBrowserSettingsFingerprint struct {
	// Any of "chrome", "edge", "firefox", "safari".
	Browsers []string `json:"browsers,omitzero"`
	// Any of "desktop", "mobile".
	Devices []string `json:"devices,omitzero"`
	// Any of 1, 2.
	HTTPVersion float64 `json:"httpVersion,omitzero"`
	// Full list of locales is available
	// [here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language).
	Locales []string `json:"locales,omitzero"`
	// Note: `operatingSystems` set to `ios` or `android` requires `devices` to include
	// `"mobile"`.
	//
	// Any of "android", "ios", "linux", "macos", "windows".
	OperatingSystems []string                                         `json:"operatingSystems,omitzero"`
	Screen           SessionNewParamsBrowserSettingsFingerprintScreen `json:"screen,omitzero"`
	paramObj
}

func (r SessionNewParamsBrowserSettingsFingerprint) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParamsBrowserSettingsFingerprint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParamsBrowserSettingsFingerprint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionNewParamsBrowserSettingsFingerprint](
		"httpVersion", 1, 2,
	)
}

type SessionNewParamsBrowserSettingsFingerprintScreen struct {
	MaxHeight param.Opt[int64] `json:"maxHeight,omitzero"`
	MaxWidth  param.Opt[int64] `json:"maxWidth,omitzero"`
	MinHeight param.Opt[int64] `json:"minHeight,omitzero"`
	MinWidth  param.Opt[int64] `json:"minWidth,omitzero"`
	paramObj
}

func (r SessionNewParamsBrowserSettingsFingerprintScreen) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParamsBrowserSettingsFingerprintScreen
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParamsBrowserSettingsFingerprintScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNewParamsBrowserSettingsViewport struct {
	Height param.Opt[int64] `json:"height,omitzero"`
	Width  param.Opt[int64] `json:"width,omitzero"`
	paramObj
}

func (r SessionNewParamsBrowserSettingsViewport) MarshalJSON() (data []byte, err error) {
	type shadow SessionNewParamsBrowserSettingsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNewParamsBrowserSettingsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionUpdateParams struct {
	// The Project ID. Can be found in
	// [Settings](https://www.browserbase.com/settings).
	ProjectID string `json:"projectId,required"`
	// Set to `REQUEST_RELEASE` to request that the session complete. Use before
	// session's timeout to avoid additional charges.
	//
	// Any of "REQUEST_RELEASE".
	Status SessionUpdateParamsStatus `json:"status,omitzero,required"`
	paramObj
}

func (r SessionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set to `REQUEST_RELEASE` to request that the session complete. Use before
// session's timeout to avoid additional charges.
type SessionUpdateParamsStatus string

const (
	SessionUpdateParamsStatusRequestRelease SessionUpdateParamsStatus = "REQUEST_RELEASE"
)

type SessionListParams struct {
	// Any of "RUNNING", "ERROR", "TIMED_OUT", "COMPLETED".
	Status SessionStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SessionListParams]'s query parameters as `url.Values`.
func (r SessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionNewUploadsParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r SessionNewUploadsParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
