
# 📚 Ngày 3: Struct, Method, Pointer trong Go

## ✅ 1. Struct – Kiểu dữ liệu có cấu trúc

Struct là cách để gom nhiều trường dữ liệu vào một kiểu:
```go
type Person struct {
  name string
  age  int
}

```

Khởi tạo struct:

```go

p := Person{name: "Eminel", age: 25}

```

## ✅ 2. Method – Hàm gắn với struct

Method giúp gắn hành vi với đối tượng:

```go

func (p Person) Greet() {
  fmt.Printf("Xin chào, tôi là %s, %d tuổi.\n", p.name, p.age)
}

```

`(p Person)` gọi là receiver – giúp method nhận dữ liệu từ struct `Person`.

## ✅ 3. Pointer receiver – Sửa dữ liệu gốc

Nếu bạn muốn method thay đổi dữ liệu bên trong struct, bạn cần dùng con trỏ:

```go

func (p *Person) HaveBirthday() {
  p.age++
}

```

Lưu ý:
- `Person` (giá trị) → tạo bản sao.
- `*Person` (pointer) → thao tác trực tiếp lên struct gốc.

## ✅ 4. Ví dụ thực tế: Rectangle

```go

type Rectangle struct {
  width, height float64
}

func (r Rectangle) Area() float64 {
  return r.width * r.height
}

func (r *Rectangle) Resize(newWidth float64) {
  r.width = newWidth
}

```

- `Area()` không thay đổi dữ liệu → dùng `Rectangle` (value).
- `Resize()` thay đổi chiều rộng → dùng `*Rectangle` (pointer).

## 🧠 Tổng kết kiến thức

| Chủ đề | Ghi nhớ ngắn |
|--------|--------------|
| Struct | Gom dữ liệu thành 1 kiểu mới |
| Method | Hàm gắn với struct |
| Pointer receiver | Thay đổi dữ liệu thật |
| Value receiver | Chỉ đọc, không ảnh hưởng tới dữ liệu gốc |

🧠 Ghi nhớ: Nếu bạn muốn method thay đổi dữ liệu gốc, hãy dùng *StructName làm receiver.