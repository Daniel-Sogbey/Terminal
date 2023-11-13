package models

type DataResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
