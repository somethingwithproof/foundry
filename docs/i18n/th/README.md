> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

ตัวรัน CI แบบอิงโปรไฟล์ที่แปลงไฟล์การตั้งค่า YAML ให้เป็นแผนการทำงานที่เรียงลำดับตามการพึ่งพา CLI ของ `anvil` เดียวกันนี้สามารถตรวจสอบแผนและรันคำสั่งได้ทั้งบนเครื่องของคุณหรือภายในงาน CI

## การออกแบบ

- การตั้งค่า (`internal/config/config.go`) กำหนดโปรไฟล์ ขั้นตอน การพึ่งพา ตัวแปรสภาพแวดล้อม ค่าหมดเวลา และการลองใหม่
- การวางแผน (`internal/plan`) ตรวจสอบความถูกต้องของลำดับการพึ่งพาและสร้างแผนการทำงาน
- การทำงาน (`internal/exec`) รันขั้นตอนต่างๆ ด้วยการทำงานพร้อมกันแบบจำกัดขอบเขต และบันทึกล็อกและผลลัพธ์
- นโยบาย (`internal/policy/policy.go`) ปฏิเสธขั้นตอนแบบ `script` เว้นแต่จะเปิดใช้งาน นโยบายนี้ไม่ใช่ขอบเขตการแยกส่วนสำหรับการรันคำสั่งเชลล์ใดๆ

## เริ่มต้นใช้งานอย่างรวดเร็ว

ใน `go.mod` กำหนดให้ใช้ Go 1.23.6 หรือใหม่กว่า จากสำเนาโค้ดนี้:

```bash
go build -o bin/anvil ./cmd/anvil
```

สร้างไฟล์ `.foundry.yaml`:

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

รัน CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| คำสั่ง | ตัวเลือก |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## การรองรับภาษาท้องถิ่น

`anvil` ฝังคำแปลสำหรับ 45 ภาษาท้องถิ่น (ภาษาท้องถิ่นระดับ Tier 1 จำนวน 20 ภาษา แอฟริกา 15 ภาษา และเอเชีย 10 ภาษา) การเลือกภาษาท้องถิ่นจะยึดตาม `LC_ALL`, `LC_MESSAGES` และ `LANG` และสามารถกำหนดทับได้ด้วย `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## การพัฒนา

```bash
make build
make test
make lint
make vet
```

## สัญญาอนุญาต

[Apache License 2.0](../../../LICENSE).
