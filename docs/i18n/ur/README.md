> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

ایک پروفائل پر مبنی CI رنر جو YAML کنفیگریشن کو انحصار کے مطابق ترتیب دیے گئے ایگزیکیوشن پلان میں تبدیل کرتا ہے۔ یہی `anvil` CLI اس پلان کا معائنہ کر سکتا ہے اور کمانڈز کو مقامی طور پر یا کسی CI جاب کے اندر چلا سکتا ہے۔

## ڈیزائن

- کنفیگریشن (`internal/config/config.go`) پروفائلز، مراحل، انحصارات، انوائرنمنٹ ویری ایبلز، ٹائم آؤٹس اور دوبارہ کوششوں کی تعریف کرتی ہے۔
- پلاننگ (`internal/plan`) انحصار کی ترتیب کی توثیق کرتی ہے اور ایگزیکیوشن پلان تیار کرتی ہے۔
- ایگزیکیوشن (`internal/exec`) محدود ہم آہنگی کے ساتھ مراحل چلاتی ہے اور لاگز اور نتائج ریکارڈ کرتی ہے۔
- پالیسی (`internal/policy/policy.go`) `script` مراحل کو مسترد کرتی ہے جب تک کہ انہیں فعال نہ کیا جائے۔ یہ پالیسی من مانی شیل کمانڈز کے لیے کوئی آئسولیشن باؤنڈری نہیں ہے۔

## فوری آغاز

`go.mod` میں Go 1.23.6 یا اس سے نیا ورژن مقرر ہے۔ اس چیک آؤٹ سے:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` بنائیں:

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

CLI چلائیں:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| کمانڈ | آپشنز |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## لوکلائزیشن

`anvil` 45 لوکیلز (20 ٹیئر 1، 15 افریقی، اور 10 ایشیائی لوکیلز) کے لیے تراجم شامل کرتا ہے۔ لوکیل کا انتخاب `LC_ALL`، `LC_MESSAGES` اور `LANG` کا احترام کرتا ہے اور `--lang` کے ذریعے اسے اوور رائیڈ کیا جا سکتا ہے:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## ڈیولپمنٹ

```bash
make build
make test
make lint
make vet
```

## لائسنس

[Apache License 2.0](../../../LICENSE)۔
