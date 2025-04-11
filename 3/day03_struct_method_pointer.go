package main

import "fmt"

// 1. Định nghĩa struct
type Person struct {
	name string
	age  int
}

// 2. Method với giá trị receiver
// Lưu ý rằng vì Greet() được định nghĩa với receiver là giá trị (Person), nên nó không thay đổi dữ liệu của struct gốc. Nó chỉ đọc và sử dụng dữ liệu để in ra màn hình.
func (p Person) Greet() {
	fmt.Printf("Xin chào, tôi là %s, %d tuổi\n", p.name, p.age)
}

// 3. Method với pointer receiver (con trỏ)
// Lưu ý rằng vì Greet() được định nghĩa với receiver là con trỏ (*Person), nên nó có thể thay đổi dữ liệu của struct gốc.
func (p *Person) HaveBirthday() {
	p.age++ // thay đổi trực tiếp giá trị age gốc
}

// Struct khác dùng pointer
type Rectangle struct {
	width, height float64
}

// Diện tích hình chữ nhật
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method thay đổi chiều rộng
func (r *Rectangle) Resize(newWidth float64) {
	r.width = newWidth
}

func main1() {
	fmt.Println("=== Ngày 3: Struct, Method, Pointer ===")

	// Khai báo và khởi tạo struct
	person := Person{name: "Eminel", age: 18}
	person.Greet() // Gọi method Greet()

	// Dùng method có pointer để thay đổi dữ liệu
	person.HaveBirthday()
	person.Greet() // Gọi lại method Greet() để xem thay đổi

	// Rectangle demo
	rect := Rectangle{width: 10, height: 5}
	fmt.Printf("Diện tích hình chữ nhật: %.2f\n", rect.Area())

	// Thay đổi chiều rộng
	rect.Resize(20)
	fmt.Printf("Diện tích hình chữ nhật sau khi thay đổi chiều rộng: %.2f\n", rect.Area())

}
