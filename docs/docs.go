// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Yaroslav Kuzmin",
            "email": "yarik1448kuzmin@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/notes": {
            "get": {
                "description": "Get all notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "List notes",
                "responses": {
                    "200": {
                        "description": "Notes",
                        "schema": {
                            "$ref": "#/definitions/http.ListNotesResponse"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {}
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Create notes",
                "parameters": [
                    {
                        "description": "Note info",
                        "name": "noteInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CreateNoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Note created",
                        "schema": {
                            "$ref": "#/definitions/models.NoteTransfer"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {}
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {}
                    }
                }
            }
        },
        "/notes/{noteID}": {
            "get": {
                "description": "Get note by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Get note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "noteID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Note",
                        "schema": {
                            "$ref": "#/definitions/models.NoteTransfer"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {}
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Update note by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Update note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "noteID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Note info",
                        "name": "noteInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UpdateNoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated notes",
                        "schema": {
                            "$ref": "#/definitions/models.NoteTransfer"
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {}
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete notes by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Delete notes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "noteID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Note deleted"
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {}
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "http.CreateNoteRequest": {
            "type": "object",
            "properties": {
                "plain_text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.ListNotesResponse": {
            "type": "object",
            "properties": {
                "notes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.NoteTransfer"
                    }
                }
            }
        },
        "http.UpdateNoteRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "plain_text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.NoteTransfer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "plain_text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "Archipelago Notes API",
	Description:      "Notes API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
