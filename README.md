# Go Programming - 7 Days Quick Learning Plan 🚀

## 🎯 Mục tiêu
Giúp bạn nắm vững kiến thức cơ bản về ngôn ngữ lập trình Go (Golang) chỉ trong **7 ngày**, thông qua lý thuyết trọng tâm, thực hành và mini project.

---

## 🗓️ Lộ trình học nhanh

| Ngày | Chủ đề                           | Mục tiêu học được                                                         |
|------|----------------------------------|---------------------------------------------------------------------------|
| 1-2  | Cú pháp cơ bản                   | Hiểu cách khai báo biến, hằng, hàm, if, for, switch                       |
| 3    | Struct, Method, Pointer          | Biết dùng struct, method, truyền con trỏ (pointer)                       |
| 4    | Slice, Map, Interface            | Làm việc với slice, map và xây dựng interface                            |
| 5    | Goroutine, Channel               | Hiểu lập trình song song với goroutine và channel                         |
| 6    | Module, tổ chức project          | Tổ chức code với module, chia file và package hợp lý                      |
| 7    | Mini Project                     | Xây dựng project nhỏ áp dụng toàn bộ kiến thức đã học                    |

---

## 📘 Nội dung chi tiết từng ngày

### ✅ Ngày 1–2: Cú pháp cơ bản
- `package main`, `func main()`
- Biến: `var`, `:=`, `const`
- Hàm: truyền và trả giá trị
- Câu lệnh điều kiện: `if`, `switch`
- Vòng lặp: `for`

### ✅ Ngày 3: Struct, Method, Pointer
- Tạo `struct` để biểu diễn đối tượng
- Gắn method cho struct
- Truyền con trỏ để cập nhật giá trị

### ✅ Ngày 4: Slice, Map, Interface
- Khai báo và thao tác với slice, map
- Định nghĩa interface và implement nó với struct cụ thể

### ✅ Ngày 5: Goroutine, Channel
- Tạo goroutine để chạy song song
- Giao tiếp giữa các goroutine bằng `chan`

### ✅ Ngày 6: Module, tổ chức project
- Sử dụng `go mod`
- Tổ chức project theo module và package
- Chia code ra nhiều file dễ quản lý

### ✅ Ngày 7: Mini Project
- REST API đơn giản với `net/http` hoặc `gin`
- Hoặc CLI tool quản lý việc cần làm
- Cấu trúc thư mục, xử lý logic, tách controller, model, utils

---

## 🛠️ Yêu cầu trước khi bắt đầu
- Đã biết một ngôn ngữ lập trình khác (C, JavaScript, Python,...)
- Cài đặt môi trường Go: https://go.dev/doc/install

---

## 📦 Gợi ý cấu trúc project
my-project/
├── main.go
├── controller/
│   └── post.go
├── model/
│   └── post.go
├── utils/
│   └── helper.go
└── go.mod
