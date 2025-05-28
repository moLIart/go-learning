package internal

import (
	"github.com/julienschmidt/httprouter"
	"github.com/moLIart/go-course/internal/handlers"
	"github.com/moLIart/go-course/internal/middlewares"
)

func RegisterHTTPRoutes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/players", middlewares.JWTAuth(handlers.CreatePlayerHandler))
	router.GET("/players", handlers.GetPlayersHandler)
	router.GET("/players/:id", handlers.GetPlayerByIDHandler)
	router.PUT("/players/:id", middlewares.JWTAuth(handlers.UpdatePlayerHandler))
	router.DELETE("/players/:id", middlewares.JWTAuth(handlers.DeletePlayerByIDHandler))

	router.POST("/rooms", middlewares.JWTAuth(handlers.CreateRoomHandler))
	router.GET("/rooms", handlers.GetRoomsHandler)
	router.GET("/rooms/:id", handlers.GetRoomByIDHandler)
	router.PUT("/rooms/:id", middlewares.JWTAuth(handlers.UpdateRoomHandler))
	router.DELETE("/rooms/:id", middlewares.JWTAuth(handlers.DeleteRoomHandler))

	router.POST("/boards", middlewares.JWTAuth(handlers.CreateBoardHandler))
	router.GET("/boards", handlers.GetBoardsHandler)
	router.GET("/boards/:id", handlers.GetBoardByIDHandler)
	router.PUT("/boards/:id", middlewares.JWTAuth(handlers.UpdateBoardHandler))
	router.DELETE("/boards/:id", middlewares.JWTAuth(handlers.DeleteBoardHandler))

	router.POST("/games", middlewares.JWTAuth(handlers.CreateGameHandler))
	router.GET("/games", handlers.GetGamesHandler)
	router.GET("/games/:id", handlers.GetGameByIDHandler)
	router.DELETE("/games/:id", middlewares.JWTAuth(handlers.DeleteGameHandler))

	router.Handler("GET", "/swagger/*any", handlers.SwaggerUIHandler())

	return router
}
