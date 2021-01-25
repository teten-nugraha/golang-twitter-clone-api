package domain

import (
	"time"
)

type Tweet struct {
	ID                 uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserRefer          string
	Tweet              string
	LikeCount          int
	CommentCount       int
	CreatedAt          time.Time
	TweetConversations []TweetConversation `gorm:"foreignKey:TweetRefer"`
}
