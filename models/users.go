package models

import (
	"time"
)

type Users struct {
    Id       int64     `gorm:"primaryKey" json:"id"`
    IdLevel  int64     `json:"id_level"`
    IdToko   int64     `json:"id_toko"`
    Nama     string    `gorm:"type:varchar(255)" json:"nama" binding:"required"`
    Username string    `gorm:"type:varchar(255)" json:"username" binding:"required"`
    Password string    `gorm:"type:varchar(255)" json:"password" binding:"required,min=8,max=30"`
    Email    string    `gorm:"type:varchar(255)" json:"email" binding:"required,email"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
