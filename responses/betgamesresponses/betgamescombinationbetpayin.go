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

func CombinationBetPayinResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse) {

	secret := redisutility.GetSecret()

	var combinationResponseParams betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponseParams
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

	response := &betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse{
		Method:    constants.BETGAMES_TRANSACTION_BET_COMBINATION_PAYIN,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    combinationResponseParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfTransactionBetCombinationPayinResponse(response, secret)

	responseChannel <- *response

}