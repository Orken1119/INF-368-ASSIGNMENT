package models 

type SuccessResponse struct {
	Result   interface{} `json:"result"`
}

type ErrorResponse struct {
	Result []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code     string     `json:"code"`
	Message  string     `json:"message"`
}

