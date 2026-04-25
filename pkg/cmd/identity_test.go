// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/micro-cli/internal/mocktest"
	"github.com/stainless-sdks/micro-cli/internal/requestflag"
)

func TestIdentitiesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--crm", "{}",
			"--default", "{}",
			"--extended", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"crm: {}\n" +
			"default: {}\n" +
			"extended: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestIdentitiesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--identity-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--crm", "{}",
			"--default", "{}",
			"--extended", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"crm: {}\n" +
			"default: {}\n" +
			"extended: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--identity-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestIdentitiesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "list",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "{select: [string], combinator: AND, crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, filter: [{foo: {foo: string}}], limit: 0, page: 0, sort: [{foo: asc}]}",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--box", "string",
			"--deleted=true",
			"--source", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(identitiesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "list",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query.select", "[string]",
			"--query.combinator", "AND",
			"--query.crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query.filter", "[{foo: {foo: string}}]",
			"--query.limit", "0",
			"--query.page", "0",
			"--query.sort", "[{foo: asc}]",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--box", "string",
			"--deleted=true",
			"--source", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query:\n" +
			"  select:\n" +
			"    - string\n" +
			"  combinator: AND\n" +
			"  crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  filter:\n" +
			"    - foo:\n" +
			"        foo: string\n" +
			"  limit: 0\n" +
			"  page: 0\n" +
			"  sort:\n" +
			"    - foo: asc\n" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"boxes:\n" +
			"  - string\n" +
			"deleted: true\n" +
			"sources:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "list",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestIdentitiesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "delete",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--identity-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestIdentitiesImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "import",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object", "{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, crm: {}, default: {}, extended: {}}",
			"--options", "{caseInsensitive: true, crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, dedupe_by: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, dedupe_type: str}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(identitiesImport)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "import",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object.id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object.crm", "{}",
			"--object.default", "{}",
			"--object.extended", "{}",
			"--options.case-insensitive=true",
			"--options.crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--options.dedupe-by", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--options.dedupe-type", "str",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"objects:\n" +
			"  - id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    crm: {}\n" +
			"    default: {}\n" +
			"    extended: {}\n" +
			"options:\n" +
			"  caseInsensitive: true\n" +
			"  crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  dedupe_by: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  dedupe_type: str\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"identities", "import",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
