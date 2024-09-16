package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error 

	db, err = InitDatabase()

	if(err!= nil){
		logger.Warnf("Database initialization error %v", err.Error());
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}


func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

