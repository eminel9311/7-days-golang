package main

import (
	"fmt"
	"userapp/models"
	"userapp/services"
	"userapp/utils"
)

func main() {
	fmt.Println("Chương trình quản lý người dùng")

	// Khởi tạo dịch vụ người dùng
	userService := services.NewUserService()

	// Thêm một số người dùng mẫu
	userService.Add(models.User{ID: 1, Name: "Nguyễn Văn A", Email: "nguyenvana@example.com", Age: 25})
	userService.Add(models.User{ID: 2, Name: "Trần Thị B", Email: "tranthib@example.com", Age: 30})
	userService.Add(models.User{ID: 3, Name: "Lê Văn C", Email: "levanc@example.com", Age: 22})

	// Hiển thị tất cả người dùng
	fmt.Println("\nDanh sách tất cả người dùng:")
	allUsers := userService.GetAll()
	utils.PrintUsers(allUsers)

	// Tìm người dùng theo ID
	fmt.Println("\nTìm người dùng có ID = 2:")
	user, found := userService.GetByID(2)
	if found {
		utils.PrintUser(user)
	} else {
		fmt.Println("Không tìm thấy người dùng!")
	}

	// Xóa người dùng
	fmt.Println("\nXóa người dùng có ID = 1")
	userService.Delete(1)

	// Hiển thị lại danh sách sau khi xóa
	fmt.Println("\nDanh sách sau khi xóa:")
	allUsers = userService.GetAll()
	utils.PrintUsers(allUsers)

	// Cập nhật thông tin người dùng
	fmt.Println("\nCập nhật thông tin người dùng ID = 2:")
	userService.Update(models.User{ID: 2, Name: "Trần Thị B (Đã cập nhật)", Email: "newemail@example.com", Age: 31})

	// Hiển thị sau khi cập nhật
	fmt.Println("\nDanh sách sau khi cập nhật:")
	allUsers = userService.GetAll()
	utils.PrintUsers(allUsers)
}
