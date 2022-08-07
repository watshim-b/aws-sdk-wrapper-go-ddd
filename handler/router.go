package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	aws_cost_handler "github.com/watshim-b/aws-sdk-wrapper-go-ddd/handler/aws/cost"
)

// ルーティング
type router struct {
	e *gin.Engine
}

func NewRouter() *router {
	r := router{}
	r.e = gin.Default()
	return &r
}

// 設定情報を取得する
func (r *router) SetConfig() {

	// 設定を行うためのインスタンスを生成
	c := cors.DefaultConfig()

	// クロスオリジン設定
	c.AllowOrigins = []string{"http://localhost:3000"}

	// リクエストを受け付けるメソッド
	c.AllowMethods = []string{"GET"}

	// リクエストを受け付けるヘッダー
	c.AllowHeaders = []string{"*"}

	// cookieなどの情報を必要とするかどうか
	c.AllowCredentials = false

	// preflightリクエストの結果をキャッシュする時間
	c.MaxAge = 24 * time.Hour

	r.e.Use(cors.New(c))
}

// ルーティングの設定を行う
func (r *router) SetRouting() {

	// テスト用のピング
	r.e.GET("ping", ping)
	r.e.GET("/aws/cost/list", aws_cost_handler.GetCostList)
}

func (r *router) Run(port string) {
	r.e.Run(port)
}
