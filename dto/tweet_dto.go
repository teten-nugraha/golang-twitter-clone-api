package dto

import "time"

type NewTweetDto struct {
	Tweet string `json:"tweet" validate:"required"`
	Maker string
}

type TweetDto struct {
	ID           uint64
	Tweet 		 string
	Maker        string
	LikeCount    int
	CommentCount int
	CreatedAt    time.Time
}
