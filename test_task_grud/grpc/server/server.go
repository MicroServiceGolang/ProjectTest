package server

import (
	"context"
	"practise/test_task_grud/config"
	"practise/test_task_grud/grpc/genprotos/grud_protos"
	"practise/test_task_grud/models"
	"practise/test_task_grud/pkg/logger"
	"practise/test_task_grud/storage"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrudServer struct {
	cfg   config.Config
	log   logger.Logger
	store storage.Storage
}

func (p *GrudServer) GetOneGrpc(ctx context.Context, req *grud_protos.IDParam) (*grud_protos.GrudParams, error) {
	id := req.ID
	idstr := strconv.FormatInt(int64(id), 10)

	resp, err := p.store.PostStore().Read(idstr)
	if err != nil {
		p.log.Error(err.Error())
		return nil, err
	}

	return Post(*resp), nil
}

func (p *GrudServer) GetAllGrpc(ctx context.Context, empty *emptypb.Empty) (*grud_protos.GrudResponse, error) {
	resp, err := p.store.PostStore().ReadAll()
	if err != nil {
		p.log.Error(err.Error())
		return nil, err
	}
	result := grud_protos.GrudResponse{
		ArrMessage: Posts(resp),
	}

	return &result, nil
}

func (p *GrudServer) UpdateGrpc(ctx context.Context, val *grud_protos.GrudParams) (*grud_protos.BooleanResponse, error) {
	result := grud_protos.BooleanResponse{Result: true}

	req := models.OpenApi{
		ID:     int(val.ID),
		UserID: int(val.UserID),
		Title:  val.Title,
		Body:   val.Body,
	}
	err := p.store.PostStore().Update(req)
	if err != nil {
		result := grud_protos.BooleanResponse{Result: false}
		return &result, err
	}

	return &result, nil
}

func (p *GrudServer) DeleteGrpc(ctx context.Context, req *grud_protos.IDParam) (*grud_protos.BooleanResponse, error) {
	result := grud_protos.BooleanResponse{Result: true}
	id := req.ID
	idstr := strconv.FormatInt(int64(id), 10)
	err := p.store.PostStore().Delete(idstr)
	if err != nil {
		result := grud_protos.BooleanResponse{Result: false}
		return &result, err
	}
	return &result, nil
}

/** Start Server **/
func StartServer(cfg config.Config, log logger.Logger, storage storage.Storage) *grpc.Server {
	log.Info("Starting gRPC Server")

	server := grpc.NewServer()

	grud_protos.RegisterGrudServiceServer(server, &GrudServer{
		cfg:   cfg,
		log:   log,
		store: storage,
	})

	reflection.Register(server)

	return server

}

func Posts(posts []models.OpenApi) []*grud_protos.GrudParams {
	views := make([]*grud_protos.GrudParams, len(posts))

	for i, post := range posts {
		views[i] = Post(post)
	}

	return views
}
func Post(o models.OpenApi) *grud_protos.GrudParams {
	res := &grud_protos.GrudParams{
		ID:     uint64(o.ID),
		UserID: uint64(o.UserID),
		Title:  o.Title,
		Body:   o.Body,
	}
	return res
}
