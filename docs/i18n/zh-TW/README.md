> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

一個以設定檔為基礎的 CI 執行器，能將 YAML 設定轉換為依相依性順序排列的執行計畫。同一個 `anvil` CLI 可用於在本機或 CI 作業中檢視計畫並執行指令。

## 設計

- 設定（`internal/config/config.go`）定義設定檔、步驟、相依性、環境變數、逾時時間與重試次數。
- 規劃（`internal/plan`）驗證相依性順序並產生執行計畫。
- 執行（`internal/exec`）以受限的並行數執行步驟，並記錄日誌與結果。
- 原則（`internal/policy/policy.go`）除非明確啟用，否則會拒絕 `script` 步驟。此原則並非任意 shell 指令的隔離邊界。

## 快速開始

`go.mod` 中宣告需使用 Go 1.23.6 或更新版本。在此簽出目錄中執行：

```bash
go build -o bin/anvil ./cmd/anvil
```

建立 `.foundry.yaml`：

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

執行 CLI：

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| 指令 | 選項 |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## 在地化

`anvil` 內嵌了 45 種語系的翻譯（包含 20 種 Tier 1 語系、15 種非洲語系與 10 種亞洲語系）。語系選取會依循 `LC_ALL`、`LC_MESSAGES` 與 `LANG` 環境變數，並可使用 `--lang` 覆寫：

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## 開發

```bash
make build
make test
make lint
make vet
```

## 授權條款

[Apache License 2.0](../../../LICENSE)。
