> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Mpampandeha CI mifototra amin'ny profil izay mamadika ny fikirakirana YAML ho drafitra fanatanterahana milamina araka ny fiankinan-doha. Io CLI `anvil` io ihany koa dia afaka mandinika ny drafitra sy manatanteraka baiko eo an-toerana na ao anatin'ny asa CI.

## Famolavolana

- Ny fikirakirana (`internal/config/config.go`) dia mamaritra ny profil, dingana, fiankinan-doha, mari-ova tontolo iainana, fe-potoana, ary fanandramana indray.
- Ny fandrafetana (`internal/plan`) dia manamarina ny filaminan'ny fiankinan-doha ary mamokatra ny drafitra fanatanterahana.
- Ny fanatanterahana (`internal/exec`) dia mandefa ireo dingana amin'ny fampandehanana miaraka voafetra ary mirakitra ny tatitra sy ny vokatra.
- Ny fitsipika (`internal/policy/policy.go`) dia mandà ny dingana `script` raha tsy hoe navela. Tsy sisintany fanavahana ho an'ny baiko shell an-tsitrapo ity fitsipika ity.

## Fanombohana haingana

Go 1.23.6 na vaovao kokoa no voalaza ao amin'ny `go.mod`. Avy amin'ity tahiry ity:

```bash
go build -o bin/anvil ./cmd/anvil
```

Mamorona `.foundry.yaml`:

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

Alefaso ny CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Baiko | Safidy |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Fampifanarahana amin'ny fiteny

`anvil` dia ahitana fandikan-teny ho an'ny fiteny eo an-toerana miisa 45 (fiteny sokajy 1 miisa 20, fiteny afrikanina miisa 15, ary fiteny aziatika miisa 10). Ny fisafidianana ny fiteny dia manaja ny `LC_ALL`, `LC_MESSAGES`, ary `LANG` ary azo soloina amin'ny alalan'ny `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Fampandrosoana

```bash
make build
make test
make lint
make vet
```

## Lisansa

[Apache License 2.0](../../../LICENSE).
