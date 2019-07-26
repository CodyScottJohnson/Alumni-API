// Package models contains application specific entities.
package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Job holds info on job history.
type Job struct {
	ID         int
	ProfileID  int `json:"-"`
	Profile    *Profile
	CompanyID  int
	Company    *Company
	Title      string
	FunctionID int
	Function   *Function
	CurrentJob bool
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	EndedAt    time.Time `json:"ended_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Job) BeforeInsert(db orm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Job) BeforeUpdate(db orm.DB) error {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Profile struct and returns validation errors.
func (p *Job) Validate() error {
	return nil
	//return validation.ValidateStruct(p,
	//validation.Field(&p.Theme, validation.Required, validation.In("default", "dark")),
	//)
}
