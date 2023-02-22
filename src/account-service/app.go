package main

import (
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
	"log"
	Migration "services.com/migration"
)

func main(){

	router := gin.Default();
	// router.Use(auth.CORSMiddleware());
	router.MaxMultipartMemory = 100 << 20  // 50 MiB

	migrateDatabase :=  func(){
		var migration Migration.Migration = &Migration.MigrationService{}
		migration.MigrateTables();
	}

	migrateDatabase();

	setRoutes := func(){
		var acountServiceApi AcountServiceApi;
		acountServiceApi.Router(router);
	}

    setRoutes();
	
	if err := Endless.ListenAndServe("localhost:8090", router); err != nil {
		log.Fatal("Failed run app: ", err)
	}
}




