// Đề bài:
// Bài toán: Tính tổng các số trong slice bằng nhiều goroutine
// Chia slice ra thành 4 phần
// Mỗi phần tính tổng riêng và gửi kết quả qua channel
// Goroutine chính cộng lại để ra kết quả cuối

package main

// Tính tổng một đoạn slice và gửi kết quả vào channel
func patialSum(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum // Gửi kết quả vào channel
}

// Tính tổng các số trong slice bằng nhiều goroutine
func main6() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // Slice chứa các số từ 1 đến 10

	// chia slice thành 4 phần
	partSize := len(numbers) / 4 // Kích thước của mỗi phần ~3

	ch := make(chan int, 4) // Tạo một channel với kích thước buffer là 4

	go patialSum(numbers[:partSize], ch)             // Tính tổng phần đầu tiên từ 0 đến 2
	go patialSum(numbers[partSize:2*partSize], ch)   // Tính tổng phần thứ hai từ 3 đến 5
	go patialSum(numbers[2*partSize:3*partSize], ch) // Tính tổng phần thứ ba từ 6 đến 8
	go patialSum(numbers[3*partSize:], ch)           // Tính tổng phần thứ tư từ 9 đến 11

	// Tổng hợp kết quả từ các goroutine
	total := 0
	for range 4 {
		partial := <-ch
		total += partial
	}

	println("Total sum:", total) // In ra tổng

}
