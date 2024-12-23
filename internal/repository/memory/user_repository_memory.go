package memory

import (
    "errors"

    "clean-architecture-go/internal/domain"
    "clean-architecture-go/internal/repository"
)

type userRepositoryMemory struct {
    users []*domain.User
}

func NewUserRepositoryMemory() repository.UserRepository {
    return &userRepositoryMemory{
        users: make([]*domain.User, 0),
    }
}

func (r *userRepositoryMemory) Save(user *domain.User) error {
    user.ID = int64(len(r.users) + 1)
    r.users = append(r.users, user)
    return nil
}

func (r *userRepositoryMemory) FindByID(id int64) (*domain.User, error) {
    for _, u := range r.users {
        if u.ID == id {
            return u, nil
        }
    }
    return nil, errors.New("user not found")
}

func (r *userRepositoryMemory) FindAll() ([]*domain.User, error) {
    return r.users, nil
}
