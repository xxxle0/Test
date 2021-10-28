package databases

import (
	"github.com/duybkit13/api/configurations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresQLClient(config configurations.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// db.AutoMigrate(&entities.Finding{})
	return db, err
}
