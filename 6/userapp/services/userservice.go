package services

import (
	"userapp/models"
)

// UserService cung cấp các phương thức để thao tác với người dùng
type UserService struct {
	users map[int]models.User
}

// NewUserService tạo một dịch vụ người dùng mới
func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]models.User),
	}
}

// Add thêm một người dùng mới
func (s *UserService) Add(user models.User) {
	s.users[user.ID] = user
}

// GetByID tìm người dùng theo ID
func (s *UserService) GetByID(id int) (models.User, bool) {
	user, found := s.users[id]
	return user, found
}

// GetAll trả về tất cả người dùng
func (s *UserService) GetAll() []models.User {
	userList := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return userList
}

// Update cập nhật thông tin người dùng
func (s *UserService) Update(user models.User) {
	if _, exists := s.users[user.ID]; exists {
		s.users[user.ID] = user
	}
}

// Delete xóa người dùng theo ID
func (s *UserService) Delete(id int) {
	delete(s.users, id)
}
