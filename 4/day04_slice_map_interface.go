package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("=== Slice ===")
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	for i, v := range nums {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	fmt.Println("\n=== Map ===")
	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	scores["Charlie"] = 88
	fmt.Println("Alice:", scores["Alice"])

	if val, ok := scores["Bob"]; ok {
		fmt.Println("Bob:", val)
	} else {
		fmt.Println("Bob không tồn tại")
	}

	delete(scores, "Bob")

	fmt.Println("\n=== Interface ===")
	var s Shape

	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 6}

	s = c
	fmt.Printf("Diện tích hình tròn: %.2f\n", s.Area())

	s = r
	fmt.Printf("Diện tích hình chữ nhật: %.2f\n", s.Area())

	fmt.Println("\n=== TESTTT ===")
	val, ok := scores["Charlie"]
	fmt.Println("hehe:", val)
	fmt.Println("ok:", ok)

}
