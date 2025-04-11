## Cấu trúc Dự án
```go
// PROJECT STRUCTURE
// ----------------
// go-bookstore/
// ├── cmd/
// │   └── api/
// │       └── main.go
// ├── internal/
// │   ├── controller/
// │   │   └── book_controller.go
// │   ├── model/
// │   │   └── book.go
// │   ├── repository/
// │   │   └── book_repository.go
// │   └── service/
// │       └── book_service.go
// ├── pkg/
// │   └── response/
// │       └── response.go
// └── go.mod
```

Dự án được tổ chức theo mô hình sạch với cấu trúc thư mục rõ ràng:

- **cmd/api/main.go**: Điểm khởi đầu của ứng dụng
- **internal/**: Chứa code cốt lõi của ứng dụng
    - **controller/**: Xử lý HTTP request/response
    - **model/**: Định nghĩa cấu trúc dữ liệu
    - **repository/**: Xử lý lưu trữ dữ liệu
    - **service/**: Logic nghiệp vụ
- **pkg/**: Mã có thể tái sử dụng cho các dự án khác


## Tính Năng Chính

- **CRUD operations cho sách**:
    - Tạo sách mới
    - Lấy tất cả sách
    - Lấy sách theo ID
    - Cập nhật thông tin sách
    - Xóa sách
- **In-memory Repository**:
    - Lưu trữ dữ liệu trong bộ nhớ (dễ thay thế bằng database thật)
    - Thread-safe với mutex
- **Chuẩn hóa response**:
    - Format phản hồi nhất quán
    - Xử lý lỗi chuyên nghiệp


## Cách Chạy Dự Án

1. Tạo cấu trúc thư mục như đã mô tả ở trên
2. Khởi tạo module: `go mod init go-bookstore`
3. Cài đặt dependencies: `go get github.com/gin-gonic/gin github.com/google/uuid`
4. Chạy ứng dụng: `go run cmd/api/main.go`
    - API sẽ chạy ở `http://localhost:8080`

## API Endpoints

- **GET /api/v1/books**: Lấy tất cả sách
- **GET /api/v1/books/:id**: Lấy sách theo ID
- **POST /api/v1/books**: Tạo sách mới
- **PUT /api/v1/books/:id**: Cập nhật thông tin sách
- **DELETE /api/v1/books/:id**: Xóa sách

Dự án này tuân theo các nguyên tắc thiết kế phần mềm: tách biệt các thành phần logic, sử dụng interfaces để dependency injection, và chuẩn hóa các phản hồi API. Bạn có thể dễ dàng mở rộng nó bằng cách thêm middleware authentication, validation logic phức tạp hơn, hoặc thay thế in-memory repository bằng cơ sở dữ liệu thực như MySQL hoặc MongoDB.

