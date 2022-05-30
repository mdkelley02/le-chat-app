package user

import (
	"context"
	"fmt"

	"github.com/mdkelley02/le-chat-app/user-service/internal/storage"
	protos "github.com/mdkelley02/le-chat-app/user-service/protos/userservice"
	"github.com/sirupsen/logrus"
)

type HandlerIF interface {
	CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.CreateUserResponse, error)
	GetUser(ctx context.Context, req *protos.GetUserRequest) (*protos.GetUserResponse, error)
	UpdateUser(ctx context.Context, req *protos.UpdateUserRequest) (*protos.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *protos.DeleteUserRequest) (*protos.DeleteUserResponse, error)
}

type Handler struct {
	storeDb storage.StorageIF
	log     *logrus.Logger
}

func NewUserHandler(storeDb storage.StorageIF, log *logrus.Logger) *Handler {
	h := Handler{
		storeDb: storeDb,
		log:     log,
	}
	return &h
}

func (h *Handler) CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.CreateUserResponse, error) {
	fmt.Println("CreateUser")
	return nil, nil
}

func (h *Handler) GetUser(ctx context.Context, req *protos.GetUserRequest) (*protos.GetUserResponse, error) {
	fmt.Println("GetUser")
	return nil, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *protos.UpdateUserRequest) (*protos.UpdateUserResponse, error) {
	fmt.Println("UpdateUser")
	return nil, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *protos.DeleteUserRequest) (*protos.DeleteUserResponse, error) {
	fmt.Println("DeleteUser")
	return nil, nil
}
