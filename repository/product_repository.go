package repository

import (
	"database/sql"
	"fmt"
	"go-api/dto"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product dto.ProductDTO) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var resultId int
	err = query.QueryRow(product.Name, product.Price).Scan(&resultId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return resultId, nil
}

func (pr *ProductRepository) GetProduct(productID int) (model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	var dbProduct model.Product
	err = query.QueryRow(productID).Scan(&dbProduct.ID, &dbProduct.Name, &dbProduct.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		return model.Product{}, err
	}

	query.Close()
	return dbProduct, nil

}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}
