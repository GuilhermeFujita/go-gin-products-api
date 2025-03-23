package usecase

import (
	"errors"
	"go-api/dto"
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	// Repository vai aqui
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product dto.ProductDTO) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    productID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}

func (pu *ProductUsecase) GetProduct(productID int) (model.Product, error) {
	return pu.repository.GetProduct(productID)
}

func (pu *ProductUsecase) UpdateProduct(productID int, productDTO dto.ProductDTO) error {
	dbProduct, err := pu.repository.GetProduct(productID)
	if err != nil {
		return err
	}

	if dbProduct.ID == 0 {
		return errors.New("Product not found")
	}

	productToUpdate := model.Product{
		ID:    dbProduct.ID,
		Name:  productDTO.Name,
		Price: productDTO.Price,
	}

	err = pu.repository.UpdateProduct(productToUpdate)
	if err != nil {
		return err
	}

	return nil
}
