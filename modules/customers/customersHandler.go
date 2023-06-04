package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
}

type RequestHandler struct {
	ctrl *CustomersControllers
}

func NewRequestHandler(ctrl *CustomersControllers) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewConstomersConstroller(
			NewUseCase(
				NewCustomersRepository(db),
			),
		),
	)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
