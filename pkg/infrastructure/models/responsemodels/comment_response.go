package responsemodels_postnrel

import "time"

type ChildComments struct {
	CommentId         uint
	PostID            uint
	UserID            uint
	UseName           string `gorm:"-"`
	UserProfileImgURL string `gorm:"-"`
	CommentText       string
	ParentCommentID   uint
	CreatedAt         time.Time
	CommentAge        string `gorm:"-"`
}

type ParentComments struct {
	CommentId          uint
	PostID             uint
	UserID             uint
	UseName            string `gorm:"-"`
	UserProfileImgURL  string `gorm:"-"`
	CommentText        string
	ParentCommentID    uint
	CreatedAt          time.Time
	CommentAge         string          `gorm:"-"`
	ChildCommentsCount uint            `gorm:"-"`
	ChildComments      []ChildComments `gorm:"-"`
}
