package main

import (
	"api-gateway/config"
	"api-gateway/handlers/routes"
	"api-gateway/healthutility"
	"api-gateway/constants"
	"api-gateway/logoutput"
	"api-gateway/messagequeue"
	"api-gateway/redisutility"
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {

	fmt.Println("--------API GATEWAY init--------")

	configuration := config.Configuration{}
	err := gonfig.GetConf(getConfigFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	t := time.Now()

	var sb strings.Builder
	sb.WriteString(configuration.Log_file_path)
	sb.WriteString("api_log_")
	sb.WriteString(t.Format("2006-01-02"))
	sb.WriteString(".txt")

	logoutput.SetLogFileLocation(sb.String())

	// https://blog.kowalczyk.info/article/1Ll7/rotate-log-files-daily-in-go.html
	// https://stackoverflow.com/questions/28796021/how-can-i-log-in-golang-to-a-file-with-log-rotation
	// ^ maybe use lumberjack package or log rotation

	f, err := os.OpenFile(sb.String(), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	//defer f.Close()

	//log.SetOutput(f)
	logoutput.Set(f)
	defer logoutput.Close() // call from main only

	log.Println("--------API GATEWAY init--------")

	log.Printf("Application refresh, setting log file to : %s", logoutput.GetLogFileLocation())

	health := healthutility.Health{
		ServiceName: "api-gateway",
		Status:      healthutility.SERVICE_STATUS_UP,
	}

	monitor := healthutility.Monitor{
		&health,
	}

	health.ServiceStarted()
	//
	////http://localhost:8085/health/
	//
	go monitor.Listen(configuration.HealthCheckPoint)

	redisutility.SetConfiguration(configuration)
	messagequeue.SetConfiguration(configuration)

	go messagequeue.Subscribe(constants.BETGAMES_REQUEST_SUBSCRIPTION_BET_PAYIN_TOPIC,
		"betgames-api-response-subscription-payin")

	setupConsumers()

	time.Sleep(5 * time.Second)

	routes.SetupRoutes(configuration, &health)

}

func setupConsumers() {

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_REFRESH_TOKEN_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_REFRESH_TOKEN)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_NEW_TOKEN_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_NEW_TOKEN)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_GET_ACCOUNT_DETAILS_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_ACCOUNT_DETAILS)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_GET_BALANCE_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_GET_BALANCE)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_TRANSACTION_BET_PAYIN_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_TRANSACTION_PAYIN)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_SUBSCRIPTION_BET_PAYIN_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_SUBSCRIPTION_PAYIN)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_TRANSACTION_BET_PAYOUT_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_TRANSACTION_PAYOUT)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_COMBINATION_BET_PAYIN_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_COMBINATION_PAYIN)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_COMBINATION_BET_PAYOUT_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_COMBINATION_PAYOUT)

	go messagequeue.Subscribe(constants.BETGAMES_RESPONSE_GENERAL_ERROR_TOPIC,
		constants.BETGAMES_QUEUE_FOR_SUBSCRIBER_GENERAL_ERROR)

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "production"
	}
	filename := []string{"config/", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}

