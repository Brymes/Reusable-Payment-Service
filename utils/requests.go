package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestPayload struct {
	Url, Auth, Method string
	Body              JsonMap
}

func (request RequestPayload) MakeRequest(response *APIResponse, logger *log.Logger) JsonMap {

	var (
		req    *http.Request
		err    error
		result JsonMap
	)

	if request.Body != nil {

		jsonPayload, err := json.Marshal(request.Body)

		if err != nil {
			logger.Println(err)
			panic(err)
		}

		body := bytes.NewBuffer(jsonPayload)
		req, err = http.NewRequest(request.Method, request.Url, body)
	} else {
		req, err = http.NewRequest(request.Method, request.Url, nil)

	}

	req.Header.Add("Content-Type", "application/json")
	if request.Auth != "" {
		req.Header.Add("Authorization", request.Auth)
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(fmt.Sprintf("Error Making Request to Payment Service: %v", err))
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Printf("Error Reading Response from Payment Service: %v", err)
		panic("Error Reading Response from Payment Service")
	}

	// FIXME for empty Body response
	err = json.Unmarshal(respBody, &result)

	if err != nil {
		logger.Println(err)
		panic("Error Reading Response from Payment Service")
	}

	// Used Switch for granular error management in the future
	switch resp.StatusCode {
	case 200, 201:
		break
	case 400, 401, 403:
		logger.Println(fmt.Sprintf("Unauthorised Request to Payment Service: %v", resp.StatusCode))
		panic("Unauthorised Request to Payment Service")
	case 500, 501, 503:
		logger.Println(fmt.Sprintf("Payment Service is Down: %v", resp.StatusCode))
		panic("Payment Service is Down")
	default:
		panic(fmt.Sprintf("Error Making Request to Payment Service: %v", resp.StatusCode))
	}

	return result
}
