package services

import (
	"Payment-Service/utils"
	"log"
)

type ServiceImpl struct{}

func (s ServiceImpl) GeneratePaymentURL(logger *log.Logger, serviceName string, payload utils.JsonMap) (response utils.APIResponse) {
	response = utils.APIResponse{Success: false}

	defer utils.HandlePanic(&response, logger, "")

	service, found := AllServices[serviceName]

	if found != true {
		response.Code = 400
		panic("Unknown Service Selected")
	}
	service.GeneratePaymentURL(&response, logger, payload)

	return
}

type Service interface {
	GeneratePaymentURL(response *utils.APIResponse, logger *log.Logger, requestBody utils.JsonMap)
}

var AllServices = map[string]Service{
	"seerbit":  &SeerBitInstance,
	"paystack": &PaystackInstance,
}
