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

func getUsers(c echo.Context) error {
	var users []Board
	db.Find(&users)
	return c.JSON(200, users)
}

func createUser(c echo.Context) error {
	user := new(Board)
	if err := c.Bind(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(201, user)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	var user Board
	db.First(&user, id)
	return c.JSON(200, user)
}

func deleteUserByID(c echo.Context) error {
	id := c.Param("id")
	var user Board
	db.Delete(&user, id)
	return c.NoContent(204)
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
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUserByID)
	e.DELETE("/users/:id", deleteUserByID)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
