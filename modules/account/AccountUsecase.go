package account

type AccountUsecase struct {
	repo *accountRepository
}

func NewUseCase(repo *accountRepository) *AccountUsecase {
	return &AccountUsecase{
		repo: repo,
	}
}

func (u AccountUsecase) create(a *Actor) error {
	return u.repo.save(a)
}

func (u AccountUsecase) getByUsername(username string) (Actor, error) {
	return u.repo.FindByUsername(username)
}

func (u AccountUsecase) Read() ([]Actor, error) {
	return u.repo.Find()
}

func (u AccountUsecase) Update(actor *Actor) error {
	return u.repo.Update(actor)
}

func (u AccountUsecase) ReadByPk(id any) (Actor, error) {
	return u.repo.FindById(id)
}
