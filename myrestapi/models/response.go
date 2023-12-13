package models

type DataResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}
