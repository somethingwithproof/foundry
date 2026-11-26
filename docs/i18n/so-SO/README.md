> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Shatiga](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Orodaha CI oo ku saleysan astaamaha kaas oo qaabeynta YAML u beddela qorshe fulineed oo ku xiran kala horreynta ku-tiirsanaanta. Isla CLI-ga `anvil` wuxuu baari karaa qorshaha wuxuuna fulin karaa amarada maxalli ahaan ama shaqo CI gudaheeda.

## Naqshadda

- Qaabeynta (`internal/config/config.go`) waxay qeexdaa astaamaha, tallaabooyinka, ku-tiirsanaanta, doorsoomayaasha deegaanka, waqti-dhammaadka, iyo isku-dayga soo celinta.
- Qorsheynta (`internal/plan`) waxay xaqiijisaa kala horreynta ku-tiirsanaanta waxayna soo saartaa qorshaha fulinta.
- Fulinta (`internal/exec`) waxay ku socodsiisaa tallaabooyinka isbarbar-socod xaddidan waxayna diiwaangelisaa diiwaannada iyo natiijooyinka.
- Siyaasadda (`internal/policy/policy.go`) way diiddaa tallaabooyinka `script` haddii aan la daarin. Siyaasaddani ma aha xuduud go'doomin u ah amarada shell-ka ee aan la xaddidin.

## Bilow degdeg ah

Go 1.23.6 ama nooc ka cusub ayaa lagu caddeeyay `go.mod`. Gudaha nuqulkan:

```bash
go build -o bin/anvil ./cmd/anvil
```

Abuur `.foundry.yaml`:

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

Socodsii CLI-ga:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Amarka | Xulashooyinka |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Deegaameynta

`anvil` wuxuu ku dhex hayaa tarjumaado loogu talagalay 45 deegaan (20 Heerka 1-aad, 15 Afrikaan, iyo 10 Aasiyaan ah). Doorashada deegaanku waxay tixgelisaa `LC_ALL`, `LC_MESSAGES`, iyo `LANG` waxaana lagu beddeli karaa `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Horumarinta

```bash
make build
make test
make lint
make vet
```

## Shatiga

[Apache License 2.0](../../../LICENSE).
