> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

مشغّل تكامل مستمر (CI) يعتمد على ملفات التعريف، ويحوّل إعدادات YAML إلى خطة تنفيذ مرتّبة وفق الاعتماديات. تستطيع أداة سطر الأوامر `anvil` نفسها فحص الخطة وتنفيذ الأوامر محليًا أو داخل مهمة CI.

## التصميم

- الإعدادات (`internal/config/config.go`) تعرّف ملفات التعريف، والخطوات، والاعتماديات، ومتغيّرات البيئة، ومُهَل الانتظار، وإعادات المحاولة.
- التخطيط (`internal/plan`) يتحقّق من صحة ترتيب الاعتماديات وينشئ خطة التنفيذ.
- التنفيذ (`internal/exec`) يشغّل الخطوات بتزامن محدود ويسجّل السجلات والنتائج.
- السياسة (`internal/policy/policy.go`) ترفض خطوات `script` ما لم تُفعَّل. لا تشكّل هذه السياسة حدود عزل لأوامر الصدفة العشوائية.

## البدء السريع

إصدار Go 1.23.6 أو أحدث مُحدَّد في `go.mod`. انطلاقًا من هذه النسخة المستنسخة:

```bash
go build -o bin/anvil ./cmd/anvil
```

أنشئ الملف `.foundry.yaml`:

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

شغّل أداة سطر الأوامر:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## واجهة سطر الأوامر

| الأمر | الخيارات |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## التوطين

تُضمِّن `anvil` ترجمات لـ 45 لغةً وإعدادًا محليًا (20 لغةً من المستوى الأول، و15 لغةً أفريقيةً، و10 لغاتٍ آسيويةٍ). يحترم اختيار اللغة متغيّرات `LC_ALL` و`LC_MESSAGES` و`LANG`، ويمكن تجاوزه باستخدام `--lang`:

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
