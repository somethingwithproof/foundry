> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licenca](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

CI pokretač zasnovan na profilima koji pretvara YAML postavke u plan izvršavanja uređen prema zavisnostima. Isti CLI alat `anvil` može pregledati plan i izvršavati naredbe lokalno ili unutar CI posla.

## Dizajn

- Postavke (`internal/config/config.go`) definišu profile, korake, zavisnosti, varijable okruženja, vremenska ograničenja i ponovne pokušaje.
- Planiranje (`internal/plan`) provjerava redoslijed zavisnosti i generiše plan izvršavanja.
- Izvršavanje (`internal/exec`) pokreće korake uz ograničenu paralelnost i bilježi dnevnike i rezultate.
- Pravila (`internal/policy/policy.go`) odbijaju `script` korake osim ako nisu uključeni. Ovo pravilo ne predstavlja granicu izolacije za proizvoljne naredbe ljuske.

## Brzi početak

Verzija Go 1.23.6 ili novija navedena je u datoteci `go.mod`. Iz ovog radnog direktorija:

```bash
go build -o bin/anvil ./cmd/anvil
```

Kreirajte datoteku `.foundry.yaml`:

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

| Naredba | Opcije |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalizacija

Alat `anvil` sadrži ugrađene prijevode za 48 jezičkih oznaka (20 Tier 1, 15 afričkih, 10 azijskih i 3 južnoslavenske). Izbor jezika prati varijable `LC_ALL`, `LC_MESSAGES` i `LANG`, a može se nadjačati zastavicom `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Razvoj

```bash
make build
make test
make lint
make vet
```

## Licenca

[Apache License 2.0](../../../LICENSE).
