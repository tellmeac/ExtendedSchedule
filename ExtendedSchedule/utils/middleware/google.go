package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const ContextTokenKey = "token"

const googleApiBaseUrl = "https://oauth2.googleapis.com"

type GoogleToken struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Iat int    `json:"iat"`
	Exp int    `json:"exp"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func (g GoogleToken) Valid() error {
	return nil
}

func GetGoogleEmail(ctx *gin.Context) (string, error) {
	jwtToken := ctx.GetString(ContextTokenKey)
	if jwtToken == "" {
		return "", errors.New("tokenID not found in context")
	}

	var result GoogleToken
	token, err := jwt.ParseWithClaims(jwtToken, result, nil)
	if err != nil {
		return "", err
	}
	return token.Claims.(GoogleToken).Email, nil
}

func GoogleOAuth2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.Request.Header.Get("Authorization")

		if bearer == "" {
			handleForbidden(ctx)
			return
		}

		tokenID := strings.TrimPrefix(bearer, "Bearer ")

		if !validate(tokenID) {
			handleForbidden(ctx)
			return
		}

		ctx.Set(ContextTokenKey, tokenID)
	}
}

func validate(tokenID string) bool {
	requestUrl := googleApiBaseUrl + "/tokeninfo?id_token=" + tokenID

	r, err := http.Get(requestUrl)

	if err == nil && r.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func handleForbidden(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
}
