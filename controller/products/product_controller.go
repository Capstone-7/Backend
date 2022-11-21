package products

import (
	"capstone/businesses/products"
	"capstone/controller/products/requests"
	response "capstone/controller/products/response"
	"capstone/helpers"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProductController struct {
	ProductUseCase products.UseCase
}

func NewProductController(productUseCase products.UseCase) *ProductController {
	return &ProductController{
		ProductUseCase: productUseCase,
	}
}

func (p *ProductController) GetAll(c echo.Context) error {
	product, err := p.ProductUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (p *ProductController) Create(c echo.Context) error {
	request := requests.Product{}
	c.Bind(&request)

	product, err := p.ProductUseCase.CreateProduct(request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    product,
	})
}

func (p *ProductController) GetProductByID(c echo.Context) error {
	product_id := c.Param("id")

	product, err := p.ProductUseCase.GetByID(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting product",
		Data:    response.FromDomain(&product),
	})
}
func (p *ProductController) UpdateProduct(c echo.Context) error {
	product_id := c.Param("id")

	request := requests.Product{}
	c.Bind(&request)

	domain := request.ToDomain()
	ObjID, _ := primitive.ObjectIDFromHex(product_id)
	domain.Id = ObjID

	product, err := p.ProductUseCase.UpdateProduct(domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success update product",
		Data:    product,
	})
}


func (p *ProductController) DeleteProduct(c echo.Context) error {
	product_id := c.Param("id")

	product, err := p.ProductUseCase.DeleteProduct(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success delete product",
		Data:    product,
	})
}