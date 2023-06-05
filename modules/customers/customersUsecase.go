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

func (u CustomersUsecase) Update(customer *Customers) error {
	return u.repo.Update(customer)
}

func (u CustomersUsecase) Delete(cutomer *Customers) error {
	return u.repo.Delete(cutomer)
}

func (u CustomersUsecase) getByEmail(email string) ([]Customers, error) {
	return u.repo.FindByEmail(email)
}
