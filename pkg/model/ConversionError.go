package model

// ConversionError provides the error and error message
type ConversionError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
