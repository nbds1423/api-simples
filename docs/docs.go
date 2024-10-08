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
        "/api/": {
            "get": {
                "description": "Return a message with text 'Hello, World!'.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "private"
                ],
                "summary": "Hello, World!",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Agent",
                        "name": "User-Agent",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authentication Token",
                        "name": "x-nbds-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Return message with text \"Pong\".",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public"
                ],
                "summary": "Check status from server.",
                "responses": {
                    "200": {
                        "description": "200 response",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "404 response",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "429": {
                        "description": "429 response",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code http\nin: uint16",
                    "type": "integer"
                },
                "message": {
                    "description": "Message\nin: string",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "NBDS API using Gin",
	Description:      "NBDS Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
