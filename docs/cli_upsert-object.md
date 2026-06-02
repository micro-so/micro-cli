## cli upsert-object

Upsert by property value

### Synopsis

Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record matches, it is patched and 200 is returned. If none match, a new record is created (with the lookup property set if absent) and 201 is returned. If multiple records match, 409 is returned and you should patch by id instead.

```
cli upsert-object [flags]
```

### Examples

```
  cli SDK upsert-object --team-id 1fa0e3b3-bd09-49e7-99c2-f37eb2a55a73 --object-type ai_chat_message --slug <value> --value <value>
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --body-param string                          JSON value (one of: { extended: object, default: object, list: object })
  -h, --help                                       help for upsert-object
  -i, --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
      --object-type string                         options: deal, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
  -s, --slug string                                [required]
  -t, --team-id string                             [required]
  -v, --value string                               [required]
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
