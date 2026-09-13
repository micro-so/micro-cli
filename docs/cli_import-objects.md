## cli import-objects

Import objects

### Synopsis

Import multiple objects in batch. Properties are keyed by slug. Automatically routes based on size: small batches complete synchronously and return 200 with the final `ImportJob`; large batches start an async job, return 202 with `status: processing` and a `Location` header, and can be polled via `GET /v2/prism/{teamId}/imports/{jobId}`.

```
cli import-objects [flags]
```

### Examples

```
  cli SDK import-objects --team-id 863ea5b9-7ed4-43bb-82aa-bc5508ef8716 --object-type document --objects '[{"list":{} }]'
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                                       help for import-objects
  -i, --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
      --object-type string                         options: comment, identity, organization, contact, action, document, engagement, deal [required]
      --objects string                             Array of objects to import with property values keyed by slug [required]
      --options string                             JSON object
  -t, --team-id string                             [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key x-api-key      Public API key generated from Micro settings. Sent as the x-api-key header and validated by AWS API Gateway in front of the service.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [cli](cli.md)	 - cli command-line interface
