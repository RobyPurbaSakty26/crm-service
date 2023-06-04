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

func (c CustomersRepository) Find() ([]Customers, error) {
	var customers []Customers
	err := c.db.Find(&customers).Error
	return customers, err

}

func (c CustomersRepository) FindById(id any) (Customers, error) {
	var customer Customers
	err := c.db.First(&customer, id).Error

	return customer, err
}
