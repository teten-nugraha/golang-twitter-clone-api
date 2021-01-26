package dto

import "time"

type CommentDto struct {
	ID        uint64
	Maker     string
	TweetId   uint64 `json:"tweet_id" validate:"required"`
	Reply     string `json:"reply" validate:"required"`
	CreatedAt time.Time
}
