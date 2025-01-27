package services

import (
	_ "github.com/lib/pq"
)

type Record struct {
	SiteID                uint   `gorm:"not null"`
	FixletID              int64  `gorm:"primaryKey"` // Changed to int64
	Name                  string `gorm:"size:255;not null"`
	Criticality           string `gorm:"size:50;not null"`
	RelevantComputerCount int    `gorm:"not null"`
}
