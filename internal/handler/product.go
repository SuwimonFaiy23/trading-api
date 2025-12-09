package handler

import (
	"net/http"
	"trading-api/internal/domain/service"
	"trading-api/internal/dto"

	"github.com/gin-gonic/gin"
	uuid "github.com/tentone/mssql-uuid"
)

type ProductHandler struct {
	productService service.ProductService
}

type ProductHandlerDependencies struct {
	ProductService service.ProductService
}

func NewProductHandler(d ProductHandlerDependencies) ProductHandler {
	return ProductHandler{
		productService: d.ProductService,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.ProductRequest
	ctx := c.Request.Context()
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.CreateProduct(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": product,
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	ctx := c.Request.Context()

	productIdStr := c.Param("id")
	productId, err := uuid.FromString(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid product id format",
			"details": err.Error(),
		})
		return
	}

	product, err := h.productService.GetProductByID(ctx, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": product,
	})
}

func (h *ProductHandler) GetListProduct(c *gin.Context) {
	ctx := c.Request.Context()

	product, err := h.productService.GetListProduct(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": product,
	})
}
