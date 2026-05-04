// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/micro-so/micro-cli/internal/mocktest"
	"github.com/micro-so/micro-cli/internal/requestflag"
)

func TestPrismObjectsDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--crm", "{}",
			"--default", "{foo: bar}",
			"--extended", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"crm: {}\n" +
			"default:\n" +
			"  foo: bar\n" +
			"extended: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--crm", "{}",
			"--default", "{foo: bar}",
			"--extended", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"crm: {}\n" +
			"default:\n" +
			"  foo: bar\n" +
			"extended: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "update",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "delete",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsBulkCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "bulk-create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object", "{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, crm: {}, default: {foo: bar}, extended: {}}",
			"--options", "{caseInsensitive: true, crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, dedupe_by: dedupe_by}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(prismObjectsDocumentsBulkCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "bulk-create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object.id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object.crm", "{}",
			"--object.default", "{foo: bar}",
			"--object.extended", "{}",
			"--options.case-insensitive=true",
			"--options.crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--options.dedupe-by", "dedupe_by",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"objects:\n" +
			"  - id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    crm: {}\n" +
			"    default:\n" +
			"      foo: bar\n" +
			"    extended: {}\n" +
			"options:\n" +
			"  caseInsensitive: true\n" +
			"  crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  dedupe_by: dedupe_by\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "bulk-create",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsDuplicate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "duplicate",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "get",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "query",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "{select: [string], combinator: AND, crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, filter: [{foo: {foo: string}}], limit: 1, page: 0, sort: [{foo: asc}]}",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--box", "string",
			"--deleted=true",
			"--source", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(prismObjectsDocumentsQuery)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "query",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query.select", "[string]",
			"--query.combinator", "AND",
			"--query.crm-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query.filter", "[{foo: {foo: string}}]",
			"--query.limit", "1",
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
			"  limit: 1\n" +
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
			"prism:objects:documents", "query",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismObjectsDocumentsRestore(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism:objects:documents", "restore",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
