> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lizenz](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Ein profilbasierter CI-Runner, der eine YAML-Konfiguration in einen nach Abhängigkeiten geordneten Ausführungsplan umwandelt. Dieselbe `anvil`-CLI kann den Plan prüfen und Befehle lokal oder in einem CI-Job ausführen.

## Aufbau

- Konfiguration (`internal/config/config.go`) definiert Profile, Schritte, Abhängigkeiten, Umgebungsvariablen, Timeouts und Wiederholungsversuche.
- Planung (`internal/plan`) validiert die Abhängigkeitsreihenfolge und erstellt den Ausführungsplan.
- Ausführung (`internal/exec`) führt Schritte mit begrenzter Parallelität aus und zeichnet Protokolle sowie Ergebnisse auf.
- Richtlinie (`internal/policy/policy.go`) lehnt `script`-Schritte ab, sofern sie nicht aktiviert sind. Diese Richtlinie stellt keine Isolationsgrenze für beliebige Shell-Befehle dar.

## Schnellstart

In `go.mod` ist Go 1.23.6 oder neuer festgelegt. Aus diesem Checkout:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` erstellen:

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

CLI ausführen:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Befehl | Optionen |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalisierung

`anvil` bettet Übersetzungen für 45 Locales (20 Tier-1-, 15 afrikanische und 10 asiatische Locales) ein. Die Locale-Auswahl berücksichtigt `LC_ALL`, `LC_MESSAGES` und `LANG` und lässt sich mit `--lang` überschreiben:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Entwicklung

```bash
make build
make test
make lint
make vet
```

## Lizenz

[Apache License 2.0](../../../LICENSE).
