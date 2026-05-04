// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/micro-so/micro-cli/internal/mocktest"
)

func TestViewsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"views", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
			"--name", "name",
			"--view-type", "view_type",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--aggregation-prop-def-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--aggregation-type", "aggregation_type",
			"--column-layout", "{foo: bar}",
			"--combinator", "AND",
			"--created-at", "created_at",
			"--crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter", "{foo: bar}",
			"--group-by", "group_by",
			"--group-hidden-option-ids", "[{}]",
			"--group-hide-empty=true",
			"--group-sort", "group_sort",
			"--icon", "icon",
			"--select", "string",
			"--sort", "{foo: bar}",
			"--sort-order", "0",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--updated-at", "updated_at",
			"--user-id", "user_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"view_type: view_type\n" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"aggregation_prop_def_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"aggregation_type: aggregation_type\n" +
			"column_layout:\n" +
			"  foo: bar\n" +
			"combinator: AND\n" +
			"created_at: created_at\n" +
			"crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"filter:\n" +
			"  - foo: bar\n" +
			"group_by: group_by\n" +
			"group_hidden_option_ids:\n" +
			"  - {}\n" +
			"group_hide_empty: true\n" +
			"group_sort: group_sort\n" +
			"icon: icon\n" +
			"select:\n" +
			"  - string\n" +
			"sort:\n" +
			"  - foo: bar\n" +
			"sort_order: 0\n" +
			"team_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"updated_at: updated_at\n" +
			"user_id: user_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"views", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
		)
	})
}

func TestViewsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"views", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
			"--view-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--aggregation-prop-def-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--aggregation-type", "aggregation_type",
			"--column-layout", "{foo: bar}",
			"--combinator", "AND",
			"--crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter", "{foo: bar}",
			"--group-by", "group_by",
			"--group-hidden-option-ids", "[{}]",
			"--group-hide-empty=true",
			"--group-sort", "group_sort",
			"--icon", "icon",
			"--name", "name",
			"--select", "string",
			"--sort", "{foo: bar}",
			"--sort-order", "0",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--user-id", "user_id",
			"--view-type", "view_type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregation_prop_def_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"aggregation_type: aggregation_type\n" +
			"column_layout:\n" +
			"  foo: bar\n" +
			"combinator: AND\n" +
			"crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"filter:\n" +
			"  - foo: bar\n" +
			"group_by: group_by\n" +
			"group_hidden_option_ids:\n" +
			"  - {}\n" +
			"group_hide_empty: true\n" +
			"group_sort: group_sort\n" +
			"icon: icon\n" +
			"name: name\n" +
			"select:\n" +
			"  - string\n" +
			"sort:\n" +
			"  - foo: bar\n" +
			"sort_order: 0\n" +
			"team_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"user_id: user_id\n" +
			"view_type: view_type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"views", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
			"--view-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestViewsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"views", "delete",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
			"--view-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestViewsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"views", "get",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--view-object-type", "action",
			"--view-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
