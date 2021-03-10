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

func GetBalanceResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestGetBalanceResponse) {

	secret := redisutility.GetSecret()

	var accountBalanceParams betgamesmodel.BetGamesRequestGetBalanceResponseParams
	if err := json.Unmarshal(paramsData, &accountBalanceParams); err != nil {
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

	response := &betgamesmodel.BetGamesRequestGetBalanceResponse{
		Method:    constants.BETGAMES_GET_BALANCE,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    accountBalanceParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfGetBalanceResponse(response, secret)

	responseChannel <- *response

}