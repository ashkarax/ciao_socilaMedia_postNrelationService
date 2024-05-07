package interface_repo_postnrel

import (
	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
)

type IPostRepo interface {
	AddNewPost(postData *requestmodels_posnrel.AddPostData) error

	GetAllActivePostByUser(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error)
	GetPostMediaById(postId *string) (*[]string, error)
	DeletePostById(postId, userId *string) error
	DeletePostMedias(postId *string) error
	EditPost(inputData *requestmodels_posnrel.EditPost) error

	GetPostCountOfUser(userId *string) (*uint, *error)
}
