> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

একটি প্রোফাইল-ভিত্তিক CI রানার যা YAML কনফিগারেশনকে নির্ভরতা-ক্রমানুসারে সাজানো এক্সিকিউশন প্ল্যানে রূপান্তরিত করে। একই `anvil` CLI প্ল্যানটি পরীক্ষা করতে পারে এবং স্থানীয়ভাবে বা কোনো CI জবের ভেতরে কমান্ড চালাতে পারে।

## নকশা

- কনফিগারেশন (`internal/config/config.go`) প্রোফাইল, ধাপ, নির্ভরতা, এনভায়রনমেন্ট ভেরিয়েবল, টাইমআউট এবং রিট্রাই নির্ধারণ করে।
- প্ল্যানিং (`internal/plan`) নির্ভরতার ক্রম যাচাই করে এবং এক্সিকিউশন প্ল্যান তৈরি করে।
- এক্সিকিউশন (`internal/exec`) সীমিত কনকারেন্সিসহ ধাপগুলো চালায় এবং লগ ও ফলাফল সংরক্ষণ করে।
- পলিসি (`internal/policy/policy.go`) সক্রিয় করা না থাকলে `script` ধাপগুলো প্রত্যাখ্যান করে। এই পলিসিটি যেকোনো শেল কমান্ডের জন্য কোনো আইসোলেশন সীমানা নয়।

## দ্রুত শুরু

`go.mod`-এ Go 1.23.6 বা তার পরবর্তী সংস্করণ নির্ধারিত রয়েছে। এই চেকআউট থেকে:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` তৈরি করুন:

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

CLI চালান:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| কমান্ড | বিকল্পসমূহ |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## স্থানীয়করণ

`anvil` ৪৫টি লোকেলের (২০টি Tier 1, ১৫টি আফ্রিকান এবং ১০টি এশীয় লোকেল) অনুবাদ এম্বেড করে। লোকেল নির্বাচন `LC_ALL`, `LC_MESSAGES` এবং `LANG` অনুসরণ করে এবং `--lang` দিয়ে এটি ওভাররাইড করা যায়:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## উন্নয়ন

```bash
make build
make test
make lint
make vet
```

## লাইসেন্স

[Apache License 2.0](../../../LICENSE)।
