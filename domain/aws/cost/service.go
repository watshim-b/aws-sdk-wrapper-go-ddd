package domain

type AWSCostService struct{}

// コスト情報の一覧を取得する
func (s AWSCostService) GetCostList(f AWSCostForm) ([]AWSCost, error) {
	acs := []AWSCost{}

	return acs, nil
}
