## cli create-metadata-property-option

Add an option to a select property

### Synopsis

Adds a single option to a `select_str` or `multiselect_str` property definition. Body must include `type` so the server knows which per-type option table to write.

```
cli create-metadata-property-option [flags]
```

### Examples

```
  cli SDK create-metadata-property-option --team-id 638bb81b-7e3c-4105-a7b3-b35949652b97 --object-type action --property-id e075c94a-7a24-4c9a-901c-f554ecee71f4 --type multiref_read_receipt --value <value>
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --color-scheme string                        string value
      --description string                         string value
  -h, --help                                       help for create-metadata-property-option
      --icon string                                string value
      --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
  -l, --list-id string                             Scope the option to a specific list/app.
      --object-type string                         options: deal, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
      --option-group string                        string value
  -p, --property-id string                         [required]
      --slug value                                 URL-safe identifier. Defaults to a slugified value.
      --sort-index string                          integer value
      --team-id string                             [required]
      --type string                                Storage type for a property definition. (options: num, str, bool, date, text, byte, select_str, multi_str, multiselect_str, jsonb, ref_identity, ref_user, ref_organization, ref_organization_user, ref_contact, ref_thread, ref_message, ref_event, ref_account, multiref_ai_chat_message, multiref_action, multiref_contact, multiref_label, multiref_thread, multiref_messages, multiref_document, multiref_identity, multiref_organization, multiref_organization_user, multiref_engagement, multiref_attendee, multiref_meeting_entry, multiref_read_receipt, multiref_account) [required]
  -v, --value string                               Display value for the option. [required]
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
