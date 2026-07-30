# AGENTS.md

`cli` is Micro's v1.0.0 command-line client. Use it to inspect or change Micro records, metadata, imports, grants, and views from a shell. Do not invent a `micro` binary or commands such as `micro login`, `micro status`, `micro get`, or `micro deals list`: they are not part of this CLI.

## Install and authenticate

```bash
# macOS/Linux
curl -fsSL https://raw.githubusercontent.com/micro-so/micro-cli/main/scripts/install.sh | bash

# Non-interactive authentication for agents and scripts
export CLI_API_KEY="..."

# Shell convenience only; pass it explicitly to every data command.
export MICRO_TEAM_ID="..."

# Verify the resolved credential source (the value is masked).
cli whoami
```

`CLI_API_KEY` is the credential environment variable. `cli configure` is interactive and stores credentials in the OS keychain when available (with a config-file fallback), so agents should use `CLI_API_KEY` instead. Flags take precedence over environment variables, which take precedence over stored credentials.

Every data command requires `--team-id "$MICRO_TEAM_ID"` (or a literal team ID). `whoami`, `configure`, and `version` are configuration commands, not data commands.

## Golden paths

Use `--no-interactive --output-format json` in scripts. Replace placeholder IDs and property slugs with values from the user's workspace.

```bash
# List a small page. Use the returned next_cursor unchanged for the next page.
cli list-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --select "name,email" --sort "-updated_at" --limit 25 \
  --no-interactive --output-format json

# Read one record, or find one by an exact property value.
cli get-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --object-id "$DEAL_ID" --select "name,amount,stage" \
  --no-interactive --output-format json

cli find-object-by-slug --team-id "$MICRO_TEAM_ID" --object-type contact \
  --slug email --value "person@example.com" \
  --no-interactive --output-format json

# Run a structured query or get an unfiltered object-type count.
cli query --team-id "$MICRO_TEAM_ID" --object-type contact \
  --query '{"select":["name","email"],"combinator":"AND"}' \
  --no-interactive --output-format json

cli count-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --no-interactive --output-format json

# Read the available property definitions before writing unfamiliar fields.
cli get-metadata-properties-by-object-type \
  --team-id "$MICRO_TEAM_ID" --object-type contact \
  --no-interactive --output-format json

# Create, patch, or upsert records. Use a new idempotency key per logical write.
cli create-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --body '{"default":{"name":"Expansion deal"}}' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

cli patch-object --team-id "$MICRO_TEAM_ID" --object-type deal \
  --object-id "$DEAL_ID" --body '{"default":{"stage":"qualified"}}' \
  --if-match "$ETAG" --idempotency-key "$IDEMPOTENCY_KEY" \
  --no-interactive --output-format json

cli upsert-object --team-id "$MICRO_TEAM_ID" --object-type contact \
  --slug email --value "person@example.com" \
  --body '{"default":{"name":"Ada Lovelace"}}' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

# Import records, then poll if the returned job is still processing.
cli import-objects --team-id "$MICRO_TEAM_ID" --object-type contact \
  --objects '[{"default":{"name":"Ada Lovelace","email":"ada@example.com"}}]' \
  --idempotency-key "$IDEMPOTENCY_KEY" --no-interactive --output-format json

cli get-import-job --team-id "$MICRO_TEAM_ID" --job-id "$JOB_ID" \
  --no-interactive --output-format json
```

The additional command families are `get-grant`/`update-grant`; metadata-property and metadata-option commands; `batch-update-objects`, `batch-delete-objects`, `restore-object`, and `duplicate-object`; and view commands including `create-view`, `get-view`, `patch-view`, `delete-view`, `list-view-records`, `reorder-view-records`, `pin-view-record`, and `unpin-view-record`.

## Output and scripting

- Global output flags are `-o, --output-format` (`pretty`, `json`, `yaml`, `table`, or `toon`) and `-q, --jq`.
- `--agent-mode` is automatically enabled in known coding-agent environments and supplies structured errors with TOON as the default output. `--no-interactive` disables prompts and the explorer UI.
- Commands accepting bodies support individual flags, `--body '<json>'`, or JSON on stdin. Individual flags win over `--body`, which wins over stdin.
- Use `--cursor "$NEXT_CURSOR"` to paginate list results; the next cursor appears as `next_cursor`. `list-objects --limit` is capped at 50.

## Safety rules

- Use `--dry-run` to preview a request without executing it; use `-d, --debug` only when request/response diagnostics are needed.
- For a mutation command that supports `--idempotency-key`, provide a fresh key for each logical operation. The server caches the first response for 24 hours; retry only the identical request with that key. A changed body with the same key returns `409 idempotency_key_mismatch`.
- `patch-object` and `delete-object` accept `--if-match`. Prefer the `etag` from a prior read; a mismatch returns `412 precondition_failed`. `--if-match "*"` only asserts that the record exists.
- Never run `delete-object`, `batch-delete-objects`, or `delete-view` unless the user explicitly asks to delete those specific resources. Prefer a narrow `--select` and a small `--limit` before any write.
- Use `docs/cli_*.md`, `cli <command> --help`, or `cli --usage` as the source of truth for command-specific flags.
