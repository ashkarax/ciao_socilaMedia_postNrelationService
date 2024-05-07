package responsemodels_postnrel

import "time"

type PostData struct {
	UserId            uint
	UserName          string
	UserProfileImgURL string
	PostId            uint
	LikeStatus        bool
	Caption           string
	CreatedAt         time.Time
	LikesCount        string
	CommentsCount     string
	PostAge           string
	MediaUrl          []string
}
