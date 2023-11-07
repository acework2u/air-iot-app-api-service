// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/confirm": {
            "post": {
                "description": "New user confirm a sign up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User confirm is sign up",
                "parameters": [
                    {
                        "description": "New User confirm information",
                        "name": "SignUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UserConfirm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/confirm-password": {
            "post": {
                "description": "refresh new user token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh user token",
                "parameters": [
                    {
                        "description": "response confirm new password",
                        "name": "ConfirmNewPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.UserConfirmNewPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "post": {
                "description": "refresh new user token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh user token",
                "parameters": [
                    {
                        "description": "response confirm code",
                        "name": "resendConfirmCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ResendConfirmCode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "refresh new user token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh user token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/resend-confirm-code": {
            "post": {
                "description": "retern Resend confirmation code for a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Resend confirm code for a new user",
                "parameters": [
                    {
                        "description": "Resend confirm code",
                        "name": "ResendConfirmCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ResendConfirmCode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "Authenticates a user and provides authorize API Calls",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User signin",
                "operationId": "Authentication",
                "parameters": [
                    {
                        "description": "User and Password ",
                        "name": "SignIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "User SignUp for use a Air IoT resource",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Sign Up",
                "parameters": [
                    {
                        "description": "New User information",
                        "name": "SignUp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/my": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Return User Information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/my/address": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "PostNewAddress user by id",
                "tags": [
                    "Users"
                ],
                "summary": "PostNewAddress user by id",
                "parameters": [
                    {
                        "description": "Customer Address",
                        "name": "CustomerAddress",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CustomerAddress"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/my/info": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "post method user information update",
                "tags": [
                    "Users"
                ],
                "summary": "update user information",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "userInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        },
        "/my/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get user information by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "get user information by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user address",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Address information",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CustomerAddress"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete user by id",
                "tags": [
                    "Users"
                ],
                "summary": "Delete user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.ResendConfirmCode": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.SignInRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.SignUpRequest": {
            "type": "object",
            "required": [
                "password",
                "phone_no",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_no": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.UserConfirm": {
            "type": "object",
            "required": [
                "confirmationCode",
                "username"
            ],
            "properties": {
                "confirmationCode": {
                    "type": "string"
                },
                "username": {
                    "description": "code for confirm",
                    "type": "string"
                }
            }
        },
        "auth.UserConfirmNewPassword": {
            "type": "object",
            "required": [
                "confirmCode",
                "password",
                "userName"
            ],
            "properties": {
                "confirmCode": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 1
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "services.CustomerAddress": {
            "type": "object",
            "required": [
                "address",
                "amphur",
                "district",
                "lastName",
                "name",
                "province",
                "tel",
                "zipcode"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "address_default": {
                    "type": "boolean"
                },
                "amphur": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "tax": {
                    "type": "string"
                },
                "tax_default": {
                    "type": "boolean"
                },
                "tax_used": {
                    "type": "boolean"
                },
                "tel": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "integer"
                }
            }
        },
        "services.UpdateInfoRequest": {
            "type": "object",
            "required": [
                "lastname",
                "mobile",
                "name"
            ],
            "properties": {
                "lastname": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                }
            }
        },
        "utils.ApiResponse": {
            "type": "object",
            "properties": {
                "message": {},
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Air IoT API Service 2023",
	Description:      "Air Smart IoT App API Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
