package repository

import (
	"avito/internal/dateMarshaller"
	"avito/internal/models"
	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (s *StatsRepository) Create(stats *models.Stats) error {

	if stats.Clicks !=0 && stats.Views !=0 {
	stats.Cpc = stats.Cost / stats.Clicks
	}

	if stats.Views !=0 {
		stats.Cpm = (stats.Cost * 1000) / stats.Views
	}

	tx := s.db.Begin()
	result := tx.Create(stats)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Commit()
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}

func (s *StatsRepository) GetStats(from, to dateMarshaller.CustomDate, order string) ([]models.Stats, error) {
	model := make([]models.Stats, 0)
	result := s.db.Order(order).Where("date BETWEEN ? and ?", from, to).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (s *StatsRepository) DeleteFromDB() error {
	result := s.db.Where("id > 0").Delete(&models.Stats{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
