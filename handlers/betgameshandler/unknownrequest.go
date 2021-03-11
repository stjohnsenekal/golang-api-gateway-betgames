package betgameshandler

import (
	"api-gateway/crypto/betgamescrypto"
	"api-gateway/models/betgamesmodel"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func unknownRequest(context echo.Context) (err error)  {

	now := time.Now()
	seconds := now.Unix()

	response := &betgamesmodel.BetGamesBaseResponse{
		Method: "unknown",
		Token: "-",
		Succcess: "0",
		ErrorCode: "0",
		ErrorText: "unknown request",
		Time: strconv.FormatInt(seconds,10),
	}

	response.Signature = betgamescrypto.CalculateMd5ChecksumOfErrorResponse(response, configuration.Secret)

	return context.XML(http.StatusBadRequest, &response)
}
