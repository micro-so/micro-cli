## cli patch-metadata-property

Update a property definition

### Synopsis

Patches the editable fields (`name`, `icon`, `enabled`) of a property definition. `type` and scoping fields are immutable; `type` must be supplied in the body so the server knows which per-type table to write.

```
cli patch-metadata-property [flags]
```

### Examples

```
  cli SDK patch-metadata-property --team-id 0b096073-a49b-49b5-8a13-ac1c438aa436 --object-type event --property-id 9563eedd-e652-46fb-86a8-454fb3549b69 --type multiref_account
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --enabled                                    boolean flag
  -h, --help                                       help for patch-metadata-property
      --icon string                                string value
      --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
  -l, --list-id string                             string value
  -n, --name string                                string value
      --object-type string                         options: comment, deal, engagement, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
  -p, --property-id string                         [required]
      --team-id string                             [required]
      --type string                                Storage type for a property definition. (options: num, str, bool, date, text, byte, select_str, multi_str, multiselect_str, jsonb, ref_identity, ref_user, ref_organization, ref_organization_user, ref_contact, ref_thread, ref_message, ref_event, ref_account, multiref_ai_chat_message, multiref_action, multiref_contact, multiref_label, multiref_thread, multiref_messages, multiref_document, multiref_identity, multiref_organization, multiref_organization_user, multiref_engagement, multiref_attendee, multiref_meeting_entry, multiref_read_receipt, multiref_account) [required]
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
