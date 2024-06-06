package interface_repo_postnrel

import (
	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
)

type ICommentRepo interface {
	CheckingCommentHierarchy(input *uint64) (bool, error)
	AddComment(input *requestmodels_posnrel.CommentRequest) error
	DeleteCommentAndReturnIsParentStat(userId, commentId *string) (bool, error)
	DeleteChildComments(parentCommentId *string) error
	EditComment(userId, commentText *string, commentId *uint64) error
	FetchParentCommentsOfPost(userId, postId, limit, offset *string) (*[]responsemodels_postnrel.ParentComments, error)
	FetchChildCommentsOfComment(parentCommentId *uint) (*[]responsemodels_postnrel.ChildComments, error)
	FindCommentCreatorId(CommentId *uint64) (*string, error)
}
