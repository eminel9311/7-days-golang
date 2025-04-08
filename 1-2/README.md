
## Ngày 1–2: Cú pháp cơ bản trong Golang

### 1. Package và Hàm Main

Mỗi chương trình Go đều bắt đầu bằng một package, và nếu bạn muốn chương trình chạy được, nó phải có package main và một hàm `func main()`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

- `package main`: Chỉ ra rằng đây là entry point (chương trình chính).
- `func main()`: Hàm main là điểm khởi đầu khi chạy chương trình.


### 2. Khai báo Biến: `var`, `:=`, `const`

- **`var`**: Khai báo có kiểu rõ ràng.

```go
var age int = 30
var name string = "Eminel"
```

- **`:=`**: Khai báo nhanh, Go tự suy luận kiểu.

```go
age := 30
name := "Eminel"
```

Dùng `:=` khi bạn đang trong hàm và muốn code gọn hơn. Dùng `var`/`const` khi khai báo ở ngoài hàm hoặc muốn rõ kiểu.
- **`const`**: Khai báo hằng số.

```go
const PI = 3.14
```


### 3. Hàm – Truyền và Trả Giá Trị

- **Định nghĩa hàm**:

```go
func add(a int, b int) int {
    return a + b
}

result := add(3, 5) // result = 8
```

- Hàm có thể trả nhiều giá trị:

```go
func swap(a, b string) (string, string) {
    return b, a
}
```


### 4. Câu Lệnh Điều Kiện: `if`, `switch`

- **`if` đơn giản**:

```go
if age &gt; 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Teen")
}
```

- **`if` có khai báo biến**:

```go
if score := 80; score &gt;= 75 {
    fmt.Println("Pass")
}
```

- **`switch`**:

```go
switch day := "mon"; day {
case "mon":
    fmt.Println("Monday")
case "tue":
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```


### 5. Vòng Lặp: `for`

Go chỉ có một loại vòng lặp duy nhất là `for`, nhưng nó rất linh hoạt.

- **Dạng C-style**:

```go
for i := 0; i &lt; 5; i++ {
    fmt.Println(i)
}
```

- **Dạng while**:

```go
i := 0
for i &lt; 5 {
    fmt.Println(i)
    i++
}
```

- **Lặp qua slice hoặc map**:

```go
nums := []int{1, 2, 3}
for index, value := range nums {
    fmt.Println(index, value)
}
```


### 🔚 Tổng Kết

| Chủ đề | Ví dụ |
| :-- | :-- |
| Biến | `var x int = 10`, `y := "hello"` |
| Hàm | `func add(a int, b int) int` |
| If | `if x &gt; 10 {}` |
| Switch | `switch x { case 1: ... }` |
| Vòng lặp | `for i := 0; i &lt; 10; i++ {}` |

