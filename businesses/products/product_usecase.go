package products

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase struct {
	ProductRepository Repository
}

func NewProductUseCase(repository Repository) UseCase {
	return &ProductUseCase{repository}
}

func (p *ProductUseCase) UpdateProduct(domain *Domain) (Domain, error) {
	product, err := p.ProductRepository.GetByID(domain.ID)
	if err != nil {
		return Domain{}, err
	}

	// Check if product code already exists
	if product.Code != domain.Code {
		_, err := p.ProductRepository.GetOneByCode(domain.Code, true)
		if err == nil {
			return Domain{}, errors.New("Product code already exists")
		}
	}

	// Shadow delete old product
	product.Deleted = primitive.NewDateTimeFromTime(time.Now())
	_, err = p.ProductRepository.Update(&product)
	if err != nil {
		return Domain{}, err
	}

	// Create new product
	domain.ID = primitive.NewObjectID()
	domain.Created = product.Created
	domain.Updated = primitive.NewDateTimeFromTime(time.Now())
	domain.Deleted = primitive.NewDateTimeFromTime(time.Time{})
	newProduct, err := p.ProductRepository.Create(domain)
	return newProduct, err
}

func (p *ProductUseCase) GetAll() ([]Domain, error) {
	return p.ProductRepository.GetAll(true)
}

func (p *ProductUseCase) GetByID(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return p.ProductRepository.GetByID(ObjId)
}

func (p *ProductUseCase) CreateProduct(domain *Domain) (Domain, error) {

	// Check if product code already exists
	_, err := p.ProductRepository.GetOneByCode(domain.Code, true)
	if err == nil {
		return Domain{}, errors.New("Product code already exists")
	}else{
		fmt.Println("err", err)
	}

	domain.Created = primitive.NewDateTimeFromTime(time.Now())
	domain.Updated = primitive.NewDateTimeFromTime(time.Now())
	domain.Deleted = primitive.NewDateTimeFromTime(time.Time{})

	product, err := p.ProductRepository.Create(domain)
	return product, err
}

func (p *ProductUseCase) DeleteProduct(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return p.ProductRepository.Delete(ObjId)
}

func (p *ProductUseCase) GetCategories() ([]string, error) {
	return p.ProductRepository.GetCategories()
}

func (p *ProductUseCase) GetCategoriesByProductType(productType string) ([]string, error) {
	return p.ProductRepository.GetCategoriesByType(productType)
}

func (p *ProductUseCase) GetProductsByCategory(category string) ([]Domain, error) {
	return p.ProductRepository.GetManyByCategory(category, true)
}

func (p *ProductUseCase) GetProductsByProductType(productType string) ([]Domain, error) {
	return p.ProductRepository.GetManyByType(productType, true)
}

func (p *ProductUseCase) GetTotalProducts() (int64, error) {
	return p.ProductRepository.CountProducts()
}

func (p *ProductUseCase) GetProductByCode(code string) (Domain, error) {
	return p.ProductRepository.GetOneByCode(code, true)
}