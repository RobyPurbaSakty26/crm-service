package customers

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Email     string
	LastName  string
	Avatar    string
	FirstName string
}
