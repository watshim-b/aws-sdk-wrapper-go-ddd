package repository

type awsCostRepository struct{}

// コネクションをはるので、errorを返す可能性がある。
func NewAWSCostRepository() (*awsCostRepository, error) {
	r := awsCostRepository{}
	return &r, nil
}

// コスト情報の一覧を取得する
func (r *awsCostRepository) GetList(asc AWSCostSearchCondition) ([]AWSCostSearchResult, error) {
	acsrs := []AWSCostSearchResult{}

	return acsrs, nil
}
