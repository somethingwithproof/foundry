> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![רישיון](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

מריץ CI מבוסס פרופילים ההופך תצורת YAML לתוכנית הרצה המסודרת לפי תלויות. אותו כלי CLI בשם `anvil` יכול לבחון את התוכנית ולהריץ פקודות מקומית או בתוך משימת CI.

## תכנון

- תצורה (`internal/config/config.go`) מגדירה פרופילים, שלבים, תלויות, משתני סביבה, מגבלות זמן וניסיונות חוזרים.
- תכנון (`internal/plan`) מאמת את סדר התלויות ומפיק את תוכנית ההרצה.
- הרצה (`internal/exec`) מריצה שלבים במקביליות מוגבלת ומתעדת יומנים ותוצאות.
- מדיניות (`internal/policy/policy.go`) דוחה שלבי `script` אלא אם הופעלו. מדיניות זו אינה מהווה גבול בידוד עבור פקודות מעטפת שרירותיות.

## התחלה מהירה

גרסת Go 1.23.6 ואילך מוגדרת בקובץ `go.mod`. מתוך עותק עבודה זה:

```bash
go build -o bin/anvil ./cmd/anvil
```

צרו את הקובץ `.foundry.yaml`:

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

הריצו את כלי ה-CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| פקודה | אפשרויות |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## לוקליזציה

`anvil` מטמיע תרגומים עבור 45 אזורי שפה (20 אזורי Tier 1, 15 אזורים אפריקאיים ו-10 אזורים אסייתיים). בחירת אזור השפה מתחשבת במשתנים `LC_ALL`, `LC_MESSAGES` ו-`LANG`, וניתן לעקוף אותה באמצעות `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## פיתוח

```bash
make build
make test
make lint
make vet
```

## רישיון

[רישיון Apache 2.0](../../../LICENSE).
