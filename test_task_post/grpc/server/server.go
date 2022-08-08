package server

import (
	"context"
	"practise/test_task_post/config"
	"practise/test_task_post/grpc/genprotos/post_protos"
	"practise/test_task_post/models"
	"practise/test_task_post/pkg/logger"
	"practise/test_task_post/services"
	"practise/test_task_post/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostServer struct {
	cfg   config.Config
	log   logger.Logger
	srv   services.Service
	store storage.Storage
}

func (p *PostServer) GetOpenApiGRPC(ctx context.Context, empty *emptypb.Empty) (*post_protos.PostResponse, error) {
	p.log.Info("Reading all posts")

	resp, err := p.srv.OpenApiSrv().Get()
	if err != nil {
		p.log.Error(err.Error())
		return nil, err
	}

	err = p.store.OpenApiStore().Create(*resp)
	if err != nil {
		p.log.Error(err.Error())
		return nil, err
	}

	result := post_protos.PostResponse{
		ArrMessage: Posts(resp.Data),
	}

	return &result, nil
}

/** Start Server **/
func StartServer(cfg config.Config, log logger.Logger, service services.Service, storage storage.Storage) *grpc.Server {
	log.Info("Starting gRPC Server")

	server := grpc.NewServer()

	post_protos.RegisterPostServiceServer(server, &PostServer{
		cfg:   cfg,
		log:   log,
		srv:   service,
		store: storage,
	})

	reflection.Register(server)

	return server

}

func Posts(posts []models.OpenApi) []*post_protos.PostParams {
	views := make([]*post_protos.PostParams, len(posts))

	for i, post := range posts {
		views[i] = Post(post)
	}

	return views
}
func Post(o models.OpenApi) *post_protos.PostParams {
	res := &post_protos.PostParams{
		ID:     uint64(o.ID),
		UserID: uint64(o.UserID),
		Title:  o.Title,
		Body:   o.Body,
	}
	return res
}
