package customers

type CustomersUsecase struct {
	repo *CustomersRepository
}

func NewUseCase(repo *CustomersRepository) *CustomersUsecase {
	return &CustomersUsecase{
		repo: repo,
	}
}

func (u CustomersUsecase) Create(customer *Customers) error {
	return u.repo.Save(customer)
}

func (u CustomersUsecase) Read() ([]Customers, error) {
	return u.repo.Find()
}

func (u CustomersUsecase) ReadByPk(id any) (Customers, error) {
	return u.repo.FindById(id)
}
