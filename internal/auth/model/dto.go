package model

import (
	"bank_api/internal/pkg/common"
)

type SignUpRequest struct {
	//server.Params
	Email     string `json:"email,required"`
	Password  string `json:"password,required"`
	FirstName string `json:"first_name,required"`
	LastName  string `json:"last_name"`
}

type SignUpResponse struct {
	common.Response
}

type EmailValidateRequest struct {
	Code int32 `json:"code,required"`
}

type EmailValidateResponse struct {
	common.Response
	Valid bool `json:"valid"`
}

type SignInRequest struct {
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type SignInResponse struct {
	common.Response
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
