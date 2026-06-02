## cli list-objects

List records of an object type

### Synopsis

Convenience list endpoint. Equivalent to `POST /v2/prism/{teamId}/{objectType}/query` with an empty body, plus query-string sugar for the common cases. Any unrecognized query parameter is interpreted as an equality filter on a property of that name; pass arrays for `in`. Values are received as strings, so non-string property filters via this endpoint may not work — use the `query` endpoint for typed comparisons or anything beyond simple equality.

```
cli list-objects [flags]
```

### Examples

```
  cli SDK list-objects --team-id 25f1d299-6309-41ac-8567-2758463e8183 --object-type event
```

### Options

```
  -c, --cursor next_cursor   Opaque cursor from a previous response's next_cursor. Pass it back unchanged to fetch the next page.
      --deleted true         Include soft-deleted records. Pass the literal string true.
  -h, --help                 help for list-objects
  -i, --include-total true   When set to true, the response includes a `total` field with the unpaginated row count. Costs an extra pass; prefer `GET .../count` for the unfiltered total.
      --limit int            Maximum number of rows to return. Capped server-side at 50.
      --list-id string       Scope properties to a specific list/app.
      --object-type string   options: deal, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
      --select id            Comma-separated property slugs to return. Use dot notation for relationships. id is always returned at the top level. Defaults to all properties.
      --sort -               Comma-separated list of slugs. Prefix with - for descending. Example: `sort=-updated_at,name`.
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
