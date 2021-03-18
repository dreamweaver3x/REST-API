package models

import (
	"avito/internal/dateMarshaller"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Stats struct {
	ID     uint                      `gorm:"primarykey" json:"-"`
	Date   dateMarshaller.CustomDate `json:"date"`
	Views  uint                      `json:"views"`
	Clicks uint                      `json:"clicks"`
	Cost   decimal.Decimal           `json:"cost" sql:"type:decimal(20,8);"`
	Cpm    decimal.Decimal           `gorm:"column:cost_per_mille" json:"cost_per_mille" sql:"type:decimal(20,8);"`
	Cpc    decimal.Decimal           `gorm:"column:cost_per_click" json:"cost_per_click" sql:"type:decimal(20,8);"`
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Stats{})
	if err != nil {
		return err
	}

	return nil
}
