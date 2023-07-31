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
        "/health": {
            "get": {
                "description": "returns health status of the server",
                "produces": [
                    "application/json"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/station": {
            "get": {
                "description": "Get all stations or one by its UUID",
                "produces": [
                    "application/json"
                ],
                "summary": "get station",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Station ID",
                        "name": "uuid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Station"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONNotFoundResult"
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes station with given UUID from the service and database",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove a Station",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Station ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONSuccessResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONNotFoundResult"
                        }
                    }
                }
            }
        },
        "/station/register": {
            "post": {
                "description": "Add new station to the service and database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new Station",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Body"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Station"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONBadReqResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONNotFoundResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Body": {
            "type": "object",
            "required": [
                "place",
                "url"
            ],
            "properties": {
                "place": {
                    "type": "string",
                    "example": "Garden"
                },
                "url": {
                    "type": "string",
                    "example": "http://localhost:8080/data"
                }
            }
        },
        "controller.JSONBadReqResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Wrong Parameter"
                }
            }
        },
        "controller.JSONNotFoundResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "message": {
                    "type": "string",
                    "example": "Not found"
                }
            }
        },
        "controller.JSONSuccessResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Wrong Parameter"
                }
            }
        },
        "models.Station": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string",
                    "example": "2023-07-29T07:52:50Z"
                },
                "place": {
                    "type": "string",
                    "example": "Garden"
                },
                "url": {
                    "type": "string",
                    "example": "http://localhost:3000/data"
                },
                "uuid": {
                    "type": "string",
                    "example": "196bf376-e82b-4893-be62-3e5b5b7902e2"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
