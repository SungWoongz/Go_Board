package entities

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	ID        uint      `json:"id" db:"id"`                 // 게시글 id
	Author    string    `json:"name" db:"name"`             // 작성자
	Boardid   int       `json:"board_id" db:"board_id"`     // 게시글 id
	Title     string    `json:"title" db:"title"`           // 제목
	Content   string    `json:"content" db:"content"`       // 내용
	View      int       `json:"view" db:"view"`             // 조회수
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 작성일
}

type User struct {
	gorm.Model
	Username  string    `json:"username" db:"username"`     // 유저 이름
	Email     string    `json:"userid" db:"userid"`         // 유저 id
	Password  string    `json:"password" db:"password"`     // 유저 pw
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 로그인 시간
}
