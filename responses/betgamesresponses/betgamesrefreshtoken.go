package betgamesresponses

import (
	"api-gateway/constants"
	"api-gateway/crypto/betgamescrypto"
	"api-gateway/models/betgamesmodel"
	"api-gateway/redisutility"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func RefreshToken(userid string, paramsData json.RawMessage, responseChannel chan betgamesmodel.BetGamesBaseResponse) {

	secret := redisutility.GetSecret()

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
		Method:    constants.BETGAMES_REFRESH_TOKEN,
		Token:     token,
		Succcess:  "1",
		ErrorCode: "0",
		ErrorText: "",
		Time:      strconv.FormatInt(seconds,10),
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfRefreshTokenResponse(response, secret)

	responseChannel <- *response

}