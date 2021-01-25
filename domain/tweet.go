package domain

import (
	"time"
)

type Tweet struct {
	ID                 uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserRefer          string
	Tweet              string
	CreatedAt          time.Time
	TweetConversations []TweetConversation `gorm:"foreignKey:TweetRefer"`
}
