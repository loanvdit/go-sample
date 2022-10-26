// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/account/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Register new account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/account/verify/:token": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Active new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create key to login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/auth/:id": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create user token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/partners": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get all partner of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Store edition partner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bank account",
                        "name": "bank_account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Register new partner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bank account",
                        "name": "bank_account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/partners/:id": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a partner",
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
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/statistic/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get details report",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions of user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Register new transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Partner Id",
                        "name": "partner_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Bank Id",
                        "name": "bank_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Bank branch Id",
                        "name": "bank_branch_id",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Amount",
                        "name": "amount",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Fee",
                        "name": "fee",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Message",
                        "name": "message",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Currency",
                        "name": "currency",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/transactions/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get transaction details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "Store edition transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction Id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Partner Id",
                        "name": "partner_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Bank Id",
                        "name": "bank_id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Bank branch Id",
                        "name": "bank_branch_id",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Amount",
                        "name": "amount",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Fee",
                        "name": "fee",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Message",
                        "name": "message",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Currency",
                        "name": "currency",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/transactions/search/:start_date/:end_date": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Search transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start date",
                        "name": "start_date",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End date",
                        "name": "end_date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/transactions/status/:id/:status": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Update transaction status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Status (0-Create; 1-Pending; 2-Success)",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Import Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image",
                        "name": "image",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.ResponseNG"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ResponseNG": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "app.ResponseOk": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "boolean"
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin API",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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