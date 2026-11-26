> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Onye na-agba ọsọ CI dabere na profaịlụ nke na-agbanwe nhazi YAML ka ọ bụrụ atụmatụ mmezu edoziri dịka ndabere si dị. Otu `anvil` CLI ahụ nwere ike inyocha atụmatụ ahụ ma mezuo iwu n'ime kọmputa gị ma ọ bụ n'ime ọrụ CI.

## Nhazi usoro

- Nhazi (`internal/config/config.go`) na-akọwa profaịlụ, nzọụkwụ, ndabere, mgbanwe gburugburu ebe obibi, oge ngwụcha, na nnwale ọzọ.
- Nhazi atụmatụ (`internal/plan`) na-akwado usoro ndabere ma na-emepụta atụmatụ mmezu.
- Mmezu (`internal/exec`) na-agba nzọụkwụ na njikọta oge nwere oke ma na-edekọ ndekọ na nsonaazụ.
- Iwu (`internal/policy/policy.go`) na-ajụ nzọụkwụ `script` ma ọ bụrụ na agbanyeghị ya. Iwu a abụghị oke nkewa maka iwu shell ọ bụla.

## Mmalite ngwa ngwa

E kwupụtara Go 1.23.6 ma ọ bụ nke dị ọhụrụ karịa na `go.mod`. Site na nbudata a:

```bash
go build -o bin/anvil ./cmd/anvil
```

Mepụta `.foundry.yaml`:

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

Gbaa CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Iwu | Nhọrọ |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Nkwado asụsụ obodo

`anvil` tinyere ntụgharị asụsụ maka asụsụ obodo 45 (asụsụ Tier 1 20, asụsụ Afrịka 15, na asụsụ Eshia 10). Nhọrọ asụsụ obodo na-asọpụrụ `LC_ALL`, `LC_MESSAGES`, na `LANG`, enwekwara ike iji `--lang` gbanwee ya:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Mmepe

```bash
make build
make test
make lint
make vet
```

## Ikikere

[Apache License 2.0](../../../LICENSE).
