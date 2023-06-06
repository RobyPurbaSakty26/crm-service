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

	err := c.db.Where("username = ?", string(username)).First(&actor).Error
	if err != nil {
		return actor, err
	}

	return actor, nil
}
