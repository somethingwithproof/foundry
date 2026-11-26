> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![Lisans](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

YAML yapılandırmasını bağımlılık sırasına göre düzenlenmiş bir yürütme planına dönüştüren profil tabanlı bir CI çalıştırıcısıdır. Aynı `anvil` CLI aracı, planı inceleyebilir ve komutları yerel olarak ya da bir CI işi içinde çalıştırabilir.

## Tasarım

- Yapılandırma (`internal/config/config.go`), profilleri, adımları, bağımlılıkları, ortam değişkenlerini, zaman aşımlarını ve yeniden denemeleri tanımlar.
- Planlama (`internal/plan`), bağımlılık sıralamasını doğrular ve yürütme planını üretir.
- Yürütme (`internal/exec`), adımları sınırlı eşzamanlılıkla çalıştırır, günlükleri ve sonuçları kaydeder.
- İlke (`internal/policy/policy.go`), etkinleştirilmediği sürece `script` adımlarını reddeder. Bu ilke, rastgele kabuk komutları için bir yalıtım sınırı değildir.

## Hızlı başlangıç

`go.mod` içinde Go 1.23.6 veya üzeri tanımlanmıştır. Bu çalışma kopyasından:

```bash
go build -o bin/anvil ./cmd/anvil
```

`.foundry.yaml` dosyasını oluşturun:

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

CLI aracını çalıştırın:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Komut | Seçenekler |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Yerelleştirme

`anvil`, 45 yerel ayar (20 Tier 1, 15 Afrika ve 10 Asya yerel ayarı) için çevirileri gömülü olarak içerir. Yerel ayar seçimi `LC_ALL`, `LC_MESSAGES` ve `LANG` değişkenlerini dikkate alır ve `--lang` ile geçersiz kılınabilir:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Geliştirme

```bash
make build
make test
make lint
make vet
```

## Lisans

[Apache License 2.0](../../../LICENSE).
