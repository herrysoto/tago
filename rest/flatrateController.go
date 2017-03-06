package rest

import (
	datos "flatrate/datos"
	"github.com/gin-gonic/gin"
)


func Getfirstcombo(c *gin.Context){

	descod,err := datos.Getfirstcombo(c)
	if err == nil{
			c.JSON(200, descod)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}


func Getcodigovalidacion(c *gin.Context){	
	codigovalidacion,err := datos.Getcodigovalidacion(c)
	if err == nil{
			c.JSON(200, codigovalidacion)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func GetHorasHombres(c *gin.Context){

	hh,err := datos.GetHorasHombres(c)
	if err == nil{
		c.JSON(200,hh)
	}else{
		c.JSON(400,gin.H{"error":"no registros en la tabla"})
	}
}


func PostHorasHombres(c *gin.Context){
	inserthh,err := datos.PostHorasHombres(c)
	if err == nil{
			c.JSON(200, inserthh)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func PutHorasHombre(c *gin.Context){
	puthh,err := datos.PutHorasHombre(c)
	if err == nil{
			c.JSON(200, puthh)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func Getnumcodigo(c *gin.Context){
	numcod,err := datos.Getnumcodigo(c)
	if err == nil{
		c.JSON(200,numcod)
	}else{
		c.JSON(400,gin.H{"error":"no registros en la tabla"})
	}
}

func Getnumcodigoop(c *gin.Context){
	numcodop,err := datos.Getnumcodigoop(c)
	if err == nil{
		c.JSON(200,numcodop)
	}else{
		c.JSON(400,gin.H{"error":"no registros en la tabla"})
	}
}

//prueba para que me traiga el input por default
func Getfirstcombo2(c *gin.Context){

	codiserv,err := datos.Getfirstcombo2(c)
	if err == nil{
			c.JSON(200, codiserv)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

//    return lista de flatratees en json

func Getsecondcombo(c *gin.Context){

	codserv,err := datos.Getsecondcombo(c)
	if err == nil{
			c.JSON(200, codserv)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func Getgrid(c *gin.Context){
	opser,err := datos.Getgrid(c)
	if err == nil{
			c.JSON(200, opser)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}

}

func BuscarOperServiciosParam(c *gin.Context){
	opserparam,err := datos.BuscarOperServiciosParam(c)
	if err == nil{
			c.JSON(200, opserparam)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}

}

func LISTAR_OPERSERVCONTENIDOS(c *gin.Context){
	opserparam1,err := datos.LISTAR_OPERSERVCONTENIDOS(c)
	if err == nil{
			c.JSON(200, opserparam1)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}

}

func PutOperaServi(c *gin.Context){
	insop,err := datos.PutOperaServi(c)
	if err == nil{
			c.JSON(200, insop)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func UpdateGrilla(c *gin.Context){
	upgrill,err := datos.UpdateGrilla(c)
	if err == nil{
			c.JSON(200, upgrill)
		}else {
			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
		}
}

func Postgrid(c *gin.Context){
	insertOS,err := datos.Postgrid(c)
	if err== nil {
		c.JSON(200,insertOS)
	}else {
		c.JSON(404,gin.H{"error":"no registros(s) en la tabla"})
	}
}

// func GetNumMaestra(c *gin.Context){
// 	nummaestra,err := datos.GetNumMaestra(c)
// 	if err == nil{
// 			c.JSON(200, nummaestra)
// 		}else {
// 			c.JSON(404, gin.H{"error": "no registros(s) en la tabla"})
// 		}

// }


//otra forma de hacer el delete
// func DeleteContenido(c *gin.Context){
// 	nummaestraa,err := datos.DeleteContenido(c)
// 	if err== nil {
// 		c.JSON(200,nummaestraa)
// 	}else {
// 		c.JSON(404,gin.H{"error":"no registros(s) en la tabla"})
// 	}
// }
