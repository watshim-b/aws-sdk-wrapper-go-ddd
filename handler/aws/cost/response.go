package handler

import (
	"github.com/watshim-b/aws-sdk-wrapper-go-ddd/handler"
)

// AWSのコストに関してのレスポンス用クラス
type AWSCostResponse struct {

	// レスポンス用のデータに指定がない場合は、interface型で受け付けられるようにしておく。
	data interface{}

	// 基底クラスの定義を移譲する
	handler.ResponseBase
}
