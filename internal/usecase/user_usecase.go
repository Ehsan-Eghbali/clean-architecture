package usecase

import "clean-architecture-go/internal/domain"

// UserUseCase اینترفیس مربوط به عملیات اصلی روی کاربران
type UserUseCase interface {
    CreateUser(user *domain.User) error
    GetUserByID(id int64) (*domain.User, error)
    GetAllUsers() ([]*domain.User, error)
}
