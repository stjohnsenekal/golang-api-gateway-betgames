package betgamesmodel

import "encoding/xml"

/* This is the most basic betgames message structure. It is used for
	(1) the most simple requests and (2) for pre-parsing the data before the message type is discovered.
	It is sufficient for all security and validity checks before routing. */
type BetGamesBaseRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the base response message for all responses sans parameters */
type BetGamesBaseResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the general error response data from our backend only. */
type ErrorResponseParams struct {
	MethodName				string `json:"methodName,omitempty" xml:"method_name,omitempty"`
	ErrorCode				string `json:"errorCode,omitempty" xml:"error_code,omitempty"`
	ErrorText				string `json:"errorText,omitempty" xml:"error_text,omitempty"`
}

/* This is the get_account_details response data */
type BetGamesAccountDetailsResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	UserId		string `json:"userId,omitempty" xml:"user_id,omitempty"`	//altered json to fit backend response
	Username	string `json:"username,omitempty" xml:"username,omitempty"`
	Currency	string `json:"currency,omitempty" xml:"currency,omitempty"`
	Info		string `json:"info,omitempty" xml:"info,omitempty"`
}

type BetGamesAccountDetailsResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesAccountDetailsResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the request_new_token response data */
type BetGamesRequestNewTokenResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	NewToken		string `json:"new_token,omitempty" xml:"new_token,omitempty"`
}

type BetGamesRequestNewTokenResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestNewTokenResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the request_new_token response data */
type BetGamesRequestGetBalanceResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	Balance		string `json:"balance,omitempty" xml:"balance,omitempty"`
}

type BetGamesRequestGetBalanceResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestGetBalanceResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_payin request data */
type BetGamesRequestTransactionBetPayinRequestParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	Amount			string `json:"amount,omitempty" xml:"amount,omitempty"`
	Currency		string `json:"currency,omitempty" xml:"currency,omitempty"`
	BetId			string `json:"bet_id,omitempty" xml:"bet_id,omitempty"`
	TransactionId	string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	Retrying		string `json:"retrying,omitempty" xml:"retrying,omitempty"`
	Bet				string `json:"bet,omitempty" xml:"bet,omitempty"`
	Odd				string `json:"odd,omitempty" xml:"odd,omitempty"`
	BetTime			string `json:"bet_time,omitempty" xml:"bet_time,omitempty"`
	Game			string `json:"game,omitempty" xml:"game,omitempty"`
	DrawCode		string `json:"draw_code,omitempty" xml:"draw_code,omitempty"`
	DrawTime		string `json:"draw_time,omitempty" xml:"draw_time,omitempty"`
}

type BetGamesRequestTransactionBetPayinRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetPayinRequestParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_payin response data */
type BetGamesRequestTransactionBetPayinResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	BalanceAfter			string `json:"balanceAfter,omitempty" xml:"balance_after,omitempty"`
	AlreadyProcessed		string `json:"alreadyProcessed,omitempty" xml:"already_processed,omitempty"`
}

type BetGamesRequestTransactionBetPayinResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetPayinResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

type OddsParams struct {
	XMLName xml.Name `json:"-" xml:"odd"`
	Name				string `json:"name,omitempty" xml:"name,omitempty"`
	Value				string `json:"value,omitempty" xml:"value,omitempty"`
	Translation			string `json:"translation,omitempty" xml:"translation,omitempty"`
}

type GamesParams struct {
	XMLName xml.Name `json:"-" xml:"game"`
	Id					string `json:"id,omitempty" xml:"id,omitempty"`
	Name				string `json:"name,omitempty" xml:"name,omitempty"`
	Translation			string `json:"translation,omitempty" xml:"translation,omitempty"`
}

type DrawParams struct {
	XMLName xml.Name `json:"-" xml:"draw"`
	Code				string `json:"code,omitempty" xml:"code,omitempty"`
	Time				string `json:"time,omitempty" xml:"time,omitempty"`
}

type BetsParams struct {
	XMLName xml.Name `json:"-" xml:"bet"`
	BetId				string `json:"bet_id,omitempty" xml:"bet_id,omitempty"`
	TransactionId		string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	Amount				string `json:"amount,omitempty" xml:"amount,omitempty"`
	Draw				DrawParams `json:"draw,omitempty" xml:"draw,omitempty"`
}

/* This is the transaction_bet_subscription_payin request data */
type BetGamesRequestTransactionBetSubscriptionPayinRequestParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	Amount				string `json:"amount,omitempty" xml:"amount,omitempty"`
	Currency			string `json:"currency,omitempty" xml:"currency,omitempty"`
	SubscriptionId		string `json:"subscription_id,omitempty" xml:"subscription_id,omitempty"`
	SubscriptionTime	string `json:"subscription_time,omitempty" xml:"subscription_time,omitempty"`
	Odd 				OddsParams `json:"odd,omitempty" xml:"odd,omitempty"`
	IsMobile			string `json:"is_mobile,omitempty" xml:"is_mobile,omitempty"`
	Game				GamesParams `json:"game,omitempty" xml:"game,omitempty"`
	Bet					[]BetsParams `json:"bet,omitempty" xml:"bet,omitempty"`

}

