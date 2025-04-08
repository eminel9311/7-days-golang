
# 📚 Ngày 4: Slice, Map, Interface trong Go

## ✅ 1. Slice – Mảng động

Slice là mảng có kích thước động, dễ thao tác hơn array.

### Khai báo và thao tác:

```go

nums := []int{1, 2, 3}
nums = append(nums, 4)

for i, v := range nums {
  fmt.Println("Index:", i, "Value:", v)
}

```

Một số hàm với slice:
- **append(slice, item)** – thêm phần tử
- **len(slice)** – độ dài
- **slice[start:end]** – cắt slice

## ✅ 2. Map – Bản đồ key-value

Map lưu trữ dữ liệu theo cặp key =&gt; value.

```go

scores := map[string]int{
  "Alice": 90,
  "Bob":   75,
}

scores["Charlie"] = 88
fmt.Println("Điểm của Alice:", scores["Alice"])

delete(scores, "Bob")

```

Kiểm tra key tồn tại:

```go

val, ok := scores["Bob"]
if ok {
  fmt.Println("Tìm thấy Bob:", val)
} else {
  fmt.Println("Không có Bob")
}

```

## ✅ 3. Interface – Giao diện

Interface định nghĩa hành vi (method) mà struct cần triển khai.

```go

type Shape interface {
  Area() float64
}

```

Struct nào implement đủ các method thì tự động là interface đó:

```go

type Circle struct {
  radius float64
}

func (c Circle) Area() float64 {
  return 3.14 * c.radius * c.radius
}

```

Dùng interface:

```go

func printArea(s Shape) {
  fmt.Println("Diện tích:", s.Area())
}

```

🧠 Tổng kết
| Kiến thức | Mô tả ngắn |
|-----------|------------|
| **Slice** | Mảng động, dùng append, range, len |
| **Map**   | Bản đồ key =&gt; value, thao tác thêm/xóa |
| **Interface** | Định nghĩa hành vi, cho phép xử lý đa hình |
