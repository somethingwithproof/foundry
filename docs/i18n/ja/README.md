> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

YAML 設定を依存関係順の実行計画に変換するプロファイルベースの CI ランナーです。同じ `anvil` CLI を使って、ローカル環境や CI ジョブ内で実行計画の確認とコマンドの実行を行えます。

## 設計

- 設定 (`internal/config/config.go`) では、プロファイル、ステップ、依存関係、環境変数、タイムアウト、リトライを定義します。
- 計画 (`internal/plan`) では、依存関係の順序を検証して実行計画を生成します。
- 実行 (`internal/exec`) では、上限付きの並列数でステップを実行し、ログと結果を記録します。
- ポリシー (`internal/policy/policy.go`) では、有効化されていない限り `script` ステップを拒否します。このポリシーは任意のシェルコマンドに対する分離境界ではありません。

## クイックスタート

`go.mod` には Go 1.23.6 以上が指定されています。このチェックアウトから以下を実行します。

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` を作成します。

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

CLI を実行します。

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| コマンド | オプション |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## ローカライズ

`anvil` には 45 ロケール（Tier 1 が 20、アフリカ地域が 15、アジア地域が 10）の翻訳が埋め込まれています。ロケール選択は `LC_ALL`、`LC_MESSAGES`、`LANG` を参照し、`--lang` で上書きできます。

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## 開発

```bash
make build
make test
make lint
make vet
```

## ライセンス

[Apache License 2.0](../../../LICENSE)。
