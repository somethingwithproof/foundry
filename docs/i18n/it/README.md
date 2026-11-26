> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licenza](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Un runner CI basato su profili che trasforma una configurazione YAML in un piano di esecuzione ordinato per dipendenze. La stessa CLI `anvil` può ispezionare il piano ed eseguire comandi in locale o all'interno di un job CI.

## Progettazione

- La configurazione (`internal/config/config.go`) definisce profili, passaggi, dipendenze, variabili d'ambiente, timeout e tentativi di ripetizione.
- La pianificazione (`internal/plan`) valida l'ordinamento delle dipendenze e produce il piano di esecuzione.
- L'esecuzione (`internal/exec`) esegue i passaggi con concorrenza limitata e registra log e risultati.
- Le policy (`internal/policy/policy.go`) rifiutano i passaggi `script` se non abilitati. Questa policy non costituisce un confine di isolamento per comandi shell arbitrari.

## Avvio rapido

In `go.mod` è dichiarato Go 1.23.6 o versione successiva. Da questo checkout:

```bash
go build -o bin/anvil ./cmd/anvil
```

Crea `.foundry.yaml`:

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

Esegui la CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Comando | Opzioni |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Localizzazione

`anvil` integra le traduzioni per 45 localizzazioni (20 Tier 1, 15 africane e 10 asiatiche). La selezione della lingua rispetta `LC_ALL`, `LC_MESSAGES` e `LANG` e può essere sovrascritta con `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Sviluppo

```bash
make build
make test
make lint
make vet
```

## Licenza

[Apache License 2.0](../../../LICENSE).
