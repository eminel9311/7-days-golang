package main

import "fmt"

func main2() {
	i, j := 32, 1240

	var p *int // Khai báo p là con trỏ kiểu int

	p = &i          // p là con trỏ trỏ đến biến i
	fmt.Println(p)  // In ra địa chỉ của i -> 0xc000192040
	fmt.Println(*p) // In ra giá trị của i thông qua con trỏ p -> 32
	*p = 100        // Thay đổi giá trị của i thông qua con trỏ p
	fmt.Println(i)  // In ra giá trị của i sau khi thay đổi -> 100

	p = &j         // p giờ trỏ đến biến j
	fmt.Println(p) // In ra địa chỉ của j -> 0xc0000120f8
	*p = *p / 2    // Thay đổi giá trị của j thông qua con trỏ p
	fmt.Println(j) // In ra giá trị của j sau khi thay đổi
}
