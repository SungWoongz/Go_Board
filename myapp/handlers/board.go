package handlers

import (
	"myapp/db"
	"myapp/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 로그인 페이지
func Login(c echo.Context) error {

	// POST 요청일 때
	if c.Request().Method == http.MethodPost {
		// 로그인 정보를 받아온다
		email := c.FormValue("email") // post요청에서 폼 데이터 값 가져옴
		password := c.FormValue("password")

		// 인증
		// db에서 해당 사용자 찾기
		var user entities.User
		result := db.Db.Where("email = ?", email).First(&user) //First(&user) 메서드는 첫 번째로 매칭되는 레코드를 user 변수에 저장

		// 사용자를 못 찾았을 때
		if result.Error != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}

		// DB에서 가져온 비밀번호와 클라이언트의 비밀번호를 비교

		// 비밀번호가 일치하지 않을 때
		if user.Password != password {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}

		// 인증 성공 시 메인 페이지로 redirect
		// TODO 세션 설정 or 토큰 생성

		return c.Redirect(http.StatusFound, "/main")
	}
	// GET 요청일 때
	return c.File("pages/login.html")
}

// 로그인 페이지 redirect
func RedirectLoginPage(c echo.Context) error {
	return c.Redirect(http.StatusFound, "login")
}

// 메인 페이지
func GetMainPage(c echo.Context) error {
	return c.File("pages/main.html")
}

// 사용자 등록
func PostRegUser(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// 데이터베이스에 사용자를 저장합니다.
	err := db.Db.Create(&user).Error
	if err != nil {
		return err
	}

	// 성공적으로 등록되었다는 메시지를 반환합니다.
	return c.JSON(http.StatusCreated, user)
}

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
