package models

import "time"

type Account struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Xpub      string     `json:"xpub" gorm:"not null; unique"`
	Gap       int        `json:"gap" gorm:"not null; default: 0"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	Addresses []Address  `json:"addresses,omitempty"`
}
