# AGENTS.md

`cli` is Micro's command-line client for working with CRM records, metadata, imports, and views. Use it when an agent needs to inspect or change Micro data from a shell; prefer it to the raw API because it supplies authentication, request validation, structured agent errors, output formatting, and dry-run diagnostics.

## Install and authenticate

```bash
# macOS/Linux
curl -fsSL https://raw.githubusercontent.com/micro-so/micro-cli/main/scripts/install.sh | bash

# API key used by the CLI
export CLI_API_KEY="..."

# Shell convenience only: the CLI has no configured team-ID environment variable.
export MICRO_TEAM_ID="..."
```

`CLI_API_KEY` is the real credential environment variable. Pass the required team on every command with `--team-id "$MICRO_TEAM_ID"` (or a literal ID). Flags override environment variables, and `cli configure` is interactive, so do not use it in agent mode.

## Golden paths

All examples use non-interactive, machine-readable output. Replace placeholder IDs and property slugs with values from the user's Micro workspace.

```bash
# List recent contacts, selecting only needed properties.
cli list-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --select "name,email" --sort "-updated_at" --limit 25 \
  --no-interactive --output-format json

# Fetch one deal with a limited field set.
cli get-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --object-id "$DEAL_ID" --select "name,amount,stage" \
  --no-interactive --output-format json

# Find one contact by an exact property value.
cli find-object-by-slug --team-id "$MICRO_TEAM_ID" --object-type contact \
  --slug email --value "person@example.com" \
  --no-interactive --output-format json

# Query contacts with a query body and selected fields.
cli query --team-id "$MICRO_TEAM_ID" --object-type contact \
  --query '{"select":["name","email"],"combinator":"AND"}' \
  --no-interactive --output-format json

# Create a deal. Use a new idempotency key for this logical operation.
cli create-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --body '{"default":{"name":"Expansion deal"}}' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

# Update a known deal. --if-match "*" requires that it already exists.
cli patch-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --object-id "$DEAL_ID" --body '{"default":{"stage":"qualified"}}' \
  --if-match "*" --idempotency-key "$IDEMPOTENCY_KEY" \
  --no-interactive --output-format json

# Create or update a contact identified by email.
cli upsert-object --team-id "$MICRO_TEAM_ID" --object-type contact \
  --slug email --value "person@example.com" \
  --body '{"default":{"name":"Ada Lovelace"}}' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

# Bulk import contacts. Reuse this same key only when retrying this same payload.
cli import-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --objects '[{"default":{"name":"Ada Lovelace","email":"ada@example.com"}}]' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

# Poll an asynchronous import using the job_id returned by import-objects.
cli get-import-job --team-id "$MICRO_TEAM_ID" --job-id "$JOB_ID" \
  --no-interactive --output-format json

# Get an unfiltered count without paging through records.
cli count-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --no-interactive --output-format json
```

## Output and scripting

- Use `--no-interactive` in scripts. Known coding-agent environments automatically enable agent mode, which defaults to TOON and returns structured errors; set `--output-format json` when a script needs JSON.
- Supported output formats are `pretty`, `json`, `yaml`, `table`, and `toon`; `--jq '<expression>'` filters output and emits JSON.
- Commands accepting request bodies support individual flags, `--body '<json>'`, or JSON on stdin. Priority is individual flags, then `--body`, then stdin: `printf '%s' '{"default":{"name":"Ada"}}' | cli create-object ...`.
- Exit code `0` means success; `1` means an API, validation, or transport error. Success is written to stdout and errors to stderr. Capture them separately when scripting.

## Safety and errors

- Always provide an idempotency key for imports and any retryable create, update, upsert, or patch. Never reuse a key with a changed request body.
- Never run `delete-object` or `batch-delete-objects` unless the user explicitly asks to delete those records.
- Prefer `list-objects` or `get-object` with `--select` and a small `--limit` to minimize data exposure and output size.
- Use `--dry-run` to inspect a request without sending it. Use `--debug` for request/response diagnostics; both write diagnostics to stderr and redact sensitive values.
- In agent mode or JSON output, errors are structured and include `error_type`, `message`, and sometimes `hints`. Check HTTP status and message before retrying; a `409 idempotency_key_mismatch` means the key was reused with a different payload.
- Command reference: `docs/cli_*.md` in this repository and <https://docs.micro.so>.
