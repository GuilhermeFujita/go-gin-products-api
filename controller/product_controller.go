package controller

import (
	"fmt"
	"go-api/dto"
	"go-api/model"
	"go-api/usecase"
	"go-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// Usecase vai aqui
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var productDTO dto.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = productDTO.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error validating product: %s", err))
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(productDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	productID, valid := utils.ValidateID(id)

	if !valid {
		response := model.Response{
			Message: "Id is not a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product.ID == 0 {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id := ctx.Param("productId")
	productID, valid := utils.ValidateID(id)

	if !valid {
		response := model.Response{
			Message: "Id is not a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err = p.productUseCase.UpdateProduct(productID, product); err != nil {
		if err.Error() == "Product not found" {
			var response model.Response
			response = model.Response{
				Message: "Product not found",
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := model.Response{
			Message: "Failed to update product",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.Status(http.StatusOK)
}
