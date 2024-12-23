package repository

import "clean-architecture-go/internal/domain"

// UserRepository اینترفیس مربوط به عملیات دیتابیس یا هر نوع ذخیره‌سازی برای User
type UserRepository interface {
    Save(user *domain.User) error
    FindByID(id int64) (*domain.User, error)
    FindAll() ([]*domain.User, error)
}
