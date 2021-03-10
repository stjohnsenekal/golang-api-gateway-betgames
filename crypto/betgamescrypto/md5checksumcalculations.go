package betgamescrypto

import (
	"api-gateway/models/betgamesmodel"
	"bytes"
	"crypto/md5"
	"fmt"
)

/* base method for calculation of check sum*/
func createChecksum(b bytes.Buffer) string {

	//prima materia
	data := b.String()
	input := []byte(data)

	//md5 checksum calculation
	calculation := md5.Sum(input)
	checksum := fmt.Sprintf("%x", calculation)

	return checksum
}

/*
func CalculateMd5Checksum(request BetGamesBaseResponse) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("success")
	b.WriteString(request.Succcess)
	b.WriteString("error_code")
	b.WriteString(request.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(request.ErrorText)
	b.WriteString("time")
	b.WriteString(request.Time)

	if request.Params.UserId != "" {
		b.WriteString("user_id")
		b.WriteString(request.Params.UserId)
	}

	if request.Params.Username != "" {
		b.WriteString("username")
		b.WriteString(request.Params.Username)
	}

	if request.Params.Currency != "" {
		b.WriteString("currency")
		b.WriteString(request.Params.Currency)
	}

	if request.Params.Info != "" {
		b.WriteString("info")
		b.WriteString(request.Params.Info)
	}

	if request.Params.NewToken != "" {
		b.WriteString("new_token")
		b.WriteString(request.Params.NewToken)
	}

	b.WriteString(configuration.Secret)

	checksum := createChecksum(b)

	return checksum
} */

