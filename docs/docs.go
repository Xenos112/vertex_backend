// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user/{tag}": {
            "get": {
                "description": "Get a user by their unique tag.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get the User By Tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag of the user",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.User": {
            "type": "object",
            "properties": {
                "banner_id": {
                    "type": "string"
                },
                "bio": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "discord_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "github_id": {
                    "type": "string"
                },
                "google_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "routes.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Gin Swagger Example API",
	Description:      "This is a sample server for Gin using Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
