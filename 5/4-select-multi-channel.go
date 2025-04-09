package main

import "time"

func main4() {
	ch1 := make(chan string) // Tạo một kênh kiểu string
	ch2 := make(chan string) // Tạo một kênh kiểu string

	go func() {
		time.Sleep(4 * time.Second) // Giả lập một tác vụ mất 1 giây
		ch1 <- "from ch1"           // Gửi thông điệp vào kênh ch1
	}()

	go func() {
		time.Sleep(5 * time.Second) // Giả lập một tác vụ mất 2 giây
		ch2 <- "from ch2"           // Gửi thông điệp vào kênh ch2
	}()

	// Sử dụng select để chờ nhận thông điệp từ cả hai kênh
	// Nếu có thông điệp đến từ ch1 hoặc ch2, sẽ in ra thông điệp đó
	// Nếu cả hai kênh đều có thông điệp đến, sẽ in ra thông điệp từ kênh nào đến trước
	// Nếu không có thông điệp nào đến trong vòng 3 giây, sẽ in ra thông báo timeout
	// Chú ý rằng select sẽ chờ cho đến khi có ít nhất một kênh có thông điệp đến
	select {
	case msg1 := <-ch1:
		println("Received", msg1)
	case msg2 := <-ch2:
		println("Received", msg2)
	case <-time.After(3 * time.Second):
		println("Timeout: No message received within 3 seconds")
	}
}
