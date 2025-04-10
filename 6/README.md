# 📘 Ngày 6: Module và Tổ chức Project trong Go

## 🎯 Mục tiêu

- Hiểu `go mod` là gì và cách sử dụng trong Go.
- Biết cách tổ chức một project theo module và package.
- Biết chia nhỏ chương trình ra nhiều file và package để dễ quản lý, mở rộng và bảo trì.

---

## 📦 1. Go Module là gì?

`go mod` là hệ thống quản lý dependency chính thức của Go (từ Go 1.11).

### ✨ Lợi ích:

- Quản lý thư viện bên ngoài rõ ràng.
- Tự động tải và lưu version dependency.
- Cho phép chia sẻ project dễ dàng hơn.

### 🔧 Một số lệnh thường dùng

| Lệnh                       | Mô tả                                  |
|----------------------------|-----------------------------------------|
| `go mod init <module>`     | Tạo file `go.mod` và khởi tạo module   |
| `go mod tidy`              | Dọn dẹp và tải các package cần thiết   |
| `go get <package>`         | Cài package bên ngoài vào project      |

---

## 🧱 2. Tổ chức Project theo Module và Package

### 📁 Ví dụ cấu trúc project:
```go
userapp/
├── go.mod
├── main.go
├── models/
│   └── user.go
├── services/
│   └── user_service.go
└── utils/
    └── string_utils.go
```


### 📌 Nguyên tắc tổ chức:

- Mỗi thư mục là một **package**.
- Một package có thể chứa nhiều file `.go` (cùng tên package).
- Không nên import vòng (import lẫn nhau).
- Tách rõ phần *model*, *logic nghiệp vụ*, và *hàm tiện ích*.

---

## 📂 3. Chia Code ra Nhiều File

Bạn có thể chia nhỏ code trong cùng một package ra nhiều file khác nhau.

Ví dụ:

```go
// models/user.go
package models

type User struct {
    ID   int
    Name string
    Age  int
}
```
```go
// services/user_service.go
package services

import (
    "fmt"
    "userapp/models"
)

func PrintUserInfo(u models.User) {
    fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
}
```
