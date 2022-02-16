package datamodels

import "time"

type Model struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
}
