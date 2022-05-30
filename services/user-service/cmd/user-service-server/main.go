package main

import (
	"context"
	"fmt"
	"net"

	userservice "github.com/mdkelley02/le-chat-app/user-service/cmd/user-service-server/services"
	configs "github.com/mdkelley02/le-chat-app/user-service/configs"
	store "github.com/mdkelley02/le-chat-app/user-service/internal/storage"
	user "github.com/mdkelley02/le-chat-app/user-service/internal/user"
	protos "github.com/mdkelley02/le-chat-app/user-service/protos/userservice"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.New()
	ctx := context.Background()
	var cfgs configs.Settings
	err := envconfig.Process(ctx, &cfgs)
	if err != nil {
		logger.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%s", cfgs.Host, cfgs.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err)
	}

	server := grpc.NewServer()
	db := store.NewDB(cfgs.PostgresHost, cfgs.PostgresPort, cfgs.PostgresUser, cfgs.PostgresPassword, cfgs.PostgresDB)
	store := store.NewStore(db, logger)
	userHandler := user.NewUserHandler(store, logger)
	userService := userservice.NewUserService(userHandler, logger)
	protos.RegisterUserServiceServer(server, userService)
	serverErrors := make(chan error, 1)
	logger.Info("serving gRPC on ", addr)

	go func() {
		serverErrors <- server.Serve(lis)
	}()

	select {
	case err := <-serverErrors:
		logger.Fatal(err)
	case <-ctx.Done():
		server.GracefulStop()
	}

}
