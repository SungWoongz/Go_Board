package main

import (
	"myapp/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Routes
	e.GET("/boards", handlers.GetAllBoard)
	e.POST("/boards", handlers.PostBoard)
	e.GET("/boards/:id", handlers.GetUserByID)
	e.DELETE("/boards/:id", handlers.DeleteUserByID)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// 게시글 작성
// curl -X POST http://localhost:1323/boards -d '{"name":"Lee","board_id":1,"title":"First board","content":"First Content","view":0}' -H "Content-Type: application/json"
// 게시글 조회
// curl http://localhost:1323/boards
// 특정 게시글 조회
// curl http://localhost:1323/boards/1
// 게시글 수정하기
// curl -X POST http://localhost:1323/boards -d '{"name":"John","board_id":2,"title":"Updated Title","content":"Updated Content","view":0}' -H "Content-Type: application/json"
// 게시글 삭제하기
// curl -X DELETE http://localhost:1323/boards/1
