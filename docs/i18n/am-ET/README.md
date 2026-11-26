> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![ፈቃድ](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

የYAML ውቅርን በጥገኝነት ቅደም ተከተል ወደተዘጋጀ የአፈጻጸም እቅድ የሚቀይር በመገለጫ ላይ የተመሠረተ የCI አስፈጻሚ ነው። ያው `anvil` CLI እቅዱን መመርመር እና ትዕዛዞችን በአካባቢያዊ ማሽን ወይም በCI ስራ ውስጥ ማስፈጸም ይችላል።

## ንድፍ

- ውቅር (`internal/config/config.go`) መገለጫዎችን፣ እርምጃዎችን፣ ጥገኝነቶችን፣ የአካባቢ ተለዋዋጮችን፣ የጊዜ ገደቦችን እና ድጋሚ ሙከራዎችን ይገልጻል።
- እቅድ ማውጣት (`internal/plan`) የጥገኝነት ቅደም ተከተልን ያረጋግጣል እና የአፈጻጸም እቅዱን ያመነጫል።
- አፈጻጸም (`internal/exec`) እርምጃዎችን በተገደበ ትይዩነት ያስፈጽማል እንዲሁም መዝገቦችን እና ውጤቶችን ይመዘግባል።
- ፖሊሲ (`internal/policy/policy.go`) ካልነቃ በስተቀር የ`script` እርምጃዎችን አይቀበልም። ይህ ፖሊሲ ለማንኛውም የshell ትዕዛዞች የመለያ ድንበር አይደለም።

## ፈጣን ጅምር

Go 1.23.6 ወይም ከዚያ በላይ በ`go.mod` ውስጥ ተገልጿል። ከዚህ ቅጂ ውስጥ፦

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` ይፍጠሩ፦

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

የCLI መሣሪያውን ያስኪዱ፦

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| ትዕዛዝ | አማራጮች |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## አካባቢያዊ ትርጉም

`anvil` ለ45 አካባቢዎች (20 ደረጃ 1፣ 15 የአፍሪካ እና 10 የእስያ አካባቢዎች) ትርጉሞችን ያካትታል። የአካባቢ ምርጫ `LC_ALL`፣ `LC_MESSAGES` እና `LANG` የተባሉትን ያከብራል እንዲሁም በ`--lang` ሊሻር ይችላል፦

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## ልማት

```bash
make build
make test
make lint
make vet
```

## ፈቃድ

[Apache License 2.0](../../../LICENSE)።
