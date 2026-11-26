> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

YAML 구성을 의존성 순서에 따른 실행 계획으로 변환하는 프로필 기반 CI 러너입니다. 동일한 `anvil` CLI로 로컬 환경이나 CI 작업 내부에서 계획을 검사하고 명령을 실행할 수 있습니다.

## 설계

- 구성(`internal/config/config.go`)은 프로필, 단계, 의존성, 환경 변수, 타임아웃 및 재시도를 정의합니다.
- 계획(`internal/plan`)은 의존성 순서를 검증하고 실행 계획을 생성합니다.
- 실행(`internal/exec`)은 제한된 동시성으로 단계를 실행하고 로그와 결과를 기록합니다.
- 정책(`internal/policy/policy.go`)은 별도로 활성화하지 않는 한 `script` 단계를 거부합니다. 이 정책은 임의의 셸 명령에 대한 격리 경계가 아닙니다.

## 빠른 시작

`go.mod`에 Go 1.23.6 이상이 선언되어 있습니다. 현재 체크아웃에서 다음을 실행합니다:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` 파일을 생성합니다:

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

CLI를 실행합니다:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| 명령 | 옵션 |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## 현지화

`anvil`은 45개 로캘(Tier 1 20개, 아프리카 15개, 아시아 10개)의 번역을 내장하고 있습니다. 로캘 선택은 `LC_ALL`, `LC_MESSAGES`, `LANG`을 따르며 `--lang`으로 재정의할 수 있습니다:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## 개발

```bash
make build
make test
make lint
make vet
```

## 라이선스

[Apache License 2.0](../../../LICENSE).
