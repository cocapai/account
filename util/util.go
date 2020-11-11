package util

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"scutrobot.buff/go_demo/model"
	"time"
)

func RandomString(n int)  string{
	var letters = []byte("fdasfaqtweqgtrjytketfasdfarj6")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}