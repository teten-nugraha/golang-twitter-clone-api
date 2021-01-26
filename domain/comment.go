package domain

import "time"

type Comment struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Maker     string
	TweetId   uint64
	Reply     string
	CreatedAt time.Time
}
