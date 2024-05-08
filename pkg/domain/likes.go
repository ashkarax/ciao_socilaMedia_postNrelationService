package domain_postNrelSvc

import "time"

type PostLikes struct {
	UserID uint `gorm:"not null"`

	PostID uint `gorm:"not null"`
	Posts  Post `gorm:"foreignkey:PostID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UniqueConstraint struct {
		UserID uint `gorm:"uniqueIndex:idx_user_post"`
		PostID uint `gorm:"uniqueIndex:idx_user_post"`
	} `gorm:"embedded;uniqueIndex:idx_user_post"`
}
