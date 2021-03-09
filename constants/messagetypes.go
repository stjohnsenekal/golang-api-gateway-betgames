package constants

/* the list of message types across the message queues */

/* betgames message types*/

const BETGAMES_REFRESH_TOKEN_MESSAGE_TYPE = "betgames_refresh_token"
const BETGAMES_REFRESH_TOKEN = "refresh_token"

const BETGAMES_REQUEST_NEW_TOKEN_MESSAGE_TYPE = "betgames_new_token"
const BETGAMES_REQUEST_NEW_TOKEN = "request_new_token"

const BETGAMES_ACCOUNT_REQUEST_MESSAGE_TYPE = "betgames_get_account_details_request"
const BETGAMES_ACCOUNT_RESPONSE_MESSAGE_TYPE = "betgames_get_account_details_response"
const BETGAMES_GET_ACCOUNT_DETAILS = "get_account_details"

const BETGAMES_GET_BALANCE_REQUEST_MESSAGE_TYPE = "betgames_get_balance_request"
const BETGAMES_GET_BALANCE_RESPONSE_MESSAGE_TYPE = "betgames_get_balance_response"
const BETGAMES_GET_BALANCE = "get_balance"

const BETGAMES_BET_PAYIN_REQUEST_MESSAGE_TYPE = "betgames_bet_payin_request"
const BETGAMES_BET_PAYIN_RESPONSE_MESSAGE_TYPE = "betgames_bet_payin_response"
const BETGAMES_TRANSACTION_BET_PAYIN = "transaction_bet_payin"

const BETGAMES_BET_SUBSCRIPTION_PAYIN_REQUEST_MESSAGE_TYPE = "betgames_bet_subscription_payin_request"
const BETGAMES_BET_SUBSCRIPTION_PAYIN_RESPONSE_MESSAGE_TYPE = "betgames_bet_subscription_payin_response"
const BETGAMES_TRANSACTION_BET_SUBSCRIPTION_PAYIN = "transaction_bet_subscription_payin"

const BETGAMES_BET_PAYOUT_REQUEST_MESSAGE_TYPE = "betgames_bet_payout_request"
const BETGAMES_BET_PAYOUT_RESPONSE_MESSAGE_TYPE = "betgames_bet_payout_response"
const BETGAMES_TRANSACTION_BET_PAYOUT = "transaction_bet_payout"

const BETGAMES_BET_COMBINATION_PAYIN_REQUEST_MESSAGE_TYPE = "betgames_bet_combination_payin_request"
const BETGAMES_BET_COMBINATION_PAYIN_RESPONSE_MESSAGE_TYPE = "betgames_bet_combination_payin_response"
const BETGAMES_TRANSACTION_BET_COMBINATION_PAYIN = "transaction_bet_combination_payin"

const BETGAMES_BET_COMBINATION_PAYOUT_REQUEST_MESSAGE_TYPE = "betgames_bet_combination_payout_request"
const BETGAMES_BET_COMBINATION_PAYOUT_RESPONSE_MESSAGE_TYPE = "betgames_bet_combination_payout_response"
const BETGAMES_TRANSACTION_BET_COMBINATION_PAYOUT = "transaction_bet_combination_payout"

const BETGAMES_ERROR_RESPONSE_MESSAGE_TYPE = "betgames_error_response"

const BETGAMES_QUEUE_FOR_SUBSCRIBER_REFRESH_TOKEN = "betgames-api-response-refresh-token"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_NEW_TOKEN = "betgames-api-response-new-token"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_ACCOUNT_DETAILS = "betgames-api-response-account-details"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_GET_BALANCE = "betgames-api-response-balance"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_TRANSACTION_PAYIN = "betgames-api-response-transaction-payin"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_SUBSCRIPTION_PAYIN = "betgames-api-response-subscription-payin"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_TRANSACTION_PAYOUT = "betgames-api-response-transaction-payout"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_COMBINATION_PAYIN = "betgames-api-response-combination-payin"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_COMBINATION_PAYOUT = "betgames-api-response-combination-payout"
const BETGAMES_QUEUE_FOR_SUBSCRIBER_GENERAL_ERROR = "betgames-api-response-error"