package store

import (
	"context"
	"errors"
	"practise/test_task_grud/models"

	"strconv"
)

type postRepo struct {
	store *Store
}

func (o *postRepo) Create(m models.Data) error {

	for _, data := range m.Data {

		query := `INSERT INTO testposts(id,user_id,title,body) VALUES($1,$2,$3,$4)`

		_, err := o.store.db.Exec(query,
			data.ID,
			data.UserID,
			data.Title,
			data.Body,
		)
		if err != nil {
			o.store.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func (o postRepo) Read(ID string) (*models.OpenApi, error) {

	var (
		openApi models.OpenApi
	)

	intID, err := strconv.Atoi(ID)
	if err != nil {
		o.store.log.Error(err.Error())
	}

	query := `SELECT 
	id,
	user_id,
	title,
	body
	FROM testposts WHERE id = $1`

	err = o.store.db.QueryRow(query, intID).Scan(
		&openApi.ID,
		&openApi.UserID,
		&openApi.Title,
		&openApi.Body,
	)
	if err != nil {
		o.store.log.Error(err.Error())
		return nil, errors.New("post not found")
	}
	return &openApi, nil
}

func (o postRepo) ReadAll() ([]models.OpenApi, error) {
	var result []models.OpenApi
	query := `SELECT 
	id,
	user_id,
	title,
	body
	FROM testposts ORDER BY id ASC`

	rows, err := o.store.db.QueryContext(context.Background(), query)
	if err != nil {
		o.store.log.Error("[SQL Storage] QueryContext: " + err.Error())
		return nil, err
	}
	for rows.Next() {
		var (
			resp models.OpenApi
		)

		err := rows.Scan(
			&resp.ID,
			&resp.UserID,
			&resp.Title,
			&resp.Body,
		)
		if err != nil {
			o.store.log.Error("[SQL Storage] Scan: " + err.Error())
			return nil, err
		}
		result = append(result, resp)
	}

	return result, nil
}

func (o postRepo) Update(val models.OpenApi) error {

	query := `
	UPDATE 
	TESTPOSTS 
	SET 
		USER_ID=$2, 
		TITLE = $3, 
		BODY=$4
	WHERE
		ID = $1
	`

	_, err := o.store.db.Exec(query, val.ID, val.UserID, val.Title, val.Body)
	if err != nil {
		o.store.log.Error("[SQL Storage] Scan: " + err.Error())
		return err
	}

	return nil
}
func (o postRepo) Delete(ID string) error {

	intID, err := strconv.Atoi(ID)
	if err != nil {
		o.store.log.Error(err.Error())
	}

	query := `Delete from testposts where id=$1`

	o.store.db.Exec(query, intID)

	return nil
}
