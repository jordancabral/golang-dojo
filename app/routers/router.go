package routers

import (
	"github.com/gin-gonic/gin"
	routers "github.com/jordancabral/golang-dojo/app/routers/api"
)

func InitRouter() *gin.Engine {

	// Init Gin
	r := gin.Default()

	// mock config CRUD
	// example curl: curl -XPOST http://localhost:8080/mock --header 'Content-Type: application/json' --data-raw '{"path":"test","Method":"GET", "Response": "app/mocks/test.json", "Code": 200, "ProxyMode": false, "ProxyUrl": "http://www.google.com"}'
	r.POST("/configs", routers.CreateConfig)
	r.GET("/configs", routers.GetAllConfigs)
	r.GET("/configs/:id", routers.GetConfig)
	r.DELETE("/configs/:id", routers.RemoveConfig)
	r.PUT("/configs/:id", routers.UpdateConfig)

	// mock endpoints
	r.Any("/mock/:path", routers.MockRoute)

	return r

}
