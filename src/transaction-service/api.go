package main  

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	// "fmt"
)

type TranserviceApi struct {

}

var transerviceApi TranserviceApi;
func (transapi *TranserviceApi) Router(router *gin.Engine){
	var route *gin.RouterGroup = router.Group("api/transactions");
    transerviceApi.CreateTransaction(route);
	transerviceApi.GetAllCustomerTransactions(route);
}

func (transapi *TranserviceApi) CreateTransaction(router *gin.RouterGroup){
	router.POST("/createTransaction", func(c *gin.Context) {

		err := transactionService.CreateTransaction(c);
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
			"data": "",
			"status": http.StatusOK,
		});
		
	});
}

func (transapi *TranserviceApi) GetAllCustomerTransactions(router *gin.RouterGroup){
	router.GET("/transactions", func(c *gin.Context) {

		transaction, err := transactionService.GetAllCustomerTransactions(c);
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
			"data": transaction,
			"status": http.StatusOK,
		});
		
	});
}
