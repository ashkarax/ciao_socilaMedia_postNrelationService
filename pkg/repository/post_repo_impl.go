package repository_postnrel

import (
	"errors"
	"fmt"
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

	query := "SELECT post_id,caption,created_at FROM posts WHERE user_id=$1 AND post_status=$2 ORDER BY created_at DESC LIMIT $3 OFFSET $4"
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
