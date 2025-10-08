package repository

import (
	"fmt"
	"os"
	"products/src/pb/products"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct {
}

const fileName string = "./products.txt"

func (r *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return productList, nil
		}
		return productList, fmt.Errorf("failed to read products file: %+v", err)
	}
	if len(data) == 0 {
		return productList, nil
	}
	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("failed to unmarshal products file: %+v", err)
	}
	return productList, nil
}

func (r *ProductRepository) SaveData(productList products.ProductList) error {
	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("failed to marshal products file: %+v", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save products file: %+v", err)
	}
	return nil
}

func (r *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := r.loadData()
	if err != nil {
		return product, err
	}
	product.Id = int32(len(productList.Products) + 1)
	productList.Products = append(productList.Products, &product)
	err = r.SaveData(productList)
	return product, err
}

func (r *ProductRepository) FindAll() (products.ProductList, error) {
	return r.loadData()
}
