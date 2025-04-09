
# 📚 Ngày 5: Goroutine, Channel

## 1. Giới thiệu về Goroutine

Goroutine là một trong những tính năng nổi bật nhất của ngôn ngữ Go, được thiết kế để hỗ trợ lập trình đồng thời (concurrent programming) một cách đơn giản và hiệu quả. Goroutine có thể được coi là một hàm hoặc phương thức chạy độc lập với chương trình chính, cho phép thực hiện nhiều tác vụ song song.

## 2. Goroutine vs Thread truyền thống

### Kích thước và tài nguyên
- **Thread:** Thường có kích thước stack cố định (thường là 1-2MB), tốn nhiều tài nguyên hệ thống
- **Goroutine:** Bắt đầu với stack nhỏ (chỉ khoảng 2KB) và có thể mở rộng/thu hẹp động theo nhu cầu

### Số lượng
- **Thread:** Hệ thống thường chỉ hỗ trợ hàng nghìn thread
- **Goroutine:** Có thể chạy hàng trăm nghìn, thậm chí hàng triệu goroutine trên một máy thông thường

### Quản lý
- **Thread:** Quản lý bởi hệ điều hành
- **Goroutine:** Quản lý bởi Go runtime (scheduler của Go), độc lập với OS thread

### Chuyển đổi ngữ cảnh (Context switching)
- **Thread:** Tốn nhiều chi phí do phải lưu trữ và khôi phục nhiều thông tin
- **Goroutine:** Chi phí thấp hơn nhiều vì được quản lý trong không gian người dùng

## 3. Go Runtime Scheduler

Scheduler của Go phân phối goroutine trên một tập các OS thread:
- Sử dụng `M:N` scheduling model
- M goroutines chạy trên N OS threads
- Mặc định N = GOMAXPROCS (thường bằng số lõi CPU)
- Tự động điều phối goroutine giữa các thread

## 4. Cú pháp và cách sử dụng Goroutine

### Cú pháp cơ bản
```go
go functionName(parameters)
```

### Ví dụ đơn giản
```go
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello() // Chạy sayHello() trong một goroutine mới
    
    // Cần một cơ chế để đợi goroutine hoàn thành
    // Nếu không, chương trình chính có thể kết thúc trước
    time.Sleep(100 * time.Millisecond)
}
```

### Sử dụng với hàm ẩn danh (anonymous function)
```go
func main() {
    go func() {
        fmt.Println("Hello from anonymous goroutine!")
    }()
    
    time.Sleep(100 * time.Millisecond)
}
```

## 5. WaitGroup để đồng bộ hóa Goroutine

Trong thực tế, thay vì sử dụng `time.Sleep()`, chúng ta thường dùng `sync.WaitGroup` để đợi các goroutine hoàn thành:

```go
func main() {
    var wg sync.WaitGroup
    
    // Thêm 1 goroutine vào WaitGroup
    wg.Add(1)
    
    go func() {
        // Báo hiệu goroutine đã hoàn thành khi thoát
        defer wg.Done()
        
        fmt.Println("Working in goroutine")
    }()
    
    // Đợi tất cả goroutine trong WaitGroup hoàn thành
    wg.Wait()
    fmt.Println("Main goroutine exits")
}
```

## 6. Chia sẻ dữ liệu giữa các Goroutine

### Lưu ý về race condition
Khi nhiều goroutine cùng truy cập vào một biến, có thể xảy ra race condition:

```go
func main() {
    counter := 0
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // Race condition
        }()
    }
    
    wg.Wait()
    fmt.Println("Counter:", counter) // Kết quả không chính xác
}
```

### Giải pháp với Mutex
```go
func main() {
    counter := 0
    var wg sync.WaitGroup
    var mu sync.Mutex
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Println("Counter:", counter) // Kết quả chính xác
}
```

## 7. Goroutine Leaks và cách tránh

Goroutine leak xảy ra khi goroutine không bao giờ kết thúc, dẫn đến rò rỉ bộ nhớ:

```go
// Goroutine leak
func leakyFunction() {
    ch := make(chan int)
    go func() {
        val := <-ch // Goroutine này sẽ chờ mãi mãi
        fmt.Println("Received:", val)
    }()
    // Không bao giờ gửi giá trị vào channel
    // Goroutine sẽ không bao giờ kết thúc
}
```

Các cách tránh leak:
- Luôn có cơ chế đóng channel
- Sử dụng context để hủy goroutine
- Đảm bảo tất cả goroutine đều có cơ chế thoát

## 8. Những điểm cần lưu ý

- Goroutine không có ID hay cách truy cập trực tiếp từ bên ngoài
- Main goroutine kết thúc sẽ dừng tất cả goroutine khác
- Nên sử dụng channel để giao tiếp giữa các goroutine
- Tránh sử dụng biến được chia sẻ nếu có thể
- Sử dụng `go run -race` để phát hiện race condition
- Goroutine rất nhẹ nhưng không phải miễn phí, cần tránh lạm dụng quá nhiều

## 9. Các pattern phổ biến

### Worker Pool
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    numJobs := 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // Khởi động 3 worker
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Gửi công việc
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Thu thập kết quả
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

### Pipeline
```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Tạo pipeline: generator -> square -> in kết quả
    for n := range square(generator(1, 2, 3, 4, 5)) {
        fmt.Println(n)
    }
}
```

Goroutine là một trong những tính năng mạnh mẽ nhất của Go, cho phép xây dựng các chương trình đồng thời một cách dễ dàng và hiệu quả. Hiểu rõ về cách hoạt động và cách sử dụng goroutine sẽ giúp bạn khai thác tối đa sức mạnh của Go trong các ứng dụng cần xử lý đồng thời.
