package rest

import (
	"net/http"
	"practise/test_task_post/pkg/status"
	"practise/test_task_post/views"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, views.R{
		Status:    status.Success,
		ErrorCode: 0,
		ErrorNote: "",
		Data:      data,
	})
}

func (s *Server) handleErrorResponse(c *gin.Context, httpCode, errCode int, err error) {
	c.JSON(httpCode, views.R{
		Status:    status.Failure,
		ErrorCode: errCode,
		ErrorNote: err.Error(),
		Data:      nil,
	})
}
