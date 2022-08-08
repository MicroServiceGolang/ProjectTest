package storage

import "practise/test_task_post/models"

type Storage interface {
	OpenApiStore() OpenApiRepo
}

type OpenApiRepo interface {
	Create(models.Data) error
}
