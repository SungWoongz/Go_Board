package main

import (
	"myapp/db"
	"myapp/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// PostgreSQL 데이터베이스 DSN 설정
	dsn := "host=13.209.44.115 user=swy password=1234 dbname=boarddb port=5432"
	// 데이터베이스 초기화
	db.InitDatabase(dsn)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/login", func(c echo.Context) error {
	// 	return c.File("pages/login.html")
	// })

	// 로그인 페이지 리다이렉트
	e.GET("/", handlers.RedirectLoginPage)
	// 로그인 핸들러
	e.GET("/login", handlers.Login)
	e.POST("/login", handlers.Login)
	// 메인페이지
	e.GET("/main", handlers.GetMainPage)

	// Routes
	// 게시판 조회
	e.GET("/boards", handlers.GetAllBoard)
	e.POST("/boards", handlers.PostBoard)
	e.GET("/boards/:id", handlers.GetUserByID)
	e.DELETE("/boards/:id", handlers.DeleteUserByID)

	// 유저 등록
	e.POST("/users", handlers.PostRegUser)

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

// 유저 등록
// curl -X POST http://localhost:1323/users -d '{"username":"sungwoong","userid":"admin","password":"admin"}' -H "Content-Type: application/json"
