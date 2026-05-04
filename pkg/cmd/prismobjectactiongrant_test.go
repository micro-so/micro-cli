// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/micro-so/micro-cli/internal/mocktest"
)

func TestPrismObjectsActionsGrantUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:actions:grant", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--action-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--team-group-id", "{foo: a}",
			"--team-id", "{foo: a}",
			"--user-id", "{foo: a}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"team_group_id:\n" +
			"  - foo: a\n" +
			"team_id:\n" +
			"  foo: a\n" +
			"user_id:\n" +
			"  - foo: a\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:actions:grant", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--action-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsActionsGrantGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:actions:grant", "get",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--action-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
