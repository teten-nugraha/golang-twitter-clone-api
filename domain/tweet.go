package domain

import (
	"time"
)

type Tweet struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserRefer uint
	Tweet     string
	CreatedAt time.Time
}
