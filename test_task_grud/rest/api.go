package rest

import (
	"errors"
	"fmt"
	"net/http"
	"practise/test_task_grud/models"
	"practise/test_task_grud/pkg/status"

	"strconv"

	"github.com/gin-gonic/gin"
)

type OpenApi struct {
	// ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Get godoc swagger
// @Router /getone [GET]
// @Summary GetOne
// @Description API for Get
// @Tags TestTask
// @Accept json
// @Produce json
// @Param id query string true "id for finding post"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 404 {object} views.R
func (s *Server) GetOne(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		s.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return

	}
	result, err := s.store.PostStore().Read(ID)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	s.handleSuccessResponse(c, result)

}

// Getall godoc swagger
// @Router /getall [GET]
// @Summary GetAll
// @Description API for Get
// @Tags TestTask
// @Accept json
// @Produce json
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 404 {object} views.R
func (s *Server) GetAll(c *gin.Context) {
	result, err := s.store.PostStore().ReadAll()
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	s.handleSuccessResponse(c, result)
}

// Update godoc swagger
// @Router /update/{id} [PUT]
// @Summary Update
// @Description API for update
// @Tags TestTask
// @Accept json
// @Param update body OpenApi true "Updated Parameters"
// @Param id query string true "ID To Update"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 404 {object} views.R
func (s *Server) Update(c *gin.Context) {

	ID := c.Query("id")
	if ID == "" {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, fmt.Errorf("id is empty"))
		return
	}
	intID, err := strconv.Atoi(ID)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, fmt.Errorf("error parsing string to int"))
		return
	}
	var req OpenApi
	if err := c.ShouldBindJSON(&req); err != nil {
		s.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}

	model := models.OpenApi{
		ID:     intID,
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
	}

	err = s.store.PostStore().Update(model)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	s.handleSuccessResponse(c, model)
}

// Delete godoc swagger
// @Router /delete/{id} [DELETE]
// @Summary Delete
// @Description API for delete
// @Tags TestTask
// @Accept json
// @Param id query string true "ID To Delete"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 404 {object} views.R
func (s *Server) Delete(c *gin.Context) {

	ID := c.Query("id")
	if ID == "" {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, fmt.Errorf("id is empty"))
		return
	}

	err := s.store.PostStore().Delete(ID)
	if err != nil {
		s.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	s.handleSuccessResponse(c, err)
}
