package rest

import (
	"net/http"
	"practise/test_task_grud/config"
	"practise/test_task_grud/pkg/logger"
	"practise/test_task_grud/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  storage.Storage
	log    logger.Logger
	cfg    config.Config
	router *gin.Engine
}

func New(st storage.Storage, log logger.Logger, cfg config.Config, router *gin.Engine) *Server {
	srv := &Server{
		log: log,
		cfg: cfg,

		router: router,
		store:  st,
	}

	srv.endpoints()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
