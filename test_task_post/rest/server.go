package rest

import (
	"net/http"
	"practise/test_task_post/config"
	"practise/test_task_post/pkg/logger"
	"practise/test_task_post/services"
	"practise/test_task_post/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store   storage.Storage
	service services.Service
	log     logger.Logger
	cfg     config.Config
	router  *gin.Engine
}

func New(st storage.Storage, s services.Service, log logger.Logger, cfg config.Config, router *gin.Engine) *Server {
	srv := &Server{
		log: log,
		cfg: cfg,

		router:  router,
		service: s,
		store:   st,
	}

	srv.endpoints()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
