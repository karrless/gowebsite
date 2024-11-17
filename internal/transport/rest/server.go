package rest

import (
	"context"
	"gowebsite/docs"
	"gowebsite/internal/transport/rest/routes"
	"gowebsite/pkg/db/postgres"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RESTServer struct {
	r    *gin.Engine
	port string
}

func NewRESTServer(ctx context.Context, db *postgres.DB, port string) *RESTServer {
	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.31.157"})
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.Title = "KarrlessGo API"
	docs.SwaggerInfo.Description = "API for Karrless.ru website"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	api := r.Group("/api")
	v1 := api.Group("/v1")

	routes.PortfolioRoutes(ctx, v1, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &RESTServer{r: r, port: port}
}

func (s *RESTServer) Run(ctx context.Context) error {
	return s.r.Run(":" + s.port)
}
