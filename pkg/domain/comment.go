package domain_postNrelSvc

import "time"

type Comment struct {
	CommentId uint `gorm:"primarykey"`

	PostID uint `gorm:"not null"`
	Posts  Post `gorm:"foreignkey:PostID"`

	UserID uint `gorm:"not null"`

	CommentText string `gorm:"not null"`

	ParentCommentID uint `gorm:"default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}
