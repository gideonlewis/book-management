// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/books": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get List Book by Params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get Book",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListBookResponseWrapper"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Create a Book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create Book",
                "parameters": [
                    {
                        "description": "Book info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BookResponseWrapper"
                        }
                    }
                }
            }
        },
        "/books/all": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get all Book by unscoped",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get Book",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListBookResponseWrapper"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get an book",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/presenter.BookResponseWrapper"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Update a Book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update Book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BookResponseWrapper"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Delete Book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Delete an Book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/borrows": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get List Borrow by Params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "GetBorrow",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListBorrowResponseWrapper"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Create a Borrow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Create Borrow",
                "parameters": [
                    {
                        "description": "Borrow info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateBorrowRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BorrowResponseWrapper"
                        }
                    }
                }
            }
        },
        "/borrows/all": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get All Borrow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Get Borrow",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListBorrowResponseWrapper"
                        }
                    }
                }
            }
        },
        "/borrows/{id}": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get an Borrow by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Get Borrow",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/presenter.BorrowResponseWrapper"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Update an Borrow by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Update Borrow",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Borrow info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateBorrowRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BorrowResponseWrapper"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Delete an Borrow by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Delete Borrow",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/examples": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get example by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Get an example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListExampleResponseWrapper"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "create a example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Create example",
                "parameters": [
                    {
                        "description": "Example info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateExampleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ExampleResponseWrapper"
                        }
                    }
                }
            }
        },
        "/examples/{id}": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get example by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Get an example",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/presenter.ExampleResponseWrapper"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Update example by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Update an example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Example info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateExampleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ExampleResponseWrapper"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Delete example by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Delete an example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get List User by params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListUserResponseWrapper"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Create a User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.UserResponseWrapper"
                        }
                    }
                }
            }
        },
        "/users/all": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get All User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ListUserResponseWrapper"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get an User by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/presenter.UserResponseWrapper"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Update an User by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User info",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.UserResponseWrapper"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Delete an User by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "available_quantity": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "total_quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "model.Borrow": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "borrow_date": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "return_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Example": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "join_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "payload.CreateBookRequest": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "payload.CreateBorrowRequest": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "payload.CreateExampleRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "payload.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "join_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateBookRequest": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateBorrowRequest": {
            "type": "object",
            "properties": {
                "return_date": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateExampleRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "presenter.BookResponseWrapper": {
            "type": "object",
            "properties": {
                "book": {
                    "$ref": "#/definitions/model.Book"
                }
            }
        },
        "presenter.BorrowResponseWrapper": {
            "type": "object",
            "properties": {
                "borrow": {
                    "$ref": "#/definitions/model.Borrow"
                }
            }
        },
        "presenter.ExampleResponseWrapper": {
            "type": "object",
            "properties": {
                "example": {
                    "$ref": "#/definitions/model.Example"
                }
            }
        },
        "presenter.ListBookResponseWrapper": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Book"
                    }
                },
                "meta": {}
            }
        },
        "presenter.ListBorrowResponseWrapper": {
            "type": "object",
            "properties": {
                "borrows": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Borrow"
                    }
                },
                "meta": {}
            }
        },
        "presenter.ListExampleResponseWrapper": {
            "type": "object",
            "properties": {
                "examples": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Example"
                    }
                },
                "meta": {}
            }
        },
        "presenter.ListUserResponseWrapper": {
            "type": "object",
            "properties": {
                "meta": {},
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "presenter.UserResponseWrapper": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Example API",
	Description:      "Transaction API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
