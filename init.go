package swagger_doc

import (
	"github.com/jimu-server/swagger-doc/docs"
	"github.com/jimu-server/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.Title = "Jimu 服务"
	docs.SwaggerInfo.Description = "api 文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	web.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
