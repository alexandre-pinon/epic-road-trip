package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type amadeusServiceSuite struct {
	suite.Suite
	amadeusService AmadeusService
	cfg            config.Config
}

func (suite *amadeusServiceSuite) SetupTest() {
	amadeusService := NewAmadeusService(suite.cfg)
	suite.amadeusService = amadeusService
}

func (suite *amadeusServiceSuite) TestGetAccessToken_Positive() {
	accessTokenResponse := model.AccessTokenResponse{
		Type:            "amadeusOAuth2Token",
		Username:        "alexandre@yakow.io",
		ApplicationName: "epic-road-tirp",
		ClientID:        "amcU3qA6gA7D68dlgR85BC4KNIt1j4vG",
		TokenType:       "Bearer",
		AccessToken:     "kzijAxGphgIbhSm3WGIhmWvhep2G",
		ExpiresIn:       1799,
		State:           "approved",
		Scope:           "",
	}
	router := gin.New()
	router.POST("/v1/security/oauth2/token", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &accessTokenResponse)
	})
	server := httptest.NewServer(router)

	accessToken, err := suite.amadeusService.GetAccessToken(server.URL)
	suite.NoError(err, "no error if result")
	suite.Equal(accessTokenResponse.AccessToken, accessToken)
}

func (suite *amadeusServiceSuite) TestGetAccessToken_InvalidCredentials_Negative() {
	accessTokenResponse := model.AccessTokenError{
		Error:            "invalid_client",
		ErrorDescription: "Client credentials are invalid",
		Code:             38187,
		Title:            "Invalid parameters",
	}
	router := gin.New()
	router.POST("/v1/security/oauth2/token", func(ctx *gin.Context) {
		ctx.JSON(http.StatusUnauthorized, &accessTokenResponse)
	})
	server := httptest.NewServer(router)

	accessToken, err := suite.amadeusService.GetAccessToken(server.URL)
	suite.Error(err, "error when bad credentials")
	suite.Equal(http.StatusUnauthorized, err.(*model.AppError).StatusCode)
	suite.Equal("invalid client credentials", err.Error())
	suite.Empty(accessToken)
}

func TestAmadeusService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()
	cfg.Env = config.Test
	suite.Run(t, &amadeusServiceSuite{cfg: *cfg})
}
