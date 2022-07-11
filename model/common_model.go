package model

type DateColumn struct {
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(message string, data interface{}) *Response {
	return &Response{
		Message: message,
		Data:    data,
	}
}
