package responsemodels_postnrel

import "time"

type PostData struct {
	UserId            uint
	UserName          string
	UserProfileImgURL string
	PostId            uint
	Caption           string
	CreatedAt         time.Time
	PostAge           string
	MediaUrl          []string `gorm:"-"`

	IsLiked       bool
	LikesCount    uint
	CommentsCount uint
}

type LikeCommentCounts struct {
	LikesCount    uint `gorm:"column:likes_count"`
	CommentsCount uint `gorm:"column:comments_count"`
}
