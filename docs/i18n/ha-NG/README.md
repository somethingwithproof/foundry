> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lasisi](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Mai gudanar da CI mai amfani da `profile` wanda ke mayar da saitin YAML zuwa tsarin aiwatarwa bisa jerin dogaro. Shi kansa `anvil` CLI yana iya duba tsarin kuma ya aiwatar da umarni a kan na'urar gida ko a cikin aikin CI.

## Tsari

- Saiti (`internal/config/config.go`) yana fayyace `profiles`, matakai, dogaro, masu canjin muhalli, lokacin ƙarewa, da sake gwadawa.
- Tsarawa (`internal/plan`) yana tantance jerin dogaro kuma ya samar da tsarin aiwatarwa.
- Aiwatarwa (`internal/exec`) yana gudanar da matakai tare da iyakantaccen aiki a lokaci guda kuma ya adana rajistan ayyuka da sakamako.
- Manufa (`internal/policy/policy.go`) tana ƙin matakan `script` sai dai idan an kunna su. Wannan manufa ba iyakar keɓewa ba ce ga kowane irin umarnin shell.

## Fara Amfani da Sauri

An ayyana Go 1.23.6 ko sabo a cikin `go.mod`. Daga wannan kwafin lambar:

```bash
go build -o bin/anvil ./cmd/anvil
```

Ƙirƙiri `.foundry.yaml`:

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

Gudanar da CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Umarni | Zaɓuɓɓuka |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Daidaita Harshe

`anvil` yana ɗauke da fassarori don harsuna 45 (20 na Mataki na 1, 15 na Afirka, da 10 na Asiya). Zaɓin harshe yana mutunta `LC_ALL`, `LC_MESSAGES`, da `LANG` kuma ana iya canza shi da `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Haɓakawa

```bash
make build
make test
make lint
make vet
```

## Lasisi

[Lasisi na Apache 2.0](../../../LICENSE).
