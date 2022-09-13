package api

import (
	"Payment-Service/config"
	"Payment-Service/services"
	serviceValidator "Payment-Service/services/validators"
	"Payment-Service/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func GetPaymentLink(c *gin.Context) {
	var request utils.JsonMap

	urlPath := strings.Split(c.Request.URL.String(), "/")

	reqType := serviceValidator.PaymentURLPayloads[urlPath[2]]

	//Call Controller
	reqBuffer, reqLogger := config.InitRequestLogger(urlPath[2])
	defer log.Println(reqBuffer)

	//Validate if request body has required fields
	if err := c.BindJSON(&reqType); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	//Bind Request Payload to request type
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	response := services.ServiceImpl{}.GeneratePaymentURL(reqLogger, urlPath[2], request)

	//Return Response
	if response.Success != true {
		c.IndentedJSON(response.Code, response)
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}
