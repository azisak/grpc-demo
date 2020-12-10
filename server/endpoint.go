package server

import (
	"context"

	"example.com/demo-grpc/user_service"
	"github.com/prometheus/common/log"
)

type Controller struct {
	user_service.UnimplementedUserServiceServer
	svc Service
}

func NewController() *Controller {
	return &Controller{
		svc: NewService(),
	}
}

func (c *Controller) GetUser(ctx context.Context, r *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	log.Info("GetUser request: ", r)
	result, err := c.svc.GetUser(r.GetId())
	return &user_service.GetUserResponse{
		User: &user_service.User{Id: result.ID, Name: result.Name},
	}, err
}
