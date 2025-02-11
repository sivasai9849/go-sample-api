package postgres

import (
    "gorm.io/gorm"
    "github.com/sivasai9849/go-advanced-api/internal/domain"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepository{
        db: db,
    }
}

func (r *userRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, domain.ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, domain.ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
    result := r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(user)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return domain.ErrUserNotFound
    }
    return nil
}

func (r *userRepository) Delete(id string) error {
    return r.db.Delete(&domain.User{}, "id = ?", id).Error
}

func (r *userRepository) List(page, limit int) ([]domain.User, error) {
    var users []domain.User
    offset := (page - 1) * limit
    err := r.db.Offset(offset).Limit(limit).Find(&users).Error
    return users, err
}