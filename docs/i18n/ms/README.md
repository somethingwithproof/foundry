> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lesen](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Pelaksana CI berasaskan profil yang menukar konfigurasi YAML kepada pelan pelaksanaan mengikut susunan kebergantungan. CLI `anvil` yang sama boleh memeriksa pelan tersebut dan melaksanakan arahan secara setempat atau di dalam tugas CI.

## Reka Bentuk

- Konfigurasi (`internal/config/config.go`) mentakrifkan profil, langkah, kebergantungan, pemboleh ubah persekitaran, tamat masa, dan percubaan semula.
- Perancangan (`internal/plan`) mengesahkan susunan kebergantungan dan menghasilkan pelan pelaksanaan.
- Pelaksanaan (`internal/exec`) menjalankan langkah dengan keserentakan terhad serta merekodkan log dan keputusan.
- Dasar (`internal/policy/policy.go`) menolak langkah `script` melainkan ia didayakan. Dasar ini bukan sempadan pengasingan bagi sebarang arahan shell.

## Mula Pantas

Go 1.23.6 atau versi yang lebih baharu diisytiharkan dalam `go.mod`. Dari salinan kerja ini:

```bash
go build -o bin/anvil ./cmd/anvil
```

Cipta `.foundry.yaml`:

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

| Arahan | Pilihan |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Penyetempatan

`anvil` membenamkan terjemahan untuk 45 lokal (20 lokal Peringkat 1, 15 lokal Afrika, dan 10 lokal Asia). Pemilihan lokal mematuhi `LC_ALL`, `LC_MESSAGES`, dan `LANG` serta boleh diganti dengan `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Pembangunan

```bash
make build
make test
make lint
make vet
```

## Lesen

[Lesen Apache 2.0](../../../LICENSE).
