package storage

import "practise/test_task_grud/models"

type Storage interface {
	PostStore() PostRepo
}

type PostRepo interface {
	Create(models.Data) error
	Read(string) (*models.OpenApi, error)
	ReadAll() ([]models.OpenApi, error)
	Update(models.OpenApi) error
	Delete(string) error
}
