package main

import (
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
	"log"
)

func main(){

	router := gin.Default();
	router.MaxMultipartMemory = 100 << 20  // 50 MiB

	setRoutes := func(){
	    var transerviceApi TranserviceApi;
		transerviceApi.Router(router);
	}

    setRoutes();
	
	if err := Endless.ListenAndServe("localhost:8091", router); err != nil {
		log.Fatal("Failed run app: ", err)
	}
}
