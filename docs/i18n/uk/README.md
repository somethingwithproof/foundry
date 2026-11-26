> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Ліцензія](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Інструмент запуску CI на основі профілів, який перетворює конфігурацію YAML на план виконання, впорядкований за залежностями. Той самий інтерфейс командного рядка `anvil` дозволяє перевіряти план і виконувати команди локально або всередині завдання CI.

## Архітектура

- Конфігурація (`internal/config/config.go`) визначає профілі, кроки, залежності, змінні середовища, тайм-аути та повторні спроби.
- Планування (`internal/plan`) перевіряє порядок залежностей і формує план виконання.
- Виконання (`internal/exec`) запускає кроки з обмеженою паралельністю та записує журнали й результати.
- Політика (`internal/policy/policy.go`) відхиляє кроки `script`, якщо їх не увімкнено. Ця політика не є межею ізоляції для довільних команд оболонки.

## Швидкий старт

У `go.mod` оголошено Go 1.23.6 або новішу версію. Із цієї локальної копії:

```bash
go build -o bin/anvil ./cmd/anvil
```

Створіть `.foundry.yaml`:

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

Запустіть CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Команда | Параметри |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Локалізація

`anvil` містить вбудовані переклади для 45 локалей (20 рівня Tier 1, 15 африканських і 10 азійських локалей). Вибір локалі враховує `LC_ALL`, `LC_MESSAGES` та `LANG` і може бути перевизначений за допомогою `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Розробка

```bash
make build
make test
make lint
make vet
```

## Ліцензія

[Apache License 2.0](../../../LICENSE).
