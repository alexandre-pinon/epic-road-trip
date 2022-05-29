package service

import (
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/stretchr/testify/suite"
)

type authServiceSuite struct {
	suite.Suite
	repo        *mocks.UserRepository
	authService AuthService
}

func (suite *authServiceSuite) SetupTest() {
	repo := new(mocks.UserRepository)
	authService := NewAuthService(repo)

	suite.repo = repo
	suite.authService = authService
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}
