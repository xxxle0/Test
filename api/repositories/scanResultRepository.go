package repositories

import (
	"context"

	"github.com/duybkit13/api/entities"
	"gorm.io/gorm"
)

type ScanResultRepositoryI interface {
	Create(ctx context.Context, scanResult entities.Result) error
	FindOne(ctx context.Context, condition map[string]interface{}) (entities.Result, error)
	FindMany(ctx context.Context, condition map[string]interface{}) ([]entities.Result, error)
	Delete(ctx context.Context, condition map[string]interface{}) error
	Update(ctx context.Context, condition map[string]interface{}, newPayload entities.Result) error
	Count(ctx context.Context, condition map[string]interface{}) int64
}

type ScanResultRepository struct {
	dbClient *gorm.DB
}

func InitScanResultRepository(dbClient *gorm.DB) ScanResultRepositoryI {
	return &ScanResultRepository{
		dbClient: dbClient,
	}
}

func (r *ScanResultRepository) Create(ctx context.Context, scanResult entities.Result) error {
	result := r.dbClient.Model(&entities.Result{}).WithContext(ctx).Create(&scanResult)
	return result.Error
}

func (r *ScanResultRepository) FindOne(ctx context.Context, condition map[string]interface{}) (entities.Result, error) {
	var scanResult entities.Result
	result := r.dbClient.Model(&entities.Result{}).Where(condition).First(&scanResult)
	if result.Error != nil {
		return scanResult, result.Error
	}
	return scanResult, nil
}

func (r *ScanResultRepository) FindMany(ctx context.Context, condition map[string]interface{}) ([]entities.Result, error) {
	var scanResults []entities.Result
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Find(&scanResults).Limit(20).Offset(0)
	if result.Error != nil {
		return scanResults, result.Error
	}
	return scanResults, nil
}

func (r *ScanResultRepository) Delete(ctx context.Context, condition map[string]interface{}) error {
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Delete(&entities.Result{})
	return result.Error
}

func (r *ScanResultRepository) Update(ctx context.Context, condition map[string]interface{}, newPayload entities.Result) error {
	result := r.dbClient.Model(&entities.Result{}).Where(condition).Updates(newPayload)
	return result.Error
}

func (r *ScanResultRepository) Count(ctx context.Context, condition map[string]interface{}) int64 {
	var count int64
	r.dbClient.Model(&entities.Result{}).Where(condition).Count(&count)
	return count
}
