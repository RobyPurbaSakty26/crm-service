package account

type AccountUsecase struct {
	repo AccountRepositoryInterface
}

type UsecaseInterface interface {
	Create(a *Actor) error
	GetByUsername(username string) (Actor, error)
	Read() ([]Actor, error)
	Update(actor *Actor) error
	ReadByPk(id any) (Actor, error)
}

func NewUseCase(repo AccountRepositoryInterface) UsecaseInterface {
	return AccountUsecase{
		repo: repo,
	}
}

func (u AccountUsecase) Create(a *Actor) error {
	return u.repo.Save(a)
}

func (u AccountUsecase) GetByUsername(username string) (Actor, error) {
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
