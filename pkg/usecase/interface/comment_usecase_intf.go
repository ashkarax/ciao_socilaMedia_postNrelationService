package interface_usecase_postnrel

import (
	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
)

type ICommentUseCase interface {
	AddNewComment(input *requestmodels_posnrel.CommentRequest) error
	DeleteComment(userId, commentId *string) error
	EditComment(userId, commentText *string, commentId *uint64) error
	FetchPostComments(userId, postId, limit, offset *string) (*[]responsemodels_postnrel.ParentComments, error)
}
