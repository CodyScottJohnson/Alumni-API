// Package models contains application specific entities.
package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Profile holds specific application settings linked to an Account.
type Address struct {
	ID        int
	Street    string
	City      string
	State     string
	Zip       string
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Address) BeforeInsert(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Address) BeforeUpdate(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Profile struct and returns validation errors.
func (p *Address) Validate() error {
	return nil
	//return validation.ValidateStruct(p,
	//validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
