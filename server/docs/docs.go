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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "allows healthcheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "healthcheck",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.AppResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login user given valid email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "email \u0026 password",
                        "name": "userLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.LoginSuccess"
                        }
                    },
                    "401": {
                        "description": "Missing/Incorrect credentials",
                        "schema": {
                            "$ref": "#/definitions/model.LoginFailureCredentials"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Logout user by removing jwt cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.LogoutSuccess"
                        }
                    },
                    "401": {
                        "description": "Missing/Expired token",
                        "schema": {
                            "$ref": "#/definitions/model.Unauthorized"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/refresh_token": {
            "post": {
                "description": "Refresh user's access token given valid refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.RefreshSuccess"
                        }
                    },
                    "401": {
                        "description": "Missing/Expired token",
                        "schema": {
                            "$ref": "#/definitions/model.Unauthorized"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register user given valid firstname, lastname, email (unique), password, phone (unique)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "firstname, lastname, email, password, phone",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.InvalidJsonBody"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.GetAllUserSuccess"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user user given valid firstname, lastname, email (unique), password, phone (unique)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "firstname, lastname, email, password, phone",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.InvalidJsonBody"
                        }
                    },
                    "401": {
                        "description": "Missing/Expired token",
                        "schema": {
                            "$ref": "#/definitions/model.Unauthorized"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user/:id": {
            "get": {
                "description": "Get user given a valid ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.GetUserByIDSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.InvalidID"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.UserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user user given valid ID, firstname, lastname, email (unique), phone (unique)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "firstname, lastname, email, phone",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid ID/body",
                        "schema": {
                            "$ref": "#/definitions/model.InvalidJsonBody"
                        }
                    },
                    "401": {
                        "description": "Missing/Expired token",
                        "schema": {
                            "$ref": "#/definitions/model.Unauthorized"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.UserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user user given valid ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.DeleteUserSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.InvalidID"
                        }
                    },
                    "401": {
                        "description": "Missing/Expired token",
                        "schema": {
                            "$ref": "#/definitions/model.Unauthorized"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.UserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.InternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AppResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ValError"
                    }
                }
            }
        },
        "model.DeleteUserSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "User deleted successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.GetAllUserSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Users retrieved successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.GetUserByIDSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.User"
                },
                "message": {
                    "type": "string",
                    "example": "User \u003cid\u003e retrieved successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.InternalServerError": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.InvalidID": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "invalid id"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.InvalidJsonBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "invalid json request body"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ValError"
                    }
                }
            }
        },
        "model.LoginFailureCredentials": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "incorrect/missing email or password"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "yoimiya.naganohara@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "model.LoginSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "Login successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.LogoutSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "Logout successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.RefreshSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "Token refreshed successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "yoimiya.naganohara@gmail.com"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "yoimiya"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "naganohara"
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8,
                    "example": "12345678"
                },
                "phone": {
                    "type": "string",
                    "example": "+33612345678"
                }
            }
        },
        "model.RegisterSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "User created successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.RoadTrip": {
            "type": "object",
            "properties": {
                "enddate": {
                    "type": "string"
                },
                "locations": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                },
                "startdate": {
                    "type": "string"
                }
            }
        },
        "model.Unauthorized": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "cookie token is empty / Token is expired"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.UpdateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "yoimiya.naganohara@gmail.com"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "yoimiya"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "naganohara"
                },
                "phone": {
                    "type": "string",
                    "example": "+33612345678"
                }
            }
        },
        "model.UpdateUserSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "User updated successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "phone": {
                    "type": "string"
                },
                "roadTrip": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RoadTrip"
                    }
                }
            }
        },
        "model.UserNotFound": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "user not found"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "valErrors": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "model.ValError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}