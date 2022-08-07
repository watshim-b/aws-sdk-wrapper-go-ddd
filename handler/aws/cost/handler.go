package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/watshim-b/aws-sdk-wrapper-go-ddd/usecase/aws/cost"
)

func GetCostList(c *gin.Context) {
	// リクエストされた情報をform煮詰める
	acf := usecase.AWSCostForm{}
	err := c.ShouldBindQuery(&acf)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// usecaseの処理を投げる
	cgu := usecase.CostGetUsecase{}
	acs, err := cgu.GetCostList(acf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// レスポンスを返却する
	c.JSON(http.StatusOK, gin.H{
		"data": acs,
	})
}
