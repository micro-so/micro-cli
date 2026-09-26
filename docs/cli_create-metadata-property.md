## cli create-metadata-property

Create a property definition

### Synopsis

Define a new property on this object type, scoped to the calling team. For `select_str` and `multiselect_str` types you may pre-seed the choices via `options`. Pass `list_id` in the body to also scope the definition to a specific list/app.

```
cli create-metadata-property [flags]
```

### Examples

```
  cli create-metadata-property --team-id ca34e81e-c9b3-4108-aade-5d4752c38ae4 --object-type action --type ref_organization_user --name <value>
```

### Options

```
      --body string                                Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -h, --help                                       help for create-metadata-property
      --icon string                                string value
      --idempotency-key idempotency_key_mismatch   A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
  -l, --list-id string                             string value
  -n, --name string                                Human-readable name. Must be unique within (team, list). [required]
      --object-type string                         options: comment, deal, engagement, identity, ai_chat_thread, ai_chat_message, document, action, event, organization, contact [required]
      --options type                               Only honored when type is `select_str` or `multiselect_str`.
      --schema                                     Print the exact JSON Schema of the request body and exit
  -s, --slug name                                  URL-safe identifier. Defaults to a slugified name. Disambiguated with a numeric suffix on conflict.
  -t, --team-id string                             [required]
      --type string                                Storage type for a property definition. (options: num, str, bool, date, text, byte, select_str, multi_str, multiselect_str, jsonb, ref_identity, ref_user, ref_organization, ref_organization_user, ref_contact, ref_thread, ref_message, ref_event, ref_account, multiref_ai_chat_message, multiref_action, multiref_contact, multiref_label, multiref_thread, multiref_messages, multiref_document, multiref_identity, multiref_organization, multiref_organization_user, multiref_engagement, multiref_attendee, multiref_meeting_entry, multiref_read_receipt, multiref_account) [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key x-api-key      Public API key generated from Micro settings. Sent as the x-api-key header and validated by AWS API Gateway in front of the service.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [cli](cli.md)	 - cli command-line interface

### Machine interface

* `cli create-metadata-property --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `cli create-metadata-property --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `cli create-metadata-property --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
