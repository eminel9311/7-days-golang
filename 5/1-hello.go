package main

import (
	"fmt"
	"time"
)

func sayHello() {
	println("Hello from Goroutine!")
}

func main1() {
	go sayHello() // Chạy sayHello() trong một goroutine mới
	fmt.Println("Hello from Main!")

	// Chương trình chính sẽ không đợi goroutine hoàn thành
	// Nếu không có gì để đợi, chương trình chính sẽ kết thúc ngay lập tức
	// Điều này có thể dẫn đến việc goroutine không được thực thi
	// Để đảm bảo rằng goroutine có thời gian để thực thi, chúng ta có thể sử dụng time.Sleep
	// Hoặc sử dụng một cơ chế đồng bộ hóa như WaitGroup
	// Ở đây, chúng ta sẽ sử dụng time.Sleep để đơn giản hóa
	time.Sleep(500 * time.Millisecond)
}
