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

func CombinationBetPayoutResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponse) {

	secret := redisutility.GetSecret()

	var combinationResponseParams betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponseParams
	if err := json.Unmarshal(paramsData, &combinationResponseParams); err != nil {
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

	response := &betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponse{
		Method:    constants.BETGAMES_TRANSACTION_BET_COMBINATION_PAYOUT,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    combinationResponseParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfTransactionBetCombinationPayoutResponse(response, secret)

	responseChannel <- *response

}