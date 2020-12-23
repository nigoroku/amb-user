package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"

	// "github.com/kzpolicy/user/controller"
	// "github.com/kzpolicy/user/middleware"
	"local.packages/controller"
	"local.packages/db"
	"local.packages/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --output models --pkgname models --wipe mysql

var identityKey = "id"

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(middleware.RecordUaAndTime)

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	authMiddleware := middleware.AuthJwt()

	// 存在しないURLの場合
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// DB接続
	db.Init()

	// ルーティング
	UserRoute := r.Group("/api/v1")
	{
		v1 := UserRoute.Group("/user")
		{
			v1.GET("/account", controller.FindAccount)
			v1.POST("/create", controller.CreateUser)
			v1.POST("/edit", controller.EditUser)
			v1.POST("/find-or-create", controller.FindOrCreateUser)
			v1.POST("/login", authMiddleware.LoginHandler)

			v1.GET("/share", controller.CreateShareToken)
			v1.POST("/share/registe", controller.RegisteShareToken)
			v1.GET("/public", controller.FindPublicUser)

			auth := v1.Group("/auth")
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)
			auth.Use(authMiddleware.MiddlewareFunc())
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run(":8081")
}
