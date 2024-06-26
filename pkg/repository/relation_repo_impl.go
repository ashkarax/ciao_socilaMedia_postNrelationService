package repository_postnrel

import (
	"database/sql"

	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	"gorm.io/gorm"
)

type RelationRepo struct {
	DB *gorm.DB
}

func NewRelationRepo(db *gorm.DB) interface_repo_postnrel.IRelationRepo {
	return &RelationRepo{DB: db}
}

func (d *RelationRepo) GetFollowerAndFollowingCountofUser(userId *string) (*uint, *uint, *error) {
	var counts struct {
		FollowersCount uint `gorm:"column:followers_count"`
		FollowingCount uint `gorm:"column:following_count"`
	}
	query := "SELECT (SELECT COUNT(*) FROM relationships WHERE following_id = $1 AND relation_type=$2) AS followers_count,(SELECT COUNT(*) FROM relationships WHERE follower_id = $1 AND relation_type=$2) AS following_count "
	err := d.DB.Raw(query, userId, "follows").Scan(&counts).Error
	if err != nil {
		return nil, nil, &err
	}
	return &counts.FollowersCount, &counts.FollowingCount, nil
}

func (d *RelationRepo) InitiateFollowRelationship(userId, userBId *string) (bool, error) {
	var inserted bool

	query := "INSERT INTO relationships (follower_id, following_id) VALUES ($1, $2) ON CONFLICT (follower_id, following_id) DO NOTHING RETURNING (xmax = 0);"
	err := d.DB.Raw(query, userId, userBId).Scan(&inserted).Error
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return inserted, nil
}

func (d *RelationRepo) InitiateUnFollowRelationship(userId, userBId *string) *error {

	query := "DELETE FROM relationships WHERE follower_id=$1 AND following_id=$2 AND relation_type=$3"
	err := d.DB.Exec(query, userId, userBId, "follows").Error
	if err != nil {
		return &err
	}
	return nil

}

func (d *RelationRepo) GetFollowersIdsOfUser(userId *string) (*[]uint64, error) {
	var userIds []uint64

	query := "SELECT follower_id FROM relationships WHERE following_id=$1 AND relation_type=$2"
	err := d.DB.Raw(query, userId, "follows").Scan(&userIds).Error
	if err != nil {
		return nil, err
	}
	return &userIds, nil
}

func (d *RelationRepo) GetFollowingsIdsOfUser(userId *string) (*[]uint64, error) {
	var userIds []uint64

	query := "SELECT following_id FROM relationships WHERE follower_id=$1 AND relation_type=$2"
	err := d.DB.Raw(query, userId, "follows").Scan(&userIds).Error
	if err != nil {
		return nil, err
	}
	return &userIds, nil
}

func (d *RelationRepo) UserAFollowingUserBorNot(userId, userBId *string) (bool, error) {
	var count uint

	query := "SELECT COUNT(*) FROM relationships WHERE follower_id = ? AND following_id = ? AND relation_type=?"
	err := d.DB.Raw(query, userId, userBId, "follows").Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}
