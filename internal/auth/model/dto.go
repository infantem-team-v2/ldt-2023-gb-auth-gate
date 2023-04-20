package model

type StandardResponse struct {
	Message    string `json:"message"`
	StatusCode uint8  `json:"status_code"`
}

type SignUpRequest struct {
	//server.CommonParams
	Email     string `json:"email,required"`
	Password  string `json:"password,required"`
	FirstName string `json:"first_name,required"`
	LastName  string `json:"last_name"`
}

type SignUpResponse struct {
	StandardResponse
}

type EmailValidateRequest struct {
	Code int32 `json:"code"`
}

type EmailValidateResponse struct {
	StandardResponse
	Valid bool `json:"valid"`
}
