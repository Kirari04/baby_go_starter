package database

import (
	"baby_starter/app"
	"baby_starter/database/model"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {

	dbFullPath := path.Join(app.ENV.WorkDir, app.ENV.Datbase)

	// create workdir if not exists
	if _, err := os.Stat(app.ENV.WorkDir); os.IsNotExist(err) {
		if err := os.MkdirAll(app.ENV.WorkDir, os.ModePerm); err != nil {
			app.LOG.Fatal().Err(err).Msg("failed to create workdir")
		}
	}

	// configure database
	db, err := gorm.Open(sqlite.Open(dbFullPath), &gorm.Config{})
	if err != nil {
		app.LOG.Fatal().Err(err).Msg("failed to open database")
	}
	app.DB = db

	// migrations
	if err := db.AutoMigrate(&model.User{}); err != nil {
		app.LOG.Fatal().Err(err).Msg("failed to run migrations")
	}
}
