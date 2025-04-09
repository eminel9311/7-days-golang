package main

import "fmt"

// hàm sayHi sẽ gửi một chuỗi vào kênh

func sayHi(ch chan string) {
	ch <- "Hello from Goroutine!"
}

func main2() {
	ch := make(chan string) // Tạo một kênh kiểu string
	go sayHi(ch)            // Gọi hàm sayHi trong một goroutine mới
	msg := <-ch             // Nhận thông điệp từ kênh
	fmt.Println(msg)
}
