> [English](../../../README.md)

# Foundry

[![CI](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/somethingwithproof/foundry/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](../../../LICENSE)

Trình chạy CI dựa trên profile giúp chuyển đổi tệp cấu hình YAML thành kế hoạch thực thi được sắp xếp theo thứ tự phụ thuộc. Chính CLI `anvil` này có thể kiểm tra kế hoạch và thực thi các lệnh ở máy cục bộ hoặc bên trong một tác vụ CI.

## Thiết kế

- Cấu hình (`internal/config/config.go`) định nghĩa các profile, bước, phần phụ thuộc, biến môi trường, thời gian chờ và số lần thử lại.
- Lập kế hoạch (`internal/plan`) kiểm tra tính hợp lệ của thứ tự phụ thuộc và tạo ra kế hoạch thực thi.
- Thực thi (`internal/exec`) chạy các bước với mức độ đồng thời giới hạn, đồng thời ghi lại nhật ký và kết quả.
- Chính sách (`internal/policy/policy.go`) từ chối các bước `script` trừ khi được bật. Chính sách này không phải là ranh giới cô lập đối với các lệnh shell tùy ý.

## Bắt đầu nhanh

Phiên bản Go 1.23.6 trở lên được khai báo trong `go.mod`. Từ bản sao mã nguồn này:

```bash
go build -o bin/anvil ./cmd/anvil
```

Tạo tệp `.foundry.yaml`:

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

Chạy CLI:

```bash
bin/anvil doctor
bin/anvil plan --profile default
bin/anvil run --profile default --jobs 4
```

## CLI

| Lệnh | Tùy chọn |
| --- | --- |
| `doctor` | `--config`, `--lang` |
| `plan` | `--config`, `--profile`, `--json`, `--lang` |
| `run` | `--config`, `--profile`, `--jobs`, `--json`, `--lang` |
| `version` | `--json`, `--lang` |

## Bản địa hóa

`anvil` nhúng sẵn bản dịch cho 45 ngôn ngữ địa phương (20 ngôn ngữ Bậc 1, 15 ngôn ngữ châu Phi và 10 ngôn ngữ châu Á). Việc chọn ngôn ngữ tuân theo `LC_ALL`, `LC_MESSAGES` và `LANG`, đồng thời có thể ghi đè bằng `--lang`:

```bash
anvil --lang de doctor
LANG=ja_JP.UTF-8 anvil version
```

## Phát triển

```bash
make build
make test
make lint
make vet
```

## Giấy phép

[Apache License 2.0](../../../LICENSE).
