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
        "/gate/{id}": {
            "get": {
                "description": "Retrieve a gate by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gate"
                ],
                "summary": "Get gate by ID",
                "operationId": "get-gate-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Gate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "Gate not found"
                    }
                }
            },
            "put": {
                "description": "Update gate of given id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gate"
                ],
                "summary": "Update ga",
                "operationId": "update-gate",
                "parameters": [
                    {
<<<<<<< HEAD
=======
                        "description": "Updated gate object",
                        "name": "gate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Gate"
                        }
                    },
                    {
>>>>>>> 844c91c ([Shaz|Madhavan] added body annotation)
                        "type": "string",
                        "description": "Gate ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Gate updated sucessfully"
                    },
                    "400": {
                        "description": "Gate not found"
                    }
                }
            }
        },
        "/gates": {
            "get": {
                "description": "get all the gate details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gate"
                ],
                "summary": "Get all gates",
                "operationId": "get-all-gate",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default = 0)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "filter by floor (default = all floor)",
                        "name": "floor",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Gate": {
            "type": "object",
            "properties": {
                "floor_number": {
                    "type": "integer"
                },
                "gate_number": {
                    "type": "integer"
                },
                "id": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
