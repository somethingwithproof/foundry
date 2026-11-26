> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licentie](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Een op profielen gebaseerde CI-runner die een YAML-configuratie omzet in een uitvoeringsplan geordend op afhankelijkheden. Dezelfde `anvil`-CLI kan het plan inspecteren en opdrachten lokaal of binnen een CI-job uitvoeren.

## Ontwerp

- Configuratie (`internal/config/config.go`) definieert profielen, stappen, afhankelijkheden, omgevingsvariabelen, time-outs en herhaalacties.
- Planning (`internal/plan`) valideert de volgorde van afhankelijkheden en genereert het uitvoeringsplan.
- Uitvoering (`internal/exec`) voert stappen uit met begrensde gelijktijdigheid en legt logboeken en resultaten vast.
- Beleid (`internal/policy/policy.go`) wijst `script`-stappen af tenzij deze zijn ingeschakeld. Dit beleid vormt geen isolatiegrens voor willekeurige shell-opdrachten.

## Snelstart

Go 1.23.6 of nieuwer is gedeclareerd in `go.mod`. Vanuit deze checkout:

```bash
go build -o bin/anvil ./cmd/anvil
```

Maak `.foundry.yaml` aan:

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

Voer de CLI uit:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Opdracht | Opties |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalisatie

`anvil` bevat ingebouwde vertalingen voor 45 locales (20 Tier 1-, 15 Afrikaanse en 10 Aziatische locales). De selectie van de locale respecteert `LC_ALL`, `LC_MESSAGES` en `LANG` en kan worden overschreven met `--lang`:

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

## Licentie

[Apache License 2.0](../../../LICENSE).
