package test

import (
	"bank_api/internal/auth/delivery/http"
	authInterface "bank_api/internal/auth/interface"
	"bank_api/internal/auth/model"
	"bank_api/internal/pkg/dependency"
	"bank_api/pkg/thttp"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	netHttp "net/http"
	"testing"
)

//=======================AUTH=======================//
func TestSignUp(t *testing.T) {
	tdc := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()
	tests := map[string]model.SignUpRequest{
		"test1": {
			Email:     "test@test.com",
			FirstName: "Test",
			LastName:  "Testov",
			Password:  "password",
		},
		"test2": {
			Email:     "aga@da.verno",
			FirstName: "Lovi",
			LastName:  "Oshibku",
			Password:  "shutka",
		},
		"test3": {
			Email:     "net.chto.eto",
			FirstName: "123123",
			LastName:  "ls23e",
			Password:  "",
		}}

	expected := map[string]model.SignUpResponse{
		"test1": {
			StandardResponse: model.StandardResponse{
				Message:    "SUCCESS",
				StatusCode: 200,
			},
		},
		"test2": {
			StandardResponse: model.StandardResponse{
				Message:    "User already exists!",
				StatusCode: 409,
			},
		},
		"test3": {
			StandardResponse: model.StandardResponse{
				Message:    "Unauthorized Error",
				StatusCode: 401,
			},
		},
	}
	au := tdc.ContainerInstance().Get("authUC").(authInterface.UseCase)
	ah := http.AuthHandler{
		AuthUC: au,
	}
	ah.SetRoutes()
	for i, test := range tests {
		fmt.Println(i)
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/sign/up", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := ah.App.Test(req, -1)
		assert.Equal(t, expected[i].StatusCode, resp.StatusCode)
	}
}
