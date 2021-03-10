package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesCombinationBetPayinRequest(userid string, request *betgamesmodel.BetGamesRequestTransactionBetCombinationPayinRequestParams, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_BET_COMBINATION_PAYIN_REQUEST_MESSAGE_TYPE,
		UserId: userid,
		Data: request,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestTransactionBetCombinationPayinResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_REQUEST_COMBINATION_BET_PAYIN_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
