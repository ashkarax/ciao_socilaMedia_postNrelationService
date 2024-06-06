package repository_postnrel

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) interface_repo_postnrel.IPostRepo {
	return &PostRepo{DB: db}
}

func (d *PostRepo) AddNewPost(postData *requestmodels_posnrel.AddPostData) error {
	var PostId string

	query := "INSERT INTO Posts (user_id,caption,created_at) VALUES ($1,$2,$3) RETURNING post_id"
	err := d.DB.Raw(query, postData.UserId, postData.Caption, time.Now()).Scan(&PostId).Error
	if err != nil {
		return err
	}

	mediaInsQuery := "INSERT INTO post_media (post_id,media_url) VALUES ($1,$2)"

	for _, url := range postData.MediaURLs {
		errIns := d.DB.Exec(mediaInsQuery, PostId, url).Error
		if errIns != nil {
			return errIns
		}
	}

	return nil
}

func (d *PostRepo) GetAllActivePostByUser(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	var response []responsemodels_postnrel.PostData

	query := "SELECT posts.*,CASE WHEN EXISTS (SELECT 1 FROM post_likes WHERE post_likes.post_id = posts.post_id AND post_likes.user_id = $1) THEN TRUE ELSE FALSE END AS is_liked FROM posts LEFT JOIN post_likes ON posts.post_id = post_likes.post_id  WHERE posts.user_id=$1 AND posts.post_status = $2 GROUP BY posts.post_id ORDER BY posts.created_at DESC LIMIT $3 OFFSET $4;"
	err := d.DB.Raw(query, *userId, "normal", *limit, *offset).Scan(&response).Error
	if err != nil {
		fmt.Println("------", err)
		return nil, err
	}
	return &response, nil
}

func (d *PostRepo) GetPostMediaById(postId *string) (*[]string, error) {
	var response []string

	query := "SELECT media_url FROM post_media WHERE post_id=$1 ORDER BY media_id DESC"
	err := d.DB.Raw(query, *postId).Scan(&response).Error
	if err != nil {
		return &response, err
	}

	return &response, nil
}

func (d *PostRepo) DeletePostById(postId, userId *string) error {

	err := d.RemovePostLikesByPostId(postId)
	if err != nil {
		return err
	}
	query := "DELETE FROM posts WHERE post_id=$1 AND user_id=$2"
	res := d.DB.Exec(query, postId, userId)
	if res.RowsAffected == 0 {
		return errors.New("enter a valid post id,rows affected 0")
	}
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *PostRepo) RemovePostLikesByPostId(postId *string) error {
	query := "DELETE FROM post_likes WHERE post_id=$1"
	err := d.DB.Exec(query, postId).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PostRepo) DeletePostMedias(postId *string) error {
	query := "DELETE FROM post_media WHERE post_id=$1"
	res := d.DB.Exec(query, postId).Error
	if res != nil {
		return res
	}
	return nil

}

func (d *PostRepo) EditPost(inputData *requestmodels_posnrel.EditPost) error {
	query := "UPDATE posts SET caption=$1 WHERE post_id=$2 AND user_id=$3;"
	res := d.DB.Exec(query, inputData.Caption, inputData.PostId, inputData.UserId)
	if res.RowsAffected == 0 {
		return errors.New("enter a valid post id")
	}
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *PostRepo) GetPostCountOfUser(userId *string) (*uint, *error) {
	var count uint
	query := "SELECT COUNT(*) FROM posts WHERE user_id=$1 AND post_status=$2"
	if err := d.DB.Raw(query, userId, "normal").Scan(&count).Error; err != nil {
		return nil, &err
	}
	return &count, nil

}

func (d *PostRepo) LikePost(postId, userId *string) (bool, error) {
	var inserted bool
	query := "INSERT INTO post_likes (user_id,post_id,created_at) VALUES (?,?,?) ON CONFLICT (user_id, post_id) DO NOTHING RETURNING (xmax = 0);"
	err := d.DB.Raw(query, userId, postId, time.Now()).Scan(&inserted).Error
	if err != nil && err != sql.ErrNoRows {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return false, errors.New("enter a valid postid: Post not found")
		}
		fmt.Println("----------", err)
		return false, err
	}
	return inserted, nil
}

func (d *PostRepo) UnLikePost(postId, userId *string) error {
	query := "DELETE FROM post_likes WHERE user_id=? AND post_id=?"
	err := d.DB.Exec(query, userId, postId).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PostRepo) GetPostLikeAndCommentsCount(postId *string) (*responsemodels_postnrel.LikeCommentCounts, error) {
	var ScannerStruct responsemodels_postnrel.LikeCommentCounts

	query := "SELECT (SELECT COUNT(*) FROM post_likes WHERE post_id = $1) AS likes_count,(SELECT COUNT(*) FROM comments WHERE post_id = $1 AND parent_comment_id = 0) AS comments_count;"
	err := d.DB.Raw(query, postId).Scan(&ScannerStruct).Error
	if err != nil {
		return nil, err
	}
	return &ScannerStruct, nil
}

func (d *PostRepo) GetMostLovedPostsFromGlobalUser(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	var response []responsemodels_postnrel.PostData

	query := "SELECT posts.*,COUNT(post_likes.post_id) AS like_count,CASE WHEN EXISTS (SELECT 1 FROM post_likes WHERE post_likes.post_id = posts.post_id AND post_likes.user_id = $1) THEN TRUE ELSE FALSE END AS is_liked FROM posts LEFT JOIN post_likes ON posts.post_id = post_likes.post_id WHERE posts.post_status = 'normal' GROUP BY posts.post_id ORDER BY like_count DESC,posts.created_at DESC LIMIT $2 OFFSET $3"
	err := d.DB.Raw(query, *userId, *limit, *offset).Scan(&response).Error
	if err != nil {
		fmt.Println("------", err)
		return nil, err
	}
	return &response, nil
}

func (d *PostRepo) GetAllActiveRelatedPostsForHomeScreen(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	var response []responsemodels_postnrel.PostData

	query := "SELECT posts.*,CASE WHEN post_likes.user_id IS NULL THEN FALSE ELSE TRUE END AS is_liked FROM posts INNER JOIN relationships ON posts.user_id =relationships.following_id LEFT JOIN (SELECT post_id, user_id FROM post_likes WHERE user_id = $1) AS post_likes ON posts.post_id = post_likes.post_id  WHERE relationships.follower_id = $1 AND posts.post_status = 'normal' ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3"
	err := d.DB.Raw(query, userId, limit, offset).Scan(&response)
	if err.Error != nil {
		return &response, err.Error
	}
	return &response, nil

}

func (d *PostRepo) GetPostCreatorId(postId *string) (*string, error) {
	var id string
	query := "SELECT user_id FROM posts WHERE post_id=?"
	result := d.DB.Raw(query, postId).Scan(&id)
	if result.RowsAffected == 0 {
		return nil, errors.New("no post found with this id,enter a valid postid")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &id, nil
}
