package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"jesus.tn79/aveonline/model"
	// "jesus.tn79/ficha_electronica/models"
)

func (con Controller) GetFacturaController(c *gin.Context) {
	var query model.FacturaDto
	query.Fecha_Inicio = c.DefaultQuery("fecha_inicio", "")
	if query.Fecha_Inicio == "" {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "no esta enviando la fecha de inicio."})
		return
	}
	query.Fecha_Fin = c.DefaultQuery("fecha_fin", "")
	if query.Fecha_Fin == "" {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "no esta enviando la fecha de fin."})
		return
	}
	_fecha_inicio, err := time.Parse("2006-01-02", query.Fecha_Inicio)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "la fecha de inicio no tiene el formato de una fecha año/mes/dia."})
		return
	}
	_fecha_fin, err := time.Parse("2006-01-02", query.Fecha_Fin)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "la fecha de fin no tiene el formato de una fecha año/mes/dia."})
		return
	}
	if _fecha_fin.Before(_fecha_inicio) {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "la fecha de inicio no puede ser mayor a la fecha fin."})
		return
	}
	data, errr := con.repo.GetFacturas(query)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: errr.Error()})
		return
	}
	if len(data) == 0 {
		c.JSON(http.StatusNotFound, model.ResponseEmpty{Data: "no se encontro la data"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// func (con Controller) CreatePromocionController(c *gin.Context) {
// 	var json model.Promocion
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "Debe enviar los datos como json.."})
// 		return
// 	}
// 	data, errr := con.repo.CreatePromocion(json)
// 	if errr != nil {
// 		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: errr.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, data)
// }
