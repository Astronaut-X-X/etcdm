package dto

type PageRequest struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Total   int    `json:"total"`
	Keyword string `json:"keyword"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessRespPage(data any, total int64) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: "success",
		Data:    map[string]any{"data": data, "total": total},
	}
}

func SuccessRespData(data any) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: "success",
		Data:    data,
	}
}

func FailResp(msg string) *BaseResponse {
	return &BaseResponse{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}

func SuccessResp() *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: "success",
		Data:    nil,
	}
}
