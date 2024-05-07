package di_postNrelSvc

import (
	"fmt"

	client_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/client"
	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	db_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/db"
	server_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/server"
	repository_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/repository"
	usecase_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/usecase"
	aws_postnrel "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/aws_s3"
	hashpassword_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/hash_password"
)

func InitializePostNRelationServer(config *config_postNrelSvc.Config) (*server_postNrelSvc.PostNrelService, error) {

	hashUtil := hashpassword_postNrelSvc.NewHashUtil()

	DB, err := db_postNrelSvc.ConnectDatabase(&config.DB, hashUtil)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}

	awsUtil := aws_postnrel.AWSS3FileUploaderSetup(config.AwsS3)

	authClient, err := client_postnrel.InitAuthServiceClient(config)
	if err != nil {
		fmt.Println("--------err--------", err)
		return nil, err
	}

	postRepo := repository_postnrel.NewPostRepo(DB)
	postUseCase := usecase_postnrel.NewPostUseCase(postRepo, awsUtil, authClient)

	relationRepo := repository_postnrel.NewRelationRepo(DB)
	relationUseCase := usecase_postnrel.NewRelationUseCase(relationRepo,postRepo,authClient)

	return server_postNrelSvc.NewPostNrelServiceServer(postUseCase, relationUseCase), nil

}
