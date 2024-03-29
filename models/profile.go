// Package models contains application specific entities.
package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Profile holds specific application settings linked to an Account.
type Profile struct {
	ID           int `json:"-"`
	AccountID    int `json:"-"`
	FirstName    string
	LastName     string
	Username     string
	ProfilePhoto string
	AddressID    int
	Address      *Address
	Email        string
	Phone        string
	Jobs         []*Job
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Profile) BeforeInsert(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Profile) BeforeUpdate(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Profile struct and returns validation errors.
func (p *Profile) Validate() error {
	return nil
	//return validation.ValidateStruct(p,
	//validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
