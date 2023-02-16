package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

type AcountServiceApi struct {

}

var acountServiceApi AcountServiceApi;
func (acctapi *AcountServiceApi) Router(router *gin.Engine){
	var route *gin.RouterGroup = router.Group("/api/account-service");
    acountServiceApi.GetCustomers(route);
	acountServiceApi.CreateAccount(route);
	accountService.CreateCustomers();
}

func (acctapi *AcountServiceApi) GetCustomers(router *gin.Engine){
	route.GET("/customers/all", func(c *gin.Context) {

		data, err := accountService.GetCustomers();
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}

func (acctapi *AcountServiceApi) CreateAccount(router *gin.Engine){
	route.POST("/create", func(c *gin.Context) {

		data, err := accountService.CreateAccount(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}





