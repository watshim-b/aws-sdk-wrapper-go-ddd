package usecase

// レスポンスの基底クラス
type ResponseBase struct {

	// エラーに関する情報を保持しておく
	err error
}
