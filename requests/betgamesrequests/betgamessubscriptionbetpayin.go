package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesSubscriptionBetPayinRequest(userid string, request *betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinRequestParams, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_BET_SUBSCRIPTION_PAYIN_REQUEST_MESSAGE_TYPE,
		UserId: userid,
		Data: request,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestTransactionBetSubscriptionPayinResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_REQUEST_SUBSCRIPTION_BET_PAYIN_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
