package flagutil

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateServerURL(t *testing.T) {
	for _, value := range []string{"/relative", "developers.micro.so", "https://", "ftp://example.com", "https://example.com/path?x=1", "https://user@example.com"} {
		if err := ValidateServerURL(value); err == nil {
			t.Errorf("ValidateServerURL(%q) accepted an invalid server", value)
		}
	}
	for _, value := range []string{"https://developers.micro.so", "http://127.0.0.1:8080"} {
		if err := ValidateServerURL(value); err != nil {
			t.Errorf("ValidateServerURL(%q) rejected a valid override: %v", value, err)
		}
	}
}

func TestServerURLValidationDoesNotExposeSecrets(t *testing.T) {
	for _, value := range []string{
		"https://user:credential-secret@example.com/path",
		"https://example.com/path?token=query-secret",
		"https://user:credential-secret@example.com/%zz?token=query-secret",
	} {
		err := ValidateServerURL(value)
		if err == nil {
			t.Fatalf("ValidateServerURL(%q) unexpectedly succeeded", value)
		}
		for _, diagnostic := range []string{err.Error(), errors.Unwrap(err).Error()} {
			if strings.Contains(diagnostic, "credential-secret") || strings.Contains(diagnostic, "query-secret") {
				t.Errorf("validation diagnostic exposed URL secret: %q", diagnostic)
			}
		}
		if reason, ok := err.(interface{ CLIReason() string }); !ok || reason.CLIReason() != "CLI_VALIDATION" {
			t.Errorf("validation reason lost: %T: %v", err, err)
		}
	}
}
