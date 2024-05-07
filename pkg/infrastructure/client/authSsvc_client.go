package client_postnrel

import (
	"fmt"

	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthServiceClient(config *config_postNrelSvc.Config) (*pb.AuthServiceClient, error) {
	cc, err := grpc.Dial(config.PortMngr.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewAuthServiceClient(cc)

	return &Client, nil
}
