package databases

import (
	"github.com/duybkit13/api/configurations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresQLClient(config configurations.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	return db, err
}
