// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/:id": {
            "get": {
                "description": "Non-authenticated endpoint fetches account at specified key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get account record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aerospike.RecordViewV1"
                        }
                    }
                }
            },
            "put": {
                "description": "Non-authenticated endpoint creates an empty record at the specified key. Overwrites any record that already exists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create account record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aerospike.AccountViewV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/credentials": {
            "post": {
                "description": "Non-authenticated endpoint that adds grafana and confluence-server users to an account. Assumes entries are pre-validated",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credentials"
                ],
                "summary": "Add credentials to an account",
                "parameters": [
                    {
                        "description": "Add credentials",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/credentials.AddCredentialsV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/credentials.AddedCredentialsV1"
                        }
                    }
                }
            }
        },
        "/credentials/check": {
            "post": {
                "description": "Non-authenticated endpoint Check credentials for validity. Returns an array of user objects with check result",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credentials"
                ],
                "summary": "Check credentials for validity",
                "parameters": [
                    {
                        "description": "Check credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/credentials.CheckCredentialsV1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/credentials.CheckCredentialsResultV1"
                        }
                    }
                }
            }
        },
        "/health/hello": {
            "get": {
                "description": "Non-authenticated endpoint that returns 200 with hello message. Used to validate that the service is responsive.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Hello sanity endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.Ping"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aerospike.AccountViewV1": {
            "type": "object",
            "properties": {
                "Alias": {
                    "description": "Optional arg. Won't be returned if missing.",
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                }
            }
        },
        "aerospike.CredentialsView1": {
            "type": "object",
            "properties": {
                "GrafanaAPIUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/aerospike.GrafanaUser"
                    }
                }
            }
        },
        "aerospike.GrafanaUser": {
            "type": "object",
            "properties": {
                "Description": {
                    "type": "string"
                }
            }
        },
        "aerospike.MetadataViewV1": {
            "type": "object",
            "properties": {
                "CreateTimeUTC": {
                    "type": "string"
                },
                "LastUpdate": {
                    "type": "string"
                },
                "PrimaryKey": {
                    "type": "string"
                },
                "Version": {
                    "type": "string"
                }
            }
        },
        "aerospike.RecordViewV1": {
            "type": "object",
            "properties": {
                "Account": {
                    "type": "object",
                    "$ref": "#/definitions/aerospike.AccountViewV1"
                },
                "Credentials": {
                    "type": "object",
                    "$ref": "#/definitions/aerospike.CredentialsView1"
                },
                "Metadata": {
                    "type": "object",
                    "$ref": "#/definitions/aerospike.MetadataViewV1"
                }
            }
        },
        "credentials.AddConfluenceServerUserV1": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "credentials.AddCredentialsV1": {
            "type": "object",
            "properties": {
                "ConfluenceServerUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.AddConfluenceServerUserV1"
                    }
                },
                "GrafanaReadUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.AddGrafanaReadUserV1"
                    }
                },
                "accountID": {
                    "type": "string"
                }
            }
        },
        "credentials.AddGrafanaReadUserV1": {
            "type": "object",
            "properties": {
                "apikey": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "credentials.AddedCredentialsV1": {
            "type": "object",
            "properties": {
                "ConfluenceServerUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/credentials.AddConfluenceServerUserV1"
                    }
                },
                "GrafanaReadUsers": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/credentials.AddGrafanaReadUserV1"
                    }
                }
            }
        },
        "credentials.CheckConfluenceServerUserV1": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "credentials.CheckConfluenceUserResultV1": {
            "type": "object",
            "properties": {
                "Cause": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "result": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "credentials.CheckCredentialsResultV1": {
            "type": "object",
            "properties": {
                "confluenceServerUserCheck": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckConfluenceUserResultV1"
                    }
                },
                "grafanaReadUserCheck": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckGrafanaReadUserResultV1"
                    }
                }
            }
        },
        "credentials.CheckCredentialsV1": {
            "type": "object",
            "properties": {
                "confluenceServerUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckConfluenceServerUserV1"
                    }
                },
                "grafanaReadUsers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/credentials.CheckGrafanaReadUserV1"
                    }
                }
            }
        },
        "credentials.CheckGrafanaReadUserResultV1": {
            "type": "object",
            "properties": {
                "Cause": {
                    "type": "string"
                },
                "apikey": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "result": {
                    "type": "boolean"
                }
            }
        },
        "credentials.CheckGrafanaReadUserV1": {
            "type": "object",
            "properties": {
                "apikey": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "health.Ping": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "hello"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Graph Snapper API",
	Description: "Takes and updates snapshots from a graph service to a document store",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
