> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licens](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

En profilbaserad CI-runner som omvandlar en YAML-konfiguration till en beroendeordnad exekveringsplan. Samma `anvil`-CLI kan inspektera planen och köra kommandon lokalt eller i ett CI-jobb.

## Design

- Konfiguration (`internal/config/config.go`) definierar profiler, steg, beroenden, miljövariabler, tidsgränser och omförsök.
- Planering (`internal/plan`) validerar beroendeordningen och skapar exekveringsplanen.
- Exekvering (`internal/exec`) kör steg med begränsad samtidighet och registrerar loggar och resultat.
- Policy (`internal/policy/policy.go`) avvisar `script`-steg om de inte är aktiverade. Denna policy utgör ingen isoleringsgräns för godtyckliga skalkommandon.

## Snabbstart

Go 1.23.6 eller senare anges i `go.mod`. Från denna utcheckning:

```bash
go build -o bin/anvil ./cmd/anvil
```

Skapa `.foundry.yaml`:

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

Kör CLI-verktyget:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Kommando | Alternativ |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalisering

`anvil` bäddar in översättningar för 45 språkområden (20 Tier 1-, 15 afrikanska och 10 asiatiska språkområden). Språkvalet följer `LC_ALL`, `LC_MESSAGES` och `LANG` och kan åsidosättas med `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Utveckling

```bash
make build
make test
make lint
make vet
```

## Licens

[Apache License 2.0](../../../LICENSE).
