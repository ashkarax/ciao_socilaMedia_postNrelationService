package di_postNrelSvc

import (
	"fmt"

	client_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/client"
	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	db_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/db"
	server_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/server"
	repository_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository"
	usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase"
	datetoage "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/DateToAge"
	aws_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/aws_s3"
	hashpassword_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/hash_password"
	kafkaproducer "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/kafka_producer"
)

func InitializePostNRelationServer(config *config_postNrelSvc.Config) (*server_postNrelSvc.PostNrelService, error) {

	hashUtil := hashpassword_postNrelSvc.NewHashUtil()

	DB, err := db_postNrelSvc.ConnectDatabase(&config.DB, hashUtil)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}

	awsUtil := aws_postnrel.AWSS3FileUploaderSetup(config.AwsS3)
	dateToAgeUtil := datetoage.NewDateToAgeUtil()
	kafkaProducer := kafkaproducer.NewKafkaProducer(config.KafkaConfig)

	authClient, err := client_postnrel.InitAuthServiceClient(config)
	if err != nil {
		fmt.Println("--------err--------", err)
		return nil, err
	}

	postRepo := repository_postnrel.NewPostRepo(DB)
	postUseCase := usecase_postnrel.NewPostUseCase(postRepo, awsUtil, dateToAgeUtil, authClient, kafkaProducer)

	relationRepo := repository_postnrel.NewRelationRepo(DB)
	relationUseCase := usecase_postnrel.NewRelationUseCase(relationRepo, postRepo, authClient, kafkaProducer)

	commentRepo := repository_postnrel.NewCommentRepo(DB)
	commentUseCase := usecase_postnrel.NewCommentUseCase(commentRepo, dateToAgeUtil, authClient, kafkaProducer,postRepo)

	return server_postNrelSvc.NewPostNrelServiceServer(postUseCase, relationUseCase, commentUseCase), nil

}
