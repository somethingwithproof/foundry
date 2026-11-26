> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Um executor de CI baseado em perfis que transforma uma configuração YAML em um plano de execução ordenado por dependências. A mesma CLI `anvil` pode inspecionar o plano e executar comandos localmente ou dentro de um job de CI.

## Arquitetura

- Configuração (`internal/config/config.go`) define perfis, etapas, dependências, variáveis de ambiente, timeouts e tentativas.
- Planejamento (`internal/plan`) valida a ordenação de dependências e gera o plano de execução.
- Execução (`internal/exec`) executa etapas com concorrência limitada e registra logs e resultados.
- Política (`internal/policy/policy.go`) rejeita etapas `script`, a menos que habilitadas. Essa política não constitui uma fronteira de isolamento para comandos shell arbitrários.

## Início rápido

Go 1.23.6 ou superior está declarado em `go.mod`. A partir deste checkout:

```bash
go build -o bin/anvil ./cmd/anvil
```

Crie `.foundry.yaml`:

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

Execute a CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Comando | Opções |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Localização

`anvil` incorpora traduções para 45 locales (20 Tier 1, 15 africanos e 10 asiáticos). A seleção de locale respeita `LC_ALL`, `LC_MESSAGES` e `LANG` e pode ser substituída com `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Desenvolvimento

```bash
make build
make test
make lint
make vet
```

## Licença

[Apache License 2.0](../../../LICENSE).
