package services

import "practise/test_task_post/models"

// const endpoint = "/generate-pdf/use-doc"

type Service interface {
	OpenApiSrv() OpenApiService
}

type OpenApiService interface {
	Get() (*models.Data, error)
}
