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
	stats.Avg.Cpc = stats.Cost / stats.Clicks
	}

	if stats.Views !=0 {
		stats.Avg.Cpm = (stats.Cost * 1000) / stats.Views
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
//Joins("inner join avgs on avgs.stats_id = stats.id")
func (s *StatsRepository) GetStats(from, to dateMarshaller.CustomDate, order string) ([]models.Stats, error) {
	model := make([]models.Stats, 0)
	result := s.db.Model(model).Preload("stats").Order(order).Where("date BETWEEN ? and ?", from, to).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	//log.Println(model[0].Avg.Cpm)
	return model, nil
}
