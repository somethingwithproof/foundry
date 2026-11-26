> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lisensi](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Runner CI berbasis profil yang mengubah konfigurasi YAML menjadi rencana eksekusi berurutan sesuai dependensi. CLI `anvil` yang sama dapat memeriksa rencana dan menjalankan perintah secara lokal maupun di dalam job CI.

## Desain

- Konfigurasi (`internal/config/config.go`) mendefinisikan profil, langkah, dependensi, variabel lingkungan, batas waktu, dan percobaan ulang.
- Perencanaan (`internal/plan`) memvalidasi urutan dependensi dan menghasilkan rencana eksekusi.
- Eksekusi (`internal/exec`) menjalankan langkah dengan konkurensi terbatas serta mencatat log dan hasilnya.
- Kebijakan (`internal/policy/policy.go`) menolak langkah `script` kecuali diaktifkan. Kebijakan ini bukan merupakan batas isolasi untuk perintah shell sembarangan.

## Mulai cepat

Go 1.23.6 atau lebih baru dideklarasikan di `go.mod`. Dari salinan kode ini:

```bash
go build -o bin/anvil ./cmd/anvil
```

Buat `.foundry.yaml`:

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

Jalankan CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Perintah | Opsi |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Pelokalan

`anvil` menyertakan terjemahan untuk 45 lokal (20 lokal Tier 1, 15 lokal Afrika, dan 10 lokal Asia). Pemilihan lokal mengikuti `LC_ALL`, `LC_MESSAGES`, dan `LANG`, serta dapat ditimpa dengan `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Pengembangan

```bash
make build
make test
make lint
make vet
```

## Lisensi

[Apache License 2.0](../../../LICENSE).
