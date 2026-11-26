> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![许可证](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

一个基于配置档（Profile）的 CI 运行器，可将 YAML 配置转换为按依赖顺序排列的执行计划。同一个 `anvil` CLI 工具既可用于查看执行计划，也可在本地或 CI 作业中执行命令。

## 设计

- 配置模块（`internal/config/config.go`）用于定义配置档、步骤、依赖关系、环境变量、超时时间及重试策略。
- 规划模块（`internal/plan`）用于校验依赖顺序并生成执行计划。
- 执行模块（`internal/exec`）以受控并发数运行各个步骤，并记录日志与执行结果。
- 策略模块（`internal/policy/policy.go`）默认拒绝 `script` 步骤，除非显式启用。该策略并非针对任意 Shell 命令的隔离边界。

## 快速开始

`go.mod` 中声明的最低 Go 版本为 Go 1.23.6。在当前代码检出目录中运行：

```bash
go build -o bin/anvil ./cmd/anvil
```

创建 `.foundry.yaml`：

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

运行 CLI：

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| 命令 | 选项 |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## 本地化

`anvil` 内置了 45 个语言环境的翻译（20 个 Tier 1、15 个非洲和 10 个亚洲语言环境）。语言环境选择遵循 `LC_ALL`、`LC_MESSAGES` 和 `LANG`，并可通过 `--lang` 参数覆盖：

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## 开发

```bash
make build
make test
make lint
make vet
```

## 许可证

[Apache License 2.0](../../../LICENSE)。
