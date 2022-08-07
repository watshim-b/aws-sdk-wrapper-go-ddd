package usecase

import (
	handler "github.com/watshim-b/aws-sdk-wrapper-go-ddd/handler/aws/cost"
)

type CostGetUsecase struct{}

func (u CostGetUsecase) Execute(handlerForm handler.AWSCostForm) (handler.AWSCostResponse, error) {
	handlerResponse := handler.AWSCostResponse{}

	return handlerResponse, nil
}
