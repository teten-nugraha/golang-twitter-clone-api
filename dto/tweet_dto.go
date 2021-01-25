package dto

import "time"

type NewTweetDto struct {
	Tweet string `json:"tweet" validate:"required"`
	Maker string
}

type TweetDto struct {
	ID        uint64
	Maker     string
	CreatedAt time.Time
}
