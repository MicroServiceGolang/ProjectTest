package srv

import (
	"encoding/json"
	"io"
	"net/http"
	"practise/test_task_post/config"
	"practise/test_task_post/models"
	"practise/test_task_post/pkg/logger"
)

type openApi struct {
	cfg    config.Config
	log    logger.Logger
	client *http.Client
}

func (o *openApi) Get() (*models.Data, error) {
	var resp *models.Data
	resHttp, err := o.client.Get(o.cfg.OpenApiURL)
	if err != nil {
		return nil, err
	}
	defer resHttp.Body.Close()
	resByte, err := io.ReadAll(resHttp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resByte, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
