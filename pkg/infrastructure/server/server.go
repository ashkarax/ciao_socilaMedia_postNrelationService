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
	pb.PostNrelServiceServer
}

func NewPostNrelServiceServer(postUseCase interface_usecase_postnrel.IPostUseCase,
	relationUseCase interface_usecase_postnrel.IRelationUseCase) *PostNrelService {
	return &PostNrelService{
		PostUseCase:     postUseCase,
		RelationUseCase: relationUseCase,
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
			LikeStatus:        (*respData)[i].LikeStatus,
			Caption:           (*respData)[i].Caption,
			LikesCount:        (*respData)[i].LikesCount,
			CommentsCount:     (*respData)[i].CommentsCount,
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

//<<<<<<<<<<<<<--------------------FROM AUTH SERVICE---------------->>>>>>>>>>>>>>>>>>>

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
