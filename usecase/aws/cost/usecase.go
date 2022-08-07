package usecase

import (
	domain "github.com/watshim-b/aws-sdk-wrapper-go-ddd/domain/aws/cost"
)

type CostGetUsecase struct{}

func (u CostGetUsecase) GetCostList(acf AWSCostForm) (AWSCostResponse, error) {
	acr := AWSCostResponse{}

	// domain層にリクエストを連携する
	domainForm := domain.AWSCostForm{}
	domainService := domain.AWSCostService{}
	cs, err := domainService.GetCostList(domainForm)
	if err != nil {
		return acr, err
	}

	// レスポンスのデータを設定
	acr.Amount = cs[0].Amount

	return acr, nil
}
