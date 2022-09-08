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