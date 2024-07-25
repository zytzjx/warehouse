package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zytzjx/warehouse/models"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Auth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwtv5.Parse(tokenString, func(token *jwtv5.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwtv5.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.Users
		models.DB.First(&user, claims["userid"])

		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Auth")
	if err != nil {
		c.Redirect(http.StatusFound, "/index")
		c.Abort()
	}

	token, err := jwtv5.Parse(tokenString, func(token *jwtv5.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
	if claims, ok := token.Claims.(jwtv5.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			//c.AbortWithStatus(http.StatusUnauthorized)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		var user models.Users
		models.DB.First(&user, claims["userid"])

		c.Set("user", user)
		c.Next()

	} else {
		//c.AbortWithStatus(http.StatusUnauthorized)
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
