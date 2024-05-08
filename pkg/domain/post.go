package domain_postNrelSvc

import "time"

type postStatus string

const (
	Normal    postStatus = "normal"
	Archieved postStatus = "archieved"
)

type Post struct {
	PostID uint `gorm:"primarykey"`

	UserID uint `gorm:"not null"`

	Caption string

	CreatedAt time.Time

	PostStatus postStatus `gorm:"default:normal"`
}

type PostMedia struct {
	MediaId uint `gorm:"primarykey"`
	
	PostID  uint `gorm:"not null"`

	Posts Post `gorm:"foreignkey:PostID"`

	MediaUrl string
}
