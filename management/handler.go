package management

import "github.com/gin-gonic/gin"

func GenCRUDResource(e *gin.Engine, c CRUD, path string, prefix string) {
	e.POST(prefix+path, func(ctx *gin.Context) {
		c.Create(ctx)
	})
	e.GET(prefix+path, func(ctx *gin.Context) {
		c.Read(ctx)
	})
	e.PUT(prefix+path, func(ctx *gin.Context) {
		c.Update(ctx)
	})
	e.DELETE(prefix+path, func(ctx *gin.Context) {
		c.Delete(ctx)
	})
	e.GET(prefix+path+"s", func(ctx *gin.Context) {
		c.List(ctx)
	})

}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement authentication
		c.Next()
	}
}

func Routes() {
	r := gin.Default()
	d := Domain{
		Name: "domain",
	}
	GenCRUDResource(r, &d, "/domain", "/api/v1")

	AddtoAdminModels(&AdminModels, &d)
	AdminRoutes(r)
	r.Run()
}
