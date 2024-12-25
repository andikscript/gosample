package model

type Response struct {
	ResponseCode string                 `json:"responseCode"`
	Description  string                 `json:"description"`
	Data         map[string]interface{} `json:"data"`
}

type ResponseList struct {
	ResponseCode string                   `json:"responseCode"`
	Description  string                   `json:"description"`
	Data         []map[string]interface{} `json:"data"`
}
