package services

import (
	"context"

	"github.com/mdkelley02/le-chat-app/user-service/internal/user"
	protos "github.com/mdkelley02/le-chat-app/user-service/protos/userservice"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	handler user.HandlerIF
	log     *logrus.Logger
}

func NewUserService(handler user.HandlerIF, log *logrus.Logger) *UserService {
	us := UserService{
		handler: handler,
		log:     log,
	}
	return &us
}

func (us *UserService) CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.CreateUserResponse, error) {
	return us.handler.CreateUser(ctx, req)
}

func (us *UserService) GetUser(ctx context.Context, req *protos.GetUserRequest) (*protos.GetUserResponse, error) {
	return us.handler.GetUser(ctx, req)
}

func (us *UserService) UpdateUser(ctx context.Context, req *protos.UpdateUserRequest) (*protos.UpdateUserResponse, error) {
	return us.handler.UpdateUser(ctx, req)
}

func (us *UserService) DeleteUser(ctx context.Context, req *protos.DeleteUserRequest) (*protos.DeleteUserResponse, error) {
	return us.handler.DeleteUser(ctx, req)
}
