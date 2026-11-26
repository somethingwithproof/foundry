> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Лиценца](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

CI покретач заснован на профилима који претвара YAML подешавања у план извршавања уређен према зависностима. Исти CLI алат `anvil` може да прегледа план и извршава команде локално или унутар CI посла.

## Дизајн

- Подешавања (`internal/config/config.go`) дефинишу профиле, кораке, зависности, променљиве окружења, временска ограничења и поновне покушаје.
- Планирање (`internal/plan`) проверава редослед зависности и генерише план извршавања.
- Извршавање (`internal/exec`) покреће кораке уз ограничену паралелност и бележи логове и резултате.
- Политика (`internal/policy/policy.go`) одбија `script` кораке осим ако нису омогућени. Ова политика не представља границу изолације за произвољне команде љуске.

## Брзи почетак

Верзија Go 1.23.6 или новија наведена је у датотеци `go.mod`. Из овог радног директоријума:

```bash
go build -o bin/anvil ./cmd/anvil
```

Направите датотеку `.foundry.yaml`:

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

Покрените CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Команда | Опције |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Локализација

Алат `anvil` садржи уграђене преводе за 48 језичких ознака (20 Tier 1, 15 афричких, 10 азијских и 3 јужнословенске). Избор језика прати променљиве `LC_ALL`, `LC_MESSAGES` и `LANG`, а може се променити заставицом `--lang`:

```bash
bin/anvil --lang de doctor
LANG=ja_JP.UTF-8 bin/anvil version
```

## Развој

```bash
make build
make test
make lint
make vet
```

## Лиценца

[Apache License 2.0](../../../LICENSE).
