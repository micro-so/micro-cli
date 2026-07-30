# AI agents & scripts

Use explicit flags and disable interactive behavior:

```bash
CLI_API_KEY="$MICRO_API_KEY" cli --no-interactive --output-format json \
  list-objects \
  --team-id "$MICRO_TEAM_ID" \
  --object-type contact
```

`CLI_API_KEY` is the environment variable for the public API key. You can also
pass it as `--api-key`, but environment variables avoid exposing credentials in
shell history and process listings. Flag values take precedence over environment
variables, which take precedence over saved credentials and configuration.

`--no-interactive` disables auto-prompts, the explorer, and TUI forms. The CLI
also detects supported coding-agent environments and enables agent mode, which
is non-interactive, defaults to TOON output, and formats API errors as
structured JSON on stderr. Pass `--agent-mode` to enable that behavior
explicitly, or `--agent-mode=false` to turn it off.

## Output

Use `--output-format json` for machine-readable JSON. There are no `--json` or
`--jsonl` flags. The supported output formats are `pretty`, `json`, `yaml`,
`table`, and `toon`.

```bash
CLI_API_KEY="$MICRO_API_KEY" cli --no-interactive --output-format json \
  get-object \
  --team-id "$MICRO_TEAM_ID" \
  --object-type contact \
  --object-id "$CONTACT_ID"
```

`--jq '<expression>'` filters a response and writes JSON, overriding
`--output-format`. Agent mode defaults to TOON output unless an explicit output
format is provided.

## Request bodies

Commands with a request body accept JSON through `--body` or stdin. For example,
these are equivalent:

```bash
CLI_API_KEY="$MICRO_API_KEY" cli --no-interactive --output-format json \
  create-object \
  --team-id "$MICRO_TEAM_ID" \
  --object-type contact \
  --body '{"default":{"email":"agent@example.com"}}'

printf '%s\n' '{"default":{"email":"agent@example.com"}}' | \
  CLI_API_KEY="$MICRO_API_KEY" cli --no-interactive --output-format json \
    create-object \
    --team-id "$MICRO_TEAM_ID" \
    --object-type contact
```

For body-capable commands, individual request flags take precedence over
`--body`, which takes precedence over stdin.

## Exit codes and streams

The CLI exits with `0` on success and `1` on errors. Successful responses are
written to stdout. Errors and diagnostics (`--debug` and `--dry-run`) are
written to stderr, so scripts can handle them independently:

```bash
if ! CLI_API_KEY="$MICRO_API_KEY" cli --no-interactive --output-format json \
  get-object \
  --team-id "$MICRO_TEAM_ID" \
  --object-type contact \
  --object-id "$CONTACT_ID" \
  > response.json 2> error.log; then
  cat error.log >&2
  exit 1
fi
```
