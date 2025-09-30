package token

import "github.com/g3techlabs/revit-api/config"

var accessTokenSecret string = config.Get("ACCESS_SECRET")
var refreshTokenSecret string = config.Get("REFRESH_SECRET")
var accessTokenExpirationInHours int = config.GetIntVariable("ACCESS_TOKEN_EXPIRATION")
var refreshTokenExpirationInHours int = config.GetIntVariable("REFRESH_TOKEN_EXPIRES_IN")
var resetPasswordTokenSecret string = config.Get("RESET_PASSWORD_SECRET")
