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
        "/portfolio/projects": {
            "get": {
                "description": "Get project list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Project list",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "Language ID",
                        "name": "tech_id",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Is active",
                        "name": "is_active",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Is archived",
                        "name": "is_archived",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Is developing",
                        "name": "is_developing",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort field",
                        "name": "sort_field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of projects",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset of projects",
                        "name": "Offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Project"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create project and write to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Create Project",
                "parameters": [
                    {
                        "description": "Project title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Project version",
                        "name": "version",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Project description",
                        "name": "description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Language ID",
                        "name": "language_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Is active",
                        "name": "isActive",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Is archived",
                        "name": "isArchived",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Is developing",
                        "name": "isDeveloping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Links",
                        "name": "links",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project ID",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            }
        },
        "/portfolio/projects/{id}": {
            "get": {
                "description": "Get project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Project nor found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Delete Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "patch": {
                "description": "Update project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Update Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Project title",
                        "name": "title",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Project version",
                        "name": "version",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Project description",
                        "name": "description",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Language ID",
                        "name": "language_id",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Is active",
                        "name": "isActive",
                        "in": "body",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Is archived",
                        "name": "isArchived",
                        "in": "body",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Is developing",
                        "name": "isDeveloping",
                        "in": "body",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    {
                        "description": "Links",
                        "name": "links",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Technology not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            }
        },
        "/portfolio/techs": {
            "get": {
                "description": "Get technology list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Technology list",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "Technology ID",
                        "name": "tech_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort field",
                        "name": "sort_field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of projects",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset of projects",
                        "name": "Offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Technology",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Technology"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create technology and write to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Create Technology",
                "parameters": [
                    {
                        "description": "Technology name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Technology svg",
                        "name": "svg",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Technology ID",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            }
        },
        "/portfolio/techs/{id}": {
            "get": {
                "description": "Get technology",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Technology",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Technology ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Technology",
                        "schema": {
                            "$ref": "#/definitions/models.Technology"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Technology not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete technology",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Delete Technology",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Technology ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Technology not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            },
            "patch": {
                "description": "Update technology",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "Update Technology",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Technology ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Technology name",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Technology svg",
                        "name": "svg",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Technology not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Project": {
            "type": "object",
            "properties": {
                "dscription": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "isArchived": {
                    "type": "boolean"
                },
                "isDeveloping": {
                    "type": "boolean"
                },
                "links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "technologies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Technology"
                    }
                },
                "title": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Technology": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "svg": {
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
