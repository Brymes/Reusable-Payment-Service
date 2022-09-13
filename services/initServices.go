package services

import (
	"Payment-Service/config"
	"Payment-Service/utils"
	"log"
	"os"
	"sync"
)

var (
	SeerBitInstance  SeerBit
	PaystackInstance Paystack
)

// FIle to Initialize Connetions to External ServiceImpl

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
		log.Fatalln("No Service Initialized")
	}
}

func InitSeerBit() int {
	var (
		response          = utils.APIResponse{}
		reqBuffer, logger = config.InitRequestLogger("SeerBit")
	)
	//defer utils.HandlePanic(&response, logger, "Unable to Initialize SeerBit")
	defer log.Println(reqBuffer, response.Message)

	privateKey, publicKey := os.Getenv("SEERBIT_PRIVATE_KEY"), os.Getenv("SEERBIT_PUBLIC_KEY")
	if privateKey == "" || publicKey == "" {
		logger.Println("SeerBit Private and Public Key not Set")
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
		reqBuffer, logger = config.InitRequestLogger("Paystack")
	)
	//defer utils.HandlePanic(&response, logger, "Unable to Initialize PayStack")
	defer log.Println(reqBuffer, response.Message)

	key := os.Getenv("PAYSTACK_KEY")

	if key == "" {
		logger.Println("Paystack Token not set")
		return 0
	} else {
		PaystackInstance.Token = "Bearer " + key
	}

	PaystackInstance.Initialize(&response, logger)
	if response.Success != true {
		return 0
	}

	return 1
}
