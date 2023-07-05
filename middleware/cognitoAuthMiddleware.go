package middleware

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/niclabs/go-nic/aws/cognito"
)

type UserInfo struct {
	Username              string    `json:"username"`
	Email                 string    `json:"email"`
	Email_verified        bool      `json:"email_verifyed"`
	Exp                   time.Time `json:"exp"`
	Iat                   time.Time `json:"iat"`
	Iss                   string    `json:"iss"`
	Jti                   string    `json:"jti"`
	Origin_jti            string    `json:"origin_jti"`
	Phone_number          string    `json:"phone_number"`
	Phone_number_verified bool      `json:"phone_number_verified"`
	Sub                   string    `json:"sub"`
	Token_use             string    `json:"token_use"`
}

func (u *UserInfo) GetVal(name string) any {
	switch name {
	default:
		return "no data"
	case "name":
		return &u.Username
	}

}

func CognitoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Extract the Cognito ID token from the request headers
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
		if err != nil {

			// txtErr := strings.Split(err.Error(), ":")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			// c.Abort()
			return
		}

		username, _ := token.Get("cognito:username")
		useremail, _ := token.Get("email")
		useremail_verified, _ := token.Get("email_verified")
		exp, _ := token.Get("exp")
		iat, _ := token.Get("iat")
		iss, _ := token.Get("iss")
		jti, _ := token.Get("jti")
		origin_jti, _ := token.Get("origin_jti")
		phone_number, _ := token.Get("phone_number")
		phone_number_verified, _ := token.Get("phone_number_verified")
		sub, _ := token.Get("sub")
		token_use, _ := token.Get("token_use")

		myData := &UserInfo{
			Username:              username.(string),
			Email:                 useremail.(string),
			Email_verified:        useremail_verified.(bool),
			Exp:                   exp.(time.Time),
			Iat:                   iat.(time.Time),
			Iss:                   iss.(string),
			Jti:                   jti.(string),
			Origin_jti:            origin_jti.(string),
			Phone_number:          phone_number.(string),
			Phone_number_verified: phone_number_verified.(bool),
			Sub:                   sub.(string),
			Token_use:             token_use.(string),
		}

		_ = myData

		// Set Variable to Context
		c.Set("UserToken", token)
		c.Set("UserId", username.(string))
		c.Set("UserEmail", useremail)
		c.Set("UserEmailVerified", useremail_verified)
		c.Set("UserPhoneVerified", phone_number_verified)
		c.Set("UserPhone", phone_number)
		c.Set("UserSub", sub)
		c.Set("UserIat", iat)
		c.Set("UserExp", exp)

		// c.JSON(http.StatusOK, gin.H{
		// 	"status":  http.StatusOK,
		// 	"message": username,
		// })

		c.Next()
	}
}
