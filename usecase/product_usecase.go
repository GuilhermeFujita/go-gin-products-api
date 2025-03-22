package usecase

import (
	"errors"
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

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID
	return product, nil
}

func (pu *ProductUsecase) GetProduct(productID int) (model.Product, error) {
	return pu.repository.GetProduct(productID)
}

func (pu *ProductUsecase) UpdateProduct(productID int, data model.Product) error {
	dbProduct, err := pu.repository.GetProduct(productID)
	if err != nil {
		return err
	}

	if dbProduct.ID == 0 {
		return errors.New("Product not found")
	}

	productToUpdate := model.Product{
		ID:    dbProduct.ID,
		Name:  data.Name,
		Price: data.Price,
	}

	err = pu.repository.UpdateProduct(productToUpdate)
	if err != nil {
		return err
	}

	return nil
}
