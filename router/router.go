package router

import (
	"TM-Product/model"
	"TM-Product/service"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderApi(e *gin.Engine) {

	e.GET("/products", products)
	e.PUT("/productsdata", productsdata)
}

func products(c *gin.Context) {

	productInfo, err := service.OrderRepo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, productInfo)
}

func productsdata(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	var productDate model.ProductDate
	err = json.Unmarshal(body, &productDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = service.OrderRepo.InsertProductData(productDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, 1)
}
