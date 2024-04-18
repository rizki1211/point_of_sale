package models

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	IdLevel  int64  `json:"id_level"`
	Nama     string `gorm:"type:varchar(255)" json:"nama" validate:"required"`
	Username string `gorm:"type:varchar(255)" json:"username" validate:"required"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required"`
}