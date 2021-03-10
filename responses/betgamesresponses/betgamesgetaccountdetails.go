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

func GetAccountDetailsResponse(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesAccountDetailsResponse) {

	secret := redisutility.GetSecret()

	var accountDetailsParams betgamesmodel.BetGamesAccountDetailsResponseParams
	if err := json.Unmarshal(paramsData, &accountDetailsParams); err != nil {
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

	response := &betgamesmodel.BetGamesAccountDetailsResponse{
		Method:    constants.BETGAMES_GET_ACCOUNT_DETAILS,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
		Params:    accountDetailsParams,
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfGetAccountDetailsResponse(response, secret)

	responseChannel <- *response

}