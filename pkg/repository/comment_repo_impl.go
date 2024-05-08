package repository_postnrel

import (
	"errors"
	"strings"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	"gorm.io/gorm"
)

type CommentRepo struct {
	DB *gorm.DB
}

func NewCommentRepo(db *gorm.DB) interface_repo_postnrel.ICommentRepo {
	return &CommentRepo{DB: db}
}

func (d *CommentRepo) CheckingCommentHierarchy(input *uint64) (bool, error) {
	var parentCommentId uint64
	query := "SELECT parent_comment_id FROM comments WHERE comment_id=?"
	err := d.DB.Raw(query, input).Scan(&parentCommentId)

	if err.Error != nil {
		return true, err.Error
	}
	if err.RowsAffected == 0 {
		return true, errors.New("no parent-comment found in this id,enter a valid parent-comment-id")
	}

	if parentCommentId != 0 {
		return true, nil
	}

	return false, nil

}

func (d *CommentRepo) AddComment(input *requestmodels_posnrel.CommentRequest) error {

	query := "INSERT INTO comments (post_id,user_id,comment_text,created_at) VALUES ($1,$2,$3,$4)"
	if input.ParentCommentId != 0 {
		query = "INSERT INTO comments (post_id,user_id,comment_text,parent_comment_id,created_at) VALUES ($1,$2,$3,$4,$5)"
	}

	var err error
	if input.ParentCommentId == 0 {
		err = d.DB.Exec(query, input.PostId, input.UserId, input.CommentText, time.Now()).Error
	} else {
		err = d.DB.Exec(query, input.PostId, input.UserId, input.CommentText, input.ParentCommentId, time.Now()).Error
	}

	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return errors.New("foreign key constraint violation: Post not found")
		}

		return err
	}
	return nil
}

func (d *CommentRepo) DeleteCommentAndReturnIsParentStat(userId, commentId *string) (bool, error) {
	var parentCommentId uint64
	query := "DELETE FROM comments WHERE user_id=$1 AND comment_id=$2 RETURNING parent_comment_id"
	result := d.DB.Raw(query, userId, commentId).Scan(&parentCommentId)
	if result.RowsAffected == 0 {
		return false, errors.New("no comment found with this id ,enter a valid commentid")
	}
	if result.Error != nil {
		return false, result.Error
	}

	if parentCommentId != 0 {
		return false, nil
	}

	return true, nil
}

func (d *CommentRepo) DeleteChildComments(parentCommentId *string) error {
	query := "DELETE FROM comments WHERE parent_comment_id=?"
	err := d.DB.Exec(query, parentCommentId).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *CommentRepo) EditComment(userId, commentText *string, commentId *uint64) error {
	query := "UPDATE comments SET comment_text=$1 WHERE user_id=$2 AND comment_id=$3"
	result := d.DB.Exec(query, commentText, userId, commentId)
	if result.RowsAffected == 0 {
		return errors.New("no comment found with this id,enter a valid commentId ")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *CommentRepo) FetchParentCommentsOfPost(userId, postId, limit, offset *string) (*[]responsemodels_postnrel.ParentComments, error) {
	var ScannerStruct []responsemodels_postnrel.ParentComments

	query := "SELECT * FROM comments WHERE user_id=$1 AND post_id=$2 AND parent_comment_id=0 LIMIT $3 OFFSET $4"
	result := d.DB.Raw(query, userId, postId, limit, offset).Scan(&ScannerStruct).Error
	if result != nil {
		return nil, result
	}
	return &ScannerStruct, nil
}

func (d *CommentRepo) FetchChildCommentsOfComment(parentCommentId *uint) (*[]responsemodels_postnrel.ChildComments, error) {
	var ScannerStruct []responsemodels_postnrel.ChildComments
	query := "SELECT * FROM comments WHERE parent_comment_id=$1"
	err := d.DB.Raw(query, parentCommentId).Scan(&ScannerStruct).Error
	if err != nil {
		return nil, err
	}
	return &ScannerStruct, nil
}

