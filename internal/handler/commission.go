package handler

import (
	"net/http"
	"trading-api/internal/domain/service"

	"github.com/gin-gonic/gin"
	uuid "github.com/tentone/mssql-uuid"
)

type CommissionHandler struct {
	commissionService service.CommissionService
}

type CommissionHandlerDependencies struct {
	CommissionService service.CommissionService
}

func NewCommissionHandler(d CommissionHandlerDependencies) CommissionHandler {
	return CommissionHandler{
		commissionService: d.CommissionService,
	}
}

func (h *CommissionHandler) GetCommissionByID(c *gin.Context) {
	ctx := c.Request.Context()
	commissionIdStr := c.Param("id")
	commissionId, err := uuid.FromString(commissionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid commission id format",
			"details": err.Error(),
		})
		return
	}

	commission, err := h.commissionService.GetCommissionByID(ctx, commissionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": commission,
	})
}

func (h *CommissionHandler) GetListCommission(c *gin.Context) {
	ctx := c.Request.Context()
	commissions, err := h.commissionService.GetListCommission(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": commissions,
	})
}
