package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func CognitoIoTAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")

		splitAuthHeader := strings.Split(idToken, " ")

		if len(splitAuthHeader) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Missing or invalid authorization header",
			})
			return
		}

		// Verified Check Key
		pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
		formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_REGION"), os.Getenv("USER_POOL_ID"))

		keySet, err := jwk.Fetch(c, formattedURL)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Server internal server",
			})
			return
		}

		token, err := jwt.Parse(
			[]byte(splitAuthHeader[1]),
			jwt.WithKeySet(keySet),
			jwt.WithValidate(true),
		)

		_ = token

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the temporary credentials in the AWS config
		//	cfg.Credentials = aws.NewCredentialsCache(stsResp.Credentials)

		// Set the AWS config in the request context
		//	ctx := context.WithValue(c.Request.Context(), aws.ConfigContextKey, &cfg)

		// Continue to the next middleware or handler with the updated context
		// c.Request = c.Request.WithContext(ctx)
		c.Next()

	} // end of return

}
