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
            "url": "https://t.me/KlenoviySirop",
            "email": "KlenoviySir@yandex.ru"
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
                "summary": "Sign in or sign up via external vendor",
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
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/check": {
            "get": {
                "description": "Validates that session is authorized",
                "tags": [
                    "Authorization"
                ],
                "summary": "Health and auth check",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/email/validate": {
            "post": {
                "description": "Validating user's email with take message on email and writing code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Validating user's email",
                "parameters": [
                    {
                        "description": "Data for validation by email from app",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EmailValidateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.EmailValidateResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/password/reset": {
            "put": {
                "description": "Resetting password by getting validated email for password change",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization",
                    "Password"
                ],
                "summary": "Resetting password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session key to identify that this is current session of password change",
                        "name": "t-session-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "New password pair with confirmation to update credentials",
                        "name": "new_pswds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/password/reset/prepare": {
            "patch": {
                "description": "Creates session for password reset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization",
                    "Password"
                ],
                "summary": "Make reset password session",
                "parameters": [
                    {
                        "description": "Session key for backend's session validation",
                        "name": "new_pswds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PrepareResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PrepareResetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/password/reset/validate": {
            "patch": {
                "description": "Validate reset password session by code that user gets on its email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization",
                    "Password"
                ],
                "summary": "Validate reset password session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session key to identify that this is current session of password change",
                        "name": "t-session-key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Code from email to validate request",
                        "name": "new_pswds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ValidateResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ValidateResetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/sign/in": {
            "post": {
                "description": "Authorization and get access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "Authorization data from user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/sign/out": {
            "delete": {
                "description": "Delete tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Sign out",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SignOutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/auth/sign/up": {
            "post": {
                "description": "Sign up with data which was in our task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Sign up with base data",
                "parameters": [
                    {
                        "description": "Authorization data from user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/calc": {
            "post": {
                "description": "Calculations with authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculator"
                ],
                "summary": "Improved calculation w/ auth",
                "parameters": [
                    {
                        "description": "Basic parameters for base calculator w/o auth",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BaseCalculateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BaseCalculateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/calc/base": {
            "post": {
                "description": "Base calculation without authorization for landing page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculator"
                ],
                "summary": "Base calculation",
                "parameters": [
                    {
                        "description": "Basic parameters for base calculator w/o auth",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BaseCalculateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BaseCalculateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.BaseCalculateRequest": {
            "type": "object"
        },
        "model.BaseCalculateResponse": {
            "type": "object",
            "properties": {
                "input_data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BasicCalculationFieldLogic"
                    }
                },
                "output_data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BasicCategoryCalculationLogic"
                    }
                }
            }
        },
        "model.BasicCalculationFieldLogic": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.BasicCategoryCalculationLogic": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BasicCalculationFieldLogic"
                    }
                }
            }
        },
        "model.BusinessDataLogic": {
            "type": "object",
            "required": [
                "inn"
            ],
            "properties": {
                "economic_activity": {
                    "type": "string",
                    "example": "Производство"
                },
                "inn": {
                    "type": "string",
                    "example": "7707083893"
                },
                "name": {
                    "type": "string",
                    "example": "ООО ИНФАНТЕМ"
                },
                "website": {
                    "type": "string",
                    "example": "infantem.tech"
                }
            }
        },
        "model.EmailValidateRequest": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                }
            }
        },
        "model.EmailValidateResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "model.PersonalDataLogic": {
            "type": "object",
            "required": [
                "full_name"
            ],
            "properties": {
                "full_name": {
                    "type": "string",
                    "example": "Иванов Иван Иванович"
                },
                "geographic": {
                    "type": "object",
                    "properties": {
                        "city": {
                            "type": "string",
                            "example": "Москва"
                        },
                        "country": {
                            "type": "string",
                            "example": "Российская Федерация"
                        }
                    }
                },
                "position": {
                    "type": "string",
                    "example": "Старший менеджер по инвестициям"
                }
            }
        },
        "model.PrepareResetPasswordRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "model.PrepareResetPasswordResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                },
                "session_key": {
                    "type": "string"
                }
            }
        },
        "model.RegistrationDataLogic": {
            "type": "object",
            "required": [
                "email",
                "password",
                "repeated_password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@mail.ru"
                },
                "password": {
                    "type": "string",
                    "example": "1234qwerty!"
                },
                "repeated_password": {
                    "type": "string",
                    "example": "1234qwerty!"
                }
            }
        },
        "model.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "password",
                "repeated_password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "repeated_password": {
                    "type": "string"
                }
            }
        },
        "model.ResetPasswordResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.SignInResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.SignOutResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.SignUpRequest": {
            "type": "object",
            "properties": {
                "auth_data": {
                    "$ref": "#/definitions/model.RegistrationDataLogic"
                },
                "business_data": {
                    "$ref": "#/definitions/model.BusinessDataLogic"
                },
                "personal_data": {
                    "$ref": "#/definitions/model.PersonalDataLogic"
                }
            }
        },
        "model.SignUpResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        },
        "model.ValidateResetPasswordRequest": {
            "type": "object",
            "required": [
                "validation_code"
            ],
            "properties": {
                "validation_code": {
                    "type": "string"
                }
            }
        },
        "model.ValidateResetPasswordResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "internal_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "gate.gb.ldt2023.infantem.tech",
	BasePath:         "",
	Schemes:          []string{"https"},
	Title:            "Core backend app for Leaders of Digital Transformation",
	Description:      "JWT token in authorization bearer",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
