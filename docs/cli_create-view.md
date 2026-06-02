## cli create-view

Create a view bundle (view + select/filter/sort)

### Synopsis

Create a view bundle (view + select/filter/sort)

```
cli create-view [flags]
```

### Examples

```
  cli SDK create-view --team-id 408d4e54-e54c-4084-9e93-0fa16e13a37a --view-object-type organization --body-param.name <value> --body-param.view-type <value>
```

### Options

```
      --body string                                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.aggregation-prop-def-id string   string value
      --body-param.aggregation-type string          string value
      --body-param.column-layout string             value
      --body-param.combinator string                options: AND, OR
      --body-param.created-at string                string value
      --body-param.filter string                    Each entry is { slug: { comparator: value } }
      --body-param.group-by string                  Property slug to group by
      --body-param.group-hidden-option-ids string   JSON value (one of: array of any | ViewBundle_group_hidden_option_ids)
      --body-param.group-hide-empty string          boolean flag
      --body-param.group-sort string                string value
      --body-param.icon string                      string value
      --body-param.id string                        string value
      --body-param.list-id string                   string value
      --body-param.name string                      [required]
      --body-param.select stringArray               Property slugs (dot-paths permitted for refs)
      --body-param.sort string                      Each entry is { slug: 'asc' | 'desc' }
      --body-param.sort-order string                integer value
      --body-param.team-id string                   string value
      --body-param.updated-at string                string value
      --body-param.user-id string                   string value
      --body-param.view-type string                 [required]
  -h, --help                                        help for create-view
  -i, --idempotency-key idempotency_key_mismatch    A unique key (UUID or any opaque string up to 255 chars) that identifies this logical request. The server caches the first response under this key for 24 hours and replays it on retry — safe to use on every POST/PUT/PATCH to make network retries deterministic. Reusing the same key with a different body returns 409 idempotency_key_mismatch. Replays include the `idempotent-replay: true` response header.
  -t, --team-id string                              [required]
  -v, --view-object-type string                     options: action, deal, document, event, identity, organization [required]
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
