package domain

// エンティティの他に値オブジェクトを定義してこのクラスの変数として利用してもよい
// ドメイン層からUsecase層に戻す際に利用するオブジェクト
type AWSCost struct {

	// 金額
	Amount int
}
