package utils

import (
	"fmt"
	"userapp/models"
)

// PrintUser hiển thị thông tin chi tiết của một người dùng
func PrintUser(user models.User) {
	fmt.Printf("ID: %d, Tên: %s, Email: %s, Tuổi: %d\n", user.ID, user.Name, user.Email, user.Age)
}

// PrintUsers hiển thị thông tin của một danh sách người dùng
func PrintUsers(users []models.User) {
	if len(users) == 0 {
		fmt.Println("Danh sách người dùng trống!")
		return
	}

	for _, user := range users {
		PrintUser(user)
	}
}
