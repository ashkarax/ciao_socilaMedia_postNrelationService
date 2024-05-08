package usecase_postnrel

import (
	"context"
	"errors"
	"fmt"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	interface_usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase/interface"
	interface_dateToAge "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/DateToAge/interface"
)

type CommentUseCase struct {
	CommentRepo   interface_repo_postnrel.ICommentRepo
	DateToAgeUtil interface_dateToAge.IDateToAge
	AuthClient    pb.AuthServiceClient
}

func NewCommentUseCase(commentRepo interface_repo_postnrel.ICommentRepo,
	dateToAgeUtil interface_dateToAge.IDateToAge,
	authClient *pb.AuthServiceClient) interface_usecase_postnrel.ICommentUseCase {
	return &CommentUseCase{
		CommentRepo:   commentRepo,
		DateToAgeUtil: dateToAgeUtil,
		AuthClient:    *authClient,
	}
}

func (r *CommentUseCase) AddNewComment(input *requestmodels_posnrel.CommentRequest) error {

	if input.ParentCommentId != 0 {
		//checking if this is a reply two another reply
		isReplyToReply, err := r.CommentRepo.CheckingCommentHierarchy(&input.ParentCommentId)
		if err != nil {
			fmt.Println("----", err)
			return err
		}
		if isReplyToReply {
			return errors.New("you can't reply to a comment-reply")
		}
	}

	err := r.CommentRepo.AddComment(input)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentUseCase) DeleteComment(userId, commentId *string) error {

	isParent, err := r.CommentRepo.DeleteCommentAndReturnIsParentStat(userId, commentId)
	if err != nil {
		return err
	}

	if isParent {
		err = r.CommentRepo.DeleteChildComments(commentId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CommentUseCase) EditComment(userId, commentText *string, commentId *uint64) error {

	err := r.CommentRepo.EditComment(userId, commentText, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentUseCase) FetchPostComments(userId, postId, limit, offset *string) (*[]responsemodels_postnrel.ParentComments, error) {
	parentComments, err := r.CommentRepo.FetchParentCommentsOfPost(userId, postId, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range *parentComments {
		childComments, err := r.CommentRepo.FetchChildCommentsOfComment(&(*parentComments)[i].CommentId)
		if err != nil {
			return nil, err
		}
		for j := range *childComments {
			context, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: fmt.Sprint((*childComments)[j].UserID)})
			if err != nil || userData.ErrorMessage != "" {
				return nil, errors.New(fmt.Sprint(err) + userData.ErrorMessage)
			}
			(*childComments)[j].UseName = userData.UserName
			(*childComments)[j].UserProfileImgURL = userData.UserProfileImgURL

			(*childComments)[j].CommentAge = *(r.DateToAgeUtil.DateTOAge(&(*childComments)[j].CreatedAt))
		}

		context, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: fmt.Sprint((*parentComments)[i].UserID)})
		if err != nil || userData.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprint(err) + userData.ErrorMessage)
		}
		(*parentComments)[i].UseName = userData.UserName
		(*parentComments)[i].UserProfileImgURL = userData.UserProfileImgURL

		(*parentComments)[i].CommentAge = *(r.DateToAgeUtil.DateTOAge(&(*parentComments)[i].CreatedAt))

		(*parentComments)[i].ChildComments = *childComments
	}

	return parentComments, nil
}
