> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licencja](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Oparty na profilach runner CI, który przekształca konfigurację YAML w plan wykonania uporządkowany według zależności. To samo narzędzie CLI `anvil` może analizować plan i wykonywać polecenia lokalnie lub w ramach zadania CI.

## Architektura

- Konfiguracja (`internal/config/config.go`) definiuje profile, kroki, zależności, zmienne środowiskowe, limity czasu i ponowne próby.
- Planowanie (`internal/plan`) weryfikuje kolejność zależności i generuje plan wykonania.
- Wykonywanie (`internal/exec`) uruchamia kroki z ograniczoną współbieżnością oraz rejestruje logi i wyniki.
- Polityka (`internal/policy/policy.go`) odrzuca kroki typu `script`, jeśli nie są włączone. Polityka ta nie stanowi granicy izolacji dla dowolnych poleceń powłoki.

## Szybki start

W pliku `go.mod` zadeklarowano wersję Go 1.23.6 lub nowszą. W katalogu roboczym repozytorium:

```bash
go build -o bin/anvil ./cmd/anvil
```

Utwórz plik `.foundry.yaml`:

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

Uruchom CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Polecenie | Opcje |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalizacja

Narzędzie `anvil` zawiera wbudowane tłumaczenia dla 45 ustawień regionalnych (20 z grupy Tier 1, 15 afrykańskich i 10 azjatyckich). Wybór języka uwzględnia zmienne `LC_ALL`, `LC_MESSAGES` i `LANG` oraz można go nadpisać flagą `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Rozwój

```bash
make build
make test
make lint
make vet
```

## Licencja

[Apache License 2.0](../../../LICENSE).
