package domain

import (
	repo "github.com/watshim-b/aws-sdk-wrapper-go-ddd/repository/aws/cost"
)

type AWSCostService struct{}

// コスト情報の一覧を取得する
func (s AWSCostService) GetCostList(f AWSCostForm) ([]AWSCost, error) {
	acs := []AWSCost{}

	// 検索条件を生成する
	repositorySearchCondition := repo.AWSCostSearchCondition{}

	// 検索用のリポジトリを作成する
	repoitory, err := repo.NewAWSCostRepository()
	if err != nil {
		return nil, err
	}

	// 検索をかける
	repoResponses, err := repoitory.GetList(repositorySearchCondition)
	if err != nil {
		return nil, err
	}

	// TODO 検索結果を返却する形に変換する
	for repoResponse := range repoResponses {
		ac := AWSCost{}
		print(repoResponse)
		acs = append(acs, ac)
	}

	ac := AWSCost{
		Amount: 100,
	}
	acs = append(acs, ac)

	return acs, nil
}
