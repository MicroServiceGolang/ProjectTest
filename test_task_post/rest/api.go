package rest

import (
	"net/http"
	"practise/test_task_post/pkg/status"

	"github.com/gin-gonic/gin"
)

type OpenApi struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Get godoc swagger
// @Router /get [GET]
// @Summary Get
// @Description API for Get
// @Tags TestTask
// @Accept json
// @Produce json
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 404 {object} views.R
func (s *Server) Get(c *gin.Context) {
	resp, err := s.service.OpenApiSrv().Get()
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeRemoteOtherEntityNotFound, err)
		return
	}

	err = s.store.OpenApiStore().Create(*resp)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	s.handleSuccessResponse(c, resp)
}
