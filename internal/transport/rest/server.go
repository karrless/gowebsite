package rest

import (
	"context"
	"gowebsite/internal/transport/rest/routes"
	"gowebsite/pkg/db/postgres"

	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	r    *gin.Engine
	port string
}

func NewRESTServer(ctx context.Context, db *postgres.DB, port string) *RESTServer {
	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.31.157"})

	api := r.Group("/api")

	v1 := api.Group("/v1")

	routes.PortfolioRoutes(ctx, v1, db)

	return &RESTServer{r: r, port: port}
}

func (s *RESTServer) Run(ctx context.Context) error {
	return s.r.Run(":" + s.port)
}
