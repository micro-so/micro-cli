## cli unpin-view-record

Unpin a record from the view

### Synopsis

Unpin a record from the view

```
cli unpin-view-record [flags]
```

### Examples

```
  cli SDK unpin-view-record --team-id 6503d82a-a518-4343-ac14-d6ec97fe5b2c --view-object-type document --view-id 15414dcd-d5c0-4917-ac31-2a81fc1c8285 --object-id 6653be40-78a4-4d1d-8358-7989174daeba
```

### Options

```
  -h, --help                      help for unpin-view-record
      --object-id string          [required]
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
