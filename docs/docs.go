// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-31 11:29:01.461470673 +0800 CST m=+0.030062326

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "A2OS Behavior API Documentation.",
        "title": "A2OS Behavior",
        "contact": {
            "name": "A2OS Dev Team",
            "url": "https://groups.google.com/group/a2os-general",
            "email": "a2os-general@googlegroups.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "0.0.1"
    },
    "host": "api.behavior.a2os.club",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/csrf": {
            "get": {
                "description": "Get CSRF token and cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "miscellaneous"
                ],
                "summary": "CSRF",
                "responses": {
                    "200": {
                        "description": "IN HEADER",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "X-CSRF-TOKEN": {
                                "type": "string",
                                "description": "CSRF Token hash value"
                            }
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Ping health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "miscellaneous"
                ],
                "summary": "PING-PONG",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/misc.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "misc.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "pong"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
