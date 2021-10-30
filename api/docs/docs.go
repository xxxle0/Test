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
        "/v1/scan-results": {
            "get": {
                "description": "Get Scan Result List with Offset \u0026 Payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scan Results"
                ],
                "summary": "Get Scan Result List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.GetScanResultListResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Scan Result With ScanResult Payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scan Results"
                ],
                "summary": "Create Scan Result",
                "parameters": [
                    {
                        "description": "The Request Payload to create Scan Result",
                        "name": "RequestPayload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateScanResultDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateScanResultResp"
                        }
                    }
                }
            }
        },
        "/v1/scan-results/{id}": {
            "get": {
                "description": "Get Scan Result Detail By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scan Results"
                ],
                "summary": "Get Scan Result Detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Scan Result ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.GetScanResultDetailResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Scan Result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scan Results"
                ],
                "summary": "Delete Scan Result",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Scan Result ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.DeleteScanResultResp"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Scan Result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scan Results"
                ],
                "summary": "Update Scan Result",
                "parameters": [
                    {
                        "description": "The Request Payload to update Scan Result",
                        "name": "RequestPayload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateScanResultDto"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Scan Result ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateScanResultResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateScanResultDto": {
            "type": "object",
            "properties": {
                "findings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.FindingDto"
                    }
                },
                "finished_at": {
                    "type": "string"
                },
                "queued_at": {
                    "type": "string"
                },
                "repository_name": {
                    "type": "string"
                },
                "scanning_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dtos.CreateScanResultResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dtos.DeleteScanResultResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dtos.FindingDto": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "object",
                    "additionalProperties": true
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": true
                },
                "ruleId": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dtos.GetScanResultDetailResp": {
            "type": "object",
            "properties": {
                "scan_result": {
                    "$ref": "#/definitions/dtos.ScanResultResp"
                }
            }
        },
        "dtos.GetScanResultListResp": {
            "type": "object",
            "properties": {
                "scan_results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.ScanResultResp"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "dtos.ScanResultResp": {
            "type": "object",
            "properties": {
                "findings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.FindingDto"
                    }
                },
                "finished_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "queued_at": {
                    "type": "string"
                },
                "repository_name": {
                    "type": "string"
                },
                "scanning_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dtos.UpdateScanResultDto": {
            "type": "object",
            "properties": {
                "findings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.FindingDto"
                    }
                },
                "finished_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "queued_at": {
                    "type": "string"
                },
                "repository_name": {
                    "type": "string"
                },
                "scanning_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dtos.UpdateScanResultResp": {
            "type": "object",
            "properties": {
                "message": {
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
	swag.Register("swagger", &s{})
}
