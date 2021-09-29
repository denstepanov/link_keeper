package handlers

import "github.com/gin-gonic/gin"

type HttpHandler struct {
}

func (h *HttpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// sign in
			auth.POST("", h.signIn)
			// sign out
			auth.GET("", h.signOut)
		}
		// users := api.Group("/users")
		// {
		// 	users.GET("/:user_id", h.get)
		//	users.GET("", h.getList)
		//	users.POST("", h.create)
		//	users.PUT("/:user_id", h.update)
		//	users.DELETE("/:user_id", h.delete)

		//	links := users.Group("/:user_id/links")
		//	{
		//		links.GET("/:link_id")
		//		links.GET("")
		//		links.POST("")
		//		links.PUT("/:link_id")
		//		links.DELETE("/:link_id")
		//	}
	}

	return router
}
