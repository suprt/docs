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
        "/api/address/geocode": {
            "post": {
                "description": "Search place address by geocode",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Geocode search",
                "operationId": "geocode-handler",
                "parameters": [
                    {
                        "description": "Latitude and Longitude of place",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.GeocodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of searched address",
                        "schema": {
                            "$ref": "#/definitions/main.SearchResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/api/address/search": {
            "post": {
                "description": "Search place by address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Address Search",
                "operationId": "search-handler",
                "parameters": [
                    {
                        "description": "Search request query",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of searched address",
                        "schema": {
                            "$ref": "#/definitions/main.SearchResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "house": {
                    "type": "string"
                },
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "main.GeocodeRequest": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lng": {
                    "type": "string"
                }
            }
        },
        "main.SearchRequest": {
            "type": "object",
            "properties": {
                "query": {
                    "type": "string"
                }
            }
        },
        "main.SearchResponse": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Address"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GeoCode API",
	Description:      "GeoCode API Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
