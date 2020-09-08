package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandanghosh/blog/config"

	"github.com/chandanghosh/blog/api/database/seeder"
	"github.com/chandanghosh/blog/api/router"
)

// Run ..
func Run() {
	config.Load()
	
	seeder.SeedDatabase()

	fmt.Printf("\t Listening [::]:%d\n", config.PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router.New()))
}
