package config

import (
	"os"

	"github.com/lucasdbr05/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database initialization")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if(os.IsNotExist(err)) {
		logger.Info("Database file not found")

		err := os.Mkdir("./db", os.ModePerm)

		if (err != nil) {
			return nil, err
		}

		file, err := os.Create(dbPath)
		
		if(err != nil) {
			return nil, err 
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if(err != nil) {
		logger.Errorf("Database opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if(err != nil) {
		logger.Errorf("Database automigrate error %v", err)
		return nil, err
	}

	return db, nil
}
