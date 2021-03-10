package betgamesrequests

import (
	"api-gateway/constants"
	"api-gateway/messagequeue"
	"api-gateway/models/betgamesmodel"
	"api-gateway/redisutility"
	"fmt"
	"time"
)

func BetgamesRefreshTokenRequest(userid string, token string, timeout *betgamesmodel.BetGamesBaseResponse, tlen int) interface{} {

	err :=  redisutility.RefreshUserToken(token)
	if err != nil {
		fmt.Println(err)
	}

	message := messagequeue.Message{
		MessageType: constants.BETGAMES_REFRESH_TOKEN_MESSAGE_TYPE,
		UserId: userid,
		Data: token,
	}

	messageResponse := make(chan betgamesmodel.BetGamesBaseResponse)
	errorResponse := make(chan betgamesmodel.BetGamesBaseResponse)

	messagequeue.SetChannelBetGamesBaseResponse(messageResponse)
	messagequeue.SetChannelBetGamesErrorResponse(errorResponse)

	messagequeue.Broadcast(message, constants.BETGAMES_RESPONSE_REFRESH_TOKEN_TOPIC)

	select {
	case result := <- messageResponse:
		return result
	case errorResult := <- errorResponse:
		return errorResult
	case <- time.After(time.Duration(tlen) * time.Second):
		return timeout
	}

}
