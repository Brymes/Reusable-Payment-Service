package services

import (
	"Payment-Service/utils"
	"log"
)

var paystackBaseURL = "https://api.paystack.co/"

type Paystack struct{ Token string }

func (payload *Paystack) Initialize(response *utils.APIResponse, logger *log.Logger) {
	defer utils.HandlePanic(response, logger, "Unable to Initialise PayStack: Cannot reach payment service")

	req := utils.RequestPayload{
		Url:    paystackBaseURL + "integration/payment_session_timeout",
		Auth:   payload.Token,
		Method: "GET",
		Body:   nil,
	}
	resp := req.MakeRequest(response, logger)

	if resp["status"].(bool) != true {
		logger.Println(resp)
		response.Success, response.Message = false, "Unable to Initialise Paystack :Invalid API Key"
		return
	}
	response.Success, response.Message = true, "Successfully Initialised Paystack"

}

func (payload *Paystack) GeneratePaymentURL(response *utils.APIResponse, logger *log.Logger, requestBody utils.JsonMap) {

	if requestBody["paymentReference"] == "" {
		paymentRef := utils.GenerateUniqueID(10)
		requestBody["paymentReference"] = paymentRef
	}

	req := utils.RequestPayload{
		Url:    paystackBaseURL + "transaction/initialize",
		Auth:   payload.Token,
		Method: "POST",
		Body:   requestBody,
	}

	resp := req.MakeRequest(response, logger)

	if resp["status"].(bool) != true {
		logger.Println(resp)
		response.Success, response.Code, response.Message = false, 500, "Temporarily Unable to generate Payment URL"
		return
	}

	url := (resp["data"].(map[string]interface{}))["authorization_url"].(string)

	res := utils.JsonMap{
		"url":              url,
		"paymentReference": requestBody["paymentReference"],
	}

	response.Data, response.Success, response.Message = res, true, "Successfully Generated PaymentURL"
	return

}
