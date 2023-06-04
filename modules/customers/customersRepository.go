package customers

import "gorm.io/gorm"

type CustomersRepository struct {
	db *gorm.DB
}

func NewCustomersRepository(db *gorm.DB) *CustomersRepository {
	return &CustomersRepository{db: db}
}

func (c CustomersRepository) Save(customers *Customers) error {
	return c.db.Create(customers).Error
}
