package models

import "gorm.io/gorm"

type Bioskop struct {
	gorm.Model
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}
