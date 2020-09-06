package seed

import (
	"fmt"
	"log"

	"github.com/skwongg/jackt/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var lifts = []models.Lift{
	{
		Name:        "Name 1",
		Description: "Hello world 1",
	},
	{
		Name:        "Name 2",
		Description: "Hello world 2",
	},
}

//Load seeds the database with fake entries.. this should be removed at some point so it stops dropping.
func Load(db *gorm.DB) {

	fmt.Println("Seeding begin...")
	if (db.Migrator().HasTable(&models.Lift{}) || db.Migrator().HasTable(&models.User{})) {
		err := db.Migrator().DropTable(&models.Lift{})
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
		err = db.Migrator().DropTable(&models.User{})
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.AutoMigrate(&models.Lift{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err := db.Create(&models.User{Nickname: users[i].Nickname,
			Email:    users[i].Email,
			Password: users[i].Password}).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i := range lifts {
		err := db.Create(&models.Lift{Name: lifts[i].Name, Description: lifts[i].Description}).Error
		if err != nil {
			log.Fatalf("cannot seed lifts table: %v", err)
		}
	}
	fmt.Println("\nSeeding Complete...")
}
