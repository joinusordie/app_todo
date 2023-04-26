package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joinusordie/app_todo/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/joinusordie/app_todo/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(h.corsSetting())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		signUp := auth.Group("/sign-up")
		{
			signUp.POST("/", h.signUp)
			signUp.GET("", h.checkUsername)
		}

		signIn := auth.Group("/sign-in")
		{
			signIn.POST("/", h.signIn)
		}
	}

	api := router.Group("/api", h.userIdentity)
	{
		account := api.Group("/account")
		{
			account.DELETE("/", h.deleteUser)
			account.GET("/", h.getUser)
		}

		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
