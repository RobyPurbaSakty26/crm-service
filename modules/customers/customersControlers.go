package customers

type CustomersControllers struct {
	useCase *CustomersUsecase
}

func NewConstomersConstroller(useCase *CustomersUsecase) *CustomersControllers {
	return &CustomersControllers{
		useCase: useCase,
	}
}

type CustomersItemsResposnse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
}

type CreateResponse struct {
	Message string                  `json:"message"`
	Data    CustomersItemsResposnse `json:"data"`
}

func (c CustomersControllers) Create(req *CreateRequest) (*CreateResponse, error) {
	customers := Customers{
		Email:     req.Email,
		LastName:  req.LastName,
		Avatar:    req.Avatar,
		FirstName: req.FirstName,
	}
	err := c.useCase.Create(&customers)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: CustomersItemsResposnse{
			ID:        customers.ID,
			Email:     customers.Email,
			LastName:  customers.LastName,
			Avatar:    customers.Avatar,
			FirstName: customers.FirstName,
		},
	}
	return res, nil

}
