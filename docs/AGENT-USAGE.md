# AI agents & scripts

The v1 CLI binary is `cli`. For automation, authenticate with `CLI_API_KEY`,
include `--team-id` on every data command, and disable interactive behavior:

```bash
CLI_API_KEY="$CLI_API_KEY" cli --no-interactive --output-format json \
  list-objects \
  --team-id "$TEAM_ID" \
  --object-type contact
```

`CLI_API_KEY` supplies the public API key. You may also pass `--api-key`, but
environment variables avoid exposing credentials in shell history and process
listings. For a person working locally, `cli configure` saves credentials
interactively; use `cli whoami` to inspect the active configuration.

`--no-interactive` disables prompts, explorer auto-launch, and TUI forms. Use
`--agent-mode` to explicitly enable structured errors and TOON output for an
agent; use `--agent-mode=false` to turn it off. Interactive `cli configure`
cannot run in agent mode, so agents should use `CLI_API_KEY` instead.

## Data commands

Every command below requires `--team-id "$TEAM_ID"` and, except where shown,
`--object-type <type>`.

```bash
# Read records
cli list-objects --team-id "$TEAM_ID" --object-type contact
cli get-object --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID"
cli find-object-by-slug --team-id "$TEAM_ID" --object-type contact \
  --slug email --value "agent@example.com"
cli query --team-id "$TEAM_ID" --object-type contact \
  --query '{"select":[],"combinator":"AND"}'
cli count-objects --team-id "$TEAM_ID" --object-type contact

# Create or update records
cli create-object --team-id "$TEAM_ID" --object-type contact \
  --body '{"default":{"email":"agent@example.com"}}'
cli patch-object --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID" \
  --body '{"default":{"name":"Agent"}}'
cli upsert-object --team-id "$TEAM_ID" --object-type contact \
  --slug email --value "agent@example.com" \
  --body '{"default":{"name":"Agent"}}'

# Import and lifecycle operations
cli import-objects --team-id "$TEAM_ID" --object-type contact \
  --objects '[{"default":{"email":"agent@example.com"}}]'
cli get-import-job --team-id "$TEAM_ID" --job-id "$JOB_ID"
cli delete-object --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID"
cli batch-delete-objects --team-id "$TEAM_ID" --object-type contact \
  --ids "$CONTACT_ID"
cli restore-object --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID"
cli duplicate-object --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID"

# Metadata and grants
cli get-metadata-properties-by-object-type --team-id "$TEAM_ID" --object-type contact
cli get-grant --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID"
cli update-grant --team-id "$TEAM_ID" --object-type contact --object-id "$CONTACT_ID" \
  --body '{"team_id":{}}'
```

## Output

Use `-o json` or `--output-format json` for machine-readable responses. The
supported output formats are `pretty`, `json`, `yaml`, `table`, and `toon`.
Use `-q '<expression>'` or `--jq '<expression>'` to filter a response; this
writes JSON and overrides the output format.

```bash
CLI_API_KEY="$CLI_API_KEY" cli --no-interactive --output-format json \
  get-object \
  --team-id "$TEAM_ID" \
  --object-type contact \
  --object-id "$CONTACT_ID"
```

Use `--dry-run` to preview a request without executing it, or `-d` / `--debug`
to log request and response diagnostics. Both write their diagnostics to stderr.

## Request bodies

Commands that accept request bodies support JSON through `--body` or stdin. For
example, these are equivalent:

```bash
CLI_API_KEY="$CLI_API_KEY" cli --no-interactive --output-format json \
  create-object \
  --team-id "$TEAM_ID" \
  --object-type contact \
  --body '{"default":{"email":"agent@example.com"}}'

printf '%s\n' '{"default":{"email":"agent@example.com"}}' | \
  CLI_API_KEY="$CLI_API_KEY" cli --no-interactive --output-format json \
    create-object \
    --team-id "$TEAM_ID" \
    --object-type contact
```

## Safe writes

Pass a unique `--idempotency-key` on write commands that support it:
`create-object`, `patch-object`, `upsert-object`, `import-objects`,
`batch-delete-objects`, `restore-object`, `duplicate-object`, and
`update-grant`. The server caches the first response for 24 hours, so retry the
same request with the same key. Reusing a key with a different request body
returns HTTP 409.

For `patch-object` and `delete-object`, pass `--if-match "$ETAG"` from a
previous GET to prevent overwriting a newer record. A mismatch returns HTTP 412.
Use `--if-match '*'` when the record only needs to exist.

## Pagination

`list-objects` accepts `--cursor` and returns `next_cursor`. Pass that value
back unchanged to fetch the next page. The `--limit` value is capped at 50.

## Exit codes and streams

Write the response to stdout and diagnostics to stderr so scripts can handle
them independently:

```bash
if ! CLI_API_KEY="$CLI_API_KEY" cli --no-interactive --output-format json \
  get-object \
  --team-id "$TEAM_ID" \
  --object-type contact \
  --object-id "$CONTACT_ID" \
  > response.json 2> error.log; then
  printf '%s\n' "cli request failed" >&2
  exit 1
fi
```
