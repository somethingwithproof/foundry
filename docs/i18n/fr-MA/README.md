> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licence](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Un exécuteur CI basé sur des profils qui transforme une configuration YAML en un plan d'exécution ordonné selon les dépendances. La même CLI `anvil` permet d'inspecter le plan et d'exécuter des commandes en local ou au sein d'un job CI.

## Conception

- La configuration (`internal/config/config.go`) définit les profils, les étapes, les dépendances, les variables d'environnement, les délais d'expiration et les tentatives de réexécution.
- La planification (`internal/plan`) valide l'ordre des dépendances et génère le plan d'exécution.
- L'exécution (`internal/exec`) lance les étapes avec une concurrence bornée et enregistre les journaux et les résultats.
- La politique (`internal/policy/policy.go`) rejette les étapes `script` sauf activation explicite. Cette politique ne constitue pas une frontière d'isolation pour des commandes shell arbitraires.

## Démarrage rapide

Go 1.23.6 ou une version ultérieure est déclaré dans `go.mod`. Depuis cette copie de travail :

```bash
go build -o bin/anvil ./cmd/anvil
```

Créez `.foundry.yaml` :

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

Exécutez la CLI :

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Commande | Options |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Localisation

`anvil` embarque des traductions pour 45 locales (20 de niveau 1, 15 africaines et 10 asiatiques). La sélection de la locale respecte `LC_ALL`, `LC_MESSAGES` et `LANG` et peut être remplacée avec `--lang` :

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Développement

```bash
make build
make test
make lint
make vet
```

## Licence

[Licence Apache 2.0](../../../LICENSE).
