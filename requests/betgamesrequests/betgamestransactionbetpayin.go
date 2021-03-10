package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesTransactionBetPayinRequest(userid string, request *betgamesmodel.BetGamesRequestTransactionBetPayinRequestParams, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_BET_PAYIN_REQUEST_MESSAGE_TYPE,
		UserId: userid,
		Data: request,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestTransactionBetPayinResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestTransactionBetPayinResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_REQUEST_TRANSACTION_BET_PAYIN_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
