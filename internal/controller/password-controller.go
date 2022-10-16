package controller

import (
	"context"
	"errors"

	"github.com/password/logger"
	api "github.com/password/password"
)

func (c *Controller) CheckAndHash(ctx context.Context, r *api.Request) (*api.Hash, error) {
	log := logger.GetLogger(ctx)

	res, err := c.service.CheckAndHash(ctx, r.GetRequest())
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to check and hash password")
	}

	return &api.Hash{Result: string(res)}, nil
}

func (c *Controller) Compare(ctx context.Context, r *api.CompareRequest) (*api.Ok, error) {
	log := logger.GetLogger(ctx)

	res, err := c.service.Compare(ctx, r.GetPassword(), r.GetHash())
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to compare hash and password")
	}

	return &api.Ok{Ok: res}, nil
}
