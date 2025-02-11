package service

import (
    "github.com/sivasai9849/go-advanced-api/internal/domain"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

func (s *UserService) Create(user *domain.User) error {
    existing, err := s.userRepo.GetByEmail(user.Email)
    if err != nil && err != domain.ErrUserNotFound {
        return err
    }
    if existing != nil {
        return domain.ErrUserExists
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    return s.userRepo.Create(user)
}

func (s *UserService) GetByID(id string) (*domain.User, error) {
    return s.userRepo.GetByID(id)
}

func (s *UserService) List(page, limit int) ([]domain.User, error) {
    return s.userRepo.List(page, limit)
}
func (s *UserService) Update(user *domain.User) error {
    existing, err := s.userRepo.GetByID(user.ID.String())
    if err != nil {
        return err
    }

    if existing.Email != user.Email {
        emailCheck, err := s.userRepo.GetByEmail(user.Email)
        if err != nil && err != domain.ErrUserNotFound {
            return err
        }
        if emailCheck != nil {
            return domain.ErrUserExists
        }
    }

    if user.Password != existing.Password {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        user.Password = string(hashedPassword)
    }

    return s.userRepo.Update(user)
}

func (s *UserService) Delete(id string) error {
    return s.userRepo.Delete(id)
}
