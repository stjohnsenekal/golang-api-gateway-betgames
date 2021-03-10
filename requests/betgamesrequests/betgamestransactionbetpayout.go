package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesTransactionBetPayoutRequest(userid string, request *betgamesmodel.BetGamesRequestTransactionBetPayoutRequestParams, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_BET_PAYOUT_REQUEST_MESSAGE_TYPE,
		UserId: userid,
		Data: request,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestTransactionBetPayoutResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestTransactionBetPayoutResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_REQUEST_TRANSACTION_BET_PAYOUT_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
