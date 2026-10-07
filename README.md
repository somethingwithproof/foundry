# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](./LICENSE)
[![Go minimum](https://img.shields.io/badge/Go_minimum-1.23.6-blue)](./go.mod)

A profile-based CI runner that turns a YAML configuration into a dependency-ordered execution plan. The same `anvil` CLI can inspect the plan and execute commands locally or inside a CI job.

## Design

- [Configuration](internal/config/config.go) defines profiles, steps, dependencies, environment variables, timeouts, and retries.
- [Planning](internal/plan) validates dependency ordering and produces the execution plan.
- [Execution](internal/exec) runs steps with bounded concurrency and records logs and results.
- [Policy](internal/policy/policy.go) rejects `script` steps unless enabled. This policy is not an isolation boundary for arbitrary shell commands.

Deterministic planning means the ordering and plan representation can be reproduced for the same configuration. It does not make external commands, services, timestamps, or build artifacts deterministic.

## Quick start

Go 1.23.6 or newer is declared in [go.mod](go.mod). From this checkout:

```bash
go build -o bin/anvil ./cmd/anvil
```

Create `.foundry.yaml`:

```yaml
version: 1
project:
  name: example
profiles:
  default:
    steps:
      - id: inspect
        type: shell
        command: ["go", "version"]
      - id: test
        type: shell
        deps: ["inspect"]
        command: ["go", "test", "./..."]
```

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

`doctor` checks the configuration and Go availability. `plan` writes `.foundry/out/plan.json` without running the steps. Use `plan` for inspection; the current CLI does not implement `run --dry-run` or `--verbose`.

## CLI

| Command | Options |
| --- | --- |
| `doctor` | `--config` |
| `plan` | `--config`, `--profile`, `--json` |
| `run` | `--config`, `--profile`, `--jobs`, `--json` |
| `version` | `--json` |

The default configuration is `.foundry.yaml`, the default profile is `default`, and run concurrency defaults to four jobs. Commands execute with the caller's privileges. Review configuration from untrusted sources before running it.

## Development

```bash
make build
make test
make lint
make vet
```

See the [Makefile](Makefile) for the exact commands and [tests](internal) for configuration, planning, and execution behavior. Linting requires golangci-lint installed separately.

## License

[Apache License 2.0](LICENSE).
