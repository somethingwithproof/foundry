> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Igikoresho cya CI gishingiye ku myirondoro gihindura igenamiterere rya YAML mo gahunda y'ishyirwa mu bikorwa ikurikiza uko ibice bikenerana. Uburyo bumwe bwa CLI `anvil` bushobora kugenzura iyo gahunda no gukoresha amabwiriza kuri mudasobwa yawe cyangwa mu murimo wa CI.

## Imiterere

- Igenamiterere (`internal/config/config.go`) risobanura imyirondoro, intambwe, ibikenerana, impinduragaciro z'ibidukikije, igihe ntarengwa, n'isubiramo.
- Igenamigambi (`internal/plan`) rigenzura ikurikirana ry'ibikenerana kandi rigatanga gahunda y'ishyirwa mu bikorwa.
- Ishyirwa mu bikorwa (`internal/exec`) rikoresha intambwe mu buryo bw'icyarimwe bufite urugero kandi rikabika inyandiko z'ibikorwa n'ibisubizo.
- Politiki (`internal/policy/policy.go`) yanga intambwe za `script` keretse zemejwe. Iyi politiki ntabwo ari urubibi rwo gutandukanya amabwiriza ya shell ayo ari yo yose.

## Gutangira vuba

Go 1.23.6 cyangwa igezweho yashyizwe muri `go.mod`. Uhereye kuri iyi kopi:

```bash
go build -o bin/anvil ./cmd/anvil
```

Kora `.foundry.yaml`:

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

Koresha CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Ibwiriza | Amahitamo |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Ishyirwa mu ndimi

`anvil` irimo ubusobanuro bw'indimi 45 (indimi 20 z'icyiciro cya 1, 15 zo muri Afurika, na 10 zo muri Aziya). Ihitamo ry'ururimi rikurikiza `LC_ALL`, `LC_MESSAGES`, na `LANG` kandi rishobora guhindurwa hakoreshejwe `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Iterambere

```bash
make build
make test
make lint
make vet
```

## Uruhushya

[Apache License 2.0](../../../LICENSE).
