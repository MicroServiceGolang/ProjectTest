package rest

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) endpoints() {
	s.router.GET("/getone", s.GetOne)
	s.router.GET("/getall", s.GetAll)
	s.router.PUT("/update/{id}", s.Update)
	s.router.DELETE("/delete/{id}", s.Delete)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
