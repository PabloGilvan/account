// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2021-08-09 09:37:29.446203913 -0300 -03 m=+0.005300623

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "description": "Creates a new Account generating am account number",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Creates an Account",
                "parameters": [
                    {
                        "description": "Object for persisting the account",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.AccountPersist"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/account.AccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "Load an account by the UUID identifier\nAccounts resources can be used to create (\"/accounts\" POST)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Load an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person's identification code",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/account.AccountResponse"
                        }
                    },
                    "400": {
                        "description": "account inactive"
                    },
                    "404": {
                        "description": "account not found"
                    }
                }
            }
        },
        "/transactions/{id}": {
            "get": {
                "description": "Persist a transaction validating the operation type and account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Persist a transaction",
                "parameters": [
                    {
                        "description": "Object for persisting the transaction",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/transaction.TransactionPersist"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/transaction.TransactionResponse"
                        }
                    },
                    "400": {
                        "description": "account inactive"
                    },
                    "404": {
                        "description": "account not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "account.AccountPersist": {
            "type": "object",
            "required": [
                "document_number"
            ],
            "properties": {
                "document_number": {
                    "type": "string"
                }
            }
        },
        "account.AccountResponse": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "createDate": {
                    "type": "string"
                },
                "document_number": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "updateDate": {
                    "type": "string"
                }
            }
        },
        "transaction.TransactionPersist": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "operation_type_id": {
                    "type": "integer"
                }
            }
        },
        "transaction.TransactionResponse": {
            "type": "object",
            "properties": {
                "transaction_identifier": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
