package repository

import (
	"avito/internal/dateMarshaller"
	"avito/internal/models"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (s *StatsRepository) Create(stats *models.Stats) error {
	stats.Cost = stats.Cost.Round(2)
	if stats.Clicks != 0 && stats.Views != 0 {
		stats.Cpc = stats.Cost.Div(decimal.NewFromInt(int64(stats.Clicks)))
		stats.Cpc = stats.Cpc.Round(2)
		stats.Cpc.StringFixedBank(2)
	}

	if stats.Views != 0 {
		stats.Cpm = stats.Cost.Mul(decimal.NewFromInt(1000)).Div(decimal.NewFromInt(int64(stats.Views)))
		stats.Cpm = stats.Cpm.Round(2)
		stats.Cpm.StringFixedBank(2)
	}
	stats.Cost.StringFixedBank(2)

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
	if err := s.checkOrder(order); err != nil {
		return nil, err
	}
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

func (s *StatsRepository) checkOrder(order string) error {
	switch order {
	case "date":
		return nil
	case "clicks":
		return nil
	case "views":
		return nil
	case "cost":
		return nil
	case "cost_per_mille":
		return nil
	case "cost_per_click":
		return nil
	default:
		return errors.New("invalid order")
	}
}
