package response

type ApiResponse struct {
	Status  string      `json:"Status"`
	Message string      `json:"message"`
	Data    interface{} `json:"Data"`
}

func NewApiResponse(status, message string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
