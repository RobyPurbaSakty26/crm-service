package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_ID  uint   `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

type RequestHandler struct {
	ctrl *AccountControllers
}

func NewRequestHandler(ctrl *AccountControllers) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewAccountConstroller(
			NewUseCase(
				newAccountRepository(db),
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

	res, err := h.ctrl.create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
