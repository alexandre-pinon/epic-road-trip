package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type authControllerSuite struct {
	suite.Suite
	cfg        config.Config
	svc        *mocks.AuthService
	ctrl       AuthController
	testServer *httptest.Server
}

func (suite *authControllerSuite) SetupTest() {
	svc := new(mocks.AuthService)
	ctrl := NewAuthController(&suite.cfg, svc)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	apiRoutes := router.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.GET("/")
		}
	}
	testServer := httptest.NewServer(router)

	suite.testServer = testServer
	suite.svc = svc
	suite.ctrl = ctrl
}

func (suite *authControllerSuite) TearDownTest() {
	defer suite.testServer.Close()
}

func TestAuthController(t *testing.T) {
	cfg := config.GetConfig()
	cfg.Env = config.Test
	suite.Run(t, &authControllerSuite{cfg: *cfg})
}
