package internal

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null;unique"`
	Status    string     `json:"status" gorm:"not null"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

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
	return nil
}
