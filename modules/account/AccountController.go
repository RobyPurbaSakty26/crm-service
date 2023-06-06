package account

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
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
	Message string                     `json:"message"`
	Data    AccountItemsLoginResposnse `json:"data"`
}

type AccountItemsLoginResposnse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_ID  uint   `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
	Token    string `json:"token"`
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

// login generate tokan and verify token
func GenerateToken(id, username, role, secret string) (string, error) {
	// ini sialisasi klaim
	claims := jwt.MapClaims{
		"sub":      id,
		"username": username,
		"role":     role,
	}

	// tandatangan token dengan kunci rahasia

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil

}

type payloadJWT struct {
	ID       string
	username string
	Role     string
}

func (c AccountControllers) verifyJWT(tokenString, secret string) (*payloadJWT, error) {
	// Memeriksa keaslian token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		fmt.Print(err)
		return nil, err
	}

	// Token valid, dapatkan informasi pengguna dari token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	userName := claims["username"].(string)
	role := claims["role"].(string)

	data := payloadJWT{

		ID:       userID,
		username: userName,
		Role:     role,
	}

	return &data, nil
}

func (c AccountControllers) login(req *loginRequest) (*responeLogin, error) {

	data, err := c.ReadByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = ComparePassword(data.Data.Password, req.Password)
	if err != nil {
		return nil, err
	}

	secret := "secret-key"

	token, err := GenerateToken(strconv.FormatUint(uint64(data.Data.ID), 10), data.Data.Username, strconv.FormatUint(uint64(data.Data.Role_ID), 10), secret)
	if err != nil {
		return nil, err
	}
	res := &responeLogin{
		Message: "Success",
		Data: AccountItemsLoginResposnse{
			ID:       data.Data.ID,
			Username: data.Data.Username,
			Password: data.Data.Password,
			Role_ID:  data.Data.Role_ID,
			Active:   data.Data.Active,
			Verified: data.Data.Verified,
			Token:    token,
		},
	}

	return res, nil
}
