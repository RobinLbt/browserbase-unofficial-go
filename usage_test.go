// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browserbaseunofficial_test

import (
	"context"
	"os"
	"testing"

	"github.com/RobinLbt/browserbase-unofficial-go"
	"github.com/RobinLbt/browserbase-unofficial-go/internal/testutil"
	"github.com/RobinLbt/browserbase-unofficial-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Mock server tests are disabled")
	createContextResponse, err := client.Contexts.New(context.TODO(), browserbaseunofficial.ContextNewParams{
		ProjectID: "projectId",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", createContextResponse.ID)
}
