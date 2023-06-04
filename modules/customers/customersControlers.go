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

type UpdateResponse struct {
	Message string                  `json:"message"`
	Data    CustomersItemsResposnse `json:"data"`
}

func (c CustomersControllers) Update(req *CreateRequest, id any) (*CreateResponse, error) {

	data, err := c.useCase.ReadByPk(id)

	if err != nil {
		return nil, err
	}
	customer := &Customers{
		ID:        data.ID,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Avatar:    data.Avatar,
	}

	if req.FirstName == "" {
		customer.FirstName = data.FirstName
	} else {
		customer.FirstName = req.FirstName
	}
	customer.LastName = req.LastName
	customer.Email = req.Email
	customer.Avatar = req.Avatar

	// Check if created_at field is set
	createdAt := data.CreatedAt
	customer.CreatedAt = createdAt

	// Check if the created_at value is set to a valid date
	// if createdAt.IsZero() || createdAt.Year() < 1 {
	// 	customer.CreatedAt = time.Now() // Assign the current time as created_at value
	// 	fmt.Println("diatas")
	// } else {
	// 	customer.CreatedAt = createdAt // Preserve the original created_at value
	// 	fmt.Println("dibawah")
	// }

	err = c.useCase.Update(customer)
	if err != nil {
		return nil, err
	}
	response := &CreateResponse{
		Message: "Success",
		Data: CustomersItemsResposnse{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return response, nil
}

type DeleteResponse struct {
	Message string
}

func (c CustomersControllers) Delete(id any) (*DeleteResponse, error) {
	data, err := c.useCase.ReadByPk(id)
	if err != nil {
		return nil, err
	}

	customer := &Customers{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Avatar:    data.Avatar,
	}

	if err := c.useCase.Delete(customer); err != nil {
		return nil, err
	}
	res := &DeleteResponse{
		Message: "Success",
	}
	return res, nil
}
