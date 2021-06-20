package models

import (
	"time"
)

type Transfer struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	InvoiceID   string     `json:"invoice_id" gorm:"not null; unique"`
	Beneficiary string     `json:"beneficiary" gorm:"not null"`
	TxIn        string     `json:"tx_in,omitempty" gorm:"not null"`
	TxOut       string     `json:"tx_out,omitempty" gorm:"not null"`
	ValueIn     string     `json:"value_in,omitempty" gorm:"not null"`
	ValueOut    string     `json:"value_out,omitempty" gorm:"not_null"`
	Rate        string     `json:"rate,omitempty" gorm:"not null"`
	Address     *Address   `json:"address,omitempty"`
	AddressID   uint       `json:"address_id" gorm:"not null"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
