package account

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AccountControllers struct {
	useCase *AccountUsecase
}

func NewAccountConstroller(useCase *AccountUsecase) *AccountControllers {
	return &AccountControllers{
		useCase: useCase,
	}
}

type AccountItemsResposnse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_ID  uint   `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

type CreateResponse struct {
	Message string                `json:"message"`
	Data    AccountItemsResposnse `json:"data"`
}

func EncryptPassword(password string) (string, error) {
	// Mengenkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c AccountControllers) create(req *CreateRequest) (*CreateResponse, error) {

	pwd, err := EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println(pwd)

	account := Actor{
		Username: req.Username,
		Password: pwd,
		Role_ID:  req.Role_ID,
		Active:   "false",
		Verified: "false",
	}
	err = c.useCase.create(&account)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: AccountItemsResposnse{
			ID:       account.ID,
			Username: account.Username,
			Password: account.Password,
			Role_ID:  account.Role_ID,
			Verified: account.Verified,
			Active:   account.Verified,
		},
	}
	return res, nil

}

func ComparePassword(hashedPassword string, password string) error {
	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

type responeLogin struct {
	Message string                `json:"message"`
	Data    AccountItemsResposnse `json:"data"`
}

type readByUsernameResponse struct {
	Message string                `json:"message"`
	Data    AccountItemsResposnse `json:"data"`
}

func (c AccountControllers) ReadByUsername(username string) (*readByUsernameResponse, error) {

	account, err := c.useCase.getByUsername(username)
	if err != nil {
		return nil, err
	}

	res := &readByUsernameResponse{
		Data: AccountItemsResposnse{
			ID:       account.ID,
			Username: account.Username,
			Password: account.Password,
			Role_ID:  account.Role_ID,
			Active:   account.Active,
			Verified: account.Verified,
		},
	}
	return res, nil
}

func (c AccountControllers) login(req *loginRequest) (*responeLogin, error) {

	data, err := c.ReadByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	// fmt.Println("DATA ", data)
	fmt.Println("USERNAME ", req.Username)
	fmt.Println("USERNAME ", req.Password)
	fmt.Println("USERNAME ", data.Data.Password)

	err = ComparePassword(data.Data.Password, req.Password)
	if err != nil {
		return nil, err
	}

	res := &responeLogin{
		Message: "Success",
		Data: AccountItemsResposnse{
			ID:       data.Data.ID,
			Username: data.Data.Username,
			Password: data.Data.Password,
			Role_ID:  data.Data.Role_ID,
			Active:   data.Data.Active,
			Verified: data.Data.Verified,
		},
	}

	return res, nil
}
