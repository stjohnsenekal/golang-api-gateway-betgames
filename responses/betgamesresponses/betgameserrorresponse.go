package betgamesresponses

import (
	"api-gateway/constants"
	"api-gateway/crypto/betgamescrypto"
	"api-gateway/models/betgamesmodel"
	"api-gateway/redisutility"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func ErrorResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesBaseResponse) {

	secret := redisutility.GetSecret()

	var errorResponseParams betgamesmodel.ErrorResponseParams
	if err := json.Unmarshal(paramsData, &errorResponseParams); err != nil {
		log.Fatal(err)
	}

	userId, err := strconv.Atoi(userid)
	if err != nil {
		fmt.Println(err)
	}

	//TODO: edge case - what if token expires in ms between backend and api
	token, err :=  redisutility.GetUserTokenForId(userId)
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now()
	seconds := now.Unix()

	response := &betgamesmodel.BetGamesBaseResponse{
		Method:    getMethodFromMessageType(errorResponseParams.MethodName),
		Token:     token,
		Succcess:  "0",
		ErrorCode: errorResponseParams.ErrorCode,
		ErrorText: errorResponseParams.ErrorText,
		Time:      strconv.FormatInt(seconds,10),
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfErrorResponse(response, secret)

	responseChannel <- *response

}

func getMethodFromMessageType(methodName string) string {
	switch methodName {

	case constants.BETGAMES_GET_BALANCE_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_GET_BALANCE
	case constants.BETGAMES_ACCOUNT_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_GET_ACCOUNT_DETAILS
	case constants.BETGAMES_BET_PAYIN_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_TRANSACTION_BET_PAYIN
	case constants.BETGAMES_BET_SUBSCRIPTION_PAYIN_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_TRANSACTION_BET_SUBSCRIPTION_PAYIN
	case constants.BETGAMES_BET_PAYOUT_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_TRANSACTION_BET_PAYOUT
	case constants.BETGAMES_BET_COMBINATION_PAYIN_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_TRANSACTION_BET_COMBINATION_PAYIN
	case constants.BETGAMES_BET_COMBINATION_PAYOUT_REQUEST_MESSAGE_TYPE:
		return constants.BETGAMES_TRANSACTION_BET_COMBINATION_PAYOUT

		//where do all the token errors go

	default:
		return "Unknown message"
	}
}