// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Docs developer",
            "url": "https://t.me/sirop_work",
            "email": "bwg.it317@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Accepts token from vendor which we process and returning pair of tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Sign in or sign up via Apple or Google",
                "parameters": [
                    {
                        "enum": [
                            "apple",
                            "google"
                        ],
                        "type": "string",
                        "description": "Vendor which is providing authorization",
                        "name": "vendor",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/server.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.CommonResponse": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "bank.onecrypto.pro",
	BasePath:         "",
	Schemes:          []string{"https"},
	Title:            "CryptoBank API",
	Description:      "TimeStamp",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
