package entity

import (
	"time"
)

type Ui struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Nome      string    `json:"nome" binding:"required" gorm:"type:varchar(32)"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256);UNIQUE"`
	Senha     string    `json:"senha" binding:"min=8,max=24,required" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
