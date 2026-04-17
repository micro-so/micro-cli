// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/micro-cli/internal/mocktest"
	"github.com/stainless-sdks/micro-cli/internal/requestflag"
)

func TestPrismCreateObject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "create-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
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
			"prism", "create-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
		)
	})
}

func TestPrismDeleteObject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "delete-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--object-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismDuplicateObject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "duplicate-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--object-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismImportObjects(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "import-objects",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "identity",
			"--object", "{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, crm: {}, default: {foo: bar}, extended: {}}",
			"--options", "{caseInsensitive: true, crm_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, dedupe_by: dedupe_by}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(prismImportObjects)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "import-objects",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "identity",
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
			"prism", "import-objects",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "identity",
		)
	})
}

func TestPrismPatchObject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "patch-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--object-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
			"prism", "patch-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--object-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPrismRestoreObject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--team-id", "string",
			"prism", "restore-object",
			"--team-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--object-type", "deal",
			"--object-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
