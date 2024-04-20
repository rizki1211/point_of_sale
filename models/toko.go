package models

import "time"

type Toko struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	NamaToko  string `json:"nama_toko" binding:"required"`
	Alamat    string `json:"alamat" binding:"required,max=250"`
	CreatedAt time.Time
	UpdatedAt time.Time
}