package internal

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusDeleted  = "deleted"
)

type CustomerUpdate struct {
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Status *string `json:"status"`
}

func (u *CustomerUpdate) TableName() string {
	return "customers"
}

func (u *CustomerUpdate) Validate() error {
	if u.Name != nil && *u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if u.Email != nil && *u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if u.Status != nil && *u.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	if u.Status != nil && *u.Status != StatusActive && *u.Status != StatusInactive && *u.Status != StatusDeleted {
		return fmt.Errorf("status must be one of: %s, %s, %s", StatusActive, StatusInactive, StatusDeleted)
	}
	return nil
}

type Customer struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null;unique"`
	Status    string     `json:"status" gorm:"not null"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (c *Customer) TableName() string {
	return "customers"
}

//func (c *CustomerUpdate) BeforeUpdate(tx *gorm.DB) error {
//	if c.updatedAt == nil {
//		now := time.Now()
//		c.updatedAt = &now
//	}
//	return nil
//}

func (c *Customer) BeforeCreate(*gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID, _ = uuid.NewV7()
	}
	if c.CreatedAt == nil {
		now := time.Now()
		c.CreatedAt = &now
	}
	if c.UpdatedAt == nil {
		now := time.Now()
		c.UpdatedAt = &now
	}
	return nil
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	if c.Email == "" {
		return fmt.Errorf("email is required")
	}
	if c.Status == "" {
		return fmt.Errorf("status is required")
	}
	if c.Status != StatusActive && c.Status != StatusInactive && c.Status != StatusDeleted {
		return fmt.Errorf("status must be one of: %s, %s, %s", StatusActive, StatusInactive, StatusDeleted)
	}
	return nil
}
