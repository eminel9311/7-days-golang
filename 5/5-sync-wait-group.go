package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done() // Đánh dấu goroutine này đã hoàn thành
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(5 * time.Second)        // Giả lập công việc mất 5 giây
	fmt.Printf("Worker %d done\n", id) // In ra thông báo khi hoàn thành
}

func main5() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)    // Thêm một goroutine vào WaitGroup
		go worker(i) // Khởi động goroutine
	}
	wg.Wait() // Chờ tất cả goroutines hoàn thành
}
