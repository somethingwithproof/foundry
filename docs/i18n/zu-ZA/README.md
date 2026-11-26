> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Isigijimi se-CI esisekelwe kumaphrofayela esiguqula ukulungiselelwa kwe-YAML kube uhlelo lokusebenza oluhlelwe ngokuncikana. I-CLI efanayo ye-`anvil` ingahlola uhlelo futhi yenze imiyalo kukhompyutha yakho noma ngaphakathi komsebenzi we-CI.

## Ukuklanywa

- Ukulungiselelwa (`internal/config/config.go`) kuchaza amaphrofayela, izinyathelo, okuncikana, iziguquguquki zendawo yokusebenza, izikhathi zokuphelelwa yisikhathi, nokuzama kabusha.
- Ukuhlela (`internal/plan`) kuqinisekisa ukulandelana kokuncikana futhi kukhiqiza uhlelo lokusebenza.
- Ukusebenzisa (`internal/exec`) kuqhuba izinyathelo ngokusebenza ngesikhathi esisodwa okulinganiselwe futhi kubhala amalogi nemiphumela.
- Inqubomgomo (`internal/policy/policy.go`) iyenqaba izinyathelo ze-`script` ngaphandle kokuthi zinikwe amandla. Le nqubomgomo ayiwona umngcele wokuhlukanisa wemiyalo ye-shell engalawuleki.

## Ukuqala ngokushesha

I-Go 1.23.6 noma entsha imenyezelwe ku-`go.mod`. Kusuka kule khophi yekhodi:

```bash
go build -o bin/anvil ./cmd/anvil
```

Dala i-`.foundry.yaml`:

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

Qalisa i-CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Umyalo | Izinketho |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Ukwenziwa ngokwendawo

I-`anvil` ifaka ukuhunyushwa kwezilimi zendawo ezingu-45 (ezingu-20 ze-Tier 1, ezingu-15 zase-Afrika, kanye nezingu-10 zase-Asia). Ukukhethwa kolimi lwendawo kuhlonipha i-`LC_ALL`, i-`LC_MESSAGES`, ne-`LANG` futhi kungashintshwa nge-`--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Ukuthuthukiswa

```bash
make build
make test
make lint
make vet
```

## Ilayisense

[Apache License 2.0](../../../LICENSE).
