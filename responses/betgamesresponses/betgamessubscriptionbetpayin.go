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

func SubscriptionBetPayinResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse) {

	secret := redisutility.GetSecret()

	var subscriptionResponseParams betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponseParams
	if err := json.Unmarshal(paramsData, &subscriptionResponseParams); err != nil {
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

	response := &betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse{
		Method:    constants.BETGAMES_TRANSACTION_BET_SUBSCRIPTION_PAYIN,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    subscriptionResponseParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfTransactionBetSubscriptionPayinResponse(response, secret)

	responseChannel <- *response

}