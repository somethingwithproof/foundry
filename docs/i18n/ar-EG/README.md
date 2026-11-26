> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

مشغّل CI بيعتمد على الـ profiles وبيحوّل ملف إعدادات YAML لخطة تنفيذ مترتبة حسب الاعتماديات. نفس أداة سطر الأوامر `anvil` تقدر تفحص الخطة وتنفّذ الأوامر محلياً أو جوّا CI job.

## التصميم

- الإعدادات (`internal/config/config.go`) بتعرّف الـ profiles، الخطوات، الاعتماديات، متغيرات البيئة، الـ timeouts، ومحاولات إعادة التشغيل (retries).
- التخطيط (`internal/plan`) بيتحقق من ترتيب الاعتماديات وبيطلّع خطة التنفيذ.
- التنفيذ (`internal/exec`) بيشغّل الخطوات بتوازي محكوم (bounded concurrency) وبيسجّل الـ logs والنتائج.
- السياسة (`internal/policy/policy.go`) بترفض خطوات `script` إلا لو اتفعّلت. السياسة دي مش حاجز عزل (isolation boundary) لأوامر الـ shell العشوائية.

## البداية السريعة

إصدار Go 1.23.6 أو أحدث متحدد في `go.mod`. من النسخة دي:

```bash
go build -o bin/anvil ./cmd/anvil
```

أنشئ ملف `.foundry.yaml`:

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

شغّل أداة الـ CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## واجهة سطر الأوامر (CLI)

| الأمر | الخيارات |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## اللغات والترجمة

`anvil` مدمج جواه ترجمات لـ 45 لغة وإعداد محلي (20 من الفئة الأولى Tier 1، و15 لغة أفريقية، و10 لغات آسيوية). اختيار اللغة بيتبع `LC_ALL` و`LC_MESSAGES` و`LANG`، وتقدر تغيّره باستخدام `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## التطوير

```bash
make build
make test
make lint
make vet
```

## الترخيص

[Apache License 2.0](../../../LICENSE).
