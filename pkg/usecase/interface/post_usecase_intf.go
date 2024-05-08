package interface_usecase_postnrel

import (
	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
)

type IPostUseCase interface {
	AddNewPost(data *[]*pb.SingleMedia, caption *string, userId *string) error

	GetAllPosts(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error)
	DeletePost(postId, userId *string) error
	EditPost(request *requestmodels_posnrel.EditPost) error
	LikePost(postId, userId *string) *error
	UnLikePost(postId, userId *string) *error
	GetMostLovedPostsFromGlobalUser(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error)
	GetAllRelatedPostsForHomeScreen(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error)
}
