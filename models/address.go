package models

import "time"

type Address struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Address   string     `json:"address" gorm:"not null; unique"`
	Account   *Account   `json:"account,omitempty"`
	AccountID uint       `json:"account_id" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	Transfers []Transfer `json:"transfers,omitempty"`
}
