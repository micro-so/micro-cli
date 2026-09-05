## cli delete-metadata-property

Delete a property definition

### Synopsis

Removes the property definition and any of its options. Fails with 409 `property_in_use` if records still reference the property.

```
cli delete-metadata-property [flags]
```

### Examples

```
  cli SDK delete-metadata-property --team-id d46ecc54-3958-4493-b59a-61646ebd4b72 --object-type identity --property-id 7c0febd9-fcea-4087-90aa-3ceebfde6a20 --type ref_account
```

### Options

```
  -h, --help                 help for delete-metadata-property
  -l, --list-id string       string value
      --object-type string   options: comment, deal, engagement, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
  -p, --property-id string   [required]
      --team-id string       [required]
      --type string          Storage type of this property definition. (options: num, str, bool, date, text, byte, select_str, multi_str, multiselect_str, jsonb, ref_identity, ref_user, ref_organization, ref_organization_user, ref_contact, ref_thread, ref_message, ref_event, ref_account, multiref_ai_chat_message, multiref_action, multiref_contact, multiref_label, multiref_thread, multiref_messages, multiref_document, multiref_identity, multiref_organization, multiref_organization_user, multiref_engagement, multiref_attendee, multiref_meeting_entry, multiref_read_receipt, multiref_account) [required]
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
