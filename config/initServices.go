package config

import (
	"Payment-Service/services"
	"Payment-Service/utils"
	"log"
	"os"
	"sync"
)

var (
	SeerBitInstance  services.SeerBit
	PaystackInstance services.Paystack
)

// FIle to Initialize Connetions to External Services

func InitServices() {
	var (
		count int
		wg    sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		count += InitSeerBit()
	}()
	go func() {
		defer wg.Done()
		count += InitPayStack()
	}()

	wg.Wait()

	if count < 1 {
		log.Fatalln("No Services Initialized")
	}
}

func InitSeerBit() int {
	var (
		response          = utils.APIResponse{}
		reqBuffer, logger = InitRequestLogger("SeerBit")
	)
	defer log.Println(reqBuffer, response.Message)

	privateKey, publicKey := os.Getenv("SEERBIT_PRIVATE_KEY"), os.Getenv("SEERBIT_PUBLIC_KEY")
	if privateKey == "" || publicKey == "" {
		return 0
	}
	key := privateKey + "." + publicKey

	SeerBitInstance.Initialize(key, &response, logger)

	if response.Success != true {
		return 0
	}

	return 1
}

func InitPayStack() int {
	var (
		response          = utils.APIResponse{}
		reqBuffer, logger = InitRequestLogger("SeerBit")
	)
	defer log.Println(reqBuffer, response.Message)

	key := os.Getenv("PAYSTACK_KEY")

	if key == "" {
		logger.Println("Paystack Token not set")
	} else {
		PaystackInstance.Token = "Bearer " + key
	}

	PaystackInstance.Initialize(&response, logger)
	if response.Success != true {
		return 0
	}

	return 1
}
