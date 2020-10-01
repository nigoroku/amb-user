package middleware

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"local.packages/models"
	"local.packages/service"
)

var identityKey = "id"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthJwt() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey:   v.UserID,
					"accountName": v.AccountName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				UserID:      claims[identityKey].(int),
				AccountName: claims["accountName"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals LoginUser
			if err := c.ShouldBind(&loginVals); err != nil {
				fmt.Println("err:", err)
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			userService := service.UserService{}
			registeredUser, err2 := userService.FindRegisteredUser(email, password)
			if err2 != nil {
				fmt.Println("err:", err2)
				return nil, jwt.ErrFailedAuthentication
			}

			return &models.User{
				Email:       registeredUser.Email,
				UserID:      registeredUser.UserID,
				AccountName: registeredUser.AccountName,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// TODO:DBから取得
			if v, ok := data.(*models.User); ok && v.Email == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	if err != nil {
		fmt.Println(err)
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
