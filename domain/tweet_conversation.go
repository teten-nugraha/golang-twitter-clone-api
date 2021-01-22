package domain

import "time"

type TweetConversation struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserRefer  uint
	TweetRefer uint
	Reply      string
	CreatedAt  time.Time
}
