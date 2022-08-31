package services

import (
	"Payment-Service/utils"
	"log"
)

type SeerBit struct {
	Token string
}

var SeerBitBaseURL = "https://seerbitapi.com/api/v2"

func (payload *SeerBit) GenerateToken(key string, response *utils.APIResponse, logger *log.Logger) {
	defer utils.HandlePanic(response, logger)

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
	response.Success = true
}