func CalculateMd5ChecksumOfErrorResponse(response *betgamesmodel.BetGamesBaseResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfPingRequest(request *betgamesmodel.BetGamesBaseRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("time")
	b.WriteString(request.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfPingResponse(response *betgamesmodel.BetGamesBaseResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfGetAccountDetailsRequest(request *betgamesmodel.BetGamesBaseRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("time")
	b.WriteString(request.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfGetAccountDetailsResponse(response *betgamesmodel.BetGamesAccountDetailsResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("user_id")
	b.WriteString(response.Params.UserId)
	b.WriteString("username")
	b.WriteString(response.Params.Username)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("info")
	b.WriteString(response.Params.Info)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfRefreshTokenRequest(request *betgamesmodel.BetGamesBaseRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("time")
	b.WriteString(request.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfRefreshTokenResponse(response *betgamesmodel.BetGamesBaseResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfRequestNewTokenRequest(request *betgamesmodel.BetGamesBaseRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("time")
	b.WriteString(request.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfRequestNewTokenResponse(response *betgamesmodel.BetGamesRequestNewTokenResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("new_token")
	b.WriteString(response.Params.NewToken)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfGetBalanceRequest(request *betgamesmodel.BetGamesBaseRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(request.Method)
	b.WriteString("token")
	b.WriteString(request.Token)
	b.WriteString("time")
	b.WriteString(request.Time)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfGetBalanceResponse(response *betgamesmodel.BetGamesRequestGetBalanceResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance")
	b.WriteString(response.Params.Balance)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetPayinRequest(response *betgamesmodel.BetGamesRequestTransactionBetPayinRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("amount")
	b.WriteString(response.Params.Amount)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("bet_id")
	b.WriteString(response.Params.BetId)
	b.WriteString("transaction_id")
	b.WriteString(response.Params.TransactionId)
	b.WriteString("retrying")
	b.WriteString(response.Params.Retrying)
	b.WriteString("bet")
	b.WriteString(response.Params.Bet)
	b.WriteString("odd")
	b.WriteString(response.Params.Odd)
	b.WriteString("bet_time")
	b.WriteString(response.Params.BetTime)
	b.WriteString("game")
	b.WriteString(response.Params.Game)
	b.WriteString("draw_code")
	b.WriteString(response.Params.DrawCode)
	b.WriteString("draw_time")
	b.WriteString(response.Params.DrawTime)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetPayinResponse(response *betgamesmodel.BetGamesRequestTransactionBetPayinResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance_after")
	b.WriteString(response.Params.BalanceAfter)
	b.WriteString("already_processed")
	b.WriteString(response.Params.AlreadyProcessed)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetPayoutRequest(response *betgamesmodel.BetGamesRequestTransactionBetPayoutRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("player_id")
	b.WriteString(response.Params.PlayerId)
	b.WriteString("amount")
	b.WriteString(response.Params.Amount)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("bet_id")
	b.WriteString(response.Params.BetId)
	b.WriteString("transaction_id")
	b.WriteString(response.Params.TransactionId)
	b.WriteString("retrying")
	b.WriteString(response.Params.Retrying)
	b.WriteString("bet_type")
	b.WriteString(response.Params.BetType)
	b.WriteString("game_id")
	b.WriteString(response.Params.GameId)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetPayoutResponse(response *betgamesmodel.BetGamesRequestTransactionBetPayoutResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance_after")
	b.WriteString(response.Params.BalanceAfter)
	b.WriteString("already_processed")
	b.WriteString(response.Params.AlreadyProcessed)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetSubscriptionPayinRequest(response *betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("amount")
	b.WriteString(response.Params.Amount)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("subscription_id")
	b.WriteString(response.Params.SubscriptionId)
	b.WriteString("subscription_time")
	b.WriteString(response.Params.SubscriptionTime)

	b.WriteString("odd")
	b.WriteString("name")
	b.WriteString(response.Params.Odd.Name)
	b.WriteString("value")
	b.WriteString(response.Params.Odd.Value)
	b.WriteString("translation")
	b.WriteString(response.Params.Odd.Translation)

	b.WriteString("is_mobile")
	b.WriteString(response.Params.IsMobile)

	b.WriteString("game")
	b.WriteString("id")
	b.WriteString(response.Params.Game.Id)
	b.WriteString("name")
	b.WriteString(response.Params.Game.Name)
	b.WriteString("translation")
	b.WriteString(response.Params.Game.Translation)

	b.WriteString("bet")

	for _, individualBet := range response.Params.Bet {
		b.WriteString("bet_id")
		b.WriteString(individualBet.BetId)
		b.WriteString("transaction_id")
		b.WriteString(individualBet.TransactionId)
		b.WriteString("amount")
		b.WriteString(individualBet.Amount)
		b.WriteString("draw")
		b.WriteString("code")
		b.WriteString(individualBet.Draw.Code)
		b.WriteString("time")
		b.WriteString(individualBet.Draw.Time)
	}

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetSubscriptionPayinResponse(response *betgamesmodel.BetGamesRequestTransactionBetSubscriptionPayinResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance_after")
	b.WriteString(response.Params.BalanceAfter)
	b.WriteString("already_processed")
	b.WriteString(response.Params.AlreadyProcessed)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetCombinationPayinRequest(response *betgamesmodel.BetGamesRequestTransactionBetCombinationPayinRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("amount")
	b.WriteString(response.Params.Amount)
	b.WriteString("odd_value")
	b.WriteString(response.Params.OddValue)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("combination_id")
	b.WriteString(response.Params.CombinationId)
	b.WriteString("combination_time")
	b.WriteString(response.Params.CombinationTime)

	b.WriteString("is_mobile")
	b.WriteString(response.Params.IsMobile)

	b.WriteString("bet")

	for _, individualBet := range response.Params.Bet {
		b.WriteString("bet_id")
		b.WriteString(individualBet.BetId)
		b.WriteString("transaction_id")
		b.WriteString(individualBet.TransactionId)

		b.WriteString("draw")
		b.WriteString("code")
		b.WriteString(individualBet.Draw.Code)
		b.WriteString("time")
		b.WriteString(individualBet.Draw.Time) /*right up to here */

		b.WriteString("game")
		b.WriteString("id")
		b.WriteString(individualBet.Game.Id)
		b.WriteString("name")
		b.WriteString(individualBet.Game.Name)
		b.WriteString("translation")
		b.WriteString(individualBet.Game.Translation)

		b.WriteString("odd")
		b.WriteString("name")
		b.WriteString(individualBet.Odd.Name)
		b.WriteString("value")
		b.WriteString(individualBet.Odd.Value)
		b.WriteString("translation")
		b.WriteString(individualBet.Odd.Translation)

	}

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetCombinationPayinResponse(response *betgamesmodel.BetGamesRequestTransactionBetCombinationPayinResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance_after")
	b.WriteString(response.Params.BalanceAfter)
	b.WriteString("already_processed")
	b.WriteString(response.Params.AlreadyProcessed)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetCombinationPayoutRequest(response *betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutRequest, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("player_id")
	b.WriteString(response.Params.PlayerId)
	b.WriteString("amount")
	b.WriteString(response.Params.Amount)
	b.WriteString("type")
	b.WriteString(response.Params.Type)
	b.WriteString("currency")
	b.WriteString(response.Params.Currency)
	b.WriteString("combination_id")
	b.WriteString(response.Params.CombinationId)

	b.WriteString("bet")

	for _, individualBet := range response.Params.Bet {
		b.WriteString("bet_id")
		b.WriteString(individualBet.BetId)
		b.WriteString("game_id")
		b.WriteString(individualBet.GameId)
		b.WriteString("transaction_id")
		b.WriteString(individualBet.TransactionId)
	}

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}

func CalculateMd5ChecksumOfTransactionBetCombinationPayoutResponse(response *betgamesmodel.BetGamesRequestTransactionBetCombinationPayoutResponse, secret string) (result string) {
	var b bytes.Buffer

	b.WriteString("method")
	b.WriteString(response.Method)
	b.WriteString("token")
	b.WriteString(response.Token)
	b.WriteString("success")
	b.WriteString(response.Succcess)
	b.WriteString("error_code")
	b.WriteString(response.ErrorCode)
	b.WriteString("error_text")
	b.WriteString(response.ErrorText)
	b.WriteString("time")
	b.WriteString(response.Time)

	b.WriteString("balance_after")
	b.WriteString(response.Params.BalanceAfter)
	b.WriteString("already_processed")
	b.WriteString(response.Params.AlreadyProcessed)

	b.WriteString(secret)

	checksum := createChecksum(b)

	return checksum
}
