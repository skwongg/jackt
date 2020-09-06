package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/skwongg/jackt/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var lifts = []models.Lift{
	models.Lift{
		Name:        "Name 1",
		Description: "Hello world 1",
	},
	models.Lift{
		Name:        "Name 2",
		Description: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Lift{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Lift{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Lift{}).Error
	//.AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&models.Lift{}).Create(&lifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed lifts table: %v", err)
		}
	}
}
