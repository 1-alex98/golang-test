package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique;notNull"`
	Password      string `gorm:"notNull"`
	Credit        float64
	AccountEntrys []AccountEntry
}

type Good struct {
	gorm.Model
	Name          string `gorm:"unique;notNull"`
	Description   string
	CurrentCourse float64
	DataPoints    []DataPoint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DataPoint struct {
	gorm.Model
	Value  float64 `gorm:"notNull"`
	GoodID uint
}

type AccountEntry struct {
	gorm.Model
	Value  float64 `gorm:"notNull"`
	UserID uint
	GoodID uint
	Good   Good
}
