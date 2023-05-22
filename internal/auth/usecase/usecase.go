package usecase

import (
	authInterface "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/auth/model"
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/pkg/terrors"
	tlogger "gb-auth-gate/pkg/tlogger"
	"gb-auth-gate/pkg/tsecure"
	"github.com/sarulabs/di"
	"math/rand"
	"time"
)

type AuthUC struct {
	logger   tlogger.ILogger
	authRepo authInterface.RelationalRepository
}

func BuildAuthUsecase(ctn di.Container) (interface{}, error) {
	return &AuthUC{
		logger:   ctn.Get("logger").(tlogger.ILogger),
		authRepo: ctn.Get("authRepo").(authInterface.RelationalRepository),
	}, nil
}

func (as *AuthUC) SignUp(params *model.SignUpRequest) (*model.SignUpResponse, error) {
	return &model.SignUpResponse{
		Response: common.Response{
			Message:      "CREATED",
			InternalCode: 201,
		},
	}, nil
}

func (as *AuthUC) ValidateEmail(params *model.EmailValidateRequest) (*model.EmailValidateResponse, error) {
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

func (as *AuthUC) SignIn(params *model.SignInRequest) (*model.SignInResponse, error) {
	return &model.SignInResponse{
		Response: common.Response{
			Message:      "SUCCESS",
			InternalCode: 200,
		},
		AccessToken: "access_token",
	}, nil
}

func (as *AuthUC) ValidateService(params *model.AuthHeadersLogic) (bool, error) {
	service, err := as.authRepo.FindServiceByName(params.PublicKey)
	if err != nil {
		return false, err
	}
	signature := tsecure.CalcSignature(
		service.PrivateKey,
		string(params.Body),
		tsecure.SHA512,
	)
	if params.Signature == signature {
		return true, nil
	}

	return false, terrors.Raise(nil, 100003)
}
