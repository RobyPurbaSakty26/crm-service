package customers

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	ID        uint
	Email     string
	LastName  string
	Avatar    string
	FirstName string
	CreatedAt time.Time
}
