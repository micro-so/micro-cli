## cli reorder-view-records

Bulk reorder pinned records

### Synopsis

Bulk reorder pinned records

```
cli reorder-view-records [flags]
```

### Examples

```
  cli SDK reorder-view-records --team-id e7e8f968-7f65-450c-a491-7ed845dafc5d --view-object-type organization --view-id 865817ce-fd37-4afa-b47c-40ddb9aea061 --object-ids '["f9213dfa-b87d-4ac9-8859-50116fd1e918","6c34d94d-1ef3-4581-892a-249058c8931a","9bb9dc7e-0c94-41b7-913d-284158b2e9d1"]'
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                                       help for reorder-view-records
  -i, --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
      --object-ids stringArray                     [required]
  -t, --team-id string                             [required]
      --view-id string                             [required]
      --view-object-type string                    options: action, deal, document, event, identity, organization [required]
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
