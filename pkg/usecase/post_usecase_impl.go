package usecase_postnrel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	interface_usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase/interface"
	interface_awss3_postnrelations "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/aws_s3/interface"
)

type PostUseCase struct {
	PostRepo   interface_repo_postnrel.IPostRepo
	AWSUtil    interface_awss3_postnrelations.IAwsS3
	AuthClient pb.AuthServiceClient
}

func NewPostUseCase(postRepo interface_repo_postnrel.IPostRepo,
	awsUtil interface_awss3_postnrelations.IAwsS3,
	authClient *pb.AuthServiceClient) interface_usecase_postnrel.IPostUseCase {
	return &PostUseCase{PostRepo: postRepo,
		AWSUtil:    awsUtil,
		AuthClient: *authClient,
	}
}

func (r *PostUseCase) AddNewPost(data *[]*pb.SingleMedia, caption *string, userId *string) error {

	BucketFolder := "ciao-socialmedia/posts/"

	sess, err := r.AWSUtil.AWSSessionInitializer()
	if err != nil {
		fmt.Println(err)
		return err
	}
	var postData requestmodels_posnrel.AddPostData

	for i, file := range *data {
		mediaURL, err := r.AWSUtil.AWSS3MediaUploader(&file.Media, &file.ContentType, sess, &BucketFolder)
		if err != nil {
			fmt.Printf("Error uploading file %d: %v\n", i+1, err)
			return err
		}
		postData.MediaURLs = append(postData.MediaURLs, *mediaURL)
	}

	postData.Caption = caption
	postData.UserId = userId

	err = r.PostRepo.AddNewPost(&postData)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostUseCase) GetAllPosts(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	postData, err := r.PostRepo.GetAllActivePostByUser(userId, limit, offset)
	if err != nil {
		return nil, err
	}
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: *userId})
	if err != nil {
		log.Fatal(err)
	}
	if userData.ErrorMessage != "" {
		return nil, errors.New(userData.ErrorMessage)
	}

	for i, split := range *postData {
		(*postData)[i].UserName = userData.UserName
		(*postData)[i].UserProfileImgURL = userData.UserProfileImgURL

		postIdString := strconv.FormatUint(uint64(split.PostId), 10)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return nil, err
		}
		(*postData)[i].MediaUrl = *postMedias

		currentTime := time.Now()
		duration := currentTime.Sub((*postData)[i].CreatedAt)

		minutes := int(duration.Minutes())
		hours := int(duration.Hours())
		days := int(duration.Hours() / 24)
		months := int(duration.Hours() / 24 / 7)

		if minutes < 60 {
			(*postData)[i].PostAge = fmt.Sprintf("%d mins ago", minutes)
		} else if hours < 24 {
			(*postData)[i].PostAge = fmt.Sprintf("%d hrs ago", hours)
		} else if days < 30 {
			(*postData)[i].PostAge = fmt.Sprintf("%d dy ago", days)
		} else {
			(*postData)[i].PostAge = fmt.Sprintf("%d weks ago", months)
		}
	}

	return postData, nil
}

func (r *PostUseCase) DeletePost(postId, userId *string) error {

	err := r.PostRepo.DeletePostById(postId, userId)
	if err != nil {
		return err
	}
	err2 := r.PostRepo.DeletePostMedias(postId)
	if err2 != nil {
		return err2
	}

	return nil
}

func (r *PostUseCase) EditPost(request *requestmodels_posnrel.EditPost) error {

	err := r.PostRepo.EditPost(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
