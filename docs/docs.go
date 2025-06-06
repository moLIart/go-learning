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
        "/boards": {
            "get": {
                "description": "Returns a list of all boards.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "Get all boards",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetBoardDto"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to encode boards",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new board with the given size and adds it to the repository.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "Create a new board (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Board size",
                        "name": "board",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBoardDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetBoardDto"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/boards/{id}": {
            "get": {
                "description": "Returns a board by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "Get board by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetBoardDto"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Board not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates the size of a board by its ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "Update board by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated board size",
                        "name": "board",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBoardDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter or request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Board not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a board by its ID.",
                "tags": [
                    "boards"
                ],
                "summary": "Delete board by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Board not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/games": {
            "get": {
                "description": "Returns a list of all games.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Get all games",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetGameDto"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to encode games",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new game and adds it to the repository.",
                "tags": [
                    "games"
                ],
                "summary": "Create a new game (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetGameDto"
                        }
                    }
                }
            }
        },
        "/games/{id}": {
            "get": {
                "description": "Returns a game by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Get game by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetGameDto"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Game not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a game by its ID.",
                "tags": [
                    "games"
                ],
                "summary": "Delete game by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Game not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates user with login and password from environment variables. Returns JWT token on success.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Could not generate token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/players": {
            "get": {
                "description": "Returns a list of all players.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get all players",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetPlayerDto"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to encode players",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new player with the given name and adds it to the repository.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Create a new player (Requires authorization)",
                "parameters": [
                    {
                        "description": "Player name",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePlayerDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPlayerDto"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/players/{id}": {
            "get": {
                "description": "Returns a player by their ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get player by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPlayerDto"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Player not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates the name of a player by their ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Update player by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated player name",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePlayerDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter or request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Player not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a player by their ID.",
                "tags": [
                    "players"
                ],
                "summary": "Delete player by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Player not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rooms": {
            "get": {
                "description": "Returns a list of all rooms.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Get all rooms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetRoomDto"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to encode rooms",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new room with the given code and adds it to the repository.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Create a new room (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Room code",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateRoomDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetRoomDto"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rooms/{id}": {
            "get": {
                "description": "Returns a room by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Get room by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetRoomDto"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Room not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates the code of a room by its ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Update room by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated room code",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateRoomDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter or request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Room not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a room by its ID.",
                "tags": [
                    "rooms"
                ],
                "summary": "Delete room by ID (Requires authorization)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Room not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateBoardDto": {
            "type": "object",
            "properties": {
                "size": {
                    "type": "integer"
                }
            }
        },
        "dto.CreatePlayerDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateRoomDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "dto.GetBoardDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "dto.GetGameDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.GetPlayerDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.GetRoomDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateBoardDto": {
            "type": "object",
            "properties": {
                "size": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdatePlayerDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateRoomDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Test API Server",
	Description:      "API Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
