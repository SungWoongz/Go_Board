package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	ID        uint      `json:"id" db:"id"`                 // 유저 id
	Author    string    `json:"name" db:"name"`             // 작성자
	Boardid   int       `json:"board_id" db:"board_id"`     // 게시글 id
	Title     string    `json:"title" db:"title"`           // 제목
	Content   string    `json:"content" db:"content"`       // 내용
	View      int       `json:"view" db:"view"`             // 조회수
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 작성일

}

var (
	db *gorm.DB
)

func init() {
	var err error
	// dsn := "postgresql://user:password@localhost/database_name?sslmode=disable" // Update with your database credentials
	dsn := "host=13.209.44.115 user=swy password=1234 dbname=boarddb port=5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	db.AutoMigrate(&Board{})
}

// 모든 게시판 글 조회
func getUsers(c echo.Context) error {
	var users []Board
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// 글 작성
func postBoard(c echo.Context) error {
	board := new(Board)
	if err := c.Bind(board); err != nil {
		return err
	}
	db.Create(&board)
	return c.JSON(http.StatusCreated, board)
}

// id로 글 상세조회
func getUserByID(c echo.Context) error {
	id := c.Param("id")
	var board Board
	db.First(&board, id)
	return c.JSON(http.StatusOK, board)
}

// 게시판 글 삭제 by id
func deleteUserByID(c echo.Context) error {
	id := c.Param("id")
	var board Board
	db.Delete(&board, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/boards", getUsers)
	e.POST("/boards", postBoard)
	e.GET("/boards/:id", getUserByID)
	e.DELETE("/boards/:id", deleteUserByID)
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
