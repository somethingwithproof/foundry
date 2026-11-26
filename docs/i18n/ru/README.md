> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Профильный CI-раннер, который преобразует YAML-конфигурацию в план выполнения с учётом порядка зависимостей. Один и тот же CLI `anvil` позволяет просматривать план и выполнять команды как локально, так и внутри задачи CI.

## Архитектура

- Конфигурация (`internal/config/config.go`) задаёт профили, шаги, зависимости, переменные окружения, таймауты и повторные попытки.
- Планирование (`internal/plan`) проверяет порядок зависимостей и формирует план выполнения.
- Выполнение (`internal/exec`) запускает шаги с ограниченным параллелизмом и сохраняет журналы и результаты.
- Политика (`internal/policy/policy.go`) отклоняет шаги `script`, если они явно не разрешены. Эта политика не является границей изоляции для произвольных команд оболочки.

## Быстрый старт

В `go.mod` указана версия Go 1.23.6 или новее. Из этой рабочей копии:

```bash
go build -o bin/anvil ./cmd/anvil
```

Создайте файл `.foundry.yaml`:

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

Запустите CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Команда | Параметры |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Локализация

`anvil` включает встроенные переводы для 45 локалей (20 локалей Tier 1, 15 африканских и 10 азиатских). Выбор локали учитывает `LC_ALL`, `LC_MESSAGES` и `LANG` и может быть переопределён с помощью `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Разработка

```bash
make build
make test
make lint
make vet
```

## Лицензия

[Apache License 2.0](../../../LICENSE).
