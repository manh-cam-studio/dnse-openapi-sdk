# Gitflow & Release Strategy for DNSE OpenAPI SDK

Mã nguồn này chứa 2 SDK độc lập cho hai ngôn ngữ là **Go** và **Rust**. Do tính đặc thù nằm chung trong 1 Monorepo, chúng ta sẽ áp dụng chiến lược Gitflow tiêu chuẩn với một số tùy biến khi Release để tương thích với các công cụ quản lý package.

## 1. Mô Hình Các Nhánh (Branching Model)

Dự án hoạt động dựa trên 2 nhánh chính, duy trì vĩnh viễn:

- **`main`**: Chứa lịch sử các phiên bản Release. Code trên nhánh này phải luôn ở trạng thái ổn định (Production-ready). Các tag phiên bản (ví dụ: `v1.0.0`) chỉ được đánh trên nhánh này.
- **`develop`**: Nhánh tích hợp cho quá trình phát triển (Integration branch). Mọi PR cho tính năng mới đều được merge vào đây trước khi chuẩn bị gom lại tạo thành 1 phiên bản Release mới sang `main`.

### Các Nhánh Phụ Trợ Chuyển Tiếp (Temporary Branches)

- **`feature/*`** (Ví dụ: `feature/rest-orders-go`, `feature/rust-websocket`): 
  - Tách ra từ: `develop`.
  - Phục vụ phát triển chức năng mới.
  - Khi hoàn thiện tạo Pull Request (PR) hợp nhất ngược trở về nhánh `develop`.

- **`release/*`** (Ví dụ: `release/v0.1.0`): 
  - Tách ra từ: `develop`.
  - Dùng để chốt phiên bản. Việc tạo nhánh release đánh dấu việc đã kết thúc thêm tính năng mới và chỉ tập trung vào sửa lỗi (bug fixes), tinh chỉnh tài liệu, nâng cấp phiên bản trong file `Cargo.toml`/`README`.
  - Kết thúc quy trình: Được merge vào đồng thời **`main`** và **`develop`**.

- **`hotfix/*`** (Ví dụ: `hotfix/fix-auth-token`): 
  - Tách ra từ: `main`.
  - Giống với nhánh release nhưng dùng trong trường hợp khẩn cấp có lỗi phát sinh nghiêm trọng ở các bản đã live. 
  - Kết thúc quy trình: Được merge ngược lại vào cả **`main`** và **`develop`**.

---

## 2. Quy Trình Cập Nhật Phiên Bản (Versioning)

Một điểm đặc thù cần lưu ý là phần nguồn thư viện **Go** nằm trong thư mục con `go/` (có chứa `go.mod` riêng lẻ). Vì vậy, hệ thống `go list` / `go get` bắt buộc phải nhận diện tags với prefix tên thư mục.

Khi phát hành một phiên bản mới, ta sẽ tách biệt vòng đời Version hoặc dùng chung Version cho cả hai nhưng **phải đẩy 2 tags khác nhau đối với Git**:

### Đối với Go SDK
- Tag bắt buộc theo chuẩn: **`go/vX.Y.Z`** (VD: `go/v0.1.0`).
- Nhờ có prefix `go/` này, khi người dùng chạy `go get github.com/manh-cam-studio/dnse-openapi-sdk/go@v0.1.0`, Go Registry sẽ hiểu và kéo mã nguồn chính xác trị trí bên trong cấu trúc monorepo của chúng ta.

### Đối với Rust SDK
- Tag theo chuẩn: **`rust/vX.Y.Z`** hoặc chung repo là **`vX.Y.Z`**.
- Rust (Cargo) không khắt khe quá về đường dẫn Git tag nếu bạn publish cấu hình qua crates.io, tuy nhiên để dễ nhận dạng, khuyến khích đánh tag là `rust/vX.Y.Z`. Cập nhật `version = "X.Y.Z"` vào trong file `Cargo.toml`.

---

## 3. GitHub Actions CI / CD Quy Trình Tự Động

**Phát Triển Hàng Ngày (`.github/workflows/ci.yml`)**:
- Mọi thao tác `push` hay tạo `pull_request` tạo cho nhánh `main` (hoặc cấu hình thêm `develop`) đều kích hoạt action này.
- Nó sẽ cài đặt song song nền tảng Go và Rust, sau đó check-format, build và chạy các bài Test để đảm bảo mã thay đổi không làm sụp đổ thư viện, gọi là Continuous Integration.

**Phát Hành Tự Động (`.github/workflows/release.yml`)**:
- Workflow này chỉ được kích hoạt khi đẩy lên Github một Tag hệ thống bắt đầu bằng `go/v*` hoặc `rust/v*` hoặc `v*`.
- GitHub Action sẽ tự động thu gom các commit liên quan và chuyển hóa thành trang Phát Hành (Release Notes) ở tab 'Releases' của Github. Bạn lúc này chỉ việc bổ sung nội dung các thay đổi vào Release Notes nếu muốn.

---

## 4. Các Bước Thực Hành Tóm Lược Một Chu Kỳ (Release Cycle)

1. Tách nhánh từ máy nội bộ: `git checkout -b feature/awesome-thing develop`.
2. Commit và đẩy Code lên (push commits)
3. Github báo CI Tests chạy OK (Tích xanh).
4. Thực hiện Merge Pull Request từ `feature/awesome-thing` vào nhánh `develop`.
5. Sau khi đủ số lượng tính năng, tách ra một nhánh Release `git checkout -b release/v0.1.0 develop`. Sửa đổi docs, versions trong code, sau đó đẩy Pull Request để hợp nhất `release` vào `main`.
6. Tại nhánh `main`, tạo và đẩy đi các tags để báo hiệu tạo bản Release chính thức:
   ```bash
   git checkout main
   git pull origin main
   git tag -a go/v0.1.0 -m "Release Go SDK v0.1.0"
   git tag -a rust/v0.1.0 -m "Release Rust SDK v0.1.0"
   git push origin --tags
   ```
7. Quá trình CI/CD sẽ hoàn tất tự động xây dựng Release UI (Trang Github Release) trên Repository cho bạn!
