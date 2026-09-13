## cli get-object

Get object

### Synopsis

Get object

```
cli get-object [flags]
```

### Examples

```
  cli SDK get-object --team-id 8376e295-a8cf-4072-a303-3c75424a5f46 --object-type event --object-id 089a27e7-6fc9-4559-bfe1-0d332cda6b98
```

### Options

```
  -h, --help                 help for get-object
      --object-id string     [required]
      --object-type string   options: comment, deal, engagement, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
  -s, --select id            Comma-separated property slugs to return. Use dot notation for relationships. id is always returned at the top level. Defaults to all properties.
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
