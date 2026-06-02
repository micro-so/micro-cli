## cli

cli command-line interface

### Synopsis

Command-line interface for cli

```
cli [flags]
```

### Options

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key x-api-key      Public API key generated from Micro settings. Sent as the x-api-key header and validated by AWS API Gateway in front of the service.
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                   help for cli
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

* [cli auth](cli_auth.md)	 - Manage authentication credentials
* [cli batch-delete-objects](cli_batch-delete-objects.md)	 - Bulk delete records (partial success)
* [cli batch-update-objects](cli_batch-update-objects.md)	 - Bulk update records (partial success)
* [cli configure](cli_configure.md)	 - Configure authentication credentials and preferences
* [cli count-objects](cli_count-objects.md)	 - Total record count for an object type
* [cli create-metadata-property](cli_create-metadata-property.md)	 - Create a property definition
* [cli create-metadata-property-option](cli_create-metadata-property-option.md)	 - Add an option to a select property
* [cli create-object](cli_create-object.md)	 - Create object
* [cli create-view](cli_create-view.md)	 - Create a view bundle (view + select/filter/sort)
* [cli delete-metadata-property](cli_delete-metadata-property.md)	 - Delete a property definition
* [cli delete-metadata-property-option](cli_delete-metadata-property-option.md)	 - Delete a property option
* [cli delete-object](cli_delete-object.md)	 - Delete object
* [cli delete-view](cli_delete-view.md)	 - Delete a view bundle
* [cli duplicate-object](cli_duplicate-object.md)	 - Duplicate object
* [cli explore](cli_explore.md)	 - Interactively browse and run commands
* [cli find-object-by-slug](cli_find-object-by-slug.md)	 - Find a record by property value
* [cli get-grant](cli_get-grant.md)	 - Get grant
* [cli get-import-job](cli_get-import-job.md)	 - Get the status of an import job
* [cli get-metadata-properties](cli_get-metadata-properties.md)	 - Get metadata properties
* [cli get-metadata-properties-by-object-type](cli_get-metadata-properties-by-object-type.md)	 - Get metadata properties by object type
* [cli get-object](cli_get-object.md)	 - Get object
* [cli get-view](cli_get-view.md)	 - Read a view bundle
* [cli import-objects](cli_import-objects.md)	 - Import objects
* [cli list-objects](cli_list-objects.md)	 - List records of an object type
* [cli list-view-records](cli_list-view-records.md)	 - List records selected by a view (filters and sorts applied; pinned record_order overlaid first)
* [cli patch-metadata-property](cli_patch-metadata-property.md)	 - Update a property definition
* [cli patch-metadata-property-option](cli_patch-metadata-property-option.md)	 - Update a property option
* [cli patch-object](cli_patch-object.md)	 - Patch object
* [cli patch-view](cli_patch-view.md)	 - Update a view bundle (select/filter/sort arrays are replaced wholesale when supplied)
* [cli pin-view-record](cli_pin-view-record.md)	 - Pin a record to the view (append to record_order)
* [cli query](cli_query.md)	 - Query
* [cli reorder-view-records](cli_reorder-view-records.md)	 - Bulk reorder pinned records
* [cli restore-object](cli_restore-object.md)	 - Restore object
* [cli unpin-view-record](cli_unpin-view-record.md)	 - Unpin a record from the view
* [cli update-grant](cli_update-grant.md)	 - Update grant
* [cli upsert-object](cli_upsert-object.md)	 - Upsert by property value
* [cli version](cli_version.md)	 - Print the CLI version
* [cli whoami](cli_whoami.md)	 - Display current authentication configuration
