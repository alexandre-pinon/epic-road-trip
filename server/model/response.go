package model

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type AppResult struct {
	Data       interface{}
	Message    string
	StatusCode int
}

type AppError struct {
	Err error
	StatusCode int
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}