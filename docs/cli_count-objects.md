## cli count-objects

Total record count for an object type

### Synopsis

Returns the total number of records of this object type that the caller can see. Avoids the page-overshoot anti-pattern — clients no longer need to keep paging until `has_more` flips false to discover the total. Currently does not apply query filters; for a filtered total, pass `include_total: true` in a POST `/query` body.

```
cli count-objects [flags]
```

### Examples

```
  cli SDK count-objects --team-id 334bc1a2-d093-4be8-8078-f7ef2fa13681 --object-type organization
```

### Options

```
  -h, --help                 help for count-objects
  -l, --list-id string       Scope the count to a specific list/app.
      --object-type string   options: deal, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
  -t, --team-id string       [required]
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
