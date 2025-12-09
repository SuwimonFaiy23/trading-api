package handler

import (
	"net/http"
	"trading-api/internal/domain/service"
	"trading-api/internal/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/tentone/mssql-uuid"
)

type AffiliateHandler struct {
	affiliateService service.AffiliateService
}

type AffiliateHandlerDependencies struct {
	AffiliateService service.AffiliateService
}

func NewAffiliateHandler(d AffiliateHandlerDependencies) AffiliateHandler {
	return AffiliateHandler{
		affiliateService: d.AffiliateService,
	}
}

func (h *AffiliateHandler) CreateAffiliate(c *gin.Context) {
	var req dto.AffiliateRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	affiliate, err := h.affiliateService.CreateAffiliate(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": affiliate,
	})
}

func (h *AffiliateHandler) GetAffiliateByID(c *gin.Context) {
	ctx := c.Request.Context()
	affiliateIdStr := c.Param("id")
	affiliateId, err := uuid.FromString(affiliateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid affiliate id format",
			"details": err.Error(),
		})
		return
	}

	affiliate, err := h.affiliateService.GetAffiliateByID(ctx, affiliateId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": affiliate,
	})
}

func (h *AffiliateHandler) GetListAffiliate(c *gin.Context) {
	ctx := c.Request.Context()

	affiliate, err := h.affiliateService.GetListAffiliate(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": affiliate,
	})
}