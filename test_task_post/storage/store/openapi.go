package store

import (
	"errors"
	"practise/test_task_post/models"
)

var (
	ErrorDuplicate = "this message is already written"
	ErrorNotFound  = "no message is found"
)

type openApiRepo struct {
	store *Store
}

func (o *openApiRepo) Create(m models.Data) error {

	for _, data := range m.Data {

		query := `INSERT INTO testposts(id,user_id,title,body) VALUES($1,$2,$3,$4)`

		_, err := o.store.db.Exec(query,
			data.ID,
			data.UserID,
			data.Title,
			data.Body,
		)
		if err != nil {
			return errors.New(ErrorDuplicate)
		}
	}
	return nil
}
