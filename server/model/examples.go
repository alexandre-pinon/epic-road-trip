package model

type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"yoimiya.naganohara@gmail.com"`
	Password string `json:"password" binding:"required" example:"12345678"`
}
type LoginSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Login successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type LoginFailureCredentials struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"incorrect/missing email or password"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type RegisterRequest struct {
	Firstname string `json:"firstname" binding:"required,min=2,max=50" example:"yoimiya"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50" example:"naganohara"`
	Email     string `json:"email" binding:"required,email" example:"yoimiya.naganohara@gmail.com"`
	Password  string `json:"password" binding:"required,min=8,max=100" example:"12345678"`
	Phone     string `json:"phone" binding:"required,e164,len=12" example:"+33612345678"`
}
type RegisterSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"User created successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type RegisterFailureInvalid struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"invalid json request body"`
	Data      struct{}   `json:"data"`
	ValErrors []ValError `json:"valErrors"`
}

type LogoutSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Logout successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type LogoutFailure struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"cookie token is empty / Token is expired"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type RefreshSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Token refreshed successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type RefreshFailure struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"cookie token is empty / Token is expired"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
