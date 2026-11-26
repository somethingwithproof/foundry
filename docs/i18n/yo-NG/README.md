> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Ìwé-àṣẹ](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Olùṣiṣẹ́ CI tí ó dá lórí profáìlì tí ó ń yí ìṣètò YAML padà sí ètò ìṣe tí a tò ní ìbámu pẹ̀lú ìgbẹ́kẹ̀lé. CLI `anvil` kan náà lè ṣe àyẹ̀wò ètò náà kí ó sì mú àwọn àṣẹ ṣẹ lórí ẹ̀rọ rẹ tàbí nínú iṣẹ́ CI.

## Ìṣètò-àpẹẹrẹ

- Ìṣètò (`internal/config/config.go`) ń ṣàlàyé àwọn profáìlì, àwọn ìgbésẹ̀, àwọn ìgbẹ́kẹ̀lé, àwọn oníyípadà àyíká, àwọn ààlà àkókò, àti àwọn àtúngbìyànjú.
- Ètò-ìṣe (`internal/plan`) ń ṣe ìdánilójú ìtòlẹ́sẹẹsẹ ìgbẹ́kẹ̀lé, ó sì ń ṣẹ̀dá ètò ìṣe náà.
- Ìmúṣẹ (`internal/exec`) ń ṣiṣẹ́ àwọn ìgbésẹ̀ pẹ̀lú ìṣiṣẹ́pọ̀ tí ó ní ààlà, ó sì ń kọ àwọn àkọsílẹ̀ àti àwọn èsì sílẹ̀.
- Ìlànà (`internal/policy/policy.go`) ń kọ̀ láti gba àwọn ìgbésẹ̀ `script` àyàfi tí a bá ṣí i sílẹ̀. Ìlànà yìí kì í ṣe ààlà ìyàsọ́tọ̀ fún àwọn àṣẹ shell èyíkéyìí.

## Ìbẹ̀rẹ̀ kíákíá

A kéde Go 1.23.6 tàbí tuntun jù bẹ́ẹ̀ lọ nínú `go.mod`. Láti inú àdàkọ yìí:

```bash
go build -o bin/anvil ./cmd/anvil
```

Ṣẹ̀dá `.foundry.yaml`:

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

Ṣiṣẹ́ CLI náà:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Àṣẹ | Àwọn àṣàyàn |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Ìsọdèdè-agbègbè

`anvil` fi àwọn ìtumọ̀ kún inú rẹ̀ fún àwọn èdè agbègbè 45 (20 ti Ìpele 1, 15 ti Áfíríkà, àti 10 ti Éṣíà). Yíyan èdè agbègbè ń tẹ̀lé `LC_ALL`, `LC_MESSAGES`, àti `LANG`, a sì lè fi `--lang` rọ́pò rẹ̀:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Ìdàgbàsókè

```bash
make build
make test
make lint
make vet
```

## Ìwé-àṣẹ

[Apache License 2.0](../../../LICENSE).
