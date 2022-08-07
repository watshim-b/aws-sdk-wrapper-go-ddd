package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	// リクエストされた情報をform煮詰める
	acf := AWSCostForm{}
	err := c.ShouldBindQuery(&acf)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// usecaseの処理を投げる
	u := uc.NewCostUsecase()
	lbs, err := u.GetCostAndUsage(acf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// レスポンスを返却する
	c.JSON(http.StatusOK, gin.H{
		"lfs": lbs,
	})
}
