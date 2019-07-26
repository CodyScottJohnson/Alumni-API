// Package models contains application specific entities.
package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Company holds info on companies
type Company struct {
	ID          int
	Name        string
	Logo        string
	LinkedinURL string
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Company) BeforeInsert(db orm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Company) BeforeUpdate(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Profile struct and returns validation errors.
func (p *Company) Validate() error {
	return nil
	//return validation.ValidateStruct(p,
	//validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
