package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at;"`
}
