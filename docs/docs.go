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
        "/": {
            "get": {
                "produces": [
                    "text/html"
                ],
                "summary": "Get all habits",
                "responses": {
                    "200": {
                        "description": "HTML page",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/habits": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/html"
                ],
                "summary": "Create a new habit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description",
                        "name": "description",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "HTML page",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/habits/{id}": {
            "get": {
                "produces": [
                    "text/html"
                ],
                "summary": "Show edit form for a habit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Habit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "HTML page",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/html"
                ],
                "summary": "Update a habit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Habit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description",
                        "name": "description",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "HTML page",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "text/html"
                ],
                "summary": "Delete a habit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Habit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "HTML page",
                        "schema": {
                            "type": "string"
                        }
                    }
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
	Title:            "Habit Tracker API",
	Description:      "This is an API for a habit tracker app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
