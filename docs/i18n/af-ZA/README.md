> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lisensie](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

'n Profielgebaseerde CI-uitvoerder wat 'n YAML-konfigurasie in 'n afhanklikheidsgeordende uitvoeringsplan omskakel. Dieselfde `anvil`-CLI kan die plan inspekteer en opdragte plaaslik of binne 'n CI-taak uitvoer.

## Ontwerp

- Konfigurasie (`internal/config/config.go`) definieer profiele, stappe, afhanklikhede, omgewingsveranderlikes, tydsbeperkings en herpogings.
- Beplanning (`internal/plan`) valideer afhanklikheidsordening en lewer die uitvoeringsplan.
- Uitvoering (`internal/exec`) voer stappe met begrensde gelyktydigheid uit en teken logs en resultate aan.
- Beleid (`internal/policy/policy.go`) verwerp `script`-stappe tensy dit geaktiveer is. Hierdie beleid is nie 'n isolasiegrens vir arbitrêre dopopdragte nie.

## Vinnige begin

Go 1.23.6 of nuwer word in `go.mod` verklaar. Vanuit hierdie werkkopie:

```bash
go build -o bin/anvil ./cmd/anvil
```

Skep `.foundry.yaml`:

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

Voer die CLI uit:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Opdrag | Opsies |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalisering

`anvil` sluit vertalings vir 45 locales (20 Vlak 1-, 15 Afrika- en 10 Asiatiese locales) in. Locale-keuse respekteer `LC_ALL`, `LC_MESSAGES` en `LANG` en kan met `--lang` oorheers word:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Ontwikkeling

```bash
make build
make test
make lint
make vet
```

## Lisensie

[Apache License 2.0](../../../LICENSE).
