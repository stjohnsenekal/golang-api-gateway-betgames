package constants

/* Here follows the list of Message Queue Topics

	in format:

	(domain).(api).(request/response).(specific information about the data)
*/

const BETGAMES_RESPONSE_REFRESH_TOKEN_TOPIC = "betgames.api.response.refresh.token"

const BETGAMES_RESPONSE_NEW_TOKEN_TOPIC = "betgames.api.response.new.token"

const BETGAMES_REQUEST_GET_ACCOUNT_DETAILS_TOPIC = "betgames.api.request.account.details"
const BETGAMES_RESPONSE_GET_ACCOUNT_DETAILS_TOPIC = "betgames.api.response.account.details"

const BETGAMES_REQUEST_GET_BALANCE_TOPIC = "betgames.api.request.account.balance"
const BETGAMES_RESPONSE_GET_BALANCE_TOPIC = "betgames.api.response.account.balance"

const BETGAMES_REQUEST_TRANSACTION_BET_PAYIN_TOPIC = "betgames.api.request.transaction.payin"
const BETGAMES_RESPONSE_TRANSACTION_BET_PAYIN_TOPIC = "betgames.api.response.transaction.payin"

const BETGAMES_REQUEST_SUBSCRIPTION_BET_PAYIN_TOPIC = "betgames.api.request.subscription.payin"
const BETGAMES_RESPONSE_SUBSCRIPTION_BET_PAYIN_TOPIC = "betgames.api.response.subscription.payin"

const BETGAMES_REQUEST_TRANSACTION_BET_PAYOUT_TOPIC = "betgames.api.request.transaction.payout"
const BETGAMES_RESPONSE_TRANSACTION_BET_PAYOUT_TOPIC = "betgames.api.response.transaction.payout"

const BETGAMES_REQUEST_COMBINATION_BET_PAYIN_TOPIC = "betgames.api.request.combination.payin"
const BETGAMES_RESPONSE_COMBINATION_BET_PAYIN_TOPIC = "betgames.api.response.combination.payin"

const BETGAMES_REQUEST_COMBINATION_BET_PAYOUT_TOPIC = "betgames.api.request.combination.payout"
const BETGAMES_RESPONSE_COMBINATION_BET_PAYOUT_TOPIC = "betgames.api.response.combination.payout"

const BETGAMES_RESPONSE_GENERAL_ERROR_TOPIC = "betgames.api.response.error"


