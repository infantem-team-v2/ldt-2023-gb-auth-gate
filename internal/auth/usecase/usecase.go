package usecase

import (
	"gb-auth-gate/internal/auth/model"
	"gb-auth-gate/internal/pkg/common"
	tlogger "gb-auth-gate/pkg/tlogger"
	"github.com/sarulabs/di"
	"math/rand"
	"time"
)

type AuthUS struct {
	logger tlogger.ILogger
}

func BuildAuthUsecase(ctn di.Container) (interface{}, error) {
	return &AuthUS{
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

func (as *AuthUS) SignUp(params *model.SignUpRequest) (*model.SignUpResponse, error) {
	return &model.SignUpResponse{
		Response: common.Response{
			Message:      "CREATED",
			InternalCode: 201,
		},
	}, nil
}

func (as *AuthUS) ValidateEmail(params *model.EmailValidateRequest) (*model.EmailValidateResponse, error) {
	rand.Seed(time.Now().UnixNano())
	valid := rand.Intn(2) == 1
	return &model.EmailValidateResponse{
		Response: common.Response{
			Message:      "SUCCESS",
			InternalCode: 200,
		},
		Valid: valid,
	}, nil
}

func (as *AuthUS) SignIn(params *model.SignInRequest) (*model.SignInResponse, error) {
	return &model.SignInResponse{
		Response: common.Response{
			Message:      "SUCCESS",
			InternalCode: 200,
		},
		AccessToken: "access_token",
	}, nil
}
