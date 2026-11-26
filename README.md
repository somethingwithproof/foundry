# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=somethingwithproof_foundry&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=somethingwithproof_foundry)
[![Go minimum](https://img.shields.io/badge/Go_minimum-1.23.6-blue)](./go.mod)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](./LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/somethingwithproof/foundry/badge)](https://scorecard.dev/viewer/?uri=github.com/somethingwithproof/foundry)

**Languages:** **English** · [Deutsch](docs/i18n/de/README.md) · [Español](docs/i18n/es/README.md) · [Français](docs/i18n/fr/README.md) · [Português (BR)](docs/i18n/pt-BR/README.md) · [日本語](docs/i18n/ja/README.md) · [한국어](docs/i18n/ko/README.md) · [简体中文](docs/i18n/zh-CN/README.md) · [العربية](docs/i18n/ar/README.md) · [हिन्दी](docs/i18n/hi/README.md)

<details>
<summary>All 48 languages</summary>

**Tier 1:** **English** · [Deutsch](docs/i18n/de/README.md) · [Español](docs/i18n/es/README.md) · [Français](docs/i18n/fr/README.md) · [Italiano](docs/i18n/it/README.md) · [Português (BR)](docs/i18n/pt-BR/README.md) · [Nederlands](docs/i18n/nl/README.md) · [Русский](docs/i18n/ru/README.md) · [Українська](docs/i18n/uk/README.md) · [日本語](docs/i18n/ja/README.md) · [한국어](docs/i18n/ko/README.md) · [简体中文](docs/i18n/zh-CN/README.md) · [繁體中文](docs/i18n/zh-TW/README.md) · [العربية](docs/i18n/ar/README.md) · [हिन्दी](docs/i18n/hi/README.md) · [Polski](docs/i18n/pl/README.md) · [Türkçe](docs/i18n/tr/README.md) · [Svenska](docs/i18n/sv/README.md) · [Tiếng Việt](docs/i18n/vi/README.md) · [Bahasa Indonesia](docs/i18n/id/README.md)

**African:** [Afrikaans (ZA)](docs/i18n/af-ZA/README.md) · [አማርኛ (ET)](docs/i18n/am-ET/README.md) · [العربية (مصر)](docs/i18n/ar-EG/README.md) · [العربية (المغرب)](docs/i18n/ar-MA/README.md) · [Français (MA)](docs/i18n/fr-MA/README.md) · [Français (SN)](docs/i18n/fr-SN/README.md) · [Hausa (NG)](docs/i18n/ha-NG/README.md) · [Igbo (NG)](docs/i18n/ig-NG/README.md) · [Malagasy (MG)](docs/i18n/mg-MG/README.md) · [Kinyarwanda (RW)](docs/i18n/rw-RW/README.md) · [Soomaali (SO)](docs/i18n/so-SO/README.md) · [Kiswahili (KE)](docs/i18n/sw-KE/README.md) · [isiXhosa (ZA)](docs/i18n/xh-ZA/README.md) · [Yorùbá (NG)](docs/i18n/yo-NG/README.md) · [isiZulu (ZA)](docs/i18n/zu-ZA/README.md)

**Asian:** [বাংলা](docs/i18n/bn/README.md) · [فارسی](docs/i18n/fa/README.md) · [עברית](docs/i18n/he/README.md) · [मराठी](docs/i18n/mr/README.md) · [தமிழ்](docs/i18n/ta/README.md) · [తెలుగు](docs/i18n/te/README.md) · [ไทย](docs/i18n/th/README.md) · [اردو](docs/i18n/ur/README.md) · [Bahasa Melayu](docs/i18n/ms/README.md) · [Filipino](docs/i18n/fil/README.md)

**South Slavic:** [Bosanski](docs/i18n/bs/README.md) · [Hrvatski](docs/i18n/hr/README.md) · [Српски](docs/i18n/sr/README.md)

</details>

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
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

The default configuration is `.foundry.yaml`, the default profile is `default`, and run concurrency defaults to four jobs. Commands execute with the caller's privileges. Review configuration from untrusted sources before running it.

## Localization

`anvil` ships with embedded translations for 48 locales: 20 Tier 1 (`en`, `de`, `es`, `fr`, `it`, `pt-BR`, `nl`, `ru`, `uk`, `ja`, `ko`, `zh-CN`, `zh-TW`, `ar`, `hi`, `pl`, `tr`, `sv`, `vi`, `id`), 15 African (`af-ZA`, `am-ET`, `ar-EG`, `ar-MA`, `fr-MA`, `fr-SN`, `ha-NG`, `ig-NG`, `mg-MG`, `rw-RW`, `so-SO`, `sw-KE`, `xh-ZA`, `yo-NG`, `zu-ZA`), 10 Asian (`bn`, `fa`, `he`, `mr`, `ta`, `te`, `th`, `ur`, `ms`, `fil`), and 3 South Slavic (`bs`, `hr`, `sr`), with RTL support for `ar`, `ar-EG`, `ar-MA`, `fa`, `he`, and `ur`. Localized READMEs for all 47 non-English locales live under `docs/i18n/<locale>/README.md`.

Locale selection follows standard POSIX precedence (`LC_ALL`, then `LC_MESSAGES`, then `LANG`) and can be overridden with `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

Unset, `C`, `POSIX`, and unrecognized locales default to English.

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
