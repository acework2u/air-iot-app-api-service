package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const SECRET_KEY = "Saijo@Denki2024#"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func userAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		// Extract the token string.
		parts := strings.Split(idToken, " ")
		if len(parts) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Missing or invalid authorization header",
			})
			return
		}

		// Parse the JWT string.
		tokenString := parts[1]
		_ = tokenString

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			return SECRET_KEY, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()

		//token,err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//      _,ok := token.Method.(*jwt.SigningMethodECDSA)
		//	  if !ok {
		//		  c.Writer.WriteHeader(http.StatusUnauthorized)
		//		  _,err := c.Writer.Write([]byte("You're Unauthorized!"))
		//		  if err != nil {
		//			  return nil, err
		//		  }
		//		  return "",nil
		//	  }
		//})
		// c.Next()

	}
}

//func verifyToken(tokenString string) error {
//	token, err := jwt.Parse([]byte(tokenString), func(token *jwt.Token) (interface{}, error) {
//		return SECRET_KEY, nil
//	})
//	return nil
//}
