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
        "/api/admin": {
            "get": {
                "description": "单个管理员",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "单个管理员",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.firstResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            },
            "post": {
                "description": "创建管理员",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "创建管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.createResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/{id}": {
            "delete": {
                "description": "删除管理员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "删除管理员",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.deleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admins": {
            "get": {
                "description": "管理员列表",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "管理员列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.listResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "admin.deleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "admin.firstResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "admin.listData": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "admin.listResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.listData"
                    }
                }
            }
        },
        "code.Failure": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务码",
                    "type": "integer"
                },
                "message": {
                    "description": "描述信息",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gin-api-mono 接口文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}