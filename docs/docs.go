// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://uzsk.iamtakagi.net",
        "contact": {
            "name": "yude",
            "email": "i@yude.jp"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/players": {
            "get": {
                "description": "Get a list of players with optional filtering and sorting, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get players",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "example": "550e8400-e29b-41d4-a716-446655440000",
                        "description": "Filter criteria",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "asc",
                        "example": "desc",
                        "description": "Sort order",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "example": 0,
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 50,
                        "example": 10,
                        "description": "Limit for pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "play_time",
                        "description": "Order by field",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Player"
                            }
                        }
                    }
                }
            }
        },
        "/servers/{name}": {
            "get": {
                "description": "Get servers registered to uzsk-api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Get servers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of target server",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Server"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.Player": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "integer"
                },
                "experience": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "initialLoginDate": {
                    "type": "string"
                },
                "lastLoginDate": {
                    "type": "string"
                },
                "totalBuildBlocks": {
                    "type": "integer"
                },
                "totalDestroyBlocks": {
                    "type": "integer"
                },
                "totalMobKills": {
                    "type": "integer"
                },
                "totalPlayTime": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "structs.Server": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "uzsk-api.iamtakagi.net",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "uzsk-api",
	Description:      "Public Web API for uzsk.iamtakagi.net",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
