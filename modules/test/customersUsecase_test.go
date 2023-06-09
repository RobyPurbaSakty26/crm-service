package test

import (
	"crm-service/modules/customers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Membuat mock struct untuk CustomerRepositoryInterface
type mockCustomerRepository struct {
	mock.Mock
}

// Implementasikan metode-metode yang dibutuhkan oleh CustomersUsecase
func (m *mockCustomerRepository) Save(customer *customers.Customers) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *mockCustomerRepository) Find() ([]customers.Customers, error) {
	args := m.Called()
	return args.Get(0).([]customers.Customers), args.Error(1)
}

func (m *mockCustomerRepository) FindById(id any) (customers.Customers, error) {
	args := m.Called(id)
	return args.Get(0).(customers.Customers), args.Error(1)
}

func (m *mockCustomerRepository) Update(customer *customers.Customers) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *mockCustomerRepository) Delete(customer *customers.Customers) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *mockCustomerRepository) FindByEmail(email, firstName string) ([]customers.Customers, error) {
	args := m.Called(email, firstName)
	return args.Get(0).([]customers.Customers), args.Error(1)
}

func TestCreate(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := customers.NewUseCase(repo)

	customer := &customers.Customers{
		ID:        1,
		Email:     "Joe@mail.com",
		FirstName: "Joe",
		LastName:  "Doe",
		Avatar:    "Joe.img",
	}

	repo.On("Save", customer).Return(nil)

	err := usecase.Create(customer)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestRead(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := customers.NewUseCase(repo)

	expectedCustomers := []customers.Customers{
		{
			ID:        1,
			Email:     "Joe@mail.com",
			FirstName: "Joe",
			LastName:  "Doe",
			Avatar:    "Joe.img",
		},
		{
			ID:        2,
			Email:     "Doe@mail.com",
			FirstName: "Doe",
			LastName:  "Joe",
			Avatar:    "Doe.img",
		},
	}

	repo.On("Find").Return(expectedCustomers, nil)

	result, err := usecase.Read()

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, result)
	repo.AssertExpectations(t)
}

// func TestReadByPk(t *testing.T) {
// 	repo := new(mockCustomerRepository)
// 	usecase := customers.NewUseCase(repo)

// 	expectedCustomer := customers.Customers{} // buat objek customer sesuai kebutuhan

// 	repo.On("FindById", any).Return(expectedCustomer, nil)

// 	result, err := usecase.ReadByPk(any)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedCustomer, result)
// 	repo.AssertExpectations(t)
// }

func TestUpdate(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := customers.NewUseCase(repo)

	customer := &customers.Customers{
		ID:        1,
		Email:     "Doe@mail.com",
		FirstName: "Doe",
		LastName:  "Joe",
		Avatar:    "Doe.img",
	}

	repo.On("Update", customer).Return(nil)

	err := usecase.Update(customer)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := customers.NewUseCase(repo)

	customer := &customers.Customers{

		ID:        2,
		Email:     "Doe@mail.com",
		FirstName: "Doe",
		LastName:  "Joe",
		Avatar:    "Doe.img",
	}

	repo.On("Delete", customer).Return(nil)

	err := usecase.Delete(customer)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetByEmail(t *testing.T) {
	repo := new(mockCustomerRepository)
	usecase := customers.NewUseCase(repo)

	email := "test@example.com"
	firstName := "John"

	expectedCustomers := []customers.Customers{
		{
			ID:        1,
			Email:     "stive@mail.com",
			FirstName: "stive",
			LastName:  "Joe",
			Avatar:    "stive.img",
		},
		{
			ID:        2,
			Email:     "Doe@mail.com",
			FirstName: "Doe",
			LastName:  "Joe",
			Avatar:    "Doe.img",
		},
	}

	repo.On("FindByEmail", email, firstName).Return(expectedCustomers, nil)

	result, err := usecase.GetByEmail(email, firstName)

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, result)
	repo.AssertExpectations(t)
}
