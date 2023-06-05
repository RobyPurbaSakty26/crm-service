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
		Verified: req.Verified,
		Active:   req.Verified,
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
