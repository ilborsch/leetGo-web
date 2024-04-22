package storage

import (
	"github.com/ilborsch/leetGo-web/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

const fileName = "storage/leetgo.db"

type Storage struct {
	db *gorm.DB
	mu sync.Mutex
}

func New() *Storage {
	conn := connectionMustLoad()
	return &Storage{db: conn}
}

func connectionMustLoad() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	runMigrations(db)
	return db
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Problem{},
		&models.Article{},
		&models.Tag{},
	)
	if err != nil {
		panic("failed to migrate models")
	}
}
