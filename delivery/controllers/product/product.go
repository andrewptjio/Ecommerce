package product

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	product "ecommerce/repository/products"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Repo product.ProductInterface
}

func NewProductControllers(pi product.ProductInterface) *ProductController {
	return &ProductController{Repo: pi}
}

func (pc ProductController) GetProductCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		product, err := pc.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := ProductResponseFormat{
			Message: "Successful Operation",
			Data:    product,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (pc ProductController) CreateProductControllers() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		newProductreq := ProductRequestFormat{}

		if err := c.Bind(&newProductreq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newProduct := entities.Product{
			Name:        newProductreq.Name,
			Price:       newProductreq.Price,
			Stock:       newProductreq.Stock,
			CategoryID:  newProductreq.CategoryID,
			Description: newProductreq.Description,
		}
		res, err := pc.Repo.Create(newProduct)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		return c.JSON(http.StatusOK, ProductResponseFormat{
			Message: "Successfull Operation",
			Data:    res,
		})
	}
}

func (pc ProductController) UpdateProductCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateProductReq := ProductRequestFormat{}
		if err := c.Bind(&updateProductReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateProduct := entities.Product{
			Name:        updateProductReq.Name,
			Price:       updateProductReq.Price,
			Stock:       updateProductReq.Stock,
			CategoryID:  updateProductReq.CategoryID,
			Description: updateProductReq.Description,
		}

		if _, err := pc.Repo.Update(updateProduct, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
func (pc ProductController) DeleteProductCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if _, err := pc.Repo.Delete(id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
func (pc ProductController) Pagination() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		category := c.QueryParam("category")
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		page, _ := strconv.Atoi(c.QueryParam("page"))
		sort := c.QueryParam("sort")
		pagination := entities.Pagination{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		}
		product, _ := pc.Repo.Pagination(name, category, pagination)

		if product.Error != nil {
			return c.JSON(http.StatusBadRequest, PaginationResponseFormat{
				Message: "Bad Request",
				Error:   product.Error,
			})
		}

		var data = product.Result

		// urlPath := c.Request().RequestURI
		// urlPath = urlPath[:11]

		return c.JSON(http.StatusOK, PaginationResponseFormat{
			Message: "Successful Operation",
			Data:    data,
		})
	}
}
