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

type ReadResponse struct {
	Data []CustomersItemsResposnse `json:"data"`
}
type ReadByIdResponse struct {
	Data CustomersItemsResposnse `json:"data"`
}

func (c CustomersControllers) Read() (*ReadResponse, error) {
	customers, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, customer := range customers {
		res.Data = append(res.Data, CustomersItemsResposnse{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Avatar,
			Avatar:    customer.Avatar,
		})
	}
	return res, nil
}

func (c CustomersControllers) ReadByPk(id any) (*ReadByIdResponse, error) {
	customer, err := c.useCase.ReadByPk(id)
	if err != nil {
		return nil, err
	}

	res := &ReadByIdResponse{
		Data: CustomersItemsResposnse{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}
