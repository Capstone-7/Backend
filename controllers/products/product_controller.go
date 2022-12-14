package products

import (
	"capstone/businesses/products"
	"capstone/controllers/products/requests"
	response "capstone/controllers/products/response"
	"capstone/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

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
		Message: "Success create product",
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

	// Validate
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	domain := request.ToDomain()
	ObjID, _ := primitive.ObjectIDFromHex(product_id)
	domain.ID = ObjID

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

// Get category list
func (p *ProductController) GetCategoryList(c echo.Context) error {
	category, err := p.ProductUseCase.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting category list",
		Data:    category,
	})
}

// Get categories by product type
func (p *ProductController) GetCategoriesByProductType(c echo.Context) error {
	product_type := c.Param("product_type")

	category, err := p.ProductUseCase.GetCategoriesByProductType(product_type)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if len(category) == 0 {
		return c.JSON(http.StatusNotFound, helpers.Response{
			Status:  http.StatusNotFound,
			Message: "Category not found",
			Data:    category,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting category list",
		Data:    category,
	})
}

// Get product by category
func (p *ProductController) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")

	product, err := p.ProductUseCase.GetProductsByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting product by category",
		Data:    product,
	})
}

// Get product by product type
func (p *ProductController) GetProductsByProductType(c echo.Context) error {
	product_type := c.Param("product_type")

	product, err := p.ProductUseCase.GetProductsByProductType(product_type)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting product by product type",
		Data:    product,
	})
}

// Get Total Product
func (p *ProductController) GetTotalProducts(c echo.Context) error {
	total, err := p.ProductUseCase.GetTotalProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting total product",
		Data:    total,
	})
}