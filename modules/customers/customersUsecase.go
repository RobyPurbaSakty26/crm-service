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
