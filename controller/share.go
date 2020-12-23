package controller

import (
	"fmt"
	"strconv"

	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ShareForm struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// CreateShareToken 限定公開用URLトークン作成
func CreateShareToken(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	shareService := service.NewShareService()
	shareToken := shareService.CreateShareToken(userID)

	c.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"share_token": shareToken,
	})
}

// RegisteShareToken 限定公開用URLのトークンを登録する
func RegisteShareToken(c *gin.Context) {

	var shareForm ShareForm
	c.BindJSON(&shareForm)

	shareService := service.NewShareService()
	err := shareService.RegisteShareToken(shareForm.UserID, shareForm.Token)

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
	})
}

// FindPublicUser 公開用URLトークンからユーザー🆔を取得する
func FindPublicUser(c *gin.Context) {

	token := c.Query("token")

	shareService := service.NewShareService()
	userID, err := shareService.FindPublicUser(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user_id": userID,
	})
}
