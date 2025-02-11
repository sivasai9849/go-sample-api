package domain

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    Name      string    `json:"name" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}