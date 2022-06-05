package model

type InternalServerError struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type Unauthorized struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"cookie token is empty / Token is expired"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
type InvalidJsonBody struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"invalid json request body"`
	Data      struct{}   `json:"data"`
	ValErrors []ValError `json:"valErrors"`
}
type InvalidID struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"invalid id"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

//* AUTH EXAMPLES *//
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

type LogoutSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Logout successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type RefreshSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Token refreshed successfully"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

//* USER EXAMPLES *//
type UserNotFound struct {
	Success   bool       `json:"success" example:"false"`
	Message   string     `json:"message" example:"user not found"`
	Data      struct{}   `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type GetAllUserSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"Users retrieved successfully"`
	Data      []User     `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type GetUserByIDSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"User <id> retrieved successfully"`
	Data      User       `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type UpdateUserRequest struct {
	Firstname string `json:"firstname" binding:"required,min=2,max=50" example:"yoimiya"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50" example:"naganohara"`
	Email     string `json:"email" binding:"required,email" example:"yoimiya.naganohara@gmail.com"`
	Phone     string `json:"phone" binding:"required,e164,len=12" example:"+33612345678"`
}
type UpdateUserSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"User updated successfully"`
	Data      []User     `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}

type DeleteUserSuccess struct {
	Success   bool       `json:"success" example:"true"`
	Message   string     `json:"message" example:"User deleted successfully"`
	Data      []User     `json:"data"`
	ValErrors []struct{} `json:"valErrors"`
}
