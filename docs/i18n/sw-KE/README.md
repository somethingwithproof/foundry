> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Mtekelezaji wa CI unaotumia wasifu ambao hugeuza usanidi wa YAML kuwa mpango wa utekelezaji uliopangwa kulingana na utegemezi. CLI hiyo hiyo ya `anvil` inaweza kukagua mpango huo na kutekeleza amri ndani ya mashine au ndani ya kazi ya CI.

## Muundo

- Usanidi (`internal/config/config.go`) hufafanua wasifu, hatua, utegemezi, vigezo vya mazingira, muda wa kuisha, na majaribio ya kurudia.
- Upangaji (`internal/plan`) huthibitisha mpangilio wa utegemezi na kuunda mpango wa utekelezaji.
- Utekelezaji (`internal/exec`) huendesha hatua kwa usambamba uliodhibitiwa na kurekodi kumbukumbu na matokeo.
- Sera (`internal/policy/policy.go`) hukataa hatua za `script` isipokuwa ziwe zimewezeshwa. Sera hii si mpaka wa utenganishaji kwa amri zozote za shell.

## Mwanzo wa haraka

Go 1.23.6 au toleo jipya zaidi limebainishwa katika `go.mod`. Kutoka kwenye nakala hii:

```bash
go build -o bin/anvil ./cmd/anvil
```

Unda `.foundry.yaml`:

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

Endesha CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Amri | Chaguo |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Ujanibishaji

`anvil` hujumuisha tafsiri za lugha na maeneo 45 (maeneo 20 ya Daraja la 1, 15 ya Afrika, na 10 ya Asia). Uteuzi wa eneo huheshimu `LC_ALL`, `LC_MESSAGES`, na `LANG` na unaweza kubatilishwa kwa kutumia `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Uendelezaji

```bash
make build
make test
make lint
make vet
```

## Leseni

[Leseni ya Apache 2.0](../../../LICENSE).