type BetGamesRequestTransactionBetSubscriptionPayinRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetSubscriptionPayinRequestParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_subscription_payin response data */
type BetGamesRequestTransactionBetSubscriptionPayinResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	BalanceAfter			string `json:"balanceAfter,omitempty" xml:"balance_after,omitempty"`
	AlreadyProcessed		string `json:"alreadyProcessed,omitempty" xml:"already_processed,omitempty"`
}

type BetGamesRequestTransactionBetSubscriptionPayinResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetSubscriptionPayinResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_payout request data */
type BetGamesRequestTransactionBetPayoutRequestParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	PlayerId			string `json:"player_id,omitempty" xml:"player_id,omitempty"`
	Amount				string `json:"amount,omitempty" xml:"amount,omitempty"`
	Currency			string `json:"currency,omitempty" xml:"currency,omitempty"`
	BetId				string `json:"bet_id,omitempty" xml:"bet_id,omitempty"`
	TransactionId 		string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	Retrying			string `json:"retrying,omitempty" xml:"retrying,omitempty"`
	BetType				string `json:"bet_type,omitempty" xml:"bet_type,omitempty"`
	GameId				string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

type BetGamesRequestTransactionBetPayoutRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetPayoutRequestParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_payout response data */
type BetGamesRequestTransactionBetPayoutResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	BalanceAfter			string `json:"balanceAfter,omitempty" xml:"balance_after,omitempty"`
	AlreadyProcessed		string `json:"alreadyProcessed,omitempty" xml:"already_processed,omitempty"`
}

type BetGamesRequestTransactionBetPayoutResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetPayoutResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

type BetCombinationParams struct {
	XMLName xml.Name `json:"-" xml:"bet"`
	BetId				string `json:"bet_id,omitempty" xml:"bet_id,omitempty"`
	TransactionId		string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	Draw				DrawParams `json:"draw,omitempty" xml:"draw,omitempty"`
	Game				GamesParams `json:"game,omitempty" xml:"game,omitempty"`
	Odd					OddsParams `json:"odd,omitempty" xml:"odd,omitempty"`
}

/* This is the transaction_bet_combination_payin request data */
type BetGamesRequestTransactionBetCombinationPayinRequestParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	Amount				string `json:"amount,omitempty" xml:"amount,omitempty"`
	OddValue			string `json:"odd_value,omitempty" xml:"odd_value,omitempty"`
	Currency			string `json:"currency,omitempty" xml:"currency,omitempty"`
	CombinationId		string `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
	CombinationTime		string `json:"combination_time,omitempty" xml:"combination_time,omitempty"`
	IsMobile			string `json:"is_mobile,omitempty" xml:"is_mobile,omitempty"`
	Bet 				[]BetCombinationParams `json:"bet,omitempty" xml:"bet,omitempty"`

}

type BetGamesRequestTransactionBetCombinationPayinRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetCombinationPayinRequestParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_combination_payin response data */
type BetGamesRequestTransactionBetCombinationPayinResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	BalanceAfter			string `json:"balanceAfter,omitempty" xml:"balance_after,omitempty"`
	AlreadyProcessed		string `json:"alreadyProcessed,omitempty" xml:"already_processed,omitempty"`
}

type BetGamesRequestTransactionBetCombinationPayinResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetCombinationPayinResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

type PayoutBets struct {
	XMLName xml.Name `json:"-" xml:"bet"`
	BetId				string `json:"bet_id,omitempty" xml:"bet_id,omitempty"`
	GameId				string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	TransactionId		string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
}

/* This is the transaction_bet_combination_payout request data */
type BetGamesRequestTransactionBetCombinationPayoutRequestParams struct { //here
	XMLName xml.Name `json:"-" xml:"params"`
	PlayerId			string `json:"player_id,omitempty" xml:"player_id,omitempty"`
	Amount				string `json:"amount,omitempty" xml:"amount,omitempty"`
	Type				string `json:"type,omitempty" xml:"type,omitempty"`
	Currency			string `json:"currency,omitempty" xml:"currency,omitempty"`
	CombinationId		string `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
	Bet 				[]PayoutBets `json:"bet,omitempty" xml:"bet,omitempty"`
}

type BetGamesRequestTransactionBetCombinationPayoutRequest struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetCombinationPayoutRequestParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the transaction_bet_combination_payout response data */
type BetGamesRequestTransactionBetCombinationPayoutResponseParams struct {
	XMLName xml.Name `json:"-" xml:"params"`
	BalanceAfter			string `json:"balanceAfter,omitempty" xml:"balance_after,omitempty"`
	AlreadyProcessed		string `json:"alreadyProcessed,omitempty" xml:"already_processed,omitempty"`
}

type BetGamesRequestTransactionBetCombinationPayoutResponse struct {
	XMLName xml.Name `json:"-" xml:"root"`
	Method		string `json:"method" xml:"method"`
	Token		string `json:"token" xml:"token"`
	Succcess 	string `json:"success" xml:"success"`
	ErrorCode 	string `json:"error_code" xml:"error_code"`
	ErrorText 	string `json:"error_text" xml:"error_text"`
	Time 		string `json:"time" xml:"time"`
	Params 		BetGamesRequestTransactionBetCombinationPayoutResponseParams `json:"params" xml:"params"`
	Signature 	string `json:"signature,omitempty" xml:"signature,omitempty"`
}

/* This is the token message structure */
type Token struct {
	Token  		string `json:"token" xml:"token"`
}

/* This is the request token structure */
type RequestToken struct {
	UserId  	int `json:"userid" xml:"userid"`
}

