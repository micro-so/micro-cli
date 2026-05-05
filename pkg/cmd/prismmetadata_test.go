// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/micro-cli/internal/mocktest"
)

func TestPrismMetadataList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:metadata", "list",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--autofill=true",
			"--crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--term", "term",
		)
	})
}
