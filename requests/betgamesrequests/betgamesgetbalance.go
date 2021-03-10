package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesGetBalanceRequest(userid string, method string, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	data := map[string]string{
		"user_id" : userid,
	}

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_GET_BALANCE_REQUEST_MESSAGE_TYPE,
		UserId: userid,
		Data: data,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestGetBalanceResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestGetBalanceResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_REQUEST_GET_BALANCE_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
