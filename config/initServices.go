package config

import (
	"Payment-Service/services"
	"Payment-Service/utils"
	"log"
	"os"
	"sync"
)

var (
	SeerBitInstance services.SeerBit
)
// FIle to Initialize Connetions to External Services

func InitServices() {
	var (
		count int
		wg    sync.WaitGroup
	)

	wg.Add(1)

	go func() {
		defer wg.Done()
		count += InitSeerBit()
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
	privateKey, publicKey := os.Getenv("SEERBIT_PRIVATE_KEY"), os.Getenv("SEERBIT_PUBLIC_KEY")
	if privateKey == "" || publicKey == "" {
		return 0
	}
	key := privateKey + "." + publicKey

	SeerBitInstance.GenerateToken(key, &response, logger)

	if response.Success != true {
		log.Println(reqBuffer, response.Message)
		return 0
	}

	return 1
}