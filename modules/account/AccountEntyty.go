package account

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	ID       uint
	Username string
	Password string
	Role_ID  uint
	Verified string
	Active   string
}
