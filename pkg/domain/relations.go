package domain_postNrelSvc

type relationType string

const (
	Follows postStatus = "follows"
	Blocked postStatus = "blocked"
)

type Relationship struct {
	FollowerID uint `gorm:"not null"`

	FollowingID uint `gorm:"not null"`

	RelationType relationType `gorm:"default:follows"`

	UniqueConstraint struct {
		FollowerID  uint `gorm:"uniqueIndex:idx_follower_following"`
		FollowingID uint `gorm:"uniqueIndex:idx_follower_following"`
	} `gorm:"embedded;uniqueIndex:idx_follower_following"`
}
