package customers

import (
	"errors"
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
	ctrl CustomerControllersInterface
}

type CustomersHandlerInterface interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadByPk(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewRequestHandler(ctrl CustomerControllersInterface) CustomersHandlerInterface {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) CustomersHandlerInterface {
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

func (h RequestHandler) Read(c *gin.Context) {

	email := c.Query("email")
	firstName := c.Query("first_name")
	if email != "" || firstName != "" {
		res, err := h.ctrl.getByEmail(email, firstName)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}

	res, err := h.ctrl.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h RequestHandler) ReadByPk(c *gin.Context) {
	CustomerID := c.Param("id")
	res, err := h.ctrl.ReadByPk(CustomerID)

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h RequestHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Update(&req, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	res, err := h.ctrl.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// func (h RequestHandler) GetByEmail(c *gin.Context) {
// 	email := c.Query("email")

// 	res, err := h.ctrl.getByEmail(email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }
