package srv

import (
	"net/http"

	"practise/test_task_post/config"
	"practise/test_task_post/pkg/logger"
	"practise/test_task_post/services"

	"time"
)

// const endpoint = "/generate-pdf/use-doc"

type Service struct {
	// url    string // http://172.25.102.137:7302/generate-pdf/use-doc
	cfg    config.Config
	log    logger.Logger
	api    *openApi
	client *http.Client
}

func New(cfg config.Config, log logger.Logger) *Service {
	return &Service{
		cfg:    cfg,
		log:    log,
		client: &http.Client{Timeout: 10 * time.Minute},
		api:    nil,
	}
}

func (s *Service) OpenApiSrv() services.OpenApiService {
	if s.api == nil {
		s.api = &openApi{s.cfg, s.log, s.client}
	}
	return s.api

}
