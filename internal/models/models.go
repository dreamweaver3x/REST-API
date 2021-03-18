package models

import (
	"avito/internal/dateMarshaller"
	"gorm.io/gorm"
)

type Stats struct {
	ID     uint                      `gorm:"primarykey"`
	Date   dateMarshaller.CustomDate `json:"date"`
	Views  uint                       `json:"views"`
	Clicks uint                       `json:"clicks"`
	Cost   uint                       `json:"cost"`
	Avg    Avg                       `gorm:"foreignKey:StatsID"`
}
type Avg struct {
	ID      uint `gorm:"primarykey"`
	StatsID uint
	Cpm     uint `gorm:"column:cost_per_mille"`
	Cpc     uint `gorm:"column:cost_per_click"`
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Stats{}, &Avg{})
	if err != nil {
		return err
	}

	return nil
}
