package test

import (
	account "crm-service/modules/account"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAccountRepository is a mock structure for AccountRepositoryInterface
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Save(a *account.Actor) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByUsername(username string) (account.Actor, error) {
	args := m.Called(username)
	return args.Get(0).(account.Actor), args.Error(1)
}

func (m *MockAccountRepository) Find() ([]account.Actor, error) {
	args := m.Called()
	return args.Get(0).([]account.Actor), args.Error(1)
}

func (m *MockAccountRepository) Update(actor *account.Actor) error {
	args := m.Called(actor)
	return args.Error(0)
}

func (m *MockAccountRepository) FindById(id interface{}) (account.Actor, error) {
	args := m.Called(id)
	return args.Get(0).(account.Actor), args.Error(1)
}

func TestAccountUsecase_Create(t *testing.T) {
	// Setup
	repo := new(MockAccountRepository)
	usecase := account.NewUseCase(repo)

	actor := &account.Actor{

		ID:       0,
		Username: "Roby",
		Password: "1234",
		Role_ID:  0,
		Verified: "true",
		Active:   "true",
	}

	repo.On("Save", actor).Return(nil)

	// Testing
	err := usecase.Create(actor)

	// Assertion
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAccountUsecase_GetByUsername(t *testing.T) {
	// Setup
	repo := new(MockAccountRepository)
	usecase := account.NewUseCase(repo)

	username := "test_username"

	expectedActor := account.Actor{
		ID:       0,
		Username: username,
		Password: "123",
		Role_ID:  0,
		Verified: "true",
		Active:   "false",
	}

	repo.On("FindByUsername", username).Return(expectedActor, nil)

	// Testing
	actor, err := usecase.GetByUsername(username)

	// Assertion
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
	repo.AssertExpectations(t)
}

func TestAccountUsecase_Read(t *testing.T) {
	// Setup
	repo := new(MockAccountRepository)
	usecase := account.NewUseCase(repo)

	expectedActors := []account.Actor{
		// Initialize expected actors
	}

	repo.On("Find").Return(expectedActors, nil)

	// Testing
	actors, err := usecase.Read()

	// Assertion
	assert.NoError(t, err)
	assert.Equal(t, expectedActors, actors)
	repo.AssertExpectations(t)
}

func TestAccountUsecase_Update(t *testing.T) {
	// Setup
	repo := new(MockAccountRepository)
	usecase := account.NewUseCase(repo)

	actor := &account.Actor{
		ID:       0,
		Username: "Joe",
		Password: "123",
		Role_ID:  0,
		Verified: "true",
		Active:   "true",
	}

	repo.On("Update", actor).Return(nil)

	// Testing
	err := usecase.Update(actor)

	// Assertion
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAccountUsecase_ReadByPk(t *testing.T) {
	// Setup
	repo := new(MockAccountRepository)
	usecase := account.NewUseCase(repo)

	id := "test_id"

	expectedActor := account.Actor{
		ID:       1,
		Username: "Joe",
		Password: "123",
		Role_ID:  2,
		Verified: "true",
		Active:   "true",
	}

	repo.On("FindById", id).Return(expectedActor, nil)

	// Testing
	actor, err := usecase.ReadByPk(id)

	// Assertion
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
}
