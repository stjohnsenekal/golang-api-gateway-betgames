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

func RequestNewToken(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesRequestNewTokenResponse) {

	secret := redisutility.GetSecret()

	var tokenResponseParams betgamesmodel.BetGamesRequestNewTokenResponseParams
	if err := json.Unmarshal(paramsData, &tokenResponseParams); err != nil {
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

	response := &betgamesmodel.BetGamesRequestNewTokenResponse{
		Method:    constants.BETGAMES_REQUEST_NEW_TOKEN,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    tokenResponseParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfRequestNewTokenResponse(response, secret)

	responseChannel <- *response

}