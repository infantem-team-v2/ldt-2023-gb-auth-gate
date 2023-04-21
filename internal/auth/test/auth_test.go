package test

import (
	_ "bank_api/internal/auth/interface"
	"bank_api/internal/auth/model"
	_ "bank_api/internal/auth/usecase"
	"bank_api/internal/pkg/common"
	"bank_api/internal/pkg/server"
	"bank_api/pkg/thttp"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	netHttp "net/http"
	"testing"
)

// ======================SIGN_UP========================//
func TestSignUp(t *testing.T) {
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
			Response: common.Response{
				Message:    "Created",
				StatusCode: 201,
			},
		},
		"test2": {
			Response: common.Response{
				Message:    "User already exists!",
				StatusCode: 409,
			},
		},
		"test3": {
			Response: common.Response{
				Message:    "Unauthorized Error",
				StatusCode: 401,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/sign/up", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].StatusCode, resp.StatusCode)
	}
}

// =========================SIGN_IN==========================//
func TestSignIn(t *testing.T) {
	tests := map[string]model.SignInRequest{
		"test1": {
			Email:    "test@test.com",
			Password: "Password1@",
		},
		"test2": {
			Email:    "aga@da.verno",
			Password: "shutka",
		},
		"test3": {
			Email:    "net.chto.eto",
			Password: "",
		}}

	expected := map[string]model.SignInResponse{
		"test1": {
			Response: common.Response{
				Message:    "Success",
				StatusCode: 200,
			},
		},
		"test2": {
			Response: common.Response{
				Message:    "Not found!",
				StatusCode: 404,
			},
		},
		"test3": {
			Response: common.Response{
				Message:    "Unauthorized Error",
				StatusCode: 401,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/sign/in", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].StatusCode, resp.StatusCode)
	}
}

// ==========================VALIDATE_EMAIL=========================//
func TestValidateEmail(t *testing.T) {
	tests := map[string]model.EmailValidateRequest{
		"test1": {
			Code: 123456,
		},
		"test2": {
			Code: 123,
		},
		"test3": {
			Code: 232323232,
		}}

	expected := map[string]model.EmailValidateResponse{
		"test1": {
			Response: common.Response{
				Message:    "Accepted",
				StatusCode: 202,
			},
		},
		"test2": {
			Response: common.Response{
				Message:    "Request validation failed",
				StatusCode: 400,
			},
		},
		"test3": {
			Response: common.Response{
				Message:    "Request validation failed",
				StatusCode: 400,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/email/validate", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].StatusCode, resp.StatusCode)
	}
}
