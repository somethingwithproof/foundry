> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

एक प्रोफ़ाइल-आधारित CI रनर जो YAML कॉन्फ़िगरेशन को निर्भरता-क्रमबद्ध निष्पादन योजना में बदलता है। यही `anvil` CLI योजना का निरीक्षण कर सकता है और स्थानीय रूप से या किसी CI जॉब के भीतर कमांड निष्पादित कर सकता है।

## डिज़ाइन

- कॉन्फ़िगरेशन (`internal/config/config.go`) प्रोफ़ाइल, चरण, निर्भरताएँ, एनवायरनमेंट वेरिएबल, टाइमआउट और पुनः प्रयास निर्धारित करता है।
- नियोजन (`internal/plan`) निर्भरता क्रम की पुष्टि करता है और निष्पादन योजना तैयार करता है।
- निष्पादन (`internal/exec`) सीमित समवर्तीता के साथ चरणों को चलाता है और लॉग तथा परिणाम रिकॉर्ड करता है।
- नीति (`internal/policy/policy.go`) `script` चरणों को तब तक अस्वीकार करती है जब तक कि उन्हें सक्षम न किया जाए। यह नीति मनमाने शेल कमांड के लिए कोई आइसोलेशन सीमा नहीं है।

## त्वरित आरंभ

`go.mod` में Go 1.23.6 या उससे नया संस्करण घोषित है। इस चेकआउट से:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` बनाएँ:

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

CLI चलाएँ:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| कमांड | विकल्प |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## स्थानीयकरण

`anvil` 45 लोकेल (20 टियर 1, 15 अफ़्रीकी और 10 एशियाई लोकेल) के लिए अनुवाद एम्बेड करता है। लोकेल चयन `LC_ALL`, `LC_MESSAGES` और `LANG` का सम्मान करता है और इसे `--lang` से ओवरराइड किया जा सकता है:

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

## लाइसेंस

[Apache License 2.0](../../../LICENSE)।
