package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"time"
)

func BetgamesRequestNewTokenRequest(userid string, token string, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	data := &betgamesmodel.BetGamesRequestNewTokenResponseParams{
		NewToken:    token,
	}

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_REQUEST_NEW_TOKEN_MESSAGE_TYPE,
		UserId: userid,
		Data: data,
	}

	messageResponse := make(chan betgamesmodel.BetGamesRequestNewTokenResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesRequestNewTokenResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_RESPONSE_NEW_TOKEN_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
