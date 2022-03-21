package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTValidator(l *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			l.Error("token is required")
			c.JSON(http.StatusUnauthorized, "Unauthorized: token is required")
			c.Abort()
			return
		}

		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)

		// Parse takes the token string and a function for looking up the key.
		// The latter is especially useful if you use multiple keys for your application.
		// The standard is to use 'kid' in the head of the token to identify
		// which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			l.Error(err)
			c.JSON(http.StatusUnauthorized, "Unauthorized: "+err.Error())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("filial", int(claims["filial"].(float64)))
		} else {
			l.Error("token is invalid")
			c.JSON(http.StatusUnauthorized, "Unauthorized: token is invalid")
			c.Abort()
			return
		}

		c.Next()
	}
}
