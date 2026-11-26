> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![الترخيص](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

مشغل CI مبني على الملفات الشخصية (profiles) يحول إعدادات YAML إلى خطة تنفيذ مرتبة حسب التبعيات. يمكن لأداة سطر الأوامر `anvil` نفسها فحص الخطة وتنفيذ الأوامر محليا أو داخل مهمة CI.

## التصميم

- الإعداد (`internal/config/config.go`) يحدد الملفات الشخصية والخطوات والتبعيات ومتغيرات البيئة ومهل الانتظار وإعادات المحاولة.
- التخطيط (`internal/plan`) يتحقق من صحة ترتيب التبعيات وينتج خطة التنفيذ.
- التنفيذ (`internal/exec`) يشغل الخطوات بتزامن محدود ويسجل السجلات والنتائج.
- السياسة (`internal/policy/policy.go`) ترفض خطوات `script` ما لم يتم تفعيلها. هذه السياسة ليست حدود عزل لأوامر الصدفة العشوائية.

## البدء السريع

يتطلب المشروع Go 1.23.6 أو أحدث كما هو محدد في `go.mod`. انطلاقا من هذه النسخة المستنسخة:

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

شغل أداة سطر الأوامر:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## واجهة سطر الأوامر

| الأمر | الخيارات |
| --- | --- |
| `doctor` | `--config`، `--lang` |
| `plan` | `--config`، `--profile`، `--json`، `--lang` |
| `run` | `--config`، `--profile`، `--jobs`، `--json`، `--lang` |
| `version` | `--json`، `--lang` |

## التوطين

تدمج أداة `anvil` ترجمات لـ 45 لغة محلية (20 من المستوى 1، و15 لغة إفريقية، و10 لغات آسيوية). يحترم اختيار اللغة متغيرات `LC_ALL` و`LC_MESSAGES` و`LANG` ويمكن تجاوزه عبر `--lang`:

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
