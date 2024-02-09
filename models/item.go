package models

import (
	"time"
)

type Item struct {
	ItemId      uint   `gorm:"PrimaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(191)"`
	Description string `gorm:"not null;type:varchar(191)"`
	Quantity    int    `gorm:"not null;type:int"`
	OrderId     uint   `gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
