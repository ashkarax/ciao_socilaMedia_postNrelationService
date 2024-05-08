package server_postNrelSvc

import (
	"context"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	interface_usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase/interface"
)

type PostNrelService struct {
	PostUseCase     interface_usecase_postnrel.IPostUseCase
	RelationUseCase interface_usecase_postnrel.IRelationUseCase
	CommentUseCase  interface_usecase_postnrel.ICommentUseCase
	pb.PostNrelServiceServer
}

func NewPostNrelServiceServer(postUseCase interface_usecase_postnrel.IPostUseCase,
	relationUseCase interface_usecase_postnrel.IRelationUseCase,
	commentUseCase interface_usecase_postnrel.ICommentUseCase) *PostNrelService {
	return &PostNrelService{
		PostUseCase:     postUseCase,
		RelationUseCase: relationUseCase,
		CommentUseCase:  commentUseCase,
	}
}

func (u *PostNrelService) AddNewPost(ctx context.Context, req *pb.RequestAddPost) (*pb.ResponseErrorMessageOnly, error) {

	err := u.PostUseCase.AddNewPost(&req.Media, &req.Caption, &req.UserId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) GetAllPostByUser(ctx context.Context, req *pb.RequestGetAllPosts) (*pb.ResponseUserPosts, error) {

	respData, err := u.PostUseCase.GetAllPosts(&req.UserId, &req.Limit, &req.OffSet)
	if err != nil {
		return &pb.ResponseUserPosts{
			ErrorMessage: err.Error(),
		}, nil
	}

	var repeatedData []*pb.PostsDataModel
	for i := range *respData {
		repeatedData = append(repeatedData, &pb.PostsDataModel{
			UserId:            uint64((*respData)[i].UserId),
			UserName:          (*respData)[i].UserName,
			UserProfileImgURL: (*respData)[i].UserProfileImgURL,
			PostId:            uint64((*respData)[i].PostId),
			LikeStatus:        (*respData)[i].IsLiked,
			Caption:           (*respData)[i].Caption,
			LikesCount:        uint64((*respData)[i].LikesCount),
			CommentsCount:     uint64((*respData)[i].CommentsCount),
			PostAge:           (*respData)[i].PostAge,
			MediaUrl:          (*respData)[i].MediaUrl,
		})
	}

	return &pb.ResponseUserPosts{
		PostsData: repeatedData,
	}, nil
}

func (u *PostNrelService) DeletePost(ctx context.Context, req *pb.RequestDeletePost) (*pb.ResponseErrorMessageOnly, error) {

	err := u.PostUseCase.DeletePost(&req.PostId, &req.UserId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ResponseErrorMessageOnly{}, nil

}

func (u *PostNrelService) EditPost(ctx context.Context, req *pb.RequestEditPost) (*pb.ResponseErrorMessageOnly, error) {

	var request requestmodels_posnrel.EditPost

	request.UserId = req.UserId
	request.PostId = req.PostId
	request.Caption = req.Caption

	err := u.PostUseCase.EditPost(&request)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) Follow(ctx context.Context, req *pb.RequestFollowUnFollow) (*pb.ResponseErrorMessageOnly, error) {

	err := u.RelationUseCase.Follow(&req.UserId, &req.UserBId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: (*err).Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}
func (u *PostNrelService) UnFollow(ctx context.Context, req *pb.RequestFollowUnFollow) (*pb.ResponseErrorMessageOnly, error) {

	err := u.RelationUseCase.UnFollow(&req.UserId, &req.UserBId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: (*err).Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) GetCountsForUserProfile(ctx context.Context, req *pb.RequestUserIdPnR) (*pb.ResponseGetCounts, error) {

	followersCount, followingCount, postsCount, err := u.RelationUseCase.GetCountsForUserProfile(&req.UserId)
	if err != nil {
		return &pb.ResponseGetCounts{
			ErrorMessage: (*err).Error(),
		}, nil
	}

	return &pb.ResponseGetCounts{
		PostCount:      uint64(*postsCount),
		FollowerCount:  uint64(*followersCount),
		FollowingCount: uint64(*followingCount),
	}, nil

}

func (u *PostNrelService) GetFollowersIds(ctx context.Context, req *pb.RequestUserIdPnR) (*pb.ResposneGetUsersIds, error) {

	userIdSlce, err := u.RelationUseCase.GetFollowersIds(&req.UserId)
	if err != nil {
		return &pb.ResposneGetUsersIds{
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ResposneGetUsersIds{
		UserIds: *userIdSlce,
	}, nil
}
func (u *PostNrelService) GetFollowingsIds(ctx context.Context, req *pb.RequestUserIdPnR) (*pb.ResposneGetUsersIds, error) {

	userIdSlce, err := u.RelationUseCase.GetFollowingsIds(&req.UserId)
	if err != nil {
		return &pb.ResposneGetUsersIds{
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ResposneGetUsersIds{
		UserIds: *userIdSlce,
	}, nil
}

func (u *PostNrelService) LikePost(ctx context.Context, req *pb.RequestLikeUnlikePost) (*pb.ResponseErrorMessageOnly, error) {
	err := u.PostUseCase.LikePost(&req.PostId, &req.UserId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: (*err).Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil

}

func (u *PostNrelService) UnLikePost(ctx context.Context, req *pb.RequestLikeUnlikePost) (*pb.ResponseErrorMessageOnly, error) {
	err := u.PostUseCase.UnLikePost(&req.PostId, &req.UserId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: (*err).Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil

}

func (u *PostNrelService) UserAFollowingUserBorNot(ctx context.Context, req *pb.RequestFollowUnFollow) (*pb.ResponseUserABrelation, error) {
	resp, err := u.RelationUseCase.UserAFollowingUserBorNot(&req.UserId, &req.UserBId)
	if err != nil {
		return &pb.ResponseUserABrelation{
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb.ResponseUserABrelation{
		BoolStat: resp,
	}, nil
}

func (u *PostNrelService) AddComment(ctx context.Context, req *pb.RequestAddComment) (*pb.ResponseErrorMessageOnly, error) {
	var input requestmodels_posnrel.CommentRequest

	input.UserId = req.UserId
	input.PostId = req.PostId
	input.CommentText = req.CommentText
	input.ParentCommentId = req.ParentCommentId

	err := u.CommentUseCase.AddNewComment(&input)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) DeleteComment(ctx context.Context, req *pb.RequestCommentDelete) (*pb.ResponseErrorMessageOnly, error) {

	err := u.CommentUseCase.DeleteComment(&req.UserId, &req.CommentId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) EditComment(ctx context.Context, req *pb.RequestEditComment) (*pb.ResponseErrorMessageOnly, error) {
	err := u.CommentUseCase.EditComment(&req.UserId, &req.CommentText, &req.CommentId)
	if err != nil {
		return &pb.ResponseErrorMessageOnly{
			ErrorMessage: err.Error(),
		}, nil
	}
	return &pb.ResponseErrorMessageOnly{}, nil
}

func (u *PostNrelService) FetchPostComments(ctx context.Context, req *pb.RequestFetchComments) (*pb.ResponseFetchComments, error) {
	respData, err := u.CommentUseCase.FetchPostComments(&req.UserId, &req.PostId, &req.Limit, &req.OffSet)
	if err != nil {
		return &pb.ResponseFetchComments{
			ErrorMessage: err.Error(),
		}, nil
	}

	var ChildComments []*pb.ChildComments
	var ParentComments []*pb.ParentComments

	for i := range *respData {
		ChildComments = []*pb.ChildComments{}
		for j := range (*respData)[i].ChildComments {
			ChildComments = append(ChildComments, &pb.ChildComments{
				CommentId:         uint64(((*respData)[i].ChildComments)[j].CommentId),
				PostId:            uint64(((*respData)[i].ChildComments)[j].PostID),
				UserId:            uint64(((*respData)[i].ChildComments)[j].UserID),
				UserName:          ((*respData)[i].ChildComments)[j].UseName,
				UserProfileImgURL: ((*respData)[i].ChildComments)[j].UserProfileImgURL,
				CommentText:       ((*respData)[i].ChildComments)[j].CommentText,
				ParentCommentID:   uint64(((*respData)[i].ChildComments)[j].ParentCommentID),
				CommentAge:        ((*respData)[i].ChildComments)[j].CommentAge,
			})
		}
		ParentComments = append(ParentComments, &pb.ParentComments{
			CommentId:         uint64((*respData)[i].CommentId),
			PostId:            uint64((*respData)[i].PostID),
			UserId:            uint64((*respData)[i].UserID),
			UserName:          (*respData)[i].UseName,
			UserProfileImgURL: (*respData)[i].UserProfileImgURL,
			CommentText:       (*respData)[i].CommentText,
			CommentAge:        (*respData)[i].CommentAge,
			ChildCommentCount: uint64(len(ChildComments)),
			ChildComments:     ChildComments,
		})

	}

	return &pb.ResponseFetchComments{
		ParentCommentsCount: uint64(len(ParentComments)),
		ParentComments:      ParentComments,
	}, nil
}

func (u *PostNrelService) GetMostLovedPostsFromGlobalUser(ctx context.Context, req *pb.RequestGetAllPosts) (*pb.ResponseUserPosts, error) {
	respData, err := u.PostUseCase.GetMostLovedPostsFromGlobalUser(&req.UserId, &req.Limit, &req.OffSet)
	if err != nil {
		return &pb.ResponseUserPosts{
			ErrorMessage: err.Error(),
		}, nil
	}

	var repeatedData []*pb.PostsDataModel
	for i := range *respData {
		repeatedData = append(repeatedData, &pb.PostsDataModel{
			UserId:            uint64((*respData)[i].UserId),
			UserName:          (*respData)[i].UserName,
			UserProfileImgURL: (*respData)[i].UserProfileImgURL,
			PostId:            uint64((*respData)[i].PostId),
			LikeStatus:        (*respData)[i].IsLiked,
			Caption:           (*respData)[i].Caption,
			LikesCount:        uint64((*respData)[i].LikesCount),
			CommentsCount:     uint64((*respData)[i].CommentsCount),
			PostAge:           (*respData)[i].PostAge,
			MediaUrl:          (*respData)[i].MediaUrl,
		})
	}

	return &pb.ResponseUserPosts{
		PostsData: repeatedData,
	}, nil
}

func (u *PostNrelService) GetAllRelatedPostsForHomeScreen(ctx context.Context, req *pb.RequestGetAllPosts) (*pb.ResponseUserPosts, error) {
	respData, err := u.PostUseCase.GetAllRelatedPostsForHomeScreen(&req.UserId, &req.Limit, &req.OffSet)
	if err != nil {
		return &pb.ResponseUserPosts{
			ErrorMessage: err.Error(),
		}, nil
	}

	var repeatedData []*pb.PostsDataModel
	for i := range *respData {
		repeatedData = append(repeatedData, &pb.PostsDataModel{
			UserId:            uint64((*respData)[i].UserId),
			UserName:          (*respData)[i].UserName,
			UserProfileImgURL: (*respData)[i].UserProfileImgURL,
			PostId:            uint64((*respData)[i].PostId),
			LikeStatus:        (*respData)[i].IsLiked,
			Caption:           (*respData)[i].Caption,
			LikesCount:        uint64((*respData)[i].LikesCount),
			CommentsCount:     uint64((*respData)[i].CommentsCount),
			PostAge:           (*respData)[i].PostAge,
			MediaUrl:          (*respData)[i].MediaUrl,
		})
	}

	return &pb.ResponseUserPosts{
		PostsData: repeatedData,
	}, nil
}
