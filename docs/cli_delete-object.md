## cli delete-object

Delete object

### Synopsis

Delete object

```
cli delete-object [flags]
```

### Examples

```
  cli SDK delete-object --team-id bb39ef04-8bee-4878-ba38-a537fd6e67c5 --object-type identity --object-id 2096d92b-1e2d-475e-b442-7992f9161602
```

### Options

```
  -h, --help                 help for delete-object
  -i, --if-match etag        Optimistic concurrency. Pass back the etag header from a previous GET of this record; the write only proceeds if the record hasn't changed since. Mismatch → 412 `precondition_failed`. Use `*` to require the record exists (any ETag accepted).
      --object-id string     [required]
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
