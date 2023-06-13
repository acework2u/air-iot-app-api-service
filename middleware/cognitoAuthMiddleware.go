package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	// "github.com/niclabs/go-nic/aws/cognito"
)

func CognitoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Cognito ID token from the request headers
		idToken := c.GetHeader("Authorization")

		splitAuthHeader := strings.Split(idToken, " ")

		fmt.Println(os.Getenv("USER_POOL_ID"))

		fmt.Println("Middleware")

		if len(splitAuthHeader) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Missing or invalid authorization header",
			})
			return
		}

		// fmt.Println("splitAuthHeader")
		// fmt.Println(splitAuthHeader[1])

		pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
		formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_REGION"), os.Getenv("USER_POOL_ID"))

		keySet, err := jwk.Fetch(c, formattedURL)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Server internal server",
			})
			return
		}

		fmt.Println("keyset")
		fmt.Println(keySet)

		token, err := jwt.Parse(
			[]byte(splitAuthHeader[1]),
			jwt.WithKeySet(keySet),
			jwt.WithValidate(true),
		)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		username, _ := token.Get("cognito:username")

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": username,
		})

		fmt.Printf("The username: %v\n", username)
		fmt.Println(token)

		// fmt.Println(splitAuthHeader)

		// cognitoClient := service.NewCognitoService()

		// fmt.Println(cognitoClient)

		// if !ok {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"status":  http.Status InternalServerError,
		// 		"message": ok,
		// 	})

		// 	return
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"status":  http.StatusUnauthorized,
		// 		"message": cognitoClient,
		// 	})
		// 	return

		// }

		// Verify the Cognito ID token
		// _, err := cognito.VerifyIDToken(idToken, "YOUR_COGNITO_USER_POOL_ID", "YOUR_COGNITO_REGION")

		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	c.Abort()
		// 	return
		// }

		// Continue to the next middleware or handler
		c.Next()
	}
}
