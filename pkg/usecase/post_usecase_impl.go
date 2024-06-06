package usecase_postnrel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	responsemodels_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	interface_usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase/interface"
	interface_dateToAge "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/DateToAge/interface"
	interface_awss3_postnrelations "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/aws_s3/interface"
	interface_kafkaproducer "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/kafka_producer/interface"
)

type PostUseCase struct {
	PostRepo      interface_repo_postnrel.IPostRepo
	AWSUtil       interface_awss3_postnrelations.IAwsS3
	DateToAgeUtil interface_dateToAge.IDateToAge
	AuthClient    pb.AuthServiceClient
	KafkaProducer interface_kafkaproducer.IKafkaProducer
}

func NewPostUseCase(postRepo interface_repo_postnrel.IPostRepo,
	awsUtil interface_awss3_postnrelations.IAwsS3,
	dateToAgeUtil interface_dateToAge.IDateToAge,
	authClient *pb.AuthServiceClient,
	kafkaProducer interface_kafkaproducer.IKafkaProducer) interface_usecase_postnrel.IPostUseCase {
	return &PostUseCase{PostRepo: postRepo,
		AWSUtil:       awsUtil,
		DateToAgeUtil: dateToAgeUtil,
		AuthClient:    *authClient,
		KafkaProducer: kafkaProducer,
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

		postIdString := fmt.Sprint(split.PostId)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return nil, err
		}
		(*postData)[i].MediaUrl = *postMedias
		LikeCommentCount, err := r.PostRepo.GetPostLikeAndCommentsCount(&postIdString)
		if err != nil {
			return nil, err
		}

		(*postData)[i].LikesCount = LikeCommentCount.LikesCount
		(*postData)[i].CommentsCount = LikeCommentCount.CommentsCount

		(*postData)[i].PostAge = *r.DateToAgeUtil.DateTOAge(&(*postData)[i].CreatedAt)
	}

	return postData, nil
}

func (r *PostUseCase) DeletePost(postId, userId *string) error {

	err := r.PostRepo.DeletePostMedias(postId)
	if err != nil {
		return err
	}
	err = r.PostRepo.DeletePostById(postId, userId)
	if err != nil {
		return err
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

func (r *PostUseCase) LikePost(postId, userId *string) *error {

	PostCreatorId, err := r.PostRepo.GetPostCreatorId(postId)
	if err != nil {
		return &err
	}
	inserted, err := r.PostRepo.LikePost(postId, userId)
	if err != nil {
		fmt.Println(err)
		return &err
	}

	if *PostCreatorId != *userId { //avoiding the case where user likes his own post
		var message requestmodels_posnrel.KafkaNotificationTopicModel
		if inserted {
			message.UserID = *PostCreatorId
			message.ActorID = *userId
			message.ActionType = "like"
			message.TargetID = *postId
			message.TargetType = "post"
			message.CreatedAt = time.Now()

			err = r.KafkaProducer.KafkaNotificationProducer(&message)
			if err != nil {
				return &err
			}
		}
	}
	return nil
}

func (r *PostUseCase) UnLikePost(postId, userId *string) *error {

	err := r.PostRepo.UnLikePost(postId, userId)
	if err != nil {
		fmt.Println(err)
		return &err
	}
	return nil
}

func (r *PostUseCase) GetMostLovedPostsFromGlobalUser(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	postData, err := r.PostRepo.GetMostLovedPostsFromGlobalUser(userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for i, split := range *postData {
		context, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: fmt.Sprint((*postData)[i].UserId)})
		if err != nil || userData.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprint(err) + userData.ErrorMessage)
		}

		(*postData)[i].UserName = userData.UserName
		(*postData)[i].UserProfileImgURL = userData.UserProfileImgURL

		postIdString := fmt.Sprint(split.PostId)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return nil, err
		}
		(*postData)[i].MediaUrl = *postMedias
		LikeCommentCount, err := r.PostRepo.GetPostLikeAndCommentsCount(&postIdString)
		if err != nil {
			return nil, err
		}

		(*postData)[i].LikesCount = LikeCommentCount.LikesCount
		(*postData)[i].CommentsCount = LikeCommentCount.CommentsCount

		(*postData)[i].PostAge = *r.DateToAgeUtil.DateTOAge(&(*postData)[i].CreatedAt)
	}

	return postData, nil
}

func (r *PostUseCase) GetAllRelatedPostsForHomeScreen(userId, limit, offset *string) (*[]responsemodels_postnrel.PostData, error) {
	postData, err := r.PostRepo.GetAllActiveRelatedPostsForHomeScreen(userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for i, split := range *postData {
		context, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: fmt.Sprint((*postData)[i].UserId)})
		if err != nil || userData.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprint(err) + userData.ErrorMessage)
		}

		(*postData)[i].UserName = userData.UserName
		(*postData)[i].UserProfileImgURL = userData.UserProfileImgURL

		postIdString := fmt.Sprint(split.PostId)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return nil, err
		}
		(*postData)[i].MediaUrl = *postMedias
		LikeCommentCount, err := r.PostRepo.GetPostLikeAndCommentsCount(&postIdString)
		if err != nil {
			return nil, err
		}

		(*postData)[i].LikesCount = LikeCommentCount.LikesCount
		(*postData)[i].CommentsCount = LikeCommentCount.CommentsCount

		(*postData)[i].PostAge = *r.DateToAgeUtil.DateTOAge(&(*postData)[i].CreatedAt)
	}

	return postData, nil
}
