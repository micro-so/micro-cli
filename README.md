# cli

Command-line interface for the *Prism* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cli](#cli)
  * [CLI Installation](#cli-installation)
  * [Shell Completion](#shell-completion)
  * [CLI Example Usage](#cli-example-usage)
  * [For AI agents](#for-ai-agents)
  * [Authentication](#authentication)
  * [Commands](#commands)
  * [Request Body Input](#request-body-input)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start CLI Installation [installation] -->
## CLI Installation

### Quick Install (Linux/macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/micro-so/micro-cli/main/scripts/install.sh | bash
```

### Quick Install (Windows PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/micro-so/micro-cli/main/scripts/install.ps1 | iex
```

### Go Install

Alternatively, install directly via Go:

```bash
go install openapi/cmd/cli@latest
```

### Manual Download

Download pre-built binaries for your platform from the [releases page](https://github.com/micro-so/micro-cli/releases).
<!-- End CLI Installation [installation] -->

<!-- Start Shell Completion [completion] -->
## Shell Completion

Shell completions are available for Bash, Zsh, Fish, and PowerShell.

### Bash

```bash
# Add to ~/.bashrc:
source <(cli completion bash)

# Or install permanently:
cli completion bash > /etc/bash_completion.d/cli
```

### Zsh

```zsh
# Add to ~/.zshrc:
source <(cli completion zsh)

# Or install permanently:
cli completion zsh > "${fpath[1]}/_cli"
```

### Fish

```fish
cli completion fish | source

# Or install permanently:
cli completion fish > ~/.config/fish/completions/cli.fish
```

### PowerShell

```powershell
cli completion powershell | Out-String | Invoke-Expression
```
<!-- End Shell Completion [completion] -->

<!-- Start CLI Example Usage [usage] -->
## CLI Example Usage

### Example

```bash
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192

```
<!-- End CLI Example Usage [usage] -->

<!-- Start For AI agents [agents] -->
## For AI agents

This CLI is built to be driven by AI coding agents as well as people: everything an agent needs is discoverable from the binary itself, and every command can be validated without credentials. Work down this ladder:

| Run | You get |
|-----|---------|
| `cli --help`, `cli restore-object --help` | Commands by category, runnable examples, flags |
| `cli --usage`, `cli restore-object --usage` | The command surface as machine-readable [KDL](https://kdl.dev): commands, aliases, flags, defaults, env vars, config keys |
| `cli batch-delete-objects --schema` | The exact JSON Schema of the command's request body (all `$ref`s bundled) — build a valid `--body` from it |
| `cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --dry-run` | The exact HTTP request (method, URL, headers, body), with no credentials or network call |
| `cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format json` (or `--jq`) | Machine-readable output |

### Discover the command surface

```bash
# Every command, flag, default, env var and config key, as KDL
cli --usage

# One command's subtree only
cli restore-object --usage
```

### Read the exact request schema

`--schema` is available on every command that accepts a request body (`--body`, stdin, or a whole-body flag where the command has one), including intent commands. It prints the JSON Schema the request is validated against and exits without calling the API.

```bash
# JSON Schema (draft 2020-12) of the request body, with every $ref bundled under $defs
cli batch-delete-objects --schema
```

### Probe before you spend

Start quota-spending commands with `--dry-run`. It validates inputs, resolves the request, redacts secrets and binary payloads, makes no network call, and exits 0. It never reads the OS keychain; credentials supplied by flag, environment, or config file are included only as `[REDACTED]`.

```bash
# Human preview: the [DRY-RUN] block is on stderr and stdout is empty
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --dry-run

# Machine preview: compact JSON on stdout and silent stderr
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --dry-run --output-format json
```

The machine form writes one object per would-be request, one per line (NDJSON for multi-request commands), with exactly this shape:

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

`body` is a parsed JSON value when the body is JSON, a string for text, `"<bytes:N>"` for binary data, and `null` when absent. An explicit caller `--jq` also selects this JSON preview protocol, but the filter is not applied to preview objects. Command-declared jq presets do not select or filter the preview.

Local mutation commands make no request under `--dry-run`: instead of a preview they emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object. `select(.request)` keeps only would-be requests; `select(.local)` keeps the local no-ops.

### Machine-readable output

```bash
# JSON on stdout
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format json

# Filter or reshape with a jq expression (always emits JSON, overrides --output-format)
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --jq '.'

# Print jq string results as plain text instead of JSON strings (like jq -r)
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --jq '.' --raw-output
```

`--output-format toon` emits [TOON](https://github.com/toon-format/spec), a compact line-oriented format that uses fewer tokens than JSON; it is the default in agent mode.

### Interactive mode
Required-input prompts and guided `configure` / `auth login` forms are enabled by default. Required-input prompts require an interactive terminal; off-TTY forms read line input from stdin. Use `--no-interactive` to force flag-only execution.

```bash
# Prompt for missing command inputs
cli restore-object --interactive

# Open the guided configuration form
cli configure --interactive

# Explicitly launch the terminal command explorer
cli explore
```

### Agent mode and structured errors
Agent mode turns on automatically when a known agent environment is detected (`CLAUDECODE`, `CURSOR_AGENT`, `CODEX`, `AIDER`, `CLINE`, `WINDSURF_AGENT`, `GITHUB_COPILOT`, `AMAZON_Q`, `GEMINI_CODE_ASSIST`, `SRC_CODY`) or with `--agent-mode` (`--agent-mode=false` disables detection).
In agent mode interactive prompts never launch, output defaults to TOON, and every failure — API errors and CLI usage errors alike — is one JSON envelope on stderr:
Outside agent mode, explicit JSON and `--jq` preserve the compatibility envelope without classification; enable agent mode to request the classified contract.

```json
{
  "error": "...",
  "error_type": "validation_error",
  "error_reason": "CLI_VALIDATION",
  "exit_code": 2,
  "message": "human-readable message",
  "hints": ["what to try next"]
}
```

`error_type` is one of `authentication_error`, `authorization_error`, `not_found`, `validation_error`, `rate_limit_error`, `server_error`, `api_error`, `connection_error`, `protocol_error`, `runtime_error`, `unsupported_error`, `async_failed`, `async_timeout`, `async_unknown_state`. Classification derives from the HTTP status and transport evidence; `error_reason` is absent for API errors. Status-less local failures may use `CLI_VALIDATION`, `CLI_CONNECTION`, `CLI_PROTOCOL`, `CLI_RUNTIME`, `CLI_UNAVAILABLE`, `CLI_AUTHENTICATION`, or the async polling reasons `CLI_ASYNC_FAILED`, `CLI_ASYNC_TIMEOUT`, and `CLI_ASYNC_UNKNOWN_STATE`. `hints` preserves server guidance first, adds the most specific local taxonomy guidance, then typed CLI and command-specific guidance, removing exact duplicates. `exit_code` is always the code for the final `error_type` shown in the envelope: 1 runtime, 2 usage, or 3 authentication/authorization.
<!-- End For AI agents [agents] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
cli --api-key "$CLI_API_KEY" restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192
```

### 2. Environment variables

Set credentials via environment variables:

| Variable | Description |
|----------|-------------|
| `CLI_API_KEY` | Public API key generated from Micro settings. Sent as the `x-api-key` header and validated by AWS API Gateway in front of the service. |

### 3. OS Keychain (recommended for workstations)

Credentials are stored securely in your operating system's keychain when you run:

```bash
cli configure
```

Secret credentials (tokens, API keys, passwords) are automatically stored in:
- **macOS**: Keychain
- **Linux**: GNOME Keyring / KWallet (via D-Bus Secret Service)
- **Windows**: Windows Credential Locker

If no keychain is available (e.g., in CI environments), credentials fall back to the config file.

### 4. Configuration file

Run the interactive `configure` command to store non-secret settings:

```bash
cli configure
```

Configuration is stored in `~/.config/cli/config.yaml`.
<!-- End Authentication [security] -->

<!-- Start Commands [operations] -->
## Commands

<details open>
<summary>Available commands</summary>

* [`restore-object`](docs/cli_restore-object.md) - Restore object
* [`duplicate-object`](docs/cli_duplicate-object.md) - Duplicate object
* [`list-objects`](docs/cli_list-objects.md) - List records of an object type
* [`create-object`](docs/cli_create-object.md) - Create object
* [`find-object-by-slug`](docs/cli_find-object-by-slug.md) - Find a record by property value
* [`upsert-object`](docs/cli_upsert-object.md) - Upsert by property value
* [`get-object`](docs/cli_get-object.md) - Get object
* [`delete-object`](docs/cli_delete-object.md) - Delete object
* [`patch-object`](docs/cli_patch-object.md) - Patch object
* [`get-grant`](docs/cli_get-grant.md) - Get grant
* [`update-grant`](docs/cli_update-grant.md) - Update grant
* [`query`](docs/cli_query.md) - Query
* [`get-metadata-properties`](docs/cli_get-metadata-properties.md) - Get metadata properties
* [`count-objects`](docs/cli_count-objects.md) - Total record count for an object type
* [`get-metadata-properties-by-object-type`](docs/cli_get-metadata-properties-by-object-type.md) - Get metadata properties by object type
* [`create-metadata-property`](docs/cli_create-metadata-property.md) - Create a property definition
* [`patch-metadata-property`](docs/cli_patch-metadata-property.md) - Update a property definition
* [`delete-metadata-property`](docs/cli_delete-metadata-property.md) - Delete a property definition
* [`create-metadata-property-option`](docs/cli_create-metadata-property-option.md) - Add an option to a select property
* [`patch-metadata-property-option`](docs/cli_patch-metadata-property-option.md) - Update a property option
* [`delete-metadata-property-option`](docs/cli_delete-metadata-property-option.md) - Delete a property option
* [`batch-update-objects`](docs/cli_batch-update-objects.md) - Bulk update records (partial success)
* [`batch-delete-objects`](docs/cli_batch-delete-objects.md) - Bulk delete records (partial success)
* [`get-import-job`](docs/cli_get-import-job.md) - Get the status of an import job
* [`import-objects`](docs/cli_import-objects.md) - Import objects
* [`create-view`](docs/cli_create-view.md) - Create a view bundle (view + select/filter/sort)
* [`get-view`](docs/cli_get-view.md) - Read a view bundle
* [`patch-view`](docs/cli_patch-view.md) - Update a view bundle (select/filter/sort arrays are replaced wholesale when supplied)
* [`delete-view`](docs/cli_delete-view.md) - Delete a view bundle
* [`list-view-records`](docs/cli_list-view-records.md) - List records selected by a view (filters and sorts applied; pinned record_order overlaid first)
* [`reorder-view-records`](docs/cli_reorder-view-records.md) - Bulk reorder pinned records
* [`pin-view-record`](docs/cli_pin-view-record.md) - Pin a record to the view (append to record_order)
* [`unpin-view-record`](docs/cli_unpin-view-record.md) - Unpin a record from the view

</details>
<!-- End Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Commands that accept a request body take it three ways, with a clear priority chain: individual field flags (highest priority), the whole body as JSON via `--body`, and JSON piped on stdin (lowest priority). Later sources never override earlier ones; a body-bearing command prints its exact request schema with `--schema`.
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Output Formats [output-formats] -->
## Output Formats

Every command supports a `--output-format` flag that controls how the response is rendered to stdout.

### Available formats

| Format | Flag | Description |
|--------|------|-------------|
| Pretty | `--output-format pretty` (default) | Aligned key-value pairs with color, nested indentation. Human-readable at a glance. |
| JSON | `--output-format json` | JSON output. Passthrough when the response is already JSON (preserves original field order and numeric precision). Falls back to typed marshaling otherwise. |
| YAML | `--output-format yaml` | YAML output via standard marshaling. |
| Table | `--output-format table` | Tabular output for array responses. |
| TOON | `--output-format toon` | [Token-Oriented Object Notation](https://github.com/toon-format/spec) — a compact, line-oriented format that typically uses 30–60% fewer tokens than JSON. Well-suited for piping responses into LLM prompts. |

```bash
# Default pretty output
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192

# Machine-readable JSON
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format json

# TOON for LLM-friendly compact output
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format toon

# Pipe JSON to jq without using --output-format
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format json | jq '.'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --jq '.'

# Reshape with any jq program; --raw-output prints string results as plain text (like jq -r)
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --jq '.' --raw-output
```

### Color control

Use `--color` to control terminal colors:

| Value | Behavior |
|-------|----------|
| `auto` (default) | Color when stdout is a TTY, plain text otherwise |
| `always` | Always colorize |
| `never` | Never colorize |

The `NO_COLOR` and `FORCE_COLOR` environment variables are also respected.

### Streaming and pagination

When using `--all` (pagination) or streaming operations, output is written incrementally as items arrive:

| Format | Streaming behavior |
|--------|-------------------|
| `json` | One compact JSON object per line ([NDJSON](https://github.com/ndjson/ndjson-spec)) |
| `yaml` | YAML documents separated by `---` |
| `toon` | One TOON-encoded object per block, separated by blank lines |
| `pretty` (default) | Pretty-printed items separated by blank lines |
<!-- End Output Formats [output-formats] -->

<!-- Start Error Handling [errors] -->
## Error Handling

The CLI uses standard exit codes to indicate success or failure:

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success |
| `1` | Runtime/API failure |
| `2` | Usage or input failure |
| `3` | Authentication or authorization failure |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --output-format json > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
This CLI uses unclassified error rendering outside agent mode: pretty and TOON print the API error text as received, while `--output-format json` and `--jq` emit the unclassified envelope (including the configure `_hint` for HTTP 401/403) plus `exit_code`. Agent mode always emits the classified JSON envelope with `exit_code`, `error_type`, optional `error_reason`, `message`, `hints`, and optional `status_code` — see [For AI agents](#for-ai-agents).
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --dry-run
```

In human output modes, stdout is empty and the `[DRY-RUN]` block goes to stderr. It includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

With `--output-format json`, or with a caller-explicit `--jq`, stderr is silent and stdout is NDJSON: one compact preview object per would-be request. The jq filter is not applied, and command-declared jq presets do not select the JSON protocol.

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

JSON bodies remain structured; text bodies are strings; binary bodies are `"<bytes:N>"`; absent bodies are `null`. Headers retain all values as arrays, with credentials replaced by `[REDACTED]`. Dry-run never reads the OS keychain, but credentials supplied by flag, environment, or config file still appear redacted. The command exits successfully without contacting the API.

Local mutation commands emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object in place of a preview; filter with `select(.request)` or `select(.local)`.

### Debug

Log request and response diagnostics while running normally:

```bash
cli restore-object --api-key test_api_key --team-id 789f763f-9f96-49ae-adef-08bcc696352d --object-type identity --object-id 2aff55d3-ece9-47d3-8d69-6723ec874192 --debug
```

Debug output goes to stderr and includes:
- Request method, URL, headers, and body preview
- Response status, headers, and body preview
- Transport errors (if any)

The command still executes normally and produces its regular output on stdout.

### Flag Precedence

If both `--dry-run` and `--debug` are set, `--dry-run` takes precedence and no network calls are made.

### Security

Sensitive information is automatically redacted in diagnostic output:
- **Headers**: `Authorization`, `Cookie`, `Set-Cookie`, `X-API-Key`, and other security headers show `[REDACTED]`
- **Body**: JSON fields named `password`, `secret`, `token`, `api_key`, `client_secret`, etc. show `[REDACTED]`
- **Binary data**: binary media and canonical base64 strings are replaced with `<bytes:N>`
- **URL query**: credential-like query parameters are replaced with `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This CLI is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=cli)
