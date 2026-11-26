> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

YAML కాన్ఫిగరేషన్‌ను డిపెండెన్సీ క్రమంలో ఉండే అమలు ప్రణాళికగా మార్చే ప్రొఫైల్-ఆధారిత CI రన్నర్. అదే `anvil` CLI ఆ ప్రణాళికను పరిశీలించగలదు మరియు స్థానికంగా లేదా CI జాబ్ లోపల ఆదేశాలను అమలు చేయగలదు.

## రూపకల్పన

- కాన్ఫిగరేషన్ (`internal/config/config.go`) ప్రొఫైల్స్, దశలు, డిపెండెన్సీలు, ఎన్విరాన్‌మెంట్ వేరియబుల్స్, టైమ్‌అవుట్‌లు, మరియు రీట్రైలను నిర్వచిస్తుంది.
- ప్రణాళిక (`internal/plan`) డిపెండెన్సీ క్రమాన్ని ధృవీకరించి అమలు ప్రణాళికను రూపొందిస్తుంది.
- అమలు (`internal/exec`) పరిమిత ఏకకాలికతతో దశలను నడుపుతుంది మరియు లాగ్‌లు, ఫలితాలను నమోదు చేస్తుంది.
- విధానం (`internal/policy/policy.go`) ప్రారంభించకపోతే `script` దశలను తిరస్కరిస్తుంది. అనియంత్రిత షెల్ ఆదేశాలకు ఈ విధానం ఐసోలేషన్ సరిహద్దు కాదు.

## త్వరిత ప్రారంభం

`go.mod` లో Go 1.23.6 లేదా అంతకంటే కొత్త వెర్షన్ ప్రకటించబడింది. ఈ చెక్‌అవుట్ నుండి:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` ను సృష్టించండి:

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

CLI ని అమలు చేయండి:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| ఆదేశం | ఎంపికలు |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## స్థానికీకరణ

`anvil` 45 లోకేల్స్ కోసం అనువాదాలను (20 టియర్ 1, 15 ఆఫ్రికన్, మరియు 10 ఆసియా లోకేల్స్) పొందుపరుస్తుంది. లోకేల్ ఎంపిక `LC_ALL`, `LC_MESSAGES`, మరియు `LANG` లను అనుసరిస్తుంది మరియు `--lang` తో భర్తీ చేయవచ్చు:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## అభివృద్ధి

```bash
make build
make test
make lint
make vet
```

## లైసెన్స్

[Apache License 2.0](../../../LICENSE).
