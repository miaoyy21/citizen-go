package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"time"
)

type User struct {
	Id        string
	Name      string
	StartTime *time.Time
}

func Test(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/pc28?charset=utf8mb4&collation=utf8mb4_general_ci&loc=Local&parseTime=true")
	if err != nil {
		log.Fatalf("gorm.Open() %s", err.Error())
	}

	users := make([]*User, 0)

	db.LogMode(true)
	if err := db.Model(&User{}).Where("start_time > ?", time.Now().UTC()).Find(&users).Error; err != nil {
		log.Fatalf("%s", err.Error())
	}

	for _, user := range users {
		log.Fatalf("USER %#v\n", user)
	}
}
