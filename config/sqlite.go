package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jtonynet/gopportunities/schemas"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not exists")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigrate error %v", err)
		return nil, err
	}

	return db, nil
}
