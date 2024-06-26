package main

import (
	"fmt"
	"log"
	"net"

	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	di_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/di"
	"github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/pb"
	"google.golang.org/grpc"
)

func main() {

	config, err := config_postNrelSvc.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := di_postNrelSvc.InitializePostNRelationServer(config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.PortMngr.RunnerPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	fmt.Println("PostNrel Service started on:", config.PortMngr.RunnerPort)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	pb.RegisterPostNrelServiceServer(grpcServer, server)

	// Log every connection attempt to the server
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Println("Error accepting connection:", err)
				continue
			}
			log.Println("New connection from:", conn.RemoteAddr())
		}
	}()

	// Serve the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start PostNrelations_service server:%v", err)

	}

}
