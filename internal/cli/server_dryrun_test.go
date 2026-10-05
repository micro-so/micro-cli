package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"openapi/internal/config"
)

func TestDefaultServerDryRun(t *testing.T) {
	t.Setenv("CLI_SERVER_URL", "")
	config.Reset()
	t.Cleanup(config.Reset)
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	err = ExecuteRoot(context.Background(), root, []string{"count-objects", "--team-id", "test-team", "--object-type", "contact", "--api-key", "dummy", "--no-interactive", "--dry-run", "-o", "json"})
	if err != nil {
		t.Fatal(err)
	}
	var preview struct {
		Request struct {
			URL string `json:"url"`
		} `json:"request"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &preview); err != nil {
		t.Fatalf("preview %q: %v", stdout.String(), err)
	}
	if want := "https://developers.micro.so/v2/prism/test-team/contact/count"; preview.Request.URL != want {
		t.Fatalf("request URL = %q, want %q", preview.Request.URL, want)
	}
}

func TestServerEnvironmentOverrideDryRun(t *testing.T) {
	t.Setenv("CLI_SERVER_URL", "https://staging.developers.micro.so")
	config.Reset()
	t.Cleanup(config.Reset)
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	err = ExecuteRoot(context.Background(), root, []string{"count-objects", "--team-id", "test-team", "--object-type", "contact", "--api-key", "dummy", "--no-interactive", "--dry-run", "-o", "json"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stdout.String(), "https://staging.developers.micro.so/v2/prism/test-team/contact/count") {
		t.Fatalf("CLI_SERVER_URL override missing from preview: %q", stdout.String())
	}
}

func TestDryRunThenLiveOnSameCommandTree(t *testing.T) {
	config.Reset()
	t.Cleanup(config.Reset)
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Method != http.MethodGet || r.URL.Path != "/v2/prism/test-team/contact/count" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"total":42}`)
	}))
	defer server.Close()
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	args := []string{"count-objects", "--team-id", "test-team", "--object-type", "contact", "--api-key", "dummy", "--no-interactive", "--server-url", server.URL, "-o", "json"}
	if err := ExecuteRoot(context.Background(), root, append(append([]string{}, args...), "--dry-run")); err != nil {
		t.Fatal(err)
	}
	if requests != 0 {
		t.Fatalf("dry run sent %d requests", requests)
	}
	if !strings.Contains(stdout.String(), server.URL) {
		t.Fatalf("preview missing explicit server: %q", stdout.String())
	}
	stdout.Reset()
	stderr.Reset()
	if err := ExecuteRoot(context.Background(), root, append(append([]string{}, args...), "--dry-run=false")); err != nil {
		t.Fatal(err)
	}
	if requests != 1 {
		t.Fatalf("live invocation sent %d requests, want 1", requests)
	}
	if !strings.Contains(stdout.String(), "42") {
		t.Fatalf("live response hidden: stdout=%q stderr=%q", stdout.String(), stderr.String())
	}
}

func TestInvalidServerURLDoesNotRenderSecrets(t *testing.T) {
	config.Reset()
	t.Cleanup(config.Reset)
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	err = ExecuteRoot(context.Background(), root, []string{"count-objects", "--team-id", "test-team", "--object-type", "contact", "--api-key", "dummy", "--no-interactive", "--server-url", "https://user:credential-secret@example.com/%zz?token=query-secret", "-o", "json"})
	if err == nil {
		t.Fatal("invalid server URL accepted")
	}
	rendered := stdout.String() + stderr.String() + err.Error()
	if strings.Contains(rendered, "credential-secret") || strings.Contains(rendered, "query-secret") {
		t.Fatalf("CLI error exposed URL secret: %q", rendered)
	}
	if !strings.Contains(rendered, "invalid --server-url") {
		t.Fatalf("CLI error lost validation context: %q", rendered)
	}
}
