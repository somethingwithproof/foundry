> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licencija](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

CI pokretač temeljen na profilima koji pretvara YAML postavke u plan izvođenja poredan po ovisnostima. Isti CLI alat `anvil` može pregledati plan i izvoditi naredbe lokalno ili unutar CI posla.

## Oblikovanje

- Postavke (`internal/config/config.go`) određuju profile, korake, ovisnosti, varijable okoline, vremenska ograničenja i ponovne pokušaje.
- Planiranje (`internal/plan`) provjerava redoslijed ovisnosti i izrađuje plan izvođenja.
- Izvođenje (`internal/exec`) pokreće korake uz ograničenu usporednost te bilježi zapisnike i rezultate.
- Pravila (`internal/policy/policy.go`) odbacuju `script` korake ako nisu izričito uključeni. To pravilo ne predstavlja granicu izolacije za proizvoljne naredbe ljuske.

## Brzi početak

Inačica Go 1.23.6 ili novija navedena je u datoteci `go.mod`. Iz ove radne mape:

```bash
go build -o bin/anvil ./cmd/anvil
```

Izradite datoteku `.foundry.yaml`:

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

Pokrenite CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Naredba | Mogućnosti |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalizacija

Alat `anvil` uključuje ugrađene prijevode za 48 jezičnih oznaka (20 Tier 1, 15 afričkih, 10 azijskih i 3 južnoslavenske). Odabir jezika poštuje varijable `LC_ALL`, `LC_MESSAGES` i `LANG`, a može se zamijeniti zastavicom `--lang`:

```bash
bin/anvil --lang de doctor
LANG=ja_JP.UTF-8 bin/anvil version
```

## Razvoj

```bash
make build
make test
make lint
make vet
```

## Licencija

[Apache License 2.0](../../../LICENSE).
