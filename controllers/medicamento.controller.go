package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jesus.tn79/aveonline/model"
	// "jesus.tn79/ficha_electronica/models"
)

func (con Controller) GetMedicamentoController(c *gin.Context) {
	data, errr := con.repo.GetMedicamentos()
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

func (con Controller) CreateMedicamentoController(c *gin.Context) {

	var json model.Medicamento
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "Debe enviar los datos como json.."})
		return
	}
	data, errr := con.repo.CreateMedicamento(json)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: errr.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
