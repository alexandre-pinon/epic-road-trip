package model

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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