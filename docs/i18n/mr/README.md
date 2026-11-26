> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![परवाना](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

प्रोफाइल-आधारित CI रनर जो YAML कॉन्फिगरेशनचे अवलंबित्व-क्रमानुसार अंमलबजावणी योजनेत रूपांतर करतो. तोच `anvil` CLI ही योजना तपासू शकतो आणि स्थानिक पातळीवर किंवा CI जॉबमध्ये कमांड्स कार्यान्वित करू शकतो.

## रचना

- कॉन्फिगरेशन (`internal/config/config.go`) प्रोफाइल्स, पायऱ्या, अवलंबित्वे, एन्व्हायर्नमेंट व्हेरिएबल्स, टाइमआउट आणि पुनर्प्रयत्न परिभाषित करते.
- नियोजन (`internal/plan`) अवलंबित्व क्रमाची पडताळणी करते आणि अंमलबजावणी योजना तयार करते.
- अंमलबजावणी (`internal/exec`) मर्यादित समवर्तीतेसह पायऱ्या चालवते आणि लॉग व निकाल नोंदवते.
- धोरण (`internal/policy/policy.go`) सक्षम केलेले नसल्यास `script` पायऱ्या नाकारते. हे धोरण अनियंत्रित शेल कमांड्ससाठी विलगीकरण सीमा नाही.

## जलद सुरुवात

`go.mod` मध्ये Go 1.23.6 किंवा त्याहून नवीन आवृत्ती घोषित केली आहे. या चेकआउटमधून:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` तयार करा:

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

CLI चालवा:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| कमांड | पर्याय |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## स्थानिकीकरण

`anvil` मध्ये 45 लोकॅल्सचे (20 टियर 1, 15 आफ्रिकन आणि 10 आशियाई लोकॅल्स) भाषांतर समाविष्ट आहे. लोकॅल निवड `LC_ALL`, `LC_MESSAGES` आणि `LANG` चे पालन करते आणि `--lang` वापरून बदलली जाऊ शकते:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## विकास

```bash
make build
make test
make lint
make vet
```

## परवाना

[Apache License 2.0](../../../LICENSE).
