package main

func main3() {
	ch := make(chan int, 3) // Tạo một kênh với kích thước buffer là 3
	ch <- 1                 // Gửi giá trị 1 vào kênh
	ch <- 2                 // Gửi giá trị 2 vào kênh
	ch <- 3                 // Gửi giá trị 3 vào kênh
	// ch <- 4                 // Nếu gửi giá trị thứ 4, chương trình sẽ bị treo vì kênh đã đầy
	// fmt.Println(<-ch) // Nhận giá trị từ kênh và in ra
	// fmt.Println(<-ch) // Nhận giá trị tiếp theo từ kênh và in ra
	// fmt.Println(<-ch) // Nhận giá trị cuối cùng từ kênh và in ra

	// ftm.Println(<-ch) // Nếu nhận giá trị thứ 4, chương trình sẽ bị treo vì kênh đã rỗng
}
