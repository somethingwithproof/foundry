> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Licencia](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Un ejecutor de CI basado en perfiles que convierte una configuración YAML en un plan de ejecución ordenado por dependencias. La misma CLI `anvil` permite inspeccionar el plan y ejecutar comandos de forma local o dentro de un trabajo de CI.

## Diseño

- La configuración (`internal/config/config.go`) define perfiles, pasos, dependencias, variables de entorno, tiempos de espera y reintentos.
- La planificación (`internal/plan`) valida el orden de las dependencias y genera el plan de ejecución.
- La ejecución (`internal/exec`) ejecuta los pasos con concurrencia acotada y registra los registros y resultados.
- La política (`internal/policy/policy.go`) rechaza los pasos `script` a menos que estén habilitados. Esta política no constituye un límite de aislamiento para comandos de shell arbitrarios.

## Inicio rápido

En `go.mod` se declara Go 1.23.6 o superior. Desde esta copia de trabajo:

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

Ejecuta la CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Comando | Opciones |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Localización

`anvil` integra traducciones para 45 configuraciones regionales (20 de nivel 1, 15 africanas y 10 asiáticas). La selección de configuración regional respeta `LC_ALL`, `LC_MESSAGES` y `LANG`, y se puede anular con `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Desarrollo

```bash
make build
make test
make lint
make vet
```

## Licencia

[Licencia Apache 2.0](../../../LICENSE).
