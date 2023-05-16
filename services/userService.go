package services

import (
	"errors"
	"server/db"
	"server/dto"

	"gorm.io/gorm"
)

type UserService struct {
	conn *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{conn: db}
}

func (s *UserService) GetUsers() []db.User {
	var users []db.User
	s.conn.Find(&users)
	return users
}

func (s *UserService) GetUser(id int) (db.User, error) {
	var user db.User
	if err := s.conn.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) CreateUser(userDto dto.CreateUserDto) (db.User, error) {
	user := db.User{Username: userDto.Username, Role: userDto.Role, Password: userDto.Password}
	err := s.conn.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) {
	s.conn.Delete(&db.User{}, id)
}

func (s *UserService) SignIn(userRequest dto.SignInDto) (db.User, error) {
	var user db.User
	if err := s.conn.Where("username = ?", userRequest.Username).First(&user).Error; err != nil {
		return user, err
	}

	if user.Password != userRequest.Password {
		return user, errors.New("Wrong password")
	}

	return user, nil
}

func (s *UserService) CurrentUser(id int) (db.User, error) {
	var user db.User
	if err := s.conn.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
