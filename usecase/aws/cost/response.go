package usecase

// AWSのコストに関してのレスポンス用クラス
type AWSCostResponse struct {

	// レスポンス用のデータに指定がない場合は、interface型で受け付けられるようにしておく。
	Amount int
}
