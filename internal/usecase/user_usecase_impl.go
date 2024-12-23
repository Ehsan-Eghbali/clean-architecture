package usecase

import (
    "errors"

    "clean-architecture-go/internal/domain"
    "clean-architecture-go/internal/repository"
)

type userUseCaseImpl struct {
    userRepo repository.UserRepository
}

// NewUserUseCase پیاده‌سازی جدیدی از UserUseCase را برمی‌گرداند
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
    return &userUseCaseImpl{
        userRepo: userRepo,
    }
}

func (uc *userUseCaseImpl) CreateUser(user *domain.User) error {
    if user.Username == "" || user.Email == "" {
        return errors.New("username or email cannot be empty")
    }
    return uc.userRepo.Save(user)
}

func (uc *userUseCaseImpl) GetUserByID(id int64) (*domain.User, error) {
    return uc.userRepo.FindByID(id)
}

func (uc *userUseCaseImpl) GetAllUsers() ([]*domain.User, error) {
    return uc.userRepo.FindAll()
}
