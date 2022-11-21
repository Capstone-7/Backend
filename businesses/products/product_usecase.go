package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductUseCase struct {
	ProductRepository Repository
}

func (p *ProductUseCase) UpdateProduct(domain *Domain) (Domain, error) {
	product, err := p.ProductRepository.GetByID(domain.Id)
	if err != nil {
		return Domain{}, err
	}
	product.Code = domain.Code
	product.Description = domain.Description
	product.Nominal = domain.Nominal
	product.Details = domain.Details
	product.Price = domain.Price
	product.Type = domain.Type
	product.ActivePeriod = domain.ActivePeriod
	product.Status = domain.Status
	product.IconUrl = domain.IconUrl
	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return p.ProductRepository.Update(&product)
}

func NewProductUseCase(repository Repository) UseCase {
	return &ProductUseCase{repository}
}

func (p *ProductUseCase) GetAll() ([]Domain, error) {
	return p.ProductRepository.GetAll()
}

func (p *ProductUseCase) GetByID(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return p.ProductRepository.GetByID(ObjId)
}

func (p *ProductUseCase) CreateProduct(domain *Domain) (Domain, error) {

	domain.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.DeletedAt = primitive.NewDateTimeFromTime(time.Time{})

	product, err := p.ProductRepository.Create(domain)
	return product, err
}

func (p *ProductUseCase) DeleteProduct(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return p.ProductRepository.Delete(ObjId)
}

