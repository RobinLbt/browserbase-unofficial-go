// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browserbaseunofficial_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RobinLbt/browserbase-unofficial-go"
	"github.com/RobinLbt/browserbase-unofficial-go/internal/testutil"
	"github.com/RobinLbt/browserbase-unofficial-go/option"
)

func TestSessionNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.New(context.TODO(), browserbaseunofficial.SessionNewParams{
		ProjectID: "projectId",
		BrowserSettings: browserbaseunofficial.SessionNewParamsBrowserSettings{
			BlockAds: browserbaseunofficial.Bool(true),
			Context: browserbaseunofficial.SessionNewParamsBrowserSettingsContext{
				ID:      "id",
				Persist: browserbaseunofficial.Bool(true),
			},
			ExtensionID: browserbaseunofficial.String("extensionId"),
			Fingerprint: browserbaseunofficial.SessionNewParamsBrowserSettingsFingerprint{
				Browsers:         []string{"chrome"},
				Devices:          []string{"desktop"},
				HTTPVersion:      1,
				Locales:          []string{"string"},
				OperatingSystems: []string{"android"},
				Screen: browserbaseunofficial.SessionNewParamsBrowserSettingsFingerprintScreen{
					MaxHeight: browserbaseunofficial.Int(0),
					MaxWidth:  browserbaseunofficial.Int(0),
					MinHeight: browserbaseunofficial.Int(0),
					MinWidth:  browserbaseunofficial.Int(0),
				},
			},
			LogSession:    browserbaseunofficial.Bool(true),
			RecordSession: browserbaseunofficial.Bool(true),
			SolveCaptchas: browserbaseunofficial.Bool(true),
			Viewport: browserbaseunofficial.SessionNewParamsBrowserSettingsViewport{
				Height: browserbaseunofficial.Int(0),
				Width:  browserbaseunofficial.Int(0),
			},
		},
		ExtensionID: browserbaseunofficial.String("extensionId"),
		KeepAlive:   browserbaseunofficial.Bool(true),
		Proxies:     map[string]any{},
		Region:      browserbaseunofficial.RegionUsWest2,
		Timeout:     browserbaseunofficial.Int(60),
	})
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Update(
		context.TODO(),
		"id",
		browserbaseunofficial.SessionUpdateParams{
			ProjectID: "projectId",
			Status:    browserbaseunofficial.SessionUpdateParamsStatusRequestRelease,
		},
	)
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.List(context.TODO(), browserbaseunofficial.SessionListParams{
		Status: browserbaseunofficial.SessionStatusRunning,
	})
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionNewUploads(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.NewUploads(
		context.TODO(),
		"id",
		browserbaseunofficial.SessionNewUploadsParams{
			File: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		},
	)
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionDebug(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Debug(context.TODO(), "id")
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionDownloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Sessions.Downloads(context.TODO(), "id")
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestSessionLogs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Logs(context.TODO(), "id")
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionRecording(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := browserbaseunofficial.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sessions.Recording(context.TODO(), "id")
	if err != nil {
		var apierr *browserbaseunofficial.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
