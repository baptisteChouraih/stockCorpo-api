package server

import (
	"stockCorpo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, suggestionHandler *handlers.SuggestionHandler) *gin.Engine {

	router := gin.Default()

	api := router.Group("/api")
	{
		users := api.Group("/user")
		{
			users.POST("", userHandler.CreateUser)
			//CRUD :
			//users.PUT (upd)
			//users.GET (get)
			//users.DELETE (del)
		}

		product := api.Group("/product")
		{
			product.POST("", productHandler.CreateProduct)
			product.PUT("", productHandler.EditProduct)
		}
		suggestion := api.Group("/suggestion")
		{
			suggestion.POST("", suggestionHandler.CreateSuggestion)
			suggestion.DELETE("", suggestionHandler.DeleteSuggestion)
			suggestion.PUT("", suggestionHandler.UpdateSuggestion)
			suggestion.GET("", suggestionHandler.GetAllSuggestions)
		}
	}
	return router
}
