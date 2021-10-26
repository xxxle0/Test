package repositories

import (
	"github.com/duybkit13/api/entities"
	"gorm.io/gorm"
)

type ScanResultRepositoryI interface {
	Create(scanResult entities.Result) error
	FindOne(condition map[string]interface{}) (entities.Result, error)
	FindMany(condition map[string]interface{}) ([]entities.Result, error)
	Delete(condition map[string]interface{}) error
	Update(condition map[string]interface{}, newPayload entities.Result) error
}

type ScanResultRepository struct {
	dbClient *gorm.DB
}

func InitScanResultRepository(dbClient *gorm.DB) ScanResultRepositoryI {
	return &ScanResultRepository{
		dbClient: dbClient,
	}
}

func (r *ScanResultRepository) Create(scanResult entities.Result) error {
	result := r.dbClient.Create(&scanResult)
	return result.Error
}

func (r *ScanResultRepository) FindOne(condition map[string]interface{}) (entities.Result, error) {
	var scanResult entities.Result
	result := r.dbClient.Model(&entities.Result{}).Where(condition).First(&scanResult)
	if result.Error != nil {
		return scanResult, result.Error
	}
	return scanResult, nil
}

func (r *ScanResultRepository) FindMany(condition map[string]interface{}) ([]entities.Result, error) {
	var scanResults []entities.Result
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Find(&scanResults)
	if result.Error != nil {
		return scanResults, result.Error
	}
	return scanResults, nil
}

func (r *ScanResultRepository) Delete(condition map[string]interface{}) error {
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Delete(&entities.Result{})
	return result.Error
}

func (r *ScanResultRepository) Update(condition map[string]interface{}, newPayload entities.Result) error {
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Updates(newPayload)
	return result.Error
}
