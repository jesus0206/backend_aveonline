package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jesus.tn79/aveonline/model"
	// "jesus.tn79/ficha_electronica/models"
)

func (con Controller) GetPromocionController(c *gin.Context) {
	data, errr := con.repo.GetPromociones()
	if errr != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: errr.Error()})
		return
	}
	if len(data) == 0 {
		c.JSON(http.StatusNotFound, model.ResponseEmpty{})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (con Controller) CreatePromocionController(c *gin.Context) {
	var json model.Promocion
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "Debe enviar los datos como json.."})
		return
	}
	data, errr := con.repo.CreatePromocion(json)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: errr.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
