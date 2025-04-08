package main

import "fmt"

func main() {
	fmt.Println("=== Ngày 1–2: Cú pháp cơ bản ===")

	// 1. Khai báo biến
	var age int = 18
	name := "Eminel"
	const PI = 3.14

	fmt.Println("Tên:", name)
	fmt.Println("Tuổi:", age)
	fmt.Println("Hằng số PI:", PI)

	// 2. Hàm + truyền/trả giá trị
	sum := add(3, 4)
	fmt.Println("Tổng:", sum)

	first, second := swap("Hello", "World")
	fmt.Println("Hoán đổi chuỗi:", first, second)

	// 3. Câu lệnh điều kiện
	if age >= 18 {
		fmt.Println("Bạn đã đủ tuổi trưởng thành")
	} else {
		fmt.Println("Bạn chưa đủ tuổi trưởng thành")
	}

	// if có khai báo biến
	if score := 80; score >= 85 {
		fmt.Println("Bạn đã đỗ")
	}

	// switch
	day := "mon"
	switch day {
	case "mon":
		fmt.Println("Thứ 2")
	case "tue":
		fmt.Println("Thứ 3")
	case "wed":
		fmt.Println("Thứ 4")
	default:
		fmt.Println("Không phải thứ 2, 3, 4")
	}

	// 4. Vòng lặp
	fmt.Println("For kiểu C-style:")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("For kiểu while:")
	i := 0
	for i < 5 {
		fmt.Println("i =", i)
		i++
	}

	fmt.Println("For kiểu range:")
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("slice[%d] = %d\n", index, value)
	}

}

// Hàm cộng 2 số
func add(a int, b int) int {
	return a + b
}

// Hàm hoán đổi 2 chuỗi
func swap(a, b string) (string, string) {
	return b, a
}
