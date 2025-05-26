package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// คือฟังก์ชันที่จะจัดการกับการสร้างข้อมูลผู้ใช้ใหม่
func GenerateJWT(userID uint) (string, error) {
	// สร้างข้อมูลผู้ใช้ใหม่
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	// สร้างข้อมูลผู้ใช้ใหม่
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
