package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const googleApiBaseUrl = "https://oauth2.googleapis.com"

type GoogleToken struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Iat string `json:"iat"`
	Exp string `json:"exp"`

	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
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

		ctx.Set("token", tokenID)
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
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "unauthorized",
	})
}
