package model

import "time"

type AppResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ValErrors []ValError  `json:"valErrors"`
}

type AppResult struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type LoginResponseData struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
