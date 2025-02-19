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
        "/welcome": {
            "get": {
                "description": "this api is like a \"hello world\"",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "RPC-like"
                ],
                "summary": "welcome api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务名称",
                        "name": "service",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "操作方法（create/update/delete）",
                        "name": "method",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "服务不存在",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2",
	Host:             "localhost:9000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "习惯养成 API 文档",
	Description:      "这是一个基于 Gin 和 RPC 风格的习惯养成 API。应当只有开发环境下能看到该文档。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
