// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://example.com/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Authenticate user with credentials and generate access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Authenticate user and generate access token",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error generating token",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/countries": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a list of all countries - For testing purpose only. Because it retrieves all the details it will take some time to load (8-12 seconds)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "countries"
                ],
                "summary": "Retrieve a list of all countries details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "country data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/countries/filter": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a filtered and sorted list of countries based on specified parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "countries"
                ],
                "summary": "Get a filtered and sorted list of countries",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Filter countries by population.",
                        "name": "population",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter countries by area.",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter countries by language.",
                        "name": "language",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc).",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "paginated list of countries",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid filter or page parameters",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error fetching or decoding countries data",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/country": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get detailed information about a specific country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "countries"
                ],
                "summary": "Get detailed information about a specific country",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the country to fetch.",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "country data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Error: Please provide a valid country name in the query parameters",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error fetching country data or decoding country data",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Credentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "main.ErrorResponse": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "assignment.snifyak.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "API_Assignment | Swagger",
	Description:      "API Documentation for all the endpoints",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
