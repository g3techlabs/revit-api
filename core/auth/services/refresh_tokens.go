package services

import (
	"github.com/g3techlabs/revit-api/core/auth/response"
	"github.com/g3techlabs/revit-api/utils/generics"
)

func (as *AuthService) RefreshTokens(refreshToken string) (*response.AuthTokensResponse, error) {
	claims, err := as.tokenService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, generics.Unauthorized("Invalid or expired token")
	}

	userId := claims.UserID
	user, err := as.userRepo.FindUserById(userId)
	if err != nil {
		return nil, generics.InternalError()
	} else if user == nil {
		return nil, generics.Unauthorized("Not authenticated")
	}

	accessToken, refreshToken, err := as.tokenService.GenerateAuthTokens(userId)
	if err != nil {
		return nil, generics.InternalError()
	}

	return &response.AuthTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
