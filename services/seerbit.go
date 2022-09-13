package services

import (
	"Payment-Service/utils"
	"log"
)

type SeerBit struct {
	Token, PublicKey string
}

var SeerBitBaseURL = "https://seerbitapi.com/api/v2"

func (payload *SeerBit) Initialize(key string, response *utils.APIResponse, logger *log.Logger) {
	defer utils.HandlePanic(response, logger, "Unable to Initialise SeerBit: Cannot reach payment service")

	body := utils.JsonMap{
		"key": key,
	}
	req := utils.RequestPayload{
		Url:    SeerBitBaseURL + "/encrypt/keys",
		Auth:   "",
		Method: "POST",
		Body:   body,
	}

	resp := req.MakeRequest(response, logger)

	payload.Token = resp["data"].(map[string]string)["EncryptedSecKey"]

	if payload.Token == "" || resp["status"].(string) != "SUCCESS" {
		logger.Println(response)
		response.Success, response.Code, response.Message = false, 500, "Unable to Initialise SeerBit"
	}

	response.Success, response.Message = true, "Successfully Initialised SeerBit"
}

func (payload *SeerBit) GeneratePaymentURL(response *utils.APIResponse, logger *log.Logger, requestBody utils.JsonMap) {
	//defer utils.HandlePanic(response, logger, "")

	requestBody["publicKey"] = payload.PublicKey

	if requestBody["paymentReference"] == "" {
		paymentRef := utils.GenerateUniqueID(10)
		requestBody["paymentReference"] = paymentRef
	}

	req := utils.RequestPayload{
		Url:    SeerBitBaseURL + "/payments",
		Auth:   payload.Token,
		Method: "POST",
		Body:   requestBody,
	}

	resp := req.MakeRequest(response, logger)

	url := resp["data"].(utils.JsonMap)["payments"].(utils.JsonMap)["redirectLink"]

	if resp["status"].(string) != "SUCCESS" || url == "" {
		logger.Println(resp)
		response.Success, response.Code, response.Message = false, 500, "Temporarily Unable to generate Payment URL"
		return
	}

	res := utils.JsonMap{
		"url":              url,
		"paymentReference": requestBody["paymentReference"],
	}

	response.Data, response.Success, response.Message = res, true, "Successfully Generated PaymentURL"

	return

}
