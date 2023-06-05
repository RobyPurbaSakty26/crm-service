package account

import "gorm.io/gorm"

type accountRepository struct {
	db *gorm.DB
}

func newAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

func (r accountRepository) save(a *Actor) error {
	return r.db.Create(a).Error
}
