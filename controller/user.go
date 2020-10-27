package controller

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAccount(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	userService := service.NewUserService()
	u, err := userService.FindUserById(userID)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"account": u,
	})
}

func CreateUser(c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)

	userService := service.NewUserService()
	u, _ := userService.FindUser(user.Email)
	if u != nil {
		// 既にメールアドレスが使用されている場合
		c.JSON(http.StatusOK, gin.H{
			"errorMessage": "入力されたメールアドレスは既に登録されています。",
		})
		return
	}

	u2, err := userService.AddUser(&user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user":    u2.Email,
	})
}

func EditUser(c *gin.Context) {

	accountName := c.PostForm("account_name")
	email := c.PostForm("email")
	introduction := c.PostForm("introduction")
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	file, _, err3 := c.Request.FormFile("image")
	if file != nil && err3 != nil {
		fmt.Println(err3)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err3,
		})
		return
	}
	var buf *bytes.Buffer
	if file != nil {
		defer file.Close()
		buf = bytes.NewBuffer(nil)
		if _, err3 := io.Copy(buf, file); err3 != nil {
			fmt.Println(err3)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ng",
				"err":     err3,
			})
			return
		}
	}

	var user models.User
	user.UserID = userID
	user.AccountName = accountName
	user.Email = email
	user.Introduction = introduction
	if buf != nil {
		user.ContentType.String = http.DetectContentType(buf.Bytes())
		user.AccountImg = buf.Bytes()
	}

	userService := service.NewUserService()

	err := userService.UpdateUser(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user":    user.Email,
	})
}

func FindOrCreateUser(c *gin.Context) {

	user := &models.User{}
	c.ShouldBindJSON(&user)

	userService := service.NewUserService()
	u, err := userService.FindUser(user.Email)

	if u == nil {
		// まだ登録されていない場合
		u2, err2 := userService.AddUser(user)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "NG",
				"err":     err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"user":    u2,
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "NG",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"user":    u,
	})
}
