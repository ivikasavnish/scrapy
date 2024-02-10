package management

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Domain struct {
	Name    string `json:"name" binding:"required" form:"name" query:"name"`
	URL     string `json:"url" binding:"required" form:"url" query:"url"`
	Enabled bool   `json:"enabled" binding:"required" form:"enabled" query:"enabled"`
}

type CRUD interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

func (d *Domain) Create(c *gin.Context) {

}

func (d *Domain) Read(c *gin.Context) {
	// Implement Read method
}

func (d *Domain) Update(c *gin.Context) {
	// Implement Update method
}

func (d *Domain) Delete(c *gin.Context) {
	// Implement Delete method
}

func (d *Domain) List(c *gin.Context) {
	// Implement List method
}

func (d *Domain) AdminRoutePath() string {
	return "/domain"
}

func AddtoAdminModels(adminmodelcontainer *[]AdminInterface, i ...AdminInterface) {
	*adminmodelcontainer = append(*adminmodelcontainer, i...)
	log.Println(len(i), "models added to admin interface", *adminmodelcontainer)
}
