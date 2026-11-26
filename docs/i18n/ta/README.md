> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

YAML உள்ளமைவைச் சார்பு வரிசைப்படுத்தப்பட்ட செயல்பாட்டுத் திட்டமாக மாற்றும் சுயவிவர அடிப்படையிலான CI இயக்கி. இதே `anvil` CLI திட்டத்தை ஆய்வு செய்யவும், உள்ளூரிலோ அல்லது CI பணிக்குள்ளோ கட்டளைகளை இயக்கவும் முடியும்.

## வடிவமைப்பு

- உள்ளமைவு (`internal/config/config.go`) சுயவிவரங்கள், படிகள், சார்புகள், சூழல் மாறிகள், காலாவதி நேரங்கள் மற்றும் மறுமுயற்சிகளை வரையறுக்கிறது.
- திட்டமிடல் (`internal/plan`) சார்பு வரிசையைச் சரிபார்த்து செயல்பாட்டுத் திட்டத்தை உருவாக்குகிறது.
- செயல்பாடு (`internal/exec`) வரையறுக்கப்பட்ட ஒருங்கிணைப்புடன் படிகளை இயக்கி பதிவுகள் மற்றும் முடிவுகளைப் பதிவு செய்கிறது.
- கொள்கை (`internal/policy/policy.go`) செயல்படுத்தப்படாதவரை `script` படிகளை நிராகரிக்கிறது. தன்னிச்சையான shell கட்டளைகளுக்கு இந்தக் கொள்கை ஒரு தனிமைப்படுத்தல் எல்லை அல்ல.

## விரைவுத் தொடக்கம்

`go.mod` கோப்பில் Go 1.23.6 அல்லது அதற்குப் பிந்தைய பதிப்பு அறிவிக்கப்பட்டுள்ளது. இந்தப் பிரதியிலிருந்து:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` கோப்பை உருவாக்கவும்:

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

CLI ஐ இயக்கவும்:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| கட்டளை | விருப்பங்கள் |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## உள்ளூர்மயமாக்கல்

45 மொழியிடங்களுக்கான (20 Tier 1, 15 ஆப்பிரிக்க மற்றும் 10 ஆசிய மொழியிடங்கள்) மொழிபெயர்ப்புகளை `anvil` உட்பொதித்துள்ளது. மொழியிடத் தேர்வு `LC_ALL`, `LC_MESSAGES` மற்றும் `LANG` ஆகியவற்றைப் பின்பற்றுகிறது; மேலும் `--lang` மூலம் இதை மாற்றலாம்:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## மேம்பாடு

```bash
make build
make test
make lint
make vet
```

## உரிமம்

[Apache License 2.0](../../../LICENSE).
