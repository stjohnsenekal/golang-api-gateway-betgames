package betgameshandler

type ErrorCodes struct {
	ErrorCodes []ErrorCode `json:"ErrorCodes"`
}

type ErrorCode struct {
	Name   string `json:"name"`
	Code   int `json:"code"`
	Value  string    `json:"value"`
}
