package exceptions

type AppError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
