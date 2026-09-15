## cli update-grant

Update grant

### Synopsis

Update grant

```
cli update-grant [flags]
```

### Examples

```
  cli update-grant --team-id 80511222-6e30-46b4-95ab-c7019a2000d3 --object-type deal --object-id ff54806d-4891-4eee-bcc0-d9330e11c67d --body '{"team_id":{}}'
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -b, --body-param string                          JSON value (one of: { "user_id": object[], ... } | { "team_id": object, ... } | { "team_group_id": object[], ... })
  -h, --help                                       help for update-grant
  -i, --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
      --object-id string                           [required]
      --object-type string                         options: comment, deal, engagement, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
      --schema                                     Print the exact JSON Schema of the request body and exit
  -t, --team-id string                             [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key x-api-key      Public API key generated from Micro settings. Sent as the x-api-key header and validated by AWS API Gateway in front of the service.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [cli](cli.md)	 - cli command-line interface

### Machine interface

* `cli update-grant --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `cli update-grant --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `cli update-grant --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
