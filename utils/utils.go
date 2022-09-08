package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"runtime/debug"
)

//House Utility Functions in a central file or special files

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"-"`
	Data    interface{} `json:"data"`
}

type JsonMap map[string]interface{}

func GenerateUniqueID(length int) string {
	bytes := make([]byte, length)

	chars := "0123456789abcdefghijklmnopqrstuvwxyz"

	if _, err := rand.Read(bytes); err != nil {
		panic("Internal Server Error whilst generating Reference")
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}

func HandlePanic(response *APIResponse, logger *log.Logger, customMessage string) {
	if err := recover(); err != nil {
		logger.Println(err)
		logger.Println(string(debug.Stack()))

		if response.Code < 400 {
			response.Code = 500
		}

		response.Success = false
		if customMessage != "" {
			response.Message = customMessage
		} else {
			response.Message = fmt.Sprintf("%v", err)
		}
	}
}
