package account

import (
	"errors"
	"net/http"
	"strings"

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

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h RequestHandler) ReadByUsername(c *gin.Context) {

	username := c.Query("username")

	if username != "" {
		res, err := h.ctrl.ReadByUsername(username)
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}

	res, err := h.ctrl.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h RequestHandler) Login(c *gin.Context) {
	var req loginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h RequestHandler) AuthMiddleware(c *gin.Context) {
	// get token from authorization
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// verify token
	token, err := h.ctrl.verifyJWT(tokenString, "secret-key")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, ErrorResponse{err.Error()})
		c.Abort()
		return
	}

	c.Set("id", token.ID)
	c.Set("username", token.username)
	c.Set("role", token.Role)

	c.Next()

}
