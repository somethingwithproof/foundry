> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![مجوز](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

یک اجراکننده CI مبتنی بر پروفایل که پیکربندی YAML را به یک برنامه اجرایی مرتب‌شده بر اساس وابستگی تبدیل می‌کند. همان ابزار خط فرمان `anvil` می‌تواند برنامه را بررسی کرده و دستورها را به‌صورت محلی یا در یک کار CI اجرا کند.

## طراحی

- پیکربندی (`internal/config/config.go`) پروفایل‌ها، مراحل، وابستگی‌ها، متغیرهای محیطی، زمان‌های انتظار و تلاش‌های مجدد را تعریف می‌کند.
- برنامه‌ریزی (`internal/plan`) ترتیب وابستگی‌ها را اعتبارسنجی کرده و برنامه اجرایی را تولید می‌کند.
- اجرا (`internal/exec`) مراحل را با هم‌زمانی محدود اجرا کرده و لاگ‌ها و نتایج را ثبت می‌کند.
- سیاست (`internal/policy/policy.go`) مراحل `script` را مگر در صورت فعال بودن رد می‌کند. این سیاست یک مرز ایزوله‌سازی برای دستورهای دلخواه شل محسوب نمی‌شود.

## شروع سریع

نسخه Go 1.23.6 یا جدیدتر در `go.mod` تعریف شده است. از همین مخزن:

```bash
go build -o bin/anvil ./cmd/anvil
```

فایل `.foundry.yaml` را بسازید:

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

ابزار خط فرمان را اجرا کنید:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## رابط خط فرمان (CLI)

| دستور | گزینه‌ها |
| --- | --- |
| `doctor` | `--config`، `--lang` |
| `plan` | `--config`، `--profile`، `--json`، `--lang` |
| `run` | `--config`، `--profile`، `--jobs`، `--json`، `--lang` |
| `version` | `--json`، `--lang` |

## بومی‌سازی

`anvil` ترجمه‌های ۴۵ زبان محلی (۲۰ زبان رده اول، ۱۵ زبان آفریقایی و ۱۰ زبان آسیایی) را در خود جای داده است. انتخاب زبان از `LC_ALL`، `LC_MESSAGES` و `LANG` پیروی می‌کند و با `--lang` قابل بازنویسی است:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## توسعه

```bash
make build
make test
make lint
make vet
```

## مجوز

[Apache License 2.0](../../../LICENSE).
