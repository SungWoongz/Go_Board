package handlers

import (
	"myapp/db"
	"myapp/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 모든 게시판 글 조회
func GetAllBoard(c echo.Context) error {
	var users []entities.Board
	// Db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// 글 작성
func PostBoard(c echo.Context) error {
	board := new(entities.Board)
	if err := c.Bind(board); err != nil {
		return err
	}
	db.Db.Create(&board)
	return c.JSON(http.StatusCreated, board)
}

// id로 글 상세조회
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var board entities.Board
	db.Db.First(&board, id)
	return c.JSON(http.StatusOK, board)
}

// 게시판 글 삭제 by id
func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	var board entities.Board
	db.Db.Delete(&board, id)
	return c.NoContent(http.StatusNoContent)
}
