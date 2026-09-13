## cli list-view-records

List records selected by a view (filters and sorts applied; pinned record_order overlaid first)

### Synopsis

List records selected by a view (filters and sorts applied; pinned record_order overlaid first)

```
cli list-view-records [flags]
```

### Examples

```
  cli SDK list-view-records --team-id 5f478bf6-b37b-450e-9850-7b22e98c2777 --view-object-type action --view-id 7872a70f-a77f-401e-8bea-203dc7aa1e06
```

### Options

```
  -c, --cursor next_cursor        Opaque cursor from a previous response's next_cursor. Pass it back unchanged to fetch the next page. When set, `page` and `limit` are derived from the cursor.
  -h, --help                      help for list-view-records
  -l, --limit int                 integer value
  -p, --page cursor               Page number (1-based). Prefer cursor.
  -t, --team-id string            [required]
      --view-id string            [required]
      --view-object-type string   options: comment, action, deal, engagement, document, event, identity, organization [required]
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
