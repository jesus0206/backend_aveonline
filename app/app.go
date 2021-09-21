package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jesus.tn79/aveonline/config"
	"jesus.tn79/aveonline/controllers"
)

var router *gin.Engine

// StartApp function runner server
func StartApp() {
	router := RouterInitial()
	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

// RouterInitial exported
func RouterInitial() *gin.Engine {
	dbSQL := config.SQLServer{
		URLBD:      os.Getenv("BD_URL"),
		NameBD:     os.Getenv("BD_NAME"),
		UserBD:     os.Getenv("BD_USER_NAME"),
		PasswordBD: os.Getenv("BD_USER_PASSWORD"),
		PortBD:     os.Getenv("BD_PORT"),
	}

	db, err := config.GetConnectionPostgres(&dbSQL)
	if err != nil {
		fmt.Println("Error en la conexion con la bd..")
		log.Fatal(err)
	}

	if os.Getenv("DEBUG") == "TRUE" {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	controller := controllers.NewController(db)

	router.GET("/medicamento", controller.GetMedicamentoController)
	router.POST("/medicamento", controller.CreateMedicamentoController)
	router.GET("/promocion", controller.GetPromocionController)
	router.POST("/promocion", controller.CreatePromocionController)
	router.GET("/factura", controller.GetFacturaController)
	return router
}
