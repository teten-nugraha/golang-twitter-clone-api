package domain

type User struct {
	ID                 uint64              `gorm:"primaryKey;autoIncrement:true"`
	Username           string              `gorm:"type:varchar(255);NOT NULL" json:"username" binding:"required"`
	Email              string              `gorm:"type:varchar(255);NOT NULL" json:"email" binding:"required"`
	Gender             string              `gorm:"type:varchar(255);NOT NULL" json:"gender" binding:"required"`
	Password           string              `gorm:"type:varchar(255);NOT NULL" json:"password" binding:"required"`
	Tweets             []Tweet             `gorm:"foreignKey:UserRefer"`
}
