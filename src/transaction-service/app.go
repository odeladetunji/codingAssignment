package main

import (
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
)

func main(){

	router := gin.Default();
	router.Use(auth.CORSMiddleware());
	router.MaxMultipartMemory = 100 << 20  // 50 MiB

	setRoutes := func(){
		transerviceApi..Router(router);
	}

    setRoutes();
	
	if err := Endless.ListenAndServe("localhost:8091", router); err != nil {
		log.Fatal("Failed run app: ", err)
	}
}
