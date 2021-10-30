package databases

import (
	"github.com/duybkit13/api/configurations"
	"github.com/duybkit13/api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresQLClient(config configurations.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if config.DBAutoMigration {
		db.AutoMigrate(&entities.Result{}, &entities.Finding{})
	}
	return db, err
}
