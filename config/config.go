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
		logger.Warn("Database initialization error %v", err);
		return err
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
