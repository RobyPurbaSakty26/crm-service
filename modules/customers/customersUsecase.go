package customers

type CustomersUsecase struct {
	repo CustomerRepositoryInterface
}

type CustomerUsecaseInterface interface {
	Create(customer *Customers) error
	Read() ([]Customers, error)
	ReadByPk(id any) (Customers, error)
	Update(customer *Customers) error
	Delete(cutomer *Customers) error
	GetByEmail(email, firstName string) ([]Customers, error)
}

func NewUseCase(repo CustomerRepositoryInterface) CustomerUsecaseInterface {
	return CustomersUsecase{
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

func (u CustomersUsecase) GetByEmail(email, firstName string) ([]Customers, error) {
	return u.repo.FindByEmail(email, firstName)
}
