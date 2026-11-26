> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lisensya](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Isang profile-based na CI runner na nagko-convert ng YAML configuration sa isang execution plan na nakaayos ayon sa dependency. Ang parehong `anvil` CLI ay maaaring magsuri ng plano at magpatakbo ng mga command nang lokal o sa loob ng isang CI job.

## Disenyo

- Tinutukoy ng configuration (`internal/config/config.go`) ang mga profile, hakbang, dependency, environment variable, timeout, at retry.
- Bineberipika ng planning (`internal/plan`) ang pagkakasunod-sunod ng mga dependency at gumagawa ng execution plan.
- Pinatatakbo ng execution (`internal/exec`) ang mga hakbang na may limitadong concurrency at nagtatala ng mga log at resulta.
- Tinatanggihan ng policy (`internal/policy/policy.go`) ang mga hakbang na `script` maliban kung naka-enable. Ang policy na ito ay hindi isang isolation boundary para sa mga arbitraryong shell command.

## Mabilisang pagsisimula

Idineklara ang Go 1.23.6 o mas bago sa `go.mod`. Mula sa checkout na ito:

```bash
go build -o bin/anvil ./cmd/anvil
```

Gumawa ng `.foundry.yaml`:

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

Patakbuhin ang CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Command | Mga Opsyon |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Lokalisasyon

May naka-embed na mga salin ang `anvil` para sa 45 locale (20 Tier 1, 15 African, at 10 Asian na locale). Sinusunod ng pagpili ng locale ang `LC_ALL`, `LC_MESSAGES`, at `LANG` at maaaring i-override gamit ang `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Pag-develop

```bash
make build
make test
make lint
make vet
```

## Lisensya

[Apache License 2.0](../../../LICENSE).
