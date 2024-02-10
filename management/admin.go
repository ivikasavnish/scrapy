package management

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

type AdminInterface interface {
	AdminRoutePath() string
}

func AdminRoutes(engine *gin.Engine) *gin.RouterGroup {
	// Implement AdminInterface method
	adminroutes := engine.Group("/admin")
	adminroutes.Use(AuthRequired())
	log.Println("Admin routes", adminroutes, len(AdminModels))
	{
		for _, model := range AdminModels {
			// Implement AdminInterface method
			log.Println(reflect.ValueOf(model), reflect.TypeOf(model))
			adminroutes.GET("/"+model.AdminRoutePath(), func(ctx *gin.Context) {
				ctx.JSON(200, model)
			})
			adminroutes.POST("/"+model.AdminRoutePath(), func(ctx *gin.Context) {
				ctx.JSON(200, model)
			})

			adminroutes.PUT("/"+model.AdminRoutePath(), func(ctx *gin.Context) {
				ctx.JSON(200, model)
			})
			adminroutes.DELETE("/"+model.AdminRoutePath(), func(ctx *gin.Context) {
				ctx.JSON(200, model)
			})
		}
	}
	return adminroutes
}
