## cli query

Query

### Synopsis

Query

```
cli query [flags]
```

### Examples

```
  cli SDK query --team-id e0f153a6-8a1c-46f2-8b7b-338bf5f8256c --object-type document --query '{"select":[],"combinator":"AND"}'
```

### Options

```
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --boxes stringArray     list of values
  -c, --cursor query          Alternative location for the opaque cursor (sibling of query). Use whichever feels more natural; if both are present, `query.cursor` wins.
      --deleted               boolean flag
  -h, --help                  help for query
      --id string             JSON value (one of: string | array of string)
      --include-total total   When true, the response includes a total field with the unpaginated row count. Costs an additional pass over the result set — for unfiltered totals prefer `GET /v2/prism/{teamId}/{objectType}/count` instead.
      --object-type string    options: deal, identity, ai_chat_thread, ai_chat_message, document, organization, contact, action, event [required]
      --query string          [required]
  -s, --sources stringArray   list of values
  -t, --team-id string        [required]
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
