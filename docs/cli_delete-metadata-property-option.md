## cli delete-metadata-property-option

Delete a property option

### Synopsis

Delete a property option

```
cli delete-metadata-property-option [flags]
```

### Examples

```
  cli SDK delete-metadata-property-option --team-id 93dec671-06ab-450d-9647-d684bc5aed13 --object-type deal --property-id 9e2daa26-45cb-4137-a8c8-d47e2bceddfc --option-id b1d01ebf-1389-40fb-bc05-f6cc721c2b82 --type ref_organization_user
```

### Options

```
  -h, --help                 help for delete-metadata-property-option
  -l, --list-id string       string value
      --object-type string   options: deal, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
      --option-id string     [required]
  -p, --property-id string   [required]
      --team-id string       [required]
      --type string          Storage type for a property definition. (options: num, str, bool, date, text, byte, select_str, multi_str, multiselect_str, jsonb, ref_identity, ref_user, ref_organization, ref_organization_user, ref_contact, ref_thread, ref_message, ref_event, ref_account, multiref_ai_chat_message, multiref_action, multiref_contact, multiref_label, multiref_thread, multiref_messages, multiref_document, multiref_identity, multiref_organization, multiref_organization_user, multiref_engagement, multiref_attendee, multiref_meeting_entry, multiref_read_receipt, multiref_account) [required]
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
