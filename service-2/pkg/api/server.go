package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api/routes"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/config"
	"google.golang.org/grpc"
)

const (
	basePath = "/api"
)

type Server struct {
	restServer *gin.Engine
	restPort   string

	grpcServer *grpc.Server
	ls         net.Listener
}

func NewServer(cfg config.Config, serviceHandler interfaces.ServiceHandler) (*Server, error) {

	// register rest server
	engine := gin.New()
	engine.Use(gin.Logger())
	// register all REST api routes
	routes.RegisterRoutes(engine.Group(basePath), serviceHandler)

	return &Server{
		restServer: engine,
		restPort:   cfg.Service2RestPort,
	}, nil
}

func (s *Server) Start() error {

	return s.restServer.Run(":" + s.restPort)

}
