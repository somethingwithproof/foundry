> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Ilayisensi](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Umkhupheli we-CI osekelwe kwiprofayile oguqula ubumbeko lwe-YAML lube sisicwangciso sokuqhuba esilandelelaniswe ngokuxhomekeka. Kwangale `anvil` CLI inye inokuhlola isicwangciso kwaye iqhube imiyalelo kwikhompyutha yakho okanye ngaphakathi komsebenzi we-CI.

## Uyilo

- Ubumbeko (`internal/config/config.go`) luchaza iiprofayile, amanyathelo, ukuxhomekeka, izinto eziguquguqukayo zemo engqongileyo, ixesha lokuphela, kunye nokuzama kwakhona.
- Ukucwangcisa (`internal/plan`) kuqinisekisa ulandelelwano lokuxhomekeka kwaye kuvelise isicwangciso sokuqhuba.
- Ukuqhuba (`internal/exec`) kuqhuba amanyathelo ngokusebenzisa ngaxeshanye okulinganiselweyo kwaye kurekhoda iilogi kunye neziphumo.
- Umgaqo-nkqubo (`internal/policy/policy.go`) uyawala amanyathelo e-`script` ngaphandle kokuba avuliwe. Lo mgaqo-nkqubo awungomda wokwahlula imiyalelo ye-shell engakhethiyo.

## Ukuqala ngokukhawuleza

I-Go 1.23.6 okanye entsha ibhengezwe kwi-`go.mod`. Ukusuka kule khopi ilayishiweyo:

```bash
go build -o bin/anvil ./cmd/anvil
```

Yenza i-`.foundry.yaml`:

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

Qhuba i-CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Umyalelo | Iinketho |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Ulungiselelo lweelwimi

I-`anvil` ifakela iinguqulelo zeelwimi zasekuhlaleni ezingama-45 (ezingama-20 ze-Tier 1, ezili-15 zase-Afrika, nezili-10 zase-Asia). Ukukhethwa kolwimi lwasekuhlaleni kuhlonipha i-`LC_ALL`, `LC_MESSAGES`, ne-`LANG` kwaye kunokubekwa ecaleni nge-`--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Uphuhliso

```bash
make build
make test
make lint
make vet
```

## Ilayisensi

[Apache License 2.0](../../../LICENSE).
