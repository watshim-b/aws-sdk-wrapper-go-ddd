package usecase

// AWSのコスト用のハンドラで利用するform
type AWSCostForm struct {
	FromDate    string `form:"from_date"`
	ToDate      string `form:"to_date"`
	Granularity string `form:"granularity"`
}
