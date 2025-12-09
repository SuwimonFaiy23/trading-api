package handler

import (
	"net/http"
	"trading-api/internal/domain/service"
	"trading-api/internal/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/tentone/mssql-uuid"
)

type OrderHandler struct {
	orderService service.OrderService
}

type OrderHandlerDependencies struct {
	OrderService service.OrderService
}

func NewOrderHandler(d OrderHandlerDependencies) OrderHandler {
	return OrderHandler{
		orderService: d.OrderService,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.OrderRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": order,
	})
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	ctx := c.Request.Context()
	orderIdStr := c.Param("id")
	orderId, err := uuid.FromString(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid order id format",
			"details": err.Error(),
		})
		return
	}

	order, err := h.orderService.GetOrderByID(ctx, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": order,
	})
}
