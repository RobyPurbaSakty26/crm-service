package account

import (
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func newAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (r accountRepository) save(a *Actor) error {
	return r.db.Create(a).Error
}

func (c accountRepository) FindByUsername(username string) (Actor, error) {
	var actor Actor

	err := c.db.Where("username = ?", username).First(&actor).Error
	if err != nil {
		return actor, err
	}

	return actor, nil
}

// get all
func (c accountRepository) Find() ([]Actor, error) {
	var actor []Actor
	err := c.db.Find(&actor).Error
	return actor, err

}

// update
func (c accountRepository) Update(actor *Actor) error {
	return c.db.Save(actor).Error
}

// find by id
func (c accountRepository) FindById(id any) (Actor, error) {
	var actor Actor
	err := c.db.First(&actor, id).Error

	return actor, err
}
