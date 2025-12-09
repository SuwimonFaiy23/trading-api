package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"trading-api/internal/domain/service"
	"trading-api/internal/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/tentone/mssql-uuid"
)

type UserHandler struct {
	userService service.UserService
}

type UserHandlerDependencies struct {
	UserService service.UserService
}

func NewUserHandler(d UserHandlerDependencies) UserHandler {
	return UserHandler{
		userService: d.UserService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateUsername(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UserRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateUsername(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userIdStr := c.Param("id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id format",
			"details": err.Error(),
		})
		return
	}

	user, err := h.userService.UpdateUser(ctx, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	userIdStr := c.Param("id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id format",
			"details": err.Error(),
		})
		return
	}

	user, err := h.userService.GetUserByID(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (h *UserHandler) AddBalanceUser(c *gin.Context) {
	var req dto.BalanceRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdStr := c.Param("id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id format",
			"details": err.Error(),
		})
		return
	}

	user, err := h.userService.AddBalanceUser(ctx, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (h *UserHandler) DeductBalanceUser(c *gin.Context) {
	var req dto.BalanceRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdStr := c.Param("id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id format",
			"details": err.Error(),
		})
		return
	}

	user, err := h.userService.DeductBalanceUser(ctx, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	result, err := h.userService.GetUserAllByPagination(c.Request.Context(), limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9._]{6,20}$`)

var usernameBlacklist = map[string]bool{
	"admin":   true,
	"root":    true,
	"system":  true,
	"owner":   true,
	"support": true,
}

func ValidateUsername(username string) error {
	username = strings.ToLower(username)

	if !usernameRegex.MatchString(username) {
		return errors.New("invalid username format")
	}

	if usernameBlacklist[username] {
		return errors.New("username is not allowed")
	}

	if strings.HasPrefix(username, ".") || strings.HasPrefix(username, "_") {
		return errors.New("username cannot start with special characters")
	}

	return nil
}
