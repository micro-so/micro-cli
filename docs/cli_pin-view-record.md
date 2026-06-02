## cli pin-view-record

Pin a record to the view (append to record_order)

### Synopsis

Pin a record to the view (append to record_order)

```
cli pin-view-record [flags]
```

### Examples

```
  cli SDK pin-view-record --team-id c22ee78c-f1a8-4032-a7b8-0bbfaaacc82d --view-object-type action --view-id 86146431-5e46-43e9-afa5-c7c54827d552 --object-id 202c347d-fd8a-4f11-8c83-677bb7c3e27c
```

### Options

```
  -h, --help                                       help for pin-view-record
  -i, --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
      --object-id string                           [required]
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
