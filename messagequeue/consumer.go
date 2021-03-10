package messagequeue

import (
	"api-gateway/constants"
	"api-gateway/models/betgamesmodel"
	"api-gateway/responses/betgamesresponses"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

/*
	documentation:
	https://www.rabbitmq.com/tutorials/amqp-concepts.html
*/

/* Asynchronous channels for responses to requests */
var betGamesBaseResponseChannel chan betgamesmodel.BetGamesBaseResponse
var betGamesErrorResponse chan betgamesmodel.BetGamesBaseResponse
var betGamesRequestNewTokenResponseChannel chan betgamesmodel.BetGamesRequestNewTokenResponse
var betGamesRequestGetBalanceResponse chan betgamesmodel.BetGamesRequestGetBalanceResponse
var betGamesAccountDetailsResponse chan betgamesmodel.BetGamesAccountDetailsResponse
var betGamesRequestTransactionBetPayinResponse chan betgamesmodel.BetGamesRequestTransactionBetPayinResponse
var betGamesRequestTransactionBetSubscriptionPayinResponse chan betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse
var betGamesRequestTransactionBetPayoutResponse chan betgamesmodel.BetGamesRequestTransactionBetPayoutResponse
var betGamesRequestTransactionBetCombinationPayinResponse chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse
var betGamesRequestTransactionBetCombinationPayoutResponse chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponse

func SetChannelBetGamesBaseResponse(provided chan betgamesmodel.BetGamesBaseResponse) {
	betGamesBaseResponseChannel = provided
}

func SetChannelBetGamesRequestNewTokenResponse(provided chan betgamesmodel.BetGamesRequestNewTokenResponse) {
	betGamesRequestNewTokenResponseChannel = provided
}

func SetChannelBetGamesRequestGetBalanceResponse(provided chan betgamesmodel.BetGamesRequestGetBalanceResponse) {
	betGamesRequestGetBalanceResponse = provided
}

func SetChannelBetGamesAccountDetailsResponse(provided chan betgamesmodel.BetGamesAccountDetailsResponse) {
	betGamesAccountDetailsResponse = provided
}

func SetChannelBetGamesRequestTransactionBetPayinResponse(provided chan betgamesmodel.BetGamesRequestTransactionBetPayinResponse) {
	betGamesRequestTransactionBetPayinResponse = provided
}

func SetChannelBetGamesRequestTransactionBetSubscriptionPayinResponse(provided chan betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse) {
	betGamesRequestTransactionBetSubscriptionPayinResponse = provided
}

func SetChannelBetGamesRequestTransactionBetPayoutResponse(provided chan betgamesmodel.BetGamesRequestTransactionBetPayoutResponse) {
	betGamesRequestTransactionBetPayoutResponse = provided
}

func SetChannelBetGamesRequestTransactionBetCombinationPayinResponse(provided chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse) {
	betGamesRequestTransactionBetCombinationPayinResponse = provided
}

func SetChannelBetGamesRequestTransactionBetCombinationPayoutResponse(provided chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponse) {
	betGamesRequestTransactionBetCombinationPayoutResponse = provided
}

func SetChannelBetGamesErrorResponse(provided chan betgamesmodel.BetGamesBaseResponse) {
	betGamesErrorResponse = provided
}

func Subscribe(topic string, queueName string)  {

	conn, err := amqp.Dial(getConnectionString())
	failOnError(err, "API CONSUMER failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "API CONSUMER failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		env.RabbitMq_Exchange_Name,
		env.RabbitMq_Exchange_Type,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "API CONSUMER failed to declare an exchange")

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "API CONSUMER failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		topic,
		env.RabbitMq_Exchange_Name,
		false,
		nil)
	failOnError(err, "API CONSUMER failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "API CONSUMER failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("API CONSUMER received a message from backend: %s", d.Body)
			routeResponse(d.Body)
		}
	}()

	log.Printf("API CONSUMER on topic %s started. Waiting for messages.", topic)
	<-forever
}

type Response struct {
	MessageType		string `json:"message_type,omitempty"`
	UserId			string `json:"user_id,omitempty"`
	Data			json.RawMessage `json:"data,omitempty"`
}

func routeResponse(t []uint8) {

	response := Response{}
	err := json.Unmarshal([]byte(t), &response)
	if err != nil {
		log.Fatal(err)
	}

	switch response.MessageType {

	/* betgames/api */
	case constants.BETGAMES_REFRESH_TOKEN_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.refresh.token: %s", response)
		betgamesresponses.RefreshToken(response.UserId, response.Data, betGamesBaseResponseChannel)
	case constants.BETGAMES_REQUEST_NEW_TOKEN_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.new.token: %s", response)
		betgamesresponses.RequestNewToken(response.UserId, response.Data, betGamesRequestNewTokenResponseChannel)
	case constants.BETGAMES_ACCOUNT_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.request.account.details: %s", response)
		betgamesresponses.GetAccountDetailsResponse(response.UserId, response.Data, betGamesAccountDetailsResponse)
	case constants.BETGAMES_GET_BALANCE_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.request.account.balance: %s", response)
		betgamesresponses.GetBalanceResponse(response.UserId, response.Data, betGamesRequestGetBalanceResponse)
	case constants.BETGAMES_BET_PAYIN_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.transaction.payin: %s", response)
		betgamesresponses.TransactionBetPayinResponse(response.UserId, response.Data, betGamesRequestTransactionBetPayinResponse)
	case constants.BETGAMES_BET_SUBSCRIPTION_PAYIN_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.subscription.payin: %s", response)
		betgamesresponses.SubscriptionBetPayinResponse(response.UserId, response.Data, betGamesRequestTransactionBetSubscriptionPayinResponse)
	case constants.BETGAMES_BET_PAYOUT_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.transaction.payout: %s", response)
		betgamesresponses.TransactionBetPayoutResponse(response.UserId, response.Data, betGamesRequestTransactionBetPayoutResponse)
	case constants.BETGAMES_BET_COMBINATION_PAYIN_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.combination.payin: %s", response)
		betgamesresponses.CombinationBetPayinResponse(response.UserId, response.Data, betGamesRequestTransactionBetCombinationPayinResponse)
	case constants.BETGAMES_BET_COMBINATION_PAYOUT_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.combination.payout: %s", response)
		betgamesresponses.CombinationBetPayoutResponse(response.UserId, response.Data, betGamesRequestTransactionBetCombinationPayoutResponse)
	case constants.BETGAMES_ERROR_RESPONSE_MESSAGE_TYPE :
		log.Printf("BACKEND RESPONSE RECEIVED: betgames.api.response.error: %s", response)
		betgamesresponses.ErrorResponse(response.UserId, response.Data, betGamesErrorResponse)
	}


}


