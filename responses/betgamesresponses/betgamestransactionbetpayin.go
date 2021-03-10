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

func TransactionBetPayinResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestTransactionBetPayinResponse) {

	secret := redisutility.GetSecret()

	var transactionResponseParams betgamesmodel.BetGamesRequestTransactionBetPayinResponseParams
	if err := json.Unmarshal(paramsData, &transactionResponseParams); err != nil {
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

	response := &betgamesmodel.BetGamesRequestTransactionBetPayinResponse{
		Method:    constants.BETGAMES_TRANSACTION_BET_PAYIN,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    transactionResponseParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfTransactionBetPayinResponse(response, secret)

	responseChannel <- *response

}