package handler

import (
	"chat/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/log-in", h.logIn)
	}

	api := router.Group("/api", h.identifyUser)
	{
		dialogs := api.Group("/dialogs")
		{
			dialogs.GET("/", h.getAllDialogs)
			dialogs.GET("/:id", h.getDialog)
		}

		api.GET("/user-status/:id", h.getUserStatus)
		api.GET("/message", h.getMessage)
		api.POST("/message", h.sendMessage)
	}
	
	return router
}
