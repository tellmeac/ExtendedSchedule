package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

// contextTokenKey locally defines what key should be used by middleware to store verified JWT token string.
const contextTokenKey = "token"

const googleApiBaseUrl = "https://oauth2.googleapis.com"

// GoogleToken claims.
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

// GetGoogleEmail returns google email, extracted from jwt token from Authorization header in request.
func GetGoogleEmail(ctx *gin.Context) (string, error) {
	jwtToken := ctx.GetString(contextTokenKey)
	if jwtToken == "" {
		return "", errors.New("tokenID not found in context")
	}

	var result GoogleToken
	_, _, err := jwt.NewParser().ParseUnverified(jwtToken, &result)
	if err != nil {
		return "", err
	}
	return result.Email, nil
}

// GoogleOAuth2 returns middleware handlers to prevent unauthorized access to API methods.
func GoogleOAuth2(debug bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.Request.Header.Get("Authorization")

		if bearer == "" {
			handleForbidden(ctx)
			return
		}

		tokenID := strings.TrimPrefix(bearer, "Bearer ")

		if !debug && !validate(tokenID) {
			handleForbidden(ctx)
			return
		}

		ctx.Set(contextTokenKey, tokenID)
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
