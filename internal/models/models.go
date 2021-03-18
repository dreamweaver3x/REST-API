package models

import (
	"avito/internal/dateMarshaller"
	"gorm.io/gorm"
)

type Stats struct {
	ID     uint                      `gorm:"primarykey" json:"-"`
	Date   dateMarshaller.CustomDate `json:"date"`
	Views  uint                      `json:"views"`
	Clicks uint                      `json:"clicks"`
	Cost   uint                      `json:"cost"`
	Cpm    uint                      `gorm:"column:cost_per_mille" json:"cost_per_mille"`
	Cpc    uint                      `gorm:"column:cost_per_click" json:"cost_per_click"`
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Stats{})
	if err != nil {
		return err
	}

	return nil
}
