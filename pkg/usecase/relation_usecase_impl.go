package usecase_postnrel

import (
	"context"
	"errors"
	"log"
	"time"

	requestmodels_posnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/models/requestmodels"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	interface_repo_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository/interface"
	interface_usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase/interface"
	interface_kafkaproducer "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/kafka_producer/interface"
)

type RelationUseCase struct {
	RelationRepo  interface_repo_postnrel.IRelationRepo
	PostRepo      interface_repo_postnrel.IPostRepo
	AuthClient    pb.AuthServiceClient
	kafkaProducer interface_kafkaproducer.IKafkaProducer
}

func NewRelationUseCase(relationRepo interface_repo_postnrel.IRelationRepo,
	postRepo interface_repo_postnrel.IPostRepo,
	authClient *pb.AuthServiceClient,
	kafkaProducer interface_kafkaproducer.IKafkaProducer) interface_usecase_postnrel.IRelationUseCase {
	return &RelationUseCase{
		RelationRepo:  relationRepo,
		PostRepo:      postRepo,
		AuthClient:    *authClient,
		kafkaProducer: kafkaProducer,
	}
}

func (r *RelationUseCase) GetCountsForUserProfile(userId *string) (*uint, *uint, *uint, *error) {

	a, b, err := r.RelationRepo.GetFollowerAndFollowingCountofUser(userId)
	if err != nil {
		return nil, nil, nil, err
	}
	c, err := r.PostRepo.GetPostCountOfUser(userId)
	if err != nil {
		return nil, nil, nil, err
	}
	return a, b, c, nil

}

func (r *RelationUseCase) Follow(userId, userBId *string) *error {
	var message requestmodels_posnrel.KafkaNotificationTopicModel

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := r.AuthClient.CheckUserExist(context, &pb.RequestUserId{
		UserId: *userBId,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.ErrorMessage != "" {
		err = errors.New(resp.ErrorMessage)
		return &err
	}
	if !resp.ExistStatus {
		err = errors.New("no user exist with this id,enter a valid userid")
		return &err
	}
	inserted, err := r.RelationRepo.InitiateFollowRelationship(userId, userBId)
	if err != nil {
		return &err
	}

	if inserted {
		message.UserID = *userBId
		message.ActorID = *userId
		message.ActionType = "follow"
		message.TargetID = "0"
		message.CreatedAt = time.Now()

		err = r.kafkaProducer.KafkaNotificationProducer(&message)
		if err != nil {
			return &err
		}
	}
	return nil
}

func (r *RelationUseCase) UnFollow(userId, userBId *string) *error {

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := r.AuthClient.CheckUserExist(context, &pb.RequestUserId{
		UserId: *userBId,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.ErrorMessage != "" {
		err = errors.New(resp.ErrorMessage)
		return &err
	}
	if !resp.ExistStatus {
		err = errors.New("no user exist with this id,enter a valid userid")
		return &err
	}
	return r.RelationRepo.InitiateUnFollowRelationship(userId, userBId)
}

func (r *RelationUseCase) GetFollowersIds(userId *string) (*[]uint64, error) {

	userIdSlice, err := r.RelationRepo.GetFollowersIdsOfUser(userId)
	if err != nil {
		return nil, err
	}
	return userIdSlice, nil
}
func (r *RelationUseCase) GetFollowingsIds(userId *string) (*[]uint64, error) {

	userIdSlice, err := r.RelationRepo.GetFollowingsIdsOfUser(userId)
	if err != nil {
		return nil, err
	}
	return userIdSlice, nil
}

func (r *RelationUseCase) UserAFollowingUserBorNot(userId, userBId *string) (bool, error) {

	resp, err := r.RelationRepo.UserAFollowingUserBorNot(userId, userBId)
	if err != nil {
		return false, err
	}
	return resp, err
}
