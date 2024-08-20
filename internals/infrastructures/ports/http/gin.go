package http

import (
	"github.com/Pr3c10us/boilerplate/internals/infrastructures/ports/http/authentication"
	"github.com/Pr3c10us/boilerplate/internals/services"
	"github.com/Pr3c10us/boilerplate/packages/configs"
	"github.com/Pr3c10us/boilerplate/packages/logger"
	"github.com/Pr3c10us/boilerplate/packages/middlewares"
	"github.com/Pr3c10us/boilerplate/packages/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type GinServer struct {
	Services    *services.Services
	Engine      *gin.Engine
	Logger      logger.Logger
	Environment *configs.EnvironmentVariables
}

func NewGinServer(services *services.Services, logger logger.Logger, environment *configs.EnvironmentVariables) *GinServer {
	ginServer := &GinServer{
		Services:    services,
		Engine:      gin.Default(),
		Logger:      logger,
		Environment: environment,
	}

	cookieStore := cookie.NewStore([]byte(environment.CookieSecret))
	gothic.Store = cookieStore

	// Middlewares
	ginServer.Engine.Use(middlewares.RequestLoggingMiddleware(logger))
	ginServer.Engine.Use(sessions.Sessions("auth", cookieStore))
	ginServer.Engine.Use(middlewares.ErrorHandlerMiddleware(logger))
	ginServer.Engine.NoRoute(middlewares.RouteNotFoundMiddleware())

	ginServer.Health()
	ginServer.Authentication()

	return ginServer
}
func (server *GinServer) Health() {
	server.Engine.GET("/health", func(c *gin.Context) {
		response.NewSuccessResponse("server up!!!", nil, nil).Send(c)
	})
}

func (server *GinServer) Authentication() {
	handler := authentication.NewAuthenticationHandler(server.Services.AuthenticationServices, server.Environment)
	oauthRoute := server.Engine.Group("/auth/:provider")
	{
		oauthRoute.GET("/", handler.InitiateAuth)
		oauthRoute.GET("/callback", handler.Callback)
	}
	//identityRoute := server.Engine.Group("/api/v1/identity")
	//{
	//	identityRoute.POST("/setup", middlewares.RiderAuthorizationMiddleware(server.Services.IdentityService, server.Environment), handler.IdentityVerification)
	//	identityRoute.POST("/device", middlewares.RiderAuthorizationMiddleware(server.Services.IdentityService, server.Environment), handler.AddRiderDevice)
	//}

}

func (server *GinServer) Run() {
	err := server.Engine.Run(server.Environment.Port)
	if err != nil {
		server.Logger.Log("panic", "failed to start server")
	}
}
