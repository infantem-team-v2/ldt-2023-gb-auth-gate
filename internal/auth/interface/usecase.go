package authInterface

import "bank_api/internal/auth/model"

type UseCase interface {
	SignUp(params *model.SignUpRequest) (*model.SignUpResponse, error)
	SignIn(params *model.SignInRequest) (*model.SignInResponse, error)
	ValidateEmail(params *model.EmailValidateRequest) (*model.EmailValidateResponse, error)
}
