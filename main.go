package main

import (
	controller "flatrate/rest"
	datos "flatrate/datos"
	"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
)
// Cors funcion para activar CORS
func Cors() gin.HandlerFunc {

	//http://stackoverflow.com/questions/31834408/golang-gin-and-go-socket-io-cors-issue
	return func(c *gin.Context) {
		fmt.Println("CORS middleware loaded...")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, WS, WSS, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(Cors())
    v1 := r.Group("api/v1")
	{
		v1.GET("/firstcombo", controller.Getfirstcombo)  //java
		v1.GET("/Getcodigovalidacion/:Codopera/:Codoperaservicio",controller.Getcodigovalidacion)
		v1.GET("/HorasHombres",controller.GetHorasHombres) //java
		v1.PUT("/PutHorasHombre", controller.PutHorasHombre)
		v1.GET("/Getnumcodigo",controller.Getnumcodigo)  //java
		v1.POST("/PostHorasHombres", controller.PostHorasHombres)
		v1.GET("/firstcombo2", controller.Getfirstcombo2) 
		v1.GET("/secondcomb/:CodServicio", controller.Getsecondcombo) //java
		v1.GET("/grid/:Operserv", controller.Getgrid) //java
		v1.GET("/BuscarOperServiciosParam/:param", controller.BuscarOperServiciosParam)
		// v1.GET("/BuscarOperServiciosParam", controller.BuscarOperServiciosParam)
		v1.GET("/LISTAROPERSERVCONTENIDOS/:param/:param1", controller.LISTAR_OPERSERVCONTENIDOS)
		v1.POST("/post", controller.Postgrid)
		v1.PUT("/PutOperaServi", controller.PutOperaServi)
		v1.GET("/Getnumcodigoop", controller.Getnumcodigoop)
		v1.DELETE("/DeleteContenido/:nummaestra", datos.DeleteContenido)
		v1.PUT("/UpdateGrilla", controller.UpdateGrilla)
		v1.GET("/Getnumcodigooperacionmaestra", controller.Getnumcodigooperacionmaestra)
		v1.POST("/Nuevocontenido", controller.Nuevocontenido)
		// v1.GET("/GetNumMaestra/:operacionmaestra/:codigooperacion/:codoperserv", controller.GetNumMaestra)
		// v1.PUT("/put/:putOpeSer/:putOpeSer1",controller.PutOperaServi)

		// otra forma de hacer el Delete
		//v1.DELETE("/DeleteContenido", controller.DeleteContenido)
	}

	r.Run(":8083")
}