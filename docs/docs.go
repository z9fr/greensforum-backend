// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://dasith.works",
        "contact": {
            "name": "API support",
            "url": "http://greenforum.com/support",
            "email": "z9fr@proton.me"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/question/create": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "create a new question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Create a new Question",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question.QuestionCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question.Question"
                        }
                    }
                }
            }
        },
        "/question/{qid}/answer/create": {
            "post": {
                "description": "Answer to a question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Answer"
                ],
                "summary": "Write Answer",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question.AnswerRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "qid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question.Answer"
                            }
                        }
                    }
                }
            }
        },
        "/user/join": {
            "post": {
                "description": "register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "login as a existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Authenticate User",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.AuthRequest"
                        }
                    }
                }
            }
        },
        "/view/questions": {
            "get": {
                "description": "Get all the posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "fetch all posts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Next Post",
                        "name": "next_post",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question.Question"
                            }
                        }
                    }
                }
            }
        },
        "/view/questions/{tag}": {
            "get": {
                "description": "find posts by using tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "get posts by tags",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag Name",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question.Question"
                            }
                        }
                    }
                }
            }
        },
        "/view/search": {
            "get": {
                "description": "Search Posts based on a keyword",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "search for posts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search Query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question.Question"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "question.Answer": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "down_vote_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "question_id": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "up_vote_count": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "view_count": {
                    "type": "integer"
                }
            }
        },
        "question.AnswerRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "question.Question": {
            "type": "object",
            "properties": {
                "answer_count": {
                    "type": "integer"
                },
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/question.Answer"
                    }
                },
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "string"
                },
                "down_vote_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_answered": {
                    "type": "boolean"
                },
                "score": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/question.Tag"
                    }
                },
                "title": {
                    "description": "QuestionID    int      ` + "`" + `gorm:\"column:question_id primaryKey\" json:\"question_id\"` + "`" + `",
                    "type": "string"
                },
                "up_vote_count": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "view_count": {
                    "type": "integer"
                }
            }
        },
        "question.QuestionCreateRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/question.Tag"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "question.Tag": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "types.AuthRequest": {
            "type": "object",
            "properties": {
                "expire": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "types.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.Account": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "is_employee": {
                    "type": "boolean"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "reputation": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "website_url": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "object",
                    "properties": {
                        "description": {
                            "type": "string"
                        },
                        "display_name": {
                            "type": "string"
                        },
                        "location": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "profile_image": {
                            "type": "string"
                        },
                        "website_url": {
                            "type": "string"
                        }
                    }
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/user.Account"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user_type": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.staging.green.dasith.works",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Green Forum Backend",
	Description:      "REST api documentation for green forum",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
